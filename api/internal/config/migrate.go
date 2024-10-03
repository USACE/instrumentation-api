package config

import "log"

type MigrateConfig struct {
	// Temporary Flyway replacement to avoid environment configuration changes
	DBUser        string   `env:"USER"`
	DBPassword    string   `env:"PASSWORD"`
	DBUrl         string   `env:"URL"`
	DBSchemas     []string `env:"SCHEMAS"`
	FileLocations string   `env:"LOCATIONS"`
}

func NewMigrateConfig() *MigrateConfig {
	var cfg MigrateConfig
	if err := parsePrefix("FLYWAY_", &cfg); err != nil {
		log.Fatal(err.Error())
	}
	return &cfg
}
