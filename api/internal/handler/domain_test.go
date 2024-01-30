package handler_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xeipuuv/gojsonschema"
)

var domainArrayLoader = gojsonschema.NewStringLoader(`{
    "type": "array",
    "items": {
        "type": "object",
        "properties": {
            "id": { "type": "string" },
            "group": { "type": "string" },
            "value": { "type": "string" },
            "description": { "type": ["string", "null"] }
        },
        "required": ["id", "group", "value", "description"],
        "additionalProperties": false
    }
}`)

func TestDomains(t *testing.T) {
	arrSchema, err := gojsonschema.NewSchema(domainArrayLoader)
	assert.Nil(t, err)

	tests := []HTTPTest{
		{
			Name:           "GetDomains",
			URL:            "/domains",
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: arrSchema,
		},
		{
			Name:           "GetDomainMap",
			URL:            "/domains/map",
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
		},
	}

	RunAll(t, tests)
}
