package handler

import (
	"net/http"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/messages"
	"github.com/USACE/instrumentation-api/api/internal/models"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

// ListInstrumentNotes returns instrument notes
func ListInstrumentNotes(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		notes, err := models.ListInstrumentNotes(db)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, notes)
	}
}

// ListInstrumentInstrumentNotes returns instrument notes for a single instrument
func ListInstrumentInstrumentNotes(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		iID, err := uuid.Parse(c.Param("instrument_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, messages.MalformedID)
		}
		notes, err := models.ListInstrumentInstrumentNotes(db, &iID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, notes)
	}
}

// GetInstrumentNote returns a single instrument note
func GetInstrumentNote(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		nID, err := uuid.Parse(c.Param("note_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, messages.MalformedID)
		}
		note, err := models.GetInstrumentNote(db, &nID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, note)
	}
}

// CreateInstrumentNote creates instrument notes
func CreateInstrumentNote(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		nc := models.InstrumentNoteCollection{}
		if err := c.Bind(&nc); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		// profile and timestamp
		p := c.Get("profile").(*models.Profile)

		t := time.Now()
		for idx := range nc.Items {
			nc.Items[idx].Creator = p.ID
			nc.Items[idx].CreateDate = t
		}
		nn, err := models.CreateInstrumentNote(db, nc.Items)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusCreated, nn)
	}
}

// UpdateInstrumentNote updates an instrument note
func UpdateInstrumentNote(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		noteID, err := uuid.Parse(c.Param("note_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		n := models.InstrumentNote{ID: noteID}
		if err := c.Bind(&n); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		// check :id in url params matches id in request body
		if noteID != n.ID {
			return echo.NewHTTPError(http.StatusBadRequest, messages.MatchRouteParam("`note_id`"))
		}
		// profile and timestamp
		p := c.Get("profile").(*models.Profile)
		t := time.Now()
		n.Updater, n.UpdateDate = &p.ID, &t

		// update
		nUpdated, err := models.UpdateInstrumentNote(db, &n)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		// return updated instrument note
		return c.JSON(http.StatusOK, nUpdated)
	}
}

// DeleteInstrumentNote deletes an instrument note
func DeleteInstrumentNote(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		noteID, err := uuid.Parse(c.Param("note_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if err := models.DeleteInstrumentNote(db, &noteID); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, make(map[string]interface{}))
	}
}
