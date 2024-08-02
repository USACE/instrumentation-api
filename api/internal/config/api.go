package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

// Config stores configuration information stored in environment variables
type ApiConfig struct {
	DBConfig
	AWSS3Config
	AWSSQSConfig
	ServerConfig
}

// GetConfig returns environment variable config
func NewApiConfig() *ApiConfig {
	var cfg ApiConfig
	if err := envconfig.Process("instrumentation", &cfg); err != nil {
		log.Fatal(err.Error())
	}
	return &cfg
}
