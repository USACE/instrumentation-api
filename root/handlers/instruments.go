package handlers

import (
	"api/root/models"
	"database/sql"
	"net/http"

	"github.com/labstack/echo"
)

// GetInstruments returns instruments
func GetInstruments(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetInstruments(db))
	}
}

// GetInstrumentGroups returns instrument groups
func GetInstrumentGroups(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetInstrumentGroups(db))
	}
}
