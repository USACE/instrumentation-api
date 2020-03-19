package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

// GetRoot returns an array of all API endpoints
func GetRoot(c echo.Context) error {
	return c.String(http.StatusOK, "HHD API Root")
}
