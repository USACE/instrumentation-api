package handler

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/USACE/instrumentation-api/api/internal/message"
	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/labstack/echo/v4"
)

// CreateProfile creates a user profile
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

// GetMyProfile returns profile for current authenticated user or 404
func (h *ApiHandler) GetMyProfile(c echo.Context) error {
	edipi := c.Get("EDIPI").(int)
	p, err := h.ProfileService.GetProfileWithTokensFromEDIPI(c.Request().Context(), edipi)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return echo.NewHTTPError(http.StatusNotFound, message.NotFound)
		}
		return echo.NewHTTPError(http.StatusInternalServerError, message.InternalServerError)
	}
	return c.JSON(http.StatusOK, &p)
}

// CreateToken returns a list of all products
func (h *ApiHandler) CreateToken(c echo.Context) error {
	edipi := c.Get("EDIPI").(int)
	p, err := h.ProfileService.GetProfileWithTokensFromEDIPI(c.Request().Context(), edipi)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "could not locate user profile with information provided")
	}
	token, err := h.ProfileService.CreateProfileToken(c.Request().Context(), p.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, token)
}

// DeleteToken deletes a token
func (h *ApiHandler) DeleteToken(c echo.Context) error {
	// Get ProfileID
	edipi := c.Get("EDIPI").(int)
	p, err := h.ProfileService.GetProfileWithTokensFromEDIPI(c.Request().Context(), edipi)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.BadRequest)
	}
	// Get Token ID
	tokenID := c.Param("token_id")
	if tokenID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Token ID")
	}
	// Delete Token
	if err := h.ProfileService.DeleteToken(c.Request().Context(), p.ID, tokenID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, make(map[string]interface{}))
}
