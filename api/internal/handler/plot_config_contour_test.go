package handler_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xeipuuv/gojsonschema"
)

var plotConfigContourSchema = fmt.Sprintf(plotConfigBaseSchema, IDSlugNameArrSchema, plotConfigContourDisplaySchema)

const plotConfigContourDisplaySchema = `{
    "timeseries_ids": { "type": "array", "items": { "type": "string" } },
    "time": { "type": "string" },
    "locf_backfill": { "type": "string" },
    "gradient_smoothing": { "type": "boolean" },
    "contour_smoothing": { "type": "boolean" },
    "show_labels": { "type": "boolean" }
}`

const plotConfigContourTimesSchema = `{
    "type": "array",
    "items": { "type": "string", "format": "date-time" }
}`

const plotConfigContourMeasurementsSchema = `{
    "type": "object",
    "properties": {
	"x": { "type": "array", "items": { "type": "number" } },
        "y": { "type": "array", "items": { "type": "number" } },
        "z": { "type": "array", "items": { "type": "number" } }
    }
}`

var plotConfigContourObjectLoader = gojsonschema.NewStringLoader(plotConfigContourSchema)
var plotConfigContourTimesLoader = gojsonschema.NewStringLoader(plotConfigContourTimesSchema)
var plotConfigContourMeasurementsLoader = gojsonschema.NewStringLoader(plotConfigContourMeasurementsSchema)

const testPlotConfigContourID = "94df34f5-ba00-4c3d-bfa7-f128a00166be"

const updatePlotConfigContourBody = `{
    "id": "94df34f5-ba00-4c3d-bfa7-f128a00166be",
    "name": "Updated Contour Plot",
    "slug": "updated-contour-plot",
    "project_id": "5b6f4f37-7755-4cf9-bd02-94f1e9bc5984",
    "plot_type": "contour",
    "display": {
        "timeseries_ids": [
            "00ae950d-5bdd-455e-a72a-56da67dafb85",
            "5842c707-b4be-4d10-a89c-1064e282e555"
        ],
	"time": "2024-06-15T18:45:47+00:00",
        "locf_backfill": "P1D",
        "gradient_smoothing": true,
        "contour_smoothing": true,
        "show_labels": true
    }
}`

const createPlotConfigContourBody = `{
    "name": "New Contour Plot",
    "project_id": "5b6f4f37-7755-4cf9-bd02-94f1e9bc5984",
    "plot_type": "contour",
    "display": {
        "timeseries_ids": [
            "00ae950d-5bdd-455e-a72a-56da67dafb85",
            "5842c707-b4be-4d10-a89c-1064e282e555"
        ],
	"time": "2024-06-15T18:45:47+00:00",
        "locf_backfill": "P1D",
        "gradient_smoothing": true,
        "contour_smoothing": true,
        "show_labels": true
    }
}`

func TestPlotConfigsContour(t *testing.T) {
	objSchema, err := gojsonschema.NewSchema(plotConfigContourObjectLoader)
	assert.Nil(t, err)
	timesSchema, err := gojsonschema.NewSchema(plotConfigContourTimesLoader)
	assert.Nil(t, err)
	measurementsSchema, err := gojsonschema.NewSchema(plotConfigContourMeasurementsLoader)
	assert.Nil(t, err)

	tests := []HTTPTest{
		{
			Name:           "ListPlotConfigContourPlotTimes",
			URL:            fmt.Sprintf("/projects/%s/plot_configs/contour_plots/%s/times", testProjectID, testPlotConfigContourID),
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: timesSchema,
		},
		{
			Name:           "GetPlotConfigContourPlotMeasurements",
			URL:            fmt.Sprintf("/projects/%s/plot_configs/contour_plots/%s/measurements?time=2024-01-01T00:00:00Z", testProjectID, testPlotConfigContourID),
			Method:         http.MethodGet,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: measurementsSchema,
		},
		{
			Name:           "UpdatePlotConfigContourPlot",
			URL:            fmt.Sprintf("/projects/%s/plot_configs/contour_plots/%s", testProjectID, testPlotConfigContourID),
			Method:         http.MethodPut,
			Body:           updatePlotConfigContourBody,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: objSchema,
		},
		{
			Name:           "CreatePlotConfigContourPlot",
			URL:            fmt.Sprintf("/projects/%s/plot_configs/contour_plots", testProjectID),
			Method:         http.MethodPost,
			Body:           createPlotConfigContourBody,
			ExpectedStatus: http.StatusCreated,
			ExpectedSchema: objSchema,
		}}

	RunAll(t, tests)
}
