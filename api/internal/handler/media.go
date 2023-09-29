package handler

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/labstack/echo/v4"
)

// A few resources that helped along the way:
// https://medium.com/@alexsante/serving-up-videos-from-s3-to-the-browser-using-go-974dfc11b738
// https://docs.min.io/docs/how-to-use-aws-sdk-for-go-with-minio-server.html
// https://stackoverflow.com/questions/24116147/how-to-download-file-in-browser-from-go-server

// cleanFilepath is a helper function to avoid path escape characters, "up one directory" path parts,
// and associated exploit attempts. Strongly based on practices in golang echo static file middleware:
// https://github.com/labstack/echo/blob/master/middleware/static.go
func (h ApiHandler) cleanFilepath(rawPath string) (string, error) {
	p, err := url.PathUnescape(rawPath)
	if err != nil {
		return "", err
	}
	return filepath.Clean("/" + p), nil
}

// GetMedia serves media, files, etc for a given project
func (h ApiHandler) GetMedia(c echo.Context) error {

	// Get Wildcard Path
	path, err := cleanFilepath(c.Request().RequestURI)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// Remove URL Route Prefix; Prefix with bucketPrefix
	// Example: If api hosted at /develop/<endpoint>... and images in bucket under prefix /instrumentation
	//          S3 Key = (1) Start with URL of request (2) remove /develop (3) prepend /instrumentation
	key := aws.String(bucketPrefix + strings.TrimPrefix(path, *routePrefix))

	output, err := s3c.GetObject(&s3.GetObjectInput{Bucket: bucket, Key: key})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// S3 Output Body to Buffer
	buff, buffErr := io.ReadAll(output.Body)
	if buffErr != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Buffered Reader
	reader := bytes.NewReader(buff)

	c.Response().Header().Set(echo.HeaderContentDisposition, "attachment")
	// Set Cache Control to 1 Year
	c.Response().Header().Set("Cache-Control", "public, max-age=31536000")
	return c.Stream(http.StatusOK, "image/jpg", reader)
}
