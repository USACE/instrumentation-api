package config

import "net/smtp"

type EmailConfig struct {
	SmtpConfig
	EmailSendMocked bool   `envconfig:"INSTRUMENTATION_EMAIL_SEND_MOCKED"`
	EmailFrom       string `envconfig:"INSTRUMENTATION_EMAIL_FROM"`
}

type SmtpConfig struct {
	SmtpHost     string `envconfig:"INSTRUMENTATION_SMTP_HOST"`
	SmtpPort     string `envconfig:"INSTRUMENTATION_SMTP_PORT"`
	SmtpAuthUser string `envconfig:"INSTRUMENTATION_SMTP_AUTH_USER"`
	SmtpAuthPass string `envconfig:"INSTRUMENTATION_SMTP_AUTH_PASS"`
}

func (cfg *SmtpConfig) SmtpAuth() smtp.Auth {
	return smtp.PlainAuth("", cfg.SmtpAuthUser, cfg.SmtpAuthPass, cfg.SmtpHost)
}

func (cfg *SmtpConfig) SmtpAddr() string {
	return cfg.SmtpHost + ":" + cfg.SmtpPort
}
