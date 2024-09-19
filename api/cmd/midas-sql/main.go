package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/USACE/instrumentation-api/api/internal/config"
	"github.com/USACE/instrumentation-api/api/internal/migrate"
	"github.com/USACE/instrumentation-api/api/migrations"
)

func main() {
	mcfg := config.NewMigrateConfig()

	seedLocal := strings.Contains(mcfg.FileLocations, "filesystem:/flyway/sql/local")
	dbConfig := migrateDBConfig(mcfg)

	var cmd string

	if len(os.Args) > 2 {
		log.Fatal("too many arguments")
	}

	if len(os.Args) > 1 {
		switch strings.ToLower(os.Args[1]) {
		case "migrate":
			cmd = "migrate"
		case "init":
			cmd = "init"
		default:
			log.Fatalf("invalid argument %s", os.Args[1])
		}
	}

	migrator := migrate.NewMigrationService(&migrate.Config{
		Init:          cmd == "init",
		DBConfig:      dbConfig,
		SeedLocal:     seedLocal,
		MigrationsDir: migrations.MigrationsDir,
	})

	migrator.Run(context.Background())
}

func migrateDBConfig(cfg *config.MigrateConfig) migrate.DBConfig {
	dbUrl := strings.TrimPrefix(cfg.DBUrl, "jdbc:postgresql://")

	dd := strings.Split(dbUrl, "/")

	if len(dd) == 0 {
		log.Fatal("invalid database url")
	}

	dbHost := dd[0]
	if len(dbHost) == 0 {
		log.Fatal("invalid database url (hostname)")
	}

	dbAddr := strings.Split(dbHost, ":")
	dbHost = dbAddr[0]

	dbPort := 5432
	if len(dbAddr) == 2 {
		var err error
		dbPort, err = strconv.Atoi(dbAddr[1])
		if err != nil {
			log.Fatal("invalid database url (port)")
		}
	}

	dbName := dd[len(dd)-1]
	if len(dbHost) == 0 {
		log.Fatal("invalid database url (database name)")
	}

	return migrate.DBConfig{
		DBUser:          cfg.DBUser,
		DBPass:          cfg.DBPassword,
		DBName:          dbName,
		DBHost:          dbHost,
		DBPort:          dbPort,
		DBSSLMode:       "disable",
		DatabaseSchemas: []string{"midas", "public"},
	}
}
