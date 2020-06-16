package handlers

import (
	"api/root/dbutils"
	"api/root/models"
	"net/http"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
)

// ListInstruments returns instruments
func ListInstruments(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		nn, err := models.ListInstruments(db)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusOK, nn)
	}
}

// GetInstrumentCount returns the total number of non deleted instruments in the system
func GetInstrumentCount(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		count, err := models.GetInstrumentCount(db)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, map[string]interface{}{"instrument_count": count})
	}
}

// GetInstrument returns a single instrument
func GetInstrument(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := uuid.Parse(c.Param("instrument_id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Malformed ID")
		}
		n, err := models.GetInstrument(db, &id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, n)
	}
}

// CreateInstrumentBulk accepts an array of instruments for bulk upload to the database
func CreateInstrumentBulk(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		ic := models.InstrumentCollection{}
		if err := c.Bind(&ic); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		// slugs already taken in the database
		slugsTaken, err := models.ListInstrumentSlugs(db)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		for idx := range ic.Items {
			// Assign UUID
			ic.Items[idx].ID = uuid.Must(uuid.NewRandom())
			// Assign Slug
			s, err := dbutils.NextUniqueSlug(ic.Items[idx].Name, slugsTaken)
			if err != nil {
				return c.JSON(http.StatusBadRequest, err)
			}
			ic.Items[idx].Slug = s
			// Add slug to array of slugs originally fetched from the database
			// to catch duplicate names/slugs from the same bulk upload
			slugsTaken = append(slugsTaken, s)
		}

		if err := models.CreateInstrumentBulk(db, ic.Items); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		// Send instrument
		return c.JSON(http.StatusCreated, ic.Items)
	}
}

// UpdateInstrument modifies an existing instrument
func UpdateInstrument(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		// id from url params
		id, err := uuid.Parse(c.Param("instrument_id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Malformed ID")
		}
		// id from request
		i := &models.Instrument{ID: id}
		if err := c.Bind(i); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		// check :id in url params matches id in request body
		if id != i.ID {
			return c.String(
				http.StatusBadRequest,
				"url parameter id does not match object id in body",
			)
		}
		// update
		iUpdated, err := models.UpdateInstrument(db, i)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		// return updated instrument
		return c.JSON(http.StatusOK, iUpdated)
	}
}

// DeleteFlagInstrument changes deleted flag true for an instrument
func DeleteFlagInstrument(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := uuid.Parse(c.Param("instrument_id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Malformed ID")
		}

		if err := models.DeleteFlagInstrument(db, &id); err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		return c.NoContent(http.StatusOK)
	}
}
