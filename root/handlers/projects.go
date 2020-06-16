package handlers

import (
	"api/root/dbutils"
	"api/root/models"
	"net/http"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
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

		for idx := range pc.Projects {
			// Assign UUID
			pc.Projects[idx].ID = uuid.Must(uuid.NewRandom())
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

		if err := models.CreateProjectBulk(db, pc.Projects); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		// Send Project
		return c.JSON(http.StatusCreated, pc.Projects)
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
		return c.NoContent(http.StatusOK)
	}
}
