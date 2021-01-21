package handlers

import (
	"net/http"
	"time"

	"github.com/USACE/instrumentation-api/dbutils"
	"github.com/USACE/instrumentation-api/models"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

// ListInstruments returns instruments
func ListInstruments(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		nn, err := models.ListInstruments(db)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
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

// CreateInstruments accepts an array of instruments for bulk upload to the database
func CreateInstruments(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		// Sanatized instruments with ID, projectID, and slug assigned
		newInstrumentCollection := func(c echo.Context) (models.InstrumentCollection, error) {
			ic := models.InstrumentCollection{}
			if err := c.Bind(&ic); err != nil {
				return models.InstrumentCollection{}, err
			}

			// Get ProjectID of Instruments
			projectID, err := uuid.Parse(c.Param("project_id"))
			if err != nil {
				return models.InstrumentCollection{}, err
			}

			// slugs already taken in the database
			slugsTaken, err := models.ListInstrumentSlugs(db)
			if err != nil {
				return models.InstrumentCollection{}, err
			}

			// profile of user creating instruments
			p, err := profileFromContext(c, db)
			if err != nil {
				return models.InstrumentCollection{}, err
			}
			// timestamp
			t := time.Now()

			for idx := range ic.Items {
				// Assign ProjectID
				ic.Items[idx].ProjectID = &projectID
				// Assign Slug
				s, err := dbutils.NextUniqueSlug(ic.Items[idx].Name, slugsTaken)
				if err != nil {
					return models.InstrumentCollection{}, err
				}
				ic.Items[idx].Slug = s
				// Assign Creator
				ic.Items[idx].Creator = p.ID
				// Assign CreateDate
				ic.Items[idx].CreateDate = t
				// Add slug to array of slugs originally fetched from the database
				// to catch duplicate names/slugs from the same bulk upload
				slugsTaken = append(slugsTaken, s)
			}

			return ic, nil
		}

		// Instruments
		ic, err := newInstrumentCollection(c)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		// Validate POST
		if c.QueryParam("dry_run") == "true" {
			v, err := models.ValidateCreateInstruments(db, ic.Items)
			if err != nil {
				return c.JSON(http.StatusBadRequest, err)
			}
			return c.JSON(http.StatusOK, v)
		}

		// Actually POST
		nn, err := models.CreateInstruments(db, ic.Items)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusCreated, nn)
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

		// profile of user creating instruments
		p, err := profileFromContext(c, db)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		// timestamp
		t := time.Now()
		i.Updater, i.UpdateDate = &p.ID, &t

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

		return c.JSON(http.StatusOK, make(map[string]interface{}))
	}
}
