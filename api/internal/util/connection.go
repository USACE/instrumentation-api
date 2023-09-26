package util

import (
	"log"
	"time"

	// _ "github.com/jackc/pgx/v4/"
	_ "github.com/jackc/pgx/v4/stdlib"

	"github.com/jmoiron/sqlx"
)

// TODO: Remove this in favor of database package
// (once refactor finished)
func initDB(connStr string) *sqlx.DB {

	log.Printf("Getting database connection")
	db, err := sqlx.Connect("pgx", connStr)

	if err != nil {
		log.Fatalf("Could not connect to database: %s", err.Error())
	}

	if db == nil {
		log.Panicf("database is nil")
	}

	db.SetMaxOpenConns(15)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Minute * 30)

	return db
}

// Connection is a database connnection
func Connection(connStr string) *sqlx.DB {
	conn := initDB(connStr)
	return conn
}
