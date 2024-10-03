package config

import (
	"log"
)

type TelemetryConfig struct {
	DBConfig
	ServerConfig
}

func NewTelemetryConfig() *TelemetryConfig {
	var cfg TelemetryConfig
	if err := parsePrefix("INSTRUMENTATION_", &cfg); err != nil {
		log.Fatal(err.Error())
	}
	return &cfg
}
