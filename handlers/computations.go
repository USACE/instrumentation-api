package handlers

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"

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

// GetComputations retrieves an array of `Formula`s associated with a particular
// instrument ID.
//
// Parameters:
// - `instrument_id`: string
func GetComputations(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		instrumentID, err := uuid.Parse(c.Param("instrument_id"))
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		formulas, err := models.GetFormulas(db, &models.Instrument{ID: instrumentID})
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, formulas)
	}
}

// CreateComputation for a given instrument.
//
// Parameters:
// - Body should be a computation model in the database.
func CreateComputation(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var formula models.Formula
		if err := c.Bind(formula); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		if err := models.CreateFormula(db, &formula); err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.NoContent(http.StatusOK)
	}
}

// DeleteComputation for a given instrument.
//
// Parameters:
// - `computation_id` should refer to the ID of a computation in the database.
func DeleteComputation(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		computationID, err := uuid.Parse(c.Param("computation_id"))
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		if err := models.DeleteFormula(db, computationID); err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.NoContent(http.StatusOK)
	}
}
