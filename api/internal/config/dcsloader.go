package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

// DcsLoaderConfig holds parameters parsed from env variables.
// Note awsSQSQueueURL private variable. Public method is AWSSQSQueueURL()
type DcsLoaderConfig struct {
	AWSS3Config
	AWSSQSConfig
	PostURL string `envconfig:"POST_URL"`
	APIKey  string `envconfig:"API_KEY"`
}

func NewDcsLoaderConfig() *DcsLoaderConfig {
	var cfg DcsLoaderConfig
	if err := envconfig.Process("loader", &cfg); err != nil {
		log.Fatal(err.Error())
	}
	return &cfg
}
