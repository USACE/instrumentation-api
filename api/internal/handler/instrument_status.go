package handler

import (
	"net/http"

	"github.com/USACE/instrumentation-api/api/internal/message"
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
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}

	ss, err := h.InstrumentStatusService.ListInstrumentStatus(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
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
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}

	s, err := h.InstrumentStatusService.GetInstrumentStatus(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
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
//	@Success 200 {object} map[string]interface{}
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /instruments/{instrument_id}/status [post]
func (h *ApiHandler) CreateOrUpdateInstrumentStatus(c echo.Context) error {
	instrumentID, err := uuid.Parse(c.Param("instrument_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}

	var sc model.InstrumentStatusCollection
	if err := c.Bind(&sc); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// Assign Fresh UUID to each Status
	for idx := range sc.Items {
		id, err := uuid.NewRandom()
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		sc.Items[idx].ID = id
	}

	if err := h.InstrumentStatusService.CreateOrUpdateInstrumentStatus(c.Request().Context(), instrumentID, sc.Items); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
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
//	@Success 200 {object} map[string]interface{}
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /instruments/{instrument_id}/status/{status_id} [delete]
func (h *ApiHandler) DeleteInstrumentStatus(c echo.Context) error {
	id, err := uuid.Parse(c.Param("status_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}
	if err := h.InstrumentStatusService.DeleteInstrumentStatus(c.Request().Context(), id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, make(map[string]interface{}))
}
