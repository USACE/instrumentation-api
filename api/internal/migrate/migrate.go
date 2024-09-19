package migrate

import (
	"bytes"
	"context"
	"crypto/sha1"
	"database/sql"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"slices"
	"sort"
	"strings"
	"time"

	"github.com/hashicorp/go-version"
	_ "github.com/jackc/pgx/v5/stdlib"
)

const databaseDriver = "pgx"

type migrationService struct {
	db  *sql.DB
	cfg *Config
}

type DBConfig struct {
	DBUser          string
	DBPass          string
	DBName          string
	DBHost          string
	DBPort          int
	DBSSLMode       string
	DatabaseSchemas []string
}

func (cfg *DBConfig) ConnStr() string {
	s := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=%s", cfg.DBUser, cfg.DBPass, cfg.DBName, cfg.DBHost, cfg.DBPort, cfg.DBSSLMode)
	if len(cfg.DatabaseSchemas) != 0 && cfg.DatabaseSchemas[0] != "" {
		s += fmt.Sprintf(" search_path=%s", strings.Join(cfg.DatabaseSchemas, ","))
	}
	return s
}

type Config struct {
	Init          bool
	SeedLocal     bool
	MigrationsDir fs.FS
	DBConfig
}

func NewMigrationService(cfg *Config) *migrationService {
	ctx := context.Background()
	db, err := sql.Open(databaseDriver, cfg.ConnStr())
	if err != nil {
		log.Fatalf("unable to connect to database: %v\n", err.Error())
	}
	if err := db.PingContext(ctx); err != nil {
		log.Fatal(err.Error())
	}

	if len(cfg.DatabaseSchemas) == 0 {
		cfg.DatabaseSchemas = append(cfg.DatabaseSchemas, "public")
	}

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer func() {
		if err := tx.Rollback(); err != nil && !errors.Is(err, sql.ErrTxDone) {
			log.Fatal(err.Error())
		}
	}()

	for _, schema := range cfg.DatabaseSchemas {
		if _, err := tx.ExecContext(ctx, "CREATE SCHEMA IF NOT EXISTS "+quoteIdentifier(schema)); err != nil {
			log.Fatalf("could not create schema %s; %s", schema, err.Error())
		}
	}

	if _, err := tx.ExecContext(ctx, `
		CREATE TABLE IF NOT EXISTS schema_migration_history (
			installed_rank integer PRIMARY KEY,
			version text UNIQUE NOT NULL,
			filename text NOT NULL,
			checksum bytea NOT NULL,
			installed_by text NOT NULL,
			installed_on timestamptz NOT NULL DEFAULT now(),
			execution_time_ms integer NOT NULL
		)
	`); err != nil {
		log.Fatalf("could not create table schema_migration_history; %s", err.Error())
	}

	if err := tx.Commit(); err != nil {
		log.Fatal(err.Error())
	}

	return &migrationService{db, cfg}
}

type MigrationFile struct {
	File   fs.DirEntry
	Prefix string
}

type SchemaMigrationHistory struct {
	InstalledRank   int
	Version         Version
	Filename        string
	Checksum        []byte
	InstalledBy     string
	InstalledOn     time.Time
	ExecutionTimeMs int
}

type Version struct {
	*version.Version
}

func (s migrationService) Run(ctx context.Context) {
	start := time.Now()

	if err := s.migrate(ctx); err != nil {
		log.Println()
		log.Print("ERROR ENCOUNTERED; ROLLED BACK MIGRATIONS")
		log.Fatal(err.Error())
	}

	end := time.Now()
	log.Println()
	log.Printf("SUCCESSFULLY COMPLETED DATABASE MIGRATIONS IN %s", end.Sub(start))
}

