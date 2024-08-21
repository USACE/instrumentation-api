package handler_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xeipuuv/gojsonschema"
)

const alertConfigInstrumentSchema = `{
    "type": "object",
    "properties": {
        "instrument_id": { "type": "string" },
        "instrument_name": { "type": "string" }
    }
}`

const alertConfigTimeseriesSchema = `{
    "type": "object",
    "properties": {
        "timeseries_id": { "type": "string" },
        "timeseries_name": { "type": "string" }
    }
}`

const alertConfigEmailSchema = `{
    "type": "object",
    "properties": {
        "id": { "type": "string" },
        "user_type": { "type": "string" },
        "username": { "type": ["string", "null"] },
        "email": { "type": "string" }
    }
}`

const alertConfigSchemaTemplate = `{
    "type": "object",
    "properties": {
        "id": { "type": "string" },
        "name": { "type": "string" },
        "body": { "type": "string" },
        "project_id": { "type": "string" },
        "alert_type_id": { "type": "string" },
        "alert_type": { "type": "string" },
        "timeseries": { "type": "array", "items": %s },
        "instruments": { "type": "array", "items": %s },
        "alert_email_subscriptions": { "type": "array", "items": %s },
        "alert_status": { "type": "string" },
        "creator_id": { "type": "string" },
        "creator_username": { "type": "string" },
        "create_date": { "type": "string", "format": "date-time" },
        "updater_id": { "type": ["string", "null"] },
        "updater_username": { "type": ["string", "null"] },
        "update_date": { "type": ["string", "null"], "format": "date-time" },
        "opts": %s
    },
    "additionalProperties": true
}`

var alertConfigSchema = fmt.Sprintf(alertConfigSchemaTemplate, alertConfigTimeseriesSchema, alertConfigInstrumentSchema, alertConfigEmailSchema, alertConfigSchedulerOptsSchema)

var alertConfigObjectLoader = gojsonschema.NewStringLoader(alertConfigSchema)

var alertConfigArrayLoader = gojsonschema.NewStringLoader(fmt.Sprintf(`{
    "type": "array",
    "items": %s
}`, alertConfigSchema))

const testCreateAlertConfigBodyEmailSubs = `[
        {   "id": "1ebf9e14-2b1c-404e-9535-6c2ee24944b6",
            "user_type": "email",
            "username": null,
            "email": "no.profile@fake.usace.army.mil"
        },
        {
            "id": "57329df6-9f7a-4dad-9383-4633b452efab",
            "user_type": "profile",
            "username": "AnthonyLambert",
            "email": "anthony.lambert@fake.usace.army.mil"
        },
        {
            "id": null,
            "user_type": null,
            "username": null,
            "email": "noprofile.newemail@fake.usace.army.mil"
        }
    ]`

const (
	testAlertConfigInstrumentID = "9e8f2ca4-4037-45a4-aaca-d9e598877439"
)

func TestAlertConfigs(t *testing.T) {
	arrSchema, err := gojsonschema.NewSchema(alertConfigArrayLoader)
	assert.Nil(t, err)

	tests := []HTTPTest{
		{
			Name:           "ListInstrumentAlertConfigs",
			URL:            fmt.Sprintf("/projects/%s/instruments/%s/alert_configs", testProjectID, testAlertConfigInstrumentID),
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: arrSchema,
		},
		{
			Name:           "ListProjectAlertConfigs",
			URL:            fmt.Sprintf("/projects/%s/alert_configs?alert_type_id=", testProjectID),
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: arrSchema,
		}}

	RunAll(t, tests)
}
