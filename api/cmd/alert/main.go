package main

import (
	"log"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/config"
	"github.com/USACE/instrumentation-api/api/internal/handlers"
	"github.com/USACE/instrumentation-api/api/internal/utils"
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
	db := utils.Connection(cfg.DBConfig.ConnStr())
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatal(err.Error())
		}
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
