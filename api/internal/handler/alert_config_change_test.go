package handler_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xeipuuv/gojsonschema"
)

const alertConfigChangeOptsSchema = `{
    "type": "object",
    "properties": {
	"warn_rate_of_change": { "type": ["number", "null"] },
	"alert_rate_of_change": { "type": "number" },
	"ignore_rate_of_change": { "type": ["number", "null"] },
	"locf_backfill": { "type": ["string", "null"] }
    },
    "additionalProperties": true
}`

var alertConfigChangeSchema = fmt.Sprintf(alertConfigSchemaTemplate, alertConfigTimeseriesSchema, alertConfigInstrumentSchema, alertConfigEmailSchema, alertConfigChangeOptsSchema)

var alertConfigChangeObjectLoader = gojsonschema.NewStringLoader(alertConfigChangeSchema)

const createAlertConfigChangeBody = `{
    "project_id": "5b6f4f37-7755-4cf9-bd02-94f1e9bc5984",
    "name": "New Test Rate of Change Alert Config",
    "body": "New Test Rate of Change Alert Config Description",
    "alert_type_id": "c37effee-6b48-4436-8d72-737ed78c1fb7",
    "opts": {
        "warn_rate_of_change": 5.0,
        "alert_rate_of_change": 10.0,
        "ignore_rate_of_change": null,
        "locf_backfill": "1 hour"
    },
    "timeseries": [
        {
            "timeseries_id": "9a3864a8-8766-4bfa-bad1-0328b166f6a8"
        }
    ],
    "alert_email_subscriptions": ` + testCreateAlertConfigBodyEmailSubs + `
}`

const updateAlertConfigChangeBody = `{
    "name": "Updated Test Rate of Change Alert 1",
    "body": "Updated Test Rate of Change Alert for demonstration purposes.",
    "project_id": "5b6f4f37-7755-4cf9-bd02-94f1e9bc5984",
    "opts": {
        "warn_rate_of_change": 8.0,
        "alert_rate_of_change": 15.0,
        "ignore_rate_of_change": 1000.0,
        "locf_backfill": null
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

const testAlertConfigChangeID = "8dc1d10e-938e-4b27-8ba8-2fd5bb957ccb"

func TestAlertConfigChanges(t *testing.T) {
	objSchema, err := gojsonschema.NewSchema(alertConfigChangeObjectLoader)
	assert.Nil(t, err)

	tests := []HTTPTest{
		{
			Name:           "CreateAlertConfigChange",
			URL:            fmt.Sprintf("/projects/%s/alert_configs/changes", testProjectID),
			Method:         http.MethodPost,
			Body:           createAlertConfigChangeBody,
			ExpectedStatus: http.StatusCreated,
			ExpectedSchema: objSchema,
		},
		{
			Name:           "UpdateAlertConfigChange",
			URL:            fmt.Sprintf("/projects/%s/alert_configs/changes/%s", testProjectID, testAlertConfigChangeID),
			Method:         http.MethodPut,
			Body:           updateAlertConfigChangeBody,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: objSchema,
		},
		{
			Name:           "GetAlertConfig",
			URL:            fmt.Sprintf("/projects/%s/alert_configs/%s", testProjectID, testAlertConfigChangeID),
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: objSchema,
		},
		{
			Name:           "DeleteAlertConfig",
			URL:            fmt.Sprintf("/projects/%s/alert_configs/%s", testProjectID, testAlertConfigChangeID),
			Method:         http.MethodDelete,
			ExpectedStatus: http.StatusOK,
		}}

	RunAll(t, tests)
}
