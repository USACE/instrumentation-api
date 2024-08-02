package handler_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xeipuuv/gojsonschema"
)

const timeseriesCwmsSchema = `{
    "type": "object",
    "properties": {
        "id": { "type": "string" },
        "slug": { "type": "string" },
        "name": { "type": "string" },
        "variable": { "type": "string" },
        "instrument_id": { "type": "string" },
        "instrument": { "type": "string"  },
        "instrument_slug": { "type": "string" },
        "parameter_id": { "type": "string" },
        "parameter": { "type": "string"  },
        "unit_id": { "type": "string" },
        "unit": { "type": "string"  },
        "is_computed": { "type": "boolean" },
        "type": { "type": "string" },
	"cwms_timeseries_id": { "type": "string" },
	"cwms_office_id": { "type": "string" }
    },
    "required": ["id", "slug", "name", "variable", "instrument_id", "parameter_id", "unit_id", "is_computed", "type", "cwms_timeseries_id", "cwms_office_id"],
    "additionalProperties": false
}`

var timeseriesCwmsArrayLoader = gojsonschema.NewStringLoader(fmt.Sprintf(`{
    "type": "array",
    "items": %s
}`, timeseriesCwmsSchema))

const (
	testTimeseriesCwmsID = "47afea78-4169-499c-be51-013ca3b53cba"
)

const createTimeseriesCwmsArrayBody = `[{
    "name": "Test CWMS Timeseries 5",
    "instrument_id": "a7540f69-c41e-43b3-b655-6e44097edb7e",
    "parameter_id": "068b59b0-aafb-4c98-ae4b-ed0365a6fbac",
    "unit_id": "f777f2e2-5e32-424e-a1ca-19d16cd8abce",
    "cwms_timeseries_id": "test timeseries",
    "cwms_office_id": "test office"
},
{
    "name": "Test CWMS Timeseries 6",
    "instrument_id": "a7540f69-c41e-43b3-b655-6e44097edb7e",
    "parameter_id": "068b59b0-aafb-4c98-ae4b-ed0365a6fbac",
    "unit_id": "f777f2e2-5e32-424e-a1ca-19d16cd8abce",
    "cwms_timeseries_id": "test timeseries",
    "cwms_office_id": "test office"
},
{
    "name": "Test CWMS Timeseries 7",
    "instrument_id": "a7540f69-c41e-43b3-b655-6e44097edb7e",
    "parameter_id": "068b59b0-aafb-4c98-ae4b-ed0365a6fbac",
    "unit_id": "f777f2e2-5e32-424e-a1ca-19d16cd8abce",
    "cwms_timeseries_id": "test timeseries",
    "cwms_office_id": "test office"
}]`

const updateTimeseriesCwmsBody = `{
    "id": "47afea78-4169-499c-be51-013ca3b53cba",
    "slug": "test-timeseries-cwms-1",
    "name": "Updated Name for Test CWMS Timeseries 1",
    "instrument_id": "a7540f69-c41e-43b3-b655-6e44097edb7e",
    "instrument": "Demo Piezometer 1",
    "parameter_id": "068b59b0-aafb-4c98-ae4b-ed0365a6fbac",
    "parameter": "stage",
    "unit_id": "f777f2e2-5e32-424e-a1ca-19d16cd8abce",
    "unit": "feet",
    "cwms_timeseries_id": "test timeseries",
    "cwms_office_id": "test office"
}`

func TestTimeseriesCwms(t *testing.T) {
	arrSchema, err := gojsonschema.NewSchema(timeseriesCwmsArrayLoader)
	assert.Nil(t, err)

	tests := []HTTPTest{
		{
			Name:           "CreateTimeseriesCwms_Array",
			URL:            fmt.Sprintf("/projects/%s/instruments/%s/timeseries/cwms", testProjectID, testInstrumentID),
			Method:         http.MethodPost,
			Body:           createTimeseriesCwmsArrayBody,
			ExpectedStatus: http.StatusCreated,
			ExpectedSchema: arrSchema,
		},
		{
			Name:           "UpdateTimeseries",
			URL:            fmt.Sprintf("/projects/%s/instruments/%s/timeseries/cwms/%s", testProjectID, testInstrumentID, testTimeseriesCwmsID),
			Method:         http.MethodPut,
			Body:           updateTimeseriesCwmsBody,
			ExpectedStatus: http.StatusOK,
		},
		{
			Name:           "ListTimeseriesCwms",
			URL:            fmt.Sprintf("/projects/%s/instruments/%s/timeseries/cwms", testProjectID, testInstrumentID),
			Method:         http.MethodGet,
			Body:           createTimeseriesCwmsArrayBody,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: arrSchema,
		}}

	RunAll(t, tests)
}
