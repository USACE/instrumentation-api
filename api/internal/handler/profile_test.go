package handler_test

import (
	"net/http"
	"testing"

	"github.com/xeipuuv/gojsonschema"
)

var profileObjectSchema = gojsonschema.NewStringLoader(`{
    "type": "object",
    "properties": {
        "id": { "type": "string" },
        "username": { "type": "string" },
        "email": { "type": "string" },
    },
    "required": ["id", "username", "email"],
    "additionalProperties": true
}`)

const testCreateProfileBody = `{
    "username": "testuser",
    "email": "test.user@gmail.com"
}`

func TestProfiles(t *testing.T) {
	tests := []HTTPTest{
		{
			Name:           "CreateProfile",
			URL:            "/profiles",
			Method:         http.MethodPost,
			Body:           testCreateProfileBody,
			ExpectedStatus: http.StatusOK,
		},
		{
			Name:           "GetMyProfile",
			URL:            "/my_profile",
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: &profileObjectSchema,
		},
		{
			Name:           "CreateToken",
			URL:            "/my_tokens",
			Method:         http.MethodPost,
			ExpectedStatus: http.StatusOK,
		}}

	RunAll(t, tests)
}
