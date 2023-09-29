package handler

import (
	"github.com/USACE/instrumentation-api/api/internal/model"

	"net/http"

	"github.com/labstack/echo/v4"
)

// ListEmailAutocomplete lists results of email autocomplete
func (h ApiHandler) ListEmailAutocomplete(c echo.Context) error {
	// Get Search String
	searchText := c.QueryParam("q")
	if searchText == "" {
		return c.JSON(http.StatusOK, make([]model.EmailAutocompleteResult, 0))
	}
	// Get Desired Number of Results; Hardcode 5 for now;
	limit := 5
	rr, err := h.EmailAutocompleteStore.ListEmailAutocomplete(c.Request().Context(), searchText, limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, rr)
}
