package handler_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/USACE/instrumentation-api/api/internal/cloud"
	"github.com/USACE/instrumentation-api/api/internal/config"
	"github.com/aws/aws-lambda-go/events"
)

func TestDcsLoaderIntegration(t *testing.T) {
	cfg := config.NewDcsLoaderConfig()
	ps := cloud.NewSQSPubsub(&cfg.AWSSQSConfig)

	r := events.S3EventRecord{
		S3: events.S3Entity{
			Bucket: events.S3Bucket{Name: cfg.AWSS3Bucket},
			Object: events.S3Object{Key: "test/test-file.csv"},
		},
	}

	eventStr, err := json.Marshal(events.S3Event{Records: []events.S3EventRecord{r}})
	if err != nil {
		t.Fatal(err.Error())
	}

	snsClone, err := json.Marshal(events.SNSEntity{Message: string(eventStr)})
	if err != nil {
		t.Fatal(err.Error())
	}

	messageID, err := ps.PublishMessage(context.Background(), snsClone)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Logf("successfully processed message: %s", messageID)
}
