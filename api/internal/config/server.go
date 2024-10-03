package config

type ServerConfig struct {
	AllowOrigins           []string `env:"ALLOWED_ORIGINS"`
	ApplicationKey         string   `env:"APPLICATION_KEY"`
	AuthDisabled           bool     `env:"AUTH_DISABLED"`
	AuthJWTMocked          bool     `env:"AUTH_JWT_MOCKED"`
	AuthAllowEmailSuffixes []string `env:"AUTH_ALLOW_EMAIL_SUFFIXES"`
	AuthPublicKey          string   `env:"AUTH_PUBLIC_KEY"`
	AuthSigningMethod      string   `env:"AUTH_SIGNING_METHOD"`
	Debug                  bool     `env:"DEBUG"`
	RequestLoggerEnabled   bool     `env:"REQUEST_LOGGER_ENABLED"`
	RoutePrefix            string   `env:"ROUTE_PREFIX"`
	ServerBaseUrl          string   `env:"SERVER_BASE_URL"`
}
