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
	survey123  *echo.Group
}

func NewTelemetryServer(cfg *config.TelemetryConfig, h *handler.TelemetryHandler) *TelemetryServer {
	e := echo.New()

	// when debug is enabled, returned errors are included in the response body
	e.Debug = cfg.Debug

	public := e.Group(cfg.RoutePrefix)
	datalogger := e.Group(cfg.RoutePrefix)
	survey123 := e.Group(cfg.RoutePrefix)

	mw := h.Middleware
	e.Use(mw.GZIP)
	public.Use(mw.CORS)
	datalogger.Use(mw.CORS, mw.DataloggerKeyAuth)
	survey123.Use(mw.CORSWhitelist)

	server := &TelemetryServer{e, telemetryGroups{
		public,
		datalogger,
		survey123,
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

	r.survey123.POST("/telemetry/survey123/measurements", h.CreateOrUpdateSurvey123Measurements)
}
