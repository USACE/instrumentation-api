package handler_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xeipuuv/gojsonschema"
)

var plotConfigProfileSchema = fmt.Sprintf(plotConfigBaseSchema, IDSlugNameArrSchema, plotConfigProfileDisplaySchema)

const plotConfigProfileDisplaySchema = `{
    "instrument_id": { "type": "string" }
}`

var plotConfigProfileObjectLoader = gojsonschema.NewStringLoader(plotConfigProfileSchema)

const testPlotConfigProfileID = "f08a15fa-448c-4066-a81f-124d54712c62"

const updatePlotConfigProfileBody = `{
    "id": "871e34da-c911-4d8f-ab68-e29ba17f8937",
    "name": "Updated Profile Plot SAA to IPI",
    "slug": "updated-profile-plot-saa-to-ipi",
    "project_id": "5b6f4f37-7755-4cf9-bd02-94f1e9bc5984",
    "plot_type": "profile",
    "display": {
        "instrument_id": "eca4040e-aecb-4cd3-bcde-3e308f0356a6"
    }
}`

const createPlotConfigProfileBody = `{
    "name": "New IPI Profile Plot",
    "project_id": "5b6f4f37-7755-4cf9-bd02-94f1e9bc5984",
    "plot_type": "profile",
    "display": {
        "instrument_id": "eca4040e-aecb-4cd3-bcde-3e308f0356a6"
    }
}`

func TestPlotConfigsProfile(t *testing.T) {
	objSchema, err := gojsonschema.NewSchema(plotConfigProfileObjectLoader)
	assert.Nil(t, err)

	tests := []HTTPTest{
		{
			Name:           "UpdatePlotConfigProfilePlot",
			URL:            fmt.Sprintf("/projects/%s/plot_configs/profile_plots/%s", testProjectID, testPlotConfigProfileID),
			Method:         http.MethodPut,
			Body:           updatePlotConfigProfileBody,
			ExpectedStatus: http.StatusOK,
			ExpectedSchema: objSchema,
		},
		{
			Name:           "CreatePlotConfigProfilePlot",
			URL:            fmt.Sprintf("/projects/%s/plot_configs/profile_plots", testProjectID),
			Method:         http.MethodPost,
			Body:           createPlotConfigProfileBody,
			ExpectedStatus: http.StatusCreated,
			ExpectedSchema: objSchema,
		}}

	RunAll(t, tests)
}
