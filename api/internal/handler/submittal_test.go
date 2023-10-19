package handler_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/USACE/instrumentation-api/api/internal/model"
)

const (
	testSubmittalID            = "b8c1c297-d1d5-4cee-b949-72299b330617"
	testSubmittalInstrumentID  = "9e8f2ca4-4037-45a4-aaca-d9e598877439"
	testSubmittalAlertConfigID = "1efd2d85-d3ee-4388-85a0-f824a761ff8b"
)

func TestSubmittals(t *testing.T) {
	tests := []HTTPTest[model.Submittal]{
		{
			Name:                 "ListProjectSubmittals",
			URL:                  fmt.Sprintf("/projects/%s/submittals?missing=false", testProjectID),
			Method:               http.MethodGet,
			ExpectedStatus:       http.StatusOK,
			ExpectedResponseType: jsonArr,
		},
		{
			Name:                 "ListInstrumentSubmittals",
			URL:                  fmt.Sprintf("/instruments/%s/submittals?missing=false", testSubmittalInstrumentID),
			Method:               http.MethodGet,
			ExpectedStatus:       http.StatusOK,
			ExpectedResponseType: jsonArr,
		},
		{
			Name:                 "ListAlertConfigSubmittals",
			URL:                  fmt.Sprintf("/alert_configs/%s/submittals?missing=false", testSubmittalAlertConfigID),
			Method:               http.MethodGet,
			ExpectedStatus:       http.StatusOK,
			ExpectedResponseType: jsonArr,
		},
		{
			Name:           "VerifyMissingSubmittal",
			URL:            fmt.Sprintf("/submittals/%s/verify_missing", testSubmittalID),
			Method:         http.MethodPut,
			ExpectedStatus: http.StatusOK,
		},
		{
			Name:           "VerifyMissingAlertConfigSubmittals",
			URL:            fmt.Sprintf("/alert_configs/%s/submittals/verify_missing", testSubmittalAlertConfigID),
			Method:         http.MethodPut,
			ExpectedStatus: http.StatusOK,
		}}

	RunAll(t, tests)
}
