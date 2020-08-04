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
	e.Use(
		middleware.CORS(),
		middleware.GzipWithConfig(middleware.GzipConfig{Level: 5}),
	)

	// Public Routes
	public := e.Group("")
	// Private Routes
	private := e.Group("")

	// JWT Middleware handles JWT Auth
	// SetCreatorUpaterFields sets context values from JWT claims for
	// creator, create_date, updater, update_date
	private.Use(
		middleware.JWTWithConfig(appconfig.JWTConfig),
		appconfig.IsLoggedIn,
	)

	// Public Routes
	// NOTE: ALL GET REQUESTS ARE ALLOWED WITHOUT AUTHENTICATION USING JWTConfig Skipper. See appconfig/jwt.go
	public.GET("instrumentation/projects", handlers.ListProjects(db))
	public.GET("instrumentation/projects/:project_id", handlers.GetProject(db))
	public.GET("instrumentation/projects/count", handlers.GetProjectCount(db))
	public.GET("instrumentation/projects/:project_id/instruments", handlers.ListProjectInstruments(db))
	public.GET("instrumentation/projects/:project_id/instruments/names", handlers.ListProjectInstrumentNames(db))
	public.GET("instrumentation/projects/:project_id/instrument_groups", handlers.ListProjectInstrumentGroups(db))
	public.GET("instrumentation/instrument_groups", handlers.ListInstrumentGroups(db))
	public.GET("instrumentation/instrument_groups/:instrument_group_id", handlers.GetInstrumentGroup(db))
	public.GET("instrumentation/instrument_groups/:instrument_group_id/instruments", handlers.ListInstrumentGroupInstruments(db))
	public.GET("instrumentation/instrument_groups/:instrument_group_id/timeseries", handlers.ListInstrumentGroupTimeseries(db))
	public.GET("instrumentation/instruments", handlers.ListInstruments(db))
	public.GET("instrumentation/instruments/count", handlers.GetInstrumentCount(db))
	public.GET("instrumentation/instruments/:instrument_id", handlers.GetInstrument(db))
	public.GET("instrumentation/instruments/notes", handlers.ListInstrumentNotes(db))
	public.GET("instrumentation/instruments/notes/:note_id", handlers.GetInstrumentNote(db))
	public.GET("instrumentation/instruments/:instrument_id/notes", handlers.ListInstrumentInstrumentNotes(db))
	public.GET("instrumentation/instruments/:instrument_id/notes/:note_id", handlers.GetInstrumentNote(db))
	public.GET("instrumentation/instruments/:instrument_id/zreference", handlers.ListInstrumentZReference(db))
	public.GET("instrumentation/instruments/:instrument_id/zreference/:zreference_id", handlers.GetInstrumentZReference(db))
	public.GET("instrumentation/instruments/:instrument_id/status", handlers.ListInstrumentStatus(db))
	public.GET("instrumentation/instruments/:instrument_id/status/:status_id", handlers.GetInstrumentStatus(db))
	public.GET("instrumentation/timeseries", handlers.ListTimeseries(db))
	public.GET("instrumentation/timeseries/:timeseries_id", handlers.GetTimeseries(db))
	public.GET("instrumentation/timeseries/:timeseries_id/measurements", handlers.ListTimeseriesMeasurements(db))
	public.GET("instrumentation/instruments/:instrument_id/timeseries", handlers.ListInstrumentTimeseries(db))
	public.GET("instrumentation/instruments/:instrument_id/timeseries/:timeseries_id/measurements", handlers.ListTimeseriesMeasurements(db))
	public.GET("instrumentation/instruments/:instrument_id/timeseries/:timeseries_id", handlers.GetTimeseries(db))
	public.GET("instrumentation/domains", handlers.GetDomains(db))
	public.GET("instrumentation/home", handlers.GetHome(db))
	public.POST("instrumentation/explorer", handlers.PostExplorer(db))

	// Authenticated Routes (Need CAC Login)
	// Projects
	private.POST("instrumentation/projects", handlers.CreateProjectBulk(db))
	private.PUT("instrumentation/projects/:project_id", handlers.UpdateProject(db))
	private.DELETE("instrumentation/projects/:project_id", handlers.DeleteFlagProject(db))
	// Project Instruments
	private.POST("instrumentation/projects/:project_id/instruments", handlers.CreateInstruments(db))
	// Instrument Groups
	private.POST("instrumentation/instrument_groups", handlers.CreateInstrumentGroupBulk(db))
	private.PUT("instrumentation/instrument_groups/:instrument_group_id", handlers.UpdateInstrumentGroup(db))
	private.DELETE("instrumentation/instrument_groups/:instrument_group_id", handlers.DeleteFlagInstrumentGroup(db))
	// Add or Remove instrument from Instrument Group
	private.POST("instrumentation/instrument_groups/:instrument_group_id/instruments", handlers.CreateInstrumentGroupInstruments(db))
	private.DELETE("instrumentation/instrument_groups/:instrument_group_id/instruments/:instrument_id", handlers.DeleteInstrumentGroupInstruments(db))
	// Instruments
	private.POST("instrumentation/instruments", handlers.CreateInstruments(db))
	private.PUT("instrumentation/instruments/:instrument_id", handlers.UpdateInstrument(db))
	private.DELETE("instrumentation/instruments/:instrument_id", handlers.DeleteFlagInstrument(db))
	// Instrument Notes(GET, PUT, DELETE work with or without instrument context in URL)
	private.POST("instrumentation/instruments/notes", handlers.CreateInstrumentNote(db))
	private.PUT("instrumentation/instruments/notes/:note_id", handlers.UpdateInstrumentNote(db))
	private.DELETE("instrumentation/instruments/notes/:note_id", handlers.DeleteInstrumentNote(db))
	private.PUT("instrumentation/instruments/:instrument_id/notes/:note_id", handlers.UpdateInstrumentNote(db))
	private.DELETE("instrumentation/instruments/:instrument_id/notes/:note_id", handlers.DeleteInstrumentNote(db))
	// Instrument ZReference
	private.POST("instrumentation/instruments/:instrument_id/zreference", handlers.CreateOrUpdateInstrumentZReference(db))
	private.DELETE("instrumentation/instruments/:instrument_id/zreference/:zreference_id", handlers.DeleteInstrumentZReference(db))
	// Instrument Status
	private.POST("instrumentation/instruments/:instrument_id/status", handlers.CreateOrUpdateInstrumentStatus(db))
	private.DELETE("instrumentation/instruments/:instrument_id/status/:status_id", handlers.DeleteInstrumentStatus(db))

	// Timeseries
	private.POST("instrumentation/timeseries", handlers.CreateTimeseries(db))
	private.PUT("instrumentation/timeseries/:timeseries_id", handlers.UpdateTimeseries(db))
	private.DELETE("instrumentation/timeseries/:timeseries_id", handlers.DeleteTimeseries(db))
	private.POST("instrumentation/timeseries/measurements", handlers.CreateOrUpdateTimeseriesMeasurements(db))

	log.Printf(
		"starting server; Running On AWS LAMBDA: %t",
		lambdaContext(),
	)
	if lambdaContext() {
		log.Fatal(gateway.ListenAndServe("localhost:3030", e))
	} else {
		log.Fatal(http.ListenAndServe("localhost:3030", e))
	}
}
