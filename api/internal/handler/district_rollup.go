package handler

import (
	"net/http"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/message"
	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// ListEvaluationDistrictRollup returns monthly evaluation statistics for the past year
func (h *ApiHandler) ListProjectEvaluationDistrictRollup(c echo.Context) error {
	id, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}

	var tw model.TimeWindow
	from, to := c.QueryParam("from_timestamp_month"), c.QueryParam("to_timestamp_month")
	if err := tw.SetWindow(from, to, time.Now().AddDate(-1, 0, 0), time.Now()); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if fiveYrsAfterStart := tw.Start.AddDate(5, 0, 0); tw.End.After(fiveYrsAfterStart) {
		return echo.NewHTTPError(http.StatusBadRequest, "maximum requested time range exceeded (5 years)")
	}

	project, err := h.DistrictRollupService.ListEvaluationDistrictRollup(c.Request().Context(), id, tw)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, project)
}

// ListMeasurementDistrictRollup returns monthly measurement statistics for the past year
func (h *ApiHandler) ListProjectMeasurementDistrictRollup(c echo.Context) error {
	id, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}

	var tw model.TimeWindow
	from, to := c.QueryParam("from_timestamp_month"), c.QueryParam("to_timestamp_month")
	if err := tw.SetWindow(from, to, time.Now().AddDate(-1, 0, 0), time.Now()); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if fiveYrsAfterStart := tw.Start.AddDate(5, 0, 0); tw.End.After(fiveYrsAfterStart) {
		return echo.NewHTTPError(http.StatusBadRequest, "maximum requested time range exceeded (5 years)")
	}

	project, err := h.DistrictRollupService.ListMeasurementDistrictRollup(c.Request().Context(), id, tw)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, project)
}
