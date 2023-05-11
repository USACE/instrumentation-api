package main

import (
	"context"
	"fmt"
	"log"

	"github.com/USACE/instrumentation-api/api/dbutils"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/kelseyhightower/envconfig"
)

const (
	// Replace sender@example.com with your "From" address
	// This address must be verified with Amazon SES
	Sender  = "sender@example.com"
	CharSet = "UTF-8"

	// TODO: Get all below vars from database email alert config
	Recipient = "recipient@example.com"
	Subject   = "Amazon SES Test (AWS SDK for Go)"
	HtmlBody  = "<h1>Amazon SES Test Email (AWS SDK for Go)</h1><p>This email was sent with " +
		"<a href='https://aws.amazon.com/ses/'>Amazon SES</a> using the " +
		"<a href='https://aws.amazon.com/sdk-for-go/'>AWS SDK for Go</a>.</p>"

	//The email body for recipients with non-HTML email clients.
	TextBody = "This email was sent with Amazon SES using the AWS SDK for Go."
)

// Config stores configuration information stored in environment variables
type Config struct {
	DBUser           string
	DBPass           string
	DBName           string
	DBHost           string
	DBSSLMode        string
	AWSDefaultRegion string `envconfig:"AWS_DEFAULT_REGION"`
}

type MyEvent struct {
	Name string `json:"name"`
}

func (c *Config) dbConnStr() string {
	return fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=%s", c.DBUser, c.DBPass, c.DBName, c.DBHost, c.DBSSLMode)
}

func HandleRequest(ctx context.Context, name MyEvent) error {
	var cfg Config
	if err := envconfig.Process("instrumentation", &cfg); err != nil {
		log.Fatal(err.Error())
	}

	db := dbutils.Connection(cfg.dbConnStr())

	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(cfg.AWSDefaultRegion)},
	))
	svc := ses.New(sess)

	// Update last_checked on read
	rows, err := db.Queryx(`
		UPDATE alert_config ac1
		SET last_checked = now()
		FROM  (
			SELECT *
			FROM alert_config a
			WHERE last_checked < now() - schedule 
		) ac2
		WHERE  ac1.id = ac2.id
		RETURNING ac2.*
	`)
	if err != nil {
		return err
	}

	for rows.Next() {
		// TODO:
		// check all of the alerts where the `last_checked` timestamp is greater than
		// the current time subtracted by its respective interval (15mins | 1hour | 1day | 1week, | 1month | quarterly)
		//
		// give each alert config a "type" of alert. depending on the type of alert, the alert config will be routed to a function
		// that handles the respective alert type, queries the data needed, and checks the corresponding condition
		//
		// for each alert that needs to be checked, run defined condition for alert and send email if alert condition is met
	}

	// Created array of "aws strings" from recipients defined in alert config
	toAddresses := []*string{aws.String(Recipient)}
	constructAndSendEmail(svc, toAddresses, Subject, HtmlBody, TextBody)

	return nil
}

// handleMeasurementUploadAlert checks that measurements for an instrument exist within
// a defined interval subtracted by the current time and sends an alert otherwise
//
// Each alert config has a corresponding instrument_id and schedule (interval)
// Query if any timeseries measurements for an instrument exist that are older than the current time minus the interval,
// and check that for each alert_config/instrument_id.
func handleMeasurementUploadAlert()

// handleQcDataEvaluationAlert checks that a "QC data evaluation" has been submitted for the instrument
// within the defined time interval
//
// Each alert config has a corresponding instrument_id and schedule (interval)
// Query `qc_instrument_evaluated`
func handleQcDataEvaluationAlert()

func constructAndSendEmail(svc *ses.SES, toAddresses []*string, subject, htmlBody, textBody string) {
	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{},
			ToAddresses: toAddresses,
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(HtmlBody),
				},
				Text: &ses.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(TextBody),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String(CharSet),
				Data:    aws.String(Subject),
			},
		},
		Source: aws.String(Sender),
	}

	result, err := svc.SendEmail(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ses.ErrCodeMessageRejected:
				log.Println(ses.ErrCodeMessageRejected, aerr.Error())
			case ses.ErrCodeMailFromDomainNotVerifiedException:
				log.Println(ses.ErrCodeMailFromDomainNotVerifiedException, aerr.Error())
			case ses.ErrCodeConfigurationSetDoesNotExistException:
				log.Println(ses.ErrCodeConfigurationSetDoesNotExistException, aerr.Error())
			default:
				log.Println(aerr.Error())
			}
		} else {
			log.Println(err.Error())
		}
	}
	log.Println(result)
}

func main() {
	lambda.Start(HandleRequest)
}
