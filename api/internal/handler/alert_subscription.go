package handler

import (
	"net/http"

	"github.com/USACE/instrumentation-api/api/internal/message"
	"github.com/USACE/instrumentation-api/api/internal/model"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// SubscribeProfileToAlerts godoc
//
//	@Summary subscribes a profile to an alert
//	@Tags alert-subscription
//	@Produce json
//	@Param project_id path string true "project uuid" Format(uuid)
//	@Param instrument_id path string true "instrument uuid" Format(uuid)
//	@Param alert_config_id path string true "alert config uuid" Format(uuid)
//	@Param key query string false "api key"
//	@Success 200 {object} model.AlertSubscription
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/instruments/{instrument_id}/alert_configs/{alert_config_id}/subscribe [post]
//	@Security Bearer
func (h *ApiHandler) SubscribeProfileToAlerts(c echo.Context) error {
	p := c.Get("profile").(model.Profile)
	profileID := p.ID

	alertConfigID, err := uuid.Parse(c.Param("alert_config_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	pa, err := h.AlertSubscriptionService.SubscribeProfileToAlerts(c.Request().Context(), alertConfigID, profileID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, pa)
}

// UnsubscribeProfileToAlerts godoc
//
//	@Summary unsubscribes a profile to an alert
//	@Tags alert-subscription
//	@Produce json
//	@Param project_id path string true "project uuid" Format(uuid)
//	@Param instrument_id path string true "instrument uuid" Format(uuid)
//	@Param alert_config_id path string true "alert config uuid" Format(uuid)
//	@Param key query string false "api key"
//	@Success 200 {object} map[string]interface{}
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/instruments/{instrument_id}/alert_configs/{alert_config_id}/unsubscribe [post]
//	@Security Bearer
func (h *ApiHandler) UnsubscribeProfileToAlerts(c echo.Context) error {
	p := c.Get("profile").(model.Profile)
	profileID := p.ID

	alertConfigID, err := uuid.Parse(c.Param("alert_config_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = h.AlertSubscriptionService.UnsubscribeProfileToAlerts(c.Request().Context(), alertConfigID, profileID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, make(map[string]interface{}))
}

// ListMyAlertSubscriptions godoc
//
//	@Summary lists all alerts subscribed to by the current profile
//	@Tags alert-subscription
//	@Produce json
//	@Param key query string false "api key"
//	@Success 200 {array} model.AlertSubscription
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /my_alert_subscriptions [get]
//	@Security Bearer
func (h *ApiHandler) ListMyAlertSubscriptions(c echo.Context) error {
	p := c.Get("profile").(model.Profile)
	profileID := p.ID
	ss, err := h.AlertSubscriptionService.ListMyAlertSubscriptions(c.Request().Context(), profileID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, ss)
}

// UpdateMyAlertSubscription godoc
//
//	@Summary updates settings for an alert subscription
//	@Tags alert-subscription
//	@Accept json
//	@Produce json
//	@Param alert_subscription_id path string true "alert subscription id" Format(uuid)
//	@Param alert_subscription body model.AlertSubscription true "alert subscription payload"
//	@Param key query string false "api key"
//	@Success 200 {array} model.AlertSubscription
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /alert_subscriptions/{alert_subscription_id} [put]
//	@Security Bearer
func (h *ApiHandler) UpdateMyAlertSubscription(c echo.Context) error {
	var s model.AlertSubscription
	if err := c.Bind(&s); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	sID, err := uuid.Parse(c.Param("alert_subscription_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if s.ID != sID {
		return echo.NewHTTPError(http.StatusBadRequest, "route parameter subscription_id does not match id in JSON payload")
	}
	p := c.Get("profile").(model.Profile)
	t, err := h.AlertSubscriptionService.GetAlertSubscriptionByID(c.Request().Context(), sID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if p.ID != t.ProfileID {
		return echo.NewHTTPError(http.StatusUnauthorized, message.Unauthorized)
	}
	sUpdated, err := h.AlertSubscriptionService.UpdateMyAlertSubscription(c.Request().Context(), s)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, sUpdated)
}
