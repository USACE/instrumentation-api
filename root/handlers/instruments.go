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
		return c.JSON(http.StatusOK, models.ListInstruments(db))
	}
}

// GetInstrument returns a single instrument
func GetInstrument(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Malformed ID")
		}
		return c.JSON(http.StatusOK, models.GetInstrument(db, id))
	}
}

// CreateInstrument creates a single instrument
func CreateInstrument(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		i := &models.Instrument{}
		if err := c.Bind(i); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		// Ensure a new UUID
		i.ID = uuid.Must(uuid.NewRandom())
		// Assign Slug
		s, err := dbutils.NextUniqueSlug(i.Name, models.ListInstrumentSlugs(db))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		i.Slug = s

		if err := models.CreateInstrument(db, i); err != nil {
			return c.JSON(http.StatusForbidden, err)
		}
		// Send instrument
		return c.JSON(http.StatusCreated, i)
	}
}

// CreateInstrumentBulk accepts an array of instruments for bulk upload to the database
func CreateInstrumentBulk(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		instruments := []models.Instrument{}
		if err := c.Bind(&instruments); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		// slugs already taken in the database
		slugs := models.ListInstrumentSlugs(db)

		for idx := range instruments {
			// Assign UUID
			instruments[idx].ID = uuid.Must(uuid.NewRandom())
			// Assign Slug
			s, err := dbutils.NextUniqueSlug(instruments[idx].Name, slugs)
			if err != nil {
				return c.JSON(http.StatusBadRequest, err)
			}
			instruments[idx].Slug = s
			// Add slug to array of slugs originally fetched from the database
			// to catch duplicate names/slugs from the same bulk upload
			slugs = append(slugs, s)
		}

		if err := models.CreateInstrumentBulk(db, instruments); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		// Send instrument
		return c.JSON(http.StatusCreated, instruments)
	}
}

// UpdateInstrument modifys an existing instrument
func UpdateInstrument(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		// id from url params
		id, err := uuid.Parse(c.Param("id"))
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
		if err := models.UpdateInstrument(db, i); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		// return whole instrument
		return c.JSON(http.StatusOK, i)
	}
}

// DeleteInstrument deletes an existing instrument by ID
func DeleteInstrument(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Malformed ID")
		}

		if err := models.DeleteInstrument(db, id); err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		return c.NoContent(http.StatusOK)
	}
}
