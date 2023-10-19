package middleware

import (
	"github.com/labstack/echo/v4"
	echomw "github.com/labstack/echo/v4/middleware"
)

var logger = echomw.Logger()

func (m *mw) Logger(next echo.HandlerFunc) echo.HandlerFunc {
	return logger(next)
}
