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
//	@Param project_slug path string true "project abbr"
//	@Param uri_path path string true "uri path of requested resource"
//	@Success 200
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_slug}/images/{uri_path} [get]
func (h *ApiHandler) GetMedia(c echo.Context) error {
	req := c.Request()
	r, err := h.BlobService.NewReaderContext(req.Context(), req.RequestURI, "")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	c.Response().Header().Set(echo.HeaderContentDisposition, "attachment")
	c.Response().Header().Set("Cache-Control", "public, max-age=31536000")
	return c.Stream(http.StatusOK, "image/jpg", r)
}

// func (h *ApiHandler) UploadMedia(c echo.Context) error {
// 	file, err := c.FormFile("image")
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, "attached form file 'image' required")
// 	}
// 	src, err := file.Open()
// 	if err != nil {
// 		return err
// 	}
// 	defer src.Close()
// 	dst, err := os.Create(file.Filename)
// 	if err != nil {
// 		return err
// 	}
// 	defer dst.Close()
//
// 	if err := img.Resize(src, dst, image.Rect(0, 0, 480, 480)); err != nil {
// 		return err
// 	}
//
// 	if err := h.BlobService.UploadContext(c.Request().Context(), src, "uploads/images/"+file.Filename, ""); err != nil {
// 		return err
// 	}
//
// 	return c.JSON(http.StatusOK, make(map[string]interface{}))
// }
