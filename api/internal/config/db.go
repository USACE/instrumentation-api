package config

import (
	"fmt"
)

type DBConfig struct {
	DBUser    string `env:"DBUSER"`
	DBPass    string `env:"DBPASS"`
	DBName    string `env:"DBNAME"`
	DBHost    string `env:"DBHOST"`
	DBSSLMode string `env:"DBSSLMODE"`
}

func (cfg *DBConfig) ConnStr() string {
	s := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=%s", cfg.DBUser, cfg.DBPass, cfg.DBName, cfg.DBHost, cfg.DBSSLMode)
	return s
}
