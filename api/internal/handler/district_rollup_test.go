package handler_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xeipuuv/gojsonschema"
)

const districtRollupSchema = `{
    "type": "object",
    "properties": {
        "alert_type_id": { "type": "string" },
        "office_id": { "type": ["string", "null"] },
        "district_initials": { "type": ["string", "null"] },
        "project_name": { "type": "string" },
        "project_id": { "type": "string" },
        "month": { "type": "string", "format": "date-time" },
        "expected_total_submittals": { "type": "number" },
        "actual_total_submittals": { "type": "number" },
        "red_submittals": { "type": "number" },
        "yellow_submittals": { "type": "number" },
        "green_submittals": { "type": "number" }
    },
    "additionalProperties": false
}`

var districtRollupArrayLoader = gojsonschema.NewStringLoader(fmt.Sprintf(`{
    "type": "array",
    "items": %s
}`, districtRollupSchema))

func TestDistrictRollup(t *testing.T) {
	arrSchema, err := gojsonschema.NewSchema(districtRollupArrayLoader)
	assert.Nil(t, err)

	tests := []HTTPTest{
		{
			Name:           "ListProjectEvaluationDistrictRollup",
			URL:            fmt.Sprintf("/projects/%s/district_rollup/evaluation_submittals?from_timestamp_month=&to_timestamp_month=", testProjectID),
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: arrSchema,
		},
		{
			Name:           "ListProjectMeasurementDistrictRollup",
			URL:            fmt.Sprintf("/projects/%s/district_rollup/measurement_submittals?from_timestamp_month=&to_timestamp_month=", testProjectID),
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: arrSchema,
		}}

	RunAll(t, tests)
}
