package handler

import (
	"net/http"

	_ "github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/labstack/echo/v4"
)

// GetMedia godoc
//
//	@Summary serves media, files, etc for a given project
//	@Tags media
//	@Produce jpeg
//	@Param uri_path path string true "uri path of requested resource"
//	@Success 200
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_slug}/images/{uri_path} [get]
func (h *ApiHandler) GetMedia(c echo.Context) error {
	req := c.Request()

	r, err := h.BlobService.NewReader(req.Context(), req.RequestURI, "")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	c.Response().Header().Set(echo.HeaderContentDisposition, "attachment")
	c.Response().Header().Set("Cache-Control", "public, max-age=31536000")
	return c.Stream(http.StatusOK, "image/jpg", r)
}
