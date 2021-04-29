package handlers

import (
	"net/http"

	"github.com/USACE/instrumentation-api/models"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

// ListProjectMembers returns project members and their role information
func ListProjectMembers(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := uuid.Parse(c.Param("project_id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Malformed ID")
		}
		mm, err := models.ListProjectMembers(db, &id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, mm)
	}
}

func AddProjectMemberRole(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		projectID, err := uuid.Parse(c.Param("project_id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Malformed ID")
		}
		profileID, err := uuid.Parse(c.Param("profile_id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Malformed ID")
		}
		roleID, err := uuid.Parse(c.Param("role_id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Malformed ID")
		}

		// profile granting role to profile_id
		grantedBy := c.Get("profile").(*models.Profile)

		r, err := models.AddProjectMemberRole(db, &projectID, &profileID, &roleID, &grantedBy.ID)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, models.DefaultMessageInternalServerError)
		}
		return c.JSON(http.StatusOK, r)
	}
}

func RemoveProjectMemberRole(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		projectID, err := uuid.Parse(c.Param("project_id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Malformed ID")
		}
		profileID, err := uuid.Parse(c.Param("profile_id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Malformed ID")
		}
		roleID, err := uuid.Parse(c.Param("role_id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Malformed ID")
		}

		if err := models.RemoveProjectMemberRole(db, &projectID, &profileID, &roleID); err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, make(map[string]interface{}))
	}
}
