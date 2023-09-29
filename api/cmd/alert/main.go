package main

import (
	"log"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/config"
	"github.com/USACE/instrumentation-api/api/internal/handler"
	"github.com/USACE/instrumentation-api/api/internal/util"
)

func checkAlerts(h *handler.AlertCheckHandler) {
	if err := h.DoAlertChecks(); err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("successfully completed alert checks at %s", time.Now())
}

func main() {
	cfg := config.NewAlertCheckConfig()
	db := util.Connection(cfg.DBConfig.ConnStr())
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatal(err.Error())
		}
	}()

	h := handler.NewAlertCheck(cfg)

	if cfg.AWSECSTriggerMocked {
		for {
			checkAlerts(h)
			time.Sleep(15 * time.Second)
		}
	} else {
		checkAlerts(h)
	}
}
