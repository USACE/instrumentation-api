package handlers

import (
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
			return c.String(http.StatusNotFound, "Malformed ID")
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

		id, err := models.CreateInstrumentGroup(db, g)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{"created": id})
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
		return c.JSON(http.StatusOK, map[string]interface{}{"deleted": id})
	}
}

// ListInstrumentGroupInstruments returns a list of instruments for a provided instrument group
func ListInstrumentGroupInstruments(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			return c.String(http.StatusNotFound, "Malformed ID")
		}
		return c.JSON(http.StatusOK, models.ListInstrumentGroupInstruments(db, id))
	}
}

// CreateInstrumentGroupInstruments adds an instrument to an instrument group
func CreateInstrumentGroupInstruments(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// instrument_group_id
		instrumentGroupID, err := uuid.Parse(c.Param("id"))
		if err != nil {
			return c.String(http.StatusNotFound, "Malformed ID")
		}
		// instrument
		i := new(models.Instrument)
		if err1 := c.Bind(i); err1 != nil {
			return c.JSON(http.StatusBadRequest, err1)
		}
		_, err2 := models.CreateInstrumentGroupInstruments(db, instrumentGroupID, i.ID)
		if err2 != nil {
			return c.JSON(http.StatusConflict, map[string]interface{}{
				"error":               "instrument already a member of this instrument group",
				"instrument_id":       i.ID,
				"instrument_group_id": instrumentGroupID,
			})
		}
		return c.JSON(http.StatusCreated, i.ID)
	}
}

// DeleteInstrumentGroupInstruments removes an instrument from an instrument group
func DeleteInstrumentGroupInstruments(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// instrument_group_id
		instrumentGroupID, err := uuid.Parse(c.Param("id"))
		if err != nil {
			return c.String(http.StatusNotFound, "Malformed ID")
		}

		// instrument
		instrumentID, err := uuid.Parse(c.Param("instrument_id"))
		if err != nil {
			return c.String(http.StatusNotFound, "Malformed ID")
		}

		if _, err := models.DeleteInstrumentGroupInstruments(db, instrumentGroupID, instrumentID); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"deleted":             instrumentID,
			"instrument_group_id": instrumentGroupID,
		})
	}
}
