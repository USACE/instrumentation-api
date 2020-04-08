package handlers

import (
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
			return c.String(http.StatusNotFound, "Malformed ID")
		}
		return c.JSON(http.StatusOK, models.GetInstrument(db, id))
	}
}

// CreateInstrument creates a single instrument
func CreateInstrument(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := uuid.Must(uuid.NewRandom())

		i := &models.Instrument{ID: id}
		if err := c.Bind(i); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		if err := models.CreateInstrument(db, i); err != nil {
			return c.String(http.StatusForbidden, err.Error())
		}

		return c.JSON(http.StatusCreated, i.ID)
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
			return c.String(http.StatusBadRequest, err.Error())
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

		return c.JSON(http.StatusOK, i.ID)
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
			return c.String(http.StatusBadRequest, "Bad Request")
		}

		return c.JSON(http.StatusOK, map[string]interface{}{"deleted": id})
	}
}
