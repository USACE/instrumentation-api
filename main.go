package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/USACE/instrumentation-api/dbutils"
	"github.com/USACE/instrumentation-api/handlers"
	"github.com/USACE/instrumentation-api/middleware"

	"github.com/apex/gateway"
	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo/v4"

	_ "github.com/lib/pq"
)

// Config stores configuration information stored in environment variables
type Config struct {
	SkipJWT       bool
	LambdaContext bool
	DBUser        string
	DBPass        string
	DBName        string
	DBHost        string
	DBSSLMode     string
	HeartbeatKey  string
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

	var cfg Config
	if err := envconfig.Process("instrumentation", &cfg); err != nil {
		log.Fatal(err.Error())
	}

	db := dbutils.Connection(
		fmt.Sprintf(
			"user=%s password=%s dbname=%s host=%s sslmode=%s binary_parameters=yes",
			cfg.DBUser, cfg.DBPass, cfg.DBName, cfg.DBHost, cfg.DBSSLMode,
		),
	)

	// All Routes
	e := echo.New()
	e.Use(middleware.CORS, middleware.GZIP)

	// Public Routes
	public := e.Group("")

	// Private Routes
	private := e.Group("")

	// Key Routes (Heartbeat Function)
	keyed := e.Group("")
	keyed.Use(middleware.KeyAuth(cfg.HeartbeatKey))

	if cfg.SkipJWT == true {
		private.Use(middleware.MockIsLoggedIn)
	} else {
		private.Use(middleware.JWT, middleware.IsLoggedIn)
	}

	// Heartbeat
	keyed.POST("instrumentation/heartbeat", handlers.DoHeartbeat(db))
	public.GET("instrumentation/heartbeat/latest", handlers.GetLatestHeartbeat(db))
	public.GET("instrumentation/heartbeats", handlers.ListHeartbeats(db))

	// AlertConfigs
	public.GET("instrumentation/projects/:project_id/instruments/:instrument_id/alert_configs", handlers.ListInstrumentAlertConfigs(db))
	public.GET("instrumentation/projects/:project_id/instruments/:instrument_id/alert_configs/:alert_config_id", handlers.GetAlertConfig(db))
	private.POST("instrumentation/projects/:project_id/instruments/:instrument_id/alert_configs", handlers.CreateInstrumentAlertConfigs(db))
	private.PUT("instrumentation/projects/:project_id/instruments/:instrument_id/alert_configs/:alert_config_id", handlers.UpdateInstrumentAlertConfig(db))
	private.DELETE("instrumentation/projects/:project_id/instruments/:instrument_id/alert_configs/:alert_config_id", handlers.DeleteInstrumentAlertConfig(db))

	// Alerts
	public.GET("instrumentation/projects/:project_id/instruments/:instrument_id/alerts", handlers.ListAlertsForInstrument(db))
	private.GET("instrumentation/my_alerts", handlers.ListMyAlerts(db)) // Private because token required to determine user (i.e. who is "me")
	private.POST("instrumentation/my_alerts/:alert_id/read", handlers.DoAlertRead(db))
	private.POST("instrumentation/my_alerts/:alert_id/unread", handlers.DoAlertUnread(db))

	// AlertSubscriptions
	private.GET("instrumentation/my_alert_subscriptions", handlers.ListMyAlertSubscriptions(db)) // Private because token required to determine user (i.e. who is "me")
	private.POST("instrumentation/projects/:project_id/instruments/:instrument_id/alert_configs/:alert_config_id/subscribe", handlers.SubscribeProfileToAlerts(db))
	private.POST("instrumentation/projects/:project_id/instruments/:instrument_id/alert_configs/:alert_config_id/unsubscribe", handlers.UnsubscribeProfileToAlerts(db))
	private.PUT("instrumentation/alert_subscriptions/:alert_subscription_id", handlers.UpdateMyAlertSubscription(db))

	// Profile
	private.GET("instrumentation/myprofile", handlers.GetMyProfile(db))
	private.GET("instrumentation/my_profile", handlers.GetMyProfile(db))
	private.POST("instrumentation/profiles", handlers.CreateProfile(db))

	// Email Autocomplete
	public.GET("instrumentation/email_autocomplete", handlers.ListEmailAutocomplete(db))

	// Projects
	public.GET("instrumentation/projects", handlers.ListProjects(db))
	public.GET("instrumentation/projects/:project_id", handlers.GetProject(db))
	public.GET("instrumentation/projects/count", handlers.GetProjectCount(db))
	public.GET("instrumentation/projects/:project_id/instruments", handlers.ListProjectInstruments(db))
	public.GET("instrumentation/projects/:project_id/instruments/names", handlers.ListProjectInstrumentNames(db))
	public.GET("instrumentation/projects/:project_id/instrument_groups", handlers.ListProjectInstrumentGroups(db))
	private.POST("instrumentation/projects", handlers.CreateProjectBulk(db))
	private.PUT("instrumentation/projects/:project_id", handlers.UpdateProject(db))
	private.DELETE("instrumentation/projects/:project_id", handlers.DeleteFlagProject(db))

	// Project Timeseries
	private.POST("instrumentation/projects/:project_id/timeseries/:timeseries_id", handlers.CreateProjectTimeseries(db))
	private.DELETE("instrumentation/projects/:project_id/timeseries/:timeseries_id", handlers.DeleteProjectTimeseries(db))

	// Instruments
	public.GET("instrumentation/instruments", handlers.ListInstruments(db))
	public.GET("instrumentation/instruments/count", handlers.GetInstrumentCount(db))
	public.GET("instrumentation/instruments/:instrument_id", handlers.GetInstrument(db))
	private.POST("instrumentation/projects/:project_id/instruments", handlers.CreateInstruments(db))
	private.POST("instrumentation/instruments", handlers.CreateInstruments(db))
	private.PUT("instrumentation/instruments/:instrument_id", handlers.UpdateInstrument(db))
	private.PUT("instrumentation/projects/:project_id/instruments/:instrument_id", handlers.UpdateInstrument(db))
	private.DELETE("instrumentation/instruments/:instrument_id", handlers.DeleteFlagInstrument(db))

	// Instrument Groups
	public.GET("instrumentation/instrument_groups", handlers.ListInstrumentGroups(db))
	public.GET("instrumentation/instrument_groups/:instrument_group_id", handlers.GetInstrumentGroup(db))
	public.GET("instrumentation/instrument_groups/:instrument_group_id/instruments", handlers.ListInstrumentGroupInstruments(db))
	public.GET("instrumentation/instrument_groups/:instrument_group_id/timeseries", handlers.ListInstrumentGroupTimeseries(db))
	private.POST("instrumentation/instrument_groups", handlers.CreateInstrumentGroup(db))
	private.PUT("instrumentation/instrument_groups/:instrument_group_id", handlers.UpdateInstrumentGroup(db))
	private.DELETE("instrumentation/instrument_groups/:instrument_group_id", handlers.DeleteFlagInstrumentGroup(db))
	// Add or Remove instrument from Instrument Group
	private.POST("instrumentation/instrument_groups/:instrument_group_id/instruments", handlers.CreateInstrumentGroupInstruments(db))
	private.DELETE("instrumentation/instrument_groups/:instrument_group_id/instruments/:instrument_id", handlers.DeleteInstrumentGroupInstruments(db))

	// Add or Remove Instrument Constants
	private.POST("instrumentation/instruments/:instrument_id/constants/:timeseries_id", handlers.CreateInstrumentConstant(db))
	private.DELETE("instrumentation/instruments/:instrument_id/constants/:timeseries_id", handlers.DeleteInstrumentConstant(db))

	// Instrument Notes(GET, PUT, DELETE work with or without instrument context in URL)
	public.GET("instrumentation/instruments/notes", handlers.ListInstrumentNotes(db))
	public.GET("instrumentation/instruments/notes/:note_id", handlers.GetInstrumentNote(db))
	public.GET("instrumentation/instruments/:instrument_id/notes", handlers.ListInstrumentInstrumentNotes(db))
	public.GET("instrumentation/instruments/:instrument_id/notes/:note_id", handlers.GetInstrumentNote(db))
	private.POST("instrumentation/instruments/notes", handlers.CreateInstrumentNote(db))
	private.PUT("instrumentation/instruments/notes/:note_id", handlers.UpdateInstrumentNote(db))
	private.DELETE("instrumentation/instruments/notes/:note_id", handlers.DeleteInstrumentNote(db))
	private.PUT("instrumentation/instruments/:instrument_id/notes/:note_id", handlers.UpdateInstrumentNote(db))
	private.DELETE("instrumentation/instruments/:instrument_id/notes/:note_id", handlers.DeleteInstrumentNote(db))

	// Instrument Status
	public.GET("instrumentation/instruments/:instrument_id/status", handlers.ListInstrumentStatus(db))
	public.GET("instrumentation/instruments/:instrument_id/status/:status_id", handlers.GetInstrumentStatus(db))
	private.POST("instrumentation/instruments/:instrument_id/status", handlers.CreateOrUpdateInstrumentStatus(db))
	private.DELETE("instrumentation/instruments/:instrument_id/status/:status_id", handlers.DeleteInstrumentStatus(db))

	// Timeseries
	public.GET("instrumentation/projects/:project_id/instruments/:instrument_id/timeseries", handlers.ListInstrumentTimeseries(db))
	public.GET("instrumentation/timeseries", handlers.ListTimeseries(db))
	public.GET("instrumentation/timeseries/:timeseries_id", handlers.GetTimeseries(db))
	public.GET("instrumentation/instruments/:instrument_id/timeseries/:timeseries_id", handlers.GetTimeseries(db))
	public.GET("instrumentation/timeseries/:timeseries_id/measurements", handlers.ListTimeseriesMeasurements(db))
	public.GET("instrumentation/instruments/:instrument_id/timeseries/:timeseries_id/measurements", handlers.ListTimeseriesMeasurements(db))

	private.POST("instrumentation/timeseries", handlers.CreateTimeseries(db))
	private.PUT("instrumentation/timeseries/:timeseries_id", handlers.UpdateTimeseries(db))
	private.DELETE("instrumentation/timeseries/:timeseries_id", handlers.DeleteTimeseries(db))

	private.POST("instrumentation/timeseries/measurements", handlers.CreateOrUpdateTimeseriesMeasurements(db))

	// Misc
	public.GET("instrumentation/domains", handlers.GetDomains(db))
	public.GET("instrumentation/home", handlers.GetHome(db))
	public.POST("instrumentation/explorer", handlers.PostExplorer(db))

	// OpenDCS Configuration
	public.GET("instrumentation/opendcs/sites", handlers.ListOpendcsSites(db))

	if cfg.LambdaContext {
		log.Print("starting server; Running On AWS LAMBDA")
		log.Fatal(gateway.ListenAndServe("localhost:3030", e))
	} else {
		log.Print("starting server")
		log.Fatal(http.ListenAndServe("localhost:3030", e))
	}
}
