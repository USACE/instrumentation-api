package model

import (
	"context"

	"github.com/google/uuid"
)

type CalculatedTimeseries struct {
	ID           uuid.UUID `json:"id" db:"id"`
	InstrumentID uuid.UUID `json:"instrument_id" db:"instrument_id"`
	ParameterID  uuid.UUID `json:"parameter_id" db:"parameter_id"`
	UnitID       uuid.UUID `json:"unit_id" db:"unit_id"`
	Slug         string    `json:"slug" db:"slug"`
	FormulaName  string    `json:"formula_name" db:"formula_name"`
	Formula      string    `json:"formula" db:"formula"`
}

const listCalculatedTimeseriesSQL = `
	SELECT
		id,
		instrument_id,
		parameter_id,
		unit_id,
		slug,
		name AS formula_name,
		COALESCE(contents, '') AS formula
	FROM v_timeseries_computed
`

const getAllCalculatedTimeseriesForInstrument = listCalculatedTimeseriesSQL + `
	WHERE instrument_id = $1
`

// GetAllCalculationsForInstrument returns all formulas associated to a given instrument ID.
func (q *Queries) GetAllCalculatedTimeseriesForInstrument(ctx context.Context, instrumentID uuid.UUID) ([]CalculatedTimeseries, error) {
	cc := make([]CalculatedTimeseries, 0)
	if err := q.db.SelectContext(ctx, &cc, getAllCalculatedTimeseriesForInstrument, instrumentID); err != nil {
		return nil, err
	}
	return cc, nil
}

const createCalculatedTimeseries = `
	INSERT INTO timeseries (
		instrument_id,
		parameter_id,
		unit_id,
		slug,
		name
	) VALUES ($1, $2, $3, slugify($4, 'timeseries'), $4)
	RETURNING id
`

func (q *Queries) CreateCalculatedTimeseries(ctx context.Context, cc CalculatedTimeseries) (uuid.UUID, error) {
	if cc.ParameterID == uuid.Nil {
		cc.ParameterID = unknownParameterID
	}
	if cc.UnitID == uuid.Nil {
		cc.UnitID = unknownUnitID
	}
	var tsID uuid.UUID
	err := q.db.GetContext(ctx, &tsID, createCalculatedTimeseries, &cc.InstrumentID, &cc.ParameterID, &cc.UnitID, &cc.FormulaName)
	return tsID, err
}

const createCalculation = `
	INSERT INTO calculation (timeseries_id, contents) VALUES ($1,$2)
`

func (q *Queries) CreateCalculation(ctx context.Context, timeseriesID uuid.UUID, contents string) error {
	_, err := q.db.ExecContext(ctx, createCalculation, timeseriesID, contents)
	return err
}

const getOneCalculation = listCalculatedTimeseriesSQL + `
	WHERE id = $1
`

func (q *Queries) GetOneCalculation(ctx context.Context, calculationID *uuid.UUID) (CalculatedTimeseries, error) {
	var defaultCc CalculatedTimeseries
	err := q.db.GetContext(ctx, &defaultCc, getOneCalculation, calculationID)
	return defaultCc, err
}

const createOrUpdateCalculation = `
	INSERT INTO calculation (timeseries_id, contents) VALUES ($1, $2)
	ON CONFLICT (timeseries_id) DO UPDATE SET contents = COALESCE(EXCLUDED.contents, $3)
`

func (q *Queries) CreateOrUpdateCalculation(ctx context.Context, timeseriesID uuid.UUID, formula, defaultFormula string) error {
	_, err := q.db.ExecContext(ctx, createOrUpdateCalculation, timeseriesID, formula, defaultFormula)
	return err
}

const deleteCalculatedTimeseries = `
	DELETE FROM timeseries WHERE id = $1 AND id IN (SELECT timeseries_id FROM calculation)
`

func (q *Queries) DeleteCalculatedTimeseries(ctx context.Context, calculationID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteCalculatedTimeseries, calculationID)
	return err
}

const createOrUpdateCalculatedTimeseries = `
	INSERT INTO timeseries (
		 id,
		 instrument_id,
		 parameter_id,
		 unit_id,
		 slug,
		 name
	) VALUES ($1, $2, $3, $4, slugify($5, 'timeseries'), $5)
	ON CONFLICT (id) DO UPDATE SET
		instrument_id = COALESCE(EXCLUDED.instrument_id, $6),
		parameter_id = COALESCE(EXCLUDED.parameter_id, $7),
		unit_id = COALESCE(EXCLUDED.unit_id, $8),
		slug = COALESCE(EXCLUDED.slug, slugify($9, 'timeseries')),
		name = COALESCE(EXCLUDED.name, $9)
`

func (q *Queries) CreateOrUpdateCalculatedTimeseries(ctx context.Context, cc CalculatedTimeseries, defaultCc CalculatedTimeseries) error {
	if _, err := q.db.ExecContext(ctx, createOrUpdateCalculatedTimeseries,
		cc.ID,
		cc.InstrumentID,
		cc.ParameterID,
		cc.UnitID,
		cc.FormulaName,
		defaultCc.InstrumentID,
		defaultCc.ParameterID,
		defaultCc.UnitID,
		defaultCc.FormulaName,
	); err != nil {
		return err
	}
	return nil
}
