package cloud

import (
	"context"
	"io"
	"net/url"
	"path/filepath"
	"strings"

	"github.com/USACE/instrumentation-api/api/internal/config"
	"github.com/aws/aws-sdk-go-v2/aws"
	s3manager "github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Blob interface {
	NewReader(rawPath, bucketName string) (io.ReadCloser, error)
	NewReaderContext(ctx context.Context, rawPath, bucketName string) (io.ReadCloser, error)
	UploadContext(ctx context.Context, r io.Reader, rawPath, bucketName string) error
}

type S3Blob struct {
	*s3.Client
	*s3manager.Uploader
	cfg s3BlobConfig
}

type s3BlobConfig struct {
	awsConfig    *config.AWSS3Config
	bucketName   string
	bucketPrefix string
	routePrefix  string
}

var _ Blob = (*S3Blob)(nil)

func uploader(u *s3manager.Uploader) {}

func NewS3Blob(cfg *config.AWSS3Config, bucketPrefix, routePrefix string) *S3Blob {
	s3Config, optFns := cfg.S3Config()
	client := s3.NewFromConfig(s3Config, optFns...)

	return &S3Blob{client, s3manager.NewUploader(client, uploader), s3BlobConfig{
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

func (s *S3Blob) NewReader(rawPath, bucketName string) (io.ReadCloser, error) {
	path, err := cleanFilepath(rawPath)
	if err != nil {
		return nil, err
	}
	key := strings.TrimPrefix(s.cfg.bucketPrefix+strings.TrimPrefix(path, s.cfg.routePrefix), "/")
	if bucketName == "" {
		bucketName = s.cfg.bucketName
	}
	output, err := s.GetObject(context.Background(), &s3.GetObjectInput{Bucket: aws.String(bucketName), Key: aws.String(key)})
	if err != nil {
		return nil, err
	}
	return output.Body, nil
}

func (s *S3Blob) NewReaderContext(ctx context.Context, rawPath, bucketName string) (io.ReadCloser, error) {
	path, err := cleanFilepath(rawPath)
	if err != nil {
		return nil, err
	}
	key := aws.String(s.cfg.bucketPrefix + strings.TrimPrefix(path, s.cfg.routePrefix))
	if bucketName == "" {
		bucketName = s.cfg.bucketName
	}

	output, err := s.GetObject(ctx, &s3.GetObjectInput{Bucket: aws.String(bucketName), Key: key})
	if err != nil {
		return nil, err
	}
	return output.Body, nil
}

func (s *S3Blob) UploadContext(ctx context.Context, r io.Reader, rawPath, bucketName string) error {
	path, err := cleanFilepath(rawPath)
	if err != nil {
		return err
	}
	key := aws.String(s.cfg.bucketPrefix + strings.TrimPrefix(path, s.cfg.routePrefix))
	if bucketName == "" {
		bucketName = s.cfg.bucketName
	}

	_, err = s.Upload(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    key,
		Body:   r,
	})
	return err
}
