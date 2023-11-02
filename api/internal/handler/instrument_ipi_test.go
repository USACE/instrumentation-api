package handler_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xeipuuv/gojsonschema"
)

const ipiSegmentArraySchema = `{
    "type": "array",
    "properties": {
        "id": { "type": "string" },
        "instrument_id": { "type": "string" },
        "length": { "type": ["number", "null"] },
        "length_timeseries_id": { "type": "string" },
        "tilt_timeseries_id": { "type": ["string", "null"] },
        "cum_dev_timeseries_id": { "type": ["string", "null"] }
    },
    "additionalProperties": false
}`

var ipiSegmentArrayLoader = gojsonschema.NewStringLoader(ipiSegmentArraySchema)

const ipiMeasurementsArraySchema = `{
    "type": "array",
    "properties": {
        "segment_id": { "type": "number" },
        "instrument_id": { "type": "string" },
        "tilt": { "type": ["number", "null"] },
        "cum_dev": { "type": ["number", "null"] },
	"elevation": { "type": ["number", "null"] }
    },
    "additionalProperties": false
}`

var ipiMeasurementsArrayLoader = gojsonschema.NewStringLoader(ipiMeasurementsArraySchema)

const (
	testIpiInstrumentID = "01ac435f-fe3c-4af1-9979-f5e00467e7f5"
	testIpiTimeAfter    = "1900-01-01T00:00:00.00Z"
	testIpiTimeBefore   = "2030-01-01T00:00:00.00Z"
)

const updateIpiSegmentsBody = `[
    {
        "id": 2,
        "instrument_id": "01ac435f-fe3c-4af1-9979-f5e00467e7f5",
        "length": 1,
        "length_timeseries_id": "e891ca7c-59b2-41bc-9d4a-43995e35b855",
        "tilt_timeseries_id": null,
        "cum_dev_timeseries_id": null
    },
    {
        "id": 3,
        "instrument_id": "01ac435f-fe3c-4af1-9979-f5e00467e7f5",
        "length": 200,
        "length_timeseries_id": "18f17db2-4bc8-44cb-a9fa-ba84d13b8444",
        "tilt_timeseries_id": null,
        "cum_dev_timeseries_id": null
    }
]`

const createIpiInstrumentBulkObjectBody = `{
    "status_id": "94578354-ffdf-4119-9663-6bd4323e58f5",
    "status": "destroyed",
    "status_time": "2001-01-01T00:00:00Z",
    "slug": "test-ipi-1",
    "name": "Test IPI 1",
    "type_id": "c81f3a5d-fc5f-47fd-b545-401fe6ee63bb",
    "type": "IPI",
    "geometry": {
        "type": "Point",
        "coordinates": [
            -80.8,
            26.7
        ]
    },
    "formula": null,
    "station": null,
    "offset": null,
    "project_id": "5b6f4f37-7755-4cf9-bd02-94f1e9bc5984",
    "opts": {
	"num_segments": 10,
	"bottom_elevation": 1000
    }
}`

func TestIpiInstruments(t *testing.T) {
	segArrSchema, err := gojsonschema.NewSchema(ipiSegmentArrayLoader)
	assert.Nil(t, err)
	measurementsArrSchema, err := gojsonschema.NewSchema(ipiMeasurementsArrayLoader)
	assert.Nil(t, err)

	tests := []HTTPTest{
		{
			Name:           "GetAllIpiSegmentsForInstrument",
			URL:            fmt.Sprintf("/instruments/ipi/%s/segments", testIpiInstrumentID),
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: segArrSchema,
		},
		{
			Name:           "GetIpiMeasurementsForInstrument",
			URL:            fmt.Sprintf("/instruments/ipi/%s/measurements?after=%s&before=%s", testIpiInstrumentID, testIpiTimeAfter, testIpiTimeBefore),
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: measurementsArrSchema,
		},
		{
			Name:           "UpdateIpiSegments",
			URL:            fmt.Sprintf("/instruments/ipi/%s/segments", testIpiInstrumentID),
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: segArrSchema,
		},
		{
			Name:           "CreateIpiInstrumentBulk_Object",
			URL:            fmt.Sprintf("/projects/%s/instruments", testProjectID),
			Method:         http.MethodPost,
			Body:           createIpiInstrumentBulkObjectBody,
			ExpectedStatus: http.StatusCreated,
		}}

	RunAll(t, tests)
}
