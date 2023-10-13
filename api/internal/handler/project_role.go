package handler

import (
	"net/http"

	"github.com/USACE/instrumentation-api/api/internal/message"
	"github.com/USACE/instrumentation-api/api/internal/model"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// ListProjectMembers godoc
//
//	@Summary lists project members and their role information
//	@Tags project-role
//	@Produce json
//	@Param project_id path string true "project uuid" Format(uuid)
//	@Success 200 {array} model.ProjectMembership
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/members [get]
func (h *ApiHandler) ListProjectMembers(c echo.Context) error {
	id, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}
	mm, err := h.ProjectRoleService.ListProjectMembers(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, mm)
}

// AddProjectMemberRole godoc
//
//	@Summary adds project members and their role information
//	@Tags project-role
//	@Produce json
//	@Param project_id path string true "project uuid" Format(uuid)
//	@Param profile_id path string true "profile uuid" Format(uuid)
//	@Param role_id path string true "role uuid" Format(uuid)
//	@Success 200 {object} model.ProjectMembership
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/members/{profile_id}/roles/{role_id} [post]
func (h *ApiHandler) AddProjectMemberRole(c echo.Context) error {
	projectID, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}
	profileID, err := uuid.Parse(c.Param("profile_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}
	roleID, err := uuid.Parse(c.Param("role_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}

	// profile granting role to profile_id
	grantedBy := c.Get("profile").(model.Profile)

	r, err := h.ProjectRoleService.AddProjectMemberRole(c.Request().Context(), projectID, profileID, roleID, grantedBy.ID)
	if err != nil {
		// return echo.NewHTTPError(http.StatusInternalServerError, message.InternalServerError)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, r)
}

// RemoveProjectMemberRole godoc
//
//	@Summary removes project members and their role information
//	@Tags project-role
//	@Produce json
//	@Param project_id path string true "project uuid" Format(uuid)
//	@Param profile_id path string true "profile uuid" Format(uuid)
//	@Param role_id path string true "role uuid" Format(uuid)
//	@Success 200 {object} map[string]interface{}
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/members/{profile_id}/roles/{role_id} [delete]
func (h *ApiHandler) RemoveProjectMemberRole(c echo.Context) error {

	projectID, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}
	profileID, err := uuid.Parse(c.Param("profile_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}
	roleID, err := uuid.Parse(c.Param("role_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}

	if err := h.ProjectRoleService.RemoveProjectMemberRole(c.Request().Context(), projectID, profileID, roleID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, make(map[string]interface{}))
}
