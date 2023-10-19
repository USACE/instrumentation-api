package middleware

import (
	"strings"

	"github.com/labstack/echo/v4"
	echomw "github.com/labstack/echo/v4/middleware"
)

var gzip = echomw.GzipWithConfig(echomw.GzipConfig{
	Level: 5,
	Skipper: func(c echo.Context) bool {
		if strings.Contains(c.Request().URL.Path, "swagger") {
			return true
		}
		return false
	},
})

func (m *mw) GZIP(next echo.HandlerFunc) echo.HandlerFunc {
	return gzip(next)
}
