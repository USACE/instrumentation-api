package middleware

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/USACE/instrumentation-api/api/models"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
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
				return c.JSON(http.StatusForbidden, models.DefaultMessageUnauthorized)
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
			return c.JSON(http.StatusForbidden, models.DefaultMessageUnauthorized)
		}
		return next(c)
	}
}

// AttachProfileID attaches ProfileID of user to context
func AttachProfileMiddleware(db *sqlx.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// If Application "Superuser" authenticated using Key Authentication (?key= query parameter),
			// lookup superuser profile; the "EDIPI" of the Superuser is consistently 79.
			// The superuser is initialized as part of database and seed data initialization
			if c.Get("ApplicationKeyAuthSuccess") == true {
				p, err := models.GetProfileFromEDIPI(db, 79)
				if err != nil {
					return c.JSON(http.StatusForbidden, models.DefaultMessageUnauthorized)
				}
				c.Set("profile", p)
				return next(c)
			}
			// If a User was authenticated via KeyAuth, lookup the user's profile using key_id
			if c.Get("KeyAuthSuccess") == true {
				keyID := c.Get("KeyAuthKeyID").(string)
				p, err := models.GetProfileFromTokenID(db, keyID)
				if err != nil {
					return c.JSON(http.StatusForbidden, models.DefaultMessageUnauthorized)
				}
				c.Set("profile", p)
				return next(c)
			}
			// If a User was authenticated using CAC (JWT), lookup Profile by EDIPI
			EDIPI := c.Get("EDIPI")
			if EDIPI == nil {
				return c.JSON(http.StatusForbidden, models.DefaultMessageUnauthorized)
			}
			p, err := models.GetProfileFromEDIPI(db, EDIPI.(int))
			if err != nil {
				return c.JSON(http.StatusForbidden, models.DefaultMessageUnauthorized)
			}
			c.Set("profile", p)

			return next(c)
		}
	}
}

func IsApplicationAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		p, ok := c.Get("profile").(*models.Profile)
		if !ok {
			return c.JSON(http.StatusForbidden, models.DefaultMessageUnauthorized)
		}
		if !p.IsAdmin {
			return c.JSON(http.StatusForbidden, models.DefaultMessageUnauthorized)
		}
		return next(c)
	}
}

func IsProjectAdminMiddleware(db *sqlx.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			p, ok := c.Get("profile").(*models.Profile)
			if !ok {
				return c.JSON(http.StatusForbidden, models.DefaultMessageUnauthorized)
			}
			// Application Admins Automatic Admin Status for All Projects
			if p.IsAdmin {
				return next(c)
			}
			// Lookup project from URL Route Parameter
			projectID, err := uuid.Parse(c.Param("project_id"))
			if err != nil {
				return c.JSON(http.StatusForbidden, models.DefaultMessageUnauthorized)
			}
			project, err := models.GetProject(db, projectID)
			if err != nil {
				return c.JSON(http.StatusForbidden, models.DefaultMessageUnauthorized)
			}
			grantingRole := fmt.Sprintf("%s.ADMIN", strings.ToUpper(project.Slug))
			for _, r := range p.Roles {
				if r == grantingRole {
					return next(c)
				}
			}
			return c.JSON(http.StatusForbidden, models.DefaultMessageUnauthorized)
		}
	}
}

func IsProjectMemberMiddleware(db *sqlx.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			p, ok := c.Get("profile").(*models.Profile)
			if !ok {
				return c.JSON(http.StatusForbidden, models.DefaultMessageUnauthorized)
			}
			projectID, err := uuid.Parse(c.Param("project_id"))
			if err != nil {
				return c.JSON(http.StatusForbidden, models.DefaultMessageUnauthorized)
			}
			project, err := models.GetProject(db, projectID)
			if err != nil {
				return c.JSON(http.StatusForbidden, models.DefaultMessageUnauthorized)
			}
			grantingRole := fmt.Sprintf("%s.MEMBER", strings.ToUpper(project.Slug))
			for _, r := range p.Roles {
				if r == grantingRole {
					return next(c)
				}
			}
			return c.JSON(http.StatusForbidden, models.DefaultMessageUnauthorized)
		}
	}
}
