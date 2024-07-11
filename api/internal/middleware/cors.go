package middleware

import (
	"github.com/labstack/echo/v4"
	echomw "github.com/labstack/echo/v4/middleware"
)

var cors = echomw.CORS()

func (m *mw) CORS(next echo.HandlerFunc) echo.HandlerFunc {
	return cors(next)
}

func (m *mw) CORSWhitelist(next echo.HandlerFunc) echo.HandlerFunc {
	cors := echomw.CORSWithConfig(echomw.CORSConfig{
		AllowOrigins: m.cfg.AllowOrigins,
	})
	return cors(next)
}
