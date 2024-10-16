package handler_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xeipuuv/gojsonschema"
)

const collectionGroupSchema = `{
    "type": "object",
    "properties": {
        "id": { "type": "string" },
        "project_id": { "type": "string" },
        "slug": { "type": "string" },
        "name": { "type": "string" },
        "creator_id": { "type": "string" },
        "create_date": { "type": "string", "format": "date-time" },
        "updater_id": {  "type": ["string", "null"] },
        "update_date": { "type": ["string", "null"], "format": "date-time" }
    },
    "required": ["id", "project_id", "name", "slug", "creator_id", "create_date", "updater_id", "update_date"],
    "additionalProperties": false
}`

var collectionGroupArrayLoader = gojsonschema.NewStringLoader(fmt.Sprintf(`{
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
        "creator_id": { "type": "string" },
        "create_date": { "type": "string", "format": "date-time" },
        "updater_id": {  "type": ["string", "null"] },
        "update_date": { "type": ["string", "null"], "format": "date-time" },
        "timeseries": {
            "type": "array",
            "items": {
                "type": "object",
                "properties": {
                    "id": { "type": "string" },
                    "slug": { "type": "string" },
                    "name": { "type": "string" },
                    "variable": { "type": "string" },
                    "instrument_id": { "type": "string" },
                    "instrument": { "type": "string" },
                    "instrument_slug": {"type": "string" },
                    "parameter_id": { "type": "string" },
                    "parameter": { "type": "string" },
                    "unit_id": { "type": "string" },
                    "unit": { "type": "string" },
                    "latest_time": {"type": "string", "format": "date-time" },
                    "latest_value": {"type": "number" },
                    "is_computed": { "type": "boolean" },
                    "type": { "type": "string" }
                },
                "required": ["id", "slug", "name", "variable", "instrument_id", "instrument", "instrument_slug", "parameter_id", "parameter", "unit_id", "unit", "latest_time", "latest_value", "is_computed", "type"],
                "additionalProperties": false
            }
        }
    },
    "required": ["id", "project_id", "name", "slug", "creator_id", "create_date", "updater_id", "update_date", "timeseries"],
    "additionalProperties": false
}`

var collectionGroupDetailsObjectLoader = gojsonschema.NewStringLoader(collectionGroupDetailsSchema)

const testCollectionGroupID = "30b32cb1-0936-42c4-95d1-63a7832a57db"

func TestCollectionGroups(t *testing.T) {
	objSchema, err := gojsonschema.NewSchema(collectionGroupDetailsObjectLoader)
	assert.Nil(t, err)
	if err != nil {
		t.Log("invalid object schema")
	}
	arrSchema, err := gojsonschema.NewSchema(collectionGroupArrayLoader)
	assert.Nil(t, err)
	if err != nil {
		t.Log("invalid array schema")
	}

	tests := []HTTPTest{
		{
			Name:           "GetCollectionGroupDetails",
			URL:            fmt.Sprintf("/projects/%s/collection_groups/%s", testProjectID, testCollectionGroupID),
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: objSchema,
		},
		{
			Name:           "ListCollectionGroups",
			URL:            fmt.Sprintf("/projects/%s/collection_groups", testProjectID),
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: arrSchema,
		},
		{
			Name:           "DeleteCollectionGroup",
			URL:            fmt.Sprintf("/projects/%s/collection_groups/%s", testProjectID, testCollectionGroupID),
			Method:         http.MethodDelete,
			ExpectedStatus: http.StatusOK,
		}}

	RunAll(t, tests)
}
