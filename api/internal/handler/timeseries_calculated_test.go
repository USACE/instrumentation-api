package handler_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/xeipuuv/gojsonschema"
)

const (
	testFormulaID           = "5b6f4f37-7755-4cf9-bd02-94f1e9bc5984"
	testFormulaInstrumentID = "a7540f69-c41e-43b3-b655-6e44097edb7e"
)

const formulaSchema = `{
    "type": "array",
    "properties": {
        "id": { "type": "string" },
        "instrument_id": { "type": "string" },
        "parameter_id": { "type": "string" },
        "unit_id": { "type": "string" },
        "slug": { "type": "string" },
        "formula_name": { "type": "string" },
        "formula": { "type": "string" },
    },
    "required": [ "id", "instrument_id", "parameter_id", "unit_id", "slug", "formula_name", "formula" ]
}`

var formulaObjectSchema = gojsonschema.NewStringLoader(formulaSchema)

const formulaIDsSchema = `{
    "type": "object",
    "properties": {
        "id": { "type": "string" },
    },
    "required": ["id"]
}`

var formulaIDsObjectSchema = gojsonschema.NewStringLoader(formulaIDsSchema)

const createFormulaBody = `{
    "instrument_id": "a7540f69-c41e-43b3-b655-6e44097edb7e",
    "parameter_id": "068b59b0-aafb-4c98-ae4b-ed0365a6fbac",
    "unit_id": "f777f2e2-5e32-424e-a1ca-19d16cd8abce",
    "slug": "test_calculation",
    "formula_name": "test calculation",
    "formula": "a + b"
}`

const updateFormulaBody = `{
    "id": "5b6f4f37-7755-4cf9-bd02-94f1e9bc5984",
    "formula_name": "NEW NAME",
    "formula": "c + d"
}`

func TestFormulas(t *testing.T) {
	tests := []HTTPTest{
		{
			Name:           "GetInstrumentFormulas",
			URL:            fmt.Sprintf("/formulas?instrument_id=%s", testFormulaInstrumentID),
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: &formulaObjectSchema,
		},
		{
			Name:           "CreateFormula",
			URL:            "/formulas",
			Method:         http.MethodPost,
			Body:           createFormulaBody,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: &formulaIDsObjectSchema,
		},
		{
			Name:           "UpdateFormula",
			URL:            fmt.Sprintf("/formulas/%s", testFormulaID),
			Method:         http.MethodPut,
			Body:           updateFormulaBody,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: &formulaIDsObjectSchema,
		},
		{
			Name:           "DeleteFormula",
			URL:            fmt.Sprintf("/formulas/%s", testFormulaID),
			Method:         http.MethodDelete,
			ExpectedStatus: http.StatusOK,
		}}

	RunAll(t, tests)
}
