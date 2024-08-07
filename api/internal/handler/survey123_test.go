package handler_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xeipuuv/gojsonschema"
)

const survey123EquivalencyTableRowSchema = `{
    "type": "object",
    "properties": {
	"field_name": { "type": "string" },
	"display_name": { "type": "string" },
	"instrument_id": { "type": ["string", "null"] },
	"timeseries_id": { "type": ["string", "null"] }
    }
}`

var survey123Schema = fmt.Sprintf(`{
    "type": "object",
    "properties": {
        "id": { "type": "string" },
	"name": { "type": "string" },
	"project_id": { "type": "string" },
	"creator_id": { "type": "string" },
	"creator_username": { "type": "string" },
	"create_date": { "type": "string" },
	"updater_id": { "type": ["string", "null"] },
	"updater_username": { "type": "string" },
	"update_date": { "type": ["string", "null"] },
	"slug": { "type": "string" },
	"errors": { "type": "array", "items": { "type": "string" } },
	"rows": { "type": "array", "items": %s }
    },
    "required": [
        "id",
	"name",
	"project_id",
	"creator_id",
	"creator_username",
	"create_date",
	"rows",
	"slug",
	"errors"
    ]
}`, survey123EquivalencyTableRowSchema)

var survey123ObjectLoader = gojsonschema.NewStringLoader(survey123Schema)
var survey123ArrayLoader = gojsonschema.NewStringLoader(fmt.Sprintf(`{
    "type": "array",
    "items": %s
}`, survey123Schema))

const survey123PreviewSchema = `{
    "type": "object",
    "properties": {
        "survey123_id": { "type": "string" },
	"update_date": { "type": "string" },
	"preview": { "type": ["object", "array", "null"] }
    }
}`

var survey123PreviewLoader = gojsonschema.NewStringLoader(survey123PreviewSchema)

const (
	testSurvey123ID = "TODO"
)

const createSurvey123Body = `{
    "project_id": "5b6f4f37-7755-4cf9-bd02-94f1e9bc5984",
    "name": "Test Create Survey123"
}`

const updateSurvey123Body = `{
    "id": "TODO",
    "name": "Updated name",
    "rows": [
        {
            "field_name": "test1__depth",
            "display_name": "depth",
            "instrument_id": "a7540f69-c41e-43b3-b655-6e44097edb7e",
            "timeseries_id": "7ee902a3-56d0-4acf-8956-67ac82c03a96"
        },
        {
            "field_name": "test1__battery",
            "display_name": "battery",
            "instrument_id": "a7540f69-c41e-43b3-b655-6e44097edb7e",
            "timeseries_id": "869465fc-dc1e-445e-81f4-9979b5fadda9"
        },
        {
            "field_name": "test1__temperature",
            "display_name": "temperature",
            "instrument_id": "a7540f69-c41e-43b3-b655-6e44097edb7e",
            "timeseries_id": "869465fc-dc1e-445e-81f4-9979b5fadda9"
        }
    ]
}`

func TestSurvey123(t *testing.T) {
	arrSchema, err := gojsonschema.NewSchema(survey123ArrayLoader)
	assert.Nil(t, err)
	previewObjSchema, err := gojsonschema.NewSchema(survey123PreviewLoader)
	assert.Nil(t, err)

	tests := []HTTPTest{
		{
			Name:           "CreateSurvey123",
			URL:            fmt.Sprintf("/projects/%s/survey123", testProjectID),
			Method:         http.MethodPost,
			Body:           createSurvey123Body,
			ExpectedStatus: http.StatusCreated,
		},
		{
			Name:           "ListSurvey123sForProject",
			URL:            fmt.Sprintf("/projects/%s/survey123", testProjectID),
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: arrSchema,
		},
		{
			Name:           "GetSurvey123Preview",
			URL:            fmt.Sprintf("/projects/%s/survey123/%s/previews", testProjectID, testSurvey123ID),
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: previewObjSchema,
		},
		{
			Name:           "UpdateSurvey123",
			URL:            fmt.Sprintf("/projects/%s/survey123/%s", testProjectID, testSurvey123ID),
			Method:         http.MethodPut,
			Body:           updateSurvey123Body,
			ExpectedStatus: http.StatusOK,
		},
		{
			Name:           "DeleteSurvey123",
			URL:            fmt.Sprintf("/projects/%s/survey123/%s", testProjectID, testSurvey123ID),
			Method:         http.MethodDelete,
			ExpectedStatus: http.StatusOK,
		}}

	RunAll(t, tests)
}
