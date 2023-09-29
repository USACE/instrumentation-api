package handler

import (
	"net/http"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/messages"
	"github.com/USACE/instrumentation-api/api/internal/model"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// ListInstrumentNotes returns instrument notes
func (h ApiHandler) ListInstrumentNotes(c echo.Context) error {
	notes, err := model.ListInstrumentNotes(db)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, notes)
}

// ListInstrumentInstrumentNotes returns instrument notes for a single instrument
func (h ApiHandler) ListInstrumentInstrumentNotes(c echo.Context) error {
	iID, err := uuid.Parse(c.Param("instrument_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, messages.MalformedID)
	}
	notes, err := model.ListInstrumentInstrumentNotes(db, &iID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, notes)
}

// GetInstrumentNote returns a single instrument note
func (h ApiHandler) GetInstrumentNote(c echo.Context) error {
	nID, err := uuid.Parse(c.Param("note_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, messages.MalformedID)
	}
	note, err := model.GetInstrumentNote(db, &nID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, note)
}

// CreateInstrumentNote creates instrument notes
func (h ApiHandler) CreateInstrumentNote(c echo.Context) error {
	nc := model.InstrumentNoteCollection{}
	if err := c.Bind(&nc); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// profile and timestamp
	p := c.Get("profile").(*model.Profile)

	t := time.Now()
	for idx := range nc.Items {
		nc.Items[idx].Creator = p.ID
		nc.Items[idx].CreateDate = t
	}
	nn, err := model.CreateInstrumentNote(db, nc.Items)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, nn)
}

// UpdateInstrumentNote updates an instrument note
func (h ApiHandler) UpdateInstrumentNote(c echo.Context) error {
	noteID, err := uuid.Parse(c.Param("note_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	n := model.InstrumentNote{ID: noteID}
	if err := c.Bind(&n); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// check :id in url params matches id in request body
	if noteID != n.ID {
		return echo.NewHTTPError(http.StatusBadRequest, messages.MatchRouteParam("`note_id`"))
	}
	// profile and timestamp
	p := c.Get("profile").(*model.Profile)
	t := time.Now()
	n.Updater, n.UpdateDate = &p.ID, &t

	// update
	nUpdated, err := model.UpdateInstrumentNote(db, &n)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// return updated instrument note
	return c.JSON(http.StatusOK, nUpdated)
}

// DeleteInstrumentNote deletes an instrument note
func (h ApiHandler) DeleteInstrumentNote(c echo.Context) error {
	noteID, err := uuid.Parse(c.Param("note_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := model.DeleteInstrumentNote(db, &noteID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, make(map[string]interface{}))
}
