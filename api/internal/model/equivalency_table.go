package model

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
)

type EquivalencyTable struct {
	DataloggerID uuid.UUID             `json:"datalogger_id" db:"datalogger_id"`
	Rows         []EquivalencyTableRow `json:"rows" db:"rows"`
}

type EquivalencyTableRow struct {
	ID           uuid.UUID  `json:"id" db:"id"`
	DataloggerID uuid.UUID  `json:"-" db:"datalogger_id"`
	SN           string     `json:"-" db:"sn"`
	Model        string     `json:"-" db:"model"`
	FieldName    string     `json:"field_name" db:"field_name"`
	DisplayName  string     `json:"display_name" db:"display_name"`
	InstrumentID *uuid.UUID `json:"instrument_id" db:"instrument_id"`
	TimeseriesID *uuid.UUID `json:"timeseries_id" db:"timeseries_id"`
}

const getIsValidEquivalencyTableTimeseries = `
	SELECT NOT EXISTS (
		SELECT id FROM v_timeseries_computed
		WHERE id = $1
		UNION ALL
		SELECT timeseries_id FROM instrument_constants
		WHERE timeseries_id = $1
	)
`

// GetIsValidEquivalencyTableTimeseries verifies that a Timeseries is not computed or constant
func (q *Queries) GetIsValidEquivalencyTableTimeseries(ctx context.Context, tsID uuid.UUID) error {
	var isValid bool
	if err := q.db.GetContext(ctx, &isValid, getIsValidEquivalencyTableTimeseries, tsID); err != nil {
		return err
	}
	if !isValid {
		return fmt.Errorf("timeseries '%s' must not be computed", tsID)
	}
	return nil
}

const getEquivalencyTable = `
	SELECT
		id,
		datalogger_id,
		field_name,
		display_name,
		instrument_id,
		timeseries_id
	FROM datalogger_equivalency_table
	WHERE datalogger_id = $1
`

// GetEquivalencyTable returns a single Datalogger EquivalencyTable
func (q *Queries) GetEquivalencyTable(ctx context.Context, dlID uuid.UUID) (EquivalencyTable, error) {
	tr := make([]EquivalencyTableRow, 0)
	err := q.db.SelectContext(ctx, &tr, getEquivalencyTable, dlID)
	et := EquivalencyTable{
		DataloggerID: dlID,
		Rows:         tr,
	}
	return et, err
}

const createEquivalencyTableRow = `
	INSERT INTO datalogger_equivalency_table
	(datalogger_id, field_name, display_name, instrument_id, timeseries_id)
	VALUES ($1, $2, $3, $4, $5)
	ON CONFLICT ON CONSTRAINT unique_datalogger_field DO NOTHING
`

func (q *Queries) CreateEquivalencyTableRow(ctx context.Context, dlID uuid.UUID, tr EquivalencyTableRow) error {
	if _, err := q.db.ExecContext(ctx, createEquivalencyTableRow,
		dlID,
		tr.FieldName,
		tr.DisplayName,
		tr.InstrumentID,
		tr.TimeseriesID,
	); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			return fmt.Errorf("timeseries_id %s is already mapped to an active datalogger", tr.TimeseriesID)
		}
		return err
	}
	return nil
}

const updateEquivalencyTableRow = `
	UPDATE datalogger_equivalency_table SET
		field_name = $3,
		display_name = $4,
		instrument_id = $5,
		timeseries_id = $6
	WHERE datalogger_id = $1
	AND id = $2
`

func (q *Queries) UpdateEquivalencyTableRow(ctx context.Context, tr EquivalencyTableRow) error {
	if _, err := q.db.ExecContext(ctx, updateEquivalencyTableRow,
		tr.DataloggerID,
		tr.ID,
		tr.FieldName,
		tr.DisplayName,
		tr.InstrumentID,
		tr.TimeseriesID,
	); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			return fmt.Errorf("timeseries_id %s is already mapped to an active datalogger", tr.TimeseriesID)
		}
		return err
	}
	return nil
}

const deleteEquivalencyTable = `
	DELETE FROM datalogger_equivalency_table WHERE datalogger_id = $1
`

// DeleteEquivalencyTable clears all rows of the EquivalencyTable for a Datalogger
func (q *Queries) DeleteEquivalencyTable(ctx context.Context, dlID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteEquivalencyTable, dlID)
	return err
}

const deleteEquivalencyTableRow = `
	DELETE FROM datalogger_equivalency_table WHERE datalogger_id = $1 AND id = $2
`

// DeleteEquivalencyTableRow deletes a single EquivalencyTable row by row id
func (q *Queries) DeleteEquivalencyTableRow(ctx context.Context, dlID, rID uuid.UUID) error {
	res, err := q.db.ExecContext(ctx, deleteEquivalencyTableRow, dlID, rID)
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
