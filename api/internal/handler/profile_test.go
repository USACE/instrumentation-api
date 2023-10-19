package handler_test

import (
	"net/http"
	"testing"

	"github.com/USACE/instrumentation-api/api/internal/model"
)

const testCreateProfileBody = `{
    "username": "testuser",
    "email": "test.user@gmail.com"
}`

const mockJwtNewUser = `Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxIiwibmFtZSI6IlVzZXIuTmV3IiwiaWF0IjoxNTE2MjM5MDIyLCJleHAiOjIwMDAwMDAwMDAsInJvbGVzIjpbIlBVQkxJQy5VU0VSIl19._WR_s6AGyq2FwHA980M8XoFbhVInvgTqstauxUfcmYs`

func TestProfiles(t *testing.T) {
	tests := []HTTPTest[model.Profile]{
		{
			Name:           "CreateProfile",
			URL:            "/profiles",
			Method:         http.MethodPost,
			Body:           testCreateProfileBody,
			ExpectedStatus: http.StatusCreated,
			authHeader:     mockJwtNewUser,
		},
		{
			Name:                 "GetMyProfile",
			URL:                  "/my_profile",
			Method:               http.MethodGet,
			ExpectedStatus:       http.StatusOK,
			ExpectedResponseType: jsonObj,
		},
		{
			Name:           "CreateToken",
			URL:            "/my_tokens",
			Method:         http.MethodPost,
			ExpectedStatus: http.StatusCreated,
		}}

	RunAll(t, tests)
}
