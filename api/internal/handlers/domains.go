package handlers

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"

	"github.com/USACE/instrumentation-api/api/internal/models"
)

// GetDomains returns all database domains in a single endpoint
func GetDomains(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		dd, err := models.GetDomains(db)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, dd)
	}
}
