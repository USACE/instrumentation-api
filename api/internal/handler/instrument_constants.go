package handler

import (
	"net/http"

	"github.com/USACE/instrumentation-api/api/internal/messages"
	"github.com/USACE/instrumentation-api/api/internal/models"
	"github.com/USACE/instrumentation-api/api/internal/utils"
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
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		// InstrumentID From RouteParams
		instrumentID, err := uuid.Parse(c.Param("instrument_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		// slugs already taken in the database
		slugsTaken, err := models.ListTimeseriesSlugsForInstrument(db, &instrumentID)
		if err != nil {
			return err
		}
		for idx := range tc.Items {
			// Verify object instrument_id matches routeParam
			if instrumentID != tc.Items[idx].InstrumentID {
				return echo.NewHTTPError(http.StatusBadRequest, messages.MatchRouteParam("`instrument_id`"))
			}
			// Assign Slug
			s, err := utils.NextUniqueSlug(tc.Items[idx].Name, slugsTaken)
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
			tc.Items[idx].Slug = s
			// Add slug to array of slugs originally fetched from the database
			// to catch duplicate names/slugs from the same bulk upload
			slugsTaken = append(slugsTaken, s)
		}
		tt, err := models.CreateInstrumentConstants(db, tc.Items)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusCreated, tt)
	}
}

// DeleteInstrumentConstant removes a timeseries as an Instrument Constant
func DeleteInstrumentConstant(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		instrumentID, err := uuid.Parse(c.Param("instrument_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		timeseriesID, err := uuid.Parse(c.Param("timeseries_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		err = models.DeleteInstrumentConstant(db, &instrumentID, &timeseriesID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, make(map[string]interface{}))
	}
}

// ListInstrumentConstants lists constants for a given instrument
func ListInstrumentConstants(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		instrumentID, err := uuid.Parse(c.Param("instrument_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		cc, err := models.ListInstrumentConstants(db, &instrumentID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, cc)
	}
}
