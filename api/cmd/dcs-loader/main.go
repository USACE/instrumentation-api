package main

import (
	"bytes"
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/config"
	ts "github.com/USACE/instrumentation-api/api/internal/timeseries"

	"github.com/google/uuid"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type handlerFunc func(context.Context, events.S3Event) error

// HandleRequest parses a CSV file hosted on S3 and stores contents instrumentation-api
func HandleRequest(cfg *config.DcsLoaderConfig) handlerFunc {

	return func(ctx context.Context, s3Event events.S3Event) error {

		sess := session.Must(session.NewSession(cfg.AWSS3Config()))
		s3Client := s3.New(sess)
		httpClient := &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return nil
			},
		}

		for _, record := range s3Event.Records {

			bucket, key := &record.S3.Bucket.Name, &record.S3.Object.Key
			log.Printf("Processing File; bucket: %s; key: %s\n", *bucket, *key)

			output, err := s3Client.GetObject(&s3.GetObjectInput{Bucket: bucket, Key: key})
			if err != nil {
				return err
			}

			defer output.Body.Close()

			mcs, mCount, err := ParseCsvMeasurementCollection(output.Body, cfg)
			if err != nil {
				return err
			}

			startPostTime := time.Now()
			if err := PostMeasurementCollectionToApi(mcs, cfg, httpClient); err != nil {
				return err
			}
			log.Printf(
				"\n\tSUCCESS; POST %d measurements across %d timeseries in %f seconds\n",
				mCount, len(mcs), time.Since(startPostTime).Seconds(),
			)
		}
		return nil
	}
}

func ParseCsvMeasurementCollection(r io.Reader, cfg *config.DcsLoaderConfig) ([]ts.MeasurementCollection, int, error) {

	mcs := make([]ts.MeasurementCollection, 0)
	mCount := 0
	reader := csv.NewReader(r)

	rows := make([][]string, 0)
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return mcs, mCount, err
		}
		rows = append(rows, row)
	}

	mcMap := make(map[uuid.UUID]*ts.MeasurementCollection)
	for _, row := range rows {
		// 0=timeseries_id, 1=time, 2=value
		tsid, err := uuid.Parse(row[0])
		if err != nil {
			return mcs, mCount, err
		}
		t, err := time.Parse(time.RFC3339, row[1])
		if err != nil {
			return mcs, mCount, err
		}
		v, err := strconv.ParseFloat(row[2], 32)
		if err != nil {
			return mcs, mCount, err
		}

		if _, ok := mcMap[tsid]; !ok {
			mcMap[tsid] = &ts.MeasurementCollection{
				TimeseriesID: tsid,
				Items:        make([]ts.Measurement, 0),
			}
		}
		mcMap[tsid].Items = append(mcMap[tsid].Items, ts.Measurement{TimeseriesID: tsid, Time: t, Value: v})
		mCount++
	}

	mcs = make([]ts.MeasurementCollection, len(mcMap))
	idx := 0
	for _, v := range mcMap {
		mcs[idx] = *v
		idx++
	}

	return mcs, mCount, nil
}

func PostMeasurementCollectionToApi(mcs []ts.MeasurementCollection, cfg *config.DcsLoaderConfig, client *http.Client) error {

	requestBodyBytes, err := json.Marshal(mcs)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s?key=%s", cfg.PostURL, cfg.APIKey), bytes.NewReader(requestBodyBytes))
	if err != nil {
		return err
	}
	defer req.Body.Close()

	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("\n\t*** Error; %s\n", err.Error())
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 201 {
		log.Printf("\n\t*** Error; Status Code: %d ***\n", resp.StatusCode)
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("Error reading response body")
			return err
		}
		log.Printf("%s\n", body)
	}
	return nil
}

func main() {

	cfg := config.GetDcsLoaderConfig()
	handler := HandleRequest(cfg)

	sessSQS := session.Must(session.NewSession(cfg.AWSSQSConfig()))
	svcSQS := sqs.New(sessSQS)

	queueURL, err := cfg.AWSSQSQueueURL(svcSQS)
	if err != nil {
		log.Fatal(err.Error())
	}
	if queueURL == "" {
		log.Fatal("Could not find queue url")
	}

	// Single memory locations to be reused by all for loop iterations
	var SNSEvt events.SNSEntity
	pSNSEvt := &SNSEvt
	var S3Evt events.S3Event
	pS3Evt := &S3Evt

	for {
		log.Println("Calling Receive Messages...")
		output, err := svcSQS.ReceiveMessage(&sqs.ReceiveMessageInput{
			AttributeNames: []*string{
				aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
			},
			MessageAttributeNames: []*string{
				aws.String(sqs.QueueAttributeNameAll),
			},
			QueueUrl:            &queueURL,
			MaxNumberOfMessages: aws.Int64(1),
			VisibilityTimeout:   aws.Int64(30),
			WaitTimeSeconds:     aws.Int64(20),
		})
		if err != nil {
			log.Printf("Error: %s\n", err.Error())
			continue
		}

		log.Printf("Received %d messages\n", len(output.Messages))
		for _, m := range output.Messages {
			log.Printf("Working on Message: %s\n", *m.MessageId)

			if err := json.Unmarshal([]byte(*m.Body), pSNSEvt); err != nil {
				log.Printf("Error: %s\n", err.Error())
				continue
			}

			if err := json.Unmarshal([]byte(pSNSEvt.Message), pS3Evt); err != nil {
				log.Printf("Error: %s\n", err.Error())
				continue
			}

			if err := handler(context.Background(), *pS3Evt); err != nil {
				log.Printf("Error: %s\n", err.Error())
				continue
			}

			svcSQS.DeleteMessage(&sqs.DeleteMessageInput{
				QueueUrl:      &queueURL,
				ReceiptHandle: m.MessageId,
			})
		}
	}
}
