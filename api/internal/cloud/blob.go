package cloud

import (
	"context"
	"io"
	"net/url"
	"path/filepath"
	"strings"

	"github.com/USACE/instrumentation-api/api/internal/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type Blob interface {
	NewReader(ctx context.Context, rawPath, bucketName string) (io.ReadCloser, error)
}

type S3Blob struct {
	*s3.S3
	cfg s3BlobConfig
}

type s3BlobConfig struct {
	awsConfig    *config.AWSS3Config
	bucketName   string
	bucketPrefix string
	routePrefix  string
}

var _ Blob = (*S3Blob)(nil)

func NewS3Blob(cfg *config.AWSS3Config, bucketPrefix, routePrefix string) *S3Blob {
	awsCfg := cfg.S3Config()
	sess := session.Must(session.NewSession(awsCfg))
	return &S3Blob{s3.New(sess), s3BlobConfig{
		awsConfig:    cfg,
		bucketName:   cfg.AWSS3Bucket,
		bucketPrefix: bucketPrefix,
		routePrefix:  routePrefix,
	}}
}

// cleanFilepath is a helper function to avoid path escape characters, "up one directory" path parts, and associated exploit attempts.
// https://github.com/labstack/echo/blob/master/middleware/static.go
func cleanFilepath(rawPath string) (string, error) {
	p, err := url.PathUnescape(rawPath)
	if err != nil {
		return "", err
	}
	return filepath.Clean("/" + p), nil
}

// GetBlob serves media from blob storage, files, etc for a given project
func (s *S3Blob) NewReader(ctx context.Context, rawPath, bucketName string) (io.ReadCloser, error) {
	path, err := cleanFilepath(rawPath)
	if err != nil {
		return nil, err
	}

	// Remove URL Route Prefix; Prefix with bucketPrefix
	// Example: If api hosted at /develop/<endpoint>... and images in bucket under prefix /instrumentation
	//          S3 Key = (1) Start with URL of request (2) remove /develop (3) prepend /instrumentation
	key := aws.String(s.cfg.bucketPrefix + strings.TrimPrefix(path, s.cfg.routePrefix))

	if bucketName == "" {
		bucketName = s.cfg.bucketName
	}

	output, err := s.GetObjectWithContext(ctx, &s3.GetObjectInput{Bucket: aws.String(bucketName), Key: key})
	if err != nil {
		return nil, err
	}
	return output.Body, nil
}
