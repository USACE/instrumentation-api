package main

import (
	"encoding/json"
	"testing"

	"github.com/USACE/go-simple-asyncer/asyncer"
	"github.com/aws/aws-lambda-go/events"
)

// TODO: move to test directory after postman tests migrated
// TestDcsLoaderIntegration should be run with the instrumentation-dcs service: https://github.com/USACE/instrumentation-dcs
func TestDcsLoaderIntegration(t *testing.T) {
	// MSG_COUNT := 10

	a, err := asyncer.NewAsyncer(asyncer.Config{Engine: "AWSSQS", Target: "local/http://localhost:9324/queue/instrumentation-dcs-goes"})
	if err != nil {
		t.Fatal(err.Error())
	}

	r := events.S3EventRecord{
		S3: events.S3Entity{
			Bucket: events.S3Bucket{Name: "corpsmap-data-incoming"},
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

	if err := a.CallAsync(snsClone); err != nil {
		t.Fatal(err.Error())
	}
}
