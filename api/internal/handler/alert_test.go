package handler_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xeipuuv/gojsonschema"
)

const alertSubAlertConfigInstrumentSchema = `{
    "type": "object",
    "properties": {
        "instrument_id": { "type": "string" },
        "instrument_name": { "type": "string" }
    }
}`

var alertSchema = fmt.Sprintf(`{
    "type": "object",
    "properties": {
        "id": { "type": "string" },
        "alert_config_id": { "type": "string" },
        "project_id": { "type": "string" },
        "project_name": { "type": "string" },
        "instruments": { "type": "array", "items": %s },
        "name": { "type": "string" },
        "body": { "type": "string" },
        "create_date": { "type": "string", "format": "date-time" }
    },
    "required": ["id", "alert_config_id", "project_id", "project_name", "instruments", "name", "body", "create_date"],
    "additionalProperties": true
}`, alertSubAlertConfigInstrumentSchema)

var alertObjectLoader = gojsonschema.NewStringLoader(alertSchema)

var alertArrayLoader = gojsonschema.NewStringLoader(fmt.Sprintf(`{
    "type": "array",
    "items": %s
}`, alertSchema))

func TestAlerts(t *testing.T) {
	objSchema, err := gojsonschema.NewSchema(alertObjectLoader)
	assert.Nil(t, err)
	arrSchema, err := gojsonschema.NewSchema(alertArrayLoader)
	assert.Nil(t, err)

	tests := []HTTPTest{
		{
			Name:           "ListAlertsForInstrument",
			URL:            fmt.Sprintf("/projects/%s/instruments/%s/alerts", testProjectID, testAlertSubInstrumentID),
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: arrSchema,
		},
		{
			Name:           "ListMyAlerts",
			URL:            "/my_alerts",
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: arrSchema,
		},
		{
			Name:           "DoAlertRead",
			URL:            fmt.Sprintf("/my_alerts/%s/read", testAlertSubAlertID),
			Method:         http.MethodPost,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: objSchema,
		},
		{
			Name:           "DoAlertUnread",
			URL:            fmt.Sprintf("/my_alerts/%s/unread", testAlertSubAlertID),
			Method:         http.MethodPost,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: objSchema,
		},
	}

	RunAll(t, tests)
}
