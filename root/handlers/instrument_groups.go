package handlers

import (
	"api/root/dbutils"
	"api/root/models"
	"net/http"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
)

// ListInstrumentGroups returns instrument groups
func ListInstrumentGroups(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.ListInstrumentGroups(db))
	}
}

// GetInstrumentGroup returns single instrument group
func GetInstrumentGroup(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Malformed ID")
		}
		return c.JSON(http.StatusOK, models.GetInstrumentGroup(db, id))
	}
}

// CreateInstrumentGroup creates a single instrument group
func CreateInstrumentGroup(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		g := &models.InstrumentGroup{}
		if err := c.Bind(g); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		// Ensure a new UUID
		g.ID = uuid.Must(uuid.NewRandom())
		// Assign Slug
		s, err := dbutils.NextUniqueSlug(g.Name, models.ListInstrumentGroupSlugs(db))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		g.Slug = s

		err = models.CreateInstrumentGroup(db, g)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		// Send back group
		return c.JSON(http.StatusCreated, g)
	}
}

// CreateInstrumentGroupBulk accepts an array of instruments for bulk upload to the database
func CreateInstrumentGroupBulk(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		groups := []models.InstrumentGroup{}
		if err := c.Bind(&groups); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		// slugs already taken in the database
		slugs := models.ListInstrumentGroupSlugs(db)

		for idx := range groups {
			// Assign UUID
			groups[idx].ID = uuid.Must(uuid.NewRandom())
			// Assign Slug
			s, err := dbutils.NextUniqueSlug(groups[idx].Name, slugs)
			if err != nil {
				return c.JSON(http.StatusBadRequest, err)
			}
			groups[idx].Slug = s
			// Add slug to array of slugs originally fetched from the database
			// to catch duplicate names/slugs from the same bulk upload
			slugs = append(slugs, s)
		}

		if err := models.CreateInstrumentGroupBulk(db, groups); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		// Send instrumentgroup
		return c.JSON(http.StatusCreated, groups)
	}
}

// DeleteInstrumentGroup deletes an instrument group and any instrument_group_instruments associations
func DeleteInstrumentGroup(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusNotFound, err)
		}
		if err = models.DeleteInstrumentGroup(db, id); err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.NoContent(http.StatusOK)
	}
}

// ListInstrumentGroupInstruments returns a list of instruments for a provided instrument group
func ListInstrumentGroupInstruments(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Malformed ID")
		}
		return c.JSON(http.StatusOK, models.ListInstrumentGroupInstruments(db, id))
	}
}

// CreateInstrumentGroupInstruments adds an instrument to an instrument group
func CreateInstrumentGroupInstruments(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// instrument_group_id
		instrumentGroupID, err := uuid.Parse(c.Param("id"))

		if err != nil || instrumentGroupID == uuid.Nil {
			return c.NoContent(http.StatusBadRequest)
		}
		// instrument
		i := new(models.Instrument)
		if err := c.Bind(i); err != nil || i.ID == uuid.Nil {
			return c.NoContent(http.StatusBadRequest)
		}

		if err := models.CreateInstrumentGroupInstruments(db, instrumentGroupID, i.ID); err != nil {
			return c.JSON(http.StatusConflict, map[string]interface{}{
				"error":               "instrument already a member of this instrument group",
				"instrument_id":       i.ID,
				"instrument_group_id": instrumentGroupID,
			})
		}

		return c.NoContent(http.StatusCreated)
	}
}

// DeleteInstrumentGroupInstruments removes an instrument from an instrument group
func DeleteInstrumentGroupInstruments(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// instrument_group_id
		instrumentGroupID, err := uuid.Parse(c.Param("id"))
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

		return c.NoContent(http.StatusOK)
	}
}
