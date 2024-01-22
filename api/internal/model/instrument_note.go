package model

import (
	"context"
	"encoding/json"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/util"
	"github.com/google/uuid"
)

// InstrumentNote is a note about an instrument
type InstrumentNote struct {
	ID           uuid.UUID `json:"id"`
	InstrumentID uuid.UUID `json:"instrument_id" db:"instrument_id"`
	Title        string    `json:"title"`
	Body         string    `json:"body"`
	Time         time.Time `json:"time"`
	AuditInfo
}

// InstrumentNoteCollection is a collection of Instrument Notes
type InstrumentNoteCollection struct {
	Items []InstrumentNote
}

// UnmarshalJSON implements UnmarshalJSON interface
// Allows unpacking object or array of objects into array of objects
func (c *InstrumentNoteCollection) UnmarshalJSON(b []byte) error {
	switch util.JSONType(b) {
	case "ARRAY":
		if err := json.Unmarshal(b, &c.Items); err != nil {
			return err
		}
	case "OBJECT":
		var n InstrumentNote
		if err := json.Unmarshal(b, &n); err != nil {
			return err
		}
		c.Items = []InstrumentNote{n}
	default:
		c.Items = make([]InstrumentNote, 0)
	}
	return nil
}

const listInstrumentNotes = `
	SELECT
		N.id            AS id,
		N.instrument_id AS instrument_id,
		N.title,
		N.body,
		N.time,
		N.creator,
		N.create_date,
		N.updater,
		N.update_date
	FROM instrument_note N
`

// ListInstrumentNotes returns an array of instruments from the database
func (q *Queries) ListInstrumentNotes(ctx context.Context) ([]InstrumentNote, error) {
	nn := make([]InstrumentNote, 0)
	if err := q.db.SelectContext(ctx, &nn, listInstrumentNotes); err != nil {
		return nil, err
	}
	return nn, nil
}

const listInstrumentInstrumentNotes = listInstrumentNotes + `
	WHERE N.instrument_id = $1
`

// ListInstrumentInstrumentNotes returns an array of instrument notes for a given instrument
func (q *Queries) ListInstrumentInstrumentNotes(ctx context.Context, instrumentID uuid.UUID) ([]InstrumentNote, error) {
	nn := make([]InstrumentNote, 0)
	if err := q.db.SelectContext(ctx, &nn, listInstrumentInstrumentNotes, instrumentID); err != nil {
		return nil, err
	}
	return nn, nil
}

const getInstrumentNotes = listInstrumentNotes + `
	WHERE N.id = $1
`

// GetInstrumentNote returns a single instrument note
func (q *Queries) GetInstrumentNote(ctx context.Context, noteID uuid.UUID) (InstrumentNote, error) {
	var n InstrumentNote
	if err := q.db.GetContext(ctx, &n, getInstrumentNotes, noteID); err != nil {
		return n, err
	}
	return n, nil
}

const createInstrumentNote = `
	INSERT INTO instrument_note (instrument_id, title, body, time, creator, create_date)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING id, instrument_id, title, body, time, creator, create_date, updater, update_date
`

func (q *Queries) CreateInstrumentNote(ctx context.Context, note InstrumentNote) (InstrumentNote, error) {
	var noteNew InstrumentNote
	err := q.db.GetContext(ctx, &noteNew, createInstrumentNote, note.InstrumentID, note.Title, note.Body, note.Time, note.CreatorID, note.CreateDate)
	return noteNew, err
}

const updateInstrumentNote = `
	UPDATE instrument_note SET
		title = $2,
		body = $3,
		time = $4,
		updater = $5,
		update_date = $6
	WHERE id = $1
	RETURNING id, instrument_id, title, body, time, creator, create_date, updater, update_date
`

// UpdateInstrumentNote updates a single instrument note
func (q *Queries) UpdateInstrumentNote(ctx context.Context, n InstrumentNote) (InstrumentNote, error) {
	var nUpdated InstrumentNote
	err := q.db.GetContext(ctx, &nUpdated, updateInstrumentNote, n.ID, n.Title, n.Body, n.Time, n.UpdaterID, n.UpdateDate)
	return nUpdated, err
}

const deleteInstrumentNote = `
	DELETE FROM instrument_note WHERE id = $1
`

// DeleteInstrumentNote deletes an instrument note
func (q *Queries) DeleteInstrumentNote(ctx context.Context, noteID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteInstrumentNote, noteID)
	return err
}
