package handlers

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/USACE/instrumentation-api/api/messages"
	"github.com/USACE/instrumentation-api/api/models"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

// CreateProfile creates a user profile
func CreateProfile(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var n models.ProfileInfo
		if err := c.Bind(&n); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		// Set EDIPI
		n.EDIPI = c.Get("EDIPI").(int)

		p, err := models.CreateProfile(db, &n)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusCreated, p)
	}
}

// GetMyProfile returns profile for current authenticated user or 404
func GetMyProfile(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		EDIPI := c.Get("EDIPI").(int)
		p, err := models.GetProfileFromEDIPI(db, EDIPI)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return echo.NewHTTPError(http.StatusNotFound, messages.NotFound)
			}
			return echo.NewHTTPError(http.StatusInternalServerError, messages.InternalServerError)
		}
		return c.JSON(http.StatusOK, &p)
	}
}

// CreateToken returns a list of all products
func CreateToken(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		EDIPI := c.Get("EDIPI").(int)
		p, err := models.GetProfileFromEDIPI(db, EDIPI)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "could not locate user profile with information provided")
		}
		token, err := models.CreateProfileToken(db, &p.ID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusCreated, token)
	}
}

// DeleteToken deletes a token
func DeleteToken(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Get ProfileID
		EDIPI := c.Get("EDIPI").(int)
		p, err := models.GetProfileFromEDIPI(db, EDIPI)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, messages.BadRequest)
		}
		// Get Token ID
		tokenID := c.Param("token_id")
		if tokenID == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "Bad Token ID")
		}
		// Delete Token
		if err := models.DeleteToken(db, &p.ID, &tokenID); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, make(map[string]interface{}))
	}
}
