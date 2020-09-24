package handlers

import (
	"net/http"

	"github.com/USACE/instrumentation-api/models"
	"github.com/lib/pq"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
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
		// Send Alerts
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

// SubscribeProfileToInstrumentAlert subscribes a profile to an alert
func SubscribeProfileToInstrumentAlert(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		p, err := myProfileFromContext(c, db)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		profileID := p.ID

		alertID, err := uuid.Parse(c.Param("alert_id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		pa, err := models.SubscribeProfileToInstrumentAlert(db, &alertID, &profileID)
		if err != nil {
			if err, ok := err.(*pq.Error); ok {
				switch err.Code {
				// Profile already subscribed to instrument alert; Get ProfileAlert and return it
				// Return a RESTful 200; i.e. nothing wrong, state of data is already "subscribed"
				case "23505":
					pa, err := models.GetAlertSubscription(db, &alertID, &profileID)
					if err != nil {
						return c.JSON(http.StatusInternalServerError, err)
					}
					return c.JSON(http.StatusOK, &pa)
				default:
					return c.JSON(http.StatusInternalServerError, err)
				}
			}
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, &pa)
	}
}

// UnsubscribeProfileToInstrumentAlert unsubscribes a profile to an alert
func UnsubscribeProfileToInstrumentAlert(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		p, err := myProfileFromContext(c, db)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		profileID := p.ID

		alertID, err := uuid.Parse(c.Param("alert_id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		err = models.UnsubscribeProfileToInstrumentAlert(db, &alertID, &profileID)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.NoContent(http.StatusOK)
	}
}

// ListMyAlertSubscriptions returns all alerts you are subscribed to and settings
func ListMyAlertSubscriptions(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		p, err := myProfileFromContext(c, db)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		profileID := p.ID
		ss, err := models.ListMyAlertSubscriptions(db, &profileID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, &ss)
	}
}

// UpdateMyAlertSubscription updates settings for an alert subscription
func UpdateMyAlertSubscription(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var s models.AlertSubscription
		if err := c.Bind(&s); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		// AlertID From Route Params
		sID, err := uuid.Parse(c.Param("alert_subscription_id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		if s.ID != sID {
			return c.String(http.StatusBadRequest, "route parameter subscription_id does not match id in JSON payload")
		}
		// Get Profile
		p, err := myProfileFromContext(c, db)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		// Verify Profile ID matches ProfileID of Subscription to be Modified
		// No Modifying anyone else's settings
		t, err := models.GetAlertSubscriptionByID(db, &sID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		if p.ID != t.ProfileID {
			return c.NoContent(http.StatusUnauthorized)
		}
		sUpdated, err := models.UpdateMyAlertSubscription(db, &s)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, sUpdated)
	}
}
