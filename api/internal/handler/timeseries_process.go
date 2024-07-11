package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/httperr"
	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type processTimeseriesType int

const (
	byTimeseriesRequest processTimeseriesType = iota
	byInstrumentRequest
	byInstrumentGroupRequest
	explorerRequest
)

// ListTimeseriesMeasurementsByTimeseries godoc
//
//	@Summary lists timeseries by timeseries uuid
//	@Tags timeseries
//	@Produce json
//	@Param instrument_id path string true "instrument uuid" Format(uuid)
//	@Param timeseries_id path string true "timeseries uuid" Format(uuid)
//	@Param after query string false "after time" Format(date-time)
//	@param before query string false "before time" Format(date-time)
//	@Param threshold query number false "downsample threshold"
//	@Success 200 {object} model.MeasurementCollection
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /timeseries/{timeseries_id}/measurements [get]
//	@Router /instruments/{instrument_id}/timeseries/{timeseries_id}/measurements [get]
func (h *ApiHandler) ListTimeseriesMeasurementsByTimeseries(c echo.Context) error {
	tsID, err := uuid.Parse(c.Param("timeseries_id"))
	if err != nil {
		return httperr.MalformedID(err)
	}

	isStored, err := h.TimeseriesService.GetStoredTimeseriesExists(c.Request().Context(), tsID)
	if err != nil {
		return httperr.InternalServerError(err)
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

		resBody, err := h.MeasurementService.ListTimeseriesMeasurements(c.Request().Context(), tsID, tw, threshold)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, resBody)
	}

	f := model.ProcessMeasurementFilter{TimeseriesID: &tsID}

	selectMeasurements := selectMeasurementsHandler(h, f, byTimeseriesRequest)
	return selectMeasurements(c)
}

// ListTimeseriesMeasurementsByInstrument godoc
//
//	@Summary lists timeseries measurements by instrument id
//	@Tags timeseries
//	@Produce json
//	@Param instrument_id path string true "instrument uuid" Format(uuid)
//	@Param after query string false "after time" Format(date-time)
//	@Param before query string false "before time" Format(date-time)
//	@Param threshold query number false "downsample threshold"
//	@Success 200 {object} model.MeasurementCollection
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /instruments/{instrument_id}/timeseries_measurements [get]
func (h *ApiHandler) ListTimeseriesMeasurementsByInstrument(c echo.Context) error {
	iID, err := uuid.Parse(c.Param("instrument_id"))
	if err != nil {
		return httperr.MalformedID(err)
	}
	f := model.ProcessMeasurementFilter{InstrumentID: &iID}

	selectMeasurements := selectMeasurementsHandler(h, f, byInstrumentRequest)
	return selectMeasurements(c)
}

// ListTimeseriesMeasurementsByInstrumentGroup godoc
//
//	@Summary lists timeseries measurements by instrument group id
//	@Tags timeseries
//	@Produce json
//	@Param instrument_group_id path string true "instrument group uuid" Format(uuid)
//	@Success 200 {object} model.MeasurementCollection
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /instrument_groups/{instrument_group_id}/timeseries_measurements [get]
func (h *ApiHandler) ListTimeseriesMeasurementsByInstrumentGroup(c echo.Context) error {
	igID, err := uuid.Parse(c.Param("instrument_group_id"))
	if err != nil {
		return httperr.MalformedID(err)
	}
	f := model.ProcessMeasurementFilter{InstrumentGroupID: &igID}

	selectMeasurements := selectMeasurementsHandler(h, f, byInstrumentGroupRequest)
	return selectMeasurements(c)
}

// ListTimeseriesMeasurementsExplorer godoc
//
//	@Summary list timeseries measurements for explorer page
//	@Tags explorer
//	@Accept json
//	@Produce json
//	@Param instrument_ids body []uuid.UUID true "array of instrument uuids"
//	@Success 200 {array} map[uuid.UUID]model.MeasurementCollectionLean
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /explorer [post]
func (h *ApiHandler) ListTimeseriesMeasurementsExplorer(c echo.Context) error {
	var iIDs []uuid.UUID
	if err := (&echo.DefaultBinder{}).BindBody(c, &iIDs); err != nil {
		return httperr.MalformedBody(err)
	}
	f := model.ProcessMeasurementFilter{InstrumentIDs: iIDs}

	selectMeasurements := selectMeasurementsHandler(h, f, explorerRequest)
	return selectMeasurements(c)
}

// ListInclinometerTimeseriesMeasurementsExplorer godoc
//
//	@Summary list inclinometer timeseries measurements for explorer page
//	@Tags explorer
//	@Accept json
//	@Produce json
//	@Param instrument_ids body []uuid.UUID true "array of inclinometer instrument uuids"
//	@Success 200 {array} map[uuid.UUID]model.InclinometerMeasurementCollectionLean
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /inclinometer_explorer [post]
func (h *ApiHandler) ListInclinometerTimeseriesMeasurementsExplorer(c echo.Context) error {
	var iIDs []uuid.UUID
	if err := (&echo.DefaultBinder{}).BindBody(c, &iIDs); err != nil {
		return httperr.MalformedBody(err)
	}
	f := model.ProcessMeasurementFilter{InstrumentIDs: iIDs}

	selectMeasurements := selectInclinometerMeasurementsHandler(h, f)
	return selectMeasurements(c)
}

func selectMeasurementsHandler(h *ApiHandler, f model.ProcessMeasurementFilter, requestType processTimeseriesType) echo.HandlerFunc {
	return func(c echo.Context) error {
		var tw model.TimeWindow
		a, b := c.QueryParam("after"), c.QueryParam("before")
		if err := tw.SetWindow(a, b, time.Now().AddDate(0, 0, -7), time.Now()); err != nil {
			return httperr.MalformedDate(err)
		}

		f.After = tw.After
		f.Before = tw.Before

		trs := c.QueryParam("threshold")

		var threshold int
		if trs != "" {
			tr, err := strconv.Atoi(trs)
			if err != nil {
				return httperr.Message(http.StatusBadRequest, "query parameter `threshold` must be non-negative int")
			}
			threshold = tr
		}

		mrc, err := h.ProcessTimeseriesService.SelectMeasurements(c.Request().Context(), f)
		if err != nil {
			return httperr.InternalServerError(err)
		}

		if requestType == byTimeseriesRequest && f.TimeseriesID != nil {
			resBody, err := mrc.CollectSingleTimeseries(threshold, *f.TimeseriesID)
			if err != nil {
				return httperr.InternalServerError(err)
			}
			return c.JSON(http.StatusOK, resBody)

		} else {
			resBody, err := mrc.GroupByInstrument(threshold)
			if err != nil {
				return httperr.InternalServerError(err)
			}
			return c.JSON(http.StatusOK, resBody)
		}
	}
}

func selectInclinometerMeasurementsHandler(h *ApiHandler, f model.ProcessMeasurementFilter) echo.HandlerFunc {
	return func(c echo.Context) error {
		var tw model.TimeWindow
		a, b := c.QueryParam("after"), c.QueryParam("before")
		if err := tw.SetWindow(a, b, time.Now().AddDate(0, 0, -7), time.Now()); err != nil {
			return httperr.MalformedDate(err)
		}

		f.After = tw.After
		f.Before = tw.Before

		mrc, err := h.ProcessTimeseriesService.SelectInclinometerMeasurements(c.Request().Context(), f)
		if err != nil {
			return httperr.InternalServerError(err)
		}

		resBody, err := mrc.GroupByInstrument()

		return c.JSON(http.StatusOK, resBody)
	}
}
