package handler

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/USACE/instrumentation-api/api/internal/messages"
	"github.com/USACE/instrumentation-api/api/internal/models"
	"github.com/google/uuid"

	"github.com/labstack/echo/v4"
)

func (r *Router) Alerts(h *Handler) {
	r.ag.public.GET("/projects/:project_id/instruments/:instrument_id/alerts", h.ListAlertsForInstrument())
	r.ag.private.GET("/my_alerts", h.ListMyAlerts())
	r.ag.private.POST("/my_alerts/:alert_id/read", h.DoAlertRead())
	r.ag.private.POST("/my_alerts/:alert_id/unread", h.DoAlertUnread())
}

// ListAlertsForInstrument godoc
// @Summary      List alerts for a single instrument
// @Description  Return all alerts associated with the provided instrument UUID
// @Tags         alert
// @Accept
// @Produce      json
// @Param        project_id	path	UUID  true  "Project ID"
// @Param        instrument_id  path    UUID  true  "Instrument ID"
// @Success      200	{array}   model.Alert
// @Failure      400	{object}  echo.HTTPError
// @Failure      404	{object}  echo.HTTPError
// @Failure      500	{object}  echo.HTTPError
// @Router       /projects/{project_id}/instruments/{instrument_id}/alerts [get]
func (h *Handler) ListAlertsForInstrument() echo.HandlerFunc {
	return func(c echo.Context) error {
		instrumentID, err := uuid.Parse(c.Param("instrument_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		aa, err := h.AlertStore.GetAllAlertsForInstrument(c.Request().Context(), instrumentID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, aa)
	}
}

// ListMyAlerts  godoc
// @Summary      List subscribed alerts for a single user
// @Description  Return all alerts a profile is subscribed to
// @Tags         alert
// @Accept
// @Produce      json
// @Success      200  {array}   model.Alert
// @Failure      400  {object}  echo.HTTPError
// @Failure      404  {object}  echo.HTTPError
// @Failure      500  {object}  echo.HTTPError
// @Router       /my_alerts [get]
func (h *Handler) ListMyAlerts() echo.HandlerFunc {
	return func(c echo.Context) error {
		p := c.Get("profile").(*models.Profile)
		profileID := p.ID
		aa, err := h.AlertStore.GetAllAlertsForProfile(c.Request().Context(), profileID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, &aa)
	}
}

// DoAlertRead   godoc
// @Summary      marks an alert as read
// @Description  marks an alert as read based on provided profile ID and alert ID.
// @Description  returning the updated alert
// @Tags         alert
// @Accept
// @Produce      json
// @Param        alert_id path UUID true "Alert ID"
// @Success      200  {object}  model.Alert
// @Failure      400  {object}  echo.HTTPError
// @Failure      404  {object}  echo.HTTPError
// @Failure      500  {object}  echo.HTTPError
// @Router       /my_alerts/{alert_id}/read [post]
func (h *Handler) DoAlertRead() echo.HandlerFunc {
	return func(c echo.Context) error {
		p := c.Get("profile").(*models.Profile)
		profileID := p.ID
		alertID, err := uuid.Parse(c.Param("alert_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		a, err := h.AlertStore.DoAlertRead(c.Request().Context(), profileID, alertID)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return echo.NewHTTPError(http.StatusNotFound, messages.NotFound)
			}
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, a)
	}
}

// DoAlertUnread godoc
// @Summary      mmarks an alert as unread for a profile
// @Description  marks an alert as unread based on provided profile ID and alert ID.
// @Description  returning the updated alert
// @Tags         alert
// @Accept
// @Produce      json
// @Param        alert_id path UUID true "Alert ID"
// @Success      200  {object}  model.Alert
// @Failure      400  {object}  echo.HTTPError
// @Failure      404  {object}  echo.HTTPError
// @Failure      500  {object}  echo.HTTPError
// @Router       /my_alerts/{alert_id}/unread [post]
func (h *Handler) DoAlertUnread() echo.HandlerFunc {
	return func(c echo.Context) error {
		p := c.Get("profile").(*models.Profile)
		profileID := p.ID
		alertID, err := uuid.Parse(c.Param("alert_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		a, err := h.AlertStore.DoAlertUnread(c.Request().Context(), profileID, alertID)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return echo.NewHTTPError(http.StatusNotFound, messages.NotFound)
			}
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, a)
	}
}
