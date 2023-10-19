package handler_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
)

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
	idTests := []HTTPTest[struct{ id uuid.UUID }]{
		{
			Name:                 "CreateCalculation",
			URL:                  "/formulas",
			Method:               http.MethodPost,
			Body:                 createCalculationBody,
			ExpectedStatus:       http.StatusOK,
			ExpectedResponseType: jsonObj,
		},
		{
			Name:                 "UpdateCalculation",
			URL:                  fmt.Sprintf("/formulas/%s", testCalculationID),
			Method:               http.MethodPut,
			Body:                 updateCalculationBody,
			ExpectedStatus:       http.StatusOK,
			ExpectedResponseType: jsonObj,
		},
	}
	RunAll(t, idTests)

	tests := []HTTPTest[model.CalculatedTimeseries]{
		{
			Name:                 "GetInstrumentCalculations",
			URL:                  fmt.Sprintf("/formulas?instrument_id=%s", testCalculationInstrumentID),
			Method:               http.MethodGet,
			ExpectedStatus:       http.StatusOK,
			ExpectedResponseType: jsonArr,
		},
		{
			Name:           "DeleteCalculation",
			URL:            fmt.Sprintf("/formulas/%s", testCalculationID),
			Method:         http.MethodDelete,
			ExpectedStatus: http.StatusOK,
		}}

	RunAll(t, tests)
}
