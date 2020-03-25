package handlers

import (
	"api/root/models"
	"database/sql"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo"
)

// GetInstruments returns instruments
func GetInstruments(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetInstruments(db))
	}
}

// GetInstrument returns a single instrument
func GetInstrument(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			return c.String(http.StatusNotFound, "Malformed ID")
		}
		return c.JSON(http.StatusOK, models.GetInstrument(db, id.String()))
	}
}

// // CreateInstrument creates
// func CreateInstrument(db *sql.DB) echo.HandlerFunc {

// }

// // UpdateInstrument modifys an existing instrument
// func UpdateInstrument(db *sql.DB) echo.HandlerFunc {

// }

// // DeleteInstrument deletes an existing instrument by ID
// func DeleteInstrument(db *sql.DB) echo.HandlerFunc {

// }
