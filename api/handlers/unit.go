package handlers

import (
	"github.com/USACE/instrumentation-api/api/models"

	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

// ListUnits returns an array of timeseries
func ListUnits(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		uu, err := models.ListUnits(db)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusOK, uu)
	}
}
