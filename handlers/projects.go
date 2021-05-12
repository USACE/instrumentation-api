package handlers

import (
	"net/http"
	"time"

	"github.com/USACE/instrumentation-api/dbutils"
	"github.com/USACE/instrumentation-api/models"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

// ListProjects returns projects
func ListProjects(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		projects, err := models.ListProjects(db)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, projects)
	}
}

func ListMyProjects(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		p := c.Get("profile").(*models.Profile)
		profileID := p.ID
		projects, err := models.ListMyProjects(db, &profileID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, projects)
	}
}

// ListProjectInstruments returns instruments associated with a project
func ListProjectInstruments(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := uuid.Parse(c.Param("project_id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Malformed ID")
		}
		nn, err := models.ListProjectInstruments(db, id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, nn)
	}
}

// ListProjectInstrumentNames returns names of all instruments associated with a project
func ListProjectInstrumentNames(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := uuid.Parse(c.Param("project_id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Malformed ID")
		}
		names, err := models.ListProjectInstrumentNames(db, &id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, names)
	}
}

// ListProjectInstrumentGroups returns instrument groups associated with a project
func ListProjectInstrumentGroups(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := uuid.Parse(c.Param("project_id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Malformed ID")
		}
		gg, err := models.ListProjectInstrumentGroups(db, id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, gg)
	}
}

// GetProjectCount returns the total number of non deleted projects in the system
func GetProjectCount(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		count, err := models.GetProjectCount(db)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, map[string]interface{}{"project_count": count})
	}
}

// GetProject returns single project
func GetProject(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := uuid.Parse(c.Param("project_id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Malformed ID")
		}
		project, err := models.GetProject(db, id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, project)
	}
}

// CreateProjectBulk accepts an array of instruments for bulk upload to the database
func CreateProjectBulk(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		pc := models.ProjectCollection{}
		if err := c.Bind(&pc); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		// slugs already taken in the database
		slugsTaken, err := models.ListProjectSlugs(db)
		if err != nil {
			return err
		}

		// profile of user creating projects
		p := c.Get("profile").(*models.Profile)

		// timestamp
		t := time.Now()

		for idx := range pc.Projects {
			// creator
			pc.Projects[idx].Creator = p.ID
			// create date
			pc.Projects[idx].CreateDate = t
			// Assign Slug
			s, err := dbutils.NextUniqueSlug(pc.Projects[idx].Name, slugsTaken)
			if err != nil {
				return c.JSON(http.StatusBadRequest, err)
			}
			pc.Projects[idx].Slug = s
			// Add slug to array of slugs originally fetched from the database
			// to catch duplicate names/slugs from the same bulk upload
			slugsTaken = append(slugsTaken, s)
		}

		pp, err := models.CreateProjectBulk(db, pc.Projects)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		// Send Project
		return c.JSON(http.StatusCreated, pp)
	}
}

// UpdateProject updates an existing project
func UpdateProject(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// id from url params
		id, err := uuid.Parse(c.Param("project_id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Malformed ID")
		}
		// id from request
		p := &models.Project{ID: id}
		if err := c.Bind(p); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		// check :id in url params matches id in request body
		if id != p.ID {
			return c.JSON(
				http.StatusBadRequest,
				map[string]interface{}{
					"err": "url parameter id does not match object id in body",
				},
			)
		}

		profile := c.Get("profile").(*models.Profile)

		// timestamp
		t := time.Now()
		p.Updater, p.UpdateDate = &profile.ID, &t

		// update
		pUpdated, err := models.UpdateProject(db, p)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		// return whole instrument
		return c.JSON(http.StatusOK, pUpdated)
	}
}

// DeleteFlagProject sets the instrument group deleted flag true
func DeleteFlagProject(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := uuid.Parse(c.Param("project_id"))
		if err != nil {
			return c.JSON(http.StatusNotFound, err)
		}
		if err := models.DeleteFlagProject(db, id); err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, make(map[string]interface{}))
	}
}

// CreateProjectTimeseries exposes a timeseries at the project level
func CreateProjectTimeseries(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		projectID, err := uuid.Parse(c.Param("project_id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		timeseriesID, err := uuid.Parse(c.Param("timeseries_id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		err = models.CreateProjectTimeseries(db, &projectID, &timeseriesID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusCreated, make(map[string]interface{}))
	}
}

// DeleteProjectTimeseries removes a timeseries from the project level
func DeleteProjectTimeseries(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		projectID, err := uuid.Parse(c.Param("project_id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		timeseriesID, err := uuid.Parse(c.Param("timeseries_id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		err = models.DeleteProjectTimeseries(db, &projectID, &timeseriesID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, make(map[string]interface{}))
	}
}
