package handler

import (
	"net/http"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/httperr"
	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// CreateAlertConfigThreshold godoc
//
//	@Summary creates one threshold alert config
//	@Tags alert-config
//	@Accept json
//	@Produce json
//	@Param project_id path string true "project uuid" Format(uuid)
//	@Param alert_config body model.AlertConfigThreshold true "threshold alert config payload"
//	@Param key query string false "api key"
//	@Success 200 {object} model.AlertConfig
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/alert_configs/thresholds [post]
//	@Security Bearer
func (h *ApiHandler) CreateAlertConfigThreshold(c echo.Context) error {
	var ac model.AlertConfigThreshold
	if err := c.Bind(&ac); err != nil {
		return httperr.MalformedBody(err)
	}
	projectID, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return httperr.MalformedID(err)
	}
	profile := c.Get("profile").(model.Profile)
	ac.ProjectID, ac.CreatorID, ac.CreateDate = projectID, profile.ID, time.Now()

	acNew, err := h.AlertConfigService.CreateAlertConfigThreshold(c.Request().Context(), ac)
	if err != nil {
		return httperr.InternalServerError(err)
	}
	return c.JSON(http.StatusCreated, acNew)
}

// UpdateAlertConfigThreshold godoc
//
//	@Summary updates an existing threshold alert config
//	@Tags alert-config
//	@Accept json
//	@Produce json
//	@Param project_id path string true "project uuid" Format(uuid)
//	@Param alert_config_id path string true "alert config uuid" Format(uuid)
//	@Param alert_config body model.AlertConfigThreshold true "threshold alert config payload"
//	@Param key query string false "api key"
//	@Success 200 {object} model.AlertConfig
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/alert_configs/thresholds/{alert_config_id} [put]
//	@Security Bearer
func (h *ApiHandler) UpdateAlertConfigThreshold(c echo.Context) error {
	var ac model.AlertConfigThreshold
	if err := c.Bind(&ac); err != nil {
		return httperr.MalformedBody(err)
	}

	projectID, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return httperr.MalformedID(err)
	}
	ac.ProjectID = projectID

	acID, err := uuid.Parse(c.Param("alert_config_id"))
	if err != nil {
		return httperr.MalformedID(err)
	}
	ac.ID = acID

	p := c.Get("profile").(model.Profile)
	t := time.Now()
	ac.UpdaterID, ac.UpdateDate = &p.ID, &t
	aUpdated, err := h.AlertConfigService.UpdateAlertConfigThreshold(c.Request().Context(), ac)
	if err != nil {
		return httperr.InternalServerError(err)
	}
	return c.JSON(http.StatusOK, aUpdated)
}
