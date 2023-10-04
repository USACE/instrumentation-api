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

func (h *ApiHandler) ListDistricts(c echo.Context) error {
	dd, err := h.ProjectService.ListDistricts(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, dd)
}

// ListProjects returns projects
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

func (h *ApiHandler) ListMyProjects(c echo.Context) error {
	p := c.Get("profile").(model.Profile)
	profileID := p.ID
	projects, err := h.ProjectService.ListProjectsForProfile(c.Request().Context(), profileID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, projects)
}

// ListProjectInstruments returns instruments associated with a project
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

// ListProjectInstrumentNames returns names of all instruments associated with a project
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

// ListProjectInstrumentGroups returns instrument groups associated with a project
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

// GetProjectCount returns the total number of non deleted projects in the system
func (h *ApiHandler) GetProjectCount(c echo.Context) error {
	count, err := h.ProjectService.GetProjectCount(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"project_count": count})
}

// GetProject returns single project
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

// CreateProjectBulk accepts an array of instruments for bulk upload to the database
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
	// Send Project
	return c.JSON(http.StatusCreated, pp)
}

// UpdateProject updates an existing project
func (h *ApiHandler) UpdateProject(c echo.Context) error {
	// id from url params
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
	// return whole instrument
	return c.JSON(http.StatusOK, pUpdated)
}

// DeleteFlagProject sets the instrument group deleted flag true
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

// CreateProjectTimeseries exposes a timeseries at the project level
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

// DeleteProjectTimeseries removes a timeseries from the project level
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
