package handler

import (
	"net/http"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/httperr"
	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// ListSurvey123sForProject godoc
//
//	@Summary lists Survey123 connections for a project
//	@Tags survey123
//	@Produce json
//	@Param project_id path string true "project uuid" Format(uuid)
//	@Success 200 {array} model.Survey123
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/survey123 [get]
func (h *ApiHandler) ListSurvey123sForProject(c echo.Context) error {
	projectID, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return httperr.MalformedID(err)
	}

	svv, err := h.Survey123Service.ListSurvey123sForProject(c.Request().Context(), projectID)
	if err != nil {
		return httperr.InternalServerError(err)
	}

	return c.JSON(http.StatusOK, svv)
}

// GetSurvey123Preview godoc
//
//	@Summary gets the most recent Survey123 raw json payload sent from the webhook API
//	@Tags survey123
//	@Produce json
//	@Param project_id path string true "project uuid" Format(uuid)
//	@Param survey123_id path string true "survey123 uuid" Format(uuid)
//	@Success 200 {object} string
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/survey123/{survey123_id}/previews [get]
func (h *ApiHandler) GetSurvey123Preview(c echo.Context) error {
	_, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return httperr.MalformedID(err)
	}
	survey123ID, err := uuid.Parse(c.Param("survey123_id"))
	if err != nil {
		return httperr.MalformedID(err)
	}

	pv, err := h.Survey123Service.GetSurvey123Preview(c.Request().Context(), survey123ID)
	if err != nil {
		return httperr.InternalServerError(err)
	}

	return c.JSON(http.StatusOK, pv)
}

// CreateSurvey123 godoc
//
//	@Summary creates a Survey123 connection with equivalency table mappings
//	@Tags survey123
//	@Accept json
//	@Produce json
//	@Param project_id path string true "project uuid" Format(uuid)
//	@Param survey123 body model.Survey123 true "survey123 payload"
//	@Success 200 {object} map[string]uuid.UUID
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/survey123 [post]
//	@Security Bearer
func (h *ApiHandler) CreateSurvey123(c echo.Context) error {
	projectID, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return httperr.MalformedID(err)
	}
	p := c.Get("profile").(model.Profile)

	var sv model.Survey123
	if err := c.Bind(&sv); err != nil {
		return httperr.MalformedBody(err)
	}
	sv.ProjectID = projectID
	sv.CreatorID = p.ID

	newID, err := h.Survey123Service.CreateSurvey123(c.Request().Context(), sv)
	if err != nil {
		return httperr.InternalServerError(err)
	}

	return c.JSON(http.StatusCreated, map[string]uuid.UUID{"id": newID})
}

// UpdateSurvey123 godoc
//
//	@Summary updates a Survey123 connection with equivalency table mappings
//	@Tags survey123
//	@Accept json
//	@Produce json
//	@Param project_id path string true "project uuid" Format(uuid)
//	@Param survey123_id path string true "survey123 uuid" Format(uuid)
//	@Param survey123 body model.Survey123 true "survey123 payload"
//	@Success 200 {object} map[string]uuid.UUID
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/survey123/{survey123_id} [put]
//	@Security Bearer
func (h *ApiHandler) UpdateSurvey123(c echo.Context) error {
	projectID, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return httperr.MalformedID(err)
	}
	survey123ID, err := uuid.Parse(c.Param("survey123_id"))
	if err != nil {
		return httperr.MalformedID(err)
	}
	var sv model.Survey123
	if err := c.Bind(&sv); err != nil {
		return httperr.MalformedBody(err)
	}
	sv.ProjectID = projectID
	sv.ID = survey123ID

	p := c.Get("profile").(model.Profile)
	t := time.Now()
	sv.UpdaterID, sv.UpdateDate = &p.ID, &t

	if err := h.Survey123Service.UpdateSurvey123(c.Request().Context(), sv); err != nil {
		return httperr.InternalServerError(err)
	}

	return c.JSON(http.StatusOK, map[string]uuid.UUID{"id": survey123ID})
}

// DeleteSurvey123 godoc
//
//	@Summary deletes a Survey123 connection with equivalency table mappings
//	@Tags survey123
//	@Produce json
//	@Param project_id path string true "project uuid" Format(uuid)
//	@Param survey123_id path string true "survey123 uuid" Format(uuid)
//	@Success 200 {object} map[string]uuid.UUID
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/survey123/{survey123_id} [delete]
//	@Security Bearer
func (h *ApiHandler) DeleteSurvey123(c echo.Context) error {
	_, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return httperr.MalformedID(err)
	}
	survey123ID, err := uuid.Parse(c.Param("survey123_id"))
	if err != nil {
		return httperr.MalformedID(err)
	}

	if err := h.Survey123Service.SoftDeleteSurvey123(c.Request().Context(), survey123ID); err != nil {
		return httperr.InternalServerError(err)
	}

	return c.JSON(http.StatusOK, map[string]uuid.UUID{"id": survey123ID})
}
