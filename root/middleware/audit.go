package middleware

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
		c.Set("action", c.Request().Method)
		c.Set("action_time", time.Now())

		// Get claims from JWT
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)

		// Set fields for creator and updater based on request time and JWT claims
		userID, err := strconv.Atoi(claims["sub"].(string))
		if err != nil {
			return c.NoContent(http.StatusForbidden)
		}
		c.Set("actor", userID)

		userRoles := claims["roles"].([]interface{})
		c.Set("actor_roles", userRoles)

		return next(c)
	}
}

// MockIsLoggedIn sets appropriate actor fields based on the time of the request and mocked user information
// Necessary for running integration tests unauthenticated
func MockIsLoggedIn(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set("action", c.Request().Method)
		c.Set("action_time", time.Now())
		c.Set("actor", 0)
		c.Set("actor_roles", "TEST")

		return next(c)
	}
}
