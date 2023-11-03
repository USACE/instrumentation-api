package handler_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xeipuuv/gojsonschema"
)

const dataloggerSchema = `{
    "type": "object",
    "properties": {
        "id": { "type": "string" }
    },
    "required": ["id"]
}`

var dataloggerObjectLoader = gojsonschema.NewStringLoader(dataloggerSchema)

// datalogger 1 for read-only tests since it's used with the mock datalogger service
const (
	testDataloggerID1     = "83a7345c-62d8-4e29-84db-c2e36f8bc40d"
	testDataloggerID2     = "c0b65315-f802-4ca5-a4dd-7e0cfcffd057"
	testDataloggerTableID = "98a77c65-e5c4-49ed-8fb4-b0ffd06add4c"
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
	objSchema, err := gojsonschema.NewSchema(dataloggerObjectLoader)
	assert.Nil(t, err)

	tests := []HTTPTest{
		{
			Name:           "CreateDatalogger",
			URL:            "/datalogger",
			Method:         http.MethodPost,
			Body:           createDataloggerBody,
			ExpectedStatus: http.StatusCreated,
			ExpectedSchema: objSchema,
		},
		{
			Name:           "ListProjectDataloggers",
			URL:            fmt.Sprintf("/dataloggers?project_id=%s", testProjectID),
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
		},
		{
			Name:           "GetDatalogger",
			URL:            fmt.Sprintf("/datalogger/%s", testDataloggerID1),
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
		},
		{
			Name:           "GetDataloggerTablePreview",
			URL:            fmt.Sprintf("/datalogger/%s/table/%s/preview ", testDataloggerID1, testDataloggerTableID),
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
		},
		{
			Name:           "ResetDataloggerTableName",
			URL:            fmt.Sprintf("/datalogger/%s/table/%s/name ", testDataloggerID1, testDataloggerTableID),
			Method:         http.MethodPut,
			ExpectedStatus: http.StatusOK,
		},
		{
			Name:           "UpdateDatalogger",
			URL:            fmt.Sprintf("/datalogger/%s", testDataloggerID2),
			Method:         http.MethodPut,
			Body:           updateDataloggerBody,
			ExpectedStatus: http.StatusOK,
		},
		{
			Name:           "CycleDataloggerKey",
			URL:            fmt.Sprintf("/datalogger/%s/key", testDataloggerID2),
			Method:         http.MethodPut,
			ExpectedStatus: http.StatusOK,
		},
		{
			Name:           "DeleteDatalogger",
			URL:            fmt.Sprintf("/datalogger/%s", testDataloggerID2),
			Method:         http.MethodDelete,
			ExpectedStatus: http.StatusOK,
		}}

	RunAll(t, tests)
}