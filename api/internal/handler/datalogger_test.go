package handler_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/xeipuuv/gojsonschema"
)

const dataloggerSchema = `{
    "type": "object",
    "properties": {
        "id": { "type": "string" },
    },
    "required": ["id"]
}`

var dataloggerObjectSchema = gojsonschema.NewStringLoader(dataloggerSchema)

// datalogger 1 for read-only tests since it's used with the mock datalogger service
const (
	testDataloggerID1 = "83a7345c-62d8-4e29-84db-c2e36f8bc40d"
	testDataloggerID2 = "c0b65315-f802-4ca5-a4dd-7e0cfcffd057"
)

const createDataloggerBody = `{
    "sn": "11111",
    "project_id": "5b6f4f37-7755-4cf9-bd02-94f1e9bc5984",
    "name": "Test Create Data Logger",
    "model_id": "6a10ef5f-b9d9-4fa0-8b1e-ea1bcc81748c"
}`

const updateDataloggerBody = `{
    "id": "c0b65315-f802-4ca5-a4dd-7e0cfcffd057",
    "name": "Updated name",
    "sn": "99999",
    "project_id": "5b6f4f37-7755-4cf9-bd02-94f1e9bc5984",
    "model": "CR1000X"
}`

func TestDatalogger(t *testing.T) {
	tests := []HTTPTest{
		{
			Name:           "CreateDataLogger",
			URL:            "/datalogger",
			Method:         http.MethodPost,
			Body:           createDataloggerBody,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: &dataloggerObjectSchema,
		},
		{
			Name:           "ListProjectDataLoggers",
			URL:            fmt.Sprintf("/dataloggers?project_id=%s", testProjectID),
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
		},
		{
			Name:           "GetDataLogger",
			URL:            fmt.Sprintf("/datalogger/%s", testDataloggerID1),
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
		},
		{
			Name:           "GetDataLoggerPreview",
			URL:            fmt.Sprintf("/datalogger/%s/preview", testDataloggerID1),
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
		},
		{
			Name:           "UpdateDataLogger",
			URL:            fmt.Sprintf("/datalogger/%s", testDataloggerID2),
			Method:         http.MethodPut,
			Body:           updateDataloggerBody,
			ExpectedStatus: http.StatusOK,
		},
		{
			Name:           "CycleDataLoggerKey",
			URL:            fmt.Sprintf("/datalogger/%s/key", testDataloggerID2),
			Method:         http.MethodPut,
			ExpectedStatus: http.StatusOK,
		},
		{
			Name:           "DeleteDataLogger",
			URL:            fmt.Sprintf("/datalogger/%s", testDataloggerID2),
			Method:         http.MethodDelete,
			ExpectedStatus: http.StatusOK,
		}}

	RunAll(t, tests)
}
