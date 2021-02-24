package handlers

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
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
func cleanFilepath(rawPath string) (string, error) {
	p, err := url.PathUnescape(rawPath)
	if err != nil {
		return "", err
	}
	return filepath.Clean("/" + p), nil
}

// GetMedia serves media, files, etc for a given project
func GetMedia(s3Config *aws.Config, bucket *string) echo.HandlerFunc {
	return func(c echo.Context) error {

		// Get Wildcard Path
		keyPath, err := cleanFilepath(c.Request().RequestURI)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		key := aws.String(keyPath)

		newSession := session.New(s3Config)
		s3Client := s3.New(newSession)

		output, err := s3Client.GetObject(&s3.GetObjectInput{Bucket: bucket, Key: key})
		if err != nil {
			return c.String(500, err.Error())
		}

		// S3 Output Body to Buffer
		buff, buffErr := ioutil.ReadAll(output.Body)
		if buffErr != nil {
			return c.String(500, err.Error())
		}

		// Buffered Reader
		reader := bytes.NewReader(buff)

		c.Response().Header().Set(echo.HeaderContentDisposition, "attachment")
		// Set Cache Control to 1 Year
		c.Response().Header().Set("Cache-Control", "public, max-age=31536000")
		return c.Stream(http.StatusOK, "image/jpg", reader)
	}
}
