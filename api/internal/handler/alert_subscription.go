package handler

import (
	"net/http"

	"github.com/USACE/instrumentation-api/api/internal/messages"
	"github.com/USACE/instrumentation-api/api/internal/model"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// SubscribeProfileToAlerts subscribes a profile to an alert
func (h ApiHandler) SubscribeProfileToAlerts(c echo.Context) error {
	p := c.Get("profile").(*model.Profile)
	profileID := p.ID

	alertConfigID, err := uuid.Parse(c.Param("alert_config_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	pa, err := h.AlertSubscriptionStore.SubscribeProfileToAlerts(c.Request().Context(), alertConfigID, profileID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, pa)
}

// UnsubscribeProfileToAlerts unsubscribes a profile to an alert
func (h ApiHandler) UnsubscribeProfileToAlerts(c echo.Context) error {
	p := c.Get("profile").(*model.Profile)
	profileID := p.ID

	alertConfigID, err := uuid.Parse(c.Param("alert_config_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = h.AlertSubscriptionStore.UnsubscribeProfileToAlerts(c.Request().Context(), alertConfigID, profileID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, make(map[string]interface{}))
}

// ListMyAlertSubscriptions returns all alerts you are subscribed to and settings
func (h ApiHandler) ListMyAlertSubscriptions(c echo.Context) error {
	p := c.Get("profile").(*model.Profile)
	profileID := p.ID
	ss, err := h.AlertSubscriptionStore.ListMyAlertSubscriptions(c.Request().Context(), profileID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, &ss)
}

// UpdateMyAlertSubscription updates settings for an alert subscription
func (h ApiHandler) UpdateMyAlertSubscription(c echo.Context) error {
	var s model.AlertSubscription
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
	p := c.Get("profile").(*model.Profile)
	// Verify Profile ID matches ProfileID of Subscription to be Modified
	// No Modifying anyone else's settings
	t, err := h.AlertSubscriptionStore.GetAlertSubscriptionByID(c.Request().Context(), sID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if p.ID != t.ProfileID {
		return echo.NewHTTPError(http.StatusUnauthorized, messages.Unauthorized)
	}
	sUpdated, err := h.AlertSubscriptionStore.UpdateMyAlertSubscription(c.Request().Context(), s)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, sUpdated)
}
