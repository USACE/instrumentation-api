package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

// Config stores configuration information stored in environment variables
type AlertCheckConfig struct {
	DBConfig
	SmtpConfig
	AWSECSTriggerMocked bool   `envconfig:"INSTRUMENTATION_AWS_ECS_TRIGGER_MOCKED"`
	EmailSendMocked     bool   `envconfig:"INSTRUMENTATION_EMAIL_SEND_MOCKED"`
	EmailFrom           string `envconfig:"INSTRUMENTATION_EMAIL_FROM"`
}

func NewAlertCheckConfig() *AlertCheckConfig {
	var cfg AlertCheckConfig
	if err := envconfig.Process("instrumentation", &cfg); err != nil {
		log.Fatal(err.Error())
	}
	return &cfg
}
