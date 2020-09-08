package handlers

import (
	"api/root/models"
	"net/http"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
)

// ListInstrumentAlerts lists alerts for a single instrument
func ListInstrumentAlerts(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		instrumentID, err := uuid.Parse(c.Param("instrument_id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		aa, err := models.ListInstrumentAlerts(db, &instrumentID)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, aa)
	}
}

// GetAlert gets a single alert
func GetAlert(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		alertID, err := uuid.Parse(c.Param("alert_id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		a, err := models.GetAlert(db, &alertID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, a)
	}
}

// CreateInstrumentAlert creates one or more alerts
func CreateInstrumentAlert(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		ac := models.AlertCollection{}
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
		aa, err := models.CreateInstrumentAlerts(db, a, &instrumentID, ac.Items)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		// Send Alertsg
		return c.JSON(http.StatusCreated, aa)
	}
}

// UpdateInstrumentAlert updates an existing alert
func UpdateInstrumentAlert(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var alert models.Alert
		if err := c.Bind(&alert); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		// Get instrument_id from Route Params
		instrumentID, err := uuid.Parse(c.Param("instrument_id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		// Get alert_id from Route Params
		alertID, err := uuid.Parse(c.Param("alert_id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		// Get action information from context
		action, err := models.NewAction(c)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		aUpdated, err := models.UpdateInstrumentAlert(db, action, &instrumentID, &alertID, &alert)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		// Send Alert
		return c.JSON(http.StatusOK, &aUpdated)
	}
}

// DeleteInstrumentAlert Deletes an Alert
func DeleteInstrumentAlert(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		alertID, err := uuid.Parse(c.Param("alert_id"))
		if err != nil {
			return c.JSON(http.StatusNotFound, err)
		}
		// Get instrument_id from Route Params
		instrumentID, err := uuid.Parse(c.Param("instrument_id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		if err := models.DeleteInstrumentAlert(db, &alertID, &instrumentID); err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.NoContent(http.StatusOK)
	}
}
