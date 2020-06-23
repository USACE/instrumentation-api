package appconfig

import (
	"net/http"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// IsLoggedIn sets appropriate actor fields
// based on the time of the request and the user claims in the JWT
func IsLoggedIn(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// If Skipping JWT, skip this middleware too
		if skipJWT(c) {
			return next(c)
		}

		// Get claims from JWT
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)

		// Set fields for creator and updater based on request time and JWT claims
		userID, err := strconv.Atoi(claims["sub"].(string))
		if err != nil {
			return c.NoContent(http.StatusForbidden)
		}

		userRoles := claims["roles"].([]interface{})

		t := time.Now()
		c.Set("actor", userID)
		c.Set("actor_roles", userRoles)
		c.Set("action", c.Request().Method)
		c.Set("action_time", t)

		return next(c)
	}
}
