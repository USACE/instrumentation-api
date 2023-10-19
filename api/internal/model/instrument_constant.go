package model

import (
	"context"

	"github.com/google/uuid"
)

const listInstrumentConstants = listTimeseries + `
	WHERE instrument_id = $1 AND id IN (
		SELECT timeseries_id
		FROM instrument_constants
		WHERE instrument_id = $1
	)
`

// ListInstrumentConstants lists constants for a given instrument id
func (q *Queries) ListInstrumentConstants(ctx context.Context, instrumentID uuid.UUID) ([]Timeseries, error) {
	tt := make([]Timeseries, 0)
	if err := q.db.SelectContext(ctx, &tt, listInstrumentConstants, instrumentID); err != nil {
		return tt, err
	}
	return tt, nil
}

const createInstrumentConstant = `
	INSERT INTO instrument_constants (instrument_id, timeseries_id) VALUES ($1, $2)
`

func (q *Queries) CreateInstrumentConstant(ctx context.Context, instrumentID, timeseriesID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, createInstrumentConstant, instrumentID, timeseriesID)
	return err
}

const deleteInstrumentConstant = `
	DELETE FROM instrument_constants WHERE instrument_id = $1 AND timeseries_id = $2
`

func (q *Queries) DeleteInstrumentConstant(ctx context.Context, instrumentID, timeseriesID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteInstrumentConstant, instrumentID, timeseriesID)
	return err
}
