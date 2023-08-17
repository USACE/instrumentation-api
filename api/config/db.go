package config

import (
	"fmt"
	"os"
)

func DBConn() string {
	return fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s sslmode=%s",
		os.Getenv("INSTRUMENTATION_DBUSER"),
		os.Getenv("INSTRUMENTATION_DBPASS"),
		os.Getenv("INSTRUMENTATION_DBNAME"),
		os.Getenv("INSTRUMENTATION_DBHOST"),
		os.Getenv("INSTRUMENTATION_DBSSLMODE"),
	)
}
