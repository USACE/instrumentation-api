package handler

import (
	"net/http"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/httperr"
	"github.com/USACE/instrumentation-api/api/internal/model"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// ListInstrumentGroups godoc
//
//	@Summary lists all instrument groups
//	@Tags instrument-group
//	@Produce json
//	@Success 200 {array} model.InstrumentGroup
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /instrument_groups [get]
func (h *ApiHandler) ListInstrumentGroups(c echo.Context) error {
	groups, err := h.InstrumentGroupService.ListInstrumentGroups(c.Request().Context())
	if err != nil {
		return httperr.InternalServerError(err)
	}
	return c.JSON(http.StatusOK, groups)
}

// GetInstrumentGroup godoc
//
//	@Summary gets a single instrument group
//	@Tags instrument-group
//	@Produce json
//	@Param instrument_group_id path string true "instrument group uuid" Format(uuid)
//	@Success 200 {object} model.InstrumentGroup
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /instrument_groups/{instrument_group_id} [get]
func (h *ApiHandler) GetInstrumentGroup(c echo.Context) error {
	id, err := uuid.Parse(c.Param("instrument_group_id"))
	if err != nil {
		return httperr.MalformedID(err)
	}
	g, err := h.InstrumentGroupService.GetInstrumentGroup(c.Request().Context(), id)
	if err != nil {
		return httperr.InternalServerError(err)
	}
	return c.JSON(http.StatusOK, g)
}

// CreateInstrumentGroup godoc
//
//	@Summary creats an instrument group from an array of instruments
//	@Tags instrument-group
//	@Produce json
//	@Param instrument_group body model.InstrumentGroup true "instrument group payload"
//	@Param key query string false "api key"
//	@Success 201 {object} model.InstrumentGroup
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /instrument_groups [post]
//	@Security Bearer
func (h *ApiHandler) CreateInstrumentGroup(c echo.Context) error {

	gc := model.InstrumentGroupCollection{}
	if err := c.Bind(&gc); err != nil {
		return httperr.MalformedBody(err)
	}

	p := c.Get("profile").(model.Profile)
	t := time.Now()

	for idx := range gc.Items {
		gc.Items[idx].CreatorID = p.ID
		gc.Items[idx].CreateDate = t
	}

	gg, err := h.InstrumentGroupService.CreateInstrumentGroup(c.Request().Context(), gc.Items)
	if err != nil {
		return httperr.InternalServerError(err)
	}
	return c.JSON(http.StatusCreated, gg)
}

// UpdateInstrumentGroup godoc
//
//	@Summary updates an existing instrument group
//	@Tags instrument-group
//	@Produce json
//	@Param instrument_group_id path string true "instrument group uuid" Format(uuid)
//	@Param instrument_group body model.InstrumentGroup true "instrument group payload"
//	@Param key query string false "api key"
//	@Success 200 {object} model.InstrumentGroup
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /instrument_groups/{instrument_group_id} [put]
//	@Security Bearer
func (h *ApiHandler) UpdateInstrumentGroup(c echo.Context) error {
	gID, err := uuid.Parse(c.Param("instrument_group_id"))
	if err != nil {
		return httperr.MalformedID(err)
	}

	g := model.InstrumentGroup{ID: gID}
	if err := c.Bind(&g); err != nil {
		return httperr.MalformedBody(err)
	}
	g.ID = gID

	p := c.Get("profile").(model.Profile)

	t := time.Now()
	g.UpdaterID, g.UpdateDate = &p.ID, &t

	gUpdated, err := h.InstrumentGroupService.UpdateInstrumentGroup(c.Request().Context(), g)
	if err != nil {
		return httperr.InternalServerError(err)
	}
	return c.JSON(http.StatusOK, gUpdated)
}

// DeleteFlagInstrumentGroup godoc
//
//	@Summary soft deletes an instrument
//	@Tags instrument-group
//	@Produce json
//	@Param instrument_group_id path string true "instrument group uuid" Format(uuid)
//	@Param key query string false "api key"
//	@Success 200 {array} model.InstrumentGroup
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /instrument_groups/{instrument_group_id} [delete]
//	@Security Bearer
func (h *ApiHandler) DeleteFlagInstrumentGroup(c echo.Context) error {
	id, err := uuid.Parse(c.Param("instrument_group_id"))
	if err != nil {
		return httperr.MalformedID(err)
	}
	if err := h.InstrumentGroupService.DeleteFlagInstrumentGroup(c.Request().Context(), id); err != nil {
		return httperr.InternalServerError(err)
	}
	return c.JSON(http.StatusOK, make(map[string]interface{}))
}

// ListInstrumentGroupInstruments godoc
//
//	@Summary lists instruments in an instrument group
//	@Tags instrument-group
//	@Produce json
//	@Param instrument_group_id path string true "instrument group uuid" Format(uuid)
//	@Success 200 {array} model.Instrument
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /instrument_groups/{instrument_group_id}/instruments [get]
func (h *ApiHandler) ListInstrumentGroupInstruments(c echo.Context) error {
	id, err := uuid.Parse(c.Param("instrument_group_id"))
	if err != nil {
		return httperr.MalformedID(err)
	}
	nn, err := h.InstrumentGroupService.ListInstrumentGroupInstruments(c.Request().Context(), id)
	if err != nil {
		return httperr.InternalServerError(err)
	}
	return c.JSON(http.StatusOK, nn)
}

// CreateInstrumentGroupInstruments godoc
//
//	@Summary adds an instrument to an instrument group
//	@Tags instrument-group
//	@Produce json
//	@Param instrument_group_id path string true "instrument group uuid" Format(uuid)
//	@Param key query string false "api key"
//	@Success 200 {object} map[string]interface{}
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /instrument_groups/{instrument_group_id}/instruments [post]
//	@Security Bearer
func (h *ApiHandler) CreateInstrumentGroupInstruments(c echo.Context) error {
	instrumentGroupID, err := uuid.Parse(c.Param("instrument_group_id"))
	if err != nil || instrumentGroupID == uuid.Nil {
		return httperr.MalformedID(err)
	}
	var i model.Instrument
	if err := c.Bind(&i); err != nil || i.ID == uuid.Nil {
		return httperr.MalformedBody(err)
	}
	if err := h.InstrumentGroupService.CreateInstrumentGroupInstruments(c.Request().Context(), instrumentGroupID, i.ID); err != nil {
		return httperr.InternalServerError(err)
	}
	return c.JSON(http.StatusCreated, make(map[string]interface{}))
}

// DeleteInstrumentGroupInstruments godoc
//
//	@Summary removes an instrument from an instrument group
//	@Tags instrument-group
//	@Produce json
//	@Param instrument_group_id path string true "instrument group uuid" Format(uuid)
//	@Param instrument_id path string true "instrument uuid" Format(uuid)
//	@Param key query string false "api key"
//	@Success 200 {object} map[string]interface{}
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /instrument_groups/{instrument_group_id}/instruments/{instrument_id} [delete]
//	@Security Bearer
func (h *ApiHandler) DeleteInstrumentGroupInstruments(c echo.Context) error {
	instrumentGroupID, err := uuid.Parse(c.Param("instrument_group_id"))
	if err != nil {
		return httperr.MalformedID(err)
	}

	instrumentID, err := uuid.Parse(c.Param("instrument_id"))
	if err != nil {
		return httperr.MalformedID(err)
	}

	if err := h.InstrumentGroupService.DeleteInstrumentGroupInstruments(c.Request().Context(), instrumentGroupID, instrumentID); err != nil {
		return httperr.InternalServerError(err)
	}

	return c.JSON(http.StatusOK, make(map[string]interface{}))
}
