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
		a, err := models.NewAction(c)
		if err != nil {
			return c.NoContent(http.StatusBadRequest)
		}
		p, err := models.GetMyProfile(db, a)
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
			return c.JSON(http.StatusBadRequest, err)
		}
		a, err := models.NewAction(c)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		r.Action = *a
		p, err := models.CreateProfile(db, &r)
		if err != nil {
			return c.NoContent(http.StatusBadRequest)
		}
		return c.JSON(http.StatusCreated, &p)
	}
}

// ListEmailAutocomplete lists results of email autocomplete
func ListEmailAutocomplete(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Get Search String
		searchText := c.QueryParam("q")
		if searchText == "" {
			return c.JSON(
				http.StatusOK,
				make([]models.EmailAutocompleteResult, 0),
			)
		}
		// Get Desired Number of Results; Hardcode 5 for now;
		limit := 5
		rr, err := models.ListEmailAutocomplete(db, &searchText, &limit)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, rr)
	}
}
