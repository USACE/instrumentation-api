package handler

import (
	"net/http"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/httperr"
	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// GetAllSaaSegmentsForInstrument godoc
//
//	@Summary gets all saa segments for an instrument
//	@Tags instrument-saa
//	@Produce json
//	@Param instrument_id path string true "instrument uuid" Format(uuid)
//	@Success 200 {array} model.SaaSegment
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /instruments/saa/{instrument_id}/segments [get]
func (h *ApiHandler) GetAllSaaSegmentsForInstrument(c echo.Context) error {
	iID, err := uuid.Parse(c.Param("instrument_id"))
	if err != nil {
		return httperr.MalformedID(err)
	}
	ss, err := h.SaaInstrumentService.GetAllSaaSegmentsForInstrument(c.Request().Context(), iID)
	if err != nil {
		return httperr.InternalServerError(err)
	}
	return c.JSON(http.StatusOK, ss)
}

// GetSaaMeasurementsForInstrument godoc
//
//	@Summary creates instrument notes
//	@Tags instrument-saa
//	@Produce json
//	@Param instrument_id path string true "instrument uuid" Format(uuid)
//	@Param after query string false "after time" Format(date-time)
//	@Param before query string true "before time" Format(date-time)
//	@Success 200 {array} model.SaaMeasurements
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /instruments/saa/{instrument_id}/measurements [get]
func (h *ApiHandler) GetSaaMeasurementsForInstrument(c echo.Context) error {
	iID, err := uuid.Parse(c.Param("instrument_id"))
	if err != nil {
		return httperr.MalformedID(err)
	}
	var tw model.TimeWindow
	a, b := c.QueryParam("after"), c.QueryParam("before")
	if err := tw.SetWindow(a, b, time.Now().AddDate(0, 0, -7), time.Now()); err != nil {
		return httperr.MalformedDate(err)
	}
	mm, err := h.SaaInstrumentService.GetSaaMeasurementsForInstrument(c.Request().Context(), iID, tw)
	if err != nil {
		return httperr.MalformedID(err)
	}
	return c.JSON(http.StatusOK, mm)
}

// UpdateSaaSegments godoc
//
//	@Summary updates multiple segments for an saa instrument
//	@Tags instrument-saa
//	@Produce json
//	@Param instrument_id path string true "instrument uuid" Format(uuid)
//	@Param instrument_segments body []model.SaaSegment true "saa instrument segments payload"
//	@Param key query string false "api key"
//	@Success 200 {array} model.SaaSegment
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /instruments/saa/{instrument_id}/segments [put]
//	@Security Bearer
func (h *ApiHandler) UpdateSaaSegments(c echo.Context) error {
	segs := make([]model.SaaSegment, 0)
	if err := c.Bind(&segs); err != nil {
		return httperr.MalformedBody(err)
	}
	if err := h.SaaInstrumentService.UpdateSaaSegments(c.Request().Context(), segs); err != nil {
		return httperr.InternalServerError(err)
	}
	return c.JSON(http.StatusOK, segs)
}
