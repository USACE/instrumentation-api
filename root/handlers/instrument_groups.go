package handlers

import (
	"api/root/models"
	"net/http"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
)

// ListInstrumentGroups returns instrument groups
func ListInstrumentGroups(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.ListInstrumentGroups(db))
	}
}

// // GetInstrumentGroup returns single instrument group
func GetInstrumentGroup(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			return c.String(http.StatusNotFound, "Malformed ID")
		}
		return c.JSON(http.StatusOK, models.GetInstrumentGroup(db, id))
	}
}

// // ListInstrumentGroupInstruments returns a list of instruments for a provided instrument group
func ListInstrumentGroupInstruments(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			return c.String(http.StatusNotFound, "Malformed ID")
		}
		return c.JSON(http.StatusOK, models.ListInstrumentGroupInstruments(db, id))
	}
}
