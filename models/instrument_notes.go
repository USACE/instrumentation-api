package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
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

	switch JSONType(b) {
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

// ListInstrumentNotes returns an array of instruments from the database
func ListInstrumentNotes(db *sqlx.DB) ([]InstrumentNote, error) {

	nn := make([]InstrumentNote, 0)
	if err := db.Select(&nn, listInstrumentNotesSQL()); err != nil {
		return make([]InstrumentNote, 0), err
	}
	return nn, nil
}

// ListInstrumentInstrumentNotes returns an array of instrument notes for a given instrument
func ListInstrumentInstrumentNotes(db *sqlx.DB, instrumentID *uuid.UUID) ([]InstrumentNote, error) {

	nn := make([]InstrumentNote, 0)
	if err := db.Select(
		&nn,
		listInstrumentNotesSQL()+" WHERE N.instrument_id = $1",
		instrumentID,
	); err != nil {
		return make([]InstrumentNote, 0), err
	}
	return nn, nil
}

// GetInstrumentNote returns a single instrument note
func GetInstrumentNote(db *sqlx.DB, id *uuid.UUID) (*InstrumentNote, error) {

	var n InstrumentNote
	if err := db.Get(&n, listInstrumentNotesSQL()+" WHERE N.id = $1", id); err != nil {
		return nil, err
	}
	return &n, nil
}

// CreateInstrumentNote creates many instrument notes from an array of instrument notes
func CreateInstrumentNote(db *sqlx.DB, notes []InstrumentNote) ([]InstrumentNote, error) {

	txn, err := db.Beginx()
	if err != nil {
		return nil, err
	}

	stmt, err := txn.Preparex(
		`INSERT INTO instrument_note (instrument_id, title, body, time, creator, create_date)
		 VALUES ($1, $2, $3, $4, $5, $6)
		 RETURNING id, instrument_id, title, body, time, creator, create_date, updater, update_date`,
	)
	if err != nil {
		return make([]InstrumentNote, 0), err
	}

	nn := make([]InstrumentNote, len(notes))
	for idx, n := range notes {
		if err := stmt.Get(&nn[idx], n.InstrumentID, n.Title, n.Body, n.Time, n.Creator, n.CreateDate); err != nil {
			return make([]InstrumentNote, 0), err
		}
	}

	if err := stmt.Close(); err != nil {
		return make([]InstrumentNote, 0), err
	}

	if err := txn.Commit(); err != nil {
		return make([]InstrumentNote, 0), err
	}

	return nn, nil
}

// UpdateInstrumentNote updates a single instrument note
func UpdateInstrumentNote(db *sqlx.DB, n *InstrumentNote) (*InstrumentNote, error) {

	var nUpdated InstrumentNote
	if err := db.QueryRowx(
		`UPDATE instrument_note
		 SET    title = $2,
			    body = $3,
			    time = $4,
			    updater = $5,
				update_date = $6
		 WHERE id = $1
		 RETURNING id, instrument_id, title, body, time, creator, create_date, updater, update_date
		`, n.ID, n.Title, n.Body, n.Time, n.Updater, n.UpdateDate,
	).StructScan(&nUpdated); err != nil {
		return nil, err
	}
	return &nUpdated, nil

}

// DeleteInstrumentNote deletes an instrument note
func DeleteInstrumentNote(db *sqlx.DB, id *uuid.UUID) error {

	if _, err := db.Exec(`DELETE FROM instrument_note WHERE id = $1`, id); err != nil {
		return err
	}

	return nil
}

// ListInstrumentNotesSQL is the base SQL to retrieve all instrument notes
func listInstrumentNotesSQL() string {
	return `SELECT N.id            AS id,
				   N.instrument_id AS instrument_id,
				   N.title,
				   N.body,
				   N.time,
				   N.creator,
				   N.create_date,
				   N.updater,
				   N.update_date
			FROM   instrument_note N
			`
}
