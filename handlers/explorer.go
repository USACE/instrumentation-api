package handlers

import (
	"net/http"
	"time"

	"github.com/USACE/instrumentation-api/models"
	"github.com/USACE/instrumentation-api/timeseries"
	ts "github.com/USACE/instrumentation-api/timeseries"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

// ExplorerRow is used for sql scanning
type ExplorerRow struct {
	InstrumentID uuid.UUID `db:"instrument_id"`
	TimeseriesID uuid.UUID `db:"timeseries_id"`
	ParameterID  uuid.UUID `db:"parameter_id"`
	UnitID       uuid.UUID `db:"unit_id"`
	Time         time.Time `db:"time"`
	Value        float32   `db:"value"`
}

// Filter encapsulates SQL query filters from a request that are used to build SQL
type Filter struct {
	InstrumentID []uuid.UUID
	ParameterID  []uuid.UUID
	TimeWindow   timeseries.TimeWindow
}

// PostExplorer retrieves timeseries information for the explorer app component
func PostExplorer(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		// Filters used in SQL Query
		var f Filter

		// Instrument IDs from POST
		if err := (&echo.DefaultBinder{}).BindBody(c, &f.InstrumentID); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		// Time Window From POST
		if err := (&echo.DefaultBinder{}).BindQueryParams(c, &f.TimeWindow); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		// If after and before are not provided; Return last 7 days of data from current time
		if (f.TimeWindow.Before == time.Time{} && f.TimeWindow.After == time.Time{}) {
			f.TimeWindow.Before = time.Now()
			f.TimeWindow.After = f.TimeWindow.Before.AddDate(0, 0, -7)
		}

		// Get Stored And Computed Timeseries With Measurements
		interval := time.Hour // Set to 1 Hour; TODO - do not hard-code interval
		tt, err := models.ComputedTimeseries(db, f.InstrumentID, &f.TimeWindow, &interval)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		// Convert Rows to Response
		response, err := explorerResponseFactory(tt)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, response)
	}
}

func PostInclinometerExplorer(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		// Filters used in SQL Query
		var f Filter

		// Instrument IDs from POST
		if err := (&echo.DefaultBinder{}).BindBody(c, &f.InstrumentID); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		// Time Window From POST
		if err := (&echo.DefaultBinder{}).BindQueryParams(c, &f.TimeWindow); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		// If after and before are not provided; Return last 7 days of data from current time
		if (f.TimeWindow.Before == time.Time{} && f.TimeWindow.After == time.Time{}) {
			f.TimeWindow.Before = time.Now()
			f.TimeWindow.After = f.TimeWindow.Before.AddDate(0, 0, -7)
		}

		// Get Stored And Computed Timeseries With Measurements
		interval := time.Hour // Set to 1 Hour; TODO - do not hard-code interval
		tt, err := models.ComputedInclinometerTimeseries(db, f.InstrumentID, &f.TimeWindow, &interval)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		// Convert Rows to Response
		response, err := explorerInclinometerResponseFactory(tt)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, response)
	}
}

// explorerResponseFactory returns the explorer-specific JSON response format
func explorerResponseFactory(tt []models.Timeseries) (map[uuid.UUID][]ts.MeasurementCollectionLean, error) {

	response := make(map[uuid.UUID][]ts.MeasurementCollectionLean)

	for _, t := range tt {
		if _, hasInstrument := response[t.InstrumentID]; !hasInstrument {
			response[t.InstrumentID] = make([]ts.MeasurementCollectionLean, 0)
		}
		mcl := ts.MeasurementCollectionLean{
			TimeseriesID: t.TimeseriesID,
			Items:        make([]ts.MeasurementLean, len(t.Measurements)),
		}
		for idx, m := range t.Measurements {
			mcl.Items[idx] = m.Lean()
		}
		response[t.InstrumentID] = append(response[t.InstrumentID], mcl)
	}

	return response, nil
}

// explorerResponseFactory returns the explorer-specific JSON response format
func explorerInclinometerResponseFactory(tt []models.InclinometerTimeseries) (map[uuid.UUID][]ts.InclinometerMeasurementCollectionLean, error) {

	response := make(map[uuid.UUID][]ts.InclinometerMeasurementCollectionLean)

	for _, t := range tt {
		if _, hasInstrument := response[t.InstrumentID]; !hasInstrument {
			response[t.InstrumentID] = make([]ts.InclinometerMeasurementCollectionLean, 0)
		}
		mcl := ts.InclinometerMeasurementCollectionLean{
			TimeseriesID: t.TimeseriesID,
			Items:        make([]ts.InclinometerMeasurementLean, len(t.Measurements)),
		}
		for idx, m := range t.Measurements {
			mcl.Items[idx] = m.InclinometerLean()
		}
		response[t.InstrumentID] = append(response[t.InstrumentID], mcl)
	}

	return response, nil
}
