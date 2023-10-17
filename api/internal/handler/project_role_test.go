package handler_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/xeipuuv/gojsonschema"
)

const projectRoleSchema = `{
    "type": "object",
    "properties": {
        "id": { "type": "string" },
        "profile_id": { "type": "string" },
        "username": { "type": ["string"] },
        "email": { "type": ["string"] },
        "role_id": { "type": ["string"] },
        "role": { "type": ["string"] }
    },
    "required": ["id", "profile_id", "username", "email", "role_id", "role"],
    "additionalProperties": false
}`

var projectRoleObjectSchema = gojsonschema.NewStringLoader(projectRoleSchema)

var projectRoleArraySchema = gojsonschema.NewStringLoader(fmt.Sprintf(`{
    "type": "array",
    "items": %s
}`, projectRoleSchema))

const (
	testMemberID = "57329df6-9f7a-4dad-9383-4633b452efab"
	testRoleID   = "37f14863-8f3b-44ca-8deb-4b74ce8a8a69"
)

func TestProjectMembership(t *testing.T) {
	tests := []HTTPTest{
		{
			Name:           "ListProjectMembers",
			URL:            fmt.Sprintf("/projects/%s/members", testProjectID),
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: &projectRoleArraySchema,
		},
		{
			Name:           "RemoveProjectMemberRole",
			URL:            fmt.Sprintf("/projects/%s/members/%s/roles/%s", testProjectID, testMemberID, testRoleID),
			Method:         http.MethodDelete,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: nil,
		},
		{
			Name:           "AddProjectMemberRole",
			URL:            fmt.Sprintf("/projects/%s/members/%s/roles/%s", testProjectID, testMemberID, testRoleID),
			Method:         http.MethodPost,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: &projectRoleObjectSchema,
		}}

	RunAll(t, tests)
}
