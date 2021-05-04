package handlers

import (
	"fmt"
	"net/http"
	"time"

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
	TimeWindow   ts.TimeWindow
}

// PostExplorer retrieves timeseries information for the explorer app component
func PostExplorer(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		// Filters used in SQL Query
		var f Filter

		// Instrument IDs from POST
		if err := c.Bind(&f.InstrumentID); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		// Time Window From POST
		a, b := c.QueryParam("after"), c.QueryParam("before")
		// If after or before are not provided; Return last 7 days of data from current time
		if a == "" || b == "" {
			f.TimeWindow.Before = time.Now()
			f.TimeWindow.After = f.TimeWindow.Before.AddDate(0, 0, -7)
		} else {
			// Attempt to parse query param "after"
			tA, err := time.Parse(time.RFC3339, a)
			if err != nil {
				return c.JSON(http.StatusBadRequest, err)
			}
			f.TimeWindow.After = tA
			// Attempt to parse query param "before"
			tB, err := time.Parse(time.RFC3339, b)
			if err != nil {
				return c.JSON(http.StatusBadRequest, err)
			}
			f.TimeWindow.Before = tB
		}

		// Get Rows from the Database
		ee, err := explorerRows(db, &f)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		// Convert Rows to Response
		response, err := explorerResponseFactory(ee)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, response)
	}
}

func explorerRows(db *sqlx.DB, f *Filter) ([]ExplorerRow, error) {

	sql := func(inClause string) string {
		return fmt.Sprintf(
			`SELECT i.id           AS instrument_id,
	                t.id           AS timeseries_id,
	                t.parameter_id AS parameter_id,
	                t.unit_id      AS unit_id,
	                m.time         AS time,
	                m.value        AS value
             FROM   timeseries_measurement m
             INNER JOIN timeseries t ON t.id = m.timeseries_id
             INNER JOIN instrument i ON i.id = t.instrument_id
			 WHERE NOT i.deleted
			 %s
			 ORDER BY instrument_id, timeseries_id, time DESC;`,
			inClause,
		)
	}

	sqlxInResult := func() (string, []interface{}, error) {
		switch {
		// Filter by Instrument IDs, Parameter IDs, Time Window
		case len(f.InstrumentID) > 0 && len(f.ParameterID) > 0:
			return sqlx.In(
				sql("AND i.id IN (?) AND t.parameter_id IN (?) AND m.time <= ? AND m.time >= ?"),
				f.InstrumentID, f.ParameterID, f.TimeWindow.Before, f.TimeWindow.After,
			)
		// Filter by Instrument IDs, Time Window Only
		case len(f.InstrumentID) > 0:
			return sqlx.In(
				sql("AND i.id IN (?) AND m.time <= ? AND m.time >= ?"),
				f.InstrumentID, f.TimeWindow.Before, f.TimeWindow.After,
			)
		default:
			return sql(""), make([]interface{}, 0), nil
		}
	}

	// SQL Things...
	var mm []ExplorerRow
	query, args, err := sqlxInResult()
	if err != nil {
		return make([]ExplorerRow, 0), err
	}
	if err := db.Select(&mm, db.Rebind(query), args...); err != nil {
		return make([]ExplorerRow, 0), err
	}

	return mm, nil
}

// explorerResponseFactory returns the explorer-specific JSON response format
func explorerResponseFactory(rr []ExplorerRow) (map[uuid.UUID][]ts.MeasurementCollectionLean, error) {

	response := make(map[uuid.UUID][]ts.MeasurementCollectionLean)
	var iID uuid.UUID
	var tsl ts.MeasurementCollectionLean
	for idx, v := range rr {
		if v.InstrumentID != iID {
			iID = v.InstrumentID // Set to New Instrument
			response[v.InstrumentID] = make([]ts.MeasurementCollectionLean, 0)
		}
		if v.TimeseriesID != tsl.TimeseriesID {
			if idx != 0 {
				response[v.InstrumentID] = append(response[v.InstrumentID], tsl)
			}
			tsl.TimeseriesID = v.TimeseriesID         // Set to New Timeseries
			tsl.Items = make([]ts.MeasurementLean, 0) // Empty the slice of measurements
		}
		// Add measurement to appropriate part of the map
		tsl.Items = append(tsl.Items, ts.MeasurementLean{v.Time: v.Value})
	}

	return response, nil
}
