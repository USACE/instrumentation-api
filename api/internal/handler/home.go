package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetHome returns information for the homepage
func (h ApiHandler) GetHome(c echo.Context) error {
	home, err := h.HomeStore.GetHome(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, home)
}
