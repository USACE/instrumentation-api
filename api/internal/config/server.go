package config

type ServerConfig struct {
	ApplicationKey       string `envconfig:"APPLICATION_KEY"`
	AuthDisabled         bool   `envconfig:"AUTH_DISABLED"`
	AuthJWTMocked        bool   `envconfig:"AUTH_JWT_MOCKED"`
	AuthPublicKey        string `envconfig:"AUTH_PUBLIC_KEY"`
	RequestLoggerEnabled bool   `envconfig:"REQUEST_LOGGER_ENABLED"`
	RoutePrefix          string `envconfig:"ROUTE_PREFIX"`
	ServerBaseUrl        string `envconfig:"SERVER_BASE_URL"`
}
