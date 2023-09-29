package handler

import (
	"net/http"

	"github.com/USACE/instrumentation-api/api/internal/messages"
	"github.com/USACE/instrumentation-api/api/internal/model"

	"github.com/labstack/echo/v4"
)

// ListOpendcsSites returns all Instruments, represented as Opendcs Sites
func (h ApiHandler) ListOpendcsSites(c echo.Context) error {
	ss, err := model.ListOpendcsSites(db)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, messages.BadRequest)
	}
	return c.XMLPretty(http.StatusOK, ss, "  ")
}
