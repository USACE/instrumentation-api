package handler_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/USACE/instrumentation-api/api/internal/model"
)

const (
	testEvaluationID           = "f7169aca-aa5f-4a0b-9fcc-609bb5c2bd7b"
	testEvaluationInstrumentID = "a7540f69-c41e-43b3-b655-6e44097edb7e"
)

const createEvaluationBody = `{
    "project_id": "5b6f4f37-7755-4cf9-bd02-94f1e9bc5984",
    "name": "New Test Evaluation",
    "body": "New Test Evaluation Description",
    "start_date": "2023-05-16T13:19:41.441328Z",
    "end_date": "2023-06-16T13:19:41.441328Z",
    "submittal_id": "f8189297-f1a6-489d-9ea7-f1a0ffc30153",
    "instruments": [
        {"instrument_id": "a7540f69-c41e-43b3-b655-6e44097edb7e"}
    ]
}`

const updateEvaluationBody = `{
    "id": "add252bf-2fa7-4824-b129-e4d0ff42dffa",
    "project_id": "5b6f4f37-7755-4cf9-bd02-94f1e9bc5984",
    "name": "Updated Test Evaluation",
    "body": "Updated Test Evaluation Description",
    "start_date": "2023-07-16T13:19:41.441328Z",
    "end_date": "2023-08-16T13:19:41.441328Z",
    "instruments": []
}`

func TestEvaluation(t *testing.T) {
	tests := []HTTPTest[model.Evaluation]{
		{
			Name:                 "GetEvaluation",
			URL:                  fmt.Sprintf("/projects/%s/evaluations/%s", testProjectID, testEvaluationID),
			Method:               http.MethodGet,
			ExpectedStatus:       http.StatusOK,
			ExpectedResponseType: jsonObj,
		},
		{
			Name:                 "ListInstrumentEvaluations",
			URL:                  fmt.Sprintf("/projects/%s/instruments/%s/evaluations", testProjectID, testEvaluationInstrumentID),
			Method:               http.MethodGet,
			ExpectedStatus:       http.StatusOK,
			ExpectedResponseType: jsonArr,
		},
		{
			Name:                 "ListProjectEvaluations",
			URL:                  fmt.Sprintf("/projects/%s/evaluations?alert_config_id=", testProjectID),
			Method:               http.MethodGet,
			ExpectedStatus:       http.StatusOK,
			ExpectedResponseType: jsonArr,
		},
		{
			Name:                 "CreateEvaluation",
			URL:                  fmt.Sprintf("/projects/%s/evaluations", testProjectID),
			Method:               http.MethodPost,
			Body:                 createEvaluationBody,
			ExpectedStatus:       http.StatusCreated,
			ExpectedResponseType: jsonObj,
		},
		{
			Name:                 "UpdateEvaluation",
			URL:                  fmt.Sprintf("/projects/%s/evaluations/%s", testProjectID, testEvaluationID),
			Method:               http.MethodPut,
			Body:                 updateEvaluationBody,
			ExpectedStatus:       http.StatusOK,
			ExpectedResponseType: jsonObj,
		},
		{
			Name:           "DeleteEvaluation",
			URL:            fmt.Sprintf("/projects/%s/evaluations/%s", testProjectID, testEvaluationID),
			Method:         http.MethodDelete,
			ExpectedStatus: http.StatusOK,
		}}

	RunAll(t, tests)
}
