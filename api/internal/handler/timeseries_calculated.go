package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"github.com/USACE/instrumentation-api/api/internal/message"
	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/USACE/instrumentation-api/api/internal/util"
)

// GetInstrumentCalculations retrieves an array of `Calculation`s associated with a particular instrument ID.
func (h *ApiHandler) GetInstrumentCalculations(c echo.Context) error {
	param := c.QueryParam("instrument_id")

	if param == "" {
		return echo.NewHTTPError(http.StatusBadRequest, message.MissingQueryParameter("`instrument_id`"))
	}

	// Instrument ID
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

// CreateCalculation creates a calculation.
func (h *ApiHandler) CreateCalculation(c echo.Context) error {
	var formula model.CalculatedTimeseries
	if err := c.Bind(&formula); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Create unique slug
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
	return c.JSON(http.StatusOK, echo.Map{
		"id": formula.ID,
	})
}

// UpdateCalculation updates a calculation.
func (h *ApiHandler) UpdateCalculation(c echo.Context) error {
	formulaID, err := uuid.Parse(c.Param("formula_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var formula model.CalculatedTimeseries
	if err := c.Bind(&formula); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Compare formula ID from Route Params to Calculation ID from Payload
	if formulaID != formula.ID {
		return echo.NewHTTPError(http.StatusBadRequest, message.BadRequest)
	}

	// Update slug when name is updated
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

	// Update in database
	if err := h.CalculatedTimeseriesService.UpdateCalculatedTimeseries(c.Request().Context(), formula); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, &formula)
}

// DeleteCalculation deletes a calculation.
func (h *ApiHandler) DeleteCalculation(c echo.Context) error {
	calculationID, err := uuid.Parse(c.Param("formula_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := h.CalculatedTimeseriesService.DeleteCalculatedTimeseries(c.Request().Context(), calculationID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}
