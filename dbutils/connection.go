package dbutils

import (
	"log"

	// _ "github.com/jackc/pgx/v4/"
	_ "github.com/jackc/pgx/v4/stdlib"

	"github.com/jmoiron/sqlx"
)

func initDB(connStr string) *sqlx.DB {

	log.Printf("Getting database connection")
	db, err := sqlx.Connect("pgx", connStr)

	if err != nil {
		log.Fatal("Could not connect to database")
		panic(err)
	}

	if db == nil {
		log.Panicf("database is nil")
	}

	db.SetMaxOpenConns(10)

	return db
}

// Connection is a database connnection
func Connection(connStr string) *sqlx.DB {
	conn := initDB(connStr)
	return conn
}
