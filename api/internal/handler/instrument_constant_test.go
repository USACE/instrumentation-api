package handler_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xeipuuv/gojsonschema"
)

const (
	testInstrumentConstantInstrumentID  = "a7540f69-c41e-43b3-b655-6e44097edb7e"
	testInstrumentConstantTimeseriesID1 = "22a734d6-dc24-451d-a462-43a32f335ae8"
	testInstrumentConstantTimeseriesID2 = "14247bc8-b264-4857-836f-182d47ebb39d"
)

const createInstrumentConstantBody = `{
    "name": "Test Instrument Constant",
    "instrument_id": "a7540f69-c41e-43b3-b655-6e44097edb7e",
    "parameter_id": "068b59b0-aafb-4c98-ae4b-ed0365a6fbac",
    "unit_id": "f777f2e2-5e32-424e-a1ca-19d16cd8abce"
}`

const updateInstrumentConstantBody = `{
    "id": "22a734d6-dc24-451d-a462-43a32f335ae8",
    "name": "Tip Depth Updated Name",
    "instrument_id": "a7540f69-c41e-43b3-b655-6e44097edb7e",
    "parameter_id": "068b59b0-aafb-4c98-ae4b-ed0365a6fbac",
    "unit_id": "f777f2e2-5e32-424e-a1ca-19d16cd8abce"
}`

func TestInstrumentConstants(t *testing.T) {
	objSchema, err := gojsonschema.NewSchema(timeseriesObjectLoader)
	assert.Nil(t, err)
	arrSchema, err := gojsonschema.NewSchema(timeseriesArrayLoader)
	assert.Nil(t, err)

	tests := []HTTPTest{
		{
			Name:           "ListInstrumentConstants",
			URL:            fmt.Sprintf("/projects/%s/instruments/%s/constants", testProjectID, testInstrumentConstantInstrumentID),
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: arrSchema,
		},
		{
			Name:           "CreateInstrumentConstant",
			URL:            fmt.Sprintf("/projects/%s/instruments/%s/constants", testProjectID, testInstrumentConstantInstrumentID),
			Method:         http.MethodPost,
			Body:           createInstrumentConstantBody,
			ExpectedStatus: http.StatusCreated,
			ExpectedSchema: arrSchema,
		},
		{
			Name:           "UpdateInstrumentConstant",
			URL:            fmt.Sprintf("/projects/%s/instruments/%s/constants/%s", testProjectID, testInstrumentConstantInstrumentID, testInstrumentConstantTimeseriesID1),
			Method:         http.MethodPut,
			Body:           updateInstrumentConstantBody,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: objSchema},
		{
			Name:           "DeleteInstrumentConstant",
			URL:            fmt.Sprintf("/projects/%s/instruments/%s/constants/%s", testProjectID, testInstrumentConstantInstrumentID, testInstrumentConstantTimeseriesID2),
			Method:         http.MethodDelete,
			ExpectedStatus: http.StatusOK,
		}}

	RunAll(t, tests)
}
