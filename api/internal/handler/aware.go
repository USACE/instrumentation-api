package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *ApiHandler) ListAwareParameters(c echo.Context) error {
	pp, err := h.AwareParameterService.ListAwareParameters(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, pp)
}

func (h *ApiHandler) ListAwarePlatformParameterConfig(c echo.Context) error {
	cc, err := h.AwareParameterService.ListAwarePlatformParameterConfig(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, cc)
}
