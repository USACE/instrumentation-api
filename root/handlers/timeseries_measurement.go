package handlers

import (
	"api/root/models"
	"database/sql"
	"net/http"

	"github.com/labstack/echo"
)

// GetTimeseries returns timeseries measurements data
func GetTimeseriesMeasurements(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetTimeseriesMeasurements(db))
	}
}
