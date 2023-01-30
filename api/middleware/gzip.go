package middleware

import (
	"github.com/labstack/echo/v4/middleware"
)

// GZIP is ready-to-go GZIP middleware based on echo middleware
var GZIP = middleware.GzipWithConfig(middleware.GzipConfig{Level: 5})
