package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/messages"
	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/USACE/instrumentation-api/api/internal/models"
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

		// Not ideal making 2 calls to database here, but need to check if timeseries
		// is computed to know when to return timeseries notes
		// Also, returning only stored timeseries is much faster with the current query

		isStored, err := models.StoredTimeseriesExists(db, &tsID)
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

			resBody, err := models.ListTimeseriesMeasurements(db, &tsID, &tw, threshold)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			return c.JSON(http.StatusOK, resBody)
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
}
