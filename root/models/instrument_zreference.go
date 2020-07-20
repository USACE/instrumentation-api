package models

import (
	"api/root/timeseries"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	// pq library
	_ "github.com/lib/pq"
)

// ZReference is a vertical reference associated with a time
// The time represents when the vertical reference went into effect
type ZReference struct {
	ID                uuid.UUID `json:"id"`
	Time              time.Time `json:"time"`
	ZReference        float32   `json:"zreference"`
	ZReferenceDatumID uuid.UUID `json:"zreference_datum_id" db:"zreference_datum_id"`
	ZReferenceDatum   string    `json:"zreference_datum" db:"zreference_datum"`
}

// AsTimeseriesMeasurement returns a timeseries.Measurement representation of ZReference
func (z ZReference) AsTimeseriesMeasurement() timeseries.Measurement {
	return timeseries.Measurement{
		Time:  z.Time,
		Value: z.ZReference,
	}
}

// ZReferenceCollection is a collection of ZReference items
type ZReferenceCollection struct {
	Items []ZReference
}

// UnmarshalJSON implements UnmarshalJSON interface
func (c *ZReferenceCollection) UnmarshalJSON(b []byte) error {

	switch JSONType(b) {
	case "ARRAY":
		if err := json.Unmarshal(b, &c.Items); err != nil {
			return err
		}
	case "OBJECT":
		var z ZReference
		if err := json.Unmarshal(b, &z); err != nil {
			return err
		}
		c.Items = []ZReference{z}
	default:
		c.Items = make([]ZReference, 0)
	}
	return nil
}

// CreateInstrumentZReferenceSQL is the base SQL to create instrument ZReference values, update on conflict
func createInstrumentZReferenceSQL() string {
	return `INSERT INTO instrument_zreference (instrument_id, time, zreference, zreference_datum_id) VALUES ($1, $2, $3, $4)
	ON CONFLICT ON CONSTRAINT instrument_unique_zreference_in_time DO UPDATE SET zreference = EXCLUDED.zreference, zreference_datum_id = EXCLUDED.zreference_datum_id;`
}

// ListInstrumentZReferenceSQL the base SQL to retrieve all InstrumentZReference
func listInstrumentZReferenceSQL() string {
	return `SELECT Z.id,
				   Z.time,
				   Z.zreference,
				   Z.zreference_datum_id,
				   D.name AS zreference_datum
			FROM   instrument_zreference Z
			INNER JOIN zreference_datum D
				ON D.id = Z.zreference_datum_id
	`
}

// ListInstrumentZReference returns all ZReference values for an instrument
func ListInstrumentZReference(db *sqlx.DB, instrumentID *uuid.UUID) ([]ZReference, error) {

	zz := make([]ZReference, 0)
	if err := db.Select(&zz, listInstrumentZReferenceSQL()+" WHERE Z.instrument_id = $1 ORDER BY Z.time DESC", instrumentID); err != nil {
		return make([]ZReference, 0), err
	}
	return zz, nil
}

// GetInstrumentZReference gets a single ZReference value
func GetInstrumentZReference(db *sqlx.DB, id *uuid.UUID) (*ZReference, error) {

	var z ZReference
	if err := db.Get(&z, listInstrumentZReferenceSQL()+" WHERE Z.id = $1", id); err != nil {
		return nil, err
	}
	return &z, nil
}

// CreateOrUpdateInstrumentZReference creates a ZReference, updates value on conflict
func CreateOrUpdateInstrumentZReference(db *sqlx.DB, instrumentID *uuid.UUID, zz []ZReference) error {

	txn, err := db.Begin()
	if err != nil {
		return err
	}
	stmt, err := txn.Prepare(createInstrumentZReferenceSQL())
	for _, z := range zz {
		if _, err := stmt.Exec(instrumentID, z.Time, z.ZReference, z.ZReferenceDatumID); err != nil {
			return err
		}
	}
	if err := stmt.Close(); err != nil {
		return err
	}
	if err := txn.Commit(); err != nil {
		return err
	}
	return nil
}

// DeleteInstrumentZReference deletes a ZReference for an instrument
func DeleteInstrumentZReference(db *sqlx.DB, id *uuid.UUID) error {
	if _, err := db.Exec(`DELETE FROM instrument_zreference WHERE id = $1`, id); err != nil {
		return err
	}
	return nil
}
