package handler_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/USACE/instrumentation-api/api/internal/model"
)

func TestDistrictRollup(t *testing.T) {
	tests := []HTTPTest[model.DistrictRollup]{
		{
			Name:                 "ListProjectEvaluationDistrictRollup",
			URL:                  fmt.Sprintf("/projects/%s/district_rollup/evaluation_submittals?from_timestamp_month=&to_timestamp_month=", testProjectID),
			Method:               http.MethodGet,
			ExpectedStatus:       http.StatusOK,
			ExpectedResponseType: jsonArr,
		},
		{
			Name:                 "ListProjectMeasurementDistrictRollup",
			URL:                  fmt.Sprintf("/projects/%s/district_rollup/measurement_submittals?from_timestamp_month=&to_timestamp_month=", testProjectID),
			Method:               http.MethodGet,
			ExpectedStatus:       http.StatusOK,
			ExpectedResponseType: jsonArr,
		}}

	RunAll(t, tests)
}
