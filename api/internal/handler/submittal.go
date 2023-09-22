package handler

import (
	"net/http"
	"strings"

	"github.com/USACE/instrumentation-api/api/internal/messages"
	"github.com/USACE/instrumentation-api/api/internal/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

// ListProjectSubmittals lists all submittals for a project
func ListProjectSubmittals(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := uuid.Parse(c.Param("project_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, messages.MalformedID)
		}

		var fmo bool
		mo := c.QueryParam("missing")
		if strings.ToLower(mo) == "true" {
			fmo = true
		}

		subs, err := models.ListProjectSubmittals(db, &id, fmo)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, subs)
	}
}

// ListInstrumentSubmittals lists all submittals for an instrument
func ListInstrumentSubmittals(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := uuid.Parse(c.Param("instrument_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, messages.MalformedID)
		}

		var fmo bool
		mo := c.QueryParam("missing")
		if strings.ToLower(mo) == "true" {
			fmo = true
		}

		subs, err := models.ListInstrumentSubmittals(db, &id, fmo)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, subs)
	}
}

// ListAlertConfigSubmittals lists all submittals for an instrument
func ListAlertConfigSubmittals(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := uuid.Parse(c.Param("alert_config_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, messages.MalformedID)
		}

		var fmo bool
		mo := c.QueryParam("missing")
		if strings.ToLower(mo) == "true" {
			fmo = true
		}

		subs, err := models.ListAlertConfigSubmittals(db, &id, fmo)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, subs)
	}
}

// DeleteFlagProject sets the instrument group deleted flag true
func VerifyMissingSubmittal(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := uuid.Parse(c.Param("submittal_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		if err := models.VerifyMissingSubmittal(db, &id); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, map[string]interface{}{"submittal_id": id})
	}
}

// DeleteFlagProject sets the instrument group deleted flag true
func VerifyMissingAlertConfigSubmittals(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := uuid.Parse(c.Param("alert_config_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		if err := models.VerifyMissingAlertConfigSubmittals(db, &id); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, map[string]interface{}{"alert_config_id": id})
	}
}
