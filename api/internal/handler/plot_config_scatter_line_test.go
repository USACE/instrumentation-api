package handler_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xeipuuv/gojsonschema"
)

const plotConfigBaseSchema = `{
    "type": "object",
    "properties": {
        "id": { "type": "string" },
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
        "report_configs": %s,
        "plot_type": { "type": "string" },
        "display": %s
    },
    "required": [
        "id", "slug", "name", "creator_id", "create_date", "updater_id", "update_date", "project_id",
        "show_masked", "show_nonvalidated", "show_comments", "auto_range", "date_range", "threshold", "report_configs", "plot_type", "display"
    ],
    "additionalProperties": false
}`

var plotConfigSchema = fmt.Sprintf(plotConfigBaseSchema, IDSlugNameArrSchema, plotConfigDisplaySchema)

var plotConfigDisplaySchema = fmt.Sprintf(`{
    "traces": %s,
    "layout": %s
}`, plotConfigTracesArrSchema, plotConfigLayoutSchema)

const plotConfigTracesArrSchema = `{
    "type": "array",
    "items": {
        "type": "object",
        "properties": {
            "timeseries_id": { "type": "string" },
            "name": { "type": "string" },
            "trace_order": { "type": "number" },
            "trace_type": { "type": "string" },
            "color": { "type": "string" },
            "line_style": { "type": "string" },
            "width": { "type": "number" },
            "show_marks": { "type": "boolean" },
            "y_axis": { "type": "string" }
        }
    }
}`

