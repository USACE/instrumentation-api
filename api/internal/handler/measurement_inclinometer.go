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

// ListInclinometerMeasurements returns a timeseries with inclinometer measurements
func (h *ApiHandler) ListInclinometerMeasurements(c echo.Context) error {

	tsID, err := uuid.Parse(c.Param("timeseries_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}

	// Time Window
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

// CreateOrUpdateProjectInclinometerMeasurements Creates or Updates a InclinometerMeasurement object or array of objects
// All timeseries must belong to the same project
func (h *ApiHandler) CreateOrUpdateProjectInclinometerMeasurements(c echo.Context) error {
	var mcc model.InclinometerMeasurementCollectionCollection
	if err := c.Bind(&mcc); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	pID, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	dd := mcc.InclinometerTimeseriesIDs()
	m, err := h.TimeseriesService.GetTimeseriesProjectMap(c.Request().Context(), dd)
	if err != nil {
		return err
	}
	for _, tID := range dd {
		ppID, ok := m[tID]
		if !ok || pID != ppID {
			return echo.NewHTTPError(http.StatusBadRequest, "all timeseries posted do not belong to project")
		}
	}
	p, ok := c.Get("profile").(model.Profile)
	if !ok {
		return echo.NewHTTPError(http.StatusBadRequest, "could not get profile")
	}

	stored, err := h.InclinometerMeasurementService.CreateOrUpdateInclinometerMeasurements(c.Request().Context(), mcc.Items, p, time.Now())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	//create inclinometer constant if doesn't exist
	if len(mcc.Items) > 0 {
		cm, err := h.MeasurementService.GetTimeseriesConstantMeasurement(c.Request().Context(), mcc.Items[0].TimeseriesID, "inclinometer-constant")
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		if cm.TimeseriesID == uuid.Nil {
			err := h.InclinometerMeasurementService.CreateTimeseriesConstant(c.Request().Context(), mcc.Items[0].TimeseriesID, "inclinometer-constant", "Meters", 20000)
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
		}

	}
	return c.JSON(http.StatusCreated, stored)
}

// DeleteInclinometerMeasurements deletes a single inclinometer measurement
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
