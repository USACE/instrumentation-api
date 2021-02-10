package handlers

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/USACE/instrumentation-api/models"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

// CreateProfile creates a user profile
func CreateProfile(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var n models.ProfileInfo
		if err := c.Bind(&n); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		// Set EDIPI
		n.EDIPI = c.Get("EDIPI").(int)

		p, err := models.CreateProfile(db, &n)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
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
				return c.JSON(http.StatusNotFound, models.DefaultMessageNotFound)
			}
			return c.JSON(http.StatusInternalServerError, models.DefaultMessageInternalServerError)
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
			return c.String(
				http.StatusBadRequest,
				"could not locate user profile with information provided",
			)
		}
		token, err := models.CreateProfileToken(db, &p.ID)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
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
			return c.JSON(http.StatusBadRequest, models.DefaultMessageBadRequest)
		}
		// Get Token ID
		tokenID := c.Param("token_id")
		if tokenID == "" {
			return c.String(http.StatusBadRequest, "Bad Token ID")
		}
		// Delete Token
		if err := models.DeleteToken(db, &p.ID, &tokenID); err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, make(map[string]interface{}))
	}
}
