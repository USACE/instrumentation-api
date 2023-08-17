package main

import (
	"log"
	"net/http"

	"github.com/USACE/instrumentation-api/api/config"
	"github.com/USACE/instrumentation-api/api/dbutils"
	"github.com/USACE/instrumentation-api/api/handlers"
	"github.com/USACE/instrumentation-api/api/middleware"
	"github.com/USACE/instrumentation-api/api/models"
	"github.com/apex/gateway"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/labstack/echo/v4"
	echomw "github.com/labstack/echo/v4/middleware"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	cfg := config.GetApiConfig()
	db := dbutils.Connection(config.DBConn())

	awsCfg := config.AWSConfig(cfg)
	sess := session.Must(session.NewSession(awsCfg))
	s3c := s3.New(sess)

	hashExtractor := func(keyID string) (string, error) {
		k, err := models.GetTokenInfoByTokenID(db, &keyID)
		if err != nil {
			return "", err
		}
		return k.Hash, nil
	}

	e := echo.New()
	e.Use(middleware.CORS, middleware.GZIP)
	public := e.Group(cfg.RoutePrefix) // TODO: /instrumentation/v1/

	// Media Routes
	public.GET("/projects/:project_slug/images/*", handlers.GetMedia(s3c, &cfg.AWSS3Bucket, "/midas", &cfg.RoutePrefix))

	// private routes; can be authenticated via cac or token
	// setting the second parameter passed to each middleware function to "true"
	// means that if `?key=` is in the query parameters, JWT middleware will automatically
	// pass (call next(c)); authentication will then be handled by keyauth
	private := e.Group(cfg.RoutePrefix)
	if cfg.AuthJWTMocked {
		private.Use(middleware.JWTMock(cfg.AuthDisabled, true))
	} else {
		private.Use(middleware.JWT(cfg.AuthDisabled, true))
	}
	// Attach keyauth middleware
	private.Use(middleware.KeyAuth(
		cfg.AuthDisabled,
		cfg.ApplicationKey,
		hashExtractor,
	))

	// keyAuth not allowed on these routes
	CACOnly := e.Group(cfg.RoutePrefix)
	if cfg.AuthJWTMocked {
		CACOnly.Use(middleware.JWTMock(cfg.AuthDisabled, false))
	} else {
		CACOnly.Use(middleware.JWT(cfg.AuthDisabled, false))
	}

	app := e.Group(cfg.RoutePrefix)
	app.Use(echomw.KeyAuthWithConfig(echomw.KeyAuthConfig{
		KeyLookup: "query:key",
		Validator: func(key string, c echo.Context) (bool, error) {
			return key == cfg.ApplicationKey, nil
		},
	}))

	// AttachProfileMiddleware attaches ProfileID to context, whether
	// authenticated by token or api key
	private.Use(middleware.EDIPIMiddleware, middleware.AttachProfileMiddleware(db))
	CACOnly.Use(middleware.EDIPIMiddleware, middleware.CACOnlyMiddleware)

	// Profile and Tokens
	CACOnly.POST("/profiles", handlers.CreateProfile(db))
	CACOnly.GET("/my_profile", handlers.GetMyProfile(db))
	CACOnly.GET("/my_projects", handlers.ListMyProjects(db))
	CACOnly.POST("/my_tokens", handlers.CreateToken(db))
	CACOnly.DELETE("/my_tokens/:token_id", handlers.DeleteToken(db))

	// Authenticated with Appkey only (routes only to be used by other components of the app)
	// Routes do not have /project/:project_id context and are typically authorized
	app.POST("/timeseries_measurements", handlers.CreateOrUpdateTimeseriesMeasurements(db))
	app.POST("/heartbeat", handlers.DoHeartbeat(db))

	// Heartbeat
	public.GET("/heartbeats", handlers.ListHeartbeats(db))
	public.GET("/heartbeat/latest", handlers.GetLatestHeartbeat(db))
	public.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{"status": "healthy"})
	})

	// Search
	public.GET("/search/:entity", handlers.Search(db))

	// Aware
	public.GET("/aware/parameters", handlers.ListAwareParameters(db))
	public.GET("/aware/data_acquisition_config", handlers.ListAwarePlatformParameterConfig(db))

	// AlertConfigs
	public.GET("/projects/:project_id/alert_configs", handlers.ListProjectAlertConfigs(db))
	public.GET("/projects/:project_id/instruments/:instrument_id/alert_configs", handlers.ListInstrumentAlertConfigs(db))
	public.GET("/projects/:project_id/alert_configs/:alert_config_id", handlers.GetAlertConfig(db))
	private.POST("/projects/:project_id/alert_configs", handlers.CreateAlertConfig(db))
	private.PUT("/projects/:project_id/alert_configs/:alert_config_id", handlers.UpdateAlertConfig(db))
	private.DELETE("/projects/:project_id/alert_configs/:alert_config_id", handlers.DeleteAlertConfig(db))

	// Submittals
	public.GET("/projects/:project_id/submittals", handlers.ListProjectSubmittals(db))
	public.GET("/instruments/:instrument_id/submittals", handlers.ListInstrumentSubmittals(db))
	public.GET("/alert_configs/:alert_config_id/submittals", handlers.ListAlertConfigSubmittals(db))
	private.PUT("/submittals/:submittal_id/verify_missing", handlers.VerifyMissingSubmittal(db))
	private.PUT("/alert_configs/:alert_config_id/submittals/verify_missing", handlers.VerifyMissingAlertConfigSubmittals(db))

	// District Rollup
	public.GET("/projects/:project_id/district_rollup/evaluation_submittals", handlers.ListProjectEvaluationDistrictRollup(db))
	public.GET("/projects/:project_id/district_rollup/measurement_submittals", handlers.ListProjectMeasurementDistrictRollup(db))

	// Alerts
	public.GET("/projects/:project_id/instruments/:instrument_id/alerts", handlers.ListAlertsForInstrument(db))
	private.GET("/my_alerts", handlers.ListMyAlerts(db))
	private.POST("/my_alerts/:alert_id/read", handlers.DoAlertRead(db))
	private.POST("/my_alerts/:alert_id/unread", handlers.DoAlertUnread(db))

	// AlertSubscriptions
	private.GET("/my_alert_subscriptions", handlers.ListMyAlertSubscriptions(db)) // Private because token required to determine user (i.e. who is "me")
	private.POST("/projects/:project_id/instruments/:instrument_id/alert_configs/:alert_config_id/subscribe", handlers.SubscribeProfileToAlerts(db))
	private.POST("/projects/:project_id/instruments/:instrument_id/alert_configs/:alert_config_id/unsubscribe", handlers.UnsubscribeProfileToAlerts(db))
	private.PUT("/alert_subscriptions/:alert_subscription_id", handlers.UpdateMyAlertSubscription(db))

	// Email Autocomplete
	public.GET("/email_autocomplete", handlers.ListEmailAutocomplete(db))

	// Evaluations
	public.GET("/projects/:project_id/evaluations", handlers.ListProjectEvaluations(db))
	public.GET("/projects/:project_id/instruments/:instrument_id/evaluations", handlers.ListInstrumentEvaluations(db))
	public.GET("/projects/:project_id/evaluations/:evaluation_id", handlers.GetEvaluation(db))
	private.POST("/projects/:project_id/evaluations", handlers.CreateEvaluation(db))
	private.PUT("/projects/:project_id/evaluations/:evaluation_id", handlers.UpdateEvaluation(db))
	private.DELETE("/projects/:project_id/evaluations/:evaluation_id", handlers.DeleteEvaluation(db))

	// Projects
	public.GET("/districts", handlers.ListDistricts(db))
	public.GET("/projects", handlers.ListProjects(db))
	public.GET("/projects/:project_id", handlers.GetProject(db))
	public.GET("/projects/count", handlers.GetProjectCount(db))
	public.GET("/projects/:project_id/instruments", handlers.ListProjectInstruments(db))
	public.GET("/projects/:project_id/instruments/names", handlers.ListProjectInstrumentNames(db))
	public.GET("/projects/:project_id/instrument_groups", handlers.ListProjectInstrumentGroups(db))
	private.POST("/projects", handlers.CreateProjectBulk(db), middleware.IsApplicationAdmin)
	private.PUT("/projects/:project_id", handlers.UpdateProject(db))
	private.DELETE("/projects/:project_id", handlers.DeleteFlagProject(db))

	// Project Membership
	// // list project memberships
	private.GET("/projects/:project_id/members", handlers.ListProjectMembers(db))
	// add role to a user
	private.POST("/projects/:project_id/members/:profile_id/roles/:role_id", handlers.AddProjectMemberRole(db))
	// remove role from a user
	private.DELETE("/projects/:project_id/members/:profile_id/roles/:role_id", handlers.RemoveProjectMemberRole(db))

	// Project Timeseries
	private.POST("/projects/:project_id/timeseries/:timeseries_id", handlers.CreateProjectTimeseries(db))
	private.DELETE("/projects/:project_id/timeseries/:timeseries_id", handlers.DeleteProjectTimeseries(db))

	// Instruments
	public.GET("/instruments", handlers.ListInstruments(db))
	public.GET("/instruments/count", handlers.GetInstrumentCount(db))
	public.GET("/instruments/:instrument_id", handlers.GetInstrument(db))
	public.GET("/instruments/:instrument_id/timeseries_measurements", handlers.ListTimeseriesMeasurementsByInstrument(db))
	private.POST("/projects/:project_id/instruments", handlers.CreateInstruments(db))
	// TODO: Remove endpoint POST /instruments (no project context)
	private.POST("/instruments", handlers.CreateInstruments(db))
	private.PUT("/projects/:project_id/instruments/:instrument_id", handlers.UpdateInstrument(db))
	private.PUT("/projects/:project_id/instruments/:instrument_id/geometry", handlers.UpdateInstrumentGeometry(db))
	private.DELETE("/projects/:project_id/instruments/:instrument_id", handlers.DeleteFlagInstrument(db))

	// Instrument Groups
	public.GET("/instrument_groups", handlers.ListInstrumentGroups(db))
	public.GET("/instrument_groups/:instrument_group_id", handlers.GetInstrumentGroup(db))
	public.GET("/instrument_groups/:instrument_group_id/instruments", handlers.ListInstrumentGroupInstruments(db))
	public.GET("/instrument_groups/:instrument_group_id/timeseries", handlers.ListInstrumentGroupTimeseries(db))
	public.GET("/instrument_groups/:instrument_group_id/timeseries_measurements", handlers.ListTimeseriesMeasurementsByInstrumentGroup(db))
	private.POST("/instrument_groups", handlers.CreateInstrumentGroup(db))
	private.PUT("/instrument_groups/:instrument_group_id", handlers.UpdateInstrumentGroup(db))
	private.DELETE("/instrument_groups/:instrument_group_id", handlers.DeleteFlagInstrumentGroup(db))
	// Add or Remove instrument from Instrument Group
	private.POST("/instrument_groups/:instrument_group_id/instruments", handlers.CreateInstrumentGroupInstruments(db))
	private.DELETE("/instrument_groups/:instrument_group_id/instruments/:instrument_id", handlers.DeleteInstrumentGroupInstruments(db))

	// Instrument Constants (same as a timeseries in structure/payload)
	public.GET("/projects/:project_id/instruments/:instrument_id/constants", handlers.ListInstrumentConstants(db))
	private.POST("/projects/:project_id/instruments/:instrument_id/constants", handlers.CreateInstrumentConstants(db))
	private.PUT("/projects/:project_id/instruments/:instrument_id/constants/:timeseries_id", handlers.UpdateTimeseries(db))
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
	public.GET("/projects/:project_id/timeseries", handlers.ListProjectTimeseries(db))
	public.GET("/projects/:project_id/instruments/:instrument_id/timeseries", handlers.ListInstrumentTimeseries(db))
	public.GET("/timeseries", handlers.ListTimeseries(db))
	public.GET("/timeseries/:timeseries_id", handlers.GetTimeseries(db))
	public.GET("/instruments/:instrument_id/timeseries/:timeseries_id", handlers.GetTimeseries(db))
	public.GET("/timeseries/:timeseries_id/measurements", handlers.ListTimeseriesMeasurementsByTimeseries(db))
	private.DELETE("/timeseries/:timeseries_id/measurements", handlers.DeleteTimeserieMeasurements(db))
	public.GET("/timeseries/:timeseries_id/inclinometer_measurements", handlers.ListInclinometerMeasurements(db))
	private.DELETE("/timeseries/:timeseries_id/inclinometer_measurements", handlers.DeleteInclinometerMeasurements(db))
	public.GET("/instruments/:instrument_id/timeseries/:timeseries_id/measurements", handlers.ListTimeseriesMeasurementsByTimeseries(db))
	// TODO: Delete timeseries endpoints without project context in URL
	private.POST("/timeseries", handlers.CreateTimeseries(db))
	private.PUT("/timeseries/:timeseries_id", handlers.UpdateTimeseries(db))
	private.DELETE("/timeseries/:timeseries_id", handlers.DeleteTimeseries(db))
	private.POST("/projects/:project_id/timeseries_measurements", handlers.CreateOrUpdateProjectTimeseriesMeasurements(db))
	private.PUT("/projects/:project_id/timeseries_measurements", handlers.UpdateTimeseriesMeasurements(db))
	private.POST("/projects/:project_id/inclinometer_measurements", handlers.CreateOrUpdateProjectInclinometerMeasurements(db))

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

	// Plotting Configurations
	public.GET("/projects/:project_id/plot_configurations", handlers.ListPlotConfigurations(db))
	public.GET("/projects/:project_id/plot_configurations/:plot_configuration_id", handlers.GetPlotConfiguration(db))
	private.POST("/projects/:project_id/plot_configurations", handlers.CreatePlotConfiguration(db))
	private.PUT("/projects/:project_id/plot_configurations/:plot_configuration_id", handlers.UpdatePlotConfiguration(db))
	private.DELETE("/projects/:project_id/plot_configurations/:plot_configuration_id", handlers.DeletePlotConfiguration(db))

	// Formulas == Calculations
	public.GET("/formulas", handlers.GetInstrumentCalculations(db))
	private.POST("/formulas", handlers.CreateCalculation(db))
	private.PUT("/formulas/:formula_id", handlers.UpdateCalculation(db))
	private.DELETE("/formulas/:formula_id", handlers.DeleteCalculation(db))

	// Misc
	public.GET("/domains", handlers.GetDomains(db))
	public.POST("/explorer", handlers.ListTimeseriesMeasurementsExplorer(db))
	public.POST("/inclinometer_explorer", handlers.PostInclinometerExplorer(db))
	public.GET("/home", handlers.GetHome(db))
	public.GET("/units", handlers.ListUnits(db))

	// OpenDCS Configuration
	public.GET("/opendcs/sites", handlers.ListOpendcsSites(db))

	// Data Logger
	private.GET("/dataloggers", handlers.ListDataLoggers(db))
	private.POST("/datalogger", handlers.CreateDataLogger(db))
	private.GET("/datalogger/:datalogger_id", handlers.GetDataLogger(db))
	private.PUT("/datalogger/:datalogger_id", handlers.UpdateDataLogger(db))
	private.PUT("/datalogger/:datalogger_id/key", handlers.CycleDataLoggerKey(db))
	private.DELETE("/datalogger/:datalogger_id", handlers.DeleteDataLogger(db))

	// Data Logger Equivalency Table
	private.GET("/datalogger/:datalogger_id/equivalency_table", handlers.GetEquivalencyTable(db))
	private.POST("/datalogger/:datalogger_id/equivalency_table", handlers.CreateEquivalencyTable(db))
	private.PUT("/datalogger/:datalogger_id/equivalency_table", handlers.UpdateEquivalencyTable(db))
	private.DELETE("/datalogger/:datalogger_id/equivalency_table", handlers.DeleteEquivalencyTable(db))
	private.DELETE("/datalogger/:datalogger_id/equivalency_table/row", handlers.DeleteEquivalencyTableRow(db))

	// Data Logger Preview
	private.GET("/datalogger/:datalogger_id/preview", handlers.GetDataLoggerPreview(db))

	if cfg.LambdaContext {
		log.Print("starting server; Running On AWS LAMBDA")
		log.Fatal(gateway.ListenAndServe("localhost:3030", e))
	} else {
		log.Print("starting server")
		log.Fatal(http.ListenAndServe(":80", e))
	}
}
