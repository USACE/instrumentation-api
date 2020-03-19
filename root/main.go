package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/apex/gateway"
	"github.com/labstack/echo"

	"api/root/handlers"
)

func dbConnStr() string {

	log.Printf("get database connection string")

	dbuser := os.Getenv("DB_USER")
	dbpass := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")
	dbhost := os.Getenv("DB_HOST")

	return fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s",
		dbuser, dbpass, dbname, dbhost,
	)
}

func initDB(connStr string) *sql.DB {

	log.Printf("Getting database connection")
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal("Could not connect to database")
		panic(err)
	}

	if db == nil {
		log.Panicf("database is nil")
	}

	log.Fatal(db)
	log.Fatal(err)

	return db
}

func main() {
	//  Here's what would typically be here:
	// lambda.Start(Handler)
	//
	// There were a few options on how to incorporate Echo v4 on Lambda.
	//
	// Landed here for now:
	//
	//     https://github.com/apex/gateway
	//     https://github.com/labstack/echo/issues/1195
	//
	// With this for local development:
	//     https://medium.com/a-man-with-no-server/running-go-aws-lambdas-locally-with-sls-framework-and-sam-af3d648d49cb
	//
	// This looks promising and is from awslabs, but Echo v4 support isn't quite there yet.
	// There is a pull request in progress, Re-evaluate in April 2020.
	//
	//    https://github.com/awslabs/aws-lambda-go-api-proxy
	//

	db := initDB(dbConnStr())

	e := echo.New()

	// Routes
	e.GET("/", handlers.GetRoot)
	// Instruments
	e.GET("/instruments", handlers.GetInstruments(db))
	// Time Series

	// Using gateway instead of net/http
	log.Fatal(gateway.ListenAndServe(":3000", e))

}
