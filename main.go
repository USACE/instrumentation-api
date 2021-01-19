package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/USACE/instrumentation-api/dbutils"
	"github.com/USACE/instrumentation-api/handlers"
	"github.com/USACE/instrumentation-api/middleware"
	"github.com/apex/gateway"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/kelseyhightower/envconfig"

	"github.com/labstack/echo/v4"

	_ "github.com/lib/pq"
)

// Config stores configuration information stored in environment variables
type Config struct {
	SkipJWT             bool
	LambdaContext       bool
	DBUser              string
	DBPass              string
	DBName              string
	DBHost              string
	DBSSLMode           string
	HeartbeatKey        string
	AWSS3Endpoint       string `envconfig:"INSTRUMENTATION_AWS_S3_ENDPOINT"`
	AWSS3Region         string `envconfig:"INSTRUMENTATION_AWS_S3_REGION"`
	AWSS3DisableSSL     bool   `envconfig:"INSTRUMENTATION_AWS_S3_DISABLE_SSL"`
	AWSS3ForcePathStyle bool   `envconfig:"INSTRUMENTATION_AWS_S3_FORCE_PATH_STYLE"`
	AWSS3Bucket         string `envconfig:"INSTRUMENTATION_AWS_S3_BUCKET"`
}

func awsConfig(cfg *Config) *aws.Config {
	awsConfig := aws.NewConfig().WithRegion(cfg.AWSS3Region)

	// Used for "minio" during development
	awsConfig.WithDisableSSL(cfg.AWSS3DisableSSL)
	awsConfig.WithS3ForcePathStyle(cfg.AWSS3ForcePathStyle)
	if cfg.AWSS3Endpoint != "" {
		awsConfig.WithEndpoint(cfg.AWSS3Endpoint)
	}

	return awsConfig
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

	// Config holding all environment variables
	var cfg Config
	if err := envconfig.Process("instrumentation", &cfg); err != nil {
		log.Fatal(err.Error())
	}

	// AWS Config used to get S3 Session/Client
	awsCfg := awsConfig(&cfg)

	db := dbutils.Connection(
		fmt.Sprintf(
			"user=%s password=%s dbname=%s host=%s sslmode=%s binary_parameters=yes",
			cfg.DBUser, cfg.DBPass, cfg.DBName, cfg.DBHost, cfg.DBSSLMode,
		),
	)

	e := echo.New()
	e.Use(middleware.CORS, middleware.GZIP)
	versioned := e.Group("/instrumentation") // TODO: /instrumentation/v1/

	// Public Media Routes (No Version)
	publicMedia := versioned.Group("") // TODO: /instrumentation
	publicMedia.GET("/projects/:project_slug/images/*", handlers.GetMedia(awsCfg, &cfg.AWSS3Bucket))

	// Key Routes (Heartbeat Function)
	keyed := versioned.Group("")
	keyed.Use(middleware.KeyAuth(cfg.HeartbeatKey))
	keyed.POST("/heartbeat", handlers.DoHeartbeat(db))

	// Public Routes
	public := versioned.Group("")

	// Private Routes (Authenticated, Authorized)
	private := versioned.Group("")
	if cfg.SkipJWT == true {
		private.Use(middleware.MockIsLoggedIn)
	} else {
		private.Use(middleware.JWT, middleware.IsLoggedIn)
	}

	// Heartbeat
	public.GET("/heartbeat/latest", handlers.GetLatestHeartbeat(db))
	public.GET("/heartbeats", handlers.ListHeartbeats(db))

	// AlertConfigs
	public.GET("/projects/:project_id/instruments/:instrument_id/alert_configs", handlers.ListInstrumentAlertConfigs(db))
	public.GET("/projects/:project_id/instruments/:instrument_id/alert_configs/:alert_config_id", handlers.GetAlertConfig(db))
	private.POST("/projects/:project_id/instruments/:instrument_id/alert_configs", handlers.CreateInstrumentAlertConfigs(db))
	private.PUT("/projects/:project_id/instruments/:instrument_id/alert_configs/:alert_config_id", handlers.UpdateInstrumentAlertConfig(db))
	private.DELETE("/projects/:project_id/instruments/:instrument_id/alert_configs/:alert_config_id", handlers.DeleteInstrumentAlertConfig(db))

	// Alerts
	public.GET("/projects/:project_id/instruments/:instrument_id/alerts", handlers.ListAlertsForInstrument(db))
	private.GET("/my_alerts", handlers.ListMyAlerts(db)) // Private because token required to determine user (i.e. who is "me")
	private.POST("/my_alerts/:alert_id/read", handlers.DoAlertRead(db))
	private.POST("/my_alerts/:alert_id/unread", handlers.DoAlertUnread(db))

	// AlertSubscriptions
	private.GET("/my_alert_subscriptions", handlers.ListMyAlertSubscriptions(db)) // Private because token required to determine user (i.e. who is "me")
	private.POST("/projects/:project_id/instruments/:instrument_id/alert_configs/:alert_config_id/subscribe", handlers.SubscribeProfileToAlerts(db))
	private.POST("/projects/:project_id/instruments/:instrument_id/alert_configs/:alert_config_id/unsubscribe", handlers.UnsubscribeProfileToAlerts(db))
	private.PUT("/alert_subscriptions/:alert_subscription_id", handlers.UpdateMyAlertSubscription(db))

	// Profile
	private.GET("/myprofile", handlers.GetMyProfile(db))
	private.GET("/my_profile", handlers.GetMyProfile(db))
	private.POST("/profiles", handlers.CreateProfile(db))

	// Email Autocomplete
	public.GET("/email_autocomplete", handlers.ListEmailAutocomplete(db))

	// Projects
	public.GET("/projects", handlers.ListProjects(db))
	public.GET("/projects/:project_id", handlers.GetProject(db))
	public.GET("/projects/count", handlers.GetProjectCount(db))
	public.GET("/projects/:project_id/instruments", handlers.ListProjectInstruments(db))
	public.GET("/projects/:project_id/instruments/names", handlers.ListProjectInstrumentNames(db))
	public.GET("/projects/:project_id/instrument_groups", handlers.ListProjectInstrumentGroups(db))
	private.POST("/projects", handlers.CreateProjectBulk(db))
	private.PUT("/projects/:project_id", handlers.UpdateProject(db))
	private.DELETE("/projects/:project_id", handlers.DeleteFlagProject(db))

	// Project Timeseries
	private.POST("/projects/:project_id/timeseries/:timeseries_id", handlers.CreateProjectTimeseries(db))
	private.DELETE("/projects/:project_id/timeseries/:timeseries_id", handlers.DeleteProjectTimeseries(db))

	// Instruments
	public.GET("/instruments", handlers.ListInstruments(db))
	public.GET("/instruments/count", handlers.GetInstrumentCount(db))
	public.GET("/instruments/:instrument_id", handlers.GetInstrument(db))
	private.POST("/projects/:project_id/instruments", handlers.CreateInstruments(db))
	private.POST("/instruments", handlers.CreateInstruments(db))
	private.PUT("/instruments/:instrument_id", handlers.UpdateInstrument(db))
	private.PUT("/projects/:project_id/instruments/:instrument_id", handlers.UpdateInstrument(db))
	private.DELETE("/instruments/:instrument_id", handlers.DeleteFlagInstrument(db))

	// Instrument Groups
	public.GET("/instrument_groups", handlers.ListInstrumentGroups(db))
	public.GET("/instrument_groups/:instrument_group_id", handlers.GetInstrumentGroup(db))
	public.GET("/instrument_groups/:instrument_group_id/instruments", handlers.ListInstrumentGroupInstruments(db))
	public.GET("/instrument_groups/:instrument_group_id/timeseries", handlers.ListInstrumentGroupTimeseries(db))
	private.POST("/instrument_groups", handlers.CreateInstrumentGroup(db))
	private.PUT("/instrument_groups/:instrument_group_id", handlers.UpdateInstrumentGroup(db))
	private.DELETE("/instrument_groups/:instrument_group_id", handlers.DeleteFlagInstrumentGroup(db))
	// Add or Remove instrument from Instrument Group
	private.POST("/instrument_groups/:instrument_group_id/instruments", handlers.CreateInstrumentGroupInstruments(db))
	private.DELETE("/instrument_groups/:instrument_group_id/instruments/:instrument_id", handlers.DeleteInstrumentGroupInstruments(db))

	// Instrument Constants (same as a timeseries in structure/payload)
	public.GET("/projects/:project_id/instruments/:instrument_id/constants", handlers.ListInstrumentConstants(db))
	private.POST("/projects/:project_id/instruments/:instrument_id/constants", handlers.CreateInstrumentConstants(db))
	private.DELETE("/projects/:project_id/instruments/:instrument_id/constants/:timeseries_id", handlers.DeleteInstrumentConstant(db))

	// Instrument Notes(GET, PUT, DELETE work with or without instrument context in URL)
	public.GET("/instruments/notes", handlers.ListInstrumentNotes(db))
	public.GET("/instruments/notes/:note_id", handlers.GetInstrumentNote(db))
	public.GET("/instruments/:instrument_id/notes", handlers.ListInstrumentInstrumentNotes(db))
	public.GET("/instruments/:instrument_id/notes/:note_id", handlers.GetInstrumentNote(db))
	private.POST("/instruments/notes", handlers.CreateInstrumentNote(db))
	private.PUT("/instruments/notes/:note_id", handlers.UpdateInstrumentNote(db))
	private.DELETE("/instruments/notes/:note_id", handlers.DeleteInstrumentNote(db))
	private.PUT("/instruments/:instrument_id/notes/:note_id", handlers.UpdateInstrumentNote(db))
	private.DELETE("/instruments/:instrument_id/notes/:note_id", handlers.DeleteInstrumentNote(db))

	// Instrument Status
	public.GET("/instruments/:instrument_id/status", handlers.ListInstrumentStatus(db))
	public.GET("/instruments/:instrument_id/status/:status_id", handlers.GetInstrumentStatus(db))
	private.POST("/instruments/:instrument_id/status", handlers.CreateOrUpdateInstrumentStatus(db))
	private.DELETE("/instruments/:instrument_id/status/:status_id", handlers.DeleteInstrumentStatus(db))

	// Timeseries
	public.GET("/projects/:project_id/instruments/:instrument_id/timeseries", handlers.ListInstrumentTimeseries(db))
	public.GET("/timeseries", handlers.ListTimeseries(db))
	public.GET("/timeseries/:timeseries_id", handlers.GetTimeseries(db))
	public.GET("/instruments/:instrument_id/timeseries/:timeseries_id", handlers.GetTimeseries(db))
	public.GET("/timeseries/:timeseries_id/measurements", handlers.ListTimeseriesMeasurements(db))
	public.GET("/instruments/:instrument_id/timeseries/:timeseries_id/measurements", handlers.ListTimeseriesMeasurements(db))

	private.POST("/timeseries", handlers.CreateTimeseries(db))
	private.PUT("/timeseries/:timeseries_id", handlers.UpdateTimeseries(db))
	private.DELETE("/timeseries/:timeseries_id", handlers.DeleteTimeseries(db))

	private.POST("/timeseries/measurements", handlers.CreateOrUpdateTimeseriesMeasurements(db))
	private.POST("/projects/:project_id/timeseries_measurements", handlers.CreateOrUpdateTimeseriesMeasurements(db))

	// Collection Groups
	public.GET("/projects/:project_id/collection_groups", handlers.ListCollectionGroups(db))
	public.GET("/projects/:project_id/collection_groups/:collection_group_id", handlers.GetCollectionGroupDetails(db))
	private.POST("/projects/:project_id/collection_groups", handlers.CreateCollectionGroup(db))
	private.PUT("/projects/:project_id/collection_groups/:collection_group_id", handlers.UpdateCollectionGroup(db))
	private.DELETE("/projects/:project_id/collection_groups/:collection_group_id", handlers.DeleteCollectionGroup(db))
	// // Collection Group; Add Timeseries to collection_group
	private.POST("/projects/:project_id/collection_groups/:collection_group_id/timeseries/:timeseries_id", handlers.AddTimeseriesToCollectionGroup(db))
	// // Collection Group; Remove Timeseries from collection_group
	private.DELETE("/projects/:project_id/collection_groups/:collection_group_id/timeseries/:timeseries_id", handlers.RemoveTimeseriesFromCollectionGroup(db))

	// Misc
	public.GET("/domains", handlers.GetDomains(db))
	public.GET("/home", handlers.GetHome(db))
	public.POST("/explorer", handlers.PostExplorer(db))

	// OpenDCS Configuration
	public.GET("/opendcs/sites", handlers.ListOpendcsSites(db))

	if cfg.LambdaContext {
		log.Print("starting server; Running On AWS LAMBDA")
		log.Fatal(gateway.ListenAndServe("localhost:3030", e))
	} else {
		log.Print("starting server")
		log.Fatal(http.ListenAndServe(":80", e))
	}
}
