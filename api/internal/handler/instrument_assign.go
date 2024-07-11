package handler

import (
	"net/http"
	"strings"

	"github.com/USACE/instrumentation-api/api/internal/httperr"
	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// AssignInstrumentToProject godoc
//
//	@Summary assigns an instrument to a project.
//	@Tags instrument
//	@Description must be Project (or Application) Admin of all existing instrument projects and project to be assigned
//	@Produce json
//	@Param project_id path string true "project uuid" Format(uuid)
//	@Param instrument_id path string true "instrument uuid" Format(uuid)
//	@Param dry_run query string false "validate request without performing action"
//	@Success 200 {object} model.InstrumentsValidation
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/instruments/{instrument_id}/assignments [post]
//	@Security Bearer
func (h *ApiHandler) AssignInstrumentToProject(c echo.Context) error {
	pID, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return httperr.MalformedID(err)
	}
	iID, err := uuid.Parse(c.Param("instrument_id"))
	if err != nil {
		return httperr.MalformedID(err)
	}
	dryRun := strings.ToLower(c.QueryParam("dry_run")) == "true"
	p := c.Get("profile").(model.Profile)

	v, err := h.InstrumentAssignService.AssignInstrumentsToProject(c.Request().Context(), p.ID, pID, []uuid.UUID{iID}, dryRun)
	if err != nil {
		return httperr.InternalServerError(err)
	}

	return c.JSON(http.StatusOK, v)
}

// UnassignInstrumentFromProject godoc
//
//	@Summary unassigns an instrument from a project.
//	@Tags instrument
//	@Description must be Project Admin of project to be unassigned
//	@Produce json
//	@Param project_id path string true "project uuid" Format(uuid)
//	@Param instrument_id path string true "instrument uuid" Format(uuid)
//	@Param action query string true "valid values are 'assign' or 'unassign'"
//	@Param dry_run query string false "validate request without performing action"
//	@Success 200 {object} model.InstrumentsValidation
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/instruments/{instrument_id}/assignments [delete]
//	@Security Bearer
func (h *ApiHandler) UnassignInstrumentFromProject(c echo.Context) error {
	pID, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return httperr.MalformedID(err)
	}
	iID, err := uuid.Parse(c.Param("instrument_id"))
	if err != nil {
		return httperr.MalformedID(err)
	}
	dryRun := strings.ToLower(c.QueryParam("dry_run")) == "true"
	p := c.Get("profile").(model.Profile)

	v, err := h.InstrumentAssignService.UnassignInstrumentsFromProject(c.Request().Context(), p.ID, pID, []uuid.UUID{iID}, dryRun)
	if err != nil {
		return httperr.InternalServerError(err)
	}

	return c.JSON(http.StatusOK, v)
}

// UpdateInstrumentProjectAssignments godoc
//
//	@Summary updates multiple project assignments for an instrument
//	@Tags instrument
//	@Description must be Project (or Application) Admin of all existing instrument projects and project to be assigned
//	@Produce json
//	@Param project_id path string true "project uuid" Format(uuid)
//	@Param instrument_id path string true "instrument uuid" Format(uuid)
//	@Param project_ids body model.InstrumentProjectAssignments true "project uuids"
//	@Param action query string true "valid values are 'assign' or 'unassign'"
//	@Param dry_run query string false "validate request without performing action"
//	@Success 200 {object} model.InstrumentsValidation
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/instruments/{instrument_id}/assignments [put]
//	@Security Bearer
func (h *ApiHandler) UpdateInstrumentProjectAssignments(c echo.Context) error {
	iID, err := uuid.Parse(c.Param("instrument_id"))
	if err != nil {
		return httperr.MalformedID(err)
	}
	p := c.Get("profile").(model.Profile)
	dryRun := strings.ToLower(c.QueryParam("dry_run")) == "true"

	pl := model.InstrumentProjectAssignments{ProjectIDs: make([]uuid.UUID, 0)}
	if err := c.Bind(&pl); err != nil {
		return httperr.MalformedBody(err)
	}

	ctx := c.Request().Context()
	switch strings.ToLower(c.QueryParam("action")) {
	case "assign":
		v, err := h.InstrumentAssignService.AssignProjectsToInstrument(ctx, p.ID, iID, pl.ProjectIDs, dryRun)
		if err != nil {
			return httperr.InternalServerError(err)
		}
		return c.JSON(http.StatusOK, v)
	case "unassign":
		v, err := h.InstrumentAssignService.UnassignProjectsFromInstrument(ctx, p.ID, iID, pl.ProjectIDs, dryRun)
		if err != nil {
			return httperr.InternalServerError(err)
		}
		return c.JSON(http.StatusOK, v)
	default:
		return httperr.Message(http.StatusBadRequest, "required query parameter 'action': valid values are 'assign' or 'unassign'")
	}
}

// UpdateProjectInstrumentAssignments godoc
//
//	@Summary updates multiple instrument assigments for a project
//	@Tags instrument
//	@Description must be Project (or Application) Admin of all existing instrument projects and project to be assigned
//	@Produce json
//	@Param project_id path string true "project uuid" Format(uuid)
//	@Param instrument_ids body model.ProjectInstrumentAssignments true "instrument uuids"
//	@Param action query string true "valid values are 'assign' or 'unassign'"
//	@Param dry_run query string false "validate request without performing action"
//	@Success 200 {object} model.InstrumentsValidation
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/instruments/assignments [put]
//	@Security Bearer
func (h *ApiHandler) UpdateProjectInstrumentAssignments(c echo.Context) error {
	pID, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return httperr.MalformedID(err)
	}
	p := c.Get("profile").(model.Profile)
	dryRun := strings.ToLower(c.QueryParam("dry_run")) == "true"

	pl := model.ProjectInstrumentAssignments{InstrumentIDs: make([]uuid.UUID, 0)}
	if err := c.Bind(&pl); err != nil {
		return httperr.MalformedBody(err)
	}

	ctx := c.Request().Context()
	switch strings.ToLower(c.QueryParam("action")) {
	case "assign":
		v, err := h.InstrumentAssignService.AssignInstrumentsToProject(ctx, p.ID, pID, pl.InstrumentIDs, dryRun)
		if err != nil {
			return httperr.InternalServerError(err)
		}
		return c.JSON(http.StatusOK, v)
	case "unassign":
		v, err := h.InstrumentAssignService.UnassignInstrumentsFromProject(ctx, p.ID, pID, pl.InstrumentIDs, dryRun)
		if err != nil {
			return httperr.InternalServerError(err)
		}
		return c.JSON(http.StatusOK, v)
	default:
		return httperr.Message(http.StatusBadRequest, "required query parameter 'action': valid values are 'assign' or 'unassign'")
	}
}
