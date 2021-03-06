package handlers

import (
	"github.com/USACE/instrumentation-api/dbutils"
	"github.com/USACE/instrumentation-api/models"
	ts "github.com/USACE/instrumentation-api/timeseries"

	"net/http"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

// ListTimeseries returns an array of timeseries
func ListTimeseries(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		tt, err := models.ListTimeseries(db)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusOK, tt)
	}
}

// GetTimeseries returns a single timeseries
func GetTimeseries(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		tsID, err := uuid.Parse(c.Param("timeseries_id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Malformed ID")
		}
		t, err := models.GetTimeseries(db, &tsID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, t)
	}
}

// ListInstrumentTimeseries lists timeseries for an instrument
func ListInstrumentTimeseries(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		pID, err := uuid.Parse(c.Param("project_id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Malformed ID")
		}
		nID, err := uuid.Parse(c.Param("instrument_id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Malformed ID")
		}
		tt, err := models.ListInstrumentTimeseries(db, &pID, &nID)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, tt)
	}
}

// ListInstrumentGroupTimeseries lists timeseries for instruments in an instrument group
func ListInstrumentGroupTimeseries(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		gID, err := uuid.Parse(c.Param("instrument_group_id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Malformed ID")
		}
		tt, err := models.ListInstrumentGroupTimeseries(db, &gID)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, tt)
	}
}

// ListProjectTimeseries lists all timeseries for a single project
func ListProjectTimeseries(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		pID, err := uuid.Parse(c.Param("project_id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Malformed ID")
		}
		tt, err := models.ListProjectTimeseries(db, &pID)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, tt)
	}
}

// CreateTimeseries accepts a timeseries object or array of timeseries objects
// Can handle objects with or without TimeseriesMeasurements
func CreateTimeseries(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		tc := models.TimeseriesCollection{}
		if err := c.Bind(&tc); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		// slugs already taken in the database
		slugsTaken, err := models.ListTimeseriesSlugs(db)
		if err != nil {
			return err
		}
		for idx := range tc.Items {
			// Assign UUID
			tc.Items[idx].ID = uuid.Must(uuid.NewRandom())
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

		tt, err := models.CreateTimeseries(db, tc.Items)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		return c.JSON(http.StatusCreated, tt)
	}
}

// UpdateTimeseries updates a single timeseries
func UpdateTimeseries(db *sqlx.DB) echo.HandlerFunc {
	// UpdateInstrumentGroup modifies an existing instrument_group
	return func(c echo.Context) error {

		// id from url params
		id, err := uuid.Parse(c.Param("timeseries_id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Malformed ID")
		}
		// id from request
		t := ts.Timeseries{}
		if err := c.Bind(&t); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		// check :id in url params matches id in request body
		if id != t.ID {
			return c.String(
				http.StatusBadRequest,
				"url parameter id does not match object id in body",
			)
		}
		// update
		tUpdated, err := models.UpdateTimeseries(db, &t)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		// return updated instrument
		return c.JSON(http.StatusOK, tUpdated)
	}
}

// DeleteTimeseries deletes a single timeseries
func DeleteTimeseries(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// id from url params
		id, err := uuid.Parse(c.Param("timeseries_id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		if err := models.DeleteTimeseries(db, &id); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusOK, make(map[string]interface{}))
	}
}
