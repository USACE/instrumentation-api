package handler_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xeipuuv/gojsonschema"
)

const equivalencyTableRowSchema = `{
    "type": "object",
    "properties": {
	"id": { "type": "string" },
	"field_name": { "type": "string" },
	"display_name": { "type": "string" },
	"instrument_id": { "type": ["string", "null"] },
	"timeseries_id": { "type": ["string", "null"] }
    }
}`

var equivalencyTableSchema = fmt.Sprintf(`{
    "type": "object",
    "properties": {
        "datalogger_id" : { "type": "string" },
        "datalogger_table_id": { "type": "string" },
	"rows": { "type": "array", "items": %s }
    },
    "required": ["datalogger_id", "datalogger_table_id"]
}`, equivalencyTableRowSchema)

var equivalencyTableLoader = gojsonschema.NewStringLoader(equivalencyTableSchema)

const testEquivalencyTableRowID = "2f1f7c3d-8b6f-4b11-917e-8f049eb6c62b"

const createEquivalencyTableBody = `{
    "datalogger_id": "83a7345c-62d8-4e29-84db-c2e36f8bc40d",
    "datalogger_table_id": "98a77c65-e5c4-49ed-8fb4-b0ffd06add4c",
    "rows": [
        {
            "field_name": "new field name",
            "display_name": "test 123",
            "instrument_id": "a7540f69-c41e-43b3-b655-6e44097edb7e",
            "timeseries_id": "5985f20a-1e37-4add-823c-545cdca49b5e"
        }
    ]
}`

const updateEquivalencyTableBody = `{
    "datalogger_id": "83a7345c-62d8-4e29-84db-c2e36f8bc40d",
    "datalogger_table_id": "98a77c65-e5c4-49ed-8fb4-b0ffd06add4c",
    "rows": [
        {
            "id": "40ceff10-cdc3-4715-a4ca-c1e570fe25de",
            "field_name": "field name 1",
            "display_name": "test 1",
            "instrument_id": "a7540f69-c41e-43b3-b655-6e44097edb7e",
            "timeseries_id": "7ee902a3-56d0-4acf-8956-67ac82c03a96"
        },
        {
            "id": "2f1f7c3d-8b6f-4b11-917e-8f049eb6c62b",
            "field_name": "changed field name",
            "display_name": "changed display name",
            "instrument_id": "a7540f69-c41e-43b3-b655-6e44097edb7e",
            "timeseries_id": "869465fc-dc1e-445e-81f4-9979b5fadda9"
        }
    ]
}`

func TestEquivalencyTable(t *testing.T) {
	objSchema, err := gojsonschema.NewSchema(equivalencyTableLoader)
	assert.Nil(t, err)

	tests := []HTTPTest{
		{
			Name:           "CreateEquivalencyTable",
			URL:            fmt.Sprintf("/datalogger/%s/tables/%s/equivalency_table", testDataloggerID1, testDataloggerTableID),
			Method:         http.MethodPost,
			Body:           createEquivalencyTableBody,
			ExpectedStatus: http.StatusCreated,
			ExpectedSchema: objSchema,
		},
		{
			Name:           "GetEquivalencyTable",
			URL:            fmt.Sprintf("/datalogger/%s/tables/%s/equivalency_table", testDataloggerID1, testDataloggerTableID),
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: objSchema,
		},
		{
			Name:           "UpdateEquivalencyTable",
			URL:            fmt.Sprintf("/datalogger/%s/tables/%s/equivalency_table", testDataloggerID1, testDataloggerTableID),
			Method:         http.MethodPut,
			Body:           updateEquivalencyTableBody,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: objSchema,
		},
		{
			Name:           "DeleteEquivalencyTableRow",
			URL:            fmt.Sprintf("/datalogger/%s/tables/%s/equivalency_table/row/%s", testDataloggerID1, testDataloggerTableID, testEquivalencyTableRowID),
			Method:         http.MethodDelete,
			ExpectedStatus: http.StatusOK,
		},
		{
			Name:           "DeleteEquivalencyTable",
			URL:            fmt.Sprintf("/datalogger/%s/tables/%s/equivalency_table", testDataloggerID1, testDataloggerTableID),
			Method:         http.MethodDelete,
			ExpectedStatus: http.StatusOK,
		}}

	RunAll(t, tests)
}
