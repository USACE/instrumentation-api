package config

type ServerConfig struct {
	AllowOrigins         []string `envconfig:"ALLOW_ORIGINS"`
	ApplicationKey       string   `envconfig:"APPLICATION_KEY"`
	AuthDisabled         bool     `envconfig:"AUTH_DISABLED"`
	AuthJWTMocked        bool     `envconfig:"AUTH_JWT_MOCKED"`
	AuthPublicKey        string   `envconfig:"AUTH_PUBLIC_KEY"`
	AuthSigningMethod    string   `envconfig:"AUTH_SIGNING_METHOD"`
	Debug                bool     `envconfig:"DEBUG"`
	RequestLoggerEnabled bool     `envconfig:"REQUEST_LOGGER_ENABLED"`
	RoutePrefix          string   `envconfig:"ROUTE_PREFIX"`
	ServerBaseUrl        string   `envconfig:"SERVER_BASE_URL"`
}
