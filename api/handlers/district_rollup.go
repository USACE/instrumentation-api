package handlers

import (
	"net/http"
	"time"

	"github.com/USACE/instrumentation-api/api/messages"
	"github.com/USACE/instrumentation-api/api/models"
	"github.com/USACE/instrumentation-api/api/timeseries"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

// ListEvaluationDistrictRollup returns monthly evaluation statistics for the past year
func ListProjectEvaluationDistrictRollup(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := uuid.Parse(c.Param("project_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, messages.MalformedID)
		}

		var tw timeseries.TimeWindow
		from, to := c.QueryParam("from_timestamp_month"), c.QueryParam("to_timestamp_month")
		if err := tw.SetWindow(from, to, time.Now(), time.Now().AddDate(-1, 0, 0)); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if fiveYrsAfterStart := tw.Start.AddDate(5, 0, 0); tw.End.After(fiveYrsAfterStart) {
			return echo.NewHTTPError(http.StatusBadRequest, "maximum requested time range exceeded (5 years)")
		}

		project, err := models.ListEvaluationDistrictRollup(db, id, tw)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, project)
	}
}

// ListMeasurementDistrictRollup returns monthly measurement statistics for the past year
func ListProjectMeasurementDistrictRollup(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := uuid.Parse(c.Param("project_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, messages.MalformedID)
		}

		var tw timeseries.TimeWindow
		from, to := c.QueryParam("from_timestamp_month"), c.QueryParam("to_timestamp_month")
		if err := tw.SetWindow(from, to, time.Now(), time.Now().AddDate(-1, 0, 0)); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if fiveYrsAfterStart := tw.Start.AddDate(5, 0, 0); tw.End.After(fiveYrsAfterStart) {
			return echo.NewHTTPError(http.StatusBadRequest, "maximum requested time range exceeded (5 years)")
		}

		project, err := models.ListMeasurementDistrictRollup(db, id, tw)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, project)
	}
}
