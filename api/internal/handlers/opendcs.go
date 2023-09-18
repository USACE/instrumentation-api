package handlers

import (
	"net/http"

	"github.com/USACE/instrumentation-api/api/internal/messages"
	"github.com/USACE/instrumentation-api/api/internal/models"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

// ListOpendcsSites returns all Instruments, represented as Opendcs Sites
func ListOpendcsSites(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		ss, err := models.ListOpendcsSites(db)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, messages.BadRequest)
		}
		return c.XMLPretty(http.StatusOK, ss, "  ")
	}
}
