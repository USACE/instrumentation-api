package handlers

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/USACE/instrumentation-api/api/messages"
	"github.com/USACE/instrumentation-api/api/models"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

// ListProjectAlertConfigs lists alert configs for a single project optionally filtered by alert_type_id
func ListProjectAlertConfigs(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		projectID, err := uuid.Parse(c.Param("project_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		var aa []models.AlertConfig
		if qp := c.QueryParam("alert_type_id"); qp != "" {
			alertTypeID, err := uuid.Parse(qp)
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
			aa, err = models.ListProjectAlertConfigsByAlertType(db, &projectID, &alertTypeID)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
		} else {
			aa, err = models.ListProjectAlertConfigs(db, &projectID)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
		}
		return c.JSON(http.StatusOK, aa)
	}
}

// ListInstrumentAlertConfigs lists alerts for a single instrument
func ListInstrumentAlertConfigs(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		instrumentID, err := uuid.Parse(c.Param("instrument_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		aa, err := models.ListInstrumentAlertConfigs(db, &instrumentID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, aa)
	}
}

// GetAlertConfig gets a single alert
func GetAlertConfig(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		acID, err := uuid.Parse(c.Param("alert_config_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		a, err := models.GetAlertConfig(db, &acID)
		if err != nil {
			if err == sql.ErrNoRows {
				return echo.NewHTTPError(http.StatusNotFound, messages.NotFound)
			}
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, a)
	}
}

// CreateAlertConfig creates one alert config
func CreateAlertConfig(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		ac := models.AlertConfig{}
		if err := c.Bind(&ac); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		projectID, err := uuid.Parse(c.Param("project_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		profile := c.Get("profile").(*models.Profile)
		ac.ProjectID, ac.Creator, ac.CreateDate = projectID, profile.ID, time.Now()

		aa, err := models.CreateAlertConfig(db, &ac)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusCreated, aa)
	}
}

// UpdateAlertConfig updates an existing alert
func UpdateAlertConfig(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var ac models.AlertConfig
		if err := c.Bind(&ac); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		acID, err := uuid.Parse(c.Param("alert_config_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		p := c.Get("profile").(*models.Profile)
		t := time.Now()
		ac.Updater, ac.UpdateDate = &p.ID, &t
		aUpdated, err := models.UpdateAlertConfig(db, &acID, &ac)
		if err != nil {
			if err == sql.ErrNoRows {
				return echo.NewHTTPError(http.StatusNotFound, messages.NotFound)
			}
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, &aUpdated)
	}
}

// DeleteAlertConfig Deletes an Alert Config
func DeleteAlertConfig(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		acID, err := uuid.Parse(c.Param("alert_config_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		if err := models.DeleteAlertConfig(db, &acID); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, make(map[string]interface{}))
	}
}
