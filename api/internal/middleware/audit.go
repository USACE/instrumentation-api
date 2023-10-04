package middleware

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/USACE/instrumentation-api/api/internal/message"
	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// EDIPIMiddleware attaches EDIPI (CAC) to Context
// Used for CAC-Only Routes
func (m *mw) EDIPI(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		key := c.QueryParam("key")
		if key == "" {
			user, ok := c.Get("user").(*jwt.Token)
			if !ok {
				return echo.NewHTTPError(http.StatusForbidden, message.Unauthorized)
			}
			claims, ok := user.Claims.(jwt.MapClaims)
			if !ok {
				return echo.NewHTTPError(http.StatusForbidden, message.Unauthorized)
			}
			edipi, err := strconv.Atoi(claims["sub"].(string))
			if err != nil {
				return echo.NewHTTPError(http.StatusForbidden, message.Unauthorized)
			}
			c.Set("EDIPI", edipi)
		}
		return next(c)
	}
}

func (m *mw) CACOnly(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		edipi := c.Get("EDIPI")
		if edipi == nil {
			return echo.NewHTTPError(http.StatusForbidden, message.Unauthorized)
		}
		return next(c)
	}
}

// AttachProfileID attaches ProfileID of user to context
func (m *mw) AttachProfile(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		// If Application "Superuser" authenticated using Key Authentication (?key= query parameter),
		// lookup superuser profile; the "EDIPI" of the Superuser is consistently 79.
		// The superuser is initialized as part of database and seed data initialization
		if c.Get("ApplicationKeyAuthSuccess") == true {
			p, err := m.ProfileService.GetProfileWithTokensFromEDIPI(ctx, 79)
			if err != nil {
				return echo.NewHTTPError(http.StatusForbidden, message.Unauthorized)
			}
			c.Set("profile", p)
			return next(c)
		}
		// If a User was authenticated via KeyAuth, lookup the user's profile using key_id
		if c.Get("KeyAuthSuccess") == true {
			keyID := c.Get("KeyAuthKeyID").(string)
			p, err := m.ProfileService.GetProfileFromTokenID(ctx, keyID)
			if err != nil {
				return echo.NewHTTPError(http.StatusForbidden, message.Unauthorized)
			}
			c.Set("profile", p)
			return next(c)
		}
		// If a User was authenticated using CAC (JWT), lookup Profile by edipi
		edipi := c.Get("EDIPI")
		if edipi == nil {
			return echo.NewHTTPError(http.StatusForbidden, message.Unauthorized)
		}
		p, err := m.ProfileService.GetProfileWithTokensFromEDIPI(ctx, edipi.(int))
		if err != nil {
			return echo.NewHTTPError(http.StatusForbidden, message.Unauthorized)
		}
		c.Set("profile", p)

		return next(c)
	}
}

func (m *mw) IsApplicationAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		p, ok := c.Get("profile").(model.Profile)
		if !ok {
			return echo.NewHTTPError(http.StatusForbidden, message.Unauthorized)
		}
		if !p.IsAdmin {
			return echo.NewHTTPError(http.StatusForbidden, message.Unauthorized)
		}
		return next(c)
	}
}

func (m *mw) IsProjectAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		p, ok := c.Get("profile").(model.Profile)
		if !ok {
			return echo.NewHTTPError(http.StatusForbidden, message.Unauthorized)
		}
		// Application Admins Automatic Admin Status for All Projects
		if p.IsAdmin {
			return next(c)
		}
		// Lookup project from URL Route Parameter
		projectID, err := uuid.Parse(c.Param("project_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusForbidden, message.Unauthorized)
		}
		project, err := m.ProjectService.GetProject(c.Request().Context(), projectID)
		if err != nil {
			return echo.NewHTTPError(http.StatusForbidden, message.Unauthorized)
		}
		grantingRole := fmt.Sprintf("%s.ADMIN", strings.ToUpper(project.Slug))
		for _, r := range p.Roles {
			if r == grantingRole {
				return next(c)
			}
		}
		return echo.NewHTTPError(http.StatusForbidden, message.Unauthorized)
	}
}

func (m *mw) IsProjectMember(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		p, ok := c.Get("profile").(model.Profile)
		if !ok {
			return echo.NewHTTPError(http.StatusForbidden, message.Unauthorized)
		}
		projectID, err := uuid.Parse(c.Param("project_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusForbidden, message.Unauthorized)
		}
		project, err := m.ProjectService.GetProject(c.Request().Context(), projectID)
		if err != nil {
			return echo.NewHTTPError(http.StatusForbidden, message.Unauthorized)
		}
		grantingRole := fmt.Sprintf("%s.MEMBER", strings.ToUpper(project.Slug))
		for _, r := range p.Roles {
			if r == grantingRole {
				return next(c)
			}
		}
		return echo.NewHTTPError(http.StatusForbidden, message.Unauthorized)
	}
}
