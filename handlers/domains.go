package handlers

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"

	"github.com/USACE/instrumentation-api/models"
)

// GetDomains returns all database domains in a single endpoint
func GetDomains(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		dd, err := models.GetDomains(db)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusOK, dd)
	}
}
