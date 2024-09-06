package migrate

import (
	"database/sql"
	"embed"
	"log"

	"github.com/USACE/instrumentation-api/api/internal/config"
	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/pressly/goose/v3"
)

//go:embed schema/*.sql
var schemaMigrations embed.FS

//go:embed repeat/*.sql
var repeatMigrations embed.FS

//go:embed seed/*.sql
var seedMigrations embed.FS

type migrationService struct {
	db *sql.DB
}

func NewMigrationService(cfg *config.DBConfig) *migrationService {
	db := model.NewDatabase(cfg)
	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatal(err.Error())
	}
	return &migrationService{db.DB.DB}
}

func (m *migrationService) RunSchemaMigrations() error {
	goose.SetBaseFS(schemaMigrations)
	return goose.Up(m.db, "schema")
}

func (m *migrationService) RunRepeatMigrations() error {
	goose.SetBaseFS(repeatMigrations)
	return goose.Up(m.db, "repeat", goose.WithNoVersioning())
}

func (m *migrationService) RunSeedMigrations() error {
	goose.SetBaseFS(seedMigrations)
	return goose.Up(m.db, "seed", goose.WithNoVersioning())
}

func (m *migrationService) RerunSeedMigrations() error {
	goose.SetBaseFS(seedMigrations)
	if err := goose.Down(m.db, "seed", goose.WithNoVersioning()); err != nil {
		return err
	}
	return goose.Up(m.db, "seed", goose.WithNoVersioning())
}
