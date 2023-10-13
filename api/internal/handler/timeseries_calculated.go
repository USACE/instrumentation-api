package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"github.com/USACE/instrumentation-api/api/internal/message"
	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/USACE/instrumentation-api/api/internal/util"
)

// GetInstrumentCalculations godoc
//
//	@Summary lists calculations associated with an instrument
//	@Tags formula
//	@Produce json
//	@Success 200 {array} model.CalculatedTimeseries
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /formulas [get]
func (h *ApiHandler) GetInstrumentCalculations(c echo.Context) error {
	param := c.QueryParam("instrument_id")
	if param == "" {
		return echo.NewHTTPError(http.StatusBadRequest, message.MissingQueryParameter("`instrument_id`"))
	}
	instrumentID, err := uuid.Parse(param)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	formulas, err := h.CalculatedTimeseriesService.GetAllCalculatedTimeseriesForInstrument(c.Request().Context(), instrumentID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, formulas)
}

// CreateCalculation godoc
//
//	@Summary creates a calculation
//	@Tags formula
//	@Produce json
//	@Success 200 {object} map[string]interface{}
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /formulas [post]
func (h *ApiHandler) CreateCalculation(c echo.Context) error {
	var formula model.CalculatedTimeseries
	if err := c.Bind(&formula); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var calculationSlug string = ""
	slugsTaken, err := h.CalculatedTimeseriesService.ListCalculatedTimeseriesSlugs(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if formula.FormulaName == "" {
		calculationSlug, err = util.NextUniqueSlug("New Formula", slugsTaken)
		formula.FormulaName = calculationSlug
	} else {
		calculationSlug, err = util.NextUniqueSlug(formula.FormulaName, slugsTaken)
	}
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	formula.Slug = calculationSlug

	if err := h.CalculatedTimeseriesService.CreateCalculatedTimeseries(c.Request().Context(), formula); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"id": formula.ID})
}

// UpdateCalculation godoc
//
//	@Summary updates a calculation
//	@Tags formula
//	@Produce json
//	@Param formula_id path string true "formula uuid" Format(uuid)
//	@Success 200 {array} model.CalculatedTimeseries
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /formulas/{formula_id} [put]
func (h *ApiHandler) UpdateCalculation(c echo.Context) error {
	formulaID, err := uuid.Parse(c.Param("formula_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var formula model.CalculatedTimeseries
	if err := c.Bind(&formula); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if formulaID != formula.ID {
		return echo.NewHTTPError(http.StatusBadRequest, message.BadRequest)
	}

	var calculationSlug string = ""
	slugsTaken, err := h.CalculatedTimeseriesService.ListCalculatedTimeseriesSlugs(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if formula.FormulaName == "" {
		calculationSlug, err = util.NextUniqueSlug("New Formula", slugsTaken)
		formula.FormulaName = calculationSlug
	} else {
		calculationSlug, err = util.NextUniqueSlug(formula.FormulaName, slugsTaken)
	}
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	formula.Slug = calculationSlug

	if err := h.CalculatedTimeseriesService.UpdateCalculatedTimeseries(c.Request().Context(), formula); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, formula)
}

// DeleteCalculation godoc
//
//	@Summary deletes a calculation
//	@Tags formula
//	@Produce json
//	@Param formula_id path string true "formula uuid" Format(uuid)
//	@Success 200 {object} map[string]interface{}
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /formulas/{formula_id} [delete]
func (h *ApiHandler) DeleteCalculation(c echo.Context) error {
	calculationID, err := uuid.Parse(c.Param("formula_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := h.CalculatedTimeseriesService.DeleteCalculatedTimeseries(c.Request().Context(), calculationID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, make(map[string]interface{}))
}
