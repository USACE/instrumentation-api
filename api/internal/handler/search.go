package handler

import (
	"context"
	"fmt"

	"github.com/USACE/instrumentation-api/api/internal/httperr"
	"github.com/USACE/instrumentation-api/api/internal/model"

	"net/http"

	"github.com/labstack/echo/v4"
)

type searchFunc func(ctx context.Context, searchText string, limit int) ([]model.SearchResult, error)

// Search godoc
//
//	@Summary allows searching using a string on different entities
//	@Tags search
//	@Produce json
//	@Param entity path string true "entity to search (i.e. projects, etc.)"
//	@Param q query string false "search string"
//	@Success 200 {array} model.SearchResult
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /search/{entity} [get]
func (h *ApiHandler) Search(c echo.Context) error {
	var fn searchFunc
	pfn := &fn
	switch entity := c.Param("entity"); entity {
	case "projects":
		*pfn = h.ProjectService.SearchProjects
	default:
		return httperr.Message(http.StatusBadRequest, fmt.Sprintf("search not implemented for entity: %s", entity))
	}

	searchText := c.QueryParam("q")
	if searchText == "" {
		return c.JSON(http.StatusOK, make([]model.SearchResult, 0))
	}
	// Get Desired Number of Results; Hardcode 5 for now;
	limit := 5
	rr, err := fn(c.Request().Context(), searchText, limit)
	if err != nil {
		return httperr.InternalServerError(err)
	}
	return c.JSON(http.StatusOK, rr)
}
