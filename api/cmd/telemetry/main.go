package main

import (
	"log"
	"net/http"

	"github.com/USACE/instrumentation-api/api/internal/config"
	"github.com/USACE/instrumentation-api/api/internal/handlers"
	"github.com/USACE/instrumentation-api/api/internal/middleware"
	"github.com/USACE/instrumentation-api/api/internal/models"
	"github.com/USACE/instrumentation-api/api/internal/utils"
	"github.com/apex/gateway"

	"github.com/labstack/echo/v4"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	cfg := config.GetTelemetryConfig()
	db := utils.Connection(cfg.DBConfig.ConnStr())

	e := echo.New()
	e.Use(middleware.CORS, middleware.GZIP)

	hashExtractor := func(model, sn string) (string, error) {
		hash, err := models.GetDataLoggerHashByModelSN(db, model, sn)
		if err != nil {
			return "", err
		}
		return hash, nil
	}

	public := e.Group(cfg.RoutePrefix)
	public.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{"status": "healthy"})
	})

	datalogger := e.Group(cfg.RoutePrefix)
	datalogger.Use(middleware.DataLoggerKeyAuth(hashExtractor))

	datalogger.POST("/telemetry/datalogger/:model/:sn", handlers.CreateOrUpdateDataLoggerMeasurements(db))

	if cfg.LambdaContext {
		log.Print("starting server; Running On AWS LAMBDA")
		if err := gateway.ListenAndServe("localhost:3030", e); err != nil {
			log.Fatal(err.Error())
		}
	} else {
		log.Print("starting server")
		if err := http.ListenAndServe(":80", e); err != nil {
			log.Fatal(err.Error())
		}
	}
}
