package handler

import (
	"net/http"

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
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, dd)
}
