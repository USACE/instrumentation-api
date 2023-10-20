package middleware

import (
	"github.com/labstack/echo/v4"
	echomw "github.com/labstack/echo/v4/middleware"
)

var cors = echomw.CORS()

func (m *mw) CORS(next echo.HandlerFunc) echo.HandlerFunc {
	return cors(next)
}
