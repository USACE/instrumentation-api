package handler

import (
	"net/http"

	"github.com/USACE/instrumentation-api/api/internal/message"
	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// ListInstrumentConstants godoc
//
//	@Summary lists constants for a given instrument
//	@Tags instrument-constant
//	@Produce json
//	@Param project_id path string true "project uuid" Format(uuid)
//	@Param instrument_id path string true "instrument uuid" Format(uuid)
//	@Success 200 {array} model.Timeseries
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/instruments/{instrument_id}/constants [get]
func (h *ApiHandler) ListInstrumentConstants(c echo.Context) error {
	instrumentID, err := uuid.Parse(c.Param("instrument_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	cc, err := h.InstrumentConstantService.ListInstrumentConstants(c.Request().Context(), instrumentID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, cc)
}

// CreateInstrumentConstants godoc
//
//	@Summary creates instrument constants (i.e. timeseries)
//	@Tags instrument-constant
//	@Produce json
//	@Param project_id path string true "project uuid" Format(uuid)
//	@Param instrument_id path string true "instrument uuid" Format(uuid)
//	@Param timeseries_collection_items body model.TimeseriesCollectionItems true "timeseries collection items payload"
//	@Param key query string false "api key"
//	@Success 200 {array} model.Timeseries
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/instruments/{instrument_id}/constants [post]
//	@Security Bearer
func (h *ApiHandler) CreateInstrumentConstants(c echo.Context) error {
	ctx := c.Request().Context()
	var tc model.TimeseriesCollectionItems
	if err := c.Bind(&tc); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	instrumentID, err := uuid.Parse(c.Param("instrument_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	for idx := range tc.Items {
		if instrumentID != tc.Items[idx].InstrumentID {
			return echo.NewHTTPError(http.StatusBadRequest, message.MatchRouteParam("`instrument_id`"))
		}
	}
	tt, err := h.InstrumentConstantService.CreateInstrumentConstants(ctx, tc.Items)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, tt)
}

// DeleteInstrumentConstant godoc
//
//	@Summary removes a timeseries as an instrument constant
//	@Tags instrument-constant
//	@Produce json
//	@Param project_id path string true "project uuid" Format(uuid)
//	@Param instrument_id path string true "instrument uuid" Format(uuid)
//	@Param timeseries_id path string true "timeseries uuid" Format(uuid)
//	@Param key query string false "api key"
//	@Success 200 {object} map[string]interface{}
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/instruments/{instrument_id}/constants/{timeseries_id} [delete]
//	@Security Bearer
func (h *ApiHandler) DeleteInstrumentConstant(c echo.Context) error {
	instrumentID, err := uuid.Parse(c.Param("instrument_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	timeseriesID, err := uuid.Parse(c.Param("timeseries_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err = h.InstrumentConstantService.DeleteInstrumentConstant(c.Request().Context(), instrumentID, timeseriesID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, make(map[string]interface{}))
}
