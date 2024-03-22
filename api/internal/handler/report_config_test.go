package handler_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xeipuuv/gojsonschema"
)

var reportConfigSchema = fmt.Sprintf(`{
    "type": "object",
    "properties": {
        "id": { "type": "string" },
        "slug": { "type": "string" },
        "name": { "type": "string" },
	"description": { "type": "string" },
        "project_id": { "type": "string" },
        "project_name": { "type": "string" },
        "creator_id": { "type": "string" },
	"creator_username": { "type": "string" },
        "create_date": { "type": "string", "format": "date-time" },
        "updater_id": {  "type": ["string", "null"] },
        "updater_username": {  "type": ["string", "null"] },
        "update_date": { "type": ["string", "null"], "format": "date-time" },
	"after": { "type": ["string", "null"], "format": "date-time" },
	"before": { "type": ["string", "null"], "format": "date-time" },
	"plot_configs": %s
    },
    "additionalProperties": false,
    "minProperties": 13
}`, IDSlugNameArrSchema)

var reportConfigObjectLoader = gojsonschema.NewStringLoader(reportConfigSchema)

var reportConfigArrayLoader = gojsonschema.NewStringLoader(fmt.Sprintf(`{
    "type": "array",
    "items": %s
}`, reportConfigSchema))

const testReportConfigID = "a6254bce-9235-4ada-afe7-8ffc3ad867e2"

const createReportConfigBody = `{
    "name": "New Test Report Config",
    "description": "create test",
    "plot_configs": [
	{"id": "cc28ca81-f125-46c6-a5cd-cc055a003c19"}
    ]
}`

const updateReportConfigBody = `{
    "name": "Updated Test Report Config",
    "description": "update test",
    "plot_configs": []
}`

func TestReportConfigurations(t *testing.T) {
	objSchema, err := gojsonschema.NewSchema(reportConfigObjectLoader)
	assert.Nil(t, err)
	arrSchema, err := gojsonschema.NewSchema(reportConfigArrayLoader)
	assert.Nil(t, err)

	tests := []HTTPTest{
		{
			Name:           "ListProjectReportConfigs",
			URL:            fmt.Sprintf("/projects/%s/report_configs", testProjectID),
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: arrSchema,
		},
		{
			Name:           "CreateReportConfig",
			URL:            fmt.Sprintf("/projects/%s/report_configs", testProjectID),
			Method:         http.MethodPost,
			Body:           createReportConfigBody,
			ExpectedStatus: http.StatusCreated,
			ExpectedSchema: objSchema,
		},
		{
			Name:           "UpdateReportConfig",
			URL:            fmt.Sprintf("/projects/%s/report_configs/%s", testProjectID, testReportConfigID),
			Method:         http.MethodPut,
			Body:           updateReportConfigBody,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: objSchema,
		},
		{
			Name:           "DeleteReportConfig",
			URL:            fmt.Sprintf("/projects/%s/report_configs/%s", testProjectID, testReportConfigID),
			Method:         http.MethodDelete,
			ExpectedStatus: http.StatusOK,
		}}

	RunAll(t, tests)
}
