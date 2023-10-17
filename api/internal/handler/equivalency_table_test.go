package handler_test

import (
	"fmt"
	"net/http"
	"testing"
)

const testEquivalencyTableRowID = "2f1f7c3d-8b6f-4b11-917e-8f049eb6c62b"

const createEquivalencyTableBody = `{
    "datalogger_id": "83a7345c-62d8-4e29-84db-c2e36f8bc40d",
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
            "timeseries_id": "d9697351-3a38-4194-9ac4-41541927e475"
        }
    ]
}`

func TestEquivalencyTable(t *testing.T) {
	tests := []HTTPTest{
		{
			Name:           "CreateEquivalencyTable",
			URL:            fmt.Sprintf("/datalogger/%s/equivalency_table", testDataloggerID1),
			Method:         http.MethodPost,
			Body:           createEquivalencyTableBody,
			ExpectedStatus: http.StatusCreated,
		},
		{
			Name:           "GetEquivalencyTable",
			URL:            fmt.Sprintf("/datalogger/%s/equivalency_table", testDataloggerID1),
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
		},
		{
			Name:           "UpdateEquivalencyTable",
			URL:            fmt.Sprintf("/datalogger/%s/equivalency_table", testDataloggerID1),
			Method:         http.MethodPut,
			Body:           updateEquivalencyTableBody,
			ExpectedStatus: http.StatusOK,
		},
		{
			Name:           "DeleteEquivalencyTableRow",
			URL:            fmt.Sprintf("/datalogger/%s/equivalency_table/row?id=%s", testDataloggerID1, testEquivalencyTableRowID),
			Method:         http.MethodDelete,
			ExpectedStatus: http.StatusOK,
		},
		{
			Name:           "DeleteEquivalencyTable",
			URL:            fmt.Sprintf("/datalogger/%s/equivalency_table", testDataloggerID1),
			Method:         http.MethodDelete,
			ExpectedStatus: http.StatusOK,
		}}

	RunAll(t, tests)
}
