package config

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/kelseyhightower/envconfig"
)

// Config stores configuration information stored in environment variables
type ApiConfig struct {
	DBConfig
	AuthDisabled        bool   `envconfig:"AUTH_DISABLED"`
	AuthJWTMocked       bool   `envconfig:"AUTH_JWT_MOCKED"`
	ApplicationKey      string `envconfig:"APPLICATION_KEY"`
	LambdaContext       bool
	HeartbeatKey        string
	RoutePrefix         string `envconfig:"ROUTE_PREFIX"`
	AWSS3Endpoint       string `envconfig:"AWS_S3_ENDPOINT"`
	AWSS3Region         string `envconfig:"AWS_S3_REGION"`
	AWSS3DisableSSL     bool   `envconfig:"AWS_S3_DISABLE_SSL"`
	AWSS3ForcePathStyle bool   `envconfig:"AWS_S3_FORCE_PATH_STYLE"`
	AWSS3Bucket         string `envconfig:"AWS_S3_BUCKET"`
}

func NewAWSConfig(cfg *ApiConfig) *aws.Config {
	awsConfig := aws.NewConfig().WithRegion(cfg.AWSS3Region)

	// Used for "minio" during development
	awsConfig.WithDisableSSL(cfg.AWSS3DisableSSL)
	awsConfig.WithS3ForcePathStyle(cfg.AWSS3ForcePathStyle)
	if cfg.AWSS3Endpoint != "" {
		awsConfig.WithEndpoint(cfg.AWSS3Endpoint)
	}
	return awsConfig
}

// GetConfig returns environment variable config
func NewApiConfig() *ApiConfig {
	var cfg ApiConfig
	if err := envconfig.Process("instrumentation", &cfg); err != nil {
		log.Fatal(err.Error())
	}
	return &cfg
}
