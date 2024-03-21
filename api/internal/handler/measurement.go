package handler

import (
	"net/http"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/model"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// CreateOrUpdateProjectTimeseriesMeasurements godoc
//
//	@Summary creates or updates one or more timeseries measurements
//	@Tags measurement
//	@Accept json
//	@Produce json
//	@Param project_id path string true "project uuid" Format(uuid)
//	@Param timeseries_measurement_collections body model.TimeseriesMeasurementCollectionCollection true "array of timeseries measurement collections"
//	@Success 200 {array} model.MeasurementCollection
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/timeseries_measurements [post]
//	@Security Bearer
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
	if err := h.TimeseriesService.AssertTimeseriesLinkedToProject(ctx, pID, dd); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	stored, err := h.MeasurementService.CreateOrUpdateTimeseriesMeasurements(ctx, mcc.Items)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, stored)
}

// CreateOrUpdateTimeseriesMeasurements godoc
//
//	@Summary creates or updates one or more timeseries measurements
//	@Tags measurement
//	@Produce json
//	@Param timeseries_measurement_collections body model.TimeseriesMeasurementCollectionCollection true "array of timeseries measurement collections"
//	@Success 200 {array} model.MeasurementCollection
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /timeseries_measurements [post]
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

// UpdateTimeseriesMeasurements godoc
//
//	@Summary overwrites all measurements witin date range with the supplied payload
//	@Tags measurement
//	@Produce json
//	@Param project_id path string true "project uuid" Format(uuid)
//	@Param after query string false "after timestamp" Format(date-time)
//	@Param before query string false "before timestamp" Format(date-time)
//	@Param timeseries_measurement_collections body model.TimeseriesMeasurementCollectionCollection true "array of timeseries measurement collections"
//	@Success 200 {array} model.MeasurementCollection
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/timeseries_measurements [put]
//	@Security Bearer
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

// DeleteTimeserieMeasurements godoc
//
//	@Summary deletes a single timeseries measurement by timestamp
//	@Tags measurement
//	@Produce json
//	@Param timeseries_id path string true "timeseries uuid" Format(uuid)
//	@Param time query string true "timestamp of measurement to delete" Format(date-time)
//	@Success 200 {object} map[string]interface{}
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /timeseries/{timeseries_id}/measurements [delete]
//	@Security Bearer
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
