package middleware

import (
	"github.com/labstack/echo/v4"
	echomw "github.com/labstack/echo/v4/middleware"
)

var gzip = echomw.GzipWithConfig(echomw.GzipConfig{Level: 5})

func (m *mw) GZIP(next echo.HandlerFunc) echo.HandlerFunc {
	return gzip(next)
}
