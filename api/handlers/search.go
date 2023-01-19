package handlers

import (
	"fmt"

	"github.com/USACE/instrumentation-api/api/models"

	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type searchFunc func(db *sqlx.DB, searchText *string, limit *int) ([]models.SearchResult, error)

// Search allows searching using a string on different entities
func Search(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Search Function
		var fn searchFunc
		pfn := &fn
		switch entity := c.Param("entity"); entity {
		case "projects":
			*pfn = models.ProjectSearch
		default:
			return c.String(http.StatusBadRequest, fmt.Sprintf("search not implemented for entity: %s", entity))
		}

		// Get Search String
		searchText := c.QueryParam("q")
		if searchText == "" {
			return c.JSON(
				http.StatusOK,
				make([]models.SearchResult, 0),
			)
		}
		// Get Desired Number of Results; Hardcode 5 for now;
		limit := 5
		rr, err := fn(db, &searchText, &limit)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, rr)
	}
}
