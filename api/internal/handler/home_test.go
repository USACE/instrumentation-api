package handler_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xeipuuv/gojsonschema"
)

var homeObjectLoader = gojsonschema.NewStringLoader(`{
    "type": "object",
    "properties": {
        "instrument_count": { "type": "number" },
        "instrument_group_count": { "type": "number" },
        "project_count": { "type": "number" },
        "new_instruments_7d": { "type": "number" },
        "new_measurements_2h": { "type": "number" }
    },
    "required": ["instrument_count", "instrument_group_count", "project_count", "new_instruments_7d", "new_measurements_2h"],
    "additionalProperties": false
}`)

func TestHome(t *testing.T) {
	objSchema, err := gojsonschema.NewSchema(homeObjectLoader)
	assert.Nil(t, err)

	tests := []HTTPTest{{
		Name:           "GetHome",
		URL:            "/home",
		Method:         http.MethodGet,
		ExpectedStatus: http.StatusOK,
		ExpectedSchema: objSchema,
	}}

	RunAll(t, tests)
}
