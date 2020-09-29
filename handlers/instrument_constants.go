package handlers

import (
	"net/http"

	"github.com/USACE/instrumentation-api/dbutils"
	"github.com/USACE/instrumentation-api/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

// CreateInstrumentConstants creates instrument constants (i.e. timeseries)
func CreateInstrumentConstants(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Get action information from context
		tc := models.TimeseriesCollection{}
		if err := c.Bind(&tc); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		// InstrumentID From RouteParams
		instrumentID, err := uuid.Parse(c.Param("instrument_id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		// slugs already taken in the database
		slugsTaken, err := models.ListTimeseriesSlugsForInstrument(db, &instrumentID)
		if err != nil {
			return err
		}
		for idx := range tc.Items {
			// Verify object instrument_id matches routeParam
			if instrumentID != tc.Items[idx].InstrumentID {
				return c.String(http.StatusBadRequest, "Object instrument_id does not match Route Param")
			}
			// Assign Slug
			s, err := dbutils.NextUniqueSlug(tc.Items[idx].Name, slugsTaken)
			if err != nil {
				return c.JSON(http.StatusBadRequest, err)
			}
			tc.Items[idx].Slug = s
			// Add slug to array of slugs originally fetched from the database
			// to catch duplicate names/slugs from the same bulk upload
			slugsTaken = append(slugsTaken, s)
		}
		tt, err := models.CreateInstrumentConstants(db, tc.Items)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusCreated, tt)
	}
}

// DeleteInstrumentConstant removes a timeseries as an Instrument Constant
func DeleteInstrumentConstant(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		instrumentID, err := uuid.Parse(c.Param("instrument_id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		timeseriesID, err := uuid.Parse(c.Param("timeseries_id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		err = models.DeleteInstrumentConstant(db, &instrumentID, &timeseriesID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, make(map[string]interface{}))
	}
}

// ListInstrumentConstants lists constants for a given instrument
func ListInstrumentConstants(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		instrumentID, err := uuid.Parse(c.Param("instrument_id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		cc, err := models.ListInstrumentConstants(db, &instrumentID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, cc)
	}
}
