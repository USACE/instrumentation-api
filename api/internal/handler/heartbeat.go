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

// DoHeartbeat godoc
//
//	@Summary creates a heartbeat entry at regular intervals
//	@Tags heartbeat
//	@Produce json
//	@Param key query string true "api key"
//	@Success 200 {object} model.Heartbeat
//	@Router /heartbeat [post]
func (h *ApiHandler) DoHeartbeat(c echo.Context) error {
	hb, err := h.HeartbeatService.DoHeartbeat(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, hb)
}

// GetLatestHeartbeat godoc
//
//	@Summary gets the latest heartbeat
//	@Tags heartbeat
//	@Produce json
//	@Success 200 {object} model.Heartbeat
//	@Router /heartbeat/latest [get]
func (h *ApiHandler) GetLatestHeartbeat(c echo.Context) error {
	hb, err := h.HeartbeatService.GetLatestHeartbeat(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, hb)
}

// ListHeartbeats godoc
//
//	@Summary returns all heartbeats
//	@Tags heartbeat
//	@Produce json
//	@Success 200 {array} model.Heartbeat
//	@Router /heartbeats [get]
func (h *ApiHandler) ListHeartbeats(c echo.Context) error {
	hh, err := h.HeartbeatService.ListHeartbeats(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, hh)
}
