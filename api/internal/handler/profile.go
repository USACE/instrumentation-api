package handler

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/USACE/instrumentation-api/api/internal/message"
	"github.com/USACE/instrumentation-api/api/internal/middleware"
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
	claims := c.Get("claims").(middleware.TokenClaims)

	p := model.ProfileInfo{
		Username: claims.PreferredUsername,
		Email:    claims.Email,
		EDIPI:    claims.CacUID,
	}

	pNew, err := h.ProfileService.CreateProfile(c.Request().Context(), p)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, message.InternalServerError)
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
	claims := c.Get("claims").(middleware.TokenClaims)
	ctx := c.Request().Context()
	var p model.Profile
	var err error
	if claims.CacUID != nil {
		p, err = h.ProfileService.GetProfileWithTokensForEDIPI(ctx, *claims.CacUID)
	} else {
		p, err = h.ProfileService.GetProfileWithTokensForEmail(ctx, claims.Email)
	}
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return echo.NewHTTPError(http.StatusNotFound, message.NotFound)
		}
		return echo.NewHTTPError(http.StatusInternalServerError, message.InternalServerError)
	}
	return c.JSON(http.StatusOK, p)
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
	claims := c.Get("claims").(middleware.TokenClaims)
	ctx := c.Request().Context()
	var p model.Profile
	var err error
	if claims.CacUID != nil {
		p, err = h.ProfileService.GetProfileWithTokensForEDIPI(ctx, *claims.CacUID)
	} else {
		p, err = h.ProfileService.GetProfileWithTokensForEmail(ctx, claims.Email)
	}
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "could not locate user profile with information provided")
	}
	token, err := h.ProfileService.CreateProfileToken(ctx, p.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, message.InternalServerError)
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
	claims := c.Get("claims").(middleware.TokenClaims)
	ctx := c.Request().Context()
	var p model.Profile
	var err error
	if claims.CacUID != nil {
		p, err = h.ProfileService.GetProfileWithTokensForEDIPI(ctx, *claims.CacUID)
	} else {
		p, err = h.ProfileService.GetProfileWithTokensForEmail(ctx, claims.Email)
	}
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.BadRequest)
	}
	tokenID := c.Param("token_id")
	if tokenID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Token ID")
	}
	if err := h.ProfileService.DeleteToken(ctx, p.ID, tokenID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, message.InternalServerError)
	}
	return c.JSON(http.StatusOK, make(map[string]interface{}))
}
