package handlers

import (
	"net/http"

	"github.com/USACE/instrumentation-api/models"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

// ListOpendcsSites returns all Instruments, represented as Opendcs Sites
func ListOpendcsSites(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		ss, err := models.ListOpendcsSites(db)
		if err != nil {
			return c.JSON(http.StatusBadRequest, models.DefaultMessageBadRequest)
		}
		return c.XMLPretty(http.StatusOK, ss, "  ")
	}
}
