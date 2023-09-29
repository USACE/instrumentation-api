package handler

import (
	"net/http"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/model"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

// allTimeseriesBelongToProject is a helper function to determine if all timeseries IDs belong to a given project ID
func (h ApiHandler) allTimeseriesBelongToProject(db *sqlx.DB, mcc *model.TimeseriesMeasurementCollectionCollection, projectID *uuid.UUID) (bool, error) {
	// timeseries IDs
	dd := mcc.TimeseriesIDs()
	m, err := model.GetTimeseriesProjectMap(db, dd)
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

// CreateOrUpdateProjectTimeseriesMeasurements Creates or Updates a TimeseriesMeasurement object or array of objects
// All timeseries must belong to the same project
func (h ApiHandler) CreateOrUpdateProjectTimeseriesMeasurements(c echo.Context) error {

	var mcc model.TimeseriesMeasurementCollectionCollection
	if err := c.Bind(&mcc); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Check :project_id from route against each timeseries' project_id in the database
	pID, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	isTrue, err := allTimeseriesBelongToProject(db, &mcc, &pID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if !isTrue {
		return echo.NewHTTPError(http.StatusBadRequest, "all timeseries posted do not belong to project")
	}
	// Post timeseries
	stored, err := model.CreateOrUpdateTimeseriesMeasurements(db, mcc.Items)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, stored)
}

// CreateOrUpdateTimeseriesMeasurements Creates or Updates a TimeseriesMeasurement object or array of objects
// Timeseries may belong to one or more projects
func (h ApiHandler) CreateOrUpdateTimeseriesMeasurements(c echo.Context) error {

	var mcc model.TimeseriesMeasurementCollectionCollection
	if err := c.Bind(&mcc); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// Post timeseries
	stored, err := model.CreateOrUpdateTimeseriesMeasurements(db, mcc.Items)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, stored)
}

// UpdateTimeseriesMeasurements Overwrites measurements with the supplied payload
// within a TimeWindow (> after, < before)
func (h ApiHandler) UpdateTimeseriesMeasurements(c echo.Context) error {
	// Time Window
	var tw model.TimeWindow
	a, b := c.QueryParam("after"), c.QueryParam("before")
	if err := tw.SetWindow(a, b, time.Now().AddDate(0, 0, -7), time.Now()); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var mcc model.TimeseriesMeasurementCollectionCollection
	if err := c.Bind(&mcc); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// Put timeseries measurments
	stored, err := model.UpdateTimeseriesMeasurements(db, mcc.Items, &tw)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, stored)
}

// DeleteTimeserieMeasurements deletes a single timeseries measurement
func (h ApiHandler) DeleteTimeserieMeasurements(c echo.Context) error {
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

	if err := model.DeleteTimeserieMeasurements(db, &id, t); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, make(map[string]interface{}))
}
