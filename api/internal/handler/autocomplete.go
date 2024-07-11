package handler

import (
	"github.com/USACE/instrumentation-api/api/internal/httperr"
	"github.com/USACE/instrumentation-api/api/internal/model"

	"net/http"

	"github.com/labstack/echo/v4"
)

// ListEmailAutocomplete godoc
//
//	@Summary lists results of email autocomplete
//	@Tags autocomplete
//	@Produce json
//	@Param q query string true "search query string"
//	@Success 200 {array} model.EmailAutocompleteResult
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /email_autocomplete [get]
func (h *ApiHandler) ListEmailAutocomplete(c echo.Context) error {
	searchText := c.QueryParam("q")
	if searchText == "" {
		return c.JSON(http.StatusOK, make([]model.EmailAutocompleteResult, 0))
	}
	limit := 5
	rr, err := h.EmailAutocompleteService.ListEmailAutocomplete(c.Request().Context(), searchText, limit)
	if err != nil {
		return httperr.InternalServerError(err)
	}
	return c.JSON(http.StatusOK, rr)
}
