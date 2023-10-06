package handler

import (
	"net/http"

	"github.com/USACE/instrumentation-api/api/internal/message"
	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// UpdateSaaInstrument(ctx context.Context, si model.SaaInstrument) error
// UpdateSaaInstrumentSegment(ctx context.Context, seg model.SaaSegment) error

func (h *ApiHandler) CreateSaaInstrument(c echo.Context) error {
	var si model.SaaInstrumentWithSegments
	if err := c.Bind(&si); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := h.SaaInstrumentService.CreateSaaInstrument(c.Request().Context(), si); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, make(map[string]interface{}))
}

func (h *ApiHandler) CreateSaaSegments(c echo.Context) error {
	var segs []model.SaaSegment
	if err := c.Bind(&segs); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := h.SaaInstrumentService.CreateSaaSegments(c.Request().Context(), segs); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, make(map[string]interface{}))
}

func (h *ApiHandler) GetOneSaaInstrumentWithSegments(c echo.Context) error {
	iID, err := uuid.Parse(c.Param("instrument_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}
	ss, err := h.SaaInstrumentService.GetOneSaaInstrumentWithSegments(c.Request().Context(), iID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, ss)
}

func (h *ApiHandler) GetAllSaaInstrumentsWithSegmentsForProject(c echo.Context) error {
	pID, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}
	ss, err := h.SaaInstrumentService.GetAllSaaInstrumentsWithSegmentsForProject(c.Request().Context(), pID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, ss)
}

func (h *ApiHandler) UpdateSaaInstrument(c echo.Context) error {
	var si model.SaaInstrumentWithSegments
	if err := c.Bind(&si); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := h.SaaInstrumentService.UpdateSaaInstrument(c.Request().Context(), si); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, si)
}

func (h *ApiHandler) UpdateSaaInstrumentSegment(c echo.Context) error {
	var seg model.SaaSegment
	if err := c.Bind(&seg); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := h.SaaInstrumentService.UpdateSaaInstrumentSegment(c.Request().Context(), seg); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, seg)
}
