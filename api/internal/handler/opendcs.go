package handler

import (
	"net/http"

	"github.com/USACE/instrumentation-api/api/internal/message"

	"github.com/labstack/echo/v4"
)

// ListOpendcsSites returns all Instruments, represented as Opendcs Sites
func (h *ApiHandler) ListOpendcsSites(c echo.Context) error {
	ss, err := h.OpendcsService.ListOpendcsSites(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.BadRequest)
	}
	return c.XMLPretty(http.StatusOK, ss, "  ")
}
