package cloud

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/USACE/instrumentation-api/api/internal/config"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	sqsTypes "github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

type Pubsub interface {
	ProcessMessagesFromBlob(handler messageHandler) error
	PublishMessage(ctx context.Context, message []byte) (string, error)
	MockPublishMessage(ctx context.Context, message []byte) (string, error)
}

type messageHandler func(r io.Reader) error

type SQSPubsub struct {
	*sqs.Client
	cfg      *config.AWSSQSConfig
	blob     Blob
	queueUrl *string
}

var _ Pubsub = (*SQSPubsub)(nil)

func NewSQSPubsub(cfg *config.AWSSQSConfig) *SQSPubsub {
	sqsCfg, optFns := cfg.SQSConfig()
	queue := sqs.NewFromConfig(sqsCfg, optFns...)

	ps := &SQSPubsub{queue, cfg, nil, nil}
	if !cfg.AWSSQSQueueNoInit {
		ps.MustInitQueueUrl()
	}

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
		return nil
	}
	urlResult, err := s.GetQueueUrl(context.Background(), &sqs.GetQueueUrlInput{QueueName: &s.cfg.AWSSQSQueueName})
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
	ctx := context.Background()

	for {
		output, err := s.ReceiveMessage(ctx, &sqs.ReceiveMessageInput{
			AttributeNames: []sqsTypes.QueueAttributeName{
				sqsTypes.QueueAttributeName(sqsTypes.MessageSystemAttributeNameSentTimestamp),
			},
			MessageAttributeNames: []string{
				string(sqsTypes.QueueAttributeNameAll),
			},
			QueueUrl:            s.queueUrl,
			MaxNumberOfMessages: 1,
			VisibilityTimeout:   30,
			WaitTimeSeconds:     20,
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

			s.DeleteMessage(ctx, &sqs.DeleteMessageInput{
				QueueUrl:      s.queueUrl,
				ReceiptHandle: m.MessageId,
			})
		}
	}
}

func (s *SQSPubsub) PublishMessage(ctx context.Context, message []byte) (string, error) {
	if s.queueUrl == nil {
		if err := s.InitQueueUrl(); err != nil {
			return "", err
		}
	}

	messageBody := string(message)
	sqsMessageInput := &sqs.SendMessageInput{
		MessageBody: &messageBody,
		QueueUrl:    s.queueUrl,
	}

	out, err := s.SendMessage(ctx, sqsMessageInput)
	if err != nil {
		return "", err
	}

	if out == nil || out.MessageId == nil {
		return "", errors.New("nil message id returned from queue")
	}

	return *out.MessageId, nil
}

func (s *SQSPubsub) MockPublishMessage(ctx context.Context, message []byte) (string, error) {
	body := events.SQSMessage{Body: string(message)}
	event := events.SQSEvent{Records: []events.SQSMessage{body}}
	b, err := json.Marshal(event)
	if err != nil {
		return "", err
	}

	if err := mockInvokeLocalLambda("http://report:8080/2015-03-31/functions/function/invocations", b); err != nil {
		return "", err
	}

	return "mock-message-id", nil
}

func mockInvokeLocalLambda(invokeUrl string, message []byte) error {
	r := bytes.NewReader(message)
	res, err := http.Post(invokeUrl, "application/json", r)
	if err != nil {
		return err
	}
	if res.StatusCode != 200 {
		return errors.New("unable to invoke locally mocked lambda")
	}
	return nil
}
