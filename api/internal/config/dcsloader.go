package config

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/kelseyhightower/envconfig"
)

// DcsLoaderConfig holds parameters parsed from env variables.
// Note awsSQSQueueURL private variable. Public method is AWSSQSQueueURL()
type DcsLoaderConfig struct {
	PostURL             string `envconfig:"POST_URL"`
	APIKey              string `envconfig:"API_KEY"`
	AWSS3Region         string `envconfig:"AWS_S3_REGION"`
	AWSS3Endpoint       string `envconfig:"AWS_S3_ENDPOINT"`
	AWSS3DisableSSL     bool   `envconfig:"AWS_S3_DISABLE_SSL"`
	AWSS3ForcePathStyle bool   `envconfig:"AWS_S3_FORCE_PATH_STYLE"`
	AWSSQSRegion        string `envconfig:"AWS_SQS_REGION"`
	AWSSQSEndpoint      string `envconfig:"AWS_SQS_ENDPOINT"`
	awsSQSQueueURL      string `envconfig:"AWS_SQS_QUEUE_URL"`
	AWSSQSQueueName     string `envconfig:"AWS_SQS_QUEUE_NAME"`
}

func NewDcsLoaderConfig() *DcsLoaderConfig {
	var cfg DcsLoaderConfig
	if err := envconfig.Process("loader", &cfg); err != nil {
		log.Fatal(err.Error())
	}
	return &cfg
}

// AWSS3Config returns a ready-to-go config to pass to session.New() for S3
// This function helps local testing against minio as an s3 stand-in
// where endpoint must be defined
func (cfg *DcsLoaderConfig) AWSS3Config() *aws.Config {
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

// AWSSQSConfig returns a ready-to-go config for session.New() for SQS Actions.
// Supports local testing using SQS stand-in elasticmq
func (cfg *DcsLoaderConfig) AWSSQSConfig() *aws.Config {
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

// AWSSQSQueueURL returns the QueueUrl for interacting with SQS
func (cfg *DcsLoaderConfig) AWSSQSQueueURL(s *sqs.SQS) (string, error) {
	// If environment variable AWS_SQS_QUEUE_URL is specified,
	// use provided queue URL without question
	if cfg.awsSQSQueueURL != "" {
		return cfg.awsSQSQueueURL, nil
	}
	// Lookup Queue URL from AWS_SQS_QUEUE_NAME
	urlResult, err := s.GetQueueUrl(&sqs.GetQueueUrlInput{QueueName: &cfg.AWSSQSQueueName})
	if err != nil {
		return "", err
	}
	return *urlResult.QueueUrl, nil
}
