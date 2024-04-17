package config

type ServerConfig struct {
	AuthDisabled         bool   `envconfig:"AUTH_DISABLED"`
	AuthJWTMocked        bool   `envconfig:"AUTH_JWT_MOCKED"`
	ApplicationKey       string `envconfig:"APPLICATION_KEY"`
	RequestLoggerEnabled bool   `envconfig:"REQUEST_LOGGER_ENABLED"`
	RoutePrefix          string `envconfig:"ROUTE_PREFIX"`
}
