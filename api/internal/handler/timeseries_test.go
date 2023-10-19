package handler_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/USACE/instrumentation-api/api/internal/model"
)

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
	tests := []HTTPTest[model.Timeseries]{
		{
			Name:                 "GetTimeseries",
			URL:                  fmt.Sprintf("/timeseries/%s", testTimeseriesID),
			Method:               http.MethodGet,
			ExpectedStatus:       http.StatusOK,
			ExpectedResponseType: jsonObj,
		},
		{
			Name:                 "ListTimeseries",
			URL:                  "/timeseries",
			Method:               http.MethodGet,
			ExpectedStatus:       http.StatusOK,
			ExpectedResponseType: jsonArr,
		},
		{
			Name:                 "ListProjectTimeseries",
			URL:                  fmt.Sprintf("/projects/%s/timeseries", testProjectID),
			Method:               http.MethodGet,
			ExpectedStatus:       http.StatusOK,
			ExpectedResponseType: jsonArr,
		},
		{
			Name:                 "ListInstrumentGroupTimeseries",
			URL:                  fmt.Sprintf("/instrument_groups/%s/timeseries", testInstrumentGroupID),
			Method:               http.MethodGet,
			ExpectedStatus:       http.StatusOK,
			ExpectedResponseType: jsonArr,
		},
		{
			Name:                 "CreateTimeseries_Object",
			URL:                  "/timeseries",
			Method:               http.MethodPost,
			Body:                 createTimeseriesObjectBody,
			ExpectedStatus:       http.StatusCreated,
			ExpectedResponseType: jsonArr,
		},
		{
			Name:                 "CreateTimeseries_Array",
			URL:                  "/timeseries",
			Method:               http.MethodPost,
			Body:                 createTimeseriesArrayBody,
			ExpectedStatus:       http.StatusCreated,
			ExpectedResponseType: jsonArr,
		},
		{
			Name:                 "UpdateTimeseries",
			URL:                  fmt.Sprintf("/timeseries/%s", testTimeseriesID),
			Method:               http.MethodPut,
			Body:                 updateTimeseriesBody,
			ExpectedStatus:       http.StatusOK,
			ExpectedResponseType: jsonObj,
		},
		{
			Name:           "DeleteTimeseries",
			URL:            fmt.Sprintf("/timeseries/%s", testTimeseriesID),
			Method:         http.MethodDelete,
			ExpectedStatus: http.StatusOK,
		}}

	RunAll(t, tests)
}
