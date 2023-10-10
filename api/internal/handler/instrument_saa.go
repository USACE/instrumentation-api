package handler

import (
	"net/http"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/message"
	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (h *ApiHandler) GetAllSaaSegmentsForInstrument(c echo.Context) error {
	iID, err := uuid.Parse(c.Param("instrument_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}
	ss, err := h.SaaInstrumentService.GetAllSaaSegmentsForInstrument(c.Request().Context(), iID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, ss)
}

func (h *ApiHandler) GetSaaMeasurementsForInstrument(c echo.Context) error {
	iID, err := uuid.Parse(c.Param("instrument_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}
	var tw model.TimeWindow
	a, b := c.QueryParam("after"), c.QueryParam("before")
	if err := tw.SetWindow(a, b, time.Now().AddDate(0, 0, -7), time.Now()); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	mm, err := h.SaaInstrumentService.GetSaaMeasurementsForInstrument(c.Request().Context(), iID, tw)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, mm)
}

func (h *ApiHandler) UpdateSaaSegments(c echo.Context) error {
	segs := make([]model.SaaSegment, 0)
	if err := c.Bind(&segs); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := h.SaaInstrumentService.UpdateSaaSegments(c.Request().Context(), segs); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, segs)
}

func (h *ApiHandler) UpdateSaaSegment(c echo.Context) error {
	var seg model.SaaSegment
	if err := c.Bind(&seg); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := h.SaaInstrumentService.UpdateSaaSegment(c.Request().Context(), seg); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, seg)
}
