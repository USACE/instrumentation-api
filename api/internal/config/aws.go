package config

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
)

type AWSS3Config struct {
	AWSS3Endpoint       string `envconfig:"AWS_S3_ENDPOINT"`
	AWSS3Region         string `envconfig:"AWS_S3_REGION"`
	AWSS3DisableSSL     bool   `envconfig:"AWS_S3_DISABLE_SSL"`
	AWSS3ForcePathStyle bool   `envconfig:"AWS_S3_FORCE_PATH_STYLE"`
	AWSS3Bucket         string `envconfig:"AWS_S3_BUCKET"`
}

func (cfg *AWSS3Config) S3Config() aws.Config {
	optFns := make([]func(*awsConfig.LoadOptions) error, 0)
	if cfg.AWSS3Region != "" {
		optFns = append(optFns, awsConfig.WithRegion(cfg.AWSS3Region))
	}
	s3Config, err := awsConfig.LoadDefaultConfig(context.Background(), optFns...)
	if err != nil {
		log.Fatal(err.Error())
	}
	return s3Config
}

type AWSSQSConfig struct {
	AWSSQSRegion      string `envconfig:"AWS_SQS_REGION"`
	AWSSQSEndpoint    string `envconfig:"AWS_SQS_ENDPOINT"`
	AWSSQSQueueURL    string `envconfig:"AWS_SQS_QUEUE_URL"`
	AWSSQSQueueName   string `envconfig:"AWS_SQS_QUEUE_NAME"`
	AWSSQSQueueNoInit bool   `envconfig:"AWS_SQS_QUEUE_NO_INIT"`
}

// AWSSQSConfig returns a ready-to-go config for session.New() for SQS Actions.
// Supports local testing using SQS stand-in elasticmq
func (cfg *AWSSQSConfig) SQSConfig() aws.Config {
	optFns := make([]func(*awsConfig.LoadOptions) error, 0)
	if cfg.AWSSQSRegion != "" {
		optFns = append(optFns, awsConfig.WithRegion(cfg.AWSSQSRegion))
	}
	s3Config, err := awsConfig.LoadDefaultConfig(context.Background(), optFns...)
	if err != nil {
		log.Fatal(err.Error())
	}
	return s3Config
}
