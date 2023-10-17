package handler_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/xeipuuv/gojsonschema"
)

const collectionGroupSchema = `{
    "type": "object",
    "properties": {
        "id": { "type": "string" },
        "project_id": { "type": "string" },
        "slug": { "type": "string" },
        "name": { "type": "string" },
        "creator": { "type": "string" },
        "create_date": { "type": "string", "format": "date-time" },
        "updater": {  "type": ["string", "null"] },
        "update_date": { "type": ["string", "null"], "format": "date-time" }
    },
    "required": ["id", "project_id", "name", "slug", "creator", "create_date", "updater", "update_date"],
    "additionalProperties": false
}`

var collectionGroupArraySchema = gojsonschema.NewStringLoader(fmt.Sprintf(`{
    "type": "array",
    "items": %s
}`, collectionGroupSchema))

const collectionGroupDetailsSchema = `{
    "type": "object",
    "properties": {
        "id": { "type": "string" },
        "project_id": { "type": "string" },
        "slug": { "type": "string" },
        "name": { "type": "string" },
        "creator": { "type": "string" },
        "create_date": { "type": "string", "format": "date-time" },
        "updater": {  "type": ["string", "null"] },
        "update_date": { "type": ["string", "null"], "format": "date-time" },
        "timeseries": {
            "type": "array",
            "items": {
                "type": "object",
                "properties": {
                    "id": { "type": "string" },
                    "slug": { "type": "string" },
                    "name": { "type": "string" },
                    "variable": {"type": "string"},
                    "project_id": {"type": "string"},
                    "project": {"type": "string"},
                    "project_slug": {"type": "string"},
                    "instrument_id": { "type": "string" },
                    "instrument": { "type": "string"  },
                    "instrument_slug": {"type": "string"},
                    "parameter_id": { "type": "string" },
                    "parameter": { "type": "string"  },
                    "unit_id": { "type": "string" },
                    "unit": { "type": "string"  },
                    "latest_time": {"type": "string", "format": "date-time"},
                    "latest_value": {"type": "number"},
                    "is_computed": {"type": "boolean"},
                },
                "required": ["id", "slug", "name", "variable", "project_id", "project", "project_slug", "instrument_id", "instrument", "instrument_slug", "parameter_id", "parameter", "unit_id", "unit", "latest_time", "latest_value", "is_computed"],
                "additionalProperties": false
            },
        },
    },
    "required": ["id", "project_id", "name", "slug", "creator", "create_date", "updater", "update_date", "timeseries"],
    "additionalProperties": false
}`

var collectionGroupDetailsObjectSchema = gojsonschema.NewStringLoader(collectionGroupDetailsSchema)

const testCollectionGroupID = "30b32cb1-0936-42c4-95d1-63a7832a57db"

func TestCollectionGroups(t *testing.T) {
	tests := []HTTPTest{
		{
			Name:           "ListCollectionGroups",
			URL:            fmt.Sprintf("/projects/%s/collection_groups", testProjectID),
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: &collectionGroupArraySchema,
		},
		{
			Name:           "GetCollectionGroupDetails",
			URL:            fmt.Sprintf("/projects/%s/collection_groups/%s", testProjectID, testCollectionGroupID),
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: &collectionGroupDetailsObjectSchema,
		},
		{
			Name:           "DeleteCollectionGroup",
			URL:            fmt.Sprintf("/projects/%s/collection_groups/%s", testProjectID, testCollectionGroupID),
			Method:         http.MethodDelete,
			ExpectedStatus: http.StatusOK,
		}}

	RunAll(t, tests)
}
