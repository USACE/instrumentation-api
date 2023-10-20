package handler

import (
	"net/http"

	_ "github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/labstack/echo/v4"
)

// ListUnits godoc
//
//	@Summary lists the available units
//	@Tags unit
//	@Produce json
//	@Success 200 {array} model.Unit
//	@Failure 400 {object} echo.HTTPError
//	@Router /units [get]
func (h *ApiHandler) ListUnits(c echo.Context) error {
	uu, err := h.UnitService.ListUnits(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, uu)
}
