package handler

import (
	"net/http"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/message"
	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/USACE/instrumentation-api/api/internal/util"

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
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
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
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}
	g, err := h.InstrumentGroupService.GetInstrumentGroup(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, g)
}

// CreateInstrumentGroup godoc
//
//	@Summary creats an instrument group from an array of instruments
//	@Tags instrument-group
//	@Produce json
//	@Param instrument_group body model.InstrumentGroup true "instrument group payload"
//	@Success 200 {object} model.InstrumentGroup
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /instrument_groups [post]
//	@Security Bearer
func (h *ApiHandler) CreateInstrumentGroup(c echo.Context) error {

	gc := model.InstrumentGroupCollection{}
	if err := c.Bind(&gc); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// slugs already taken in the database
	slugsTaken, err := h.InstrumentGroupService.ListInstrumentGroupSlugs(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// profile
	p := c.Get("profile").(model.Profile)

	// timestamp
	t := time.Now()

	for idx := range gc.Items {
		// Creator
		gc.Items[idx].Creator = p.ID
		// CreateDate
		gc.Items[idx].CreateDate = t
		// Assign Slug
		s, err := util.NextUniqueSlug(gc.Items[idx].Name, slugsTaken)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		gc.Items[idx].Slug = s
		// Add slug to array of slugs originally fetched from the database
		// to catch duplicate names/slugs from the same bulk upload
		slugsTaken = append(slugsTaken, s)
	}

	gg, err := h.InstrumentGroupService.CreateInstrumentGroup(c.Request().Context(), gc.Items)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// Send instrumentgroup
	return c.JSON(http.StatusCreated, gg)
}

// UpdateInstrumentGroup godoc
//
//	@Summary updates an existing instrument group
//	@Tags instrument-group
//	@Produce json
//	@Param instrument_group_id path string true "instrument group uuid" Format(uuid)
//	@Param instrument_group body model.InstrumentGroup true "instrument group payload"
//	@Success 200 {object} model.InstrumentGroup
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /instrument_groups/{instrument_group_id} [put]
//	@Security Bearer
func (h *ApiHandler) UpdateInstrumentGroup(c echo.Context) error {

	// id from url params
	id, err := uuid.Parse(c.Param("instrument_group_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}
	// id from request
	g := model.InstrumentGroup{ID: id}
	if err := c.Bind(&g); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// check :id in url params matches id in request body
	if id != g.ID {
		return echo.NewHTTPError(http.StatusBadRequest, message.MatchRouteParam("`id`"))
	}

	// profile information and timestamp
	p := c.Get("profile").(model.Profile)

	t := time.Now()
	g.Updater, g.UpdateDate = &p.ID, &t

	// update
	gUpdated, err := h.InstrumentGroupService.UpdateInstrumentGroup(c.Request().Context(), g)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// return updated instrument
	return c.JSON(http.StatusOK, gUpdated)
}

// DeleteFlagInstrumentGroup godoc
//
//	@Summary soft deletes an instrument
//	@Tags instrument-group
//	@Produce json
//	@Param instrument_group_id path string true "instrument group uuid" Format(uuid)
//	@Success 200 {array} model.InstrumentGroup
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /instrument_groups/{instrument_group_id} [delete]
//	@Security Bearer
func (h *ApiHandler) DeleteFlagInstrumentGroup(c echo.Context) error {
	id, err := uuid.Parse(c.Param("instrument_group_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	if err := h.InstrumentGroupService.DeleteFlagInstrumentGroup(c.Request().Context(), id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
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
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}
	nn, err := h.InstrumentGroupService.ListInstrumentGroupInstruments(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, nn)
}

// CreateInstrumentGroupInstruments godoc
//
//	@Summary adds an instrument to an instrument group
//	@Tags instrument-group
//	@Produce json
//	@Param instrument_group_id path string true "instrument group uuid" Format(uuid)
//	@Success 200 {object} map[string]interface{}
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /instrument_groups/{instrument_group_id}/instruments [post]
//	@Security Bearer
func (h *ApiHandler) CreateInstrumentGroupInstruments(c echo.Context) error {
	instrumentGroupID, err := uuid.Parse(c.Param("instrument_group_id"))
	if err != nil || instrumentGroupID == uuid.Nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.BadRequest)
	}
	i := new(model.Instrument)
	if err := c.Bind(i); err != nil || i.ID == uuid.Nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.BadRequest)
	}
	if err := h.InstrumentGroupService.CreateInstrumentGroupInstruments(c.Request().Context(), instrumentGroupID, i.ID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
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
//	@Success 200 {object} map[string]interface{}
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /instrument_groups/{instrument_group_id}/instruments/{instrument_id} [delete]
//	@Security Bearer
func (h *ApiHandler) DeleteInstrumentGroupInstruments(c echo.Context) error {
	// instrument_group_id
	instrumentGroupID, err := uuid.Parse(c.Param("instrument_group_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}

	// instrument
	instrumentID, err := uuid.Parse(c.Param("instrument_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}

	if err := h.InstrumentGroupService.DeleteInstrumentGroupInstruments(c.Request().Context(), instrumentGroupID, instrumentID); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, make(map[string]interface{}))
}
