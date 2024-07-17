package handler_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xeipuuv/gojsonschema"
)

var plotConfigBullseyeSchema = fmt.Sprintf(plotConfigBaseSchema, IDSlugNameArrSchema, plotConfigBullseyeDisplaySchema)

const plotConfigBullseyeDisplaySchema = `{
    "x_axis_timeseries_id": { "type": "string" },
    "y_axis_timeseries_id": { "type": "string" }
}`

var plotConfigBullseyeObjectLoader = gojsonschema.NewStringLoader(plotConfigBullseyeSchema)

const testPlotConfigBullseyeID = "871e34da-c911-4d8f-ab68-e29ba17f8937"

const updatePlotConfigBullseyeBody = `{
    "id": "871e34da-c911-4d8f-ab68-e29ba17f8937",
    "name": "Updated Bullseye Plot",
    "slug": "updated-bullseye-plot",
    "project_id": "5b6f4f37-7755-4cf9-bd02-94f1e9bc5984",
    "plot_type": "bullseye",
    "display": {
        "x_axis_timeseries_id": "4759bdac-656e-47c3-b403-d3118cf57342",
        "y_axis_timeseries_id": "2598aa5f-cb8f-4ab7-8ebf-6de0c30bce70"
    }
}`

const createPlotConfigBullseyeBody = `{
    "name": "New Bullseye Plot",
    "project_id": "5b6f4f37-7755-4cf9-bd02-94f1e9bc5984",
    "plot_type": "bullseye",
    "display": {
        "x_axis_timeseries_id": "4759bdac-656e-47c3-b403-d3118cf57342",
        "y_axis_timeseries_id": "2598aa5f-cb8f-4ab7-8ebf-6de0c30bce70"
    }
}`

func TestPlotConfigsBullseye(t *testing.T) {
	objSchema, err := gojsonschema.NewSchema(plotConfigBullseyeObjectLoader)
	assert.Nil(t, err)

	tests := []HTTPTest{
		{
			Name:           "UpdatePlotConfigBullseyePlot",
			URL:            fmt.Sprintf("/projects/%s/plot_configs/bullseye_plots/%s", testProjectID, testPlotConfigBullseyeID),
			Method:         http.MethodPut,
			Body:           updatePlotConfigBullseyeBody,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: objSchema,
		},
		{
			Name:           "CreatePlotConfigBullseyePlot",
			URL:            fmt.Sprintf("/projects/%s/plot_configs/bullseye_plots", testProjectID),
			Method:         http.MethodPost,
			Body:           createPlotConfigBullseyeBody,
			ExpectedStatus: http.StatusCreated,
			ExpectedSchema: objSchema,
		}}

	RunAll(t, tests)
}