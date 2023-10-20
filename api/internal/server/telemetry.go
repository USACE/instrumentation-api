package server

import (
	"net/http"

	"github.com/USACE/instrumentation-api/api/internal/config"
	"github.com/USACE/instrumentation-api/api/internal/handler"
	"github.com/labstack/echo/v4"
)

type TelemetryServer struct {
	e *echo.Echo
	telemetryGroups
}

type telemetryGroups struct {
	public     *echo.Group
	datalogger *echo.Group
}

func NewTelemetryServer(cfg *config.TelemetryConfig, h *handler.TelemetryHandler) *TelemetryServer {
	e := echo.New()

	public := e.Group(cfg.RoutePrefix)
	datalogger := e.Group(cfg.RoutePrefix)

	e.Use(h.Middleware.CORS, h.Middleware.GZIP)
	datalogger.Use(h.Middleware.DataloggerKeyAuth)

	server := &TelemetryServer{e, telemetryGroups{
		public,
		datalogger,
	}}
	server.RegisterRoutes(h)

	return server
}

func (server *TelemetryServer) Start() error {
	return http.ListenAndServe(":80", server.e)
}

func (r *TelemetryServer) RegisterRoutes(h *handler.TelemetryHandler) {
	r.public.GET("/health", h.Healthcheck)
	r.datalogger.POST("/telemetry/datalogger/:model/:sn", h.CreateOrUpdateDataloggerMeasurements)
}
