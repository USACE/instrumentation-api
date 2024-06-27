package middleware

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/USACE/instrumentation-api/api/internal/message"
	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type TokenClaims struct {
	PreferredUsername string
	Email             string
	SubjectDN         *string
	CacUID            *int
}

func mapClaims(user *jwt.Token) (TokenClaims, error) {
	claims, ok := user.Claims.(jwt.MapClaims)
	if !ok {
		return TokenClaims{}, errors.New("unable to map claims")
	}

	// common claims, required
	pu, ok := claims["preferred_username"].(string)
	if !ok || pu == "" {
		return TokenClaims{}, errors.New("error parsing token claims: email")
	}
	email, ok := claims["email"].(string)
	if !ok || email == "" {
		return TokenClaims{}, errors.New("error parsing token claims: email")
	}

	// cac-specific claims, for cac users only
	dnClaim, exists := claims["subjectDN"]
	var dn *string
	if exists && dnClaim != nil {
		dnStr, ok := dnClaim.(string)
		if !ok {
			return TokenClaims{}, errors.New("error parsing token claims: subjectDN")
		}
		dn = &dnStr
	}

	cacUIDClaim, exists := claims["cacUID"]
	var cacUID *int
	if exists && cacUIDClaim != nil {
		cacUIDClaims, err := strconv.Atoi(cacUIDClaim.(string))
		if err != nil {
			return TokenClaims{}, errors.New("error parsing token claims: cacUID")
		}
		cacUID = &cacUIDClaims
	}

	return TokenClaims{
		PreferredUsername: pu,
		Email:             email,
		SubjectDN:         dn,
		CacUID:            cacUID,
	}, nil
}

func (m *mw) AttachClaims(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		key := c.QueryParam("key")
		if key != "" {
			return next(c)
		}

		user, ok := c.Get("user").(*jwt.Token)
		if !ok {
			return echo.NewHTTPError(http.StatusForbidden, message.Unauthorized)
		}

		claims, err := mapClaims(user)
		if err != nil {
			return echo.NewHTTPError(http.StatusForbidden, message.Unauthorized)
		}
		c.Set("claims", claims)

		return next(c)
	}
}

func (m *mw) RequireClaims(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		claims, ok := c.Get("claims").(TokenClaims)
		if !ok || (claims.CacUID == nil && claims.Email == "") {
			return echo.NewHTTPError(http.StatusForbidden, message.Unauthorized)
		}
		return next(c)
	}
}

// AttachProfile attaches the Profile of a user to request context
func (m *mw) AttachProfile(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		// If Application "Superuser" authenticated using Key Authentication (?key= query parameter),
		// lookup superuser profile; the "EDIPI" of the Superuser is consistently 79.
		// The superuser is initialized as part of database and seed data initialization
		if c.Get("ApplicationKeyAuthSuccess") == true {
			p, err := m.ProfileService.GetProfileWithTokensForEDIPI(ctx, 79)
			if err != nil {
				return echo.NewHTTPError(http.StatusForbidden, message.Unauthorized)
			}
			c.Set("profile", p)
			return next(c)
		}
		// If a User was authenticated via KeyAuth, lookup the user's profile using key_id
		if c.Get("KeyAuthSuccess") == true {
			keyID := c.Get("KeyAuthKeyID").(string)
			p, err := m.ProfileService.GetProfileWithTokensForTokenID(ctx, keyID)
			if err != nil {
				return echo.NewHTTPError(http.StatusForbidden, message.Unauthorized)
			}
			c.Set("profile", p)
			return next(c)
		}
		claims, ok := c.Get("claims").(TokenClaims)
		if !ok {
			return echo.NewHTTPError(http.StatusForbidden, message.Unauthorized)
		}

		if claims.CacUID != nil {
			p, err := m.ProfileService.GetProfileWithTokensForEDIPI(ctx, *claims.CacUID)
			if err != nil {
				return echo.NewHTTPError(http.StatusForbidden, message.Unauthorized)
			}
			if p.Username != claims.PreferredUsername || p.Email != claims.Email {
				if err := m.ProfileService.UpdateProfileForEDIPI(ctx, claims.PreferredUsername, claims.Email, *claims.CacUID); err != nil {
					return echo.NewHTTPError(http.StatusForbidden, message.Unauthorized)
				}
				p.Username = claims.PreferredUsername
				p.Email = claims.Email
			}
			c.Set("profile", p)
		} else if claims.Email != "" {
			p, err := m.ProfileService.GetProfileWithTokensForEmail(ctx, claims.Email)
			if err != nil {
				return echo.NewHTTPError(http.StatusForbidden, message.Unauthorized)
			}
			if p.Username != claims.PreferredUsername {
				if err := m.ProfileService.UpdateProfileForEmail(ctx, claims.PreferredUsername, claims.Email); err != nil {
					return echo.NewHTTPError(http.StatusForbidden, message.Unauthorized)
				}
				p.Username = claims.PreferredUsername
			}
			c.Set("profile", p)
		} else {
			return echo.NewHTTPError(http.StatusForbidden, message.Unauthorized)
		}

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
