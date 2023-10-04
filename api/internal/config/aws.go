package config

import (
	"github.com/aws/aws-sdk-go/aws"
)

type AWSS3Config struct {
	AWSS3Endpoint       string `envconfig:"AWS_S3_ENDPOINT"`
	AWSS3Region         string `envconfig:"AWS_S3_REGION"`
	AWSS3DisableSSL     bool   `envconfig:"AWS_S3_DISABLE_SSL"`
	AWSS3ForcePathStyle bool   `envconfig:"AWS_S3_FORCE_PATH_STYLE"`
	AWSS3Bucket         string `envconfig:"AWS_S3_BUCKET"`
}

func (cfg *AWSS3Config) S3Config() *aws.Config {
	awsConfig := aws.NewConfig()
	awsConfig.WithDisableSSL(cfg.AWSS3DisableSSL)
	awsConfig.WithS3ForcePathStyle(cfg.AWSS3ForcePathStyle)
	if cfg.AWSS3Region != "" {
		awsConfig.WithRegion(cfg.AWSS3Region)
	}
	if cfg.AWSS3Endpoint != "" {
		awsConfig.WithEndpoint(cfg.AWSS3Endpoint)
	}
	return awsConfig
}

type AWSSQSConfig struct {
	AWSSQSRegion    string `envconfig:"AWS_SQS_REGION"`
	AWSSQSEndpoint  string `envconfig:"AWS_SQS_ENDPOINT"`
	AWSSQSQueueURL  string `envconfig:"AWS_SQS_QUEUE_URL"`
	AWSSQSQueueName string `envconfig:"AWS_SQS_QUEUE_NAME"`
}

// AWSSQSConfig returns a ready-to-go config for session.New() for SQS Actions.
// Supports local testing using SQS stand-in elasticmq
func (cfg *AWSSQSConfig) SQSConfig() *aws.Config {
	awsConfig := aws.NewConfig()
	if cfg.AWSSQSRegion != "" {
		awsConfig.WithRegion(cfg.AWSSQSRegion)
	}
	if cfg.AWSSQSEndpoint != "" {
		awsConfig.WithEndpoint(cfg.AWSSQSEndpoint)
	}
	if cfg.AWSSQSRegion != "" {
		awsConfig.WithRegion(cfg.AWSSQSRegion)
	}
	return awsConfig
}
