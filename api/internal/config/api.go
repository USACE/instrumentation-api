package config

import (
	"log"
	"os"
)

// Config stores configuration information stored in environment variables
type ApiConfig struct {
	BuildTag string
	DBConfig
	AWSS3Config
	AWSSQSConfig
	ServerConfig
}

// GetConfig returns environment variable config
func NewApiConfig() *ApiConfig {
	var cfg ApiConfig
	cfg.BuildTag = os.Getenv("BUILD_TAG")
	if err := parsePrefix("INSTRUMENTATION_", &cfg); err != nil {
		log.Fatal(err.Error())
	}

	switch cfg.BuildTag {
	case "local":
		cfg.ServerBaseUrl = "http://localhost:8080"
	case "dev":
		cfg.ServerBaseUrl = "https://develop-midas-api.rsgis.dev"
	case "test":
		cfg.ServerBaseUrl = "https://midas-test.cwbi.us/api"
	case "prod":
		cfg.ServerBaseUrl = "https://midas.sec.usace.army.mil/api"
	default:
	}

	return &cfg
}
