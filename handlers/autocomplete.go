package handlers

import (
	"github.com/USACE/instrumentation-api/models"

	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

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
