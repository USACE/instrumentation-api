package handlers

import (
	"net/http"
	"strconv"

	"github.com/USACE/instrumentation-api/api/messages"
	"github.com/USACE/instrumentation-api/api/models"
	"github.com/USACE/instrumentation-api/api/timeseries"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

// Request types
const (
	byTimeseriesRequest = iota
	byInstrumentRequest
	byInstrumentGroupRequest
	explorerRequest
)

func ListTimeseriesMeasurementsByTimeseries(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		tsID, err := uuid.Parse(c.Param("timeseries_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, messages.MalformedID)
		}
		f := models.MeasurementsFilter{TimeseriesID: &tsID}

		selectMeasurements := selectMeasurementsHandler(db, &f, byTimeseriesRequest)
		return selectMeasurements(c)
	}
}

func ListTimeseriesMeasurementsByInstrument(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		iID, err := uuid.Parse(c.Param("instrument_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, messages.MalformedID)
		}
		f := models.MeasurementsFilter{InstrumentID: &iID}

		selectMeasurements := selectMeasurementsHandler(db, &f, byInstrumentRequest)
		return selectMeasurements(c)
	}
}

func ListTimeseriesMeasurementsByInstrumentGroup(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		igID, err := uuid.Parse(c.Param("instrument_group_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, messages.MalformedID)
		}
		f := models.MeasurementsFilter{InstrumentGroupID: &igID}

		selectMeasurements := selectMeasurementsHandler(db, &f, byInstrumentGroupRequest)
		return selectMeasurements(c)
	}
}

func ListTimeseriesMeasurementsExplorer(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var iIDs []uuid.UUID

		// Instrument IDs from POST
		if err := (&echo.DefaultBinder{}).BindBody(c, &iIDs); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		f := models.MeasurementsFilter{InstrumentIDs: iIDs}

		selectMeasurements := selectMeasurementsHandler(db, &f, explorerRequest)
		return selectMeasurements(c)
	}
}

func selectMeasurementsHandler(db *sqlx.DB, f *models.MeasurementsFilter, requestType int) echo.HandlerFunc {
	return func(c echo.Context) error {
		var tw timeseries.TimeWindow
		a, b := c.QueryParam("after"), c.QueryParam("before")
		if err := tw.SetWindow(a, b); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		f.After = tw.After
		f.Before = tw.Before

		trs := c.QueryParam("threshold")

		var threshold int
		if trs != "" {
			tr, err := strconv.Atoi(trs)
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
			threshold = tr
		}

		mrc, err := models.SelectMeasurements(db, f)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		if requestType == byTimeseriesRequest {
			resBody, err := mrc.CollectSingleTimeseries(threshold, f.TimeseriesID)
			if err != nil {
				if err.Error() == messages.NotFound {
					return c.JSON(
						http.StatusOK,
						timeseries.MeasurementCollection{
							TimeseriesID: *f.TimeseriesID,
							Items:        make([]timeseries.Measurement, 0),
						},
					)
				}
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			return c.JSON(http.StatusOK, resBody)

		} else {
			resBody, err := mrc.GroupByInstrument(threshold)
			if err != nil {
				if err.Error() == messages.NotFound {
					return c.JSON(http.StatusOK, make([]map[string]interface{}, 0))
				}
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			return c.JSON(http.StatusOK, resBody)
		}
	}
}
