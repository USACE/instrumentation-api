package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"api/root/appconfig"
	"api/root/dbutils"
	"api/root/handlers"

	"github.com/apex/gateway"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	_ "github.com/lib/pq"
)

func lambdaContext() bool {

	value, exists := os.LookupEnv("LAMBDA")

	if exists && strings.ToUpper(value) == "TRUE" {
		return true
	}
	return false
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

	db := dbutils.Connection()

	e := echo.New()
	e.Use(middleware.CORS())

	// JWT Middleware handles JWT Auth
	// SetCreatorUpaterFields sets context values from JWT claims for
	// creator, create_date, updater, update_date
	e.Use(
		middleware.JWTWithConfig(appconfig.JWTConfig),
		appconfig.IsLoggedIn,
	)

	// Public Routes
	// NOTE: ALL GET REQUESTS ARE ALLOWED WITHOUT AUTHENTICATION USING JWTConfig Skipper. See appconfig/jwt.go
	e.GET("instrumentation/projects", handlers.ListProjects(db))
	e.GET("instrumentation/projects/:project_id", handlers.GetProject(db))
	e.GET("instrumentation/projects/count", handlers.GetProjectCount(db))
	e.GET("instrumentation/projects/:project_id/instruments", handlers.ListProjectInstruments(db))
	e.GET("instrumentation/projects/:project_id/instrument_groups", handlers.ListProjectInstrumentGroups(db))
	e.GET("instrumentation/instrument_groups", handlers.ListInstrumentGroups(db))
	e.GET("instrumentation/instrument_groups/:instrument_group_id", handlers.GetInstrumentGroup(db))
	e.GET("instrumentation/instrument_groups/:instrument_group_id/instruments", handlers.ListInstrumentGroupInstruments(db))
	e.GET("instrumentation/instrument_groups/:instrument_group_id/timeseries", handlers.ListInstrumentGroupTimeseries(db))
	e.GET("instrumentation/instruments", handlers.ListInstruments(db))
	e.GET("instrumentation/instruments/count", handlers.GetInstrumentCount(db))
	e.GET("instrumentation/instruments/:instrument_id", handlers.GetInstrument(db))
	e.GET("instrumentation/instruments/notes", handlers.ListInstrumentNotes(db))
	e.GET("instrumentation/instruments/notes/:note_id", handlers.GetInstrumentNote(db))
	e.GET("instrumentation/instruments/:instrument_id/notes", handlers.ListInstrumentInstrumentNotes(db))
	e.GET("instrumentation/instruments/:instrument_id/notes/:note_id", handlers.GetInstrumentNote(db))
	e.GET("instrumentation/instruments/:instrument_id/zreference", handlers.ListInstrumentZReference(db))
	e.GET("instrumentation/instruments/:instrument_id/zreference/:zreference_id", handlers.GetInstrumentZReference(db))
	e.GET("instrumentation/instruments/:instrument_id/status", handlers.ListInstrumentStatus(db))
	e.GET("instrumentation/instruments/:instrument_id/status/:status_id", handlers.GetInstrumentStatus(db))
	e.GET("instrumentation/timeseries", handlers.ListTimeseries(db))
	e.GET("instrumentation/timeseries/:timeseries_id", handlers.GetTimeseries(db))
	e.GET("instrumentation/timeseries/:timeseries_id/measurements", handlers.ListTimeseriesMeasurements(db))
	e.GET("instrumentation/instruments/:instrument_id/timeseries", handlers.ListInstrumentTimeseries(db))
	e.GET("instrumentation/instruments/:instrument_id/timeseries/:timeseries_id/measurements", handlers.ListTimeseriesMeasurements(db))
	e.GET("instrumentation/instruments/:instrument_id/timeseries/:timeseries_id", handlers.GetTimeseries(db))
	e.GET("instrumentation/domains", handlers.GetDomains(db))
	e.GET("instrumentation/home", handlers.GetHome(db))
	e.POST("instrumentation/explorer", handlers.PostExplorer(db))

	// Authenticated Routes (Need CAC Login)
	// Projects
	e.POST("instrumentation/projects", handlers.CreateProjectBulk(db))
	e.PUT("instrumentation/projects/:project_id", handlers.UpdateProject(db))
	e.DELETE("instrumentation/projects/:project_id", handlers.DeleteFlagProject(db))
	// Instrument Groups
	e.POST("instrumentation/instrument_groups", handlers.CreateInstrumentGroupBulk(db))
	e.PUT("instrumentation/instrument_groups/:instrument_group_id", handlers.UpdateInstrumentGroup(db))
	e.DELETE("instrumentation/instrument_groups/:instrument_group_id", handlers.DeleteFlagInstrumentGroup(db))
	// Add or Remove instrument from Instrument Group
	e.POST("instrumentation/instrument_groups/:instrument_group_id/instruments", handlers.CreateInstrumentGroupInstruments(db))
	e.DELETE("instrumentation/instrument_groups/:instrument_group_id/instruments/:instrument_id", handlers.DeleteInstrumentGroupInstruments(db))
	// Instruments
	e.POST("instrumentation/instruments", handlers.CreateInstrumentBulk(db))
	e.PUT("instrumentation/instruments/:instrument_id", handlers.UpdateInstrument(db))
	e.DELETE("instrumentation/instruments/:instrument_id", handlers.DeleteFlagInstrument(db))
	// Instrument Notes(GET, PUT, DELETE work with or without instrument context in URL)
	e.POST("instrumentation/instruments/notes", handlers.CreateInstrumentNote(db))
	e.PUT("instrumentation/instruments/notes/:note_id", handlers.UpdateInstrumentNote(db))
	e.DELETE("instrumentation/instruments/notes/:note_id", handlers.DeleteInstrumentNote(db))
	e.PUT("instrumentation/instruments/:instrument_id/notes/:note_id", handlers.UpdateInstrumentNote(db))
	e.DELETE("instrumentation/instruments/:instrument_id/notes/:note_id", handlers.DeleteInstrumentNote(db))
	// Instrument ZReference
	e.POST("instrumentation/instruments/:instrument_id/zreference", handlers.CreateOrUpdateInstrumentZReference(db))
	e.DELETE("instrumentation/instruments/:instrument_id/zreference/:zreference_id", handlers.DeleteInstrumentZReference(db))
	// Instrument Status
	e.POST("instrumentation/instruments/:instrument_id/status", handlers.CreateOrUpdateInstrumentStatus(db))
	e.DELETE("instrumentation/instruments/:instrument_id/status/:status_id", handlers.DeleteInstrumentStatus(db))

	// Timeseries
	e.POST("instrumentation/timeseries", handlers.CreateTimeseries(db))
	e.PUT("instrumentation/timeseries/:timeseries_id", handlers.UpdateTimeseries(db))
	e.DELETE("instrumentation/timeseries/:timeseries_id", handlers.DeleteTimeseries(db))
	e.POST("instrumentation/timeseries/measurements", handlers.CreateOrUpdateTimeseriesMeasurements(db))

	log.Printf(
		"starting server; Running On AWS LAMBDA: %t",
		lambdaContext(),
	)
	if lambdaContext() {
		log.Fatal(gateway.ListenAndServe(":3030", e))
	} else {
		log.Fatal(http.ListenAndServe(":3030", e))
	}
}
