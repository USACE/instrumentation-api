package handler

import (
	"net/http"
	"strings"

	"github.com/USACE/instrumentation-api/api/internal/messages"
	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// ListProjectSubmittals lists all submittals for a project
func (h ApiHandler) ListProjectSubmittals(c echo.Context) error {
	id, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, messages.MalformedID)
	}

	var fmo bool
	mo := c.QueryParam("missing")
	if strings.ToLower(mo) == "true" {
		fmo = true
	}

	subs, err := model.ListProjectSubmittals(db, &id, fmo)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, subs)
}

// ListInstrumentSubmittals lists all submittals for an instrument
func (h ApiHandler) ListInstrumentSubmittals(c echo.Context) error {
	id, err := uuid.Parse(c.Param("instrument_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, messages.MalformedID)
	}

	var fmo bool
	mo := c.QueryParam("missing")
	if strings.ToLower(mo) == "true" {
		fmo = true
	}

	subs, err := model.ListInstrumentSubmittals(db, &id, fmo)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, subs)
}

// ListAlertConfigSubmittals lists all submittals for an instrument
func (h ApiHandler) ListAlertConfigSubmittals(c echo.Context) error {
	id, err := uuid.Parse(c.Param("alert_config_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, messages.MalformedID)
	}

	var fmo bool
	mo := c.QueryParam("missing")
	if strings.ToLower(mo) == "true" {
		fmo = true
	}

	subs, err := model.ListAlertConfigSubmittals(db, &id, fmo)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, subs)
}

// DeleteFlagProject sets the instrument group deleted flag true
func (h ApiHandler) VerifyMissingSubmittal(c echo.Context) error {
	id, err := uuid.Parse(c.Param("submittal_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	if err := model.VerifyMissingSubmittal(db, &id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"submittal_id": id})
}

// DeleteFlagProject sets the instrument group deleted flag true
func (h ApiHandler) VerifyMissingAlertConfigSubmittals(c echo.Context) error {
	id, err := uuid.Parse(c.Param("alert_config_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	if err := model.VerifyMissingAlertConfigSubmittals(db, &id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"alert_config_id": id})
}
