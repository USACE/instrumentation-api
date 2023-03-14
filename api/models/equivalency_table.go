package models

import (
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jmoiron/sqlx"
)

type EquivalencyTable struct {
	DataLoggerID uuid.UUID             `json:"datalogger_id" db:"datalogger_id"`
	Rows         []EquivalencyTableRow `json:"rows" db:"rows"`
}

type EquivalencyTableRow struct {
	ID           uuid.UUID  `json:"id" db:"id"`
	DataLoggerID uuid.UUID  `json:"-" db:"datalogger_id"`
	SN           string     `json:"-" db:"sn"`
	Model        string     `json:"-" db:"model"`
	FieldName    string     `json:"field_name" db:"field_name"`
	DisplayName  string     `json:"display_name" db:"display_name"`
	InstrumentID *uuid.UUID `json:"instrument_id" db:"instrument_id"`
	TimeseriesID *uuid.UUID `json:"timeseries_id" db:"timeseries_id"`
}

// ValidEquivalencyTableTimeseries verifies that a Timeseries is not computed or constant
func ValidEquivalencyTableTimeseries(txn *sqlx.Tx, tsID *uuid.UUID) error {
	stmt, err := txn.Preparex(`
		SELECT NOT EXISTS (
			SELECT id FROM v_timeseries_computed
			WHERE id = $1
			AND id NOT IN (
				SELECT timeseries_id FROM instrument_constants
			)
		)
	`)
	if err != nil {
		return err
	}

	var isValid bool
	if err := stmt.Get(&isValid, &tsID); err != nil {
		return err
	}
	if !isValid {
		return fmt.Errorf("timeseries '%s' must not be computed", tsID)
	}
	if err := stmt.Close(); err != nil {
		return err
	}

	return nil
}

// GetEquivalencyTable returns a single DataLogger EquivalencyTable
func GetEquivalencyTable(db *sqlx.DB, dlID *uuid.UUID) (*EquivalencyTable, error) {
	var tr []EquivalencyTableRow
	if err := db.Select(&tr, `
		SELECT
			id,
			datalogger_id,
			field_name,
			display_name,
			instrument_id,
			timeseries_id
		FROM datalogger_equivalency_table
		WHERE datalogger_id = $1
	`, &dlID); err != nil {
		return nil, err
	}

	if tr == nil {
		tr = make([]EquivalencyTableRow, 0)
	}

	return &EquivalencyTable{
		DataLoggerID: *dlID,
		Rows:         tr,
	}, nil
}

// CreateEquivalencyTable creates EquivalencyTable rows
// If a row with the given datalogger id or field name already exists the row will be ignored
func CreateEquivalencyTable(db *sqlx.DB, t *EquivalencyTable) error {
	txn, err := db.Beginx()
	if err != nil {
		return err
	}
	defer txn.Rollback()

	for _, r := range t.Rows {
		if err = ValidEquivalencyTableTimeseries(txn, r.TimeseriesID); err != nil {
			return err
		}

		stmt, err := txn.Preparex(`
			INSERT INTO datalogger_equivalency_table
			(datalogger_id, field_name, display_name, instrument_id, timeseries_id)
			VALUES ($1, $2, $3, $4, $5)
			ON CONFLICT ON CONSTRAINT unique_datalogger_field DO NOTHING
		`)
		if err != nil {
			return err
		}
		if _, err := stmt.Exec(
			&t.DataLoggerID,
			&r.FieldName,
			&r.DisplayName,
			&r.InstrumentID,
			&r.TimeseriesID,
		); err != nil {
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
				// TODO: possibly query and return the active eqivalency table row mapped to this timeseries

				return fmt.Errorf("timeseries_id %s is already mapped to an active datalogger", r.TimeseriesID)
			}
			log.Printf("error is %s", err)
			return err
		}
		if err := stmt.Close(); err != nil {
			return err
		}
	}
	if err := txn.Commit(); err != nil {
		return err
	}

	return nil
}

// UpdateEquivalencyTable updates rows of an EquivalencyTable
func UpdateEquivalencyTable(db *sqlx.DB, t *EquivalencyTable) error {
	txn, err := db.Beginx()
	if err != nil {
		return err
	}
	defer txn.Rollback()

	for _, r := range t.Rows {
		if err = ValidEquivalencyTableTimeseries(txn, r.TimeseriesID); err != nil {
			return err
		}

		stmt, err := txn.Preparex(`
			UPDATE datalogger_equivalency_table SET
				field_name = $3,
				display_name = $4,
				instrument_id = $5,
				timeseries_id = $6
			WHERE datalogger_id = $1
			AND id = $2
		`)
		if err != nil {
			return err
		}
		if _, err := stmt.Exec(
			&t.DataLoggerID,
			&r.ID,
			&r.FieldName,
			&r.DisplayName,
			&r.InstrumentID,
			&r.TimeseriesID,
		); err != nil {
			return err
		}
		if err := stmt.Close(); err != nil {
			return err
		}
	}
	if err := txn.Commit(); err != nil {
		return err
	}

	return nil
}

// DeleteEquivalencyTable clears all rows of the EquivalencyTable for a Datalogger
func DeleteEquivalencyTable(db *sqlx.DB, dlID *uuid.UUID) error {
	_, err := db.Exec(`DELETE FROM datalogger_equivalency_table WHERE datalogger_id = $1`, &dlID)
	if err != nil {
		return err
	}
	return nil
}

// DeleteEquivalencyTableRow deletes a single EquivalencyTable row by row id
func DeleteEquivalencyTableRow(db *sqlx.DB, dlID *uuid.UUID, rID *uuid.UUID) error {
	res, err := db.Exec(`
		DELETE FROM datalogger_equivalency_table WHERE datalogger_id = $1 AND id = $2
	`, &dlID, &rID)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("row %s not found for datalogger %s", rID, dlID)
	}
	return nil
}
