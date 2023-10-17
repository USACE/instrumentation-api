package handler_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/xeipuuv/gojsonschema"
)

const alertSubSchema = `{
    "type": "object",
    "properties": {
        "id": { "type": "string" },
        "alert_config_id": { "type": "string" },
        "profile_id": { "type": "string" },
        "mute_ui": { "type": "boolean" },
        "mute_notify": { "type": "boolean" }
    },
    "required": ["id", "alert_config_id", "profile_id", "mute_ui", "mute_notify"],
    "additionalProperties": false
}`

var alertSubObjectSchema = gojsonschema.NewStringLoader(alertSubSchema)

var alertSubArraySchema = gojsonschema.NewStringLoader(fmt.Sprintf(`{
    "type": "array",
    "items": %s
}`, alertSubSchema))

const alertSubAlertConfigInstrumentSchema = `{
    "type": "object",
    "properties": {
        "instrument_id": { "type": "string" },
        "instrument_name": { "type": "string" }
    }
}`

var alertSubAlertConfigSchema = fmt.Sprintf(`{
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

var alertSubAlertConfigObjectSchema = gojsonschema.NewStringLoader(alertSubAlertConfigSchema)

var alertSubAlertConfigArraySchema = gojsonschema.NewStringLoader(fmt.Sprintf(`{
    "type": "array",
    "items": %s
}`, alertSubAlertConfigSchema))

const (
	testAlertSubID            = "197d6140-f273-4c50-a87f-dec3f809663b"
	testAlertSubAlertID       = "e070be13-ef17-40f3-99c8-fef3ee1b9fb5"
	testAlertSubInstrumentID  = "a7540f69-c41e-43b3-b655-6e44097edb7e"
	testAlertSubAlertConfigID = "6f3dfe9f-4664-4c78-931f-32ffac6d2d43"
)

const updateInstrumentAlertSubscriptionBody = `{
    "id": "197d6140-f273-4c50-a87f-dec3f809663b",
    "alert_config_id": "243e9d32-2cba-4f12-9abe-63adc09fc5dd",
    "profile_id": "96c01ff3-edc9-44f0-8690-191cc2281a12",
    "mute_ui": false,
    "mute_notify": true
}`

func TestAlertSubscriptions(t *testing.T) {
	tests := []HTTPTest{
		{
			Name:           "SubscribeProfileToInstrumentAlert",
			URL:            fmt.Sprintf("/projects/%s/instruments/%s/alert_configs/%s/subscribe", testProjectID, testAlertSubInstrumentID, testAlertSubAlertConfigID),
			Method:         http.MethodPost,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: &alertSubObjectSchema,
		},
		{
			Name:           "ListMyInstrumentAlertSubscriptions",
			URL:            "/my_alert_subscriptions",
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: &alertSubArraySchema,
		},
		{
			Name:           "UpdateInstrumentAlertSubscription",
			URL:            fmt.Sprintf("/alert_subscriptions/%s", testAlertSubID),
			Method:         http.MethodPut,
			Body:           updateInstrumentAlertSubscriptionBody,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: &alertSubObjectSchema,
		},
		{
			Name:           "ListAlertsForInstrument",
			URL:            fmt.Sprintf("/projects/%s/instruments/%s/alerts", testProjectID, testAlertSubInstrumentID),
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: &alertSubAlertConfigArraySchema,
		},
		{
			Name:           "ListMyAlerts",
			URL:            "/my_alerts",
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: &alertSubAlertConfigArraySchema,
		},
		{
			Name:           "DoAlertRead",
			URL:            fmt.Sprintf("/my_alerts/%s/read", testAlertSubAlertID),
			Method:         http.MethodPost,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: &alertSubAlertConfigObjectSchema,
		},
		{
			Name:           "DoAlertUnread",
			URL:            fmt.Sprintf("/my_alerts/%s/unread", testAlertSubAlertID),
			Method:         http.MethodPost,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: &alertSubAlertConfigObjectSchema,
		},
		{
			Name:           "UnsubscribeProfileToInstrumentAlert",
			URL:            fmt.Sprintf("/projects/%s/instruments/%s/alert_configs/%s/unsubscribe", testProjectID, testAlertSubInstrumentID, testAlertSubAlertConfigID),
			Method:         http.MethodPost,
			ExpectedStatus: http.StatusOK,
		}}

	RunAll(t, tests)
}
