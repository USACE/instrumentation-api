package handler

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/message"
	"github.com/USACE/instrumentation-api/api/internal/model"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// GetAllAlertConfigsForProject godoc
//
//	@Summary lists alert configs for a project
//	@Tags alert-config
//	@Produce json
//	@Param project_id path string true "project uuid" Format(uuid)
//	@Success 200 {array} model.AlertConfig
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/alert_configs [get]
func (h *ApiHandler) GetAllAlertConfigsForProject(c echo.Context) error {
	projectID, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	var aa []model.AlertConfig
	if qp := c.QueryParam("alert_type_id"); qp != "" {
		alertTypeID, err := uuid.Parse(qp)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		aa, err = h.AlertConfigService.GetAllAlertConfigsForProjectAndAlertType(c.Request().Context(), projectID, alertTypeID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	} else {
		aa, err = h.AlertConfigService.GetAllAlertConfigsForProject(c.Request().Context(), projectID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}
	return c.JSON(http.StatusOK, aa)
}

// ListInstrumentAlertConfigs godoc
//
//	@Summary lists alerts for a single instrument
//	@Tags alert-config
//	@Produce json
//	@Param project_id path string true "project uuid" Format(uuid)
//	@Param instrument_id path string true "instrument uuid" Format(uuid)
//	@Success 200 {array} model.AlertConfig
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/instruments/{instrument_id}/alert_configs [get]
func (h *ApiHandler) ListInstrumentAlertConfigs(c echo.Context) error {
	instrumentID, err := uuid.Parse(c.Param("instrument_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	aa, err := h.AlertConfigService.GetAllAlertConfigsForInstrument(c.Request().Context(), instrumentID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, aa)
}

// GetAlertConfig godoc
//
//	@Summary gets a single alert
//	@Tags alert-config
//	@Produce json
//	@Param project_id path string true "project uuid" Format(uuid)
//	@Param alert_config_id path string true "alert config uuid" Format(uuid)
//	@Success 200 {object} model.AlertConfig
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/alert_configs/{alert_config_id} [get]
func (h *ApiHandler) GetAlertConfig(c echo.Context) error {
	acID, err := uuid.Parse(c.Param("alert_config_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	a, err := h.AlertConfigService.GetOneAlertConfig(c.Request().Context(), acID)
	if err != nil {
		if err == sql.ErrNoRows {
			return echo.NewHTTPError(http.StatusNotFound, message.NotFound)
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, a)
}

// CreateAlertConfig godoc
//
//	@Summary creates one alert config
//	@Tags alert-config
//	@Accept json
//	@Produce json
//	@Param project_id path string true "project uuid" Format(uuid)
//	@Param alert_config body model.AlertConfig true "alert config payload"
//	@Success 200 {object} model.AlertConfig
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/alert_configs [post]
//	@Security Bearer
func (h *ApiHandler) CreateAlertConfig(c echo.Context) error {
	ac := model.AlertConfig{}
	if err := c.Bind(&ac); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	projectID, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	profile := c.Get("profile").(model.Profile)
	ac.ProjectID, ac.Creator, ac.CreateDate = projectID, profile.ID, time.Now()

	acNew, err := h.AlertConfigService.CreateAlertConfig(c.Request().Context(), ac)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, acNew)
}

// UpdateAlertConfig godoc
//
//	@Summary updates an existing alert config
//	@Tags alert-config
//	@Accept json
//	@Produce json
//	@Param project_id path string true "project uuid" Format(uuid)
//	@Param alert_config_id path string true "alert config uuid" Format(uuid)
//	@Param alert_config body model.AlertConfig true "alert config payload"
//	@Success 200 {array} model.AlertConfig
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/alert_configs/{alert_config_id} [put]
//	@Security Bearer
func (h *ApiHandler) UpdateAlertConfig(c echo.Context) error {
	var ac model.AlertConfig
	if err := c.Bind(&ac); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	acID, err := uuid.Parse(c.Param("alert_config_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	p := c.Get("profile").(model.Profile)
	t := time.Now()
	ac.Updater, ac.UpdateDate = &p.ID, &t
	aUpdated, err := h.AlertConfigService.UpdateAlertConfig(c.Request().Context(), acID, ac)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return echo.NewHTTPError(http.StatusNotFound, message.NotFound)
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, aUpdated)
}

// DeleteAlertConfig godoc
//
//	@Summary deletes an alert config
//	@Tags alert-config
//	@Produce json
//	@Param project_id path string true "Project ID" Format(uuid)
//	@Param alert_config_id path string true "instrument uuid" Format(uuid)
//	@Success 200 {array} model.AlertConfig
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/alert_configs/{alert_config_id} [delete]
//	@Security Bearer
func (h *ApiHandler) DeleteAlertConfig(c echo.Context) error {
	acID, err := uuid.Parse(c.Param("alert_config_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	if err := h.AlertConfigService.DeleteAlertConfig(c.Request().Context(), acID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, make(map[string]interface{}))
}
