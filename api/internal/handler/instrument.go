package handler

import (
	"net/http"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/message"
	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/USACE/instrumentation-api/api/internal/util"
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
//	@Success 200 {array} model.IDAndSlug
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/instruments [post]
//	@Router /instruments [post]
//	@Security Bearer
func (h *ApiHandler) CreateInstruments(c echo.Context) error {
	ctx := c.Request().Context()
	newInstrumentCollection := func(c echo.Context) (model.InstrumentCollection, error) {
		ic := model.InstrumentCollection{}
		if err := c.Bind(&ic); err != nil {
			return model.InstrumentCollection{}, err
		}

		// Get ProjectID of Instruments
		projectID, err := uuid.Parse(c.Param("project_id"))
		if err != nil {
			return model.InstrumentCollection{}, err
		}

		// slugs already taken in the database
		slugsTaken, err := h.InstrumentService.ListInstrumentSlugs(ctx)
		if err != nil {
			return model.InstrumentCollection{}, err
		}

		// profile of user creating instruments
		p := c.Get("profile").(model.Profile)

		// timestamp
		t := time.Now()

		for idx := range ic.Items {
			ic.Items[idx].ProjectID = &projectID
			s, err := util.NextUniqueSlug(ic.Items[idx].Name, slugsTaken)
			if err != nil {
				return model.InstrumentCollection{}, err
			}
			ic.Items[idx].Slug = s
			ic.Items[idx].Creator = p.ID
			ic.Items[idx].CreateDate = t
			slugsTaken = append(slugsTaken, s)
		}

		return ic, nil
	}

	ic, err := newInstrumentCollection(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if c.QueryParam("dry_run") == "true" {
		v, err := h.InstrumentService.ValidateCreateInstruments(ctx, ic.Items)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, v)
	}

	nn, err := h.InstrumentService.CreateInstruments(ctx, ic.Items)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
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
//	@Success 200 {object} model.Instrument
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/instruments/{instrument_id} [put]
//	@Security Bearer
func (h *ApiHandler) UpdateInstrument(c echo.Context) error {

	// instrument_id from url params
	iID, err := uuid.Parse(c.Param("instrument_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}
	// project_id from url params
	pID, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}
	// instrument from request payload
	i := model.Instrument{ID: iID, ProjectID: &pID}
	if err := c.Bind(&i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// check project_id in url params matches project_id in request body
	if pID != *i.ProjectID {
		return echo.NewHTTPError(http.StatusBadRequest, message.MatchRouteParam("`project_id`"))
	}
	// check instrument_id in url params matches instrument_id in request body
	if iID != i.ID {
		return echo.NewHTTPError(http.StatusBadRequest, message.MatchRouteParam("`instrument_id`"))
	}

	// profile of user creating instruments
	p := c.Get("profile").(model.Profile)

	// timestamp
	t := time.Now()
	i.Updater, i.UpdateDate = &p.ID, &t

	// update
	iUpdated, err := h.InstrumentService.UpdateInstrument(c.Request().Context(), i)
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
		return echo.NewHTTPError(http.StatusBadRequest, message.BadRequest)
	}

	return c.JSON(http.StatusOK, make(map[string]interface{}))
}
