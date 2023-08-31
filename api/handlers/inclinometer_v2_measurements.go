package handlers

import (
	"net/http"
	"time"

	"github.com/USACE/instrumentation-api/api/messages"
	"github.com/USACE/instrumentation-api/api/models"
	"github.com/USACE/instrumentation-api/api/timeseries"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

// allInclinometerV2TimeseriesBelongToProject is a helper function to determine if all timeseries IDs belong to a given project ID
func allInclinometerV2TimeseriesBelongToProject(db *sqlx.DB, mcs *models.InclinometerV2MeasurementCollections, projectID *uuid.UUID) (bool, error) {
	dd := make([]uuid.UUID, len(mcs.Items))
	for i := range mcs.Items {
		dd[i] = mcs.Items[i].TimeseriesID
	}

	m, err := models.GetTimeseriesProjectMap(db, dd)
	if err != nil {
		return false, err
	}
	for _, tID := range dd {
		pID, ok := m[tID]
		if !ok {
			return false, nil
		}
		if pID != *projectID {
			return false, nil
		}
	}
	return true, nil
}

func ListInclinometerV2Measurements(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		tsID, err := uuid.Parse(c.Param("timeseries_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, messages.MalformedID)
		}

		var tw timeseries.TimeWindow
		a, b := c.QueryParam("after"), c.QueryParam("before")
		if err = tw.SetWindow(a, b, time.Now().AddDate(0, 0, -7), time.Now()); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		mc, err := models.ListInclinometerV2Measurements(db, &tsID, &tw)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, mc)
	}
}

// CreateOrUpdateProjectInclinometerV2Measurements Creates or Updates a InclinometerV2Measurement object or array of objects
// All timeseries must belong to the same project
func CreateOrUpdateProjectInclinometerV2Measurements(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		var mcs models.InclinometerV2MeasurementCollections
		if err := c.Bind(&mcs); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		pID, err := uuid.Parse(c.Param("project_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		isTrue, err := allInclinometerV2TimeseriesBelongToProject(db, &mcs, &pID)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if !isTrue {
			return echo.NewHTTPError(http.StatusBadRequest, "all timeseries posted do not belong to project")
		}

		if err := models.CreateOrUpdateInclinometerV2Measurements(db, mcs); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusCreated, make(map[string]interface{}))
	}
}

// DeleteInclinometerV2Measurements deletes a single inclinometerv2 measurement
func DeleteInclinometerV2Measurements(db *sqlx.DB) echo.HandlerFunc {
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

		if err := models.DeleteInclinometerV2Measurements(db, &id, t); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, make(map[string]interface{}))
	}
}
