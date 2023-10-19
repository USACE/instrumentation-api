package handler_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/USACE/instrumentation-api/api/internal/model"
)

const (
	testInstrumentID = "a7540f69-c41e-43b3-b655-6e44097edb7e"
)

const updateInstrumentBody = `{
    "id": "a7540f69-c41e-43b3-b655-6e44097edb7e",
    "status_id": "94578354-ffdf-4119-9663-6bd4323e58f5",
    "status": "destroyed",
    "status_time": "2001-01-01T00:00:00Z",
    "slug": "demo-piezometer-1",
    "name": "Demo Piezometer 1 Updated Name",
    "type_id": "1bb4bf7c-f5f8-44eb-9805-43b07ffadbef",
    "type": "Piezometer",
    "geometry": {
        "type": "Point",
        "coordinates": [
            -80.8,
            26.7
        ]
    },
    "station": null,
    "offset": null,
    "project_id": "5b6f4f37-7755-4cf9-bd02-94f1e9bc5984",
    "zreference": 44.5,
    "zreference_datum_id": "72113f9a-982d-44e5-8fc1-8e595dafd344",
    "zreference_datum": "North American Vertical Datum of 1988 (NAVD 88)",
    "zreference_time": "2006-06-01T00:00:00Z"
}`

const updateInstrumentGeometryBody = `{
    "type": "Point",
    "coordinates": [
        -78.0,
        25.0
    ]
}`

const createInstrumentBulkArrayBody = `[{
    "status_id": "94578354-ffdf-4119-9663-6bd4323e58f5",
    "status": "destroyed",
    "status_time": "2001-01-01T00:00:00Z",
    "slug": "demo-piezometer-2",
    "formula": null,
    "name": "Demo Piezometer 2",
    "type_id": "1bb4bf7c-f5f8-44eb-9805-43b07ffadbef",
    "type": "Piezometer",
    "geometry": {
        "type": "Point",
        "coordinates": [
            -80.8,
            26.7
        ]
    },
    "station": null,
    "offset": null,
    "project_id": "5b6f4f37-7755-4cf9-bd02-94f1e9bc5984"
},
{
    "status_id": "94578354-ffdf-4119-9663-6bd4323e58f5",
    "status": "destroyed",
    "status_time": "2001-01-01T00:00:00Z",
    "slug": "demo-piezometer-3",
    "name": "Demo Piezometer 3",
    "type_id": "1bb4bf7c-f5f8-44eb-9805-43b07ffadbef",
    "type": "Piezometer",
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
    "project_id": "5b6f4f37-7755-4cf9-bd02-94f1e9bc5984"
},
{
    "status_id": "94578354-ffdf-4119-9663-6bd4323e58f5",
    "status": "destroyed",
    "status_time": "2001-01-01T00:00:00Z",
    "slug": "demo-piezometer-4",
    "name": "Demo Piezometer 4",
    "type_id": "1bb4bf7c-f5f8-44eb-9805-43b07ffadbef",
    "type": "Piezometer",
    "geometry": {
        "type": "Point",
        "coordinates": [
            -80.8,
            26.7
        ]
    },
    "station": null,
    "offset": null,
    "project_id": "5b6f4f37-7755-4cf9-bd02-94f1e9bc5984"
}]`

const validateCreateInstrumentArrayBody = `[{
    "status_id": "94578354-ffdf-4119-9663-6bd4323e58f5",
    "status": "destroyed",
    "status_time": "2001-01-01T00:00:00Z",
    "slug": "demo-piezometer-2",
    "name": "Demo Piezometer 2",
    "type_id": "1bb4bf7c-f5f8-44eb-9805-43b07ffadbef",
    "type": "Piezometer",
    "geometry": {
        "type": "Point",
        "coordinates": [
            -80.8,
            26.7
        ]
    },
    "station": null,
    "offset": null,
    "project_id": "5b6f4f37-7755-4cf9-bd02-94f1e9bc5984",
    "zreference": 44.5,
    "zreference_datum_id": "72113f9a-982d-44e5-8fc1-8e595dafd344",
    "zreference_datum": "North American Vertical Datum of 1988 (NAVD 88)",
    "zreference_time": "2006-06-01T00:00:00Z"
},
{
    "status_id": "94578354-ffdf-4119-9663-6bd4323e58f5",
    "status": "destroyed",
    "status_time": "2001-01-01T00:00:00Z",
    "slug": "demo-piezometer-3",
    "name": "Demo Piezometer 3",
    "type_id": "1bb4bf7c-f5f8-44eb-9805-43b07ffadbef",
    "type": "Piezometer",
    "geometry": {
        "type": "Point",
        "coordinates": [
            -80.8,
            26.7
        ]
    },
    "station": null,
    "offset": null,
    "project_id": "5b6f4f37-7755-4cf9-bd02-94f1e9bc5984",
    "zreference": 44.5,
    "zreference_datum_id": "72113f9a-982d-44e5-8fc1-8e595dafd344",
    "zreference_datum": "North American Vertical Datum of 1988 (NAVD 88)",
    "zreference_time": "2006-06-01T00:00:00Z"
},
{
    "status_id": "94578354-ffdf-4119-9663-6bd4323e58f5",
    "status": "destroyed",
    "status_time": "2001-01-01T00:00:00Z",
    "slug": "demo-piezometer-4",
    "name": "Demo Piezometer 4",
    "type_id": "1bb4bf7c-f5f8-44eb-9805-43b07ffadbef",
    "type": "Piezometer",
    "geometry": {
        "type": "Point",
        "coordinates": [
            -80.8,
            26.7
        ]
    },
    "station": null,
    "offset": null,
    "project_id": "5b6f4f37-7755-4cf9-bd02-94f1e9bc5984",
    "zreference": 44.5,
    "zreference_datum_id": "72113f9a-982d-44e5-8fc1-8e595dafd344",
    "zreference_datum": "North American Vertical Datum of 1988 (NAVD 88)",
    "zreference_time": "2006-06-01T00:00:00Z"
}]`

