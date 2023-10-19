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
    "email": "test.user@gmail.com"
}`

const mockJwtNewUser = `Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxIiwibmFtZSI6IlVzZXIuTmV3IiwiaWF0IjoxNTE2MjM5MDIyLCJleHAiOjIwMDAwMDAwMDAsInJvbGVzIjpbIlBVQkxJQy5VU0VSIl19._WR_s6AGyq2FwHA980M8XoFbhVInvgTqstauxUfcmYs`

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
