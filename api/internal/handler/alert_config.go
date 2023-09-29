package handler

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/messages"
	"github.com/USACE/instrumentation-api/api/internal/model"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// ListProjectAlertConfigs lists alert configs for a single project optionally filtered by alert_type_id
func (h ApiHandler) GetAllAlertConfigsForProject(c echo.Context) error {
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
		aa, err = h.AlertConfigStore.GetAllAlertConfigsForProjectAndAlertType(c.Request().Context(), projectID, alertTypeID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	} else {
		aa, err = h.AlertConfigStore.GetAllAlertConfigsForProject(c.Request().Context(), projectID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}
	if len(aa) == 0 {
		return echo.NewHTTPError(http.StatusNotFound, messages.NotFound)
	}
	return c.JSON(http.StatusOK, aa)
}

// ListInstrumentAlertConfigs lists alerts for a single instrument
func (h ApiHandler) ListInstrumentAlertConfigs(c echo.Context) error {
	instrumentID, err := uuid.Parse(c.Param("instrument_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	aa, err := h.AlertConfigStore.GetAllAlertConfigsForInstrument(c.Request().Context(), instrumentID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if len(aa) == 0 {
		return echo.NewHTTPError(http.StatusNotFound, messages.NotFound)
	}
	return c.JSON(http.StatusOK, aa)
}

// GetAlertConfig gets a single alert
func (h ApiHandler) GetAlertConfig(c echo.Context) error {
	acID, err := uuid.Parse(c.Param("alert_config_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	a, err := h.AlertConfigStore.GetOneAlertConfig(c.Request().Context(), acID)
	if err != nil {
		if err == sql.ErrNoRows {
			return echo.NewHTTPError(http.StatusNotFound, messages.NotFound)
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, a)
}

// CreateAlertConfig creates one alert config
func (h ApiHandler) CreateAlertConfig(c echo.Context) error {
	ac := model.AlertConfig{}
	if err := c.Bind(&ac); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	projectID, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	profile := c.Get("profile").(*model.Profile)
	ac.ProjectID, ac.Creator, ac.CreateDate = projectID, profile.ID, time.Now()

	aa, err := h.AlertConfigStore.CreateAlertConfig(c.Request().Context(), ac)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, aa)
}

// UpdateAlertConfig updates an existing alert
func (h ApiHandler) UpdateAlertConfig(c echo.Context) error {
	var ac model.AlertConfig
	if err := c.Bind(&ac); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	acID, err := uuid.Parse(c.Param("alert_config_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	p := c.Get("profile").(*model.Profile)
	t := time.Now()
	ac.Updater, ac.UpdateDate = &p.ID, &t
	aUpdated, err := h.AlertConfigStore.UpdateAlertConfig(c.Request().Context(), acID, ac)
	if err != nil {
		if err == sql.ErrNoRows {
			return echo.NewHTTPError(http.StatusNotFound, messages.NotFound)
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, &aUpdated)
}

// DeleteAlertConfig Deletes an Alert Config
func (h ApiHandler) DeleteAlertConfig(c echo.Context) error {
	acID, err := uuid.Parse(c.Param("alert_config_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	if err := h.AlertConfigStore.DeleteAlertConfig(c.Request().Context(), acID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, make(map[string]interface{}))
}
