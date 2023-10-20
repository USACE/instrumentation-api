package handler

import (
	"net/http"

	_ "github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/labstack/echo/v4"
)

// GetHome godoc
//
//	@Summary gets information for the homepage
//	@Tags home
//	@Produce json
//	@Success 200 {object} model.Home
//	@Failure 500 {object} echo.HTTPError
//	@Router /home [get]
func (h *ApiHandler) GetHome(c echo.Context) error {
	home, err := h.HomeService.GetHome(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, home)
}
