package config

import "net/smtp"

type SmtpConfig struct {
	SmtpHost     string `env:"SMTP_HOST"`
	SmtpPort     string `env:"SMTP_PORT"`
	SmtpAuthUser string `env:"SMTP_AUTH_USER"`
	SmtpAuthPass string `env:"SMTP_AUTH_PASS"`
}

func (cfg *SmtpConfig) SmtpAuth() smtp.Auth {
	return smtp.PlainAuth("", cfg.SmtpAuthUser, cfg.SmtpAuthPass, cfg.SmtpHost)
}

func (cfg *SmtpConfig) SmtpAddr() string {
	return cfg.SmtpHost + ":" + cfg.SmtpPort
}
