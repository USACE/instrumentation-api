package handler

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/USACE/instrumentation-api/api/internal/httperr"
	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/labstack/echo/v4"
)

// CreateProfile godoc
//
//	@Summary creates a user profile
//	@Tags profile
//	@Produce json
//	@Success 200 {object} model.Profile
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /profiles [post]
//	@Security ClaimsOnly
func (h *ApiHandler) CreateProfile(c echo.Context) error {
	claims := c.Get("claims").(model.ProfileClaims)

	if !claims.X509Presented {
		return httperr.Forbidden(errors.New("invalid value for claim x509_presented"))
	}
	if claims.CacUID == nil {
		return httperr.Forbidden(errors.New("unable to create profile; cacUID claim is nil"))
	}

	p := model.ProfileInfo{
		Username:    claims.PreferredUsername,
		DisplayName: claims.Name,
		Email:       claims.Email,
		EDIPI:       *claims.CacUID,
	}

	pNew, err := h.ProfileService.CreateProfile(c.Request().Context(), p)
	if err != nil {
		return httperr.InternalServerError(err)
	}
	return c.JSON(http.StatusCreated, pNew)
}

// GetMyProfile godoc
//
//	@Summary gets profile for current authenticated user
//	@Tags profile
//	@Produce json
//	@Success 200 {object} model.Profile
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /my_profile [get]
//	@Security ClaimsOnly
func (h *ApiHandler) GetMyProfile(c echo.Context) error {
	ctx := c.Request().Context()
	claims := c.Get("claims").(model.ProfileClaims)

	p, err := h.ProfileService.GetProfileWithTokensForClaims(ctx, claims)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return h.CreateProfile(c)
		}
		return httperr.InternalServerError(err)
	}

	pValidated, err := h.ProfileService.UpdateProfileForClaims(ctx, p, claims)
	if err != nil {
		return httperr.InternalServerError(err)
	}

	return c.JSON(http.StatusOK, pValidated)
}

// CreateToken godoc
//
//	@Summary creates token for a profile
//	@Tags profile
//	@Produce json
//	@Success 200 {object} model.Token
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /my_tokens [post]
//	@Security ClaimsOnly
func (h *ApiHandler) CreateToken(c echo.Context) error {
	claims := c.Get("claims").(model.ProfileClaims)
	ctx := c.Request().Context()

	p, err := h.ProfileService.GetProfileWithTokensForClaims(ctx, claims)
	if err != nil {
		return httperr.InternalServerError(err)
	}
	token, err := h.ProfileService.CreateProfileToken(ctx, p.ID)
	if err != nil {
		return httperr.InternalServerError(err)
	}
	return c.JSON(http.StatusCreated, token)
}

// DeleteToken godoc
//
//	@Summary deletes a token for a profile
//	@Tags profile
//	@Produce json
//	@Param token_id path string true "token uuid" Format(uuid)
//	@Success 200 {object} map[string]interface{}
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /my_tokens/{token_id} [delete]
//	@Security ClaimsOnly
func (h *ApiHandler) DeleteToken(c echo.Context) error {
	claims := c.Get("claims").(model.ProfileClaims)
	ctx := c.Request().Context()

	tokenID := c.Param("token_id")
	if tokenID == "" {
		return httperr.Message(http.StatusBadRequest, "bad token id")
	}

	p, err := h.ProfileService.GetProfileWithTokensForClaims(ctx, claims)
	if err != nil {
		return httperr.InternalServerError(err)
	}
	if err := h.ProfileService.DeleteToken(ctx, p.ID, tokenID); err != nil {
		return httperr.InternalServerError(err)
	}
	return c.JSON(http.StatusOK, make(map[string]interface{}))
}
