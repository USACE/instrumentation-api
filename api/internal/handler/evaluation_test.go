package handler_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xeipuuv/gojsonschema"
)

const evaluationInstrumentSchema = `{
    "type": "object",
    "properties": {
        "instrument_id": { "type": "string" },
        "instrument_name": { "type": "string" }
    }
}`

var evaluationSchema = fmt.Sprintf(`{
    "type": "object",
    "properties": {
        "id": { "type": "string" },
        "name": { "type": "string" },
        "body": { "type": "string" },
        "project_id": { "type": "string" },
        "project_name": { "type": "string" },
        "alert_config_id": { "type": ["string", "null"] },
        "submittal_id": { "type": ["string", "null"] },
        "alert_config_name": { "type": ["string", "null"] },
        "start_date": { "type": "string", "format": "date-time" },
        "end_date": { "type": "string", "format": "date-time" },
        "instruments": { "type": "array", "items": %s },
        "creator": { "type": "string" },
        "creator_username": { "type": "string" },
        "create_date": { "type": "string", "format": "date-time" },
        "updater": { "type": ["string", "null"] },
        "updater_username": { "type": ["string", "null"] },
        "update_date": { "type": ["string", "null"], "format": "date-time" }
    },
    "additionalProperties": false
}`, evaluationInstrumentSchema)

var evaluationObjectLoader = gojsonschema.NewStringLoader(evaluationSchema)

var evaluationArrayLoader = gojsonschema.NewStringLoader(fmt.Sprintf(`{
    "type": "array",
    "items": %s
}`, evaluationSchema))

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
	objSchema, err := gojsonschema.NewSchema(evaluationObjectLoader)
	assert.Nil(t, err)
	arrSchema, err := gojsonschema.NewSchema(evaluationArrayLoader)
	assert.Nil(t, err)

	tests := []HTTPTest{
		{
			Name:           "GetEvaluation",
			URL:            fmt.Sprintf("/projects/%s/evaluations/%s", testProjectID, testEvaluationID),
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: objSchema,
		},
		{
			Name:           "ListInstrumentEvaluations",
			URL:            fmt.Sprintf("/projects/%s/instruments/%s/evaluations", testProjectID, testEvaluationInstrumentID),
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: arrSchema,
		},
		{
			Name:           "ListProjectEvaluations",
			URL:            fmt.Sprintf("/projects/%s/evaluations?alert_config_id=", testProjectID),
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: arrSchema,
		},
		{
			Name:           "CreateEvaluation",
			URL:            fmt.Sprintf("/projects/%s/evaluations", testProjectID),
			Method:         http.MethodPost,
			Body:           createEvaluationBody,
			ExpectedStatus: http.StatusCreated,
			ExpectedSchema: objSchema,
		},
		{
			Name:           "UpdateEvaluation",
			URL:            fmt.Sprintf("/projects/%s/evaluations/%s", testProjectID, testEvaluationID),
			Method:         http.MethodPut,
			Body:           updateEvaluationBody,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: objSchema,
		},
		{
			Name:           "DeleteEvaluation",
			URL:            fmt.Sprintf("/projects/%s/evaluations/%s", testProjectID, testEvaluationID),
			Method:         http.MethodDelete,
			ExpectedStatus: http.StatusOK,
		}}

	RunAll(t, tests)
}
