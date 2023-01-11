package handlers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"

	"github.com/USACE/instrumentation-api/dbutils"
	"github.com/USACE/instrumentation-api/models"
)

// GetInstrumentCalculations retrieves an array of `Calculation`s associated with a particular
// instrument ID.
//
// Query Parameters:
// - `instrument_id`: string
func GetInstrumentCalculations(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		param := c.QueryParam("instrument_id")

		if param == "" {
			return c.String(
				http.StatusBadRequest,
				"Missing query parameter 'instrument_id'",
			)
		}

		// Instrument ID
		instrumentID, err := uuid.Parse(param)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		formulas, err := models.GetInstrumentCalculations(db, &models.Instrument{ID: instrumentID})
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, formulas)
	}
}

// CreateCalculation creates a calculation.
//
// Parameters:
// - Body should be a calculation model in the database.
func CreateCalculation(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var formula models.Calculation
		if err := c.Bind(&formula); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		// Create unique slug
		var calculationSlug string = ""
		slugsTaken, err := models.ListCalculationSlugs(db)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		if formula.FormulaName == "" {
			calculationSlug, err = dbutils.NextUniqueSlug("New Formula", slugsTaken)
			formula.FormulaName = calculationSlug
		} else {
			calculationSlug, err = dbutils.NextUniqueSlug(formula.FormulaName, slugsTaken)
		}
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		formula.Slug = calculationSlug

		if err := models.CreateCalculation(db, &formula); err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, echo.Map{
			"id": formula.ID,
		})
	}
}

// UpdateCalculation updates a calculation.
//
// Paramaters:
// - `formula_id` should refer to the ID of a calculation in the database.
func UpdateCalculation(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		formulaID, err := uuid.Parse(c.Param("formula_id"))
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		var formula models.Calculation
		if err := c.Bind(&formula); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		// Compare formula ID from Route Params to Calculation ID from Payload
		if formulaID != formula.ID {
			return c.JSON(http.StatusBadRequest, models.DefaultMessageBadRequest)
		}

		// Update slug when name is updated
		var calculationSlug string = ""
		slugsTaken, err := models.ListCalculationSlugs(db)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		if formula.FormulaName == "" {
			calculationSlug, err = dbutils.NextUniqueSlug("New Formula", slugsTaken)
			formula.FormulaName = calculationSlug
		} else {
			calculationSlug, err = dbutils.NextUniqueSlug(formula.FormulaName, slugsTaken)
		}
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		formula.Slug = calculationSlug

		// Update in database
		if err := models.UpdateCalculation(db, &formula); err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, &formula)
	}
}

// DeleteCalculation deletes a calculation.
//
// Parameters:
// - `formula_id` should refer to the ID of a calculation in the database.
func DeleteCalculation(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		calculationID, err := uuid.Parse(c.Param("formula_id"))
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		if err := models.DeleteCalculation(db, calculationID); err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.NoContent(http.StatusOK)
	}
}
