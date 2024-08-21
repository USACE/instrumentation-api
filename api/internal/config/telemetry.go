package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type TelemetryConfig struct {
	DBConfig
	ServerConfig
	EmailConfig
}

func NewTelemetryConfig() *TelemetryConfig {
	var cfg TelemetryConfig
	if err := envconfig.Process("instrumentation", &cfg); err != nil {
		log.Fatal(err.Error())
	}
	return &cfg
}
