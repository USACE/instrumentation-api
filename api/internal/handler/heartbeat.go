package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// DoHeartbeat triggers regular-interval tasks
func (h ApiHandler) DoHeartbeat(c echo.Context) error {
	// Create a Record of Heartbeat
	hb, err := h.HeartbeatStore.DoHeartbeat(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, hb)
}

// GetLatestHeartbeat returns the latest heartbeat entry
func (h ApiHandler) GetLatestHeartbeat(c echo.Context) error {
	hb, err := h.HeartbeatStore.GetLatestHeartbeat(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, hb)
}

// ListHeartbeats returns all heartbeats
func (h ApiHandler) ListHeartbeats(c echo.Context) error {
	hh, err := h.HeartbeatStore.ListHeartbeats(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, hh)
}
