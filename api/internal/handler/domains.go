package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetDomains returns all database domains in a single endpoint
func (h *ApiHandler) GetDomains(c echo.Context) error {
	dd, err := h.DomainService.GetDomains(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, dd)
}
