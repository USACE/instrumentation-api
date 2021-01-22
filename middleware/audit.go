package middleware

import (
	"net/http"
	"strconv"

	"github.com/USACE/instrumentation-api/models"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

// EDIPIMiddleware attaches EDIPI (CAC) to Context
// Used for CAC-Only Routes
func EDIPIMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// If key is in query parameters; count on keyauth
		key := c.QueryParam("key")
		if key == "" {
			user := c.Get("user").(*jwt.Token)
			claims := user.Claims.(jwt.MapClaims)
			// Get EDIPI
			EDIPI, err := strconv.Atoi(claims["sub"].(string))
			if err != nil {
				return c.NoContent(http.StatusForbidden)
			}
			c.Set("EDIPI", EDIPI)
		}
		return next(c)
	}
}

func CACOnlyMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		EDIPI := c.Get("EDIPI")
		if EDIPI == nil {
			return c.NoContent(http.StatusForbidden)
		}
		return next(c)

	}
}

// AttachProfileID attaches ProfileID of user to context
// If UsedKeyAuth, looks profile id up by token_id used
// Otherwise, looks-up ProfileID by EDIPI in JWT
func AttachProfileMiddleware(db *sqlx.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// If User authenticated via KeyAuth, lookup profile via key_id
			if c.Get("KeyAuthSuccess") == true {
				keyID := c.Get("KeyAuthKeyID").(string)
				p, err := models.GetProfileFromTokenID(db, keyID)
				if err != nil {
					return c.NoContent(http.StatusForbidden)
				}
				c.Set("profile", p)
				return next(c)
			}
			// Lookup Profile by EDIPI
			EDIPI := c.Get("EDIPI")
			if EDIPI == nil {
				return c.NoContent(http.StatusForbidden)
			}
			p, err := models.GetProfileFromEDIPI(db, EDIPI.(int))
			if err != nil {
				return c.NoContent(http.StatusForbidden)
			}
			c.Set("profile", p)
			// userRoles := claims["roles"].([]interface{})
			// c.Set("actor_roles", userRoles)
			return next(c)
		}
	}
}
