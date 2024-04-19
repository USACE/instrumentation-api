package cloud

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"

	"github.com/USACE/instrumentation-api/api/internal/config"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type Pubsub interface {
	ProcessMessagesFromBlob(handler messageHandler) error
	PublishMessage(ctx context.Context, message json.Marshaler) (string, error)
}

type messageHandler func(r io.Reader) error

type SQSPubsub struct {
	*sqs.SQS
	cfg      *config.AWSSQSConfig
	blob     Blob
	queueUrl *string
}

var _ Pubsub = (*SQSPubsub)(nil)

func NewSQSPubsub(cfg *config.AWSSQSConfig) *SQSPubsub {
	awsCfg := cfg.SQSConfig()
	sess := session.Must(session.NewSession(awsCfg))
	queue := sqs.New(sess)

	ps := &SQSPubsub{queue, cfg, nil, nil}
	ps.MustInitQueueUrl()

	return ps
}

func (s *SQSPubsub) WithBlob(blob Blob) *SQSPubsub {
	s.blob = blob
	return s
}

func (s *SQSPubsub) MustInitQueueUrl() {
	if err := s.InitQueueUrl(); err != nil {
		log.Fatalf(err.Error())
	}
}

func (s *SQSPubsub) InitQueueUrl() error {
	if s.cfg.AWSSQSQueueURL != "" {
		s.queueUrl = &s.cfg.AWSSQSQueueURL
	}
	urlResult, err := s.GetQueueUrl(&sqs.GetQueueUrlInput{QueueName: &s.cfg.AWSSQSQueueName})
	if err != nil {
		return err
	}
	if urlResult == nil || urlResult.QueueUrl == nil {
		return errors.New("queue url is nil")
	}
	s.queueUrl = urlResult.QueueUrl
	return nil
}

func (s *SQSPubsub) ProcessMessagesFromBlob(handler messageHandler) error {
	if s.blob == nil {
		return errors.New("blob must not be nil")
	}

	if s.queueUrl == nil {
		if err := s.InitQueueUrl(); err != nil {
			return err
		}
	}

	var entity events.SNSEntity
	var evt events.S3Event

	for {
		output, err := s.ReceiveMessage(&sqs.ReceiveMessageInput{
			AttributeNames: []*string{
				aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
			},
			MessageAttributeNames: []*string{
				aws.String(sqs.QueueAttributeNameAll),
			},
			QueueUrl:            s.queueUrl,
			MaxNumberOfMessages: aws.Int64(1),
			VisibilityTimeout:   aws.Int64(30),
			WaitTimeSeconds:     aws.Int64(20),
		})
		if err != nil {
			return err
		}

		log.Printf("Received %d messages\n", len(output.Messages))
		for _, m := range output.Messages {
			log.Printf("Working on Message: %s\n", *m.MessageId)

			if err := json.Unmarshal([]byte(*m.Body), &entity); err != nil {
				log.Printf("Error: %s\n", err.Error())
				continue
			}
			if err := json.Unmarshal([]byte(entity.Message), &evt); err != nil {
				log.Printf("Error: %s\n", err.Error())
				continue
			}
			for _, record := range evt.Records {
				bucket, key := record.S3.Bucket.Name, record.S3.Object.Key
				log.Printf("Processing File; bucket: %s; key: %s\n", bucket, key)

				r, err := s.blob.NewReader(key, bucket)
				if err != nil {
					return err
				}
				defer func(rc io.ReadCloser) {
					if err := rc.Close(); err != nil {
						log.Printf("could not close file: %s", err.Error())
					}
				}(r)
				if err := handler(r); err != nil {
					log.Printf("message processing failed")
					continue
				}
			}

			s.DeleteMessage(&sqs.DeleteMessageInput{
				QueueUrl:      s.queueUrl,
				ReceiptHandle: m.MessageId,
			})
		}
	}
}

func (s *SQSPubsub) PublishMessage(ctx context.Context, message json.Marshaler) (string, error) {
	if s.queueUrl == nil {
		if err := s.InitQueueUrl(); err != nil {
			return "", err
		}
	}

	b, err := message.MarshalJSON()
	if err != nil {
		return "", err
	}
	messageBody := string(b)
	sqsMessageInput := &sqs.SendMessageInput{
		MessageBody: &messageBody,
	}

	out, err := s.SendMessageWithContext(ctx, sqsMessageInput)
	if err != nil {
		return "", err
	}

	if out.MessageId == nil {
		return "", errors.New("nil message id returned from queue")
	}

	return *out.MessageId, nil
}
