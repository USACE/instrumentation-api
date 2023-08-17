package main

import (
	"log"
	"time"

	"github.com/USACE/instrumentation-api/api/config"
	"github.com/USACE/instrumentation-api/api/dbutils"
	"github.com/USACE/instrumentation-api/api/handlers"
	"github.com/jmoiron/sqlx"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func checkAlerts(db *sqlx.DB, cfg *config.AlertCheckConfig) {
	smtpCfg := config.GetSmtpConfig(cfg)

	if err := handlers.DoAlertChecks(db, cfg, smtpCfg); err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("successfully completed alert checks at %s", time.Now())
}

func main() {
	cfg := config.GetAlertCheckConfig()
	db := dbutils.Connection(config.DBConn())
	defer func() error {
		if err := db.Close(); err != nil {
			log.Fatal(err.Error())
		}
		return nil
	}()

	if cfg.AWSECSTriggerMocked {
		for {
			checkAlerts(db, cfg)
			time.Sleep(15 * time.Second)
		}
	} else {
		checkAlerts(db, cfg)
	}
}
