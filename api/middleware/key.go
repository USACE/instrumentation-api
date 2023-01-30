package middleware

import (
	"github.com/USACE/instrumentation-api/api/passwords"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// HashExtractor returns a hash (string) to be compared with user supplied key
type HashExtractor func(keyID string) (string, error)

// KeyAuth returns a ready-to-go key auth middleware
func KeyAuth(isDisabled bool, appKey string, h HashExtractor) echo.MiddlewareFunc {
	return middleware.KeyAuthWithConfig(
		middleware.KeyAuthConfig{
			// If Auth Manually Disabled via Environment Variable
			// or "?key=..." is not in QueryParams (counting on other auth middleware
			// further down the chain); Skip this middleware
			Skipper: func(c echo.Context) bool {
				if isDisabled || c.QueryParam("key") == "" {
					return true
				}
				return false
			},
			// Compare key passed via query parameters with hash stored in the database
			Validator: func(key string, c echo.Context) (bool, error) {
				// If Key is Master ApplicationKey; Grant Access
				////////////////////////////////////////////////
				if key == appKey {
					c.Set("ApplicationKeyAuthSuccess", true)
					return true, nil
				}
				// Check Key against stored hash in the database
				////////////////////////////////////////////////
				// If key_id not provided as query parameter; Deny Access
				keyID := c.QueryParam("key_id")
				if keyID == "" {
					return false, nil
				}
				// Lookup hash in database using key_id; Deny Access if not found or error
				hash, err := h(keyID)
				if err != nil {
					return false, nil
				}
				// Compare provided key with key hash; Deny Access if error
				match, err := passwords.ComparePasswordAndHash(key, hash)
				if err != nil {
					return false, err
				}
				// Compare provided key with key hash; Allow access if match
				if match {
					c.Set("KeyAuthSuccess", true)
					c.Set("KeyAuthKeyID", keyID)
					return true, nil
				}
				// Deny Access otherwise
				return false, nil
			},
			KeyLookup: "query:key",
		},
	)
}
