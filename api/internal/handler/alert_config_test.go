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

const alertConfigEmailSchema = `{
    "type": "object",
    "properties": {
        "id": { "type": "string" },
        "user_type": { "type": "string" },
        "username": { "type": ["string", "null"] },
        "email": { "type": "string" }
    }
}`

var alertConfigSchema = fmt.Sprintf(`{
    "type": "object",
    "properties": {
        "id": { "type": "string" },
        "name": { "type": "string" },
        "body": { "type": "string" },
        "project_id": { "type": "string" },
        "alert_type_id": { "type": "string" },
        "alert_type": { "type": "string" },
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
}`, alertConfigInstrumentSchema, alertConfigEmailSchema, alertConfigSchedulerOptsSchema)

var alertConfigObjectLoader = gojsonschema.NewStringLoader(alertConfigSchema)

var alertConfigArrayLoader = gojsonschema.NewStringLoader(fmt.Sprintf(`{
    "type": "array",
    "items": %s
}`, alertConfigSchema))

const (
	testAlertConfigID           = "1efd2d85-d3ee-4388-85a0-f824a761ff8b"
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
