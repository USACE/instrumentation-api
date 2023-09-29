package handler

import (
	"context"
	"fmt"

	"github.com/USACE/instrumentation-api/api/internal/model"

	"net/http"

	"github.com/labstack/echo/v4"
)

func (r *ApiRouter) Autocomplete(h *ApiHandler) {
	r.g.public.GET("/search/:entity", h.Search)
}

type searchfunc func(ctx context.Context, searchText string, limit int) ([]model.SearchResult, error)

// Search allows searching using a string on different entities
func (h ApiHandler) Search(c echo.Context) error {
	// Search Function
	var fn searchFunc
	pfn := &fn
	switch entity := c.Param("entity"); entity {
	case "projects":
		*pfn = model.ProjectSearch
	default:
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("search not implemented for entity: %s", entity))
	}

	// Get Search String
	searchText := c.QueryParam("q")
	if searchText == "" {
		return c.JSON(http.StatusOK, make([]model.SearchResult, 0))
	}
	// Get Desired Number of Results; Hardcode 5 for now;
	limit := 5
	rr, err := fn(db, &searchText, &limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, rr)
}
