package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// ListUnits returns an array of timeseries
func (h ApiHandler) ListUnits(c echo.Context) error {
	uu, err := h.UnitStore.ListUnits(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, uu)
}
