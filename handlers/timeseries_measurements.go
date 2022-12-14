package handlers

import (
	"net/http"
	"sort"
	"time"

	"github.com/USACE/instrumentation-api/models"
	"github.com/USACE/instrumentation-api/timeseries"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

// allTimeseriesBelongToProject is a helper function to determine if all timeseries IDs belong to a given project ID
func allTimeseriesBelongToProject(db *sqlx.DB, mcc *models.TimeseriesMeasurementCollectionCollection, projectID *uuid.UUID) (bool, error) {
	// timeseries IDs
	dd := mcc.TimeseriesIDs()
	m, err := models.GetTimeseriesProjectMap(db, dd)
	if err != nil {
		return false, err
	}
	for _, tID := range dd {
		pID, ok := m[tID]
		// timeseries does not exist; therefore does not belong to project
		if !ok {
			return false, nil
		}
		// timeseries' project_id in database does not match projectID
		if pID != *projectID {
			return false, nil
		}
	}
	return true, nil
}

// ListTimeseriesMeasurements TODO (possible optimization):
// Due to the amount of repeated data for each time entry (and the inherent size of those payloads)
// it is possible to decreate repetition for some data points for regularized computed intervals,
// the value can be omitted only if the measurement:
//
// 		- IS the same as the previous value
// 		- IS NOT the first measurement
// 		- IS NOT the last measurement
// 		- IS NOT the same as the value directly ahead
//
// A downside to this approach is that it may be harder to have distinction between measurements
// that are not present in the timeseries vs repeating measurements.
//
// E.g.
//
// if i != 0
// 		&& i-1 != len(measurements)
// 		&& measurements[i-1] == measurements[i]
// 		&& measurements[i+1] == measurements[i] {
//	continue
// }
//

// ListTimeseriesMeasurements returns a timeseries with measurements
func ListTimeseriesMeasurements(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		tsID, err := uuid.Parse(c.Param("timeseries_id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Malformed ID")
		}

		// Time Window
		var tw timeseries.TimeWindow
		a, b := c.QueryParam("after"), c.QueryParam("before")
		// If after or before are not provided return last 7 days of data from current time
		if a == "" || b == "" {
			tw.Before = time.Now()
			tw.After = tw.Before.AddDate(0, 0, -7)
		} else {
			// Attempt to parse query param "after"
			tA, err := time.Parse(time.RFC3339, a)
			if err != nil {
				return c.JSON(http.StatusBadRequest, err)
			}
			tw.After = tA
			// Attempt to parse query param "before"
			tB, err := time.Parse(time.RFC3339, b)
			if err != nil {
				return c.JSON(http.StatusBadRequest, err)
			}
			tw.Before = tB
		}

		// "interval" query parameter
		// If parameter is omitted or 0, resampling not applied
		p := c.QueryParam("interval")
		interval, err := time.ParseDuration(p)
		if p != "" && err != nil {
			return c.String(
				http.StatusBadRequest,
				"Invalid interval. Valid time units are \"ns\", \"us\", \"ms\", \"s\", \"m\", \"h\" E.g. \"5h30m5s\"",
			)
		}

		// Bind Timeseries id to struct
		ts, err := models.GetTimeseries(db, &tsID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		// If timeseries NOT computed, query stored measurements
		if !ts.IsComputed {
			mc, err := models.ListTimeseriesMeasurements(db, &ts.ID, &tw)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, err.Error())
			}
			sort.Slice(mc.Items, func(i, j int) bool {
				return mc.Items[i].Time.Before(mc.Items[j].Time)
			})

			return c.JSON(http.StatusOK, mc)
		}

		// If timeseries IS computed, calulate measurements
		ct, err := models.ComputedTimeseriesWithMeasurements(db, &ts.ID, &ts.InstrumentID, &tw, &interval)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		tsms := make([]models.Measurement, 0)

		// Trim “masked”, “validated”, and “annotation” fields as they only apply to stored timeseries
		for _, t := range ct {
			// Sort by time
			sort.Slice(t.Measurements, func(i, j int) bool {
				return t.Measurements[i].Time.Before(t.Measurements[j].Time)
			})
			for _, m := range t.Measurements {
				tsms = append(tsms, models.Measurement{
					Time:  m.Time,
					Value: m.Value,
				})
			}
		}
		mc := models.MeasurementCollection{
			TimeseriesID: tsID,
			Items:        tsms,
		}
		return c.JSON(http.StatusOK, mc)
	}
}

// CreateOrUpdateProjectTimeseriesMeasurements Creates or Updates a TimeseriesMeasurement object or array of objects
// All timeseries must belong to the same project
func CreateOrUpdateProjectTimeseriesMeasurements(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		var mcc models.TimeseriesMeasurementCollectionCollection
		if err := c.Bind(&mcc); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		// Check :project_id from route against each timeseries' project_id in the database
		pID, err := uuid.Parse(c.Param("project_id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		isTrue, err := allTimeseriesBelongToProject(db, &mcc, &pID)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		if !isTrue {
			return c.String(http.StatusBadRequest, "all timeseries posted do not belong to project")
		}
		// Post timeseries
		stored, err := models.CreateOrUpdateTimeseriesMeasurements(db, mcc.Items)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusCreated, stored)
	}
}

// CreateOrUpdateTimeseriesMeasurements Creates or Updates a TimeseriesMeasurement object or array of objects
// Timeseries may belong to one or more projects
func CreateOrUpdateTimeseriesMeasurements(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		var mcc models.TimeseriesMeasurementCollectionCollection
		if err := c.Bind(&mcc); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		// Post timeseries
		stored, err := models.CreateOrUpdateTimeseriesMeasurements(db, mcc.Items)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusCreated, stored)
	}
}

// UpdateTimeseriesMeasurements Overwrites measurements with the supplied payload
// within a TimeWindow (> after, < before)
func UpdateTimeseriesMeasurements(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Time Window
		var tw timeseries.TimeWindow
		a, b := c.QueryParam("after"), c.QueryParam("before")
		// If after or before are not provided return last 7 days of data from current time
		if a == "" || b == "" {
			tw.Before = time.Now()
			tw.After = tw.Before.AddDate(0, 0, -7)
		} else {
			// Attempt to parse query param "after"
			tA, err := time.Parse(time.RFC3339, a)
			if err != nil {
				return c.JSON(http.StatusBadRequest, err)
			}
			tw.After = tA
			// Attempt to parse query param "before"
			tB, err := time.Parse(time.RFC3339, b)
			if err != nil {
				return c.JSON(http.StatusBadRequest, err)
			}
			tw.Before = tB
		}

		var mcc models.TimeseriesMeasurementCollectionCollection
		if err := c.Bind(&mcc); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		// Put timeseries measurments
		stored, err := models.UpdateTimeseriesMeasurements(db, mcc.Items, &tw)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusOK, stored)
	}
}

// DeleteTimeserieMeasurements deletes a single timeseries measurement
func DeleteTimeserieMeasurements(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// id from url params
		id, err := uuid.Parse(c.Param("timeseries_id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		timeString := c.QueryParam("time")

		t, err := time.Parse(time.RFC3339, timeString)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		if err := models.DeleteTimeserieMeasurements(db, &id, t); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusOK, make(map[string]interface{}))
	}
}
