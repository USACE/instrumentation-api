package handler

import (
	"net/http"

	"github.com/USACE/instrumentation-api/api/internal/messages"
	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// CreateInstrumentConstants creates instrument constants (i.e. timeseries)
func (h ApiHandler) CreateInstrumentConstants(c echo.Context) error {
	ctx := c.Request().Context()
	// Get action information from context
	var tc model.TimeseriesCollectionItems
	if err := c.Bind(&tc); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// InstrumentID From RouteParams
	instrumentID, err := uuid.Parse(c.Param("instrument_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// slugs already taken in the database
	slugsTaken, err := h.TimeseriesStore.ListTimeseriesSlugsForInstrument(ctx, instrumentID)
	if err != nil {
		return err
	}
	for idx := range tc.Items {
		// Verify object instrument_id matches routeParam
		if instrumentID != tc.Items[idx].InstrumentID {
			return echo.NewHTTPError(http.StatusBadRequest, messages.MatchRouteParam("`instrument_id`"))
		}
		// Assign Slug
		s, err := h.TimeseriesStore.NextUniqueSlugTimeseries(ctx, tc.Items[idx].Name, slugsTaken)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		tc.Items[idx].Slug = s
		// Add slug to array of slugs originally fetched from the database
		// to catch duplicate names/slugs from the same bulk upload
		slugsTaken = append(slugsTaken, s)
	}
	tt, err := h.InstrumentConstantStore.CreateInstrumentConstants(ctx, tc.Items)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, tt)
}

// DeleteInstrumentConstant removes a timeseries as an Instrument Constant
func (h ApiHandler) DeleteInstrumentConstant(c echo.Context) error {
	instrumentID, err := uuid.Parse(c.Param("instrument_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	timeseriesID, err := uuid.Parse(c.Param("timeseries_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err = h.InstrumentConstantStore.DeleteInstrumentConstant(c.Request().Context(), instrumentID, timeseriesID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, make(map[string]interface{}))
}

// ListInstrumentConstants lists constants for a given instrument
func (h ApiHandler) ListInstrumentConstants(c echo.Context) error {
	instrumentID, err := uuid.Parse(c.Param("instrument_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	cc, err := h.InstrumentConstantStore.ListInstrumentConstants(c.Request().Context(), instrumentID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, cc)
}
