package handler

import (
	"net/http"
	"strings"

	"github.com/USACE/instrumentation-api/api/internal/message"
	_ "github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// ListProjectSubmittals godoc
//
//	@Summary lists all submittals for a project
//	@Tags submittal
//	@Produce json
//	@Param project_id path string true "project uuid" Format(uuid)
//	@Param missing query bool false "filter by missing projects only"
//	@Success 200 {array} model.Submittal
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /projects/{project_id}/submittals [get]
func (h *ApiHandler) ListProjectSubmittals(c echo.Context) error {
	id, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}

	var fmo bool
	mo := c.QueryParam("missing")
	if strings.ToLower(mo) == "true" {
		fmo = true
	}

	subs, err := h.SubmittalService.ListProjectSubmittals(c.Request().Context(), id, fmo)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, subs)
}

// ListInstrumentSubmittals godoc
//
//	@Summary lists all submittals for an instrument
//	@Tags submittal
//	@Produce json
//	@Param instrument_id path string true "instrument uuid" Format(uuid)
//	@Param missing query bool false "filter by missing projects only"
//	@Success 200 {array} model.Submittal
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /instruments/{instrument_id}/submittals [get]
func (h *ApiHandler) ListInstrumentSubmittals(c echo.Context) error {
	id, err := uuid.Parse(c.Param("instrument_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}

	var fmo bool
	mo := c.QueryParam("missing")
	if strings.ToLower(mo) == "true" {
		fmo = true
	}

	subs, err := h.SubmittalService.ListInstrumentSubmittals(c.Request().Context(), id, fmo)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, subs)
}

// ListAlertConfigSubmittals godoc
//
//	@Summary lists all submittals for an instrument
//	@Tags submittal
//	@Produce json
//	@Param alert_config_id path string true "alert config uuid" Format(uuid)
//	@Success 200 {array} model.Submittal
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /alert_configs/{alert_config_id}/submittals [get]
func (h *ApiHandler) ListAlertConfigSubmittals(c echo.Context) error {
	id, err := uuid.Parse(c.Param("alert_config_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}

	var fmo bool
	mo := c.QueryParam("missing")
	if strings.ToLower(mo) == "true" {
		fmo = true
	}

	subs, err := h.SubmittalService.ListAlertConfigSubmittals(c.Request().Context(), id, fmo)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, subs)
}

// VerifyMissingSubmittal godoc
//
//	@Summary verifies the specified submittal is "missing" and will not be completed
//	@Tags submittal
//	@Produce json
//	@Param submittal_id path string true "submittal uuid" Format(uuid)
//	@Success 200 {object} map[string]interface{}
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /submittals/{submittal_id}/verify_missing [put]
func (h *ApiHandler) VerifyMissingSubmittal(c echo.Context) error {
	id, err := uuid.Parse(c.Param("submittal_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	if err := h.SubmittalService.VerifyMissingSubmittal(c.Request().Context(), id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"submittal_id": id})
}

// VerifyMissingAlertConfigSubmittals godoc
//
//	@Summary verifies all current submittals for the alert config are "missing" and will not be completed
//	@Tags submittal
//	@Produce json
//	@Param alert_config_id path string true "alert config uuid" Format(uuid)
//	@Success 200 {object} map[string]interface{}
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /alert_configs/{alert_config_id}/submittals/verify_missing [put]
func (h *ApiHandler) VerifyMissingAlertConfigSubmittals(c echo.Context) error {
	id, err := uuid.Parse(c.Param("alert_config_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	if err := h.SubmittalService.VerifyMissingAlertConfigSubmittals(c.Request().Context(), id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"alert_config_id": id})
}
