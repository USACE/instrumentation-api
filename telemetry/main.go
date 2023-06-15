package main

import (
	"log"
	"net/http"

	"github.com/USACE/instrumentation-api/api/config"
	"github.com/USACE/instrumentation-api/api/dbutils"
	"github.com/USACE/instrumentation-api/api/handlers"
	"github.com/USACE/instrumentation-api/api/middleware"
	"github.com/USACE/instrumentation-api/api/models"
	"github.com/apex/gateway"

	"github.com/labstack/echo/v4"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	// Environment Variable Config
	cfg := config.GetTelemetryConfig()
	db := dbutils.Connection(config.DBConnStr(&cfg.DBConfig))

	e := echo.New()
	e.Use(middleware.CORS, middleware.GZIP)

	hashExtractor := func(model, sn string) (string, error) {
		hash, err := models.GetDataLoggerHashByModelSN(db, model, sn)
		if err != nil {
			return "", err
		}
		return hash, nil
	}

	// Healthcheck
	public := e.Group(cfg.RoutePrefix)
	public.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{"status": "healthy"})
	})

	// Datalogger Telemetry
	datalogger := e.Group(cfg.RoutePrefix)
	datalogger.Use(middleware.DataLoggerKeyAuth(hashExtractor))

	datalogger.POST("/telemetry/datalogger/:model/:sn", handlers.CreateOrUpdateDataLoggerMeasurements(db))

	if cfg.LambdaContext {
		log.Print("starting server; Running On AWS LAMBDA")
		log.Fatal(gateway.ListenAndServe("localhost:3030", e))
	} else {
		log.Print("starting server")
		log.Fatal(http.ListenAndServe(":80", e))
	}
}
