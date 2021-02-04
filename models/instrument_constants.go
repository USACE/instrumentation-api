package models

import (
	ts "github.com/USACE/instrumentation-api/timeseries"
	"github.com/google/uuid"

	"github.com/jmoiron/sqlx"
)

// ListInstrumentConstants lists constants for a given instrument id
func ListInstrumentConstants(db *sqlx.DB, id *uuid.UUID) ([]ts.Timeseries, error) {
	// ListInstrumentTimeseries returns an array of timeseries for an instrument
	tt := make([]ts.Timeseries, 0)
	if err := db.Select(&tt,
		`SELECT * FROM v_timeseries
		 WHERE instrument_id = $1 AND id IN (
			SELECT timeseries_id
			FROM instrument_constants
			WHERE instrument_id = $1
		)`, id,
	); err != nil {
		return make([]ts.Timeseries, 0), err
	}
	return tt, nil
}

// CreateInstrumentConstants creates many instrument constants from an array of instrument constants
// An InstrumentConstant is structurally the same as a timeseries and saved in the same tables
func CreateInstrumentConstants(db *sqlx.DB, tt []ts.Timeseries) ([]ts.Timeseries, error) {
	txn, err := db.Beginx()
	if err != nil {
		return nil, err
	}
	// Create Timeseries
	stmt1, err := txn.Preparex(
		`INSERT INTO timeseries (instrument_id, slug, name, parameter_id, unit_id)
		 VALUES ($1, $2, $3, $4, $5)
		 RETURNING id, instrument_id, slug, name, parameter_id, unit_id`,
	)
	if err != nil {
		return make([]ts.Timeseries, 0), err
	}
	// Designate Timeseries as Instrument Constant
	stmt2, err := txn.Preparex(
		`INSERT INTO instrument_constants (instrument_id, timeseries_id) VALUES ($1, $2)`,
	)
	if err != nil {
		return make([]ts.Timeseries, 0), err
	}
	// Create timeseries for each item in array and designate each as instrument_constant
	uu := make([]ts.Timeseries, len(tt))
	for idx, t := range tt {
		// create timeseries
		if err := stmt1.Get(&uu[idx], t.InstrumentID, t.Slug, t.Name, t.ParameterID, t.UnitID); err != nil {
			return make([]ts.Timeseries, 0), err
		}
		// designate as instrument_constant
		if _, err := stmt2.Exec(&uu[idx].InstrumentID, &uu[idx].ID); err != nil {
			return make([]ts.Timeseries, 0), err
		}
	}
	if err := stmt1.Close(); err != nil {
		return make([]ts.Timeseries, 0), err
	}
	if err := stmt2.Close(); err != nil {
		return make([]ts.Timeseries, 0), err
	}
	if err := txn.Commit(); err != nil {
		return make([]ts.Timeseries, 0), err
	}
	return uu, nil
}

// DeleteInstrumentConstant removes a timeseries as an Instrument Constant; Does not delete underlying timeseries
func DeleteInstrumentConstant(db *sqlx.DB, instrumentID *uuid.UUID, timeseriesID *uuid.UUID) error {
	txn, err := db.Beginx()
	if err != nil {
		return err
	}
	// Create Timeseries
	stmt1, err := txn.Preparex(
		`DELETE FROM instrument_constants WHERE instrument_id = $1 AND timeseries_id = $2`,
	)
	if err != nil {
		return err
	}
	// Designate Timeseries as Instrument Constant
	stmt2, err := txn.Preparex(
		`DELETE FROM timeseries WHERE instrument_id = $1 AND id = $2`,
	)
	if err != nil {
		return err
	}
	if _, err := stmt1.Exec(instrumentID, timeseriesID); err != nil {
		return err
	}
	if _, err := stmt2.Exec(instrumentID, timeseriesID); err != nil {
		return err
	}
	if err := stmt1.Close(); err != nil {
		return err
	}
	if err := stmt2.Close(); err != nil {
		return err
	}
	if err := txn.Commit(); err != nil {
		return err
	}
	return nil
}
