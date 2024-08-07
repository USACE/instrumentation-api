package model

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
)

type DataloggerEquivalencyTable struct {
	DataloggerID        uuid.UUID                                  `json:"datalogger_id" db:"datalogger_id"`
	DataloggerTableID   uuid.UUID                                  `json:"datalogger_table_id" db:"datalogger_table_id"`
	DataloggerTableName string                                     `json:"datalogger_table_name" db:"datalogger_table_name"`
	Rows                dbJSONSlice[DataloggerEquivalencyTableRow] `json:"rows" db:"fields"`
}

type DataloggerEquivalencyTableRow struct {
	ID           uuid.UUID  `json:"id" db:"id"`
	FieldName    string     `json:"field_name" db:"field_name"`
	DisplayName  string     `json:"display_name" db:"display_name"`
	InstrumentID *uuid.UUID `json:"instrument_id" db:"instrument_id"`
	TimeseriesID *uuid.UUID `json:"timeseries_id" db:"timeseries_id"`
}

const getIsValidDataloggerTable = `
	SELECT NOT EXISTS (
		SELECT * FROM datalogger_table WHERE id = $1 AND table_name = 'preparse'
	)
`

// GetIsValidDataloggerTable verifies that a datalogger table is not "preparse" (read-only)
func (q *Queries) GetIsValidDataloggerTable(ctx context.Context, dataloggerTableID uuid.UUID) error {
	var isValid bool
	if err := q.db.GetContext(ctx, &isValid, getIsValidDataloggerTable, dataloggerTableID); err != nil {
		return err
	}
	if !isValid {
		return fmt.Errorf("table preparse is read only %s", dataloggerTableID)
	}
	return nil
}

const getIsValidEquivalencyTableTimeseriesBatch = `
	SELECT NOT EXISTS (
		SELECT id FROM v_timeseries_computed
		WHERE id IN (?)
		UNION ALL
		SELECT timeseries_id FROM instrument_constants
		WHERE timeseries_id IN (?)
	)
`

func (q *Queries) GetIsValidEquivalencyTableTimeseriesBatch(ctx context.Context, timeseriesIDs []uuid.UUID) error {
	if len(timeseriesIDs) == 0 {
		return nil
	}

	query, args, err := sqlIn(getIsValidEquivalencyTableTimeseriesBatch, timeseriesIDs, timeseriesIDs)
	if err != nil {
		return err
	}
	query = q.db.Rebind(query)

	var isValid bool
	if err := q.db.GetContext(ctx, &isValid, query, args...); err != nil {
		return err
	}
	if !isValid {
		return errors.New("comuted or constant timeseries not allowed")
	}
	return nil
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
		return fmt.Errorf("timeseries '%s' must not be computed or constant", tsID)
	}
	return nil
}

const getEquivalencyTable = `
	SELECT
		datalogger_id,
		datalogger_table_id,
		datalogger_table_name,
		fields
	FROM v_datalogger_equivalency_table
	WHERE datalogger_table_id = $1
`

// GetEquivalencyTable returns a single Datalogger EquivalencyTable
func (q *Queries) GetEquivalencyTable(ctx context.Context, dataloggerTableID uuid.UUID) (DataloggerEquivalencyTable, error) {
	var et DataloggerEquivalencyTable
	err := q.db.GetContext(ctx, &et, getEquivalencyTable, dataloggerTableID)
	return et, err
}

const createOrUpdateEquivalencyTableRow = `
	INSERT INTO datalogger_equivalency_table
	(datalogger_id, datalogger_table_id, field_name, display_name, instrument_id, timeseries_id)
	VALUES ($1, $2, $3, $4, $5, $6)
	ON CONFLICT ON CONSTRAINT datalogger_equivalency_table_datalogger_table_id_field_name_key
	DO UPDATE SET display_name = EXCLUDED.display_name, instrument_id = EXCLUDED.instrument_id, timeseries_id = EXCLUDED.timeseries_id
`

func (q *Queries) CreateOrUpdateEquivalencyTableRow(ctx context.Context, dataloggerID, dataloggerTableID uuid.UUID, tr DataloggerEquivalencyTableRow) error {
	if _, err := q.db.ExecContext(ctx, createOrUpdateEquivalencyTableRow,
		dataloggerID,
		dataloggerTableID,
		tr.FieldName,
		tr.DisplayName,
		tr.InstrumentID,
		tr.TimeseriesID,
	); err != nil {
		return err
	}
	return nil
}

const updateEquivalencyTableRow = `
	UPDATE datalogger_equivalency_table SET
		field_name = $2,
		display_name = $3,
		instrument_id = $4,
		timeseries_id = $5
	WHERE id = $1
`

func (q *Queries) UpdateEquivalencyTableRow(ctx context.Context, tr DataloggerEquivalencyTableRow) error {
	if _, err := q.db.ExecContext(ctx, updateEquivalencyTableRow,
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
	DELETE FROM datalogger_equivalency_table WHERE datalogger_table_id = $1
`

// DeleteEquivalencyTable clears all rows of the EquivalencyTable for a datalogger table
func (q *Queries) DeleteEquivalencyTable(ctx context.Context, dataloggerTableID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteEquivalencyTable, dataloggerTableID)
	return err
}

const deleteEquivalencyTableRow = `
	DELETE FROM datalogger_equivalency_table WHERE id = $1
`

// DeleteEquivalencyTableRow deletes a single EquivalencyTable row by row id
func (q *Queries) DeleteEquivalencyTableRow(ctx context.Context, rowID uuid.UUID) error {
	res, err := q.db.ExecContext(ctx, deleteEquivalencyTableRow, rowID)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("row not found %s", rowID)
	}
	return nil
}
