package handler

import (
	"net/http"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/httperr"
	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

const timeRangeErrMessage = "maximum requested time range exceeded (5 years)"

// ListEvaluationDistrictRollup godoc
//
//	@Summary lists monthly evaluation statistics for a district by project id
//	@Tags district-rollup
//	@Produce json
//	@Param project_id path string true "project id" Format(uuid)
//	@Success 200 {array} model.DistrictRollup
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/district_rollup/evaluation_submittals [get]
func (h *ApiHandler) ListProjectEvaluationDistrictRollup(c echo.Context) error {
	id, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		httperr.MalformedID(err)
	}

	var tw model.TimeWindow
	from, to := c.QueryParam("from_timestamp_month"), c.QueryParam("to_timestamp_month")
	if err := tw.SetWindow(from, to, time.Now().AddDate(-1, 0, 0), time.Now()); err != nil {
		return httperr.MalformedDate(err)
	}
	if fiveYrsAfterStart := tw.After.AddDate(5, 0, 0); tw.Before.After(fiveYrsAfterStart) {
		return httperr.Message(http.StatusBadRequest, timeRangeErrMessage)
	}

	project, err := h.DistrictRollupService.ListEvaluationDistrictRollup(c.Request().Context(), id, tw)
	if err != nil {
		return httperr.InternalServerError(err)
	}
	return c.JSON(http.StatusOK, project)
}

// ListMeasurementDistrictRollup godoc
//
//	@Summary lists monthly measurement statistics for a district by project id
//	@Tags district-rollup
//	@Produce json
//	@Param project_id path string true "project id" Format(uuid)
//	@Success 200 {array} model.DistrictRollup
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/district_rollup/measurement_submittals [get]
func (h *ApiHandler) ListProjectMeasurementDistrictRollup(c echo.Context) error {
	id, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return httperr.MalformedID(err)
	}

	var tw model.TimeWindow
	from, to := c.QueryParam("from_timestamp_month"), c.QueryParam("to_timestamp_month")
	if err := tw.SetWindow(from, to, time.Now().AddDate(-1, 0, 0), time.Now()); err != nil {
		return httperr.MalformedDate(err)
	}
	if fiveYrsAfterStart := tw.After.AddDate(5, 0, 0); tw.Before.After(fiveYrsAfterStart) {
		return httperr.Message(http.StatusBadRequest, timeRangeErrMessage)
	}

	project, err := h.DistrictRollupService.ListMeasurementDistrictRollup(c.Request().Context(), id, tw)
	if err != nil {
		return httperr.InternalServerError(err)
	}
	return c.JSON(http.StatusOK, project)
}
