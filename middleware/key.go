package middleware

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// KeyAuth returns ready-to-go key auth middleware based on a known token
func KeyAuth(validKey string) echo.MiddlewareFunc {
	return middleware.KeyAuth(
		func(key string, c echo.Context) (bool, error) {
			return key == validKey, nil
		},
	)
}
