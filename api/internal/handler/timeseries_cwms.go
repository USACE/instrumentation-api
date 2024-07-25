package handler

import (
	"net/http"

	"github.com/USACE/instrumentation-api/api/internal/httperr"
	"github.com/USACE/instrumentation-api/api/internal/model"
	_ "github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// ListTimeseriesCwmsMeasurements(ctx context.Context, timeseriesID uuid.UUID) (model.MeasurementCollection, error)

// CreateTimeseriesCwms godoc
//
//	@Summary creates cwms timeseries
//	@Tags timeseries-cwms
//	@Produce json
//	@Param project_id path string true "project uuid" Format(uuid)
//	@Param instrument_id path string true "instrument uuid" Format(uuid)
//	@Param timeseries_cwms_arr body []model.TimeseriesCwms true "array of cwms timeseries to create"
//	@Success 200 {array} model.TimeseriesCwms
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/instruments/{instrument_id}/timeseries/cwms [post]
func (h *ApiHandler) CreateTimeseriesCwms(c echo.Context) error {
	_, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return httperr.MalformedID(err)
	}
	instrumentID, err := uuid.Parse(c.Param("instrument_id"))
	if err != nil {
		return httperr.MalformedID(err)
	}

	var tcc []model.TimeseriesCwms
	if err := c.Bind(&tcc); err != nil {
		return httperr.MalformedBody(err)
	}

	tss, err := h.TimeseriesCwmsService.CreateTimeseriesCwmsBatch(c.Request().Context(), instrumentID, tcc)
	if err != nil {
		return httperr.InternalServerError(err)
	}

	return c.JSON(http.StatusCreated, tss)
}

// UpdateTimeseriesCwms godoc
//
//	@Summary updates cwms timeseries
//	@Tags timeseries-cwms
//	@Produce json
//	@Param project_id path string true "project uuid" Format(uuid)
//	@Param instrument_id path string true "instrument uuid" Format(uuid)
//	@Param timeseries_id path string true "timeseries uuid" Format(uuid)
//	@Param timeseries_cwms body model.TimeseriesCwms true "cwms timeseries to update"
//	@Success 200 {array} model.TimeseriesCwms
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/instruments/{instrument_id}/timeseries/cwms/{timeseries_id} [put]
func (h *ApiHandler) UpdateTimeseriesCwms(c echo.Context) error {
	_, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return httperr.MalformedID(err)
	}
	instrumentID, err := uuid.Parse(c.Param("instrument_id"))
	if err != nil {
		return httperr.MalformedID(err)
	}
	timeseriesID, err := uuid.Parse(c.Param("timeseries_id"))
	if err != nil {
		return httperr.MalformedID(err)
	}

	var tc model.TimeseriesCwms
	if err := c.Bind(&tc); err != nil {
		return httperr.MalformedBody(err)
	}
	tc.InstrumentID = instrumentID
	tc.ID = timeseriesID

	if err := h.TimeseriesCwmsService.UpdateTimeseriesCwms(c.Request().Context(), tc); err != nil {
		return httperr.InternalServerError(err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"id": tc.ID})
}
