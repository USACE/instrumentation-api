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
	ProjectID       uuid.UUID             `json:"project_id" db:"project_id"`
	InstrumentID    uuid.UUID             `json:"instrument_id" db:"instrument_id"`
	AwareID         uuid.UUID             `json:"aware_id" db:"aware_id"`
	AwareParameters map[string]*uuid.UUID `json:"aware_parameters"`
}

type AwarePlatformParameterEnabled struct {
	ProjectID         uuid.UUID  `json:"project_id" db:"project_id"`
	InstrumentID      uuid.UUID  `json:"instrument_id" db:"instrument_id"`
	AwareID           uuid.UUID  `json:"aware_id" db:"aware_id"`
	AwareParameterKey string     `json:"aware_parameter_key" db:"aware_parameter_key"`
	TimeseriesID      *uuid.UUID `json:"timeseries_id" db:"timeseries_id"`
}

// ListAwareParameters returns aware parameters
func (q *Queries) ListAwareParameters(ctx context.Context) ([]AwareParameter, error) {
	c := `
		SELECT id, key, parameter_id, unit_id FROM aware_parameter
	`
	pp := make([]AwareParameter, 0)
	if err := q.db.SelectContext(ctx, &pp, c); err != nil {
		return make([]AwareParameter, 0), err
	}
	return pp, nil
}

func (q *Queries) ListAwarePlatformParameterEnabled(ctx context.Context) ([]AwarePlatformParameterEnabled, error) {
	c := `
		SELECT project_id, instrument_id, aware_id, aware_parameter_key, timeseries_id
		FROM v_aware_platform_parameter_enabled
		ORDER BY project_id, aware_id, aware_parameter_key
	`
	aa := make([]AwarePlatformParameterEnabled, 0)
	if err := q.db.SelectContext(ctx, &aa, c); err != nil {
		return make([]AwarePlatformParameterEnabled, 0), err
	}
	return aa, nil
}
