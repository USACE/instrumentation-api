package handlers

import (
	"api/root/models"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
)

// GetTimeseries returns timeseries measurements data
func GetTimeseriesMeasurements(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetTimeseriesMeasurements(db))
	}
}
