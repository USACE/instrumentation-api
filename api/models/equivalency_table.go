package models

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type EquivalencyTable struct {
	DataLoggerID uuid.UUID             `json:"datalogger_id" db:"datalogger_id"`
	Rows         []EquivalencyTableRow `json:"rows" db:"rows"`
}

type EquivalencyTableRow struct {
	ID           uuid.UUID  `json:"id" db:"id"`
	DataLoggerID uuid.UUID  `json:"-" db:"datalogger_id"`
	FieldName    string     `json:"field_name" db:"field_name"`
	DisplayName  string     `json:"display_name" db:"display_name"`
	InstrumentID *uuid.UUID `json:"instrument_id" db:"instrument_id"`
	TimeseriesID *uuid.UUID `json:"timeseries_id" db:"timeseries_id"`
}

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

func CreateEquivalencyTable(db *sqlx.DB, t *EquivalencyTable) error {
	txn, err := db.Beginx()
	if err != nil {
		return err
	}
	defer txn.Rollback()

	for _, r := range t.Rows {
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

func UpdateEquivalencyTable(db *sqlx.DB, t *EquivalencyTable) error {
	txn, err := db.Beginx()
	if err != nil {
		return err
	}
	defer txn.Rollback()

	for _, r := range t.Rows {
		stmt, err := txn.Preparex(`
			UPDATE datalogger_equivalency_table SET
				field_name = $2,
				display_name = $3,
				instrument_id = $4,
				timeseries_id = $5
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

func DeleteEquivalencyTable(db *sqlx.DB, dlID *uuid.UUID) error {
	_, err := db.Exec(`DELETE FROM datalogger_equivalency_table WHERE datalogger_id = $1`, &dlID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteEquivalencyTableRow(db *sqlx.DB, dlID *uuid.UUID, rID *uuid.UUID) error {
	if _, err := db.Exec(`
		DELETE FROM datalogger_equivalency_table WHERE datalogger_id = $1 AND id = $2
	`, &rID); err != nil {
		return err
	}
	return nil
}
