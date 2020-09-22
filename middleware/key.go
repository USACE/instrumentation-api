package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// KeyAuth returns ready-to-go key auth middleware based on a known token
func KeyAuth(validKey string) echo.MiddlewareFunc {
	return middleware.KeyAuth(
		func(key string, c echo.Context) (bool, error) {
			return key == validKey, nil
		},
	)
}
