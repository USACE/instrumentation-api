package middleware

import (
	"errors"
	"strconv"
	"strings"

	"github.com/USACE/instrumentation-api/api/internal/httperr"
	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func mapClaims(user *jwt.Token) (model.ProfileClaims, error) {
	claims, ok := user.Claims.(jwt.MapClaims)
	if !ok {
		return model.ProfileClaims{}, errors.New("unable to map claims")
	}

	preferredUsername, ok := claims["preferred_username"].(string)
	if !ok || preferredUsername == "" {
		return model.ProfileClaims{}, errors.New("error parsing token claims: email")
	}
	email, ok := claims["email"].(string)
	if !ok || email == "" {
		return model.ProfileClaims{}, errors.New("error parsing token claims: email")
	}
	name, ok := claims["name"].(string)
	if !ok || name == "" {
		return model.ProfileClaims{}, errors.New("error parsing token claims: name")
	}

	dnClaim, exists := claims["subjectDN"]
	var subjectDN *string
	if exists && dnClaim != nil {
		dnStr, ok := dnClaim.(string)
		if !ok {
			return model.ProfileClaims{}, errors.New("error parsing token claims: subjectDN")
		}
		subjectDN = &dnStr
	}

	cacUIDClaim, exists := claims["cacUID"]
	var cacUID *int
	if exists && cacUIDClaim != nil {
		cacUIDClaims, err := strconv.Atoi(cacUIDClaim.(string))
		if err != nil {
			return model.ProfileClaims{}, errors.New("error parsing token claims: cacUID")
		}
		cacUID = &cacUIDClaims
	}

	x509, _ := claims["x509_presented"].(string)
	x509Presented := false
	if strings.ToLower(x509) == "true" {
		x509Presented = true
	}

	return model.ProfileClaims{
		PreferredUsername: preferredUsername,
		Name:              name,
		Email:             email,
		SubjectDN:         subjectDN,
		CacUID:            cacUID,
		X509Presented:     x509Presented,
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
			return httperr.Forbidden(errors.New("could not get `user` jwt from echo context"))
		}

		claims, err := mapClaims(user)
		if err != nil {
			return httperr.Forbidden(err)
		}

		if email := strings.ToLower(claims.Email); !strings.HasSuffix(email, "usace.army.mil") && !strings.HasSuffix(email, "erdc.dren.mil") && email != "midas@rsgis.dev" {
			return httperr.Forbidden(errors.New("email forbidden"))
		}

		c.Set("claims", claims)

		return next(c)
	}
}

func (m *mw) RequireClaims(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		_, ok := c.Get("claims").(model.ProfileClaims)
		if !ok {
			return httperr.Forbidden(errors.New("no valid claims for user"))
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
				return httperr.Forbidden(err)
			}
			c.Set("profile", p)
			return next(c)
		}

		// If a User was authenticated via KeyAuth, lookup the user's profile using key_id
		if c.Get("KeyAuthSuccess") == true {
			keyID := c.Get("KeyAuthKeyID").(string)
			p, err := m.ProfileService.GetProfileWithTokensForTokenID(ctx, keyID)
			if err != nil {
				return httperr.Forbidden(err)
			}
			c.Set("profile", p)
			return next(c)
		}

		claims, ok := c.Get("claims").(model.ProfileClaims)
		if !ok {
			return httperr.Forbidden(errors.New("could not bind claims from context"))
		}

		p, err := m.ProfileService.GetProfileWithTokensForClaims(ctx, claims)
		if err != nil {
			return httperr.Forbidden(err)
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
			return httperr.Unauthorized(errors.New("could not bind profile from context"))
		}
		if !p.IsAdmin {
			return httperr.ForbiddenRole(errors.New("attempted application admin access failure"))
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
			return httperr.Unauthorized(errors.New("could not bind profile from context"))
		}
		if p.IsAdmin {
			return next(c)
		}
		projectID, err := uuid.Parse(c.Param("project_id"))
		if err != nil {
			return httperr.MalformedID(err)
		}
		authorized, err := m.ProjectRoleService.IsProjectAdmin(c.Request().Context(), p.ID, projectID)
		if err != nil || !authorized {
			return httperr.ForbiddenRole(err)
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
			return httperr.Unauthorized(errors.New("could not bind profile from context"))
		}
		if p.IsAdmin {
			return next(c)
		}
		projectID, err := uuid.Parse(c.Param("project_id"))
		if err != nil {
			return httperr.MalformedID(err)
		}
		authorized, err := m.ProjectRoleService.IsProjectMember(c.Request().Context(), p.ID, projectID)
		if err != nil || !authorized {
			return httperr.ForbiddenRole(err)
		}
		return next(c)
	}
}
