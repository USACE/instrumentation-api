package model

import (
	"context"

	"github.com/google/uuid"
)

// AwareParameter struct
type AwareParameter struct {
	ID          uuid.UUID `json:"id"`
	Key         string    `json:"key"`
	ParameterID uuid.UUID `json:"parameter_id" db:"parameter_id"`
	UnitID      uuid.UUID `json:"unit_id" db:"unit_id"`
}

// AwarePlatformParameterConfig holds information about which parameters are "enabled" for given instrument(s)
// { projectID: <uuid4>, instrument_id: <uuid4>, aware_id: <uuid4>, aware_parameters: { <string>: <uuid4> } }
// aware_parameters is a map of <aware_parameter_key> : <timeseries_id>
type AwarePlatformParameterConfig struct {
	InstrumentID    uuid.UUID             `json:"instrument_id" db:"instrument_id"`
	AwareID         uuid.UUID             `json:"aware_id" db:"aware_id"`
	AwareParameters map[string]*uuid.UUID `json:"aware_parameters"`
}

type AwarePlatformParameterEnabled struct {
	InstrumentID      uuid.UUID  `json:"instrument_id" db:"instrument_id"`
	AwareID           uuid.UUID  `json:"aware_id" db:"aware_id"`
	AwareParameterKey string     `json:"aware_parameter_key" db:"aware_parameter_key"`
	TimeseriesID      *uuid.UUID `json:"timeseries_id" db:"timeseries_id"`
}

const listAwareParameters = `
	SELECT id, key, parameter_id, unit_id FROM aware_parameter
`

// ListAwareParameters returns aware parameters
func (q *Queries) ListAwareParameters(ctx context.Context) ([]AwareParameter, error) {
	pp := make([]AwareParameter, 0)
	if err := q.db.SelectContext(ctx, &pp, listAwareParameters); err != nil {
		return nil, err
	}
	return pp, nil
}

const listAwarePlatformParameterEnabled = `
	SELECT instrument_id, aware_id, aware_parameter_key, timeseries_id
	FROM v_aware_platform_parameter_enabled
	ORDER BY project_id, aware_id, aware_parameter_key
`

func (q *Queries) ListAwarePlatformParameterEnabled(ctx context.Context) ([]AwarePlatformParameterEnabled, error) {
	aa := make([]AwarePlatformParameterEnabled, 0)
	if err := q.db.SelectContext(ctx, &aa, listAwarePlatformParameterEnabled); err != nil {
		return nil, err
	}
	return aa, nil
}

const createAwarePlatform = `
	INSERT INTO aware_platform (instrument_id, aware_id) VALUES ($1, $2)
`

func (q *Queries) CreateAwarePlatform(ctx context.Context, instrumentID, awareID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, createAwarePlatform, &instrumentID, &awareID)
	return err
}
