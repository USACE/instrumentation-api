package config

import (
	"log"
)

// DcsLoaderConfig holds parameters parsed from env variables.
// Note awsSQSQueueURL private variable. Public method is AWSSQSQueueURL()
type DcsLoaderConfig struct {
	AWSS3Config
	AWSSQSConfig
	PostURL string `env:"POST_URL"`
	APIKey  string `env:"API_KEY"`
}

func NewDcsLoaderConfig() *DcsLoaderConfig {
	var cfg DcsLoaderConfig
	if err := parsePrefix("LOADER_", &cfg); err != nil {
		log.Fatal(err.Error())
	}
	return &cfg
}
