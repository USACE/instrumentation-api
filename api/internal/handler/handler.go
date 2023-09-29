package handler

import (
	"github.com/USACE/instrumentation-api/api/internal/config"
	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/USACE/instrumentation-api/api/internal/store"
	"github.com/labstack/echo/v4"
)

type router struct {
	e *echo.Echo
	g groups
}

type groups struct {
	public  *echo.Group
	private *echo.Group
	cacOnly *echo.Group
}

type ApiRouter struct {
	*router
}

type ApiHandler struct {
	// s3mediaStore         MediaStoreHandler
	AlertStore                   store.AlertStore
	AlertConfigStore             store.AlertConfigStore
	AlertSubscriptionStore       store.AlertSubscriptionStore
	EmailAutocompleteStore       store.EmailAutocompleteStore
	AwareParameterStore          store.AwareParameterStore
	CollectionGroupStore         store.CollectionGroupStore
	DataloggerStore              store.DataloggerStore
	DistrictRollupStore          store.DistrictRollupStore
	DomainStore                  store.DomainStore
	EquivalencyTableStore        store.EquivalencyTableStore
	EvaluationStore              store.EvaluationStore
	HeartbeatStore               store.HeartbeatStore
	HomeStore                    store.HomeStore
	InstrumentStore              store.InstrumentStore
	InstrumentConstantStore      store.InstrumentConstantStore
	InstrumentGroupStore         store.InstrumentGroupStore
	InstrumentNoteStore          store.InstrumentNoteStore
	InstrumentStatusStore        store.InstrumentStatusStore
	MeasurementStore             store.MeasurementStore
	InclinometerMeasurementStore store.InclinometerMeasurementStore
	OpendcsStore                 store.OpendcsStore
	PlotConfigStore              store.PlotConfigStore
	ProfileStore                 store.ProfileStore
	ProjectRoleStore             store.ProjectRoleStore
	ProjectStore                 store.ProjectStore
	SubmittalStore               store.SubmittalStore
	TimeseriesStore              store.TimeseriesStore
	CalculatedTimeseriesStore    store.CalculatedTimeseriesStore
	ProcessTimeseriesStore       store.ProcessTimeseriesStore
	UnitStore                    store.UnitStore
}

func NewApi(cfg *config.ApiConfig) *ApiHandler {
	db := model.NewDatabase(&cfg.DBConfig)
	q := db.Queries()

	// mediaStore := NewMediaStore(cfg)

	return &ApiHandler{
		// s3mediaStore:         NewS3MediaStore(mediaStore),
		AlertStore:                   store.NewAlertStore(db, q),
		AlertConfigStore:             store.NewAlertConfigStore(db, q),
		AlertSubscriptionStore:       store.NewAlertSubscriptionStore(db, q),
		EmailAutocompleteStore:       store.NewEmailAutocompleteStore(db, q),
		AwareParameterStore:          store.NewAwareParameterStore(db, q),
		CollectionGroupStore:         store.NewCollectionGroupStore(db, q),
		DataloggerStore:              store.NewDataloggerStore(db, q),
		DistrictRollupStore:          store.NewDistrictRollupStore(db, q),
		DomainStore:                  store.NewDomainStore(db, q),
		EquivalencyTableStore:        store.NewEquivalencyTableStore(db, q),
		EvaluationStore:              store.NewEvaluationStore(db, q),
		HeartbeatStore:               store.NewHeartbeatStore(db, q),
		HomeStore:                    store.NewHomeStore(db, q),
		InstrumentStore:              store.NewInstrumentStore(db, q),
		InstrumentConstantStore:      store.NewInstrumentConstantStore(db, q),
		InstrumentGroupStore:         store.NewInstrumentGroupStore(db, q),
		InstrumentNoteStore:          store.NewInstrumentNoteStore(db, q),
		InstrumentStatusStore:        store.NewInstrumentStatusStore(db, q),
		MeasurementStore:             store.NewMeasurementStore(db, q),
		InclinometerMeasurementStore: store.NewInclinometerMeasurementStore(db, q),
		OpendcsStore:                 store.NewOpendcsStore(db, q),
		PlotConfigStore:              store.NewPlotConfigStore(db, q),
		ProfileStore:                 store.NewProfileStore(db, q),
		ProjectRoleStore:             store.NewProjectRoleStore(db, q),
		ProjectStore:                 store.NewProjectStore(db, q),
		SubmittalStore:               store.NewSubmittalStore(db, q),
		TimeseriesStore:              store.NewTimeseriesStore(db, q),
		CalculatedTimeseriesStore:    store.NewCalculatedTimeseriesStore(db, q),
		ProcessTimeseriesStore:       store.NewProcessTimeseriesStore(db, q),
		UnitStore:                    store.NewUnitStore(db, q),
	}
}

type TelemetryRouter struct {
	*router
}

type TelemetryHandler struct {
	DataloggerStore          store.DataloggerStore
	DataloggerTelemetryStore store.DataloggerTelemetryStore
	EquivalencyTableStore    store.EquivalencyTableStore
	MeasurementStore         store.MeasurementStore
}

func NewTelemetry(cfg *config.TelemetryConfig) *TelemetryHandler {
	db := model.NewDatabase(&cfg.DBConfig)
	q := db.Queries()

	return &TelemetryHandler{
		DataloggerStore:          store.NewDataloggerStore(db, q),
		DataloggerTelemetryStore: store.NewDataloggerTelemetryStore(db, q),
		EquivalencyTableStore:    store.NewEquivalencyTableStore(db, q),
		MeasurementStore:         store.NewMeasurementStore(db, q),
	}
}

type AlertCheckHandler struct {
	AlertCheckStore store.AlertCheckStore
}

func NewAlertCheck(cfg *config.AlertCheckConfig) *AlertCheckHandler {
	db := model.NewDatabase(&cfg.DBConfig)
	q := db.Queries()

	return &AlertCheckHandler{
		AlertCheckStore: store.NewAlertCheckStore(db, q, cfg),
	}
}

type DcsLoaderHandler struct {
	TimeseriesStore  store.TimeseriesStore
	MeasurementStore store.MeasurementStore
}

func NewDcsLoader(cfg *config.DcsLoaderConfig) *DcsLoaderHandler {
	// inject aws services

	return &DcsLoaderHandler{}
}
