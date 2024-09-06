package config

import "log"

// Config stores configuration information stored in environment variables
type AlertCheckConfig struct {
	DBConfig
	SmtpConfig
	TriggerMocked   bool   `env:"AWS_ECS_TRIGGER_MOCKED"`
	EmailSendMocked bool   `env:"EMAIL_SEND_MOCKED"`
	EmailFrom       string `env:"EMAIL_FROM"`
}

func NewAlertCheckConfig() *AlertCheckConfig {
	var cfg AlertCheckConfig
	if err := parsePrefix("INSTRUMENTATION_", &cfg); err != nil {
		log.Fatalf(err.Error())
	}
	return &cfg
}
