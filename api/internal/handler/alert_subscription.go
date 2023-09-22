package handler

import (
	"net/http"

	"github.com/USACE/instrumentation-api/api/internal/messages"
	"github.com/USACE/instrumentation-api/api/internal/models"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

// SubscribeProfileToAlerts subscribes a profile to an alert
func SubscribeProfileToAlerts(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		p := c.Get("profile").(*models.Profile)
		profileID := p.ID

		alertConfigID, err := uuid.Parse(c.Param("alert_config_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		pa, err := models.SubscribeProfileToAlerts(db, &alertConfigID, &profileID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, pa)
	}
}

// UnsubscribeProfileToAlerts unsubscribes a profile to an alert
func UnsubscribeProfileToAlerts(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		p := c.Get("profile").(*models.Profile)
		profileID := p.ID

		alertConfigID, err := uuid.Parse(c.Param("alert_config_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if err = models.UnsubscribeProfileToAlerts(db, &alertConfigID, &profileID); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, make(map[string]interface{}))
	}
}

// ListMyAlertSubscriptions returns all alerts you are subscribed to and settings
func ListMyAlertSubscriptions(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		p := c.Get("profile").(*models.Profile)
		profileID := p.ID
		ss, err := models.ListMyAlertSubscriptions(db, &profileID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, &ss)
	}
}

// UpdateMyAlertSubscription updates settings for an alert subscription
func UpdateMyAlertSubscription(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var s models.AlertSubscription
		if err := c.Bind(&s); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		// alertConfigID From Route Params
		sID, err := uuid.Parse(c.Param("alert_subscription_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if s.ID != sID {
			return echo.NewHTTPError(http.StatusBadRequest, "route parameter subscription_id does not match id in JSON payload")
		}
		// Get Profile
		p := c.Get("profile").(*models.Profile)
		// Verify Profile ID matches ProfileID of Subscription to be Modified
		// No Modifying anyone else's settings
		t, err := models.GetAlertSubscriptionByID(db, &sID)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if p.ID != t.ProfileID {
			return echo.NewHTTPError(http.StatusUnauthorized, messages.Unauthorized)
		}
		sUpdated, err := models.UpdateMyAlertSubscription(db, &s)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, sUpdated)
	}
}
