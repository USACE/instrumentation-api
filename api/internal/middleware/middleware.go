package middleware

import (
	"github.com/USACE/instrumentation-api/api/internal/config"
	"github.com/USACE/instrumentation-api/api/internal/service"
	"github.com/labstack/echo/v4"
)

type Middleware interface {
	AttachClaims(next echo.HandlerFunc) echo.HandlerFunc
	RequireClaims(next echo.HandlerFunc) echo.HandlerFunc
	AttachProfile(next echo.HandlerFunc) echo.HandlerFunc
	IsApplicationAdmin(next echo.HandlerFunc) echo.HandlerFunc
	IsProjectAdmin(next echo.HandlerFunc) echo.HandlerFunc
	IsProjectMember(next echo.HandlerFunc) echo.HandlerFunc
	GZIP(next echo.HandlerFunc) echo.HandlerFunc
	CORS(next echo.HandlerFunc) echo.HandlerFunc
	CORSWhitelist(next echo.HandlerFunc) echo.HandlerFunc
	JWT(next echo.HandlerFunc) echo.HandlerFunc
	JWTSkipIfKey(next echo.HandlerFunc) echo.HandlerFunc
	KeyAuth(next echo.HandlerFunc) echo.HandlerFunc
	AppKeyAuth(next echo.HandlerFunc) echo.HandlerFunc
	DataloggerKeyAuth(next echo.HandlerFunc) echo.HandlerFunc
	RequestLogger(next echo.HandlerFunc) echo.HandlerFunc
}

type mw struct {
	cfg                        *config.ServerConfig
	ProfileService             service.ProfileService
	ProjectRoleService         service.ProjectRoleService
	DataloggerTelemetryService service.DataloggerTelemetryService
}

var _ Middleware = (*mw)(nil)

func NewMiddleware(cfg *config.ServerConfig, profileService service.ProfileService, projectRoleService service.ProjectRoleService, dataloggerTelemetryService service.DataloggerTelemetryService) *mw {
	return &mw{cfg, profileService, projectRoleService, dataloggerTelemetryService}
}
