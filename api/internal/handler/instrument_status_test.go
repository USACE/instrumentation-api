package handler_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/xeipuuv/gojsonschema"
)

const instrumentStatusSchema = `{
    "type": "object",
    "properties": {
        "id": { "type": "string" },
        "time": { "type": "string" },
        "status_id": { "type": "string" },
        "status": { "type": "string" }
    },
    "required": ["id", "time", "status_id", "status"],
    "additionalProperties": false
}`

var instrumentStatusObjectSchema = gojsonschema.NewStringLoader(instrumentStatusSchema)

var instrumentStatusArraySchema = gojsonschema.NewStringLoader(fmt.Sprintf(`{
    "type": "array",
    "items": %s
}`, instrumentStatusSchema))

const testInstrumentStatusID = "4ed5e9ac-40dc-4bca-b44f-7b837ec1b0fc"

const createInstrumentStatusArrayBody = `[{
    "time": "2002-01-01T00:00:00Z",
    "status_id": "03a2bf9a-bbd8-4031-8f4e-13e8c77807f1"
},
{
    "time": "2003-01-01T00:00:00Z",
    "status_id": "c9ee4acb-9623-4fde-bf36-7668afe463d4"
},
{
    "time": "2004-01-01T00:00:00Z",
    "status_id": "e26ba2ef-9b52-4c71-97df-9e4b6cf4174d"
},
{
    "time": "2005-01-01T00:00:00Z",
    "status_id": "94578354-ffdf-4119-9663-6bd4323e58f5"
}]`

const createInstrumentStatusObjectBody = `{
    "time": "2018-01-01T00:00:00Z",
    "status_id": "03a2bf9a-bbd8-4031-8f4e-13e8c77807f1"
}`

func TestInstrumentStatus(t *testing.T) {
	tests := []HTTPTest{
		{
			Name:           "GetInstrumentStatus",
			URL:            fmt.Sprintf("/instruments/%s/status/%s", testInstrumentID, testInstrumentStatusID),
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: &instrumentStatusObjectSchema,
		},
		{
			Name:           "ListInstrumentStatus",
			URL:            fmt.Sprintf("/instruments/%s/status", testInstrumentID),
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: &instrumentArraySchema,
		},
		{
			Name:           "CreateInstrumentStatus_Array",
			URL:            fmt.Sprintf("/instruments/%s/status", testInstrumentID),
			Method:         http.MethodPost,
			Body:           createInstrumentStatusArrayBody,
			ExpectedStatus: http.StatusCreated,
		},
		{
			Name:           "CreateInstrumentStatus_Object",
			URL:            fmt.Sprintf("/instruments/%s/status", testInstrumentID),
			Method:         http.MethodPost,
			Body:           createInstrumentStatusObjectBody,
			ExpectedStatus: http.StatusCreated,
		},
		{
			Name:           "DeleteInstrumentStatus",
			URL:            fmt.Sprintf("/instruments/%s/status/%s", testInstrumentID, testInstrumentStatusID),
			Method:         http.MethodDelete,
			ExpectedStatus: http.StatusOK,
		}}

	RunAll(t, tests)
}
