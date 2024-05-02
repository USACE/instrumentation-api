package handler_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xeipuuv/gojsonschema"
)

const globalOverridesSchema = `{
    "type": "object",
    "properties": {
        "date_range": {
            "type": "object",
            "properties": {
                "value": { "type": ["string", "null"] },
                "enabled": { "type": "boolean" }
            }
        },
        "show_masked": {
            "type": "object",
            "properties": {
                "value": { "type": "boolean" },
                "enabled": { "type": "boolean" }
            }
        },
        "show_nonvalidated": {
            "type": "object",
            "properties": {
                "value": { "type": "boolean" },
                "enabled": { "type": "boolean" }
            }
        }
    }
}`

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
	"global_overrides": %s,
	"plot_configs": %s
    },
    "additionalProperties": false,
    "required": [
        "id","slug","name","description","project_id","project_name","creator_id",
        "creator_username","create_date","global_overrides","plot_configs"
    ]
}`, globalOverridesSchema, IDSlugNameArrSchema)

var reportConfigObjectLoader = gojsonschema.NewStringLoader(reportConfigSchema)

var reportConfigArrayLoader = gojsonschema.NewStringLoader(fmt.Sprintf(`{
    "type": "array",
    "items": %s
}`, reportConfigSchema))

const reportDownloadJobSchema = `{
    "type": "object",
    "properties": {
        "id": { "type": "string" },
        "report_config_id": { "type": "string" },
        "creator": { "type": "string" },
        "create_date": { "type": "string" },
        "status": { "type": "string" },
        "file_key": { "type": ["string", "null"] },
        "file_expiry": { "type": ["string", "null"] },
        "progress": { "type": "number" },
        "progress_update_date": { "type": "string" }
    }
}`

var reportDownloadJobObjectLoader = gojsonschema.NewStringLoader(reportDownloadJobSchema)

const testReportConfigID = "a6254bce-9235-4ada-afe7-8ffc3ad867e2"
const testJobID = "e90dbcc9-7bf4-4402-80ea-c0cdbbb91c6d"
const testUpdateJobID = "61b69ef2-2c73-4143-930d-3832400ba8f2"

const createReportConfigBody = `{
    "name": "New Test Report Config",
    "description": "create test",
    "plot_configs": [
	{"id": "cc28ca81-f125-46c6-a5cd-cc055a003c19"}
    ],
    "global_overrides": {
        "date_range": {
             "value": "2023-01-01 2024-01-01",
             "enabled": true
        },
        "show_masked": {
             "value": true,
             "enabled": true
        },
        "show_nonvalidated": {
             "value": true,
             "enabled": true
        }
    }
}`

const updateReportConfigBody = `{
    "name": "Updated Test Report Config",
    "description": "update test",
    "plot_configs": [],
    "global_overrides": {
        "date_range": {
             "value": null,
             "enabled": false
        },
        "show_masked": {
             "value": true,
             "enabled": false
        },
        "show_nonvalidated": {
             "value": true,
             "enabled": false
        }
    }
}`

const updateReportDownloadJobBody = `{
	"status": "FAIL",
	"progress": 0
}`

func TestReportConfigs(t *testing.T) {
	objSchema, err := gojsonschema.NewSchema(reportConfigObjectLoader)
	assert.Nil(t, err)
	arrSchema, err := gojsonschema.NewSchema(reportConfigArrayLoader)
	assert.Nil(t, err)
	jobObjSchema, err := gojsonschema.NewSchema(reportDownloadJobObjectLoader)
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
			Name:           "GetReportDownloadJob",
			URL:            fmt.Sprintf("/projects/%s/report_configs/%s/jobs/%s", testProjectID, testReportConfigID, testJobID),
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: jobObjSchema,
		},
		{
			Name:           "CreateReportDownloadJob",
			URL:            fmt.Sprintf("/projects/%s/report_configs/%s/jobs", testProjectID, testReportConfigID),
			Method:         http.MethodPost,
			ExpectedStatus: http.StatusCreated,
			ExpectedSchema: jobObjSchema,
		},
		{
			Name:           "UpdateReportDownloadJob",
			URL:            fmt.Sprintf("/projects/%s/report_configs/%s/jobs/%s?key=%s", testProjectID, testReportConfigID, testUpdateJobID, mockAppKey),
			Body:           updateReportDownloadJobBody,
			Method:         http.MethodPut,
			ExpectedStatus: http.StatusOK,
		},
		{
			Name:           "UpdateReportConfig",
			URL:            fmt.Sprintf("/projects/%s/report_configs/%s", testProjectID, testReportConfigID),
			Method:         http.MethodPut,
			Body:           updateReportConfigBody,
			ExpectedStatus: http.StatusOK,
		},
		{
			Name:           "DeleteReportConfig",
			URL:            fmt.Sprintf("/projects/%s/report_configs/%s", testProjectID, testReportConfigID),
			Method:         http.MethodDelete,
			ExpectedStatus: http.StatusOK,
		}}

	RunAll(t, tests)
}
