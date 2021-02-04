package handlers

import (
	"net/http"

	"github.com/USACE/instrumentation-api/models"
	"github.com/google/uuid"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

// ListAlertsForInstrument lists alerts for a single instrument
func ListAlertsForInstrument(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		instrumentID, err := uuid.Parse(c.Param("instrument_id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		aa, err := models.ListAlertsForInstrument(db, &instrumentID)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, aa)
	}
}

// ListMyAlerts returns all alerts a profile is subscribed to
func ListMyAlerts(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		p := c.Get("profile").(*models.Profile)
		profileID := p.ID
		aa, err := models.ListMyAlerts(db, &profileID)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, &aa)
	}
}

// DoAlertRead marks an alert as read for a profile
func DoAlertRead(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		p := c.Get("profile").(*models.Profile)
		profileID := p.ID
		alertID, err := uuid.Parse(c.Param("alert_id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		a, err := models.DoAlertRead(db, &profileID, &alertID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, a)
	}
}

// DoAlertUnread marks an alert as unread for a profile
func DoAlertUnread(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		p := c.Get("profile").(*models.Profile)
		profileID := p.ID
		alertID, err := uuid.Parse(c.Param("alert_id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		a, err := models.DoAlertUnread(db, &profileID, &alertID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, a)
	}
}
