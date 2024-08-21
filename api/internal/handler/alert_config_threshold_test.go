package handler_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xeipuuv/gojsonschema"
)

const alertConfigThresholdOptsSchema = `{
    "type": "object",
    "properties": {
	"alert_low_value": { "type": ["number", "null"] },
        "alert_high_value": { "type": ["number", "null"] },
        "warn_low_value": { "type": ["number", "null"] },
        "warn_high_value": { "type": ["number", "null"] },
        "ignore_low_value": { "type": ["number", "null"] },
        "ignore_high_value": { "type": ["number", "null"] },
	"variance": { "type": "number" }
    },
    "additionalProperties": true
}`

var alertConfigThresholdSchema = fmt.Sprintf(alertConfigSchemaTemplate, alertConfigTimeseriesSchema, alertConfigInstrumentSchema, alertConfigEmailSchema, alertConfigThresholdOptsSchema)

var alertConfigThresholdObjectLoader = gojsonschema.NewStringLoader(alertConfigThresholdSchema)

const createAlertConfigThresholdBody = `{
    "project_id": "5b6f4f37-7755-4cf9-bd02-94f1e9bc5984",
    "name": "New Test Threshold Alert Config",
    "body": "New Test Threshold Alert Config Description",
    "alert_type_id": "bb15e7c2-8eae-452c-92f7-e720dc5c9432",
    "opts": {
        "alert_low_value": 10.0,
        "alert_high_value": 50.0,
        "warn_low_value": 15.0,
        "warn_high_value": 45.0,
        "ignore_low_value": 5.0,
        "ignore_high_value": 55.0,
        "variance": 1.0
    },
    "timeseries": [
        {
            "timeseries_id": "869465fc-dc1e-445e-81f4-9979b5fadda9"
        }
    ],
    "alert_email_subscriptions": ` + testCreateAlertConfigBodyEmailSubs + `
}`

const updateAlertConfigThresholdBody = `{
    "name": "Updated Test Threshold Alert 1",
    "body": "Updated Test Threshold Alert for demonstration purposes.",
    "project_id": "5b6f4f37-7755-4cf9-bd02-94f1e9bc5984",
    "opts": {
        "alert_low_value": 20.0,
        "alert_high_value": 60.0,
        "warn_low_value": 25.0,
        "warn_high_value": 55.0,
        "ignore_low_value": 10.0,
        "ignore_high_value": 65.0,
        "variance": 2.0
    },
    "timeseries": [],
    "alert_email_subscriptions": [
        {
            "id": "57329df6-9f7a-4dad-9383-4633b452efab",
            "user_type": "profile",
            "username": "AnthonyLambert",
            "email": "anthony.lambert@fake.usace.army.mil"
        }
    ]
}`

const testAlertConfigThresholdID = "7fae1097-4e96-453f-bf36-7a3427cfb0d7"

func TestAlertConfigThresholds(t *testing.T) {
	objSchema, err := gojsonschema.NewSchema(alertConfigThresholdObjectLoader)
	assert.Nil(t, err)

	tests := []HTTPTest{
		{
			Name:           "CreateAlertConfigThreshold",
			URL:            fmt.Sprintf("/projects/%s/alert_configs/thresholds", testProjectID),
			Method:         http.MethodPost,
			Body:           createAlertConfigThresholdBody,
			ExpectedStatus: http.StatusCreated,
			ExpectedSchema: objSchema,
		},
		{
			Name:           "UpdateAlertConfigThreshold",
			URL:            fmt.Sprintf("/projects/%s/alert_configs/thresholds/%s", testProjectID, testAlertConfigThresholdID),
			Method:         http.MethodPut,
			Body:           updateAlertConfigThresholdBody,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: objSchema,
		},
		{
			Name:           "GetAlertConfig",
			URL:            fmt.Sprintf("/projects/%s/alert_configs/%s", testProjectID, testAlertConfigThresholdID),
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: objSchema,
		},
		{
			Name:           "DeleteAlertConfig",
			URL:            fmt.Sprintf("/projects/%s/alert_configs/%s", testProjectID, testAlertConfigThresholdID),
			Method:         http.MethodDelete,
			ExpectedStatus: http.StatusOK,
		}}

	RunAll(t, tests)
}
