package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

// GetTimeseries returns timeseries endpoints
func GetTimeseries(c echo.Context) error {
	return c.String(http.StatusOK, "HHD API Timeseries Endpoint")
}
