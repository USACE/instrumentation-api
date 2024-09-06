package config

import "fmt"

type DBConfig struct {
	DBUser    string `env:"DBUSER"`
	DBPass    string `env:"DBPASS"`
	DBName    string `env:"DBNAME"`
	DBHost    string `env:"DBHOST"`
	DBSSLMode string `env:"DBSSLMODE"`
}

func (c *DBConfig) ConnStr() string {
	return fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=%s", c.DBUser, c.DBPass, c.DBName, c.DBHost, c.DBSSLMode)
}
