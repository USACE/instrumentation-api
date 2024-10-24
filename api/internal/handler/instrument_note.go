package handler

import (
	"net/http"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/httperr"
	"github.com/USACE/instrumentation-api/api/internal/model"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// ListInstrumentNotes godoc
//
//	@Summary gets all instrument notes
//	@Tags instrument-note
//	@Produce json
//	@Success 200 {array} model.InstrumentNote
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /instruments/notes [get]
func (h *ApiHandler) ListInstrumentNotes(c echo.Context) error {
	notes, err := h.InstrumentNoteService.ListInstrumentNotes(c.Request().Context())
	if err != nil {
		return httperr.InternalServerError(err)
	}
	return c.JSON(http.StatusOK, notes)
}

// ListInstrumentInstrumentNotes godoc
//
//	@Summary gets instrument notes for a single instrument
//	@Tags instrument-note
//	@Produce json
//	@Param instrument_id path string true "instrument uuid" Format(uuid)
//	@Success 200 {array} model.InstrumentNote
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /instruments/{instrument_id}/notes [get]
func (h *ApiHandler) ListInstrumentInstrumentNotes(c echo.Context) error {
	iID, err := uuid.Parse(c.Param("instrument_id"))
	if err != nil {
		return httperr.MalformedID(err)
	}
	notes, err := h.InstrumentNoteService.ListInstrumentInstrumentNotes(c.Request().Context(), iID)
	if err != nil {
		return httperr.InternalServerError(err)
	}
	return c.JSON(http.StatusOK, notes)
}

// GetInstrumentNote godoc
//
//	@Summary gets a single instrument note by id
//	@Tags instrument-note
//	@Produce json
//	@Param note_id path string true "note uuid" Format(uuid)
//	@Success 200 {object} model.InstrumentNote
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /instruments/notes/{note_id} [get]
func (h *ApiHandler) GetInstrumentNote(c echo.Context) error {
	nID, err := uuid.Parse(c.Param("note_id"))
	if err != nil {
		return httperr.MalformedID(err)
	}
	note, err := h.InstrumentNoteService.GetInstrumentNote(c.Request().Context(), nID)
	if err != nil {
		return httperr.InternalServerError(err)
	}
	return c.JSON(http.StatusOK, note)
}

// CreateInstrumentNote godoc
//
//	@Summary creates instrument notes
//	@Tags instrument-note
//	@Produce json
//	@Param instrument_note body model.InstrumentNoteCollection true "instrument note collection payload"
//	@Param key query string false "api key"
//	@Success 200 {array} model.InstrumentNote
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /instruments/notes [post]
//	@Security Bearer
func (h *ApiHandler) CreateInstrumentNote(c echo.Context) error {
	nc := model.InstrumentNoteCollection{}
	if err := c.Bind(&nc); err != nil {
		return httperr.MalformedBody(err)
	}
	// profile and timestamp
	p := c.Get("profile").(model.Profile)

	t := time.Now()
	for idx := range nc.Items {
		nc.Items[idx].CreatorID = p.ID
		nc.Items[idx].CreateDate = t
	}
	nn, err := h.InstrumentNoteService.CreateInstrumentNote(c.Request().Context(), nc.Items)
	if err != nil {
		return httperr.InternalServerError(err)
	}

	return c.JSON(http.StatusCreated, nn)
}

// UpdateInstrumentNote godoc
//
//	@Summary updates an instrument note by id
//	@Tags instrument-note
//	@Produce json
//	@Param note_id path string true "note uuid" Format(uuid)
//	@Param instrument_note body model.InstrumentNote true "instrument note collection payload"
//	@Param key query string false "api key"
//	@Success 200 {array} model.AlertConfig
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /instruments/notes/{note_id} [put]
//	@Security Bearer
func (h *ApiHandler) UpdateInstrumentNote(c echo.Context) error {
	noteID, err := uuid.Parse(c.Param("note_id"))
	if err != nil {
		return httperr.MalformedID(err)
	}
	n := model.InstrumentNote{ID: noteID}
	if err := c.Bind(&n); err != nil {
		return httperr.MalformedBody(err)
	}
	n.ID = noteID

	p := c.Get("profile").(model.Profile)
	t := time.Now()
	n.UpdaterID, n.UpdateDate = &p.ID, &t

	nUpdated, err := h.InstrumentNoteService.UpdateInstrumentNote(c.Request().Context(), n)
	if err != nil {
		return httperr.InternalServerError(err)
	}
	return c.JSON(http.StatusOK, nUpdated)
}

// DeleteInstrumentNote godoc
//
//	@Summary deletes an instrument note
//	@Tags instrument-note
//	@Produce json
//	@Param instrument_id path string true "instrument uuid" Format(uuid)
//	@Param note_id path string true "note uuid" Format(uuid)
//	@Param key query string false "api key"
//	@Success 200 {object} map[string]interface{}
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /instruments/{instrument_id}/notes/{note_id} [delete]
//	@Security Bearer
func (h *ApiHandler) DeleteInstrumentNote(c echo.Context) error {
	noteID, err := uuid.Parse(c.Param("note_id"))
	if err != nil {
		return httperr.MalformedID(err)
	}
	if err := h.InstrumentNoteService.DeleteInstrumentNote(c.Request().Context(), noteID); err != nil {
		return httperr.InternalServerError(err)
	}
	return c.JSON(http.StatusOK, make(map[string]interface{}))
}
