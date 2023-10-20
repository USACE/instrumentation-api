package handler

import (
	"net/http"

	"github.com/USACE/instrumentation-api/api/internal/message"
	_ "github.com/USACE/instrumentation-api/api/internal/model"

	"github.com/labstack/echo/v4"
)

// ListOpendcsSites godoc
//
//	@Summary lists all instruments, represented as opendcs sites
//	@Tags opendcs
//	@Produce xml
//	@Success 200 {array} model.Site
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /opendcs/sites [get]
func (h *ApiHandler) ListOpendcsSites(c echo.Context) error {
	ss, err := h.OpendcsService.ListOpendcsSites(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.BadRequest)
	}
	return c.XMLPretty(http.StatusOK, ss, "  ")
}
