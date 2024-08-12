package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

// Config stores configuration information stored in environment variables
type AlertCheckSchedulerConfig struct {
	DBConfig
	EmailConfig
	TriggerMocked bool `envconfig:"INSTRUMENTATION_AWS_ECS_TRIGGER_MOCKED"`
}

func NewAlertCheckConfig() *AlertCheckSchedulerConfig {
	var cfg AlertCheckSchedulerConfig
	if err := envconfig.Process("instrumentation", &cfg); err != nil {
		log.Fatal(err.Error())
	}
	return &cfg
}
