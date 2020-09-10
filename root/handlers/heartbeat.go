package handlers

import (
	"api/root/models"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
)

// DoHeartbeat triggers regular-interval tasks
func DoHeartbeat(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		h, err := models.DoHeartbeat(db)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, h)
	}
}

// GetLatestHeartbeat returns the latest heartbeat entry
func GetLatestHeartbeat(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		h, err := models.GetLatestHeartbeat(db)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, h)
	}
}
