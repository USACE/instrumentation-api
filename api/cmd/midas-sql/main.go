package main

import (
	"context"
	"log"
	"strconv"
	"strings"

	"github.com/USACE/instrumentation-api/api/internal/config"
	"github.com/USACE/instrumentation-api/api/internal/migrate"
	"github.com/USACE/instrumentation-api/api/migrations"
)

func main() {
	mcfg := config.NewMigrateConfig()

	seedLocal := strings.Contains(mcfg.FileLocations, "filesystem:/flyway/sql/local")
	dbUrl := strings.TrimPrefix(mcfg.DBUrl, "jdbc:postgresql://")

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

	migrator := migrate.NewMigrationService(&migrate.Config{
		Init:      false,
		SeedLocal: seedLocal,
		DBConfig: migrate.DBConfig{
			DBUser:          mcfg.DBUser,
			DBPass:          mcfg.DBPassword,
			DBName:          dbName,
			DBHost:          dbHost,
			DBPort:          dbPort,
			DBSSLMode:       "disable",
			DatabaseSchemas: []string{"midas", "public"},
		},
		MigrationsDir: migrations.MigrationsDir,
	})

	migrator.Run(context.Background())
}
