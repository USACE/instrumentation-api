package handler_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/USACE/instrumentation-api/api/internal/model"
)

func TestAlerts(t *testing.T) {
	tests := []HTTPTest[model.Alert]{
		{
			Name:                 "ListAlertsForInstrument",
			URL:                  fmt.Sprintf("/projects/%s/instruments/%s/alerts", testProjectID, testAlertSubInstrumentID),
			Method:               http.MethodGet,
			ExpectedStatus:       http.StatusOK,
			ExpectedResponseType: jsonArr,
		},
		{
			Name:                 "ListMyAlerts",
			URL:                  "/my_alerts",
			Method:               http.MethodGet,
			ExpectedStatus:       http.StatusOK,
			ExpectedResponseType: jsonArr,
		},
		{
			Name:                 "DoAlertRead",
			URL:                  fmt.Sprintf("/my_alerts/%s/read", testAlertSubAlertID),
			Method:               http.MethodPost,
			ExpectedStatus:       http.StatusOK,
			ExpectedResponseType: jsonObj,
		},
		{
			Name:                 "DoAlertUnread",
			URL:                  fmt.Sprintf("/my_alerts/%s/unread", testAlertSubAlertID),
			Method:               http.MethodPost,
			ExpectedStatus:       http.StatusOK,
			ExpectedResponseType: jsonObj,
		},
	}

	RunAll(t, tests)
}
