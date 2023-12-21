package handler

import (
	"net/http"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/message"
	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// GetAllIpiSegmentsForInstrument godoc
//
//	@Summary gets all ipi segments for an instrument
//	@Tags instrument-ipi
//	@Produce json
//	@Param instrument_id path string true "instrument uuid" Format(uuid)
//	@Success 200 {array} model.IpiSegment
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /instruments/ipi/{instrument_id}/segments [get]
func (h *ApiHandler) GetAllIpiSegmentsForInstrument(c echo.Context) error {
	iID, err := uuid.Parse(c.Param("instrument_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}
	ss, err := h.IpiInstrumentService.GetAllIpiSegmentsForInstrument(c.Request().Context(), iID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, ss)
}

// GetIpiMeasurementsForInstrument godoc
//
//	@Summary creates instrument notes
//	@Tags instrument-ipi
//	@Produce json
//	@Param instrument_id path string true "instrument uuid" Format(uuid)
//	@Param after query string false "after time" Format(date-time)
//	@Param before query string true "before time" Format(date-time)
//	@Success 200 {array} model.IpiMeasurements
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /instruments/ipi/{instrument_id}/measurements [get]
func (h *ApiHandler) GetIpiMeasurementsForInstrument(c echo.Context) error {
	iID, err := uuid.Parse(c.Param("instrument_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}
	var tw model.TimeWindow
	a, b := c.QueryParam("after"), c.QueryParam("before")
	if err := tw.SetWindow(a, b, time.Now().AddDate(0, 0, -7), time.Now()); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	mm, err := h.IpiInstrumentService.GetIpiMeasurementsForInstrument(c.Request().Context(), iID, tw)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, mm)
}

// UpdateIpiSegments godoc
//
//	@Summary updates multiple segments for an ipi instrument
//	@Tags instrument-ipi
//	@Produce json
//	@Param instrument_id path string true "instrument uuid" Format(uuid)
//	@Param instrument_segments body []model.IpiSegment true "ipi instrument segments payload"
//	@Success 200 {array} model.IpiSegment
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /instruments/ipi/{instrument_id}/segments [put]
//	@Security Bearer
func (h *ApiHandler) UpdateIpiSegments(c echo.Context) error {
	segs := make([]model.IpiSegment, 0)
	if err := c.Bind(&segs); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := h.IpiInstrumentService.UpdateIpiSegments(c.Request().Context(), segs); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, segs)
}
