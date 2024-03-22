package handler_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xeipuuv/gojsonschema"
)

var plotConfigSchema = fmt.Sprintf(`{
    "type": "object",
    "properties": {
        "id": { "type": "string" },
        "timeseries_id": {
            "type": "array",
            "items": { "type": "string" }
        },
        "slug": { "type": "string" },
        "name": { "type": "string" },
        "creator_id": { "type": "string" },
        "create_date": { "type": "string", "format": "date-time" },
        "updater_id": {  "type": ["string", "null"] },
        "update_date": { "type": ["string", "null"], "format": "date-time" },
        "project_id": { "type": ["string", "null"] },
        "show_masked": { "type": "boolean" },
        "show_nonvalidated": { "type": "boolean" },
        "show_comments": { "type": "boolean" },
        "auto_range": { "type": "boolean" },
        "date_range": { "type": "string" },
        "threshold": { "type": "number" },
	"report_configs": %s
    },
    "required": ["id", "slug", "name", "creator_id", "create_date", "updater_id", "update_date", "project_id", "timeseries_id"],
    "additionalProperties": false
}`, IDSlugNameArrSchema)

var plotConfigObjectLoader = gojsonschema.NewStringLoader(plotConfigSchema)

var plotConfigArrayLoader = gojsonschema.NewStringLoader(fmt.Sprintf(`{
    "type": "array",
    "items": %s
}`, plotConfigSchema))

const testPlotConfigID = "64879f68-6a2c-4d78-8e8b-5e9b9d2e0d6a"

const updatePlotConfigRemoveTimeseriesBody = `{
    "id": "64879f68-6a2c-4d78-8e8b-5e9b9d2e0d6a",
    "name": "PZ-1A PLOT",
    "slug": "pz-1a-plot",
    "project_id": "5b6f4f37-7755-4cf9-bd02-94f1e9bc5984",
    "timeseries_id": [
        "9a3864a8-8766-4bfa-bad1-0328b166f6a8"
    ],
    "creator_id": "00000000-0000-0000-0000-000000000000",
    "create_date": "2021-02-26T15:54:18.982835Z",
    "updater_id": null,
    "update_date": null
}`

const updatePlotConfigAddManyTimeseriesBody = `{
    "id": "64879f68-6a2c-4d78-8e8b-5e9b9d2e0d6a",
    "name": "PZ-1A PLOT",
    "slug": "pz-1a-plot",
    "project_id": "5b6f4f37-7755-4cf9-bd02-94f1e9bc5984",
    "timeseries_id": [
        "8f4ca3a3-5971-4597-bd6f-332d1cf5af7c",
        "869465fc-dc1e-445e-81f4-9979b5fadda9",
        "9a3864a8-8766-4bfa-bad1-0328b166f6a8",
        "7ee902a3-56d0-4acf-8956-67ac82c03a96",
        "d9697351-3a38-4194-9ac4-41541927e475",
        "22a734d6-dc24-451d-a462-43a32f335ae8"
    ],
    "creator_id": "00000000-0000-0000-0000-000000000000",
    "create_date": "2021-02-26T16:21:07.925124Z",
    "updater_id": null,
    "update_date": null
}`

const createPlotConfigBody = `{
    "name": "Test Create Plot Configuration",
    "project_id": "5b6f4f37-7755-4cf9-bd02-94f1e9bc5984",
    "timeseries_id": [
        "8f4ca3a3-5971-4597-bd6f-332d1cf5af7c",
        "9a3864a8-8766-4bfa-bad1-0328b166f6a8",
        "7ee902a3-56d0-4acf-8956-67ac82c03a96",
        "d9697351-3a38-4194-9ac4-41541927e475",
        "22a734d6-dc24-451d-a462-43a32f335ae8"
    ]
}`

func TestPlotConfigurations(t *testing.T) {
	objSchema, err := gojsonschema.NewSchema(plotConfigObjectLoader)
	assert.Nil(t, err)
	arrSchema, err := gojsonschema.NewSchema(plotConfigArrayLoader)
	assert.Nil(t, err)

	tests := []HTTPTest{
		{
			Name:           "GetPlotConfiguration",
			URL:            fmt.Sprintf("/projects/%s/plot_configurations/%s", testProjectID, testPlotConfigID),
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: objSchema,
		},
		{
			Name:           "ListPlotConfigurations",
			URL:            fmt.Sprintf("/projects/%s/plot_configurations", testProjectID),
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: arrSchema,
		},
		{
			Name:           "UpdatePlotConfiguration - Add Many Timeseries",
			URL:            fmt.Sprintf("/projects/%s/plot_configurations/%s", testProjectID, testPlotConfigID),
			Method:         http.MethodPut,
			Body:           updatePlotConfigAddManyTimeseriesBody,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: objSchema,
		},
		{
			Name:           "UpdatePlotConfiguration - Remove Timeseries",
			URL:            fmt.Sprintf("/projects/%s/plot_configurations/%s", testProjectID, testPlotConfigID),
			Method:         http.MethodPut,
			Body:           updatePlotConfigRemoveTimeseriesBody,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: objSchema,
		},
		{
			Name:           "CreatePlotConfiguration",
			URL:            fmt.Sprintf("/projects/%s/plot_configurations", testProjectID),
			Method:         http.MethodPost,
			Body:           createPlotConfigBody,
			ExpectedStatus: http.StatusCreated,
			ExpectedSchema: objSchema,
		},
		{
			Name:           "DeletePlotConfiguration",
			URL:            fmt.Sprintf("/projects/%s/plot_configurations/%s", testProjectID, testPlotConfigID),
			Method:         http.MethodDelete,
			ExpectedStatus: http.StatusOK,
		}}

	RunAll(t, tests)
}
