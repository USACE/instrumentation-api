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

const mockJwtNewCacUser = `Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6Ikw0YXFVRmd6YV9RVjhqc1ZOa281OW5GVzl6bGh1b0JGX3RxdlpkTUZkajQifQ.eyJzdWIiOiJmOGRjYWZlYS0yNDNlLTRiODktOGQ3ZC1mYTAxOTE4MTMwZjUiLCJ0eXAiOiJCZWFyZXIiLCJhbGxvd2VkLW9yaWdpbnMiOlsiaHR0cDovL2xvY2FsaG9zdDozMDAwIl0sIm5hbWUiOiJOZXcgVXNlciIsInByZWZlcnJlZF91c2VybmFtZSI6InRlc3QgbmV3IHVzZXIiLCJnaXZlbl9uYW1lIjoiTmV3IiwiZmFtaWx5X25hbWUiOiJVc2VyIiwiZW1haWwiOiJuZXcubS51c2VyQGZha2UudXNhY2UuYXJteS5taWwiLCJzdWJqZWN0RE4iOiJ1c2VyLm5ldy5tLjEiLCJjYWNVSUQiOiIxIiwieDUwOV9wcmVzZW50ZWQiOiJ0cnVlIn0.r-ABZpji9UIWdXIZmW1daxclJcH7qghXnBkacOX0g6U`

func TestProfiles(t *testing.T) {
	objSchema, err := gojsonschema.NewSchema(profileObjectLoader)
	assert.Nil(t, err)

	tests := []HTTPTest{
		{
			Name:           "CreateCacProfile",
			URL:            "/profiles",
			Method:         http.MethodPost,
			Body:           testCreateProfileBody,
			ExpectedStatus: http.StatusCreated,
			authHeader:     mockJwtNewCacUser,
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
