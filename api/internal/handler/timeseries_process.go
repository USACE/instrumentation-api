package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/messages"
	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// Request types
const (
	byTimeseriesRequest = iota
	byInstrumentRequest
	byInstrumentGroupRequest
	explorerRequest
)

func (h ApiHandler) ListTimeseriesMeasurementsByTimeseries(c echo.Context) error {
	tsID, err := uuid.Parse(c.Param("timeseries_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, messages.MalformedID)
	}

	// Not ideal making 2 calls to database here, but need to check if timeseries
	// is computed to know when to return timeseries notes
	// Also, returning only stored timeseries is much faster with the current query

	isStored, err := model.StoredTimeseriesExists(db, &tsID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if isStored {
		var tw model.TimeWindow
		a, b := c.QueryParam("after"), c.QueryParam("before")
		if err := tw.SetWindow(a, b, time.Now().AddDate(0, 0, -7), time.Now()); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		trs := c.QueryParam("threshold")

		var threshold int
		if trs != "" {
			tr, err := strconv.Atoi(trs)
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
			threshold = tr
		}

		resBody, err := model.ListTimeseriesMeasurements(db, &tsID, &tw, threshold)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, resBody)
	}

	f := model.MeasurementsFilter{TimeseriesID: &tsID}

	selectMeasurements := selectMeasurementsHandler(db, &f, byTimeseriesRequest)
	return selectMeasurements(c)
}

func (h ApiHandler) ListTimeseriesMeasurementsByInstrument(c echo.Context) error {
	iID, err := uuid.Parse(c.Param("instrument_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, messages.MalformedID)
	}
	f := model.MeasurementsFilter{InstrumentID: &iID}

	selectMeasurements := selectMeasurementsHandler(db, &f, byInstrumentRequest)
	return selectMeasurements(c)
}

func (h ApiHandler) ListTimeseriesMeasurementsByInstrumentGroup(c echo.Context) error {
	igID, err := uuid.Parse(c.Param("instrument_group_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, messages.MalformedID)
	}
	f := model.MeasurementsFilter{InstrumentGroupID: &igID}

	selectMeasurements := selectMeasurementsHandler(db, &f, byInstrumentGroupRequest)
	return selectMeasurements(c)
}

func (h ApiHandler) ListTimeseriesMeasurementsExplorer(c echo.Context) error {
	var iIDs []uuid.UUID

	// Instrument IDs from POST
	if err := (&echo.DefaultBinder{}).BindBody(c, &iIDs); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	f := model.MeasurementsFilter{InstrumentIDs: iIDs}

	selectMeasurements := selectMeasurementsHandler(db, &f, explorerRequest)
	return selectMeasurements(c)
}

func (h ApiHandler) selectMeasurementsHandler(c echo.Context) error {
	var tw model.TimeWindow
	a, b := c.QueryParam("after"), c.QueryParam("before")
	if err := tw.SetWindow(a, b, time.Now().AddDate(0, 0, -7), time.Now()); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	f.After = tw.Start
	f.Before = tw.End

	trs := c.QueryParam("threshold")

	var threshold int
	if trs != "" {
		tr, err := strconv.Atoi(trs)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		threshold = tr
	}

	mrc, err := model.SelectMeasurements(db, f)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if requestType == byTimeseriesRequest {
		resBody, err := mrc.CollectSingleTimeseries(threshold, f.TimeseriesID)
		if err != nil {
			if err.Error() == messages.NotFound {
				return c.JSON(
					http.StatusOK,
					model.MeasurementCollection{
						TimeseriesID: *f.TimeseriesID,
						Items:        make([]model.Measurement, 0),
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
