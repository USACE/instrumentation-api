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

func (h ApiHandler) ListDistricts(c echo.Context) error {
	dd, err := model.ListDistricts(db)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, dd)
}

// ListProjects returns projects
func (h ApiHandler) ListProjects(c echo.Context) error {
	id := c.QueryParam("federal_id")
	if id != "" {
		projects, err := model.ListProjectsByFederalID(db, id)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, messages.InternalServerError)
		}
		return c.JSON(http.StatusOK, projects)
	}

	projects, err := model.ListProjects(db)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, projects)
}

func (h ApiHandler) ListMyProjects(c echo.Context) error {
	p := c.Get("profile").(*model.Profile)
	profileID := p.ID
	projects, err := model.ListMyProjects(db, &profileID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, projects)
}

// ListProjectInstruments returns instruments associated with a project
func (h ApiHandler) ListProjectInstruments(c echo.Context) error {
	id, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, messages.MalformedID)
	}
	nn, err := model.ListProjectInstruments(db, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, nn)
}

// ListProjectInstrumentNames returns names of all instruments associated with a project
func (h ApiHandler) ListProjectInstrumentNames(c echo.Context) error {
	id, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, messages.MalformedID)
	}
	names, err := model.ListProjectInstrumentNames(db, &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, names)
}

// ListProjectInstrumentGroups returns instrument groups associated with a project
func (h ApiHandler) ListProjectInstrumentGroups(c echo.Context) error {
	id, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, messages.MalformedID)
	}
	gg, err := model.ListProjectInstrumentGroups(db, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, gg)
}

// GetProjectCount returns the total number of non deleted projects in the system
func (h ApiHandler) GetProjectCount(c echo.Context) error {
	count, err := model.GetProjectCount(db)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"project_count": count})
}

// GetProject returns single project
func (h ApiHandler) GetProject(c echo.Context) error {
	id, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, messages.MalformedID)
	}
	project, err := model.GetProject(db, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, project)
}

// CreateProjectBulk accepts an array of instruments for bulk upload to the database
func (h ApiHandler) CreateProjectBulk(c echo.Context) error {

	pc := model.ProjectCollection{}
	if err := c.Bind(&pc); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// slugs already taken in the database
	slugsTaken, err := model.ListProjectSlugs(db)
	if err != nil {
		return err
	}

	// profile of user creating projects
	p := c.Get("profile").(*model.Profile)

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

	pp, err := model.CreateProjectBulk(db, pc.Projects)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// Send Project
	return c.JSON(http.StatusCreated, pp)
}

// UpdateProject updates an existing project
func (h ApiHandler) UpdateProject(c echo.Context) error {
	// id from url params
	id, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, messages.MalformedID)
	}
	// id from request
	p := &model.Project{ID: id}
	if err := c.Bind(p); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if id != p.ID {
		return echo.NewHTTPError(http.StatusBadRequest, messages.MatchRouteParam("`id`"))
	}

	profile := c.Get("profile").(*model.Profile)

	// timestamp
	t := time.Now()
	p.Updater, p.UpdateDate = &profile.ID, &t

	// update
	pUpdated, err := model.UpdateProject(db, p)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// return whole instrument
	return c.JSON(http.StatusOK, pUpdated)
}

// DeleteFlagProject sets the instrument group deleted flag true
func (h ApiHandler) DeleteFlagProject(c echo.Context) error {
	id, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	if err := model.DeleteFlagProject(db, id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, make(map[string]interface{}))
}

// CreateProjectTimeseries exposes a timeseries at the project level
func (h ApiHandler) CreateProjectTimeseries(c echo.Context) error {
	projectID, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	timeseriesID, err := uuid.Parse(c.Param("timeseries_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err = model.CreateProjectTimeseries(db, &projectID, &timeseriesID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, make(map[string]interface{}))
}

// DeleteProjectTimeseries removes a timeseries from the project level
func (h ApiHandler) DeleteProjectTimeseries(c echo.Context) error {
	projectID, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	timeseriesID, err := uuid.Parse(c.Param("timeseries_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err = model.DeleteProjectTimeseries(db, &projectID, &timeseriesID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, make(map[string]interface{}))
}
