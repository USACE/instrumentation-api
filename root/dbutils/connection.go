package dbutils

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
)

func dbConnStr() string {

	log.Printf("get database connection string")

	dbuser := os.Getenv("DB_USER")
	dbpass := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")
	dbhost := os.Getenv("DB_HOST")
	sslmode := os.Getenv("DB_SSLMODE")

	return fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s sslmode=%s binary_parameters=yes",
		dbuser, dbpass, dbname, dbhost, sslmode,
	)
}

func initDB(connStr string) *sqlx.DB {

	log.Printf("Getting database connection")
	db, err := sqlx.Open("postgres", connStr)

	if err != nil {
		log.Fatal("Could not connect to database")
		panic(err)
	}

	if db == nil {
		log.Panicf("database is nil")
	}

	return db
}

func Connection() *sqlx.DB {
	conn := initDB(dbConnStr())
	return conn
}
