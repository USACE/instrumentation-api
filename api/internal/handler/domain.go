package handler

import (
	"net/http"

	"github.com/USACE/instrumentation-api/api/internal/httperr"
	_ "github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/labstack/echo/v4"
)

// GetDomains godoc
//
//	@Summary lists all domains
//	@Tags domain
//	@Produce json
//	@Success 200 {array} model.Domain
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /domains [get]
func (h *ApiHandler) GetDomains(c echo.Context) error {
	dd, err := h.DomainService.GetDomains(c.Request().Context())
	if err != nil {
		return httperr.InternalServerError(err)
	}
	return c.JSON(http.StatusOK, dd)
}

// GetDomainMap godoc
//
//	@Summary Get map with domain group as key
//	@Tags domain
//	@Produce json
//	@Success 200 {object} model.DomainMap
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /domains/map [get]
func (h *ApiHandler) GetDomainMap(c echo.Context) error {
	dm, err := h.DomainService.GetDomainMap(c.Request().Context())
	if err != nil {
		return httperr.InternalServerError(err)
	}
	return c.JSON(http.StatusOK, dm)
}
