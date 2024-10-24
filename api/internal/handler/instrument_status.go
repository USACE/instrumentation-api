package handler

import (
	"net/http"

	"github.com/USACE/instrumentation-api/api/internal/httperr"
	"github.com/USACE/instrumentation-api/api/internal/model"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// ListInstrumentStatus godoc
//
//	@Summary lists all Status for an instrument
//	@Tags instrument-status
//	@Produce json
//	@Param instrument_id path string true "instrument uuid" Format(uuid)
//	@Success 200 {array} model.InstrumentStatus
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /instruments/{instrument_id}/status [get]
func (h *ApiHandler) ListInstrumentStatus(c echo.Context) error {
	id, err := uuid.Parse(c.Param("instrument_id"))
	if err != nil {
		return httperr.MalformedID(err)
	}

	ss, err := h.InstrumentStatusService.ListInstrumentStatus(c.Request().Context(), id)
	if err != nil {
		return httperr.InternalServerError(err)
	}
	return c.JSON(http.StatusOK, ss)
}

// GetInstrumentStatus godoc
//
//	@Summary gets a single status
//	@Tags instrument-status
//	@Produce json
//	@Param instrument_id path string true "instrument uuid" Format(uuid)
//	@Param status_id path string true "status uuid" Format(uuid)
//	@Success 200 {array} model.AlertConfig
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /instruments/{instrument_id}/status/{status_id} [get]
func (h *ApiHandler) GetInstrumentStatus(c echo.Context) error {
	id, err := uuid.Parse(c.Param("status_id"))
	if err != nil {
		return httperr.MalformedID(err)
	}

	s, err := h.InstrumentStatusService.GetInstrumentStatus(c.Request().Context(), id)
	if err != nil {
		return httperr.ServerErrorOrNotFound(err)
	}
	return c.JSON(http.StatusOK, s)
}

// CreateOrUpdateInstrumentStatus godoc
//
//	@Summary creates a status for an instrument
//	@Tags instrument-status
//	@Produce json
//	@Param instrument_id path string true "instrument uuid" Format(uuid)
//	@Param instrument_status body model.InstrumentStatusCollection true "instrument status collection paylaod"
//	@Param key query string false "api key"
//	@Success 200 {object} map[string]interface{}
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /instruments/{instrument_id}/status [post]
//	@Security Bearer
func (h *ApiHandler) CreateOrUpdateInstrumentStatus(c echo.Context) error {
	instrumentID, err := uuid.Parse(c.Param("instrument_id"))
	if err != nil {
		return httperr.MalformedID(err)
	}

	var sc model.InstrumentStatusCollection
	if err := c.Bind(&sc); err != nil {
		return httperr.MalformedBody(err)
	}
	for idx := range sc.Items {
		id, err := uuid.NewRandom()
		if err != nil {
			return httperr.InternalServerError(err)
		}
		sc.Items[idx].ID = id
	}

	if err := h.InstrumentStatusService.CreateOrUpdateInstrumentStatus(c.Request().Context(), instrumentID, sc.Items); err != nil {
		return httperr.InternalServerError(err)
	}
	return c.JSON(http.StatusCreated, make(map[string]interface{}))
}

// DeleteInstrumentStatus godoc
//
//	@Summary deletes a status for an instrument
//	@Tags instrument-status
//	@Produce json
//	@Param instrument_id path string true "instrument uuid" Format(uuid)
//	@Param status_id path string true "project uuid" Format(uuid)
//	@Param key query string false "api key"
//	@Success 200 {object} map[string]interface{}
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /instruments/{instrument_id}/status/{status_id} [delete]
//	@Security Bearer
func (h *ApiHandler) DeleteInstrumentStatus(c echo.Context) error {
	id, err := uuid.Parse(c.Param("status_id"))
	if err != nil {
		return httperr.MalformedID(err)
	}
	if err := h.InstrumentStatusService.DeleteInstrumentStatus(c.Request().Context(), id); err != nil {
		return httperr.InternalServerError(err)
	}
	return c.JSON(http.StatusOK, make(map[string]interface{}))
}
