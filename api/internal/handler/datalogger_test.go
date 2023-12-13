package handler_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xeipuuv/gojsonschema"
)

const dataloggerTableSchema = `{
    "type": "object",
    "properties": {
	"id": { "type": "string" },
        "name": { "type": "string" }
    }
}`

var dataloggerSchema = fmt.Sprintf(`{
    "type": "object",
    "properties": {
        "id": { "type": "string" },
	"name": { "type": "string" },
	"sn": { "type": "string" },
	"project_id": { "type": "string" },
	"creator": { "type": "string" },
	"creator_username": { "type": "string" },
	"create_date": { "type": "string" },
	"updater": { "type": ["string", "null"] },
	"updater_username": { "type": "string" },
	"update_date": { "type": ["string", "null"] },
	"slug": { "type": "string" },
	"model_id": { "type": "string" },
	"model": { "type": "string" },
	"errors": { "type": "array", "items": { "type": "string" } },
	"tables": { "type": "array", "items": %s },
	"key": { "type": "string" }
    },
    "required": [
        "id",
	"name",
	"sn",
	"project_id",
	"creator",
	"creator_username",
	"create_date",
	"slug",
	"model_id",
	"model",
	"errors",
	"tables"
    ]
}`, dataloggerTableSchema)

var dataloggerObjectLoader = gojsonschema.NewStringLoader(dataloggerSchema)
var dataloggerArrayLoader = gojsonschema.NewStringLoader(fmt.Sprintf(`{
    "type": "array",
    "items": %s
}`, dataloggerSchema))

const dataloggerPreviewSchema = `{
    "type": "object",
    "properties": {
        "datalogger_table_id": { "type": "string" },
	"update_date": { "type": "string" },
	"preview": { "type": ["object", "array", "null"] }
    }
}`

var dataloggerPreviewLoader = gojsonschema.NewStringLoader(dataloggerPreviewSchema)

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
	arrSchema, err := gojsonschema.NewSchema(dataloggerArrayLoader)
	assert.Nil(t, err)
	previewObjSchema, err := gojsonschema.NewSchema(dataloggerPreviewLoader)
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
			ExpectedSchema: arrSchema,
		},
		{
			Name:           "GetDatalogger",
			URL:            fmt.Sprintf("/datalogger/%s", testDataloggerID1),
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: objSchema,
		},
		{
			Name:           "GetDataloggerTablePreview",
			URL:            fmt.Sprintf("/datalogger/%s/tables/%s/preview", testDataloggerID1, testDataloggerTableID),
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: previewObjSchema,
		},
		{
			Name:           "ResetDataloggerTableName",
			URL:            fmt.Sprintf("/datalogger/%s/tables/%s/name", testDataloggerID1, testDataloggerTableID),
			Method:         http.MethodPut,
			ExpectedStatus: http.StatusOK,
		},
		{
			Name:           "UpdateDatalogger",
			URL:            fmt.Sprintf("/datalogger/%s", testDataloggerID2),
			Method:         http.MethodPut,
			Body:           updateDataloggerBody,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: objSchema,
		},
		{
			Name:           "CycleDataloggerKey",
			URL:            fmt.Sprintf("/datalogger/%s/key", testDataloggerID2),
			Method:         http.MethodPut,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: objSchema,
		},
		{
			Name:           "DeleteDatalogger",
			URL:            fmt.Sprintf("/datalogger/%s", testDataloggerID2),
			Method:         http.MethodDelete,
			ExpectedStatus: http.StatusOK,
		}}

	RunAll(t, tests)
}
