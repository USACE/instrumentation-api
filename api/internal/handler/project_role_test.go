package handler_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/USACE/instrumentation-api/api/internal/model"
)

const (
	testMemberID = "57329df6-9f7a-4dad-9383-4633b452efab"
	testRoleID   = "37f14863-8f3b-44ca-8deb-4b74ce8a8a69"
)

func TestProjectMembership(t *testing.T) {
	tests := []HTTPTest[model.ProjectMembership]{
		{
			Name:                 "ListProjectMembers",
			URL:                  fmt.Sprintf("/projects/%s/members", testProjectID),
			Method:               http.MethodGet,
			ExpectedStatus:       http.StatusOK,
			ExpectedResponseType: jsonArr,
		},
		{
			Name:           "RemoveProjectMemberRole",
			URL:            fmt.Sprintf("/projects/%s/members/%s/roles/%s", testProjectID, testMemberID, testRoleID),
			Method:         http.MethodDelete,
			ExpectedStatus: http.StatusOK,
		},
		{
			Name:                 "AddProjectMemberRole",
			URL:                  fmt.Sprintf("/projects/%s/members/%s/roles/%s", testProjectID, testMemberID, testRoleID),
			Method:               http.MethodPost,
			ExpectedStatus:       http.StatusOK,
			ExpectedResponseType: jsonObj,
		}}

	RunAll(t, tests)
}
