package dbutils

import (
	"log"

	"github.com/jmoiron/sqlx"
)

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

// Connection is a database connnection
func Connection(connStr string) *sqlx.DB {
	conn := initDB(connStr)
	return conn
}
