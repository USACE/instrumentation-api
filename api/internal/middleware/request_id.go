package middleware

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	echomw "github.com/labstack/echo/v4/middleware"
)

var reqID = echomw.RequestIDWithConfig(echomw.RequestIDConfig{
	Generator: func() string { return uuid.NewString() },
})

func (m *mw) RequestID(next echo.HandlerFunc) echo.HandlerFunc {
	return reqID(next)
}
