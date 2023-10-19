package handler_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
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

var alertSubObjectLoader = gojsonschema.NewStringLoader(alertSubSchema)

var alertSubArrayLoader = gojsonschema.NewStringLoader(fmt.Sprintf(`{
    "type": "array",
    "items": %s
}`, alertSubSchema))

const (
	testAlertSubID            = "197d6140-f273-4c50-a87f-dec3f809663b"
	testAlertSubAlertID       = "e070be13-ef17-40f3-99c8-fef3ee1b9fb5"
	testAlertSubInstrumentID  = "a7540f69-c41e-43b3-b655-6e44097edb7e"
	testAlertSubAlertConfigID = "6f3dfe9f-4664-4c78-931f-32ffac6d2d43"
)

const updateInstrumentAlertSubscriptionBody = `{
    "id": "197d6140-f273-4c50-a87f-dec3f809663b",
    "alert_config_id": "1efd2d85-d3ee-4388-85a0-f824a761ff8b",
    "profile_id": "57329df6-9f7a-4dad-9383-4633b452efab",
    "mute_ui": false,
    "mute_notify": true
}`

func TestAlertSubscriptions(t *testing.T) {
	objSchema, err := gojsonschema.NewSchema(alertSubObjectLoader)
	assert.Nil(t, err)
	arrSchema, err := gojsonschema.NewSchema(alertSubArrayLoader)
	assert.Nil(t, err)

	tests := []HTTPTest{
		{
			Name:           "SubscribeProfileToInstrumentAlert",
			URL:            fmt.Sprintf("/projects/%s/instruments/%s/alert_configs/%s/subscribe", testProjectID, testAlertSubInstrumentID, testAlertSubAlertConfigID),
			Method:         http.MethodPost,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: objSchema,
		},
		{
			Name:           "ListMyInstrumentAlertSubscriptions",
			URL:            "/my_alert_subscriptions",
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: arrSchema,
		},
		{
			Name:           "UpdateMyAlertSubscription",
			URL:            fmt.Sprintf("/alert_subscriptions/%s", testAlertSubID),
			Method:         http.MethodPut,
			Body:           updateInstrumentAlertSubscriptionBody,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: objSchema,
		},
		{
			Name:           "UnsubscribeProfileToInstrumentAlert",
			URL:            fmt.Sprintf("/projects/%s/instruments/%s/alert_configs/%s/unsubscribe", testProjectID, testAlertSubInstrumentID, testAlertSubAlertConfigID),
			Method:         http.MethodPost,
			ExpectedStatus: http.StatusOK,
		}}

	RunAll(t, tests)
}
