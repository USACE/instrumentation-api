package server

import (
	"net/http"

	_ "github.com/USACE/instrumentation-api/api/docs"
	"github.com/USACE/instrumentation-api/api/internal/config"
	"github.com/USACE/instrumentation-api/api/internal/handler"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type ApiServer struct {
	e *echo.Echo
	apiGroups
}

type apiGroups struct {
	public  *echo.Group
	private *echo.Group
	cacOnly *echo.Group
	app     *echo.Group
}

func NewApiServer(cfg *config.ApiConfig, h *handler.ApiHandler) *ApiServer {
	e := echo.New()
	mw := h.Middleware

	e.Use(mw.CORS, mw.GZIP)

	public := e.Group(cfg.RoutePrefix)
	private := e.Group(cfg.RoutePrefix) // cac or token (KeyAuth)
	cacOnly := e.Group(cfg.RoutePrefix)
	app := e.Group(cfg.RoutePrefix)

	private.Use(mw.JWTSkipIfKey, mw.KeyAuth, mw.EDIPI, mw.AttachProfile)
	cacOnly.Use(mw.JWT, mw.EDIPI, mw.CACOnly)
	app.Use(mw.AppKeyAuth)

	server := &ApiServer{e, apiGroups{
		public,
		private,
		cacOnly,
		app,
	}}

	public.GET("/swagger/*", echoSwagger.WrapHandler)

	server.RegisterRoutes(h)
	return server
}

func (server *ApiServer) Start() error {
	return http.ListenAndServe(":80", server.e)
}

func (server *ApiServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	server.e.ServeHTTP(w, r)
}

func (r *ApiServer) RegisterRoutes(h *handler.ApiHandler) {
	// Alert
	r.public.GET("/projects/:project_id/instruments/:instrument_id/alerts", h.ListAlertsForInstrument)
	r.private.GET("/my_alerts", h.ListMyAlerts)
	r.private.POST("/my_alerts/:alert_id/read", h.DoAlertRead)
	r.private.POST("/my_alerts/:alert_id/unread", h.DoAlertUnread)

	//AlertConfig
	r.public.GET("/projects/:project_id/alert_configs", h.GetAllAlertConfigsForProject)
	r.public.GET("/projects/:project_id/instruments/:instrument_id/alert_configs", h.ListInstrumentAlertConfigs)
	r.public.GET("/projects/:project_id/alert_configs/:alert_config_id", h.GetAlertConfig)
	r.private.POST("/projects/:project_id/alert_configs", h.CreateAlertConfig)
	r.private.PUT("/projects/:project_id/alert_configs/:alert_config_id", h.UpdateAlertConfig)
	r.private.DELETE("/projects/:project_id/alert_configs/:alert_config_id", h.DeleteAlertConfig)

	// AlertSubscription
	r.private.GET("/my_alert_subscriptions", h.ListMyAlertSubscriptions) // Private because token required to determine user (i.e. who is "me")
	r.private.POST("/projects/:project_id/instruments/:instrument_id/alert_configs/:alert_config_id/subscribe", h.SubscribeProfileToAlerts)
	r.private.POST("/projects/:project_id/instruments/:instrument_id/alert_configs/:alert_config_id/unsubscribe", h.UnsubscribeProfileToAlerts)
	r.private.PUT("/alert_subscriptions/:alert_subscription_id", h.UpdateMyAlertSubscription)

	// Autocomplete
	r.public.GET("/email_autocomplete", h.ListEmailAutocomplete)

	// AwareParameters
	r.public.GET("/aware/parameters", h.ListAwareParameters)
	r.public.GET("/aware/data_acquisition_config", h.ListAwarePlatformParameterConfig)

	// CollectionGroup
	r.public.GET("/projects/:project_id/collection_groups", h.ListCollectionGroups)
	r.public.GET("/projects/:project_id/collection_groups/:collection_group_id", h.GetCollectionGroupDetails)
	r.private.POST("/projects/:project_id/collection_groups", h.CreateCollectionGroup)
	r.private.PUT("/projects/:project_id/collection_groups/:collection_group_id", h.UpdateCollectionGroup)
	r.private.DELETE("/projects/:project_id/collection_groups/:collection_group_id", h.DeleteCollectionGroup)
	r.private.POST("/projects/:project_id/collection_groups/:collection_group_id/timeseries/:timeseries_id", h.AddTimeseriesToCollectionGroup)
	r.private.DELETE("/projects/:project_id/collection_groups/:collection_group_id/timeseries/:timeseries_id", h.RemoveTimeseriesFromCollectionGroup)

	// Datalogger
	r.private.GET("/dataloggers", h.ListDataloggers)
	r.private.POST("/datalogger", h.CreateDatalogger)
	r.private.GET("/datalogger/:datalogger_id", h.GetDatalogger)
	r.private.PUT("/datalogger/:datalogger_id", h.UpdateDatalogger)
	r.private.PUT("/datalogger/:datalogger_id/key", h.CycleDataloggerKey)
	r.private.DELETE("/datalogger/:datalogger_id", h.DeleteDatalogger)
	r.private.GET("/datalogger/:datalogger_id/table/:datalogger_table_id/preview", h.GetDataloggerTablePreview)
	r.private.PUT("/datalogger/:datalogger_id/table/:datalogger_table_id/name ", h.ResetDataloggerTableName)

	// DistrictRollup
	r.public.GET("/projects/:project_id/district_rollup/evaluation_submittals", h.ListProjectEvaluationDistrictRollup)
	r.public.GET("/projects/:project_id/district_rollup/measurement_submittals", h.ListProjectMeasurementDistrictRollup)

	// Domain
	r.public.GET("/domains", h.GetDomains)

	// EquivalencyTable
	r.private.GET("/datalogger/:datalogger_id/table/:datalogger_table_id/equivalency_table", h.GetEquivalencyTable)
	r.private.POST("/datalogger/:datalogger_id/table/:datalogger_table_id/equivalency_table", h.CreateEquivalencyTable)
	r.private.PUT("/datalogger/:datalogger_id/table/:datalogger_table_id/equivalency_table", h.UpdateEquivalencyTable)
	r.private.DELETE("/datalogger/:datalogger_id/table/:datalogger_table_id/equivalency_table", h.DeleteEquivalencyTable)
	r.private.DELETE("/datalogger/:datalogger_id/table/:datalogger_table_id/equivalency_table/row/:row_id", h.DeleteEquivalencyTableRow)

	// Evaluation
	r.public.GET("/projects/:project_id/evaluations", h.ListProjectEvaluations)
	r.public.GET("/projects/:project_id/instruments/:instrument_id/evaluations", h.ListInstrumentEvaluations)
	r.public.GET("/projects/:project_id/evaluations/:evaluation_id", h.GetEvaluation)
	r.private.POST("/projects/:project_id/evaluations", h.CreateEvaluation)
	r.private.PUT("/projects/:project_id/evaluations/:evaluation_id", h.UpdateEvaluation)
	r.private.DELETE("/projects/:project_id/evaluations/:evaluation_id", h.DeleteEvaluation)

	// Explorer
	r.public.POST("/explorer", h.ListTimeseriesMeasurementsExplorer)
	r.public.POST("/inclinometer_explorer", h.ListInclinometerTimeseriesMeasurementsExplorer)

	// Heartbeat
	r.public.GET("/health", h.Healthcheck)
	r.app.POST("/heartbeat", h.DoHeartbeat)
	r.public.GET("/heartbeats", h.ListHeartbeats)
	r.public.GET("/heartbeat/latest", h.GetLatestHeartbeat)

	// Home
	r.public.GET("/home", h.GetHome)

	// Instrument
	r.public.GET("/instruments", h.ListInstruments)
	r.public.GET("/instruments/count", h.GetInstrumentCount)
	r.public.GET("/instruments/:instrument_id", h.GetInstrument)
	r.public.GET("/instruments/:instrument_id/timeseries_measurements", h.ListTimeseriesMeasurementsByInstrument)
	r.private.POST("/projects/:project_id/instruments", h.CreateInstruments)
	// TODO: Remove endpoint POST /instruments (no project context)
	r.private.POST("/instruments", h.CreateInstruments)
	r.private.PUT("/projects/:project_id/instruments/:instrument_id", h.UpdateInstrument)
	r.private.PUT("/projects/:project_id/instruments/:instrument_id/geometry", h.UpdateInstrumentGeometry)
	r.private.DELETE("/projects/:project_id/instruments/:instrument_id", h.DeleteFlagInstrument)

	// InstrumentConstant
	r.public.GET("/projects/:project_id/instruments/:instrument_id/constants", h.ListInstrumentConstants)
	r.private.POST("/projects/:project_id/instruments/:instrument_id/constants", h.CreateInstrumentConstants)
	r.private.PUT("/projects/:project_id/instruments/:instrument_id/constants/:timeseries_id", h.UpdateTimeseries)
	r.private.DELETE("/projects/:project_id/instruments/:instrument_id/constants/:timeseries_id", h.DeleteInstrumentConstant)

	// InstrumentGroup
	r.public.GET("/instrument_groups", h.ListInstrumentGroups)
	r.public.GET("/instrument_groups/:instrument_group_id", h.GetInstrumentGroup)
	r.public.GET("/instrument_groups/:instrument_group_id/instruments", h.ListInstrumentGroupInstruments)
	r.public.GET("/instrument_groups/:instrument_group_id/timeseries", h.ListInstrumentGroupTimeseries)
	r.public.GET("/instrument_groups/:instrument_group_id/timeseries_measurements", h.ListTimeseriesMeasurementsByInstrumentGroup)
	r.private.POST("/instrument_groups", h.CreateInstrumentGroup)
	r.private.PUT("/instrument_groups/:instrument_group_id", h.UpdateInstrumentGroup)
	r.private.DELETE("/instrument_groups/:instrument_group_id", h.DeleteFlagInstrumentGroup)
	r.private.POST("/instrument_groups/:instrument_group_id/instruments", h.CreateInstrumentGroupInstruments)
	r.private.DELETE("/instrument_groups/:instrument_group_id/instruments/:instrument_id", h.DeleteInstrumentGroupInstruments)

	// InstrumentNote
	r.public.GET("/instruments/notes", h.ListInstrumentNotes)
	r.public.GET("/instruments/notes/:note_id", h.GetInstrumentNote)
	r.public.GET("/instruments/:instrument_id/notes", h.ListInstrumentInstrumentNotes)
	r.public.GET("/instruments/:instrument_id/notes/:note_id", h.GetInstrumentNote)
	r.private.POST("/instruments/notes", h.CreateInstrumentNote)
	r.private.PUT("/instruments/notes/:note_id", h.UpdateInstrumentNote)
	r.private.DELETE("/instruments/notes/:note_id", h.DeleteInstrumentNote)
	r.private.PUT("/instruments/:instrument_id/notes/:note_id", h.UpdateInstrumentNote)
	r.private.DELETE("/instruments/:instrument_id/notes/:note_id", h.DeleteInstrumentNote)

	// InstrumentStatus
	r.public.GET("/instruments/:instrument_id/status", h.ListInstrumentStatus)
	r.public.GET("/instruments/:instrument_id/status/:status_id", h.GetInstrumentStatus)
	r.private.POST("/instruments/:instrument_id/status", h.CreateOrUpdateInstrumentStatus)
	r.private.DELETE("/instruments/:instrument_id/status/:status_id", h.DeleteInstrumentStatus)

	// IpiInstruemnt
	r.public.GET("/instruments/ipi/:instrument_id/segments", h.GetAllIpiSegmentsForInstrument)
	r.public.GET("/instruments/ipi/:instrument_id/measurements", h.GetIpiMeasurementsForInstrument)
	r.private.PUT("/instruments/ipi/:instrument_id/segments", h.UpdateIpiSegments)

	// Measurement
	r.private.POST("/projects/:project_id/timeseries_measurements", h.CreateOrUpdateProjectTimeseriesMeasurements)
	r.app.POST("/timeseries_measurements", h.CreateOrUpdateTimeseriesMeasurements)
	r.private.PUT("/projects/:project_id/timeseries_measurements", h.UpdateTimeseriesMeasurements)
	r.private.DELETE("/timeseries/:timeseries_id/measurements", h.DeleteTimeserieMeasurements)

	// InclinometerMeasurement
	r.public.GET("/timeseries/:timeseries_id/inclinometer_measurements", h.ListInclinometerMeasurements)
	r.private.POST("/projects/:project_id/inclinometer_measurements", h.CreateOrUpdateProjectInclinometerMeasurements)
	r.private.DELETE("/timeseries/:timeseries_id/inclinometer_measurements", h.DeleteInclinometerMeasurements)

	// Media
	r.public.GET("/projects/:project_slug/images/*", h.GetMedia)

	// Opendcs
	r.public.GET("/opendcs/sites", h.ListOpendcsSites)

	// PlotConfig
	r.public.GET("/projects/:project_id/plot_configurations", h.ListPlotConfigs)
	r.public.GET("/projects/:project_id/plot_configurations/:plot_configuration_id", h.GetPlotConfig)
	r.private.POST("/projects/:project_id/plot_configurations", h.CreatePlotConfig)
	r.private.PUT("/projects/:project_id/plot_configurations/:plot_configuration_id", h.UpdatePlotConfig)
	r.private.DELETE("/projects/:project_id/plot_configurations/:plot_configuration_id", h.DeletePlotConfig)

	// Profile
	r.cacOnly.POST("/profiles", h.CreateProfile)
	r.cacOnly.GET("/my_profile", h.GetMyProfile)
	r.cacOnly.GET("/my_projects", h.ListMyProjects)
	r.cacOnly.POST("/my_tokens", h.CreateToken)
	r.cacOnly.DELETE("/my_tokens/:token_id", h.DeleteToken)

	// ProjectRole
	r.private.GET("/projects/:project_id/members", h.ListProjectMembers)
	r.private.POST("/projects/:project_id/members/:profile_id/roles/:role_id", h.AddProjectMemberRole)
	r.private.DELETE("/projects/:project_id/members/:profile_id/roles/:role_id", h.RemoveProjectMemberRole)

	// Project
	r.public.GET("/districts", h.ListDistricts)
	r.public.GET("/projects", h.ListProjects)
	r.public.GET("/projects/:project_id", h.GetProject)
	r.public.GET("/projects/count", h.GetProjectCount)
	r.public.GET("/projects/:project_id/instruments", h.ListProjectInstruments)
	r.public.GET("/projects/:project_id/instruments/names", h.ListProjectInstrumentNames)
	r.public.GET("/projects/:project_id/instrument_groups", h.ListProjectInstrumentGroups)
	r.private.POST("/projects", h.CreateProjectBulk, h.Middleware.IsApplicationAdmin)
	r.private.PUT("/projects/:project_id", h.UpdateProject)
	r.private.DELETE("/projects/:project_id", h.DeleteFlagProject)
	r.private.POST("/projects/:project_id/timeseries/:timeseries_id", h.CreateProjectTimeseries)
	r.private.DELETE("/projects/:project_id/timeseries/:timeseries_id", h.DeleteProjectTimeseries)

	// Search
	r.public.GET("/search/:entity", h.Search)

	// SaaInstrument
	r.public.GET("/instruments/saa/:instrument_id/segments", h.GetAllSaaSegmentsForInstrument)
	r.public.GET("/instruments/saa/:instrument_id/measurements", h.GetSaaMeasurementsForInstrument)
	r.private.PUT("/instruments/saa/:instrument_id/segments", h.UpdateSaaSegments)

	// Submittal
	r.public.GET("/projects/:project_id/submittals", h.ListProjectSubmittals)
	r.public.GET("/instruments/:instrument_id/submittals", h.ListInstrumentSubmittals)
	r.public.GET("/alert_configs/:alert_config_id/submittals", h.ListAlertConfigSubmittals)
	r.private.PUT("/submittals/:submittal_id/verify_missing", h.VerifyMissingSubmittal)
	r.private.PUT("/alert_configs/:alert_config_id/submittals/verify_missing", h.VerifyMissingAlertConfigSubmittals)

	// Timeseries
	// TODO: Delete timeseries endpoints without project context in URL
	r.public.GET("/timeseries", h.ListTimeseries)
	r.public.GET("/timeseries/:timeseries_id", h.GetTimeseries)
	r.public.GET("/instruments/:instrument_id/timeseries/:timeseries_id", h.GetTimeseries)
	r.public.GET("/projects/:project_id/timeseries", h.ListProjectTimeseries)
	r.public.GET("/projects/:project_id/instruments/:instrument_id/timeseries", h.ListInstrumentTimeseries)
	r.private.POST("/timeseries", h.CreateTimeseries)
	r.private.PUT("/timeseries/:timeseries_id", h.UpdateTimeseries)
	r.private.DELETE("/timeseries/:timeseries_id", h.DeleteTimeseries)

	// CalculatedTimeseries
	r.public.GET("/formulas", h.GetInstrumentCalculations)
	r.private.POST("/formulas", h.CreateCalculation)
	// TODO: This PUT should really be a PATCH to conform to the REST spec
	// Will need to coordinate this with the web client
	r.private.PUT("/formulas/:formula_id", h.UpdateCalculation)
	r.private.DELETE("/formulas/:formula_id", h.DeleteCalculation)

	// ProcessTimeseries
	r.public.GET("/timeseries/:timeseries_id/measurements", h.ListTimeseriesMeasurementsByTimeseries)
	r.public.GET("/instruments/:instrument_id/timeseries/:timeseries_id/measurements", h.ListTimeseriesMeasurementsByTimeseries)

	// Unit
	r.public.GET("/units", h.ListUnits)
}
