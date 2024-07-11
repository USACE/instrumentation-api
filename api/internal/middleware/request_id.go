package middleware

import (
	"github.com/labstack/echo/v4"
	echomw "github.com/labstack/echo/v4/middleware"
)

var reqID = echomw.RequestID()

func (m *mw) RequestID(next echo.HandlerFunc) echo.HandlerFunc {
	return reqID(next)
}
