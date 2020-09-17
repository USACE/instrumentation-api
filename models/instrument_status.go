package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	// pq library
	_ "github.com/lib/pq"
)

// InstrumentStatus is an instrument status
type InstrumentStatus struct {
	ID       uuid.UUID `json:"id"`
	Time     time.Time `json:"time"`
	StatusID uuid.UUID `json:"status_id" db:"status_id"`
	Status   string    `json:"status"`
}

// InstrumentStatusCollection is a collection of instrument status
type InstrumentStatusCollection struct {
	Items []InstrumentStatus
}

// UnmarshalJSON implements the UnmarshalJSONinterface
func (c *InstrumentStatusCollection) UnmarshalJSON(b []byte) error {

	switch JSONType(b) {
	case "ARRAY":
		if err := json.Unmarshal(b, &c.Items); err != nil {
			return err
		}
	case "OBJECT":
		var s InstrumentStatus
		if err := json.Unmarshal(b, &s); err != nil {
			return err
		}
		c.Items = []InstrumentStatus{s}
	default:
		c.Items = make([]InstrumentStatus, 0)
	}
	return nil
}

// CreateInstrumentStatusSQL is the base SQL to create instrument status values, update on conflict
func createInstrumentStatusSQL() string {
	return `INSERT INTO instrument_status (instrument_id, status_id, time) VALUES ($1, $2, $3)
	        ON CONFLICT ON CONSTRAINT instrument_unique_status_in_time DO UPDATE SET status_id = EXCLUDED.status_id;`
}

// ListInstrumentStatusSQL the base SQL to retrieve all status for all instruments
func listInstrumentStatusSQL() string {
	return `SELECT S.id,
				   S.status_id,
				   D.name         AS status,
				   S.time
			FROM instrument_status S
			INNER JOIN status D
				ON D.id = S.status_id
	`
}

// ListInstrumentStatus returns all status values for an instrument
func ListInstrumentStatus(db *sqlx.DB, id *uuid.UUID) ([]InstrumentStatus, error) {

	ss := make([]InstrumentStatus, 0)
	if err := db.Select(&ss, listInstrumentStatusSQL()+" WHERE S.instrument_id = $1 ORDER BY time DESC", id); err != nil {
		return make([]InstrumentStatus, 0), err
	}
	return ss, nil
}

// GetInstrumentStatus gets a single status
func GetInstrumentStatus(db *sqlx.DB, id *uuid.UUID) (*InstrumentStatus, error) {

	var s InstrumentStatus
	if err := db.Get(&s, listInstrumentStatusSQL()+" WHERE S.id = $1", id); err != nil {
		return nil, err
	}
	return &s, nil
}

// CreateOrUpdateInstrumentStatus creates a Instrument Status, updates value on conflict
func CreateOrUpdateInstrumentStatus(db *sqlx.DB, instrumentID *uuid.UUID, ss []InstrumentStatus) error {

	txn, err := db.Begin()
	if err != nil {
		return err
	}
	stmt2, err := txn.Prepare(createInstrumentStatusSQL())
	for _, s := range ss {
		if _, err := stmt2.Exec(instrumentID, s.StatusID, s.Time); err != nil {
			return err
		}
	}
	if err := stmt2.Close(); err != nil {
		return err
	}
	if err := txn.Commit(); err != nil {
		return err
	}
	return nil
}

// DeleteInstrumentStatus deletes a status for an instrument
func DeleteInstrumentStatus(db *sqlx.DB, id *uuid.UUID) error {
	if _, err := db.Exec(`DELETE FROM instrument_status WHERE id = $1`, id); err != nil {
		return err
	}
	return nil
}
