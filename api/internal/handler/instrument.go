package handler

import (
	"net/http"
	"strings"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/httperr"
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
		return httperr.InternalServerError(err)
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
		return httperr.InternalServerError(err)
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
		return httperr.MalformedID(err)
	}
	n, err := h.InstrumentService.GetInstrument(c.Request().Context(), id)
	if err != nil {
		return httperr.ServerErrorOrNotFound(err)
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
//	@Param key query string false "api key"
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
		return httperr.MalformedID(err)
	}

	ic := model.InstrumentCollection{}
	if err := c.Bind(&ic); err != nil {
		return httperr.MalformedBody(err)
	}

	p := c.Get("profile").(model.Profile)
	t := time.Now()

	instrumentNames := make([]string, len(ic))
	for idx := range ic {
		instrumentNames[idx] = ic[idx].Name
		var prj model.IDSlugName
		prj.ID = projectID
		ic[idx].Projects = []model.IDSlugName{prj}
		ic[idx].CreatorID = p.ID
		ic[idx].CreateDate = t
	}

	if strings.ToLower(c.QueryParam("dry_run")) == "true" {
		v, err := h.InstrumentAssignService.ValidateInstrumentNamesProjectUnique(ctx, projectID, instrumentNames)
		if err != nil {
			return httperr.InternalServerError(err)
		}
		if !v.IsValid {
			return c.JSON(http.StatusBadRequest, v)
		}
		return c.JSON(http.StatusOK, v)
	}

	nn, err := h.InstrumentService.CreateInstruments(ctx, ic)
	if err != nil {
		return httperr.InternalServerError(err)
	}

	return c.JSON(http.StatusCreated, nn)
}

// UpdateInstrument godoc
//
//	@Summary updates an existing instrument
//	@Tags instrument
//	@Produce json
//	@Param project_id path string true "project uuid" Format(uuid)
//	@Param instrument_id path string true "instrument uuid" Format(uuid)
//	@Param instrument body model.Instrument true "instrument payload"
//	@Param key query string false "api key"
//	@Success 200 {object} model.Instrument
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/instruments/{instrument_id} [put]
//	@Security Bearer
func (h *ApiHandler) UpdateInstrument(c echo.Context) error {
	iID, err := uuid.Parse(c.Param("instrument_id"))
	if err != nil {
		return httperr.MalformedID(err)
	}
	pID, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return httperr.MalformedID(err)
	}

	var i model.Instrument
	if err := c.Bind(&i); err != nil {
		return httperr.MalformedBody(err)
	}
	i.ID = iID

	p := c.Get("profile").(model.Profile)

	t := time.Now()
	i.UpdaterID, i.UpdateDate = &p.ID, &t

	// update
	iUpdated, err := h.InstrumentService.UpdateInstrument(c.Request().Context(), pID, i)
	if err != nil {
		return httperr.InternalServerError(err)
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
//	@Param key query string false "api key"
//	@Success 200 {object} model.Instrument
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/instruments/{instrument_id}/geometry [put]
//	@Security Bearer
func (h *ApiHandler) UpdateInstrumentGeometry(c echo.Context) error {
	projectID, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return httperr.MalformedID(err)
	}
	instrumentID, err := uuid.Parse(c.Param("instrument_id"))
	if err != nil {
		return httperr.MalformedID(err)
	}
	var geom geojson.Geometry
	if err := c.Bind(&geom); err != nil {
		return httperr.MalformedBody(err)
	}
	// profile of user creating instruments
	p := c.Get("profile").(model.Profile)

	instrument, err := h.InstrumentService.UpdateInstrumentGeometry(c.Request().Context(), projectID, instrumentID, geom, p)
	if err != nil {
		return httperr.InternalServerError(err)
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
//	@Param key query string false "api key"
//	@Success 200 {object} map[string]interface{}
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/instruments/{instrument_id} [delete]
//	@Security Bearer
func (h *ApiHandler) DeleteFlagInstrument(c echo.Context) error {
	pID, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return httperr.MalformedID(err)
	}

	iID, err := uuid.Parse(c.Param("instrument_id"))
	if err != nil {
		return httperr.MalformedID(err)
	}

	if err := h.InstrumentService.DeleteFlagInstrument(c.Request().Context(), pID, iID); err != nil {
		return httperr.InternalServerError(err)
	}

	return c.JSON(http.StatusOK, make(map[string]interface{}))
}
