package config

import "fmt"

type DBConfig struct {
	DBUser    string
	DBPass    string
	DBName    string
	DBHost    string
	DBSSLMode string
}

func (c *DBConfig) ConnStr() string {
	return fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=%s", c.DBUser, c.DBPass, c.DBName, c.DBHost, c.DBSSLMode)
}
