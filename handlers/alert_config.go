package handlers

import (
	"net/http"

	"github.com/USACE/instrumentation-api/models"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

// ListInstrumentAlertConfigs lists alerts for a single instrument
func ListInstrumentAlertConfigs(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		instrumentID, err := uuid.Parse(c.Param("instrument_id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		aa, err := models.ListInstrumentAlertConfigs(db, &instrumentID)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, aa)
	}
}

// GetAlertConfig gets a single alert
func GetAlertConfig(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		alertConfigID, err := uuid.Parse(c.Param("alert_config_id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		a, err := models.GetAlertConfig(db, &alertConfigID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, a)
	}
}

// CreateInstrumentAlertConfigs creates one or more alerts
func CreateInstrumentAlertConfigs(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		ac := models.AlertConfigCollection{}
		if err := c.Bind(&ac); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		// Get instrument_id from Route Params
		instrumentID, err := uuid.Parse(c.Param("instrument_id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		// Get action information from context
		a, err := models.NewAction(c)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		aa, err := models.CreateInstrumentAlertConfigs(db, a, &instrumentID, ac.Items)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		// Send Alerts
		return c.JSON(http.StatusCreated, aa)
	}
}

// UpdateInstrumentAlertConfig updates an existing alert
func UpdateInstrumentAlertConfig(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var alert models.AlertConfig
		if err := c.Bind(&alert); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		// Get instrument_id from Route Params
		instrumentID, err := uuid.Parse(c.Param("instrument_id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		// Get alert_config_id from Route Params
		alertID, err := uuid.Parse(c.Param("alert_config_id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		// Get action information from context
		action, err := models.NewAction(c)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		aUpdated, err := models.UpdateInstrumentAlertConfig(db, action, &instrumentID, &alertID, &alert)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		// Send Alert
		return c.JSON(http.StatusOK, &aUpdated)
	}
}

// DeleteInstrumentAlertConfig Deletes an Alert
func DeleteInstrumentAlertConfig(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		alertID, err := uuid.Parse(c.Param("alert_config_id"))
		if err != nil {
			return c.JSON(http.StatusNotFound, err)
		}
		// Get instrument_id from Route Params
		instrumentID, err := uuid.Parse(c.Param("instrument_id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		if err := models.DeleteInstrumentAlertConfig(db, &alertID, &instrumentID); err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.NoContent(http.StatusOK)
	}
}
