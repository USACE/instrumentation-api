package config

import (
	"log"
	"net/smtp"

	"github.com/kelseyhightower/envconfig"
)

// Config stores configuration information stored in environment variables
type AlertCheckConfig struct {
	DBConfig
	AWSECSTriggerMocked bool   `envconfig:"INSTRUMENTATION_AWS_ECS_TRIGGER_MOCKED"`
	EmailSendMocked     bool   `envconfig:"INSTRUMENTATION_EMAIL_SEND_MOCKED"`
	EmailFrom           string `envconfig:"INSTRUMENTATION_EMAIL_FROM"`
	SmtpHost            string `envconfig:"INSTRUMENTATION_SMTP_HOST"`
	SmtpPort            string `envconfig:"INSTRUMENTATION_SMTP_PORT"`
	SmtpAuthUser        string `envconfig:"INSTRUMENTATION_SMTP_AUTH_USER"`
	SmtpAuthPass        string `envconfig:"INSTRUMENTATION_SMTP_AUTH_PASS"`
}

func GetAlertCheckConfig() *AlertCheckConfig {
	var cfg AlertCheckConfig
	if err := envconfig.Process("instrumentation", &cfg); err != nil {
		log.Fatal(err.Error())
	}
	return &cfg
}

type SmtpConfig struct {
	SmtpAuth smtp.Auth
	SmtpAddr string
}

func GetSmtpConfig(c *AlertCheckConfig) *SmtpConfig {
	return &SmtpConfig{
		SmtpAuth: smtp.PlainAuth("", c.SmtpAuthUser, c.SmtpAuthPass, c.SmtpHost),
		SmtpAddr: c.SmtpHost + ":" + c.SmtpPort,
	}
}
