package handler

import (
	"net/http"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/cloud"
	"github.com/USACE/instrumentation-api/api/internal/config"
	"github.com/USACE/instrumentation-api/api/internal/middleware"
	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/USACE/instrumentation-api/api/internal/service"
)

func newHttpClient() *http.Client {
	return &http.Client{
		Timeout: time.Second * 60,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return nil
		},
	}
}

type ApiHandler struct {
	Middleware                     middleware.Middleware
	BlobService                    cloud.Blob
	AlertService                   service.AlertService
	AlertConfigService             service.AlertConfigService
	AlertSubscriptionService       service.AlertSubscriptionService
	EmailAutocompleteService       service.EmailAutocompleteService
	AwareParameterService          service.AwareParameterService
	CollectionGroupService         service.CollectionGroupService
	DataloggerService              service.DataloggerService
	DataloggerTelemetryService     service.DataloggerTelemetryService
	DistrictRollupService          service.DistrictRollupService
	DomainService                  service.DomainService
	EquivalencyTableService        service.EquivalencyTableService
	EvaluationService              service.EvaluationService
	HeartbeatService               service.HeartbeatService
	HomeService                    service.HomeService
	InstrumentService              service.InstrumentService
	InstrumentAssignService        service.InstrumentAssignService
	InstrumentConstantService      service.InstrumentConstantService
	InstrumentGroupService         service.InstrumentGroupService
	InstrumentNoteService          service.InstrumentNoteService
	InstrumentStatusService        service.InstrumentStatusService
	IpiInstrumentService           service.IpiInstrumentService
	MeasurementService             service.MeasurementService
	InclinometerMeasurementService service.InclinometerMeasurementService
	OpendcsService                 service.OpendcsService
	PlotConfigService              service.PlotConfigService
	ProfileService                 service.ProfileService
	ProjectRoleService             service.ProjectRoleService
	ProjectService                 service.ProjectService
	ReportConfigService            service.ReportConfigService
	SaaInstrumentService           service.SaaInstrumentService
	SubmittalService               service.SubmittalService
	TimeseriesService              service.TimeseriesService
	TimeseriesCwmsService          service.TimeseriesCwmsService
	CalculatedTimeseriesService    service.CalculatedTimeseriesService
	ProcessTimeseriesService       service.ProcessTimeseriesService
	UnitService                    service.UnitService
}

