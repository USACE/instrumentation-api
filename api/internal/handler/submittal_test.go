package handler_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xeipuuv/gojsonschema"
)

const submittalSchema = `{
    "type": "object",
    "properties": {
        "id": { "type": "string" },
        "alert_config_id": { "type": "string" },
        "alert_config_name": { "type": "string" },
        "alert_type_id": { "type": "string" },
        "alert_type_name": { "type": "string" },
        "project_id": { "type": "string" },
        "submittal_status_id": { "type": "string" },
        "submittal_status_name": { "type": "string" },
        "create_date": { "type": "string", "format": "date-time" },
        "due_date": { "type": "string", "format": "date-time" },
        "completion_date": { "type": ["string", "null"], "format": "date-time" },
        "marked_as_missing": { "type": "boolean" },
        "warning_sent": { "type": "boolean" }
    },
    "additionalProperties": true
}`

var submittalArrayLoader = gojsonschema.NewStringLoader(fmt.Sprintf(`{
    "type": "array",
    "items": %s
}`, submittalSchema))

const (
	testSubmittalID            = "b8c1c297-d1d5-4cee-b949-72299b330617"
	testSubmittalInstrumentID  = "9e8f2ca4-4037-45a4-aaca-d9e598877439"
	testSubmittalAlertConfigID = "1efd2d85-d3ee-4388-85a0-f824a761ff8b"
)

func TestSubmittals(t *testing.T) {
	arrSchema, err := gojsonschema.NewSchema(submittalArrayLoader)
	assert.Nil(t, err)

	tests := []HTTPTest{
		{
			Name:           "ListProjectSubmittals",
			URL:            fmt.Sprintf("/projects/%s/submittals?missing=false", testProjectID),
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: arrSchema,
		},
		{
			Name:           "ListInstrumentSubmittals",
			URL:            fmt.Sprintf("/instruments/%s/submittals?missing=false", testSubmittalInstrumentID),
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: arrSchema,
		},
		{
			Name:           "ListAlertConfigSubmittals",
			URL:            fmt.Sprintf("/alert_configs/%s/submittals?missing=false", testSubmittalAlertConfigID),
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: arrSchema,
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
