package handlers

import (
	"api/root/models"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
)

// GetTimeseries returns timeseries endpoints
func GetTimeseries(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetTimeseries(db))
	}
}