package handler_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xeipuuv/gojsonschema"
)

const instrumentGroupSchema = `{
    "type": "object",
    "properties": {
        "id": { "type": "string" },
        "slug": { "type": "string" },
        "name": { "type": "string" },
        "description": { "type": "string" },
        "creator": { "type": "string" },
        "create_date": { "type": "string", "format": "date-time" },
        "updater": {  "type": ["string", "null"] },
        "update_date": { "type": ["string", "null"], "format": "date-time" },
        "project_id": { "type": ["string", "null"] },
        "instrument_count": { "type": "number" },
        "timeseries_count": { "type": "number" }
    },
    "required": ["id", "slug", "name", "description", "creator", "create_date", "updater", "update_date", "project_id"],
    "additionalProperties": false
}`

var instrumentGroupObjectLoader = gojsonschema.NewStringLoader(instrumentGroupSchema)

var instrumentGroupArrayLoader = gojsonschema.NewStringLoader(fmt.Sprintf(`{
    "type": "array",
    "items": %s
}`, instrumentGroupSchema))

const (
	testInstrumentGroupID           = "d0916e8a-39a6-4f2f-bd31-879881f8b40c"
	testInstrumentGroupInstrumentID = "9e8f2ca4-4037-45a4-aaca-d9e598877439"
	testInstrumentGroupTimeAfter    = "1900-01-01T00:00:00.00Z"
	testInstrumentGroupTimeBefore   = "2025-12-31T00:00:00.00Z"
	testInstrumentGroupThreshold    = "1000"
)

const createInstrumentGroupBulkArrayBody = `[{
    "name": "Test Instrument Group 100",
    "project_id": "5b6f4f37-7755-4cf9-bd02-94f1e9bc5984"
},
{
    "name": "Test Instrument Group 101",
    "project_id": "5b6f4f37-7755-4cf9-bd02-94f1e9bc5984"
},
{
    "name": "Test Instrument Group 102",
    "project_id": "5b6f4f37-7755-4cf9-bd02-94f1e9bc5984"
}]`

const createInstrumentGroupBulkObjectBody = `{
    "name": "Test Instrument Group 500",
    "project_id": "5b6f4f37-7755-4cf9-bd02-94f1e9bc5984",
    "description": "This is a sample instrument group created by integration tests"
}`

const updateInstrumentGroupBody = `{
    "id": "d0916e8a-39a6-4f2f-bd31-879881f8b40c",
    "name": "Updated Name for Instrument Group 1",
    "description": "A sample instrument group created by integration tests",
    "project_id": null
}`

const addInstrumentToInstrumentGroupBody = `{"id": "9e8f2ca4-4037-45a4-aaca-d9e598877439"}`

func TestInstrumentGroups(t *testing.T) {
	objSchema, err := gojsonschema.NewSchema(instrumentGroupObjectLoader)
	assert.Nil(t, err)
	arrSchema, err := gojsonschema.NewSchema(instrumentGroupArrayLoader)
	assert.Nil(t, err)

	tests := []HTTPTest{
		{
			Name:           "GetInstrumentGroup",
			URL:            fmt.Sprintf("/instrument_groups/%s", testInstrumentGroupID),
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: objSchema,
		},
		{
			Name:           "ListInstrumentGroups",
			URL:            "/instrument_groups",
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: arrSchema,
		},
		{
			Name:           "CreateInstrumentGroupBulk_Array",
			URL:            "/instrument_groups",
			Method:         http.MethodPost,
			Body:           createInstrumentGroupBulkArrayBody,
			ExpectedStatus: http.StatusCreated,
			ExpectedSchema: arrSchema,
		},
		{
			Name:           "CreateInstrumentGroupBulk_Object",
			URL:            "/instrument_groups",
			Method:         http.MethodPost,
			Body:           createInstrumentGroupBulkObjectBody,
			ExpectedStatus: http.StatusCreated,
			ExpectedSchema: arrSchema,
		},
		{
			Name:           "UpdateInstrumentGroup",
			URL:            fmt.Sprintf("/instrument_groups/%s", testInstrumentGroupID),
			Method:         http.MethodPut,
			Body:           updateInstrumentGroupBody,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: objSchema,
		},
		{
			Name:           "DeleteInstrumentGroup",
			URL:            fmt.Sprintf("/instrument_groups/%s", testInstrumentGroupID),
			Method:         http.MethodDelete,
			ExpectedStatus: http.StatusOK,
		},
		{
			Name:           "ListInstrumentGroupMeasurements",
			URL:            fmt.Sprintf("/instrument_groups/%s/timeseries_measurements?after=%s&before=%s&threshold=%s", testInstrumentGroupID, testInstrumentGroupTimeAfter, testInstrumentGroupTimeBefore, testInstrumentGroupThreshold),
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
		},
		{
			Name:           "Add Instrument to InstrumentGroup",
			URL:            fmt.Sprintf("/instrument_groups/%s/instruments", testInstrumentGroupID),
			Method:         http.MethodPost,
			Body:           addInstrumentToInstrumentGroupBody,
			ExpectedStatus: http.StatusCreated,
		},
		{
			Name:           "Remove Instrument from InstrumentGroup",
			URL:            fmt.Sprintf("/instrument_groups/%s/instruments/%s", testInstrumentGroupID, testInstrumentGroupInstrumentID),
			Method:         http.MethodDelete,
			ExpectedStatus: http.StatusOK,
		},
		{
			Name:           "ListProjectInstrumentGroups",
			URL:            fmt.Sprintf("/projects/%s/instrument_groups", testProjectID),
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: arrSchema,
		}}

	RunAll(t, tests)
}
