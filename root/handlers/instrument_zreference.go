package handlers

import (
	"api/root/models"
	"net/http"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
)

// ListInstrumentZReference lists all ZReference for an instrument
func ListInstrumentZReference(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := uuid.Parse(c.Param("instrument_id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Malformed ID")
		}

		zz, err := models.ListInstrumentZReference(db, &id)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, zz)
	}
}

// GetInstrumentZReference returns a single zreference
func GetInstrumentZReference(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := uuid.Parse(c.Param("zreference_id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Malformed ID")
		}

		z, err := models.GetInstrumentZReference(db, &id)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, z)
	}
}

// CreateOrUpdateInstrumentZReference creates a ZReference for an instrument
func CreateOrUpdateInstrumentZReference(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		instrumentID, err := uuid.Parse(c.Param("instrument_id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Malformed ID")
		}

		var zc models.ZReferenceCollection
		if err := c.Bind(&zc); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		// Assign Fresh UUID to each ZReference
		for idx := range zc.Items {
			id, err := uuid.NewRandom()
			if err != nil {
				return c.JSON(http.StatusInternalServerError, err)
			}
			zc.Items[idx].ID = id
		}

		if err := models.CreateOrUpdateInstrumentZReference(db, &instrumentID, zc.Items); err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.NoContent(http.StatusCreated)
	}
}

// DeleteInstrumentZReference deletes a ZReference for an instrument
func DeleteInstrumentZReference(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := uuid.Parse(c.Param("zreference_id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Malformed ID")
		}
		if err := models.DeleteInstrumentZReference(db, &id); err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.NoContent(http.StatusOK)
	}
}
