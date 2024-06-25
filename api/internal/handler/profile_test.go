package handler_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xeipuuv/gojsonschema"
)

var profileObjectLoader = gojsonschema.NewStringLoader(`{
    "type": "object",
    "properties": {
        "id": { "type": "string" },
        "username": { "type": "string" },
        "email": { "type": "string" }
    },
    "required": ["id", "username", "email"],
    "additionalProperties": true
}`)

const testCreateProfileBody = `{
    "username": "testuser",
    "email": "test.user@fake.usace.army.mil"
}`

const mockJwtNewUser = `Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxIiwibmFtZSI6Ik5ld1VzZXIiLCJnaXZlbl9uYW1lIjoiTmV3IFVzZXIiLCJwcmVmZXJyZWRfbmFtZSI6Ik5ldyBVc2VyIiwiY2FjVUlEIjoiMSIsIng1MDlfcHJlc2VudGVkIjp0cnVlLCJpYXQiOjE1MTYyMzkwMjIsImV4cCI6MjAwMDAwMDAwMCwicm9sZXMiOltdfQ.ElWDNEZu7EVMKzm7DaZctRXgJmLZy8658AOAteaY2Cs`

func TestProfiles(t *testing.T) {
	objSchema, err := gojsonschema.NewSchema(profileObjectLoader)
	assert.Nil(t, err)

	tests := []HTTPTest{
		{
			Name:           "CreateProfile",
			URL:            "/profiles",
			Method:         http.MethodPost,
			Body:           testCreateProfileBody,
			ExpectedStatus: http.StatusCreated,
			authHeader:     mockJwtNewUser,
		},
		{
			Name:           "GetMyProfile",
			URL:            "/my_profile",
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: objSchema,
		},
		{
			Name:           "CreateToken",
			URL:            "/my_tokens",
			Method:         http.MethodPost,
			ExpectedStatus: http.StatusCreated,
		}}

	RunAll(t, tests)
}
