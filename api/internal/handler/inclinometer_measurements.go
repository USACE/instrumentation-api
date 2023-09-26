package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/messages"
	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/USACE/instrumentation-api/api/internal/models"

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
			return echo.NewHTTPError(http.StatusBadRequest, messages.MalformedID)
		}

		// Time Window
		var tw model.TimeWindow
		a, b := c.QueryParam("after"), c.QueryParam("before")
		if err = tw.SetWindow(a, b, time.Now().AddDate(0, 0, -7), time.Now()); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		im, err := models.ListInclinometerMeasurements(db, &tsID, &tw)

		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		cm, err := models.ConstantMeasurement(db, &tsID, "inclinometer-constant")
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		for idx := range im.Inclinometers {
			values, err := models.ListInclinometerMeasurementValues(db, &tsID, im.Inclinometers[idx].Time, cm.Value)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}

			jsonValues, err := json.Marshal(values)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
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
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		// Check :project_id from route against each timeseries' project_id in the database
		pID, err := uuid.Parse(c.Param("project_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		isTrue, err := allInclinometerTimeseriesBelongToProject(db, &mcc, &pID)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if !isTrue {
			return echo.NewHTTPError(http.StatusBadRequest, "all timeseries posted do not belong to project")
		}

		// Post inclinometers
		p := c.Get("profile").(*models.Profile)
		stored, err := models.CreateOrUpdateInclinometerMeasurements(db, mcc.Items, p, time.Now())
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		//create inclinometer constant if doesn't exist
		if len(mcc.Items) > 0 {
			cm, err := models.ConstantMeasurement(db, &mcc.Items[0].TimeseriesID, "inclinometer-constant")
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}

			if cm.TimeseriesID == uuid.Nil {
				err := models.CreateTimeseriesConstant(db, &mcc.Items[0].TimeseriesID, "inclinometer-constant", "Meters", 20000)
				if err != nil {
					return echo.NewHTTPError(http.StatusBadRequest, err.Error())
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
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		timeString := c.QueryParam("time")

		t, err := time.Parse(time.RFC3339, timeString)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		if err := models.DeleteInclinometerMeasurements(db, &id, t); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, make(map[string]interface{}))
	}
}