const createInstrumentBulkObjectBody = `{
    "status_id": "94578354-ffdf-4119-9663-6bd4323e58f5",
    "status": "destroyed",
    "status_time": "2001-01-01T00:00:00Z",
    "slug": "demo-piezometer-5",
    "name": "Demo Piezometer 5",
    "type_id": "1bb4bf7c-f5f8-44eb-9805-43b07ffadbef",
    "type": "Piezometer",
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
    "project_id": "5b6f4f37-7755-4cf9-bd02-94f1e9bc5984"
}`

const validateCreateInstrumentBulkObjectBody = `{
    "status_id": "94578354-ffdf-4119-9663-6bd4323e58f5",
    "status": "destroyed",
    "status_time": "2001-01-01T00:00:00Z",
    "slug": "demo-piezometer-5",
    "name": "Demo Piezometer 5",
    "type_id": "1bb4bf7c-f5f8-44eb-9805-43b07ffadbef",
    "type": "Piezometer",
    "geometry": {
        "type": "Point",
        "coordinates": [
            -80.8,
            26.7
        ]
    },
    "station": null,
    "offset": null,
    "project_id": "5b6f4f37-7755-4cf9-bd02-94f1e9bc5984",
    "zreference": 44.5,
    "zreference_datum_id": "72113f9a-982d-44e5-8fc1-8e595dafd344",
    "zreference_datum": "North American Vertical Datum of 1988 (NAVD 88)",
    "zreference_time": "2006-06-01T00:00:00Z"
}`

func TestInstruments(t *testing.T) {
	countTest := []HTTPTest[model.InstrumentCount]{{
		Name:                 "GetInstrumentCount",
		URL:                  "/instruments/count",
		Method:               http.MethodGet,
		ExpectedStatus:       http.StatusOK,
		ExpectedResponseType: jsonObj,
	}}

	tests := []HTTPTest[model.Instrument]{
		{
			Name:                 "GetInstrument",
			URL:                  fmt.Sprintf("/instruments/%s", testInstrumentID),
			Method:               http.MethodGet,
			ExpectedStatus:       http.StatusOK,
			ExpectedResponseType: jsonObj,
		},
		{
			Name:                 "UpdateInstrument",
			URL:                  fmt.Sprintf("/projects/%s/instruments/%s", testProjectID, testInstrumentID),
			Method:               http.MethodPut,
			Body:                 updateInstrumentBody,
			ExpectedStatus:       http.StatusOK,
			ExpectedResponseType: jsonObj,
		},
		{
			Name:                 "UpdateInstrumentGeometry",
			URL:                  fmt.Sprintf("/projects/%s/instruments/%s/geometry", testProjectID, testInstrumentID),
			Method:               http.MethodPut,
			Body:                 updateInstrumentGeometryBody,
			ExpectedStatus:       http.StatusOK,
			ExpectedResponseType: jsonObj,
		},
		{
			Name:                 "ListInstruments",
			URL:                  "/instruments",
			Method:               http.MethodGet,
			ExpectedStatus:       http.StatusOK,
			ExpectedResponseType: jsonArr,
		},
		{
			Name:                 "ListProjectInstruments",
			URL:                  fmt.Sprintf("/projects/%s/instruments", testProjectID),
			Method:               http.MethodGet,
			ExpectedStatus:       http.StatusOK,
			ExpectedResponseType: jsonArr,
		},
		{
			Name:                 "ListInstrumentGroupInstruments",
			URL:                  fmt.Sprintf("/instrument_groups/%s/instruments", testInstrumentGroupID),
			Method:               http.MethodGet,
			ExpectedStatus:       http.StatusOK,
			ExpectedResponseType: jsonArr,
		},
		{
			Name:           "CreateInstrumentBulk_Array",
			URL:            fmt.Sprintf("/projects/%s/instruments", testProjectID),
			Method:         http.MethodPost,
			Body:           createInstrumentBulkArrayBody,
			ExpectedStatus: http.StatusCreated,
		},
		{
			Name:           "ValidateCreateInstrument_Array",
			URL:            fmt.Sprintf("/projects/%s/instruments?dry_run=true", testProjectID),
			Method:         http.MethodPost,
			Body:           validateCreateInstrumentArrayBody,
			ExpectedStatus: http.StatusOK,
		},
		{
			Name:           "CreateInstrumentBulk_Object",
			URL:            fmt.Sprintf("/projects/%s/instruments", testProjectID),
			Method:         http.MethodPost,
			Body:           createInstrumentBulkObjectBody,
			ExpectedStatus: http.StatusCreated,
		},
		{
			Name:           "ValidateCreateInstrument_Object",
			URL:            fmt.Sprintf("/projects/%s/instruments?dry_run=true", testProjectID),
			Method:         http.MethodPost,
			Body:           validateCreateInstrumentBulkObjectBody,
			ExpectedStatus: http.StatusOK,
		},
		{
			Name:           "DeleteInstrument",
			URL:            fmt.Sprintf("/projects/%s/instruments/%s", testProjectID, testInstrumentID),
			Method:         http.MethodDelete,
			ExpectedStatus: http.StatusOK,
		}}

	RunAll(t, countTest)
	RunAll(t, tests)
}
