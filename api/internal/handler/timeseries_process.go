package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/message"
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

func (h *ApiHandler) ListTimeseriesMeasurementsByTimeseries(c echo.Context) error {
	tsID, err := uuid.Parse(c.Param("timeseries_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}

	// TODO: move business logic to service layer

	isStored, err := h.TimeseriesService.GetStoredTimeseriesExists(c.Request().Context(), tsID)
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

func (h *ApiHandler) ListTimeseriesMeasurementsByInstrument(c echo.Context) error {
	iID, err := uuid.Parse(c.Param("instrument_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}
	f := model.ProcessMeasurementFilter{InstrumentID: &iID}

	selectMeasurements := selectMeasurementsHandler(h, f, byInstrumentRequest)
	return selectMeasurements(c)
}

func (h *ApiHandler) ListTimeseriesMeasurementsByInstrumentGroup(c echo.Context) error {
	igID, err := uuid.Parse(c.Param("instrument_group_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}
	f := model.ProcessMeasurementFilter{InstrumentGroupID: &igID}

	selectMeasurements := selectMeasurementsHandler(h, f, byInstrumentGroupRequest)
	return selectMeasurements(c)
}

func (h *ApiHandler) ListTimeseriesMeasurementsExplorer(c echo.Context) error {
	var iIDs []uuid.UUID
	if err := (&echo.DefaultBinder{}).BindBody(c, &iIDs); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	f := model.ProcessMeasurementFilter{InstrumentIDs: iIDs}

	selectMeasurements := selectMeasurementsHandler(h, f, explorerRequest)
	return selectMeasurements(c)
}

func selectMeasurementsHandler(h *ApiHandler, f model.ProcessMeasurementFilter, requestType processTimeseriesType) echo.HandlerFunc {
	return func(c echo.Context) error {
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

		mrc, err := h.ProcessTimeseriesService.SelectMeasurements(c.Request().Context(), f)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		if requestType == byTimeseriesRequest && f.TimeseriesID != nil {
			resBody, err := mrc.CollectSingleTimeseries(threshold, *f.TimeseriesID)
			if err != nil {
				if err.Error() == message.NotFound {
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
				if err.Error() == message.NotFound {
					return c.JSON(http.StatusOK, make([]map[string]interface{}, 0))
				}
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			return c.JSON(http.StatusOK, resBody)
		}
	}
}
