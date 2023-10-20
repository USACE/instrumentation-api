package handler_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xeipuuv/gojsonschema"
)

var unitArrayLoader = gojsonschema.NewStringLoader(`{
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
            "measure": { "type": ["string", "null"] }
        },
        "required": ["id", "name", "abbreviation", "unit_family_id", "unit_family", "measure_id", "measure"],
        "additionalProperties": false
    }
}`)

func TestUnits(t *testing.T) {
	arrSchema, err := gojsonschema.NewSchema(unitArrayLoader)
	assert.Nil(t, err)

	tests := []HTTPTest{{
		Name:           "ListUnits",
		URL:            "/units",
		Method:         http.MethodGet,
		ExpectedStatus: http.StatusOK,
		ExpectedSchema: arrSchema,
	}}

	RunAll(t, tests)
}
