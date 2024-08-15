package handler_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xeipuuv/gojsonschema"
)

const instrumentValidationSchema = `{
    "type": "object",
    "properties": {
	"is_valid": { "type": "boolean" },
	"errors": { "type": "array", "items": { "type": "string" } }
    },
    "required": ["is_valid", "errors"],
    "additionalProperties": false
}`

var instrumentValidationObjectLoader = gojsonschema.NewStringLoader(instrumentValidationSchema)

const (
	testAssignProjectID = "d559abfd-7ec7-4d0d-97bd-a04018f01e4c"
)

const updateProjectInstrumentAssignmentsPayload = `{
    "instrument_ids": ["` + testInstrumentID + `"]
}`

const updateInstrumentProjectAssignmentsPayload = `{
    "project_ids": ["` + testAssignProjectID + `"]
}`

func TestInstrumentAssignments(t *testing.T) {
	objSchema, err := gojsonschema.NewSchema(instrumentValidationObjectLoader)
	assert.Nil(t, err)

	tests := []HTTPTest{
		{
			Name:           "AssignInstrumentToProject",
			URL:            fmt.Sprintf("/projects/%s/instruments/%s/assignments", testAssignProjectID, testInstrumentID),
			Method:         http.MethodPost,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: objSchema,
		},
		{
			Name:           "UnassignInstrumentFromProject",
			URL:            fmt.Sprintf("/projects/%s/instruments/%s/assignments", testAssignProjectID, testInstrumentID),
			Method:         http.MethodDelete,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: objSchema,
		},
		{
			Name:           "UpdateProjectInstrumentAssignments - Assign",
			URL:            fmt.Sprintf("/projects/%s/instruments/assignments?action=assign", testInstrumentID),
			Method:         http.MethodPut,
			Body:           updateProjectInstrumentAssignmentsPayload,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: objSchema,
		},
		{
			Name:           "UpdateProjectInstrumentAssignments - Unassign",
			URL:            fmt.Sprintf("/projects/%s/instruments/assignments?action=unassign", testInstrumentID),
			Method:         http.MethodPut,
			Body:           updateProjectInstrumentAssignmentsPayload,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: objSchema,
		},
		{
			Name:           "UpdateInstrumentProjectAssignments - Assign",
			URL:            fmt.Sprintf("/projects/%s/instruments/%s/assignments?action=assign", testProjectID, testInstrumentID),
			Method:         http.MethodPut,
			Body:           updateInstrumentProjectAssignmentsPayload,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: objSchema,
		},
		{
			Name:           "UpdateInstrumentProjectAssignments - Unassign",
			URL:            fmt.Sprintf("/projects/%s/instruments/%s/assignments?action=unassign", testProjectID, testInstrumentID),
			Method:         http.MethodPut,
			Body:           updateInstrumentProjectAssignmentsPayload,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: objSchema,
		}}

	RunAll(t, tests)
}
