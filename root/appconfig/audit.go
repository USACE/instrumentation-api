package appconfig

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// SetCreatorUpdaterFields sets appropriate Creator, CreateDate, Updater, UpdateDate fields
// based on the time of the request and the user claims in the JWT
func SetCreatorUpdaterFields(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// If Skipping JWT, skip this middleware too
		if skipJWT(c) {
			return next(c)
		}

		// Get claims from JWT
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)

		// Set fields for creator and updater based on request time and JWT claims
		userID := claims["sub"].(int)
		t := time.Now()
		c.Set("creater", userID)
		c.Set("updater", userID)
		c.Set("update_date", t)
		c.Set("create_date", t)

		return next(c)
	}
}
