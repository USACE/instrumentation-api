package handler_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xeipuuv/gojsonschema"
)

const alertConfigSchedulerOptsSchema = `{
    "type": "object",
    "properties": {
        "start_date": { "type": "string" },
        "schedule_interval": { "type": "string" },
        "mute_consecutive_alerts": { "type": "boolean" },
        "remind_interval": { "type": ["string", "null"] },
        "warning_interval": { "type": ["string", "null"] },
        "last_checked": { "type": ["string", "null"], "format": "date-time" },
        "last_reminded": { "type": ["string", "null"], "format": "date-time" }
    },
    "additionalProperties": true
}`

var alertConfigSchedulerSchema = fmt.Sprintf(alertConfigSchemaTemplate, alertConfigTimeseriesSchema, alertConfigInstrumentSchema, alertConfigEmailSchema, alertConfigSchedulerOptsSchema)

var alertConfigSchedulerObjectLoader = gojsonschema.NewStringLoader(alertConfigSchedulerSchema)

const createAlertConfigSchedulerBody = `{
    "project_id": "5b6f4f37-7755-4cf9-bd02-94f1e9bc5984",
    "name": "New Test Alert Config",
    "body": "New Test Alert Config Description",
    "alert_type_id": "97e7a25c-d5c7-4ded-b272-1bb6e5914fe3",
    "opts": {
        "start_date": "2023-05-16T13:19:41.441328Z",
        "schedule_interval": "P1D",
        "mute_consecutive_alerts": true,
        "warning_interval": "PT1H"
    },
    "instruments": [
        {
            "instrument_id": "a7540f69-c41e-43b3-b655-6e44097edb7e"
        }
    ],
    "alert_email_subscriptions": ` + testCreateAlertConfigBodyEmailSubs + `
}`

const updateAlertConfigSchedulerBody = `{
    "name": "Updated Test Alert 1",
    "body": "Updated Alert for demonstration purposes.",
    "project_id": "5b6f4f37-7755-4cf9-bd02-94f1e9bc5984",
    "opts": {
        "start_date": "2023-05-16T13:19:41.441328Z",
        "schedule_interval": "P3D",
        "mute_consecutive_alerts": false,
        "remind_interval": "P1D"
    },
    "instruments": [],
    "alert_email_subscriptions": [
        {
            "id": "57329df6-9f7a-4dad-9383-4633b452efab",
            "user_type": "profile",
            "username": "AnthonyLambert",
            "email": "anthony.lambert@fake.usace.army.mil"
        }
    ]
}`

const (
	testAlertConfigSchedulerID = "1efd2d85-d3ee-4388-85a0-f824a761ff8b"
)

func TestAlertConfigSchedulers(t *testing.T) {
	objSchema, err := gojsonschema.NewSchema(alertConfigSchedulerObjectLoader)
	assert.Nil(t, err)

	tests := []HTTPTest{
		{
			Name:           "CreateAlertConfigScheduler",
			URL:            fmt.Sprintf("/projects/%s/alert_configs/schedulers", testProjectID),
			Method:         http.MethodPost,
			Body:           createAlertConfigSchedulerBody,
			ExpectedStatus: http.StatusCreated,
			ExpectedSchema: objSchema,
		},
		{
			Name:           "UpdateAlertConfigScheduler",
			URL:            fmt.Sprintf("/projects/%s/alert_configs/schedulers/%s", testProjectID, testAlertConfigSchedulerID),
			Method:         http.MethodPut,
			Body:           updateAlertConfigSchedulerBody,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: objSchema,
		},
		{
			Name:           "GetAlertConfig",
			URL:            fmt.Sprintf("/projects/%s/alert_configs/%s", testProjectID, testAlertConfigSchedulerID),
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: objSchema,
		},
		{
			Name:           "DeleteAlertConfig",
			URL:            fmt.Sprintf("/projects/%s/alert_configs/%s", testProjectID, testAlertConfigSchedulerID),
			Method:         http.MethodDelete,
			ExpectedStatus: http.StatusOK,
		}}

	RunAll(t, tests)
}
