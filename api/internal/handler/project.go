package handler

import (
	"database/sql"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/message"
	"github.com/USACE/instrumentation-api/api/internal/model"

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
	ctx := c.Request().Context()

	fedID := c.QueryParam("federal_id")
	if fedID != "" {
		projects, err := h.ProjectService.ListProjectsByFederalID(ctx, fedID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, message.InternalServerError)
		}
		return c.JSON(http.StatusOK, projects)
	}

	projects, err := h.ProjectService.ListProjects(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, projects)
}

// ListMyProjects godoc
//
//	@Summary lists projects where current profile is an admin or member with optional filter by project role
//	@Tags project
//	@Produce json
//	@Param role query string false "role"
//	@Success 200 {array} model.Project
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /my_projects [get]
//	@Security CacOnly
func (h *ApiHandler) ListMyProjects(c echo.Context) error {
	ctx := c.Request().Context()

	p := c.Get("profile").(model.Profile)

	if p.IsAdmin {
		projects, err := h.ProjectService.ListProjects(ctx)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, message.InternalServerError)
		}
		return c.JSON(http.StatusOK, projects)
	}

	role := c.QueryParam("role")
	if role != "" {
		role = strings.ToLower(role)
		if role == "admin" || role == "member" {
			projects, err := h.ProjectService.ListProjectsForProfileRole(ctx, p.ID, role)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			return c.JSON(http.StatusOK, projects)
		}
		return echo.NewHTTPError(http.StatusBadRequest, "role parameter must be 'admin' or 'member'")
	}

	projects, err := h.ProjectService.ListProjectsForProfile(ctx, p.ID)
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
//	@Success 200 {array} model.IDSlugName
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects [post]
//	@Security Bearer
func (h *ApiHandler) CreateProjectBulk(c echo.Context) error {
	var pc model.ProjectCollection
	if err := c.Bind(&pc); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	p := c.Get("profile").(model.Profile)
	t := time.Now()

	for idx := range pc {
		if pc[idx].Name == "" {
			return echo.NewHTTPError(http.StatusBadRequest, errors.New("project name required"))
		}
		pc[idx].CreatorID = p.ID
		pc[idx].CreateDate = t
	}

	pp, err := h.ProjectService.CreateProjectBulk(c.Request().Context(), pc)
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
	var p model.Project
	if err := c.Bind(&p); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	p.ID = id
	profile := c.Get("profile").(model.Profile)

	t := time.Now()
	p.UpdaterID, p.UpdateDate = &profile.ID, &t

	pUpdated, err := h.ProjectService.UpdateProject(c.Request().Context(), p)
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

// UploadProjectImage godoc
//
//	@Summary uploades a picture for a project
//	@Tags project
//	@Accept jpeg
//	@Accept png
//	@Produce json
//	@Param project_id path string true "project id" Format(uuid)
//	@Success 200 {object} map[string]interface{}
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/images [post]
//	@Security Bearer
func (h *ApiHandler) UploadProjectImage(c echo.Context) error {
	projectID, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}
	fh, err := c.FormFile("image")
	if err != nil || fh == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "attached form file 'image' required")
	}
	if fh.Size > 2000000 {
		return echo.NewHTTPError(http.StatusBadRequest, "image exceeds max size of 2MB")
	}

	if err := h.ProjectService.UploadProjectImage(c.Request().Context(), projectID, *fh, h.BlobService.UploadContext); err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return echo.NewHTTPError(http.StatusBadRequest, message.BadRequest)
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, make(map[string]interface{}))
}
