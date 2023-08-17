package config

import (
	"log"
	"net/smtp"
	"os"

	"github.com/kelseyhightower/envconfig"
)

// Config stores configuration information stored in environment variables
type AlertCheckConfig struct {
	AWSECSTriggerMocked bool   `envconfig:"INSTRUMENTATION_AWS_ECS_TRIGGER_MOCKED"`
	EmailSendMocked     bool   `envconfig:"INSTRUMENTATION_EMAIL_SEND_MOCKED"`
	EmailFrom           string `envconfig:"INSTRUMENTATION_EMAIL_FROM"`
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
		SmtpAuth: smtp.PlainAuth(
			"",
			os.Getenv("INSTRUMENTATION_SMTP_AUTH_USER"),
			os.Getenv("INSTRUMENTATION_SMTP_AUTH_PASS"),
			os.Getenv("INSTRUMENTATION_SMTP_HOST"),
		),
		SmtpAddr: os.Getenv("INSTRUMENTATION_SMTP_HOST") + ":" + os.Getenv("INSTRUMENTATION_SMTP_PORT"),
	}
}
