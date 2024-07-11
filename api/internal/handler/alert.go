package handler

import (
	"net/http"

	"github.com/USACE/instrumentation-api/api/internal/httperr"
	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"

	"github.com/labstack/echo/v4"
)

// ListAlertsForInstrument godoc
//
//	@Summary lists alerts for a single instrument
//	@Description list all alerts associated an instrument
//	@Tags alert
//	@Produce json
//	@Param project_id path string true "project uuid" Format(uuid)
//	@Param instrument_id path string true "instrument uuid" Format(uuid)
//	@Success 200 {array} model.Alert
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/instruments/{instrument_id}/alerts [get]
func (h *ApiHandler) ListAlertsForInstrument(c echo.Context) error {
	instrumentID, err := uuid.Parse(c.Param("instrument_id"))
	if err != nil {
		return httperr.BadRequest(err)
	}
	aa, err := h.AlertService.GetAllAlertsForInstrument(c.Request().Context(), instrumentID)
	if err != nil {
		return httperr.InternalServerError(err)
	}
	return c.JSON(http.StatusOK, aa)
}

// ListMyAlerts godoc
//
//	@Summary lists subscribed alerts for a single user
//	@Description list all alerts a profile is subscribed to
//	@Tags alert
//	@Produce json
//	@Param key query string false "api key"
//	@Success 200 {array} model.Alert
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /my_alerts [get]
//	@Security Bearer
func (h *ApiHandler) ListMyAlerts(c echo.Context) error {
	p := c.Get("profile").(model.Profile)
	profileID := p.ID
	aa, err := h.AlertService.GetAllAlertsForProfile(c.Request().Context(), profileID)
	if err != nil {
		return httperr.InternalServerError(err)
	}
	return c.JSON(http.StatusOK, aa)
}

// DoAlertRead godoc
//
//	@Summary marks an alert as read
//	@Description marks an alert as read for a profile
//	@Description returning the updated alert
//	@Tags alert
//	@Produce json
//	@Param alert_id path string true "alert uuid" Format(uuid)
//	@Param key query string false "api key"
//	@Success 200 {object} model.Alert
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /my_alerts/{alert_id}/read [post]
//	@Security Bearer
func (h *ApiHandler) DoAlertRead(c echo.Context) error {
	p := c.Get("profile").(model.Profile)
	profileID := p.ID
	alertID, err := uuid.Parse(c.Param("alert_id"))
	if err != nil {
		return httperr.MalformedID(err)
	}
	a, err := h.AlertService.DoAlertRead(c.Request().Context(), profileID, alertID)
	if err != nil {
		return httperr.InternalServerError(err)
	}
	return c.JSON(http.StatusOK, a)
}

// DoAlertUnread godoc
//
//	@Summary marks an alert as unread for a profile
//	@Description marks an alert as unread based on provided profile ID and alert ID.
//	@Description returning the updated alert
//	@Tags alert
//	@Produce json
//	@Param alert_id path string true "alert uuid" Format(uuid)
//	@Param key query string false "api key"
//	@Success 200 {object} model.Alert
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /my_alerts/{alert_id}/unread [post]
//	@Security Bearer
func (h *ApiHandler) DoAlertUnread(c echo.Context) error {
	p := c.Get("profile").(model.Profile)
	profileID := p.ID
	alertID, err := uuid.Parse(c.Param("alert_id"))
	if err != nil {
		return httperr.MalformedID(err)
	}
	a, err := h.AlertService.DoAlertUnread(c.Request().Context(), profileID, alertID)
	if err != nil {
		return httperr.InternalServerError(err)
	}
	return c.JSON(http.StatusOK, a)
}
