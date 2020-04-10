package handlers

import (
	"api/root/models"
	"database/sql"
	"net/http"

	"github.com/labstack/echo"
)

// GetTimeseries returns timeseries endpoints
func GetTimeseries(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetTimeseries(db))
	}
}