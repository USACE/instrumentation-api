package handlers

import (
	"net/http"
	"time"

	"github.com/USACE/instrumentation-api/api/models"
	"github.com/USACE/instrumentation-api/api/timeseries"

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

func PostInclinometerExplorer(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		// Filters used in SQL Query
		var f Filter

		// Instrument IDs from POST
		if err := (&echo.DefaultBinder{}).BindBody(c, &f.InstrumentID); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		// Get timeWindow from query params
		var tw timeseries.TimeWindow
		a, b := c.QueryParam("after"), c.QueryParam("before")
		if err := tw.SetWindow(a, b, time.Now().AddDate(0, 0, -7), time.Now()); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		f.TimeWindow = tw

		// Get Stored And Computed Timeseries With Measurements
		interval := time.Hour // Set to 1 Hour; TODO - do not hard-code interval
		tt, err := models.ComputedInclinometerTimeseries(db, f.InstrumentID, &f.TimeWindow, &interval)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		// Convert Rows to Response
		response, err := explorerInclinometerResponseFactory(tt)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, response)
	}
}

// explorerResponseFactory returns the explorer-specific JSON response format
func ExplorerResponseFactory(tt []models.Timeseries) (map[uuid.UUID][]timeseries.MeasurementCollectionLean, error) {

	response := make(map[uuid.UUID][]timeseries.MeasurementCollectionLean)

	for _, t := range tt {
		if _, hasInstrument := response[t.InstrumentID]; !hasInstrument {
			response[t.InstrumentID] = make([]timeseries.MeasurementCollectionLean, 0)
		}
		mcl := timeseries.MeasurementCollectionLean{
			TimeseriesID: t.TimeseriesID,
			Items:        make([]timeseries.MeasurementLean, len(t.Measurements)),
		}
		for idx, m := range t.Measurements {
			mcl.Items[idx] = m.Lean()
		}
		response[t.InstrumentID] = append(response[t.InstrumentID], mcl)
	}

	return response, nil
}

// explorerResponseFactory returns the explorer-specific JSON response format
func explorerInclinometerResponseFactory(tt []models.InclinometerTimeseries) (map[uuid.UUID][]timeseries.InclinometerMeasurementCollectionLean, error) {

	response := make(map[uuid.UUID][]timeseries.InclinometerMeasurementCollectionLean)

	for _, t := range tt {
		if _, hasInstrument := response[t.InstrumentID]; !hasInstrument {
			response[t.InstrumentID] = make([]timeseries.InclinometerMeasurementCollectionLean, 0)
		}
		mcl := timeseries.InclinometerMeasurementCollectionLean{
			TimeseriesID: t.TimeseriesID,
			Items:        make([]timeseries.InclinometerMeasurementLean, len(t.Measurements)),
		}
		for idx, m := range t.Measurements {
			mcl.Items[idx] = m.InclinometerLean()
		}
		response[t.InstrumentID] = append(response[t.InstrumentID], mcl)
	}

	return response, nil
}
