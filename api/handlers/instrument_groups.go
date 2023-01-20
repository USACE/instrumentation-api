package handlers

import (
	"net/http"
	"time"

	"github.com/USACE/instrumentation-api/api/dbutils"
	"github.com/USACE/instrumentation-api/api/models"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

// ListInstrumentGroups returns instrument groups
func ListInstrumentGroups(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		groups, err := models.ListInstrumentGroups(db)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, groups)
	}
}

// GetInstrumentGroup returns single instrument group
func GetInstrumentGroup(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := uuid.Parse(c.Param("instrument_group_id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Malformed ID")
		}
		g, err := models.GetInstrumentGroup(db, id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, g)
	}
}

// CreateInstrumentGroup accepts an array of instruments for bulk upload to the database
func CreateInstrumentGroup(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		gc := models.InstrumentGroupCollection{}
		if err := c.Bind(&gc); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		// slugs already taken in the database
		slugsTaken, err := models.ListInstrumentGroupSlugs(db)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		// profile
		p := c.Get("profile").(*models.Profile)

		// timestamp
		t := time.Now()

		for idx := range gc.Items {
			// Creator
			gc.Items[idx].Creator = p.ID
			// CreateDate
			gc.Items[idx].CreateDate = t
			// Assign Slug
			s, err := dbutils.NextUniqueSlug(gc.Items[idx].Name, slugsTaken)
			if err != nil {
				return c.JSON(http.StatusBadRequest, err)
			}
			gc.Items[idx].Slug = s
			// Add slug to array of slugs originally fetched from the database
			// to catch duplicate names/slugs from the same bulk upload
			slugsTaken = append(slugsTaken, s)
		}

		gg, err := models.CreateInstrumentGroup(db, gc.Items)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		// Send instrumentgroup
		return c.JSON(http.StatusCreated, gg)
	}
}

// UpdateInstrumentGroup modifies an existing instrument_group
func UpdateInstrumentGroup(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		// id from url params
		id, err := uuid.Parse(c.Param("instrument_group_id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Malformed ID")
		}
		// id from request
		g := models.InstrumentGroup{ID: id}
		if err := c.Bind(&g); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		// check :id in url params matches id in request body
		if id != g.ID {
			return c.String(
				http.StatusBadRequest,
				"url parameter id does not match object id in body",
			)
		}

		// profile information and timestamp
		p := c.Get("profile").(*models.Profile)

		t := time.Now()
		g.Updater, g.UpdateDate = &p.ID, &t

		// update
		gUpdated, err := models.UpdateInstrumentGroup(db, &g)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		// return updated instrument
		return c.JSON(http.StatusOK, gUpdated)
	}
}

// DeleteFlagInstrumentGroup sets the instrument group deleted flag true
func DeleteFlagInstrumentGroup(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := uuid.Parse(c.Param("instrument_group_id"))
		if err != nil {
			return c.JSON(http.StatusNotFound, err)
		}
		if err := models.DeleteFlagInstrumentGroup(db, id); err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, make(map[string]interface{}))
	}
}

// ListInstrumentGroupInstruments returns a list of instruments for a provided instrument group
func ListInstrumentGroupInstruments(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := uuid.Parse(c.Param("instrument_group_id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Malformed ID")
		}
		nn, err := models.ListInstrumentGroupInstruments(db, id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, nn)
	}
}

// CreateInstrumentGroupInstruments adds an instrument to an instrument group
func CreateInstrumentGroupInstruments(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// instrument_group_id
		instrumentGroupID, err := uuid.Parse(c.Param("instrument_group_id"))

		if err != nil || instrumentGroupID == uuid.Nil {
			return c.JSON(http.StatusBadRequest, models.DefaultMessageBadRequest)
		}
		// Instrument
		i := new(models.Instrument)
		if err := c.Bind(i); err != nil || i.ID == uuid.Nil {
			return c.JSON(http.StatusBadRequest, models.DefaultMessageBadRequest)
		}

		if err := models.CreateInstrumentGroupInstruments(db, instrumentGroupID, i.ID); err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusCreated, make(map[string]interface{}))
	}
}

// DeleteInstrumentGroupInstruments removes an instrument from an instrument group
func DeleteInstrumentGroupInstruments(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// instrument_group_id
		instrumentGroupID, err := uuid.Parse(c.Param("instrument_group_id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Malformed ID")
		}

		// instrument
		instrumentID, err := uuid.Parse(c.Param("instrument_id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Malformed ID")
		}

		if err := models.DeleteInstrumentGroupInstruments(db, instrumentGroupID, instrumentID); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		return c.JSON(http.StatusOK, make(map[string]interface{}))
	}
}
