package handler_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/USACE/instrumentation-api/api/internal/model"
)

const (
	testProjectID           = "5b6f4f37-7755-4cf9-bd02-94f1e9bc5984"
	testProjectTimeseriesID = "8f4ca3a3-5971-4597-bd6f-332d1cf5af7c"
	testProjectFederalID    = "NIST001"
)

const createProjectBulkArrayBody = `[{
    "name": "Test Project 100000",
    "federal_id": null
},
{
    "name": "Test Project 100001",
    "federal_id": null
},
{
    "name": "Test Project 100002",
    "federal_id": null
}]`

const createProjectBulkObjectBody = `{
    "name": "Test Project 500000",
    "federal_id": null
}`

const updateProjectBody = `{
    "id": "5b6f4f37-7755-4cf9-bd02-94f1e9bc5984",
    "federal_id": null,
    "name": "Blue Water Reservoir"
}`

func TestProjects(t *testing.T) {

	districtTest := []HTTPTest[model.District]{{
		Name:                 "ListDistricts",
		URL:                  "/districts",
		Method:               http.MethodGet,
		ExpectedStatus:       http.StatusOK,
		ExpectedResponseType: jsonArr,
	}}
	RunAll(t, districtTest)

	countTest := []HTTPTest[model.ProjectCount]{{
		Name:                 "GetProjectCount",
		URL:                  "/projects/count",
		Method:               http.MethodGet,
		ExpectedStatus:       http.StatusOK,
		ExpectedResponseType: jsonObj,
	}}
	RunAll(t, countTest)

	namesTest := []HTTPTest[string]{{
		Name:                 "ListProjectInstrumentNames",
		URL:                  fmt.Sprintf("/projects/%s/instruments/names", testProjectID),
		Method:               http.MethodGet,
		ExpectedStatus:       http.StatusOK,
		ExpectedResponseType: jsonArr,
	}}
	RunAll(t, namesTest)

	tests := []HTTPTest[model.Project]{
		{
			Name:                 "GetProject",
			URL:                  fmt.Sprintf("/projects/%s", testProjectID),
			Method:               http.MethodGet,
			ExpectedStatus:       http.StatusOK,
			ExpectedResponseType: jsonObj,
		},
		{
			Name:           "CreateProjectTimeseries",
			URL:            fmt.Sprintf("/projects/%s/timeseries/%s", testProjectID, testProjectTimeseriesID),
			Method:         http.MethodPost,
			ExpectedStatus: http.StatusCreated,
		},
		{
			Name:           "DeleteProjectTimeseries",
			URL:            fmt.Sprintf("/projects/%s/timeseries/%s", testProjectID, testProjectTimeseriesID),
			Method:         http.MethodDelete,
			ExpectedStatus: http.StatusOK,
		},
		{
			Name:                 "ListProjects",
			URL:                  "/projects",
			Method:               http.MethodGet,
			ExpectedStatus:       http.StatusOK,
			ExpectedResponseType: jsonArr,
		},
		{
			Name:                 "ListProjectsByFederalID",
			URL:                  fmt.Sprintf("/projects?federal_id=%s", testProjectFederalID),
			Method:               http.MethodGet,
			ExpectedStatus:       http.StatusOK,
			ExpectedResponseType: jsonArr,
		},
		{
			Name:           "CreateProjectBulk_Array",
			URL:            "/projects",
			Method:         http.MethodPost,
			Body:           createProjectBulkArrayBody,
			ExpectedStatus: http.StatusCreated,
		},
		{
			Name:           "CreateProjectBulk_Object",
			URL:            "/projects",
			Method:         http.MethodPost,
			Body:           createProjectBulkObjectBody,
			ExpectedStatus: http.StatusCreated,
		},
		{
			Name:                 "UpdateProject",
			URL:                  fmt.Sprintf("/projects/%s", testProjectID),
			Method:               http.MethodPut,
			Body:                 updateProjectBody,
			ExpectedStatus:       http.StatusOK,
			ExpectedResponseType: jsonObj,
		},
		{
			Name:           "DeleteProject",
			URL:            fmt.Sprintf("/projects/%s", testProjectID),
			Method:         http.MethodDelete,
			ExpectedStatus: http.StatusOK,
		},
		{
			Name:                 "ListProjectInstrumentGroups",
			URL:                  fmt.Sprintf("/projects/%s/instrument_groups", testProjectID),
			Method:               http.MethodGet,
			ExpectedStatus:       http.StatusOK,
			ExpectedResponseType: jsonArr,
		},
	}
	RunAll(t, tests)
}
