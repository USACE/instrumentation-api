package handler_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xeipuuv/gojsonschema"
)

const calculationArraySchema = `{
    "type": "array",
    "properties": {
        "id": { "type": "string" },
        "instrument_id": { "type": "string" },
        "parameter_id": { "type": "string" },
        "unit_id": { "type": "string" },
        "slug": { "type": "string" },
        "formula_name": { "type": "string" },
        "formula": { "type": "string" }
    },
    "required": ["id", "instrument_id", "parameter_id", "unit_id", "slug", "formula_name", "formula"]
}`

var calculationArrayLoader = gojsonschema.NewStringLoader(calculationArraySchema)

const calculationIDsObjectSchema = `{
    "type": "object",
    "properties": {
        "id": { "type": "string" }
    },
    "required": ["id"]
}`

var calculationIDsObjectLoader = gojsonschema.NewStringLoader(calculationIDsObjectSchema)

const (
	testCalculationID           = "5b6f4f37-7755-4cf9-bd02-94f1e9bc5984"
	testCalculationInstrumentID = "a7540f69-c41e-43b3-b655-6e44097edb7e"
)

const createCalculationBody = `{
    "instrument_id": "a7540f69-c41e-43b3-b655-6e44097edb7e",
    "parameter_id": "068b59b0-aafb-4c98-ae4b-ed0365a6fbac",
    "unit_id": "f777f2e2-5e32-424e-a1ca-19d16cd8abce",
    "slug": "test_calculation",
    "formula_name": "test calculation",
    "formula": "a + b"
}`

const updateCalculationBody = `{
    "id": "5b6f4f37-7755-4cf9-bd02-94f1e9bc5984",
    "formula_name": "NEW NAME",
    "formula": "c + d"
}`

func TestFormulas(t *testing.T) {
	objSchema, err := gojsonschema.NewSchema(calculationIDsObjectLoader)
	assert.Nil(t, err)
	arrSchema, err := gojsonschema.NewSchema(calculationArrayLoader)
	assert.Nil(t, err)

	tests := []HTTPTest{
		{
			Name:           "CreateCalculation",
			URL:            "/formulas",
			Method:         http.MethodPost,
			Body:           createCalculationBody,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: objSchema,
		},
		{
			Name:           "UpdateCalculation",
			URL:            fmt.Sprintf("/formulas/%s", testCalculationID),
			Method:         http.MethodPut,
			Body:           updateCalculationBody,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: objSchema,
		},
		{
			Name:           "GetInstrumentCalculations",
			URL:            fmt.Sprintf("/formulas?instrument_id=%s", testCalculationInstrumentID),
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: arrSchema,
		},
		{
			Name:           "DeleteCalculation",
			URL:            fmt.Sprintf("/formulas/%s", testCalculationID),
			Method:         http.MethodDelete,
			ExpectedStatus: http.StatusOK,
		}}

	RunAll(t, tests)
}