func NewApi(cfg *config.ApiConfig) *ApiHandler {
	db := model.NewDatabase(&cfg.DBConfig)
	q := db.Queries()
	ps := cloud.NewSQSPubsub(&cfg.AWSSQSConfig)

	profileService := service.NewProfileService(db, q)
	projectRoleService := service.NewProjectRoleService(db, q)
	dataloggerTelemetryService := service.NewDataloggerTelemetryService(db, q)
	mw := middleware.NewMiddleware(&cfg.ServerConfig, profileService, projectRoleService, dataloggerTelemetryService)

	return &ApiHandler{
		Middleware:                     mw,
		BlobService:                    cloud.NewS3Blob(&cfg.AWSS3Config, "/instrumentation", cfg.RoutePrefix),
		AlertService:                   service.NewAlertService(db, q),
		AlertConfigService:             service.NewAlertConfigService(db, q),
		AlertSubscriptionService:       service.NewAlertSubscriptionService(db, q),
		EmailAutocompleteService:       service.NewEmailAutocompleteService(db, q),
		AwareParameterService:          service.NewAwareParameterService(db, q),
		CollectionGroupService:         service.NewCollectionGroupService(db, q),
		DataloggerService:              service.NewDataloggerService(db, q),
		DataloggerTelemetryService:     dataloggerTelemetryService,
		DistrictRollupService:          service.NewDistrictRollupService(db, q),
		DomainService:                  service.NewDomainService(db, q),
		EquivalencyTableService:        service.NewEquivalencyTableService(db, q),
		EvaluationService:              service.NewEvaluationService(db, q),
		HeartbeatService:               service.NewHeartbeatService(db, q),
		HomeService:                    service.NewHomeService(db, q),
		InstrumentService:              service.NewInstrumentService(db, q),
		InstrumentAssignService:        service.NewInstrumentAssignService(db, q),
		InstrumentConstantService:      service.NewInstrumentConstantService(db, q),
		InstrumentGroupService:         service.NewInstrumentGroupService(db, q),
		InstrumentNoteService:          service.NewInstrumentNoteService(db, q),
		InstrumentStatusService:        service.NewInstrumentStatusService(db, q),
		IpiInstrumentService:           service.NewIpiInstrumentService(db, q),
		MeasurementService:             service.NewMeasurementService(db, q),
		InclinometerMeasurementService: service.NewInclinometerMeasurementService(db, q),
		OpendcsService:                 service.NewOpendcsService(db, q),
		PlotConfigService:              service.NewPlotConfigService(db, q),
		ProfileService:                 profileService,
		ProjectRoleService:             service.NewProjectRoleService(db, q),
		ProjectService:                 service.NewProjectService(db, q),
		ReportConfigService:            service.NewReportConfigService(db, q, ps, cfg.AuthJWTMocked),
		SaaInstrumentService:           service.NewSaaInstrumentService(db, q),
		SubmittalService:               service.NewSubmittalService(db, q),
		TimeseriesService:              service.NewTimeseriesService(db, q),
		TimeseriesCwmsService:          service.NewTimeseriesCwmsService(db, q),
		CalculatedTimeseriesService:    service.NewCalculatedTimeseriesService(db, q),
		ProcessTimeseriesService:       service.NewProcessTimeseriesService(db, q),
		UnitService:                    service.NewUnitService(db, q),
	}
}

type TelemetryHandler struct {
	Middleware                 middleware.Middleware
	DataloggerService          service.DataloggerService
	DataloggerTelemetryService service.DataloggerTelemetryService
	EquivalencyTableService    service.EquivalencyTableService
	MeasurementService         service.MeasurementService
}

func NewTelemetry(cfg *config.TelemetryConfig) *TelemetryHandler {
	db := model.NewDatabase(&cfg.DBConfig)
	q := db.Queries()

	profileService := service.NewProfileService(db, q)
	projectRoleService := service.NewProjectRoleService(db, q)
	dataloggerTelemetryService := service.NewDataloggerTelemetryService(db, q)
	mw := middleware.NewMiddleware(&cfg.ServerConfig, profileService, projectRoleService, dataloggerTelemetryService)

	return &TelemetryHandler{
		Middleware:                 mw,
		DataloggerService:          service.NewDataloggerService(db, q),
		DataloggerTelemetryService: dataloggerTelemetryService,
		EquivalencyTableService:    service.NewEquivalencyTableService(db, q),
		MeasurementService:         service.NewMeasurementService(db, q),
	}
}

type AlertCheckHandler struct {
	AlertCheckService service.AlertCheckService
}

func NewAlertCheck(cfg *config.AlertCheckConfig) *AlertCheckHandler {
	db := model.NewDatabase(&cfg.DBConfig)
	q := db.Queries()

	return &AlertCheckHandler{
		AlertCheckService: service.NewAlertCheckService(db, q, cfg),
	}
}

type DcsLoaderHandler struct {
	PubsubService    cloud.Pubsub
	DcsLoaderService service.DcsLoaderService
}

func NewDcsLoader(cfg *config.DcsLoaderConfig) *DcsLoaderHandler {
	s3Blob := cloud.NewS3Blob(&cfg.AWSS3Config, "", "")
	ps := cloud.NewSQSPubsub(&cfg.AWSSQSConfig).WithBlob(s3Blob)
	apiClient := newHttpClient()

	return &DcsLoaderHandler{
		PubsubService:    ps,
		DcsLoaderService: service.NewDcsLoaderService(apiClient, cfg),
	}
}
