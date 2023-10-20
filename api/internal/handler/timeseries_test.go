package handler_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xeipuuv/gojsonschema"
)

const timeseriesSchema = `{
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
        "is_computed": { "type": "boolean" }
    },
    "required": ["id", "slug", "name", "variable", "instrument_id", "parameter_id", "unit_id", "is_computed"],
    "additionalProperties": false
}`

var timeseriesObjectLoader = gojsonschema.NewStringLoader(timeseriesSchema)

var timeseriesArrayLoader = gojsonschema.NewStringLoader(fmt.Sprintf(`{
    "type": "array",
    "items": %s
}`, timeseriesSchema))

const (
	testTimeseriesID = "869465fc-dc1e-445e-81f4-9979b5fadda9"
)

const createTimeseriesObjectBody = `{
    "name": "Test Timeseries 4",
    "instrument_id": "a7540f69-c41e-43b3-b655-6e44097edb7e",
    "parameter_id": "068b59b0-aafb-4c98-ae4b-ed0365a6fbac",
    "unit_id": "f777f2e2-5e32-424e-a1ca-19d16cd8abce"
}`

const createTimeseriesArrayBody = `[{
    "name": "Test Timeseries 5",
    "instrument_id": "a7540f69-c41e-43b3-b655-6e44097edb7e",
    "parameter_id": "068b59b0-aafb-4c98-ae4b-ed0365a6fbac",
    "unit_id": "f777f2e2-5e32-424e-a1ca-19d16cd8abce"
},
{
    "name": "Test Timeseries 6",
    "instrument_id": "a7540f69-c41e-43b3-b655-6e44097edb7e",
    "parameter_id": "068b59b0-aafb-4c98-ae4b-ed0365a6fbac",
    "unit_id": "f777f2e2-5e32-424e-a1ca-19d16cd8abce"
},
{
    "name": "Test Timeseries 7",
    "instrument_id": "a7540f69-c41e-43b3-b655-6e44097edb7e",
    "parameter_id": "068b59b0-aafb-4c98-ae4b-ed0365a6fbac",
    "unit_id": "f777f2e2-5e32-424e-a1ca-19d16cd8abce"
}]`

const updateTimeseriesBody = `{
    "id": "869465fc-dc1e-445e-81f4-9979b5fadda9",
    "slug": "test-timeseries-1",
    "name": "New Name for Test Timeseries 1",
    "instrument_id": "a7540f69-c41e-43b3-b655-6e44097edb7e",
    "instrument": "Demo Piezometer 1",
    "parameter_id": "068b59b0-aafb-4c98-ae4b-ed0365a6fbac",
    "parameter": "stage",
    "unit_id": "f777f2e2-5e32-424e-a1ca-19d16cd8abce",
    "unit": "feet"
}`

func TestTimeseries(t *testing.T) {
	objSchema, err := gojsonschema.NewSchema(timeseriesObjectLoader)
	assert.Nil(t, err)
	arrSchema, err := gojsonschema.NewSchema(timeseriesArrayLoader)
	assert.Nil(t, err)

	tests := []HTTPTest{
		{
			Name:           "GetTimeseries",
			URL:            fmt.Sprintf("/timeseries/%s", testTimeseriesID),
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: objSchema,
		},
		{
			Name:           "ListTimeseries",
			URL:            "/timeseries",
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: arrSchema,
		},
		{
			Name:           "ListProjectTimeseries",
			URL:            fmt.Sprintf("/projects/%s/timeseries", testProjectID),
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: arrSchema,
		},
		{
			Name:           "ListInstrumentGroupTimeseries",
			URL:            fmt.Sprintf("/instrument_groups/%s/timeseries", testInstrumentGroupID),
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: arrSchema,
		},
		{
			Name:           "CreateTimeseries_Object",
			URL:            "/timeseries",
			Method:         http.MethodPost,
			Body:           createTimeseriesObjectBody,
			ExpectedStatus: http.StatusCreated,
			ExpectedSchema: arrSchema,
		},
		{
			Name:           "CreateTimeseries_Array",
			URL:            "/timeseries",
			Method:         http.MethodPost,
			Body:           createTimeseriesArrayBody,
			ExpectedStatus: http.StatusCreated,
			ExpectedSchema: arrSchema,
		},
		{
			Name:           "UpdateTimeseries",
			URL:            fmt.Sprintf("/timeseries/%s", testTimeseriesID),
			Method:         http.MethodPut,
			Body:           updateTimeseriesBody,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: objSchema,
		},
		{
			Name:           "DeleteTimeseries",
			URL:            fmt.Sprintf("/timeseries/%s", testTimeseriesID),
			Method:         http.MethodDelete,
			ExpectedStatus: http.StatusOK,
		}}

	RunAll(t, tests)
}
