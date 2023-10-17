package handler_test

import (
	"net/http"
	"testing"

	"github.com/xeipuuv/gojsonschema"
)

var unitArraySchema = gojsonschema.NewStringLoader(`{
    "type": "array",
    "items": {
        "type": "object",
        "properties": {
            "id": { "type": "string" },
            "name": { "type": "string" },
            "abbreviation": { "type": "string" },
            "unit_family_id": { "type": "string" },
            "unit_family": { "type": "string" },
            "measure_id": { "type": "string" },
            "measure": { "type": ["string", "null"] },
        },
        "required": ["id", "name", "abbreviation", "unit_family_id", "unit_family", "measure_id", "measure"],
        "additionalProperties": false
    }
}`)

func TestUnits(t *testing.T) {
	tests := []HTTPTest{{
		Name:           "ListUnits",
		URL:            "/units",
		Method:         http.MethodGet,
		ExpectedStatus: http.StatusOK,
		ExpectedSchema: &unitArraySchema,
	}}

	RunAll(t, tests)
}
