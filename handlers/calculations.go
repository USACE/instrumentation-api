package handlers

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"

	"github.com/USACE/instrumentation-api/dbutils"
	"github.com/USACE/instrumentation-api/models"
)

// This is an endpoint for debugging at this time
func ComputedTimeseries(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Instrument ID
		instrumentID, err := uuid.Parse(c.Param("instrument_id"))
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		instrumentIDs := make([]uuid.UUID, 1)
		instrumentIDs[0] = instrumentID
		// Time Window
		timeWindow := models.TimeWindow{
			After:  time.Date(2020, 1, 3, 0, 0, 0, 0, time.UTC),
			Before: time.Date(2021, 1, 5, 0, 0, 0, 0, time.UTC),
		}
		// Interval - Hard Code at 1 Hour
		interval := time.Hour

		tt, err := models.ComputedTimeseries(db, instrumentIDs, &timeWindow, &interval)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, &tt)
	}
}

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

		slugsTaken, err := models.ListCalculationSlugs(db)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		calculationSlug, err := dbutils.NextUniqueSlug(formula.FormulaName, slugsTaken)
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
		calculationID, err := uuid.Parse(c.Param("formula_id"))
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		var calculation models.Calculation
		if err := c.Bind(&calculation); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		// Compare Calculation ID from Route Params to Calculation ID from Payload
		if calculationID != calculation.ID {
			return c.JSON(http.StatusBadRequest, models.DefaultMessageBadRequest)
		}

		if err := models.UpdateCalculation(db, &calculation); err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, &calculation)
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
