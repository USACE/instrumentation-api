package handler_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/USACE/instrumentation-api/api/internal/model"
)

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
	tests := []HTTPTest[model.AlertSubscription]{
		{
			Name:                 "SubscribeProfileToInstrumentAlert",
			URL:                  fmt.Sprintf("/projects/%s/instruments/%s/alert_configs/%s/subscribe", testProjectID, testAlertSubInstrumentID, testAlertSubAlertConfigID),
			Method:               http.MethodPost,
			ExpectedStatus:       http.StatusOK,
			ExpectedResponseType: jsonObj,
		},
		{
			Name:                 "ListMyInstrumentAlertSubscriptions",
			URL:                  "/my_alert_subscriptions",
			Method:               http.MethodGet,
			ExpectedStatus:       http.StatusOK,
			ExpectedResponseType: jsonArr,
		},
		{
			Name:                 "UpdateMyAlertSubscription",
			URL:                  fmt.Sprintf("/alert_subscriptions/%s", testAlertSubID),
			Method:               http.MethodPut,
			Body:                 updateInstrumentAlertSubscriptionBody,
			ExpectedStatus:       http.StatusOK,
			ExpectedResponseType: jsonObj,
		},
		{
			Name:           "UnsubscribeProfileToInstrumentAlert",
			URL:            fmt.Sprintf("/projects/%s/instruments/%s/alert_configs/%s/unsubscribe", testProjectID, testAlertSubInstrumentID, testAlertSubAlertConfigID),
			Method:         http.MethodPost,
			ExpectedStatus: http.StatusOK,
		}}

	RunAll(t, tests)
}
