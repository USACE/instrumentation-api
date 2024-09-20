package handler_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xeipuuv/gojsonschema"
)

const measurementSchema = `{
    "type": "object",
    "properties": {
        "id": { "type": "string" },
        "time": { "type": "string" },
        "value": { "type": "number" }
    },
    "required": ["time", "value"],
    "additionalProperties": true
}`

var measurementCollectionSchema = fmt.Sprintf(`{
    "type": "object",
    "properties": {
        "timeseries_id": {"type": "string"},
        "items": { 
            "type": "array",
            "items": %s
        }
    },
    "required": ["items", "timeseries_id"],
    "additionalProperties": false
}`, measurementSchema)

var measurementObjectLoader = gojsonschema.NewStringLoader(measurementSchema)

var measurementCollectionObjectLoader = gojsonschema.NewStringLoader(measurementCollectionSchema)

var measurementCollectionArrayLoader = gojsonschema.NewStringLoader(fmt.Sprintf(`{
	"type": "array",
	"items": %s
}`, measurementCollectionSchema))

const (
	testMeasurementTimeAfter  = "1900-01-01T00:00:00.00Z"
	testMeasurementTimeBefore = "2021-01-01T00:00:00.00Z"
)

const createMeasurementsObjectBody = `{
    "timeseries_id": "869465fc-dc1e-445e-81f4-9979b5fadda9",
    "items": [
    	    {"time": "2020-06-01T00:00:00Z", "value": 10.00},
            {"time": "2020-06-02T01:00:00Z", "value": null},
            {"time": "2020-06-03T02:00:00Z", "value": 10.20},
            {"time": "2020-06-04T03:00:00Z", "value": 10.30},
            {"time": "2020-06-05T04:00:00Z", "value": 10.40}
    	]
}`

const createMeasurementsArrayBody = `[
    {
        "timeseries_id": "869465fc-dc1e-445e-81f4-9979b5fadda9",
        "items": [
		{"time": "2020-06-01T00:00:00Z", "value": 10.00},
	        {"time": "2020-06-02T01:00:00Z", "value": 11.10},
	        {"time": "2020-06-03T02:00:00Z", "value": 10.20},
	        {"time": "2020-06-04T03:00:00Z", "value": null},
	        {"time": "2020-06-05T04:00:00Z", "value": 10.40}
	]
    },
    {
        "timeseries_id": "9a3864a8-8766-4bfa-bad1-0328b166f6a8",
        "items": [
		{"time": "2020-06-01T00:00:00Z", "value": 10.00},
	        {"time": "2020-06-02T01:00:00Z", "value": null},
	        {"time": "2020-06-03T02:00:00Z", "value": 10.20},
	        {"time": "2020-06-04T03:00:00Z", "value": 10.30},
	        {"time": "2020-06-05T04:00:00Z", "value": 10.40}
	]
    },
    {
        "timeseries_id": "7ee902a3-56d0-4acf-8956-67ac82c03a96",
        "items": [
		{"time": "2020-06-01T00:00:00Z", "value": 10.00},
	        {"time": "2020-06-02T01:00:00Z", "value": 11.10},
	        {"time": "2020-06-03T02:00:00Z", "value": 10.20},
	        {"time": "2020-06-04T03:00:00Z", "value": 10.30},
	        {"time": "2020-06-05T04:00:00Z", "value": null}
	]
    }
]`

const updateMeasurementsBody = `{
    "timeseries_id": "869465fc-dc1e-445e-81f4-9979b5fadda9",
    "items": [
        {
            "time": "2020-01-23T00:00:00Z",
            "value": 130.05,
            "masked": true,
            "validated": true,
            "annotation": "test"
        }
    ]
}`

func TestMeasurements(t *testing.T) {
	objSchema, err := gojsonschema.NewSchema(measurementCollectionObjectLoader)
	assert.Nil(t, err)

	tests := []HTTPTest{
		{
			Name:           "ListTimeseriesMeasurements",
			URL:            fmt.Sprintf("/timeseries/%s/measurements?after=%s&before=%s", testTimeseriesID, testMeasurementTimeAfter, testMeasurementTimeBefore),
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: objSchema,
		},
		{
			Name:           "CreateTimeseriesMeasurements_Object",
			URL:            fmt.Sprintf("/projects/%s/timeseries_measurements", testProjectID),
			Method:         http.MethodPost,
			Body:           createMeasurementsObjectBody,
			ExpectedStatus: http.StatusCreated,
		},
		{
			Name:           "CreateTimeseriesMeasurements_Array",
			URL:            fmt.Sprintf("/projects/%s/timeseries_measurements", testProjectID),
			Method:         http.MethodPost,
			Body:           createMeasurementsArrayBody,
			ExpectedStatus: http.StatusCreated,
		},
		{
			Name:           "UpdateTimeseriesMeasurements",
			URL:            fmt.Sprintf("/projects/%s/timeseries_measurements?after=%s&before=%s", testProjectID, testMeasurementTimeAfter, testMeasurementTimeBefore),
			Method:         http.MethodPut,
			Body:           updateMeasurementsBody,
			ExpectedStatus: http.StatusOK,
		}}

	RunAll(t, tests)
}
