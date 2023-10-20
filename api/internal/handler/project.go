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

// ListDistricts godoc
//
//	@Summary lists all districts
//	@Tags project
//	@Produce json
//	@Success 200 {array} model.District
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /districts [get]
func (h *ApiHandler) ListDistricts(c echo.Context) error {
	dd, err := h.ProjectService.ListDistricts(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, dd)
}

// ListProjects godoc
//
//	@Summary lists all projects optionally filtered by federal id
//	@Tags project
//	@Produce json
//	@Param federal_id query string false "federal id"
//	@Success 200 {array} model.Project
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects [get]
func (h *ApiHandler) ListProjects(c echo.Context) error {
	id := c.QueryParam("federal_id")
	if id != "" {
		projects, err := h.ProjectService.ListProjectsByFederalID(c.Request().Context(), id)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, message.InternalServerError)
		}
		return c.JSON(http.StatusOK, projects)
	}

	projects, err := h.ProjectService.ListProjects(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, projects)
}

// ListMyProjects godoc
//
//	@Summary lists projects for current profile
//	@Tags project
//	@Produce json
//	@Success 200 {array} model.Project
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /my_projects [get]
//	@Security CacOnly
func (h *ApiHandler) ListMyProjects(c echo.Context) error {
	p := c.Get("profile").(model.Profile)
	profileID := p.ID
	projects, err := h.ProjectService.ListProjectsForProfile(c.Request().Context(), profileID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, projects)
}

// ListProjectInstruments godoc
//
//	@Summary lists instruments associated with a project
//	@Tags project
//	@Produce json
//	@Param project_id path string true "project uuid" Format(uuid)
//	@Success 200 {array} model.Project
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/instruments [get]
func (h *ApiHandler) ListProjectInstruments(c echo.Context) error {
	id, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}
	nn, err := h.ProjectService.ListProjectInstruments(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, nn)
}

// ListProjectInstrumentNames godoc
//
//	@Summary lists names of all instruments associated with a project
//	@Tags project
//	@Produce json
//	@Param project_id path string true "project uuid" Format(uuid)
//	@Success 200 {array} string
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/instruments/names [get]
func (h *ApiHandler) ListProjectInstrumentNames(c echo.Context) error {
	id, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}
	names, err := h.ProjectService.ListProjectInstrumentNames(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, names)
}

// ListProjectInstrumentGroups godoc
//
//	@Summary lists instrument groups associated with a project
//	@Tags project
//	@Produce json
//	@Param project_id path string true "project uuid" Format(uuid)
//	@Success 200 {array} model.InstrumentGroup
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/instrument_groups [get]
func (h *ApiHandler) ListProjectInstrumentGroups(c echo.Context) error {
	id, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}
	gg, err := h.ProjectService.ListProjectInstrumentGroups(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, gg)
}

// GetProjectCount godoc
//
//	@Summary gets the total number of non-deleted projects in the system
//	@Tags project
//	@Produce json
//	@Success 200 {object} model.ProjectCount
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/count [get]
func (h *ApiHandler) GetProjectCount(c echo.Context) error {
	pc, err := h.ProjectService.GetProjectCount(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, pc)
}

// GetProject godoc
//
//	@Summary gets a single project
//	@Tags project
//	@Produce json
//	@Param project_id path string true "project uuid" Format(uuid)
//	@Success 200 {object} model.Project
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id} [get]
func (h *ApiHandler) GetProject(c echo.Context) error {
	id, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}
	project, err := h.ProjectService.GetProject(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, project)
}

// CreateProjectBulk godoc
//
//	@Summary accepts an array of instruments for bulk upload to the database
//	@Tags project
//	@Produce json
//	@Param project_collection body model.ProjectCollection true "project collection payload"
//	@Success 200 {array} model.IDAndSlug
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects [post]
//	@Security Bearer[admin]
func (h *ApiHandler) CreateProjectBulk(c echo.Context) error {
	pc := model.ProjectCollection{}
	if err := c.Bind(&pc); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// slugs already taken in the database
	slugsTaken, err := h.ProjectService.ListProjectSlugs(c.Request().Context())
	if err != nil {
		return err
	}

	// profile of user creating projects
	p := c.Get("profile").(model.Profile)

	// timestamp
	t := time.Now()

	for idx := range pc.Projects {
		// creator
		pc.Projects[idx].Creator = p.ID
		// create date
		pc.Projects[idx].CreateDate = t
		// Assign Slug
		s, err := util.NextUniqueSlug(pc.Projects[idx].Name, slugsTaken)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		pc.Projects[idx].Slug = s
		// Add slug to array of slugs originally fetched from the database
		// to catch duplicate names/slugs from the same bulk upload
		slugsTaken = append(slugsTaken, s)
	}

	pp, err := h.ProjectService.CreateProjectBulk(c.Request().Context(), pc.Projects)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, pp)
}

// UpdateProject godoc
//
//	@Summary updates an existing project
//	@Tags project
//	@Produce json
//	@Param project_id path string true "project uuid" Format(uuid)
//	@Param project body model.Project true "project payload"
//	@Success 200 {object} model.Project
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id} [put]
//	@Security Bearer
func (h *ApiHandler) UpdateProject(c echo.Context) error {
	id, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}
	p := &model.Project{ID: id}
	if err := c.Bind(p); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if id != p.ID {
		return echo.NewHTTPError(http.StatusBadRequest, message.MatchRouteParam("`id`"))
	}

	profile := c.Get("profile").(model.Profile)

	t := time.Now()
	p.Updater, p.UpdateDate = &profile.ID, &t

	pUpdated, err := h.ProjectService.UpdateProject(c.Request().Context(), *p)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, pUpdated)
}

// DeleteFlagProject godoc
//
//	@Summary soft deletes a project
//	@Tags project
//	@Produce json
//	@Param project_id path string true "project id" Format(uuid)
//	@Success 200 {object} map[string]interface{}
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id} [delete]
//	@Security Bearer
func (h *ApiHandler) DeleteFlagProject(c echo.Context) error {
	id, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	if err := h.ProjectService.DeleteFlagProject(c.Request().Context(), id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, make(map[string]interface{}))
}

// CreateProjectTimeseries godoc
//
//	@Summary exposes a timeseries at the project level
//	@Tags project
//	@Produce json
//	@Param project_id path string true "project id" Format(uuid)
//	@Param instrument_id path string true "timeseries uuid" Format(uuid)
//	@Success 200 {object} map[string]interface{}
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/timeseries/{timeseries_id} [post]
//	@Security Bearer
func (h *ApiHandler) CreateProjectTimeseries(c echo.Context) error {
	projectID, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	timeseriesID, err := uuid.Parse(c.Param("timeseries_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err = h.ProjectService.CreateProjectTimeseries(c.Request().Context(), projectID, timeseriesID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, make(map[string]interface{}))
}

// DeleteProjectTimeseries godoc
//
//	@Summary removes a timeseries from the project level
//	@Tags project
//	@Produce json
//	@Param project_id path string true "project id" Format(uuid)
//	@Param instrument_id path string true "timeseries uuid" Format(uuid)
//	@Success 200 {object} map[string]interface{}
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/timeseries/{timeseries_id} [delete]
//	@Security Bearer
func (h *ApiHandler) DeleteProjectTimeseries(c echo.Context) error {
	projectID, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	timeseriesID, err := uuid.Parse(c.Param("timeseries_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err = h.ProjectService.DeleteProjectTimeseries(c.Request().Context(), projectID, timeseriesID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, make(map[string]interface{}))
}
