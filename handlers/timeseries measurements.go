package handlers

import (
	"github.com/USACE/instrumentation-api/models"
	"github.com/USACE/instrumentation-api/timeseries"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
)

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
		// If after or before are not provided
		// Return last 14 days of data from current time
		if a == "" || b == "" {
			tw.Before = time.Now()
			tw.After = tw.Before.AddDate(0, 0, -14)
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

		mc, err := models.ListTimeseriesMeasurements(db, &tsID, &tw)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, mc)
	}
}

// CreateOrUpdateTimeseriesMeasurements Creates or Updates a TimeseriesMeasurement object or array of objects
func CreateOrUpdateTimeseriesMeasurements(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		var mcc models.TimeseriesMeasurementCollectionCollection
		if err := c.Bind(&mcc); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		if err := models.CreateOrUpdateTimeseriesMeasurements(db, mcc.Items); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.NoContent(http.StatusCreated)
	}
}
