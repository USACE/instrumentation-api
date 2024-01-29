package handler_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xeipuuv/gojsonschema"
)

const districtSchema = `{
    "type": "object",
    "properties": {
	"agency": { "type": "string" },
        "id": { "type": "string" },
        "name": { "type": "string" },
        "initials": { "type": "string" },
        "division_name": { "type": "string" },
        "division_initials": { "type": "string" },
        "office_id": { "type": ["string", "null"] }
    },
    "additionalProperties": false
}`

var districtArrayLoader = gojsonschema.NewStringLoader(fmt.Sprintf(`{
    "type": "array",
    "items": %s
}`, districtSchema))

const projectSchema = `{
    "type": "object",
    "properties": {
        "id": { "type": "string" },
        "federal_id": { "type": ["string", "null"] },
        "image": { "type": ["string", "null"] },
        "office_id": { "type": [ "string", "null"] },
        "district_id": { "type": [ "string", "null"] },
        "slug": { "type": "string" },
        "name": { "type": "string" },
        "creator_id": { "type": "string" },
        "creator_username": { "type": "string" },
        "create_date": { "type": "string", "format": "date-time" },
        "updater_id": {  "type": ["string", "null"] },
        "updater_username": {  "type": ["string", "null"] },
	"update_date": { "type": ["string", "null"], "format": "date-time" },
        "instrument_count": { "type": "number" },
        "instrument_group_count": { "type": "number" }
    },
    "required": ["id", "federal_id", "image", "office_id", "slug", "name", "creator_id", "create_date", "updater_id", "update_date", "instrument_count", "instrument_group_count"],
    "additionalProperties": false
}`

var projectObjectLoader = gojsonschema.NewStringLoader(projectSchema)

var projectArrayLoader = gojsonschema.NewStringLoader(fmt.Sprintf(`{
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

var projectCountObjectLoader = gojsonschema.NewStringLoader(projectCountSchema)

const (
	testProjectID           = "5b6f4f37-7755-4cf9-bd02-94f1e9bc5984"
	testProjectTimeseriesID = "8f4ca3a3-5971-4597-bd6f-332d1cf5af7c"
	testProjectFederalID    = "NIST001"
)

const createProjectBulkBody = `[{
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

const updateProjectBody = `{
    "id": "5b6f4f37-7755-4cf9-bd02-94f1e9bc5984",
    "federal_id": null,
    "name": "Blue Water Reservoir"
}`

func TestProjects(t *testing.T) {
	districtArrSchema, err := gojsonschema.NewSchema(districtArrayLoader)
	assert.Nil(t, err)
	countObjSchema, err := gojsonschema.NewSchema(projectCountObjectLoader)
	assert.Nil(t, err)
	objSchema, err := gojsonschema.NewSchema(projectObjectLoader)
	assert.Nil(t, err)
	arrSchema, err := gojsonschema.NewSchema(projectArrayLoader)
	assert.Nil(t, err)

	tests := []HTTPTest{
		{
			Name:           "ListDistricts",
			URL:            "/districts",
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: districtArrSchema,
		},
		{
			Name:           "GetProjectCount",
			URL:            "/projects/count",
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: countObjSchema,
		},
		{
			Name:           "GetProject",
			URL:            fmt.Sprintf("/projects/%s", testProjectID),
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: objSchema,
		},
		{
			Name:           "ListProjects",
			URL:            "/projects",
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: arrSchema,
		},
		{
			Name:           "ListProjectsByFederalID",
			URL:            fmt.Sprintf("/projects?federal_id=%s", testProjectFederalID),
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: arrSchema,
		},
		{
			Name:           "ListProjectsForProfileRole",
			URL:            fmt.Sprintf("/my_projects?role=%s", "admin"),
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: arrSchema,
		},
		{
			Name:           "CreateProjectBulk",
			URL:            "/projects",
			Method:         http.MethodPost,
			Body:           createProjectBulkBody,
			ExpectedStatus: http.StatusCreated,
		},
		{
			Name:           "UpdateProject",
			URL:            fmt.Sprintf("/projects/%s", testProjectID),
			Method:         http.MethodPut,
			Body:           updateProjectBody,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: objSchema,
		},
		{
			Name:           "DeleteProject",
			URL:            fmt.Sprintf("/projects/%s", testProjectID),
			Method:         http.MethodDelete,
			ExpectedStatus: http.StatusOK,
		},
	}
	RunAll(t, tests)
}
