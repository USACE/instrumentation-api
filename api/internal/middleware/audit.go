package middleware

import (
	"net/http"
	"strconv"

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
			p, err := m.ProfileService.GetProfileWithTokensFromTokenID(ctx, keyID)
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

// IsApplicationAdmin checks that a profile is an application admin
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

// IsProjectAdmin checks that a profile is an admin for the project_id URL path parameter
// ApplicationAdmin has automatic member/admin status for all projects
func (m *mw) IsProjectAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		p, ok := c.Get("profile").(model.Profile)
		if !ok {
			return echo.NewHTTPError(http.StatusForbidden, message.Unauthorized)
		}
		if p.IsAdmin {
			return next(c)
		}
		projectID, err := uuid.Parse(c.Param("project_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusForbidden, message.Unauthorized)
		}
		authorized, err := m.ProjectRoleService.IsProjectAdmin(c.Request().Context(), p.ID, projectID)
		if err != nil || !authorized {
			return echo.NewHTTPError(http.StatusForbidden, message.Unauthorized)
		}
		return next(c)
	}
}

// IsProjectMember checks that a profile is a member or admin for the project_id URL path parameter
// ApplicationAdmin has automatic member/admin status for all projects
func (m *mw) IsProjectMember(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		p, ok := c.Get("profile").(model.Profile)
		if !ok {
			return echo.NewHTTPError(http.StatusForbidden, message.Unauthorized)
		}
		if p.IsAdmin {
			return next(c)
		}
		projectID, err := uuid.Parse(c.Param("project_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusForbidden, message.Unauthorized)
		}
		authorized, err := m.ProjectRoleService.IsProjectMember(c.Request().Context(), p.ID, projectID)
		if err != nil || !authorized {
			return echo.NewHTTPError(http.StatusForbidden, message.Unauthorized)
		}
		return next(c)
	}
}
