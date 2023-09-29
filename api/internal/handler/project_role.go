package handler

import (
	"net/http"

	"github.com/USACE/instrumentation-api/api/internal/messages"
	"github.com/USACE/instrumentation-api/api/internal/model"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// ListProjectMembers returns project members and their role information
func (h ApiHandler) ListProjectMembers(c echo.Context) error {
	id, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, messages.MalformedID)
	}
	mm, err := model.ListProjectMembers(db, &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, mm)
}

func (h ApiHandler) AddProjectMemberRole(c echo.Context) error {

	projectID, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, messages.MalformedID)
	}
	profileID, err := uuid.Parse(c.Param("profile_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, messages.MalformedID)
	}
	roleID, err := uuid.Parse(c.Param("role_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, messages.MalformedID)
	}

	// profile granting role to profile_id
	grantedBy := c.Get("profile").(*model.Profile)

	r, err := model.AddProjectMemberRole(db, &projectID, &profileID, &roleID, &grantedBy.ID)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, messages.InternalServerError)
	}
	return c.JSON(http.StatusOK, r)
}

func (h ApiHandler) RemoveProjectMemberRole(c echo.Context) error {

	projectID, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, messages.MalformedID)
	}
	profileID, err := uuid.Parse(c.Param("profile_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, messages.MalformedID)
	}
	roleID, err := uuid.Parse(c.Param("role_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, messages.MalformedID)
	}

	if err := model.RemoveProjectMemberRole(db, &projectID, &profileID, &roleID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, make(map[string]interface{}))
}