func (s migrationService) migrate(ctx context.Context) error {
	defer func() {
		if err := s.db.Close(); err != nil {
			log.Fatal(err.Error())
		}
	}()

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if err := tx.Rollback(); err != nil && !errors.Is(err, sql.ErrTxDone) {
			log.Fatalf(err.Error())
		}
	}()

	history, err := s.querySchemaHistory(ctx)
	if err != nil {
		return err
	}

	sh := make(map[string]SchemaMigrationHistory)
	for _, s := range history {
		sh[s.Filename] = s
	}

	schemas, err := fs.ReadDir(s.cfg.MigrationsDir, "schema")
	if err != nil {
		return err
	}

	if len(schemas) == 0 {
		return errors.New("schemas directory must not be empty")
	}
	schemaMigrations := make([]MigrationFile, 0)
	for _, f := range schemas {
		schemaMigrations = append(schemaMigrations, MigrationFile{File: f, Prefix: "schema"})
	}

	if s.cfg.SeedLocal {
		seed, err := fs.ReadDir(s.cfg.MigrationsDir, "seed")
		if err != nil {
			return err
		}
		for _, f := range seed {
			schemaMigrations = append(schemaMigrations, MigrationFile{File: f, Prefix: "seed"})
		}
	}

	hasher := sha1.New()
	migrationsNew := make([]SchemaMigrationHistory, 0)
	var mostRecent *SchemaMigrationHistory
	if len(history) > 0 {
		mostRecent = &history[len(history)-1]
	}

	slices.SortFunc(schemaMigrations, sortFileBySemver)

	log.Printf("VALIDATING %d EXISTING VERSIONED MIGRATIONS...", len(history))
	log.Println()
	for _, m := range schemaMigrations {
		n := m.Prefix + "/" + m.File.Name()

		v, err := extractVersion(n)
		if err != nil {
			return err
		}
		if err != nil {
			return fmt.Errorf("invalid version prefix: '%s'", n)
		}

		m, exists := sh[n]
		if !exists {
			if mostRecent != nil && v.LessThan(mostRecent.Version.Version) {
				return fmt.Errorf(
					"migration %s applied out of order: most recent applied version is %s",
					n, mostRecent.Version,
				)
			}
			mNew := SchemaMigrationHistory{
				Version:  Version{v},
				Filename: n,
			}
			migrationsNew = append(migrationsNew, mNew)
			continue
		}

		// check hash of all files up to version in FS against hash of most recent version in schema_migration_history
		// if they don't match, exit and tell the user they can't edit existing migrations
		content, err := fs.ReadFile(s.cfg.MigrationsDir, n)
		if err != nil {
			return err
		}
		if _, err := hasher.Write(content); err != nil {
			return err
		}
		checksum := hasher.Sum(nil)

		if !bytes.Equal(checksum, m.Checksum) {
			return fmt.Errorf("checksums for %s did not match:\n\twant: %x\n\thave: %x", n, m.Checksum, checksum)
		}
		log.Printf("migration %s validated", n)
	}

	installedRank := len(history) + 1

	stmt, err := tx.PrepareContext(ctx, `
		INSERT INTO schema_migration_history
		(installed_rank, version, filename, checksum, execution_time_ms, installed_by)
		VALUES ($1, $2, $3, $4, $5, CURRENT_USER)
	`)
	if err != nil {
		return err
	}

	log.Println()
	log.Printf("APPLYING %d NEW VERSIONED MIGRATIONS...", len(migrationsNew))
	log.Println()
	for _, m := range migrationsNew {
		content, err := fs.ReadFile(s.cfg.MigrationsDir, m.Filename)
		if err != nil {
			return err
		}
		startExec := time.Now()
		if !s.cfg.Init {
			if _, err := tx.ExecContext(ctx, string(content)); err != nil {
				return err
			}
		}
		endExec := time.Now()

		if _, err := hasher.Write(content); err != nil {
			return err
		}

		m.InstalledRank = installedRank
		m.Checksum = hasher.Sum(nil)

		if _, err := stmt.ExecContext(
			ctx,
			m.InstalledRank, m.Version, m.Filename, m.Checksum, endExec.Sub(startExec).Milliseconds(),
		); err != nil {
			return err
		}
		installedRank++
		log.Printf("migration %s applied", m.Filename)
	}

	repeat, err := fs.ReadDir(s.cfg.MigrationsDir, "repeat")
	if err != nil {
		return err
	}

	sort.Slice(repeat, func(i, j int) bool {
		return repeat[i].Name() < repeat[j].Name()
	})

	log.Println()
	log.Printf("APPLYING %d REPEAT MIGRATIONS...", len(repeat))
	log.Println()
	for _, f := range repeat {
		n := "repeat/" + f.Name()
		content, err := fs.ReadFile(s.cfg.MigrationsDir, n)
		if err != nil {
			return err
		}
		if _, err := tx.ExecContext(ctx, string(content)); err != nil {
			return err
		}
		log.Printf("migration %s applied", n)
	}

	return tx.Commit()
}

func (s migrationService) querySchemaHistory(ctx context.Context) ([]SchemaMigrationHistory, error) {
	rows, err := s.db.QueryContext(ctx, `
		SELECT installed_rank, version, filename, checksum, installed_by, installed_on, execution_time_ms
		FROM schema_migration_history
		ORDER BY filename ASC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var history []SchemaMigrationHistory
	for rows.Next() {
		var h SchemaMigrationHistory
		if err := rows.Scan(&h.InstalledRank, &h.Version.Version, &h.Filename, &h.Checksum, &h.InstalledBy, &h.InstalledOn, &h.ExecutionTimeMs); err != nil {
			return nil, err
		}
		history = append(history, h)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return history, nil
}

func sortFileBySemver(a, b MigrationFile) int {
	if a.File == nil || b.File == nil {
		log.Fatal("internal error: nil migration")
	}
	va, err := extractVersion(a.File.Name())
	if err != nil {
		log.Fatal("invalid version for %s", a.File.Name())
	}
	vb, err := extractVersion(b.File.Name())
	if err != nil {
		log.Fatal("invalid version for %s", b.File.Name())
	}

	o := va.Compare(vb)
	if o == 0 {
		log.Fatalf("duplicate versions for %s", vb.String())
	}
	return o
}

func extractVersion(s string) (*version.Version, error) {
	n, err := extractPrefix(s, "V", "__")
	if err != nil {
		return nil, err
	}

	v, err := version.NewVersion(n)
	if err != nil {
		return nil, fmt.Errorf("file %s could not be parsed: error %s; skipping...", s, err.Error())
	}
	return v, nil
}

func extractPrefix(s, beginDelim, endDelim string) (string, error) {
	start := strings.Index(s, beginDelim)
	if start == -1 {
		return "", fmt.Errorf("starting sequence '%s' not found", beginDelim)
	}
	start++ // Move index to the character after beginDelim

	end := strings.Index(s, endDelim)
	if end == -1 || end <= start {
		return "", fmt.Errorf("ending sequence '%s' not found or in wrong position", endDelim)
	}

	return s[start:end], nil
}

func quoteIdentifier(name string) string {
	end := strings.IndexRune(name, 0)
	if end > -1 {
		name = name[:end]
	}
	return `"` + strings.Replace(name, `"`, `""`, -1) + `"`
}
