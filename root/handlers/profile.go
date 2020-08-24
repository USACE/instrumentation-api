package handlers

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"

	"api/root/models"
)

// GetMyProfile returns profile for credentials or returns 404
func GetMyProfile(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var a models.Action
		if err := c.Bind(&a); err != nil {
			return c.NoContent(http.StatusBadRequest)
		}
		p, err := models.GetMyProfile(db, &a)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return c.NoContent(http.StatusNotFound)
			}
			return c.NoContent(http.StatusInternalServerError)
		}
		return c.JSON(http.StatusOK, &p)
	}
}

// CreateProfile creates a new profile
func CreateProfile(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var r models.CreateProfileRequest
		if err := c.Bind(&r); err != nil {
			return c.NoContent(http.StatusBadRequest)
		}
		p, err := models.CreateProfile(db, &r)
		if err != nil {
			return c.NoContent(http.StatusBadRequest)
		}
		return c.JSON(http.StatusCreated, &p)
	}
}
