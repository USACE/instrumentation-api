package handler_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xeipuuv/gojsonschema"
)

const saaSegmentArraySchema = `{
    "type": "array",
    "properties": {
        "id": { "type": "string" },
        "instrument_id": { "type": "string" },
        "length": { "type": [ "number", "null" ] },
        "x_timeseries_id": { "type": [ "string", "null" ] },
        "y_timeseries_id": { "type": [ "string", "null" ] },
        "z_timeseries_id": { "type": [ "string", "null" ] },
        "temp_timeseries_id": { "type": [ "string", "null" ] }
    },
    "additionalProperties": false
}`

var saaSegmentArrayLoader = gojsonschema.NewStringLoader(saaSegmentArraySchema)

const saaMeasurementsArraySchema = `{
    "type": "array",
    "properties": {
        "segment_id": { "type": "number" },
        "instrument_id": { "type": "string" },
        "x": { "type": [ "number", "null" ] },
        "y": { "type": [ "number", "null" ] },
        "z": { "type": [ "number", "null" ] },
        "temp": { "type": [ "number", "null" ] },
        "x_increment": { "type": [ "number", "null" ] },
        "y_increment": { "type": [ "number", "null" ] },
        "z_increment": { "type": [ "number", "null" ] },
        "temp_increment": { "type": [ "number", "null" ] },
        "x_cum_dev": { "type": [ "number", "null" ] },
        "y_cum_dev": { "type": [ "number", "null" ] },
        "z_cum_dev": { "type": [ "number", "null" ] },
        "temp_cum_dev": { "type": [ "number", "null" ] }
    },
    "additionalProperties": false
}`

var saaMeasurementsArrayLoader = gojsonschema.NewStringLoader(saaMeasurementsArraySchema)

const (
	testSaaInstrumentID = "eca4040e-aecb-4cd3-bcde-3e308f0356a6"
	testSaaTimeAfter    = "1900-01-01T00:00:00.00Z"
	testSaaTimeBefore   = "2030-01-01T00:00:00.00Z"
)

const updateSaaSegmentsBody = `[
    {
        "id": 5,
        "instrument_id": "eca4040e-aecb-4cd3-bcde-3e308f0356a6",
        "length": 1,
        "x_timeseries_id": "eec831d1-56a5-47ef-85eb-02c7622d6cb8",
        "y_timeseries_id": "8b3762ef-a852-4edc-8e87-746a92eaac9d",
        "z_timeseries_id": "ecfa267b-339b-4bb8-b7ae-eda550257878",
        "temp_timeseries_id": "a31a24c4-aa8e-4e52-9895-43cdb69fe703"
    },
    {
        "id": 6,
        "instrument_id": "eca4040e-aecb-4cd3-bcde-3e308f0356a6",
        "length": 200,
        "x_timeseries_id": "23bda2f6-c479-48e0-a0c2-db48c3b08c3c",
        "y_timeseries_id": "eb25ab9f-af8b-4383-839a-7d24899e02c4",
        "z_timeseries_id": "8e641473-d7bf-433c-a24b-55fa065ca0c3",
        "temp_timeseries_id": "21cfe121-d29d-40a2-b04f-6be71ba479fe"
    }
]`

func TestSaaInstruments(t *testing.T) {
	segArrSchema, err := gojsonschema.NewSchema(saaSegmentArrayLoader)
	assert.Nil(t, err)
	measurementsArrSchema, err := gojsonschema.NewSchema(saaMeasurementsArrayLoader)
	assert.Nil(t, err)

	tests := []HTTPTest{
		{
			Name:           "GetAllSaaSegmentsForInstrument",
			URL:            fmt.Sprintf("/instruments/saa/%s/segments", testSaaInstrumentID),
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: segArrSchema,
		},
		{
			Name:           "GetSaaMeasurementsForInstrument",
			URL:            fmt.Sprintf("/instruments/saa/%s/measurements?after=%s&before=%s", testSaaInstrumentID, testSaaTimeAfter, testSaaTimeBefore),
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: measurementsArrSchema,
		},
		{
			Name:           "UpdateSaaSegments",
			URL:            fmt.Sprintf("/instruments/saa/%s/segments", testSaaInstrumentID),
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: segArrSchema,
		}}

	RunAll(t, tests)
}
