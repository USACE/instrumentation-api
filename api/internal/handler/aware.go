package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (r *ApiRouter) AwareParameters(h *ApiHandler) {
	r.g.public.GET("/aware/parameters", h.ListAwareParameters)
	r.g.public.GET("/aware/data_acquisition_config", h.ListAwarePlatformParameterConfig)
}

func (h ApiHandler) ListAwareParameters(c echo.Context) error {
	pp, err := h.AwareParameterStore.ListAwareParameters(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, pp)
}

func (h ApiHandler) ListAwarePlatformParameterConfig(c echo.Context) error {
	cc, err := h.AwareParameterStore.ListAwarePlatformParameterConfig(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, cc)
}
