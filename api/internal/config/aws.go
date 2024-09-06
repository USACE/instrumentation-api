package config

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type AWSS3Config struct {
	AWSS3Endpoint       string `env:"AWS_S3_ENDPOINT"`
	AWSS3Region         string `env:"AWS_S3_REGION"`
	AWSS3DisableSSL     bool   `env:"AWS_S3_DISABLE_SSL"`
	AWSS3ForcePathStyle bool   `env:"AWS_S3_FORCE_PATH_STYLE"`
	AWSS3Bucket         string `env:"AWS_S3_BUCKET"`
}

func (cfg *AWSS3Config) S3Config() (aws.Config, []func(*s3.Options)) {
	optFns := make([]func(*awsConfig.LoadOptions) error, 0)
	if cfg.AWSS3Region != "" {
		optFns = append(optFns, awsConfig.WithRegion(cfg.AWSS3Region))
	}
	s3Config, err := awsConfig.LoadDefaultConfig(context.Background(), optFns...)
	if err != nil {
		log.Fatal(err.Error())
	}
	s3OptFns := make([]func(*s3.Options), 0)
	if cfg.AWSS3Endpoint != "" {
		s3OptFns = append(s3OptFns, func(o *s3.Options) {
			o.BaseEndpoint = aws.String(cfg.AWSS3Endpoint)
		})
	}
	if cfg.AWSS3ForcePathStyle {
		s3OptFns = append(s3OptFns, func(o *s3.Options) {
			o.UsePathStyle = true
		})
	}

	return s3Config, s3OptFns
}

type AWSSQSConfig struct {
	AWSSQSRegion      string `env:"AWS_SQS_REGION"`
	AWSSQSEndpoint    string `env:"AWS_SQS_ENDPOINT"`
	AWSSQSQueueURL    string `env:"AWS_SQS_QUEUE_URL"`
	AWSSQSQueueName   string `env:"AWS_SQS_QUEUE_NAME"`
	AWSSQSQueueNoInit bool   `env:"AWS_SQS_QUEUE_NO_INIT"`
}

func (cfg *AWSSQSConfig) SQSConfig() (aws.Config, []func(*sqs.Options)) {
	awsOptFns := make([]func(*awsConfig.LoadOptions) error, 0)
	if cfg.AWSSQSRegion != "" {
		awsOptFns = append(awsOptFns, awsConfig.WithRegion(cfg.AWSSQSRegion))
	}
	sqsConfig, err := awsConfig.LoadDefaultConfig(context.Background(), awsOptFns...)
	if err != nil {
		log.Fatal(err.Error())
	}
	sqsOptFns := make([]func(*sqs.Options), 0)
	if cfg.AWSSQSEndpoint != "" {
		sqsOptFns = append(sqsOptFns, func(o *sqs.Options) {
			o.BaseEndpoint = aws.String(cfg.AWSSQSEndpoint)
		})
	}
	return sqsConfig, sqsOptFns
}
