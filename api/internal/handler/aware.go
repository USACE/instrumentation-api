package handler

import (
	"net/http"

	"github.com/USACE/instrumentation-api/api/internal/httperr"
	_ "github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/labstack/echo/v4"
)

// ListAwareParameters godoc
//
//	@Summary lists alert configs for a project
//	@Tags aware
//	@Produce json
//	@Success 200 {array} model.AwareParameter
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /aware/parameters [get]
func (h *ApiHandler) ListAwareParameters(c echo.Context) error {
	pp, err := h.AwareParameterService.ListAwareParameters(c.Request().Context())
	if err != nil {
		return httperr.InternalServerError(err)
	}
	return c.JSON(http.StatusOK, pp)
}

// ListAwarePlatformParameterConfig godoc
//
//	@Summary lists alert configs for a project
//	@Tags aware
//	@Produce json
//	@Success 200 {array} model.AwarePlatformParameterConfig
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /aware/data_acquisition_config [get]
func (h *ApiHandler) ListAwarePlatformParameterConfig(c echo.Context) error {
	cc, err := h.AwareParameterService.ListAwarePlatformParameterConfig(c.Request().Context())
	if err != nil {
		return httperr.InternalServerError(err)
	}
	return c.JSON(http.StatusOK, cc)
}
