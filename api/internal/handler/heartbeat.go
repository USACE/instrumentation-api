package handler

import (
	"net/http"

	_ "github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/labstack/echo/v4"
)

// Healthcheck godoc
//
//	@Summary checks the health of the api server
//	@Tags heartbeat
//	@Produce json
//	@Success 200 {array} map[string]interface{}
//	@Router /health [get]
func (h *ApiHandler) Healthcheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{"status": "healthy"})
}

func (h *TelemetryHandler) Healthcheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{"status": "healthy"})
}

// DoHeartbeat triggers regular-interval tasks
func (h *ApiHandler) DoHeartbeat(c echo.Context) error {
	hb, err := h.HeartbeatService.DoHeartbeat(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, hb)
}

// GetLatestHeartbeat returns the latest heartbeat entry
func (h *ApiHandler) GetLatestHeartbeat(c echo.Context) error {
	hb, err := h.HeartbeatService.GetLatestHeartbeat(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, hb)
}

// ListHeartbeats returns all heartbeats
func (h *ApiHandler) ListHeartbeats(c echo.Context) error {
	hh, err := h.HeartbeatService.ListHeartbeats(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, hh)
}