const plotConfigLayoutSchema = `{
    "type": "object",
    "properties": {
        "custom_shapes": {
            "type": "array",
            "items": {
                "type": "object",
                "properties": {
                    "enabled": "boolean",
                    "name": "string",
                    "data_point": "number",
                    "color": "string"
                }
            }
        },
        "y_axis_title": { "type": ["string", "null"] },
        "y2_axis_title": { "type": ["string", "null"] }
    }
}`

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
    "plot_type": "scatter-line",
    "display": {
        "traces": [
            {
                "timeseries_id": "9a3864a8-8766-4bfa-bad1-0328b166f6a8",
                "name": "update test trace 1",
                "trace_order": 0,
                "color": "#0066ff"
            }
        ],
        "layout": {
            "custom_shapes": [],
            "yaxis_title": "Custom Y Axis Title",
            "secondary_axis_title": "test second axis title"
        }
    }
}`

const updatePlotConfigAddManyTimeseriesBody = `{
    "id": "64879f68-6a2c-4d78-8e8b-5e9b9d2e0d6a",
    "name": "PZ-1A PLOT",
    "slug": "pz-1a-plot",
    "project_id": "5b6f4f37-7755-4cf9-bd02-94f1e9bc5984",
    "plot_type": "scatter-line",
    "display": {
        "traces": [
            {
                "timeseries_id": "9a3864a8-8766-4bfa-bad1-0328b166f6a8",
                "name": "update test trace 1",
                "trace_order": 0,
                "color": "#0066ff"
            },
            {
                "timeseries_id": "8f4ca3a3-5971-4597-bd6f-332d1cf5af7c",
                "name": "update test trace 2",
                "trace_order": 1,
                "color": "#ff0000"
            },
            {
                "timeseries_id": "869465fc-dc1e-445e-81f4-9979b5fadda9",
                "name": "update test trace 3",
                "trace_order": 2,
                "color": "#ffaa00"
            },
            {
                "timeseries_id": "7ee902a3-56d0-4acf-8956-67ac82c03a96",
                "name": "update test trace 4",
                "trace_order": 3,
                "color": "#0000ff"
            },
            {
                "timeseries_id": "d9697351-3a38-4194-9ac4-41541927e475",
                "name": "update test trace 5",
                "trace_order": 4,
                "color": "#00ff00"
            },
            {
                "timeseries_id": "22a734d6-dc24-451d-a462-43a32f335ae8",
                "name": "update test trace 6",
                "trace_order": 5,
                "color": "#aa00aa"
            }
        ],
        "layout": {
            "custom_shapes": [],
            "yaxis_title": "Custom Y Axis Title",
            "secondary_axis_title": "test second axis title"
        }
    }
}`

const createPlotConfigBody = `{
    "name": "Test Create Plot Config",
    "project_id": "5b6f4f37-7755-4cf9-bd02-94f1e9bc5984",
    "plot_type": "scatter-line",
    "display": {
        "traces": [
            {
                "timeseries_id": "9a3864a8-8766-4bfa-bad1-0328b166f6a8",
                "name": "update test trace 1",
                "trace_order": 0,
                "color": "#0066ff"
            },
            {
                "timeseries_id": "8f4ca3a3-5971-4597-bd6f-332d1cf5af7c",
                "name": "update test trace 2",
                "trace_order": 1,
                "color": "#ff0000"
            },
            {
                "timeseries_id": "869465fc-dc1e-445e-81f4-9979b5fadda9",
                "name": "update test trace 3",
                "trace_order": 2,
                "color": "#ffaa00"
            },
            {
                "timeseries_id": "7ee902a3-56d0-4acf-8956-67ac82c03a96",
                "name": "update test trace 4",
                "trace_order": 3,
                "color": "#0000ff"
            },
            {
                "timeseries_id": "d9697351-3a38-4194-9ac4-41541927e475",
                "name": "update test trace 5",
                "trace_order": 4,
                "color": "#00ff00"
            },
            {
                "timeseries_id": "22a734d6-dc24-451d-a462-43a32f335ae8",
                "name": "update test trace 6",
                "trace_order": 5,
                "color": "#aa00aa"
            }
        ],
        "layout": {
            "custom_shapes": [
                {
                    "enabled": true,
                    "name": "test custom shape",
                    "data_point": 123,
                    "color": "#123abc"
                }
            ],
            "yaxis_title": "New Custom Y Axis Title",
            "secondary_axis_title": "test second axis title"
        }
    }
}`

func TestPlotConfigs(t *testing.T) {
	objSchema, err := gojsonschema.NewSchema(plotConfigObjectLoader)
	assert.Nil(t, err)
	arrSchema, err := gojsonschema.NewSchema(plotConfigArrayLoader)
	assert.Nil(t, err)

	tests := []HTTPTest{
		{
			Name:           "GetPlotConfig",
			URL:            fmt.Sprintf("/projects/%s/plot_configs/%s", testProjectID, testPlotConfigID),
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: objSchema,
		},
		{
			Name:           "ListPlotConfigs",
			URL:            fmt.Sprintf("/projects/%s/plot_configs", testProjectID),
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: arrSchema,
		},
		{
			Name:           "UpdatePlotConfigScatterLinePlot - Add Many Timeseries",
			URL:            fmt.Sprintf("/projects/%s/plot_configs/scatter_line_plots/%s", testProjectID, testPlotConfigID),
			Method:         http.MethodPut,
			Body:           updatePlotConfigAddManyTimeseriesBody,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: objSchema,
		},
		{
			Name:           "UpdatePlotConfigScatterLinePlot - Remove Timeseries",
			URL:            fmt.Sprintf("/projects/%s/plot_configs/scatter_line_plots/%s", testProjectID, testPlotConfigID),
			Method:         http.MethodPut,
			Body:           updatePlotConfigRemoveTimeseriesBody,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: objSchema,
		},
		{
			Name:           "CreatePlotConfigScatterLinePlot",
			URL:            fmt.Sprintf("/projects/%s/plot_configs/scatter_line_plots", testProjectID),
			Method:         http.MethodPost,
			Body:           createPlotConfigBody,
			ExpectedStatus: http.StatusCreated,
			ExpectedSchema: objSchema,
		},
		{
			Name:           "DeletePlotConfig",
			URL:            fmt.Sprintf("/projects/%s/plot_configs/%s", testProjectID, testPlotConfigID),
			Method:         http.MethodDelete,
			ExpectedStatus: http.StatusOK,
		}}

	RunAll(t, tests)
}
