package main

import (
	"context"
	"log"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/config"
	"github.com/USACE/instrumentation-api/api/internal/handler"
)

func main() {
	cfg := config.NewAlertCheckConfig()
	h := handler.NewAlertCheck(cfg)

	if cfg.TriggerMocked {
		for {
			checkAlerts(h)
			time.Sleep(15 * time.Second)
		}
	} else {
		checkAlerts(h)
	}
}

func checkAlerts(h *handler.AlertCheckHandler) {
	if err := h.AlertCheckService.DoAlertSchedulerChecks(context.Background()); err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("successfully completed alert checks at %s", time.Now())
}
