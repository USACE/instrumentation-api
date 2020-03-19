package handlers

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo"
	_ "github.com/lib/pq"

	"api/root/models"
)

// GetInstruments returns all instruments in the database
func GetInstruments(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetInstruments(db))
	}
}
