package handler

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/USACE/instrumentation-api/api/internal/message"
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
func (h *ApiHandler) CreateProfile(c echo.Context) error {
	var n model.ProfileInfo
	if err := c.Bind(&n); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	n.EDIPI = c.Get("EDIPI").(int)

	p, err := h.ProfileService.CreateProfile(c.Request().Context(), n)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, p)
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
func (h *ApiHandler) GetMyProfile(c echo.Context) error {
	edipi := c.Get("EDIPI").(int)
	p, err := h.ProfileService.GetProfileWithTokensFromEDIPI(c.Request().Context(), edipi)
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
func (h *ApiHandler) CreateToken(c echo.Context) error {
	edipi := c.Get("EDIPI").(int)
	ctx := c.Request().Context()
	p, err := h.ProfileService.GetProfileWithTokensFromEDIPI(ctx, edipi)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "could not locate user profile with information provided")
	}
	token, err := h.ProfileService.CreateProfileToken(ctx, p.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
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
func (h *ApiHandler) DeleteToken(c echo.Context) error {
	// Get ProfileID
	edipi := c.Get("EDIPI").(int)
	ctx := c.Request().Context()
	p, err := h.ProfileService.GetProfileWithTokensFromEDIPI(ctx, edipi)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.BadRequest)
	}
	// Get Token ID
	tokenID := c.Param("token_id")
	if tokenID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Token ID")
	}
	// Delete Token
	if err := h.ProfileService.DeleteToken(ctx, p.ID, tokenID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, make(map[string]interface{}))
}
