package models

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

// Timeseries is a timeseries data structure
type Timeseries struct {
	ID           uuid.UUID               `json:"id"`
	Slug         string                  `json:"slug"`
	Name         string                  `json:"name"`
	InstrumentID uuid.UUID               `json:"instrument_id" db:"instrument_id"`
	Instrument   string                  `json:"instrument,omitempty"`
	ParameterID  uuid.UUID               `json:"parameter_id" db:"parameter_id"`
	Parameter    string                  `json:"parameter,omitempty"`
	UnitID       uuid.UUID               `json:"unit_id" db:"unit_id"`
	Unit         string                  `json:"unit,omitempty"`
	Values       []TimeseriesMeasurement `json:"values,omitempty"`
}

// TimeseriesCollection is a collection of Timeseries items
type TimeseriesCollection struct {
	Items []Timeseries
}

// UnmarshalJSON implements UnmarshalJSON interface
func (c *TimeseriesCollection) UnmarshalJSON(b []byte) error {
	switch JSONType(b) {
	case "ARRAY":
		if err := json.Unmarshal(b, &c.Items); err != nil {
			return err
		}
	case "OBJECT":
		var t Timeseries
		if err := json.Unmarshal(b, &t); err != nil {
			return err
		}
		c.Items = []Timeseries{t}
	default:
		c.Items = make([]Timeseries, 0)
	}
	return nil
}

// ListTimeseriesSlugs lists used timeseries slugs in the database
func ListTimeseriesSlugs(db *sqlx.DB) ([]string, error) {

	ss := make([]string, 0)
	if err := db.Select(&ss, "SELECT slug FROM timeseries"); err != nil {
		return make([]string, 0), err
	}
	return ss, nil
}

// ListTimeseries lists all timeseries
func ListTimeseries(db *sqlx.DB) ([]Timeseries, error) {

	tt := make([]Timeseries, 0)
	if err := db.Select(&tt, listTimeseriesSQL()); err != nil {
		return make([]Timeseries, 0), err
	}
	return tt, nil
}

// ListInstrumentTimeseries returns an array of timeseries for an instrument
func ListInstrumentTimeseries(db *sqlx.DB, instrumentID *uuid.UUID) ([]Timeseries, error) {
	tt := make([]Timeseries, 0)
	if err := db.Select(&tt, listTimeseriesSQL()+" WHERE I.id = $1", instrumentID); err != nil {
		return make([]Timeseries, 0), err
	}
	return tt, nil
}

// ListInstrumentGroupTimeseries returns an array of timeseries for instruments that belong to an instrument_group
func ListInstrumentGroupTimeseries(db *sqlx.DB, instrumentGroupID *uuid.UUID) ([]Timeseries, error) {

	var tt []Timeseries
	if err := db.Select(
		&tt,
		`SELECT *
		 FROM   timeseries
		 WHERE  instrument_id IN (
			SELECT instrument_id
			FROM   instrument_group_instruments
			WHERE  instrument_group_id = $1
		)`, instrumentGroupID,
	); err != nil {
		return make([]Timeseries, 0), err
	}
	return tt, nil
}

// GetTimeseries returns a single timeseries without measurements
func GetTimeseries(db *sqlx.DB, id *uuid.UUID) (*Timeseries, error) {

	var t Timeseries
	if err := db.Get(&t, listTimeseriesSQL()+" WHERE T.id = $1", id); err != nil {
		return nil, err
	}
	return &t, nil
}

// CreateTimeseries creates many timeseries from an array of timeseries
func CreateTimeseries(db *sqlx.DB, tt []Timeseries) error {

	txn, err := db.Begin()
	if err != nil {
		return err
	}

	// Create Timeseries
	stmt, err := txn.Prepare(pq.CopyIn("timeseries", "slug", "name", "instrument_id", "parameter_id", "unit_id"))
	if err != nil {
		return err
	}

	// Iterate Timeseries
	for _, t := range tt {
		if _, err := stmt.Exec(t.Slug, t.Name, t.InstrumentID, t.ParameterID, t.UnitID); err != nil {
			return err
		}
	}
	if _, err := stmt.Exec(); err != nil {
		return err
	}
	if err := stmt.Close(); err != nil {
		return err
	}
	if err := txn.Commit(); err != nil {
		return err
	}

	return nil
}

// UpdateTimeseries updates a timeseries
func UpdateTimeseries(db *sqlx.DB, t *Timeseries) (*Timeseries, error) {

	var tUpdated Timeseries
	if err := db.QueryRowx(
		`UPDATE timeseries
		 SET    name = $2,
			    instrument_id = $3,
			    parameter_id = $4,
			    unit_id = $5
		 WHERE id = $1
		 RETURNING *
		`, t.ID, t.Name, t.InstrumentID, t.ParameterID, t.UnitID,
	).StructScan(&tUpdated); err != nil {
		return nil, err
	}

	return &tUpdated, nil
}

// DeleteTimeseries deletes a timeseries and cascade deletes all measurements
func DeleteTimeseries(db *sqlx.DB, id *uuid.UUID) error {
	if _, err := db.Exec("DELETE from timeseries WHERE id = $1", id); err != nil {
		return err
	}
	return nil
}

func listTimeseriesSQL() string {
	return `SELECT T.id as id,
	               T.SLUG as slug, 
				   T.NAME as name,
				   I.ID   as instrument_id,
				   I.Name as instrument,
				   P.ID   as parameter_id,
				   P.Name as parameter,
				   U.ID   as unit_id,
	               U.Name as unit
            FROM timeseries T
	        INNER JOIN instrument I
				    ON I.id = T.instrument_id
	        INNER JOIN parameter P
				    ON P.id = T.parameter_id
	        INNER JOIN unit U
				    ON U.id = T.unit_id
	`
}
