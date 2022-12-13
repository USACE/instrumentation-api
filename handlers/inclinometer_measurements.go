package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/USACE/instrumentation-api/models"
	"github.com/USACE/instrumentation-api/timeseries"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

// allInclinometerTimeseriesBelongToProject is a helper function to determine if all timeseries IDs belong to a given project ID
func allInclinometerTimeseriesBelongToProject(db *sqlx.DB, mcc *models.InclinometerMeasurementCollectionCollection, projectID *uuid.UUID) (bool, error) {
	// timeseries IDs
	dd := mcc.InclinometerTimeseriesIDs()
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

// ListInclinometerMeasurements returns a timeseries with inclinometer measurements
func ListInclinometerMeasurements(db *sqlx.DB) echo.HandlerFunc {
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

		im, err := models.ListInclinometerMeasurements(db, &tsID, &tw)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		cm, err := models.ConstantMeasurement(db, &tsID, "inclinometer-constant")
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		for idx := range im.Inclinometers {
			values, err := models.ListInclinometerMeasurementValues(db, &tsID, im.Inclinometers[idx].Time, cm.Value)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, err)
			}

			jsonValues, err := json.Marshal(values)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, err)
			}
			im.Inclinometers[idx].Values = jsonValues
		}

		return c.JSON(http.StatusOK, im)
	}
}

// CreateOrUpdateProjectInclinometerMeasurements Creates or Updates a InclinometerMeasurement object or array of objects
// All timeseries must belong to the same project
func CreateOrUpdateProjectInclinometerMeasurements(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		var mcc models.InclinometerMeasurementCollectionCollection
		if err := c.Bind(&mcc); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		// Check :project_id from route against each timeseries' project_id in the database
		pID, err := uuid.Parse(c.Param("project_id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		isTrue, err := allInclinometerTimeseriesBelongToProject(db, &mcc, &pID)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		if !isTrue {
			return c.String(http.StatusBadRequest, "all timeseries posted do not belong to project")
		}

		// Post inclinometers
		p := c.Get("profile").(*models.Profile)
		stored, err := models.CreateOrUpdateInclinometerMeasurements(db, mcc.Items, p, time.Now())
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		//create inclinometer constant if doesn't exist
		if len(mcc.Items) > 0 {
			cm, err := models.ConstantMeasurement(db, &mcc.Items[0].TimeseriesID, "inclinometer-constant")
			if err != nil {
				return c.JSON(http.StatusBadRequest, err)
			}

			if cm.TimeseriesID == uuid.Nil {
				err := models.CreateTimeseriesConstant(db, &mcc.Items[0].TimeseriesID, "inclinometer-constant", "Meters", 20000)
				if err != nil {
					return c.JSON(http.StatusBadRequest, err)
				}
			}

		}

		return c.JSON(http.StatusCreated, stored)
	}
}

// DeleteInclinometerMeasurements deletes a single inclinometer measurement
func DeleteInclinometerMeasurements(db *sqlx.DB) echo.HandlerFunc {
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

		if err := models.DeleteInclinometerMeasurements(db, &id, t); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusOK, make(map[string]interface{}))
	}
}
