package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/message"
	"github.com/USACE/instrumentation-api/api/internal/model"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// ListInclinometerMeasurements godoc
//
//	@Summary lists all measurements for an inclinometer
//	@Tags measurement-inclinometer
//	@Produce json
//	@Param timeseries_id path string true "timeseries uuid" Format(uuid)
//	@Param after query string false "after timestamp" Format(date-time)
//	@Param before query string false "before timestamp" Format(date-time)
//	@Success 200 {object} model.InclinometerMeasurementCollection
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /timeseries/{timeseries_id}/inclinometer_measurements [get]
func (h *ApiHandler) ListInclinometerMeasurements(c echo.Context) error {

	tsID, err := uuid.Parse(c.Param("timeseries_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}

	var tw model.TimeWindow
	a, b := c.QueryParam("after"), c.QueryParam("before")
	if err = tw.SetWindow(a, b, time.Now().AddDate(0, 0, -7), time.Now()); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	im, err := h.InclinometerMeasurementService.ListInclinometerMeasurements(c.Request().Context(), tsID, tw)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	cm, err := h.MeasurementService.GetTimeseriesConstantMeasurement(c.Request().Context(), tsID, "inclinometer-constant")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	for idx := range im.Inclinometers {
		values, err := h.InclinometerMeasurementService.ListInclinometerMeasurementValues(c.Request().Context(), tsID, im.Inclinometers[idx].Time, cm.Value)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		jsonValues, err := json.Marshal(values)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		im.Inclinometers[idx].Values = jsonValues
	}

	return c.JSON(http.StatusOK, im)
}

// CreateOrUpdateProjectInclinometerMeasurements godoc
//
//	@Summary creates or updates one or more inclinometer measurements
//	@Tags measurement-inclinometer
//	@Produce json
//	@Param project_id path string true "project uuid" Format(uuid)
//	@Param timeseries_measurement_collections body model.InclinometerMeasurementCollectionCollection true "inclinometer measurement collections"
//	@Success 200 {array} model.InclinometerMeasurementCollection
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/inclinometer_measurements [post]
//	@Security Bearer
func (h *ApiHandler) CreateOrUpdateProjectInclinometerMeasurements(c echo.Context) error {
	var mcc model.InclinometerMeasurementCollectionCollection
	if err := c.Bind(&mcc); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	pID, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()

	dd := mcc.TimeseriesIDs()
	if err := h.TimeseriesService.AssertTimeseriesLinkedToProject(ctx, pID, dd); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	p, ok := c.Get("profile").(model.Profile)
	if !ok {
		return echo.NewHTTPError(http.StatusBadRequest, "could not get profile")
	}

	stored, err := h.InclinometerMeasurementService.CreateOrUpdateInclinometerMeasurements(ctx, mcc.Items, p, time.Now())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	//create inclinometer constant if doesn't exist
	if len(mcc.Items) > 0 {
		cm, err := h.MeasurementService.GetTimeseriesConstantMeasurement(ctx, mcc.Items[0].TimeseriesID, "inclinometer-constant")
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		if cm.TimeseriesID == uuid.Nil {
			err := h.InclinometerMeasurementService.CreateTimeseriesConstant(ctx, mcc.Items[0].TimeseriesID, "inclinometer-constant", "Meters", 20000)
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
		}

	}
	return c.JSON(http.StatusCreated, stored)
}

// DeleteInclinometerMeasurements godoc
//
//	@Summary deletes a single inclinometer measurement by timestamp
//	@Tags measurement-inclinometer
//	@Produce json
//	@Param timeseries_id path string true "timeseries uuid" Format(uuid)
//	@Param time query string true "timestamp of measurement to delete" Format(date-time)
//	@Success 200 {object} map[string]interface{}
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /timeseries/{timeseries_id}/inclinometer_measurements [delete]
//	@Security Bearer
func (h *ApiHandler) DeleteInclinometerMeasurements(c echo.Context) error {
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

	if err := h.InclinometerMeasurementService.DeleteInclinometerMeasurement(c.Request().Context(), id, t); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, make(map[string]interface{}))
}
