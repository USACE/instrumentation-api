package handlers

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/USACE/instrumentation-api/models"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func edipiFromContext(c echo.Context) int {
	return c.Get("actor").(int)
}

// profileFromContext is a helper function that uses the current context to return
// the corresponding profile
func profileFromContext(c echo.Context, db *sqlx.DB) (*models.Profile, error) {
	if edipi, ok := c.Get("actor").(int); ok {
		p, err := models.GetProfileFromEDIPI(db, edipi)
		if err != nil {
			return nil, err
		}
		return p, nil
	}
	return nil, nil
}

// CreateProfile creates a user profile
func CreateProfile(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var n models.ProfileInfo
		if err := c.Bind(&n); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		// Set EDIPI
		n.EDIPI = edipiFromContext(c)

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
		p, err := profileFromContext(c, db)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return c.NoContent(http.StatusNotFound)
			}
			return c.NoContent(http.StatusInternalServerError)
		}
		return c.JSON(http.StatusOK, &p)
	}
}

// CreateToken returns a list of all products
func CreateToken(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		p, err := profileFromContext(c, db)
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
		p, err := profileFromContext(c, db)
		if err != nil {
			return c.NoContent(http.StatusBadRequest)
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
