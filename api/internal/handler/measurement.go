package handler

import (
	"net/http"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/model"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// CreateOrUpdateProjectTimeseriesMeasurements Creates or Updates a TimeseriesMeasurement object or array of objects
// All timeseries must belong to the same project
func (h *ApiHandler) CreateOrUpdateProjectTimeseriesMeasurements(c echo.Context) error {
	ctx := c.Request().Context()
	var mcc model.TimeseriesMeasurementCollectionCollection
	if err := c.Bind(&mcc); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	pID, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	dd := mcc.TimeseriesIDs()
	m, err := h.TimeseriesService.GetTimeseriesProjectMap(ctx, dd)
	if err != nil {
		return err
	}
	for _, tID := range dd {
		ppID, ok := m[tID]
		if !ok || ppID != pID {
			return echo.NewHTTPError(http.StatusBadRequest, "all timeseries posted do not belong to project")
		}
	}

	stored, err := h.MeasurementService.CreateOrUpdateTimeseriesMeasurements(ctx, mcc.Items)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, stored)
}

// CreateOrUpdateTimeseriesMeasurements Creates or Updates a TimeseriesMeasurement object or array of objects
// Timeseries may belong to one or more projects
func (h *ApiHandler) CreateOrUpdateTimeseriesMeasurements(c echo.Context) error {
	var mcc model.TimeseriesMeasurementCollectionCollection
	if err := c.Bind(&mcc); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	stored, err := h.MeasurementService.CreateOrUpdateTimeseriesMeasurements(c.Request().Context(), mcc.Items)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, stored)
}

// UpdateTimeseriesMeasurements Overwrites measurements with the supplied payload
// within a TimeWindow (> after, < before)
func (h *ApiHandler) UpdateTimeseriesMeasurements(c echo.Context) error {
	var tw model.TimeWindow
	a, b := c.QueryParam("after"), c.QueryParam("before")
	if err := tw.SetWindow(a, b, time.Now().AddDate(0, 0, -7), time.Now()); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var mcc model.TimeseriesMeasurementCollectionCollection
	if err := c.Bind(&mcc); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	stored, err := h.MeasurementService.UpdateTimeseriesMeasurements(c.Request().Context(), mcc.Items, tw)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, stored)
}

// DeleteTimeserieMeasurements deletes a single timeseries measurement
func (h *ApiHandler) DeleteTimeserieMeasurements(c echo.Context) error {
	// id from url params
	id, err := uuid.Parse(c.Param("timeseries_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	timeString := c.QueryParam("time")

	t, err := time.Parse(time.RFC3339, timeString)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := h.MeasurementService.DeleteTimeserieMeasurements(c.Request().Context(), id, t); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, make(map[string]interface{}))
}
