package handler_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/xeipuuv/gojsonschema"
)

const districtSchema = `{
    "type": "object",
    "properties": {
        "id": { "type": "string" },
        "name": { "type": "string" },
        "initials": { "type": "string" },
        "division_name": { "type": "string" },
        "division_initials": { "type": "string" },
        "office_id": { "type": ["string", "null"] }
    },
    "additionalProperties": false
}`

var districtArraySchema = gojsonschema.NewStringLoader(fmt.Sprintf(`{
    "type": "array",
    "items": %s
}`, districtSchema))

const projectSchema = `{
    "type": "object",
    "properties": {
        "id": { "type": "string" },
        "federal_id": { "type": ["string", "null"] },
        "image": { "type": ["string", "null"]},
        "office_id": { "type": [ "string", "null" ]},
        "slug": { "type": "string" },
        "name": { "type": "string" },
        "creator": { "type": "string" },
        "create_date": { "type": "string", "format": "date-time" },
        "updater": {  "type": ["string", "null"] },
        "update_date": { "type": ["string", "null"], "format": "date-time" },
        "instrument_count": {"type": "number"},
        "instrument_group_count": {"type": "number"},
        "timeseries": {
            "type": "array",
            "items": { "type": "string" },
        },
    },
    "required": ["id", "federal_id", "image", "office_id", "slug", "name", "creator", "create_date", "updater", "update_date", "instrument_count", "instrument_group_count", "timeseries"],
    "additionalProperties": false
}`

var projectObjectSchema = gojsonschema.NewStringLoader(projectSchema)

var projectArraySchema = gojsonschema.NewStringLoader(fmt.Sprintf(`{
    "type": "array",
    "items": %s
}`, projectSchema))

const projectCountSchema = `{
    "type": "object",
    "properties": {
        "project_count": { "type": "number" }
    },
    "required": ["project_count"],
    "additionalProperties": false
}`

var projectCountObjectSchema = gojsonschema.NewStringLoader(projectCountSchema)

var projectInstrumentNamesArraySchema = gojsonschema.NewStringLoader(`{
    "type": "array",
    "items": { "type": "string" }
}`)

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
	tests := []HTTPTest{
		{
			Name:           "ListDistricts",
			URL:            "/districts",
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: &districtArraySchema,
		},
		{
			Name:           "GetProject",
			URL:            fmt.Sprintf("/projects/%s", testProjectID),
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: &projectObjectSchema,
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
			Name:           "GetProjectCount",
			URL:            "/projects",
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: &projectCountObjectSchema,
		},
		{
			Name:           "ListProjects",
			URL:            "/projects",
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: &projectArraySchema,
		},
		{
			Name:           "ListProjectsByFederalID",
			URL:            fmt.Sprintf("/projects?federal_id=%s", testProjectFederalID),
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: &projectArraySchema,
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
			Name:           "UpdateProject",
			URL:            fmt.Sprintf("/projects/%s", testProjectID),
			Method:         http.MethodPut,
			Body:           updateProjectBody,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: &projectObjectSchema,
		},
		{
			Name:           "DeleteProject",
			URL:            fmt.Sprintf("/projects/%s", testProjectID),
			Method:         http.MethodDelete,
			ExpectedStatus: http.StatusOK,
		},
		{
			Name:           "ListProjectInstrumentGroups",
			URL:            fmt.Sprintf("/projects/%s/instrument_groups", testProjectID),
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: &projectArraySchema,
		},
		{
			Name:           "ListProjectInstruments",
			URL:            fmt.Sprintf("/projects/%s/instruments", testProjectID),
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: &instrumentArraySchema,
		},
		{
			Name:           "ListProjectInstrumentNames",
			URL:            fmt.Sprintf("/projects/%s/instruments/names", testProjectID),
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: &projectInstrumentNamesArraySchema,
		}}

	RunAll(t, tests)
}
