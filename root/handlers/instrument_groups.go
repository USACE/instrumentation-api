package handlers

import (
	"api/root/models"
	"database/sql"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo"
)

// GetInstrumentGroups returns instrument groups
func GetInstrumentGroups(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetInstrumentGroups(db))
	}
}

// GetInstrumentGroup returns single instrument group
func GetInstrumentGroup(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			return c.String(http.StatusNotFound, "Malformed ID")
		}
		return c.JSON(http.StatusOK, models.GetInstrumentGroup(db, id.String()))
	}
}

// GetInstrumentGroupInstruments returns a list of instruments for a provided instrument group
func GetInstrumentGroupInstruments(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			return c.String(http.StatusNotFound, "Malformed ID")
		}
		return c.JSON(http.StatusOK, models.GetInstrumentGroupInstruments(db, id.String()))
	}
}
