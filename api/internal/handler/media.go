package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetMedia serves media, files, etc for a given project
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
