package middleware

import (
	"context"

	"github.com/USACE/instrumentation-api/api/internal/password"

	"github.com/labstack/echo/v4"
	echomw "github.com/labstack/echo/v4/middleware"
)

func (m *mw) AppKeyAuth(next echo.HandlerFunc) echo.HandlerFunc {
	ka := echomw.KeyAuthWithConfig(echomw.KeyAuthConfig{
		KeyLookup: "query:key",
		Validator: func(key string, c echo.Context) (bool, error) {
			return key == m.cfg.ApplicationKey, nil
		},
	})
	return ka(next)
}

type HashExtractorFunc func(keyID string) (string, error)

// KeyAuth returns a ready-to-go key auth middleware
func keyAuth(isDisabled bool, appKey string, h HashExtractorFunc) echo.MiddlewareFunc {
	return echomw.KeyAuthWithConfig(
		echomw.KeyAuthConfig{
			KeyLookup: "query:key",
			// If Auth Manually Disabled via Environment Variable
			// or "?key=..." is not in QueryParams (counting on other auth middleware
			// further down the chain); Skip this middleware
			Skipper: func(c echo.Context) bool {
				if isDisabled || c.QueryParam("key") == "" {
					return true
				}
				return false
			},
			// Compare key passed via query parameters with hash serviced in the database
			Validator: func(key string, c echo.Context) (bool, error) {
				// If Key is Master ApplicationKey; Grant Access
				if key == appKey {
					c.Set("ApplicationKeyAuthSuccess", true)
					return true, nil
				}
				// Check Key against serviced hash in the database
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
				match, err := password.ComparePasswordAndHash(key, hash)
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
		},
	)
}

func getHashExtractorFunc(ctx context.Context, m *mw) HashExtractorFunc {
	return func(keyID string) (string, error) {
		k, err := m.ProfileService.GetTokenInfoByTokenID(ctx, keyID)
		if err != nil {
			return "", err
		}
		return k.Hash, nil
	}
}

func (m *mw) KeyAuth(next echo.HandlerFunc) echo.HandlerFunc {
	ka := keyAuth(m.cfg.AuthDisabled, m.cfg.ApplicationKey, getHashExtractorFunc(context.TODO(), m))
	return ka(next)
}

type DataloggerHashExtractorFunc func(modelName, sn string) (string, error)

func getDataloggerHashExtractorFunc(ctx context.Context, m *mw) DataloggerHashExtractorFunc {
	return func(modelName, sn string) (string, error) {
		hash, err := m.DataloggerTelemetryService.GetDataloggerHashByModelSN(ctx, modelName, sn)
		if err != nil {
			return "", err
		}
		return hash, nil
	}
}

// DataLoggerKeyAuth returns key auth for data logger model / serial number lookup
func dataloggerKeyAuth(h DataloggerHashExtractorFunc) echo.MiddlewareFunc {
	return echomw.KeyAuthWithConfig(
		echomw.KeyAuthConfig{
			KeyLookup: "header:X-Api-Key",
			Validator: func(key string, c echo.Context) (bool, error) {
				modelName := c.Param("model")
				if modelName == "" {
					return false, nil
				}
				sn := c.Param("sn")
				if sn == "" {
					return false, nil
				}

				hash, err := h(modelName, sn)
				if err != nil {
					return false, nil
				}

				match, err := password.ComparePasswordAndHash(key, hash)
				if err != nil {
					return false, err
				}

				if match {
					return true, nil
				}

				return false, nil
			},
		},
	)
}

func (m *mw) DataloggerKeyAuth(next echo.HandlerFunc) echo.HandlerFunc {
	ka := dataloggerKeyAuth(getDataloggerHashExtractorFunc(context.TODO(), m))
	return ka(next)
}
