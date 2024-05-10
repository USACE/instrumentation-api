package handler

import (
	"net/http"
	"strings"

	"github.com/USACE/instrumentation-api/api/internal/message"
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
//	@Success 200 {object} model.InstrumentsValidation
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
	dryRun := c.QueryParam("dry_run") == "true"
	p := c.Get("profile").(model.Profile)

	v, err := h.InstrumentAssignService.AssignInstrumentsToProject(c.Request().Context(), p.ID, pID, []uuid.UUID{iID}, dryRun)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
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
//	@Success 200 {object} model.InstrumentsValidation
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
	dryRun := strings.ToLower(c.QueryParam("dry_run")) == "true"
	p := c.Get("profile").(model.Profile)

	v, err := h.InstrumentAssignService.UnassignInstrumentsFromProject(c.Request().Context(), p.ID, pID, []uuid.UUID{iID}, dryRun)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, v)
}

type ProjectIDsPayload struct {
	ProjectIDs []uuid.UUID `json:"project_ids"`
}

// AssignProjectsToInstrument godoc
//
//	@Summary assigns multiple projects to an instruments
//	@Tags instrument
//	@Description must be Project (or Application) Admin of all existing instrument projects and project to be assigned
//	@Produce json
//	@Param instrument_id path string true "instrument uuid" Format(uuid)
//	@Param project_ids body  true "project uuids" Format(uuid)
//	@Success 200 {object} model.InstrumentsValidation
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /instruments/{instrument_id}/assignments [post]
//	@Security Bearer
func (h *ApiHandler) AssignProjectsToInstrument(c echo.Context) error {
	iID, err := uuid.Parse(c.Param("instrument_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}
	p := c.Get("profile").(model.Profile)
	dryRun := strings.ToLower(c.QueryParam("dry_run")) == "true"

	pl := ProjectIDsPayload{ProjectIDs: make([]uuid.UUID, 0)}
	if err := c.Bind(&pl); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	v, err := h.InstrumentAssignService.AssignProjectsToInstrument(c.Request().Context(), p.ID, iID, pl.ProjectIDs, dryRun)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, v)
}

// UnassignProjectsFromInstrument godoc
//
//	@Summary unassigns multiple projects from an instrument
//	@Tags instrument
//	@Description must be Project (or Application) Admin of all projects to be uassigned
//	@Produce json
//	@Param instrument_id path string true "instrument uuid" Format(uuid)
//	@Param project_ids body  true "project uuids" Format(uuid)
//	@Success 200 {object} model.InstrumentsValidation
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /instruments/{instrument_id}/assignments [post]
//	@Security Bearer
func (h *ApiHandler) UnssignProjectsFromInstrument(c echo.Context) error {
	iID, err := uuid.Parse(c.Param("instrument_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}
	p := c.Get("profile").(model.Profile)
	dryRun := strings.ToLower(c.QueryParam("dry_run")) == "true"

	pl := ProjectIDsPayload{ProjectIDs: make([]uuid.UUID, 0)}
	if err := c.Bind(&pl); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	v, err := h.InstrumentAssignService.UnassignProjectsFromInstrument(c.Request().Context(), p.ID, iID, pl.ProjectIDs, dryRun)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, v)
}

type InstrumentIDsPayload struct {
	InstrumentIDs []uuid.UUID `json:"instrument_ids"`
}

// AssignInstrumentsToProject godoc
//
//	@Summary assigns multiple instruments to a project
//	@Tags instrument
//	@Description must be Project (or Application) Admin of all existing instrument projects and project to be assigned
//	@Produce json
//	@Param project_id path string true "project uuid" Format(uuid)
//	@Param instrument_ids body true "instrument uuids" Format(uuid)
//	@Success 200 {object} map[string]interface{}
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/instruments/assignments [post]
//	@Security Bearer
func (h *ApiHandler) AssignInstrumentsToProject(c echo.Context) error {
	pID, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}
	p := c.Get("profile").(model.Profile)
	dryRun := strings.ToLower(c.QueryParam("dry_run")) == "true"

	pl := InstrumentIDsPayload{InstrumentIDs: make([]uuid.UUID, 0)}
	if err := c.Bind(&pl); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	v, err := h.InstrumentAssignService.AssignInstrumentsToProject(c.Request().Context(), p.ID, pID, pl.InstrumentIDs, dryRun)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, v)
}

// UnassignInstrumentsFromProject godoc
//
//	@Summary unassigns multiple instruments from a project
//	@Tags instrument
//	@Description must be Project (or Application) Admin of all projects to be unassigned
//	@Produce json
//	@Param project_id path string true "project uuid" Format(uuid)
//	@Param instrument_ids body true "instrument uuids" Format(uuid)
//	@Success 200 {object} map[string]interface{}
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/instruments/assignments [post]
//	@Security Bearer
func (h *ApiHandler) UnassignInstrumentsToProject(c echo.Context) error {
	pID, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}
	p := c.Get("profile").(model.Profile)
	dryRun := strings.ToLower(c.QueryParam("dry_run")) == "true"

	pl := InstrumentIDsPayload{InstrumentIDs: make([]uuid.UUID, 0)}
	if err := c.Bind(&pl); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	v, err := h.InstrumentAssignService.UnassignInstrumentsFromProject(c.Request().Context(), p.ID, pID, pl.InstrumentIDs, dryRun)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, v)
}
