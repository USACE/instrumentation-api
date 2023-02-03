package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/USACE/instrumentation-api/api/dbutils"
	"github.com/USACE/instrumentation-api/api/handlers"
	"github.com/USACE/instrumentation-api/api/middleware"
	"github.com/USACE/instrumentation-api/api/models"
	"github.com/apex/gateway"
	"github.com/kelseyhightower/envconfig"

	"github.com/labstack/echo/v4"

	_ "github.com/jackc/pgx/v4/stdlib"
)

// Config stores configuration information stored in environment variables
type Config struct {
	AuthDisabled   bool   `envconfig:"AUTH_DISABLED"`
	AuthJWTMocked  bool   `envconfig:"AUTH_JWT_MOCKED"`
	ApplicationKey string `envconfig:"APPLICATION_KEY"`
	LambdaContext  bool
	DBUser         string
	DBPass         string
	DBName         string
	DBHost         string
	DBSSLMode      string
	HeartbeatKey   string
	RoutePrefix    string `envconfig:"ROUTE_PREFIX"`
}

func (c *Config) dbConnStr() string {
	return fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=%s", c.DBUser, c.DBPass, c.DBName, c.DBHost, c.DBSSLMode)
}

func main() {

	var cfg Config
	if err := envconfig.Process("instrumentation", &cfg); err != nil {
		log.Fatal(err.Error())
	}

	db := dbutils.Connection(cfg.dbConnStr())

	hashExtractor := func(keyID string) (string, error) {
		k, err := models.GetTokenInfoByTokenID(db, &keyID)
		if err != nil {
			return "", err
		}
		return k.Hash, nil
	}

	e := echo.New()
	e.Use(middleware.CORS, middleware.GZIP)

	// Datalogger Telemetry
	datalogger := e.Group(cfg.RoutePrefix)
	datalogger.Use(middleware.KeyAuth(
		cfg.AuthDisabled,
		cfg.ApplicationKey,
		hashExtractor,
	))
	datalogger.POST("/telemetry/measurements", handlers.CreateOrUpdateDataLoggerMeasurements(db))

	if cfg.LambdaContext {
		log.Print("starting server; Running On AWS LAMBDA")
		log.Fatal(gateway.ListenAndServe("localhost:3030", e))
	} else {
		log.Print("starting server")
		log.Fatal(http.ListenAndServe(":80", e))
	}
}
