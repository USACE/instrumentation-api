package handler

import (
	"net/http"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/messages"
	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/USACE/instrumentation-api/api/internal/util"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// ListInstrumentGroups returns instrument groups
func (h ApiHandler) ListInstrumentGroups(c echo.Context) error {
	groups, err := h.InstrumentGroupStore.ListInstrumentGroups(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, groups)
}

// GetInstrumentGroup returns single instrument group
func (h ApiHandler) GetInstrumentGroup(c echo.Context) error {
	id, err := uuid.Parse(c.Param("instrument_group_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, messages.MalformedID)
	}
	g, err := h.InstrumentGroupStore.GetInstrumentGroup(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, g)
}

// CreateInstrumentGroup accepts an array of instruments for bulk upload to the database
func (h ApiHandler) CreateInstrumentGroup(c echo.Context) error {

	gc := model.InstrumentGroupCollection{}
	if err := c.Bind(&gc); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// slugs already taken in the database
	slugsTaken, err := h.InstrumentGroupStore.ListInstrumentGroupSlugs(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// profile
	p := c.Get("profile").(*model.Profile)

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

	gg, err := h.InstrumentGroupStore.CreateInstrumentGroup(c.Request().Context(), gc.Items)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// Send instrumentgroup
	return c.JSON(http.StatusCreated, gg)
}

// UpdateInstrumentGroup modifies an existing instrument_group
func (h ApiHandler) UpdateInstrumentGroup(c echo.Context) error {

	// id from url params
	id, err := uuid.Parse(c.Param("instrument_group_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, messages.MalformedID)
	}
	// id from request
	g := model.InstrumentGroup{ID: id}
	if err := c.Bind(&g); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// check :id in url params matches id in request body
	if id != g.ID {
		return echo.NewHTTPError(http.StatusBadRequest, messages.MatchRouteParam("`id`"))
	}

	// profile information and timestamp
	p := c.Get("profile").(*model.Profile)

	t := time.Now()
	g.Updater, g.UpdateDate = &p.ID, &t

	// update
	gUpdated, err := h.InstrumentGroupStore.UpdateInstrumentGroup(c.Request().Context(), &g)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// return updated instrument
	return c.JSON(http.StatusOK, gUpdated)
}

// DeleteFlagInstrumentGroup sets the instrument group deleted flag true
func (h ApiHandler) DeleteFlagInstrumentGroup(c echo.Context) error {
	id, err := uuid.Parse(c.Param("instrument_group_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	if err := h.InstrumentGroupStore.DeleteFlagInstrumentGroup(c.Request().Context(), id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, make(map[string]interface{}))
}

// ListInstrumentGroupInstruments returns a list of instruments for a provided instrument group
func (h ApiHandler) ListInstrumentGroupInstruments(c echo.Context) error {
	id, err := uuid.Parse(c.Param("instrument_group_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, messages.MalformedID)
	}
	nn, err := h.InstrumentGroupStore.ListInstrumentGroupInstruments(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, nn)
}

// CreateInstrumentGroupInstruments adds an instrument to an instrument group
func (h ApiHandler) CreateInstrumentGroupInstruments(c echo.Context) error {
	// instrument_group_id
	instrumentGroupID, err := uuid.Parse(c.Param("instrument_group_id"))

	if err != nil || instrumentGroupID == uuid.Nil {
		return echo.NewHTTPError(http.StatusBadRequest, messages.BadRequest)
	}
	// Instrument
	i := new(model.Instrument)
	if err := c.Bind(i); err != nil || i.ID == uuid.Nil {
		return echo.NewHTTPError(http.StatusBadRequest, messages.BadRequest)
	}

	if err := h.InstrumentGroupStore.CreateInstrumentGroupInstruments(c.Request().Context(), instrumentGroupID, i.ID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, make(map[string]interface{}))
}

// DeleteInstrumentGroupInstruments removes an instrument from an instrument group
func (h ApiHandler) DeleteInstrumentGroupInstruments(c echo.Context) error {
	// instrument_group_id
	instrumentGroupID, err := uuid.Parse(c.Param("instrument_group_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, messages.MalformedID)
	}

	// instrument
	instrumentID, err := uuid.Parse(c.Param("instrument_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, messages.MalformedID)
	}

	if err := h.InstrumentGroupStore.DeleteInstrumentGroupInstruments(c.Request().Context(), instrumentGroupID, instrumentID); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, make(map[string]interface{}))
}
