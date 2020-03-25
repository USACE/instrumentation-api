package handlers

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo"

	"api/root/models"
)

// GetDomains returns all database domains in a single endpoint
func GetDomains(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetDomains(db))
	}
}
