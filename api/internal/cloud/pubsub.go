package cloud

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"

	"github.com/USACE/instrumentation-api/api/internal/config"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type Pubsub interface {
	ProcessMessages(ctx context.Context, handler messageHandler) error
}

type SQSPubsub struct {
	*sqs.SQS
	cfg  *config.AWSSQSConfig
	blob Blob
}

var _ Pubsub = (*SQSPubsub)(nil)

func NewSQSPubsub(cfg *config.AWSSQSConfig) *SQSPubsub {
	awsCfg := cfg.SQSConfig()
	sess := session.Must(session.NewSession(awsCfg))
	return &SQSPubsub{sqs.New(sess), cfg, nil}
}

func (s *SQSPubsub) WithBlob(blob Blob) *SQSPubsub {
	s.blob = blob
	return s
}

func queueURL(ctx context.Context, s *SQSPubsub) (string, error) {
	if s.cfg.AWSSQSQueueURL != "" {
		return s.cfg.AWSSQSQueueURL, nil
	}
	urlResult, err := s.GetQueueUrlWithContext(ctx, &sqs.GetQueueUrlInput{QueueName: &s.cfg.AWSSQSQueueName})
	if err != nil {
		return "", err
	}
	if urlResult == nil || urlResult.QueueUrl == nil {
		return "", fmt.Errorf("queue url is nil")
	}
	return *urlResult.QueueUrl, nil
}

type messageHandler func(ctx context.Context, r io.Reader) error

func (s *SQSPubsub) ProcessMessages(ctx context.Context, handler messageHandler) error {
	if s.blob == nil {
		return fmt.Errorf("blob must not be nil")
	}

	url, err := queueURL(ctx, s)
	if err != nil {
		return err
	}

	var entity events.SNSEntity
	var evt events.S3Event

	for {
		output, err := s.ReceiveMessageWithContext(ctx, &sqs.ReceiveMessageInput{
			AttributeNames: []*string{
				aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
			},
			MessageAttributeNames: []*string{
				aws.String(sqs.QueueAttributeNameAll),
			},
			QueueUrl:            aws.String(url),
			MaxNumberOfMessages: aws.Int64(1),
			VisibilityTimeout:   aws.Int64(30),
			WaitTimeSeconds:     aws.Int64(20),
		})
		if err != nil {
			return err
		}

		fmt.Printf("Received %d messages\n", len(output.Messages))
		for _, m := range output.Messages {
			fmt.Printf("Working on Message: %s\n", *m.MessageId)

			if err := json.Unmarshal([]byte(*m.Body), &entity); err != nil {
				fmt.Printf("Error: %s\n", err.Error())
				continue
			}
			if err := json.Unmarshal([]byte(entity.Message), &evt); err != nil {
				fmt.Printf("Error: %s\n", err.Error())
				continue
			}
			for _, record := range evt.Records {
				bucket, key := record.S3.Bucket.Name, record.S3.Object.Key
				log.Printf("Processing File; bucket: %s; key: %s\n", bucket, key)

				r, err := s.blob.NewReader(ctx, key, bucket)
				if err != nil {
					return err
				}
				defer func(rc io.ReadCloser) {
					if err := rc.Close(); err != nil {
						fmt.Printf("could not close file: %s", err.Error())
					}
				}(r)
				if err := handler(ctx, r); err != nil {
					log.Printf("message processing failed")
					continue
				}
			}

			s.DeleteMessage(&sqs.DeleteMessageInput{
				QueueUrl:      aws.String(url),
				ReceiptHandle: m.MessageId,
			})
		}
	}
}
