package middleware

import (
	"context"
	"log/slog"
	"os"

	"github.com/USACE/instrumentation-api/api/internal/util"
	"github.com/labstack/echo/v4"
	echomw "github.com/labstack/echo/v4/middleware"
)

var slogger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

var requestLogger = echomw.RequestLoggerWithConfig(echomw.RequestLoggerConfig{
	LogStatus:   true,
	LogURI:      true,
	LogError:    true,
	HandleError: true, // forwards error to the global error handler, so it can decide appropriate status code
	LogValuesFunc: func(c echo.Context, v echomw.RequestLoggerValues) error {
		urlRedact := util.RedactRequest{URL: v.URI}
		if err := urlRedact.RedactQueryParam("key"); err != nil {
			return err
		}
		if v.Error == nil {
			slogger.LogAttrs(context.Background(), slog.LevelInfo, "REQUEST",
				slog.String("uri", urlRedact.URL),
				slog.Int("status", v.Status),
			)
		} else {
			slogger.LogAttrs(context.Background(), slog.LevelError, "REQUEST_ERROR",
				slog.String("uri", urlRedact.URL),
				slog.Int("status", v.Status),
				slog.String("err", v.Error.Error()),
			)
		}
		return nil
	},
})

func (m *mw) RequestLogger(next echo.HandlerFunc) echo.HandlerFunc {
	return requestLogger(next)
}
