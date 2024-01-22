package handler

import (
	"net/http"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/message"
	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/paulmach/orb/geojson"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// ListInstruments godoc
//
//	@Summary lists all instruments
//	@Tags instrument
//	@Produce json
//	@Success 200 {array} model.Instrument
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /instruments [get]
func (h *ApiHandler) ListInstruments(c echo.Context) error {
	nn, err := h.InstrumentService.ListInstruments(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, nn)
}

// GetInstrumentCount godoc
//
//	@Summary gets the total number of non deleted instruments in the system
//	@Tags instrument
//	@Produce json
//	@Success 200 {object} model.InstrumentCount
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /instruments/count [get]
func (h *ApiHandler) GetInstrumentCount(c echo.Context) error {
	ic, err := h.InstrumentService.GetInstrumentCount(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, ic)
}

// GetInstrument godoc
//
//	@Summary gets a single instrument by id
//	@Tags instrument
//	@Produce json
//	@Param instrument_id path string true "instrument uuid" Format(uuid)
//	@Success 200 {object} model.Instrument
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /instruments/{instrument_id} [get]
func (h *ApiHandler) GetInstrument(c echo.Context) error {
	id, err := uuid.Parse(c.Param("instrument_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}
	n, err := h.InstrumentService.GetInstrument(c.Request().Context(), id)
	if err != nil {
		if err.Error() == message.NotFound {
			return echo.NewHTTPError(http.StatusBadRequest, message.NotFound)
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, n)
}

// CreateInstruments godoc
//
//	@Summary accepts an array of instruments for bulk upload to the database
//	@Tags instrument
//	@Produce json
//	@Param project_id path string true "project id" Format(uuid)
//	@Param instrument_id path string true "instrument id" Format(uuid)
//	@Param instrument body model.InstrumentCollection true "instrument collection payload"
//	@Success 200 {array} model.IDSlugName
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/instruments [post]
//	@Security Bearer
func (h *ApiHandler) CreateInstruments(c echo.Context) error {
	ctx := c.Request().Context()

	projectID, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	ic := model.InstrumentCollection{}
	if err := c.Bind(&ic); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	p := c.Get("profile").(model.Profile)
	t := time.Now()

	for idx := range ic {
		var prj model.IDSlugName
		prj.ID = projectID
		ic[idx].Projects = []model.IDSlugName{prj}
		ic[idx].CreatorID = p.ID
		ic[idx].CreateDate = t
	}

	if c.QueryParam("dry_run") == "true" {
		v, err := h.InstrumentService.ValidateCreateInstruments(ctx, ic)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, v)
	}

	nn, err := h.InstrumentService.CreateInstruments(ctx, ic)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, nn)
}

// AssignInstrumentToProject godoc
//
//	@Summary assigns an instrument to a project.
//	@Tags instrument
//	@Description must be Project (or Application) Admin of all existing instrument projects and project to be assigned
//	@Produce json
//	@Param project_id path string true "project uuid" Format(uuid)
//	@Param instrument_id path string true "instrument uuid" Format(uuid)
//	@Success 200 {object} map[string]interface{}
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/instruments/{instrument_id}/assignments [post]
//	@Security Bearer
func (h *ApiHandler) AssignInstrumentToProject(c echo.Context) error {
	pID, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}
	iID, err := uuid.Parse(c.Param("instrument_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}

	ctx := c.Request().Context()
	p := c.Get("profile").(model.Profile)
	if !p.IsAdmin {
		autorized, err := h.InstrumentService.AssignerIsAuthorized(ctx, iID, p.ID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, message.InternalServerError)
		}
		if !autorized {
			return echo.NewHTTPError(http.StatusUnauthorized, message.Unauthorized)
		}
	}

	if err := h.InstrumentService.AssignInstrumentToProject(ctx, pID, iID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, make(map[string]interface{}))
}

// UnassignInstrumentFromProject godoc
//
//	@Summary unassigns an instrument from a project.
//	@Tags instrument
//	@Description must be Project Admin of project to be unassigned
//	@Produce json
//	@Param project_id path string true "project uuid" Format(uuid)
//	@Param instrument_id path string true "instrument uuid" Format(uuid)
//	@Success 200 {object} map[string]interface{}
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/instruments/{instrument_id}/assignments [delete]
//	@Security Bearer
func (h *ApiHandler) UnassignInstrumentFromProject(c echo.Context) error {
	pID, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}
	iID, err := uuid.Parse(c.Param("instrument_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}

	if err := h.InstrumentService.UnassignInstrumentFromProject(c.Request().Context(), pID, iID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, make(map[string]interface{}))
}

// UpdateInstrument godoc
//
//	@Summary updates an existing instrument
//	@Tags instrument
//	@Produce json
//	@Param project_id path string true "project uuid" Format(uuid)
//	@Param instrument_id path string true "instrument uuid" Format(uuid)
//	@Param instrument body model.Instrument true "instrument payload"
//	@Success 200 {object} model.Instrument
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/instruments/{instrument_id} [put]
//	@Security Bearer
func (h *ApiHandler) UpdateInstrument(c echo.Context) error {
	iID, err := uuid.Parse(c.Param("instrument_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}
	pID, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}

	var i model.Instrument
	if err := c.Bind(&i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	i.ID = iID

	p := c.Get("profile").(model.Profile)

	t := time.Now()
	i.UpdaterID, i.UpdateDate = &p.ID, &t

	// update
	iUpdated, err := h.InstrumentService.UpdateInstrument(c.Request().Context(), pID, i)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// return updated instrument
	return c.JSON(http.StatusOK, iUpdated)
}

// UpdateInstrumentGeometry godoc
//
//	@Summary updates the geometry of an instrument
//	@Tags instrument
//	@Produce json
//	@Param project_id path string true "project uuid" Format(uuid)
//	@Param instrument_id path string true "instrument uuid" Format(uuid)
//	@Param instrument body model.Instrument true "instrument payload"
//	@Success 200 {object} model.Instrument
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/instruments/{instrument_id}/geometry [put]
//	@Security Bearer
func (h *ApiHandler) UpdateInstrumentGeometry(c echo.Context) error {
	projectID, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	instrumentID, err := uuid.Parse(c.Param("instrument_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	var geom geojson.Geometry
	if err := c.Bind(&geom); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// profile of user creating instruments
	p, ok := c.Get("profile").(model.Profile)
	if !ok {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	instrument, err := h.InstrumentService.UpdateInstrumentGeometry(c.Request().Context(), projectID, instrumentID, geom, p)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, instrument)
}

// DeleteFlagInstrument godoc
//
//	@Summary soft deletes an instrument
//	@Tags instrument
//	@Produce json
//	@Param project_id path string true "project uuid" Format(uuid)
//	@Param instrument_id path string true "instrument uuid" Format(uuid)
//	@Success 200 {object} map[string]interface{}
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/instruments/{instrument_id} [delete]
//	@Security Bearer
func (h *ApiHandler) DeleteFlagInstrument(c echo.Context) error {
	pID, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}

	iID, err := uuid.Parse(c.Param("instrument_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}

	if err := h.InstrumentService.DeleteFlagInstrument(c.Request().Context(), pID, iID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, make(map[string]interface{}))
}
