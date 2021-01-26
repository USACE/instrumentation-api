package models

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

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

type awarePlatformParameterEnabled struct {
	ProjectID         uuid.UUID  `json:"project_id" db:"project_id"`
	InstrumentID      uuid.UUID  `json:"instrument_id" db:"instrument_id"`
	AwareID           uuid.UUID  `json:"aware_id" db:"aware_id"`
	AwareParameterKey string     `json:"aware_parameter_key" db:"aware_parameter_key"`
	TimeseriesID      *uuid.UUID `json:"timeseries_id" db:"timeseries_id"`
}

func ListAwareParameters(db *sqlx.DB) ([]AwareParameter, error) {
	pp := make([]AwareParameter, 0)
	if err := db.Select(
		&pp, "SELECT id, key, parameter_id, unit_id FROM aware_parameter",
	); err != nil {
		return make([]AwareParameter, 0), err
	}
	return pp, nil
}

func listAwarePlatformParameterEnabled(db *sqlx.DB) ([]awarePlatformParameterEnabled, error) {

	sql := `SELECT project_id, instrument_id, aware_id, aware_parameter_key, timeseries_id
			FROM v_aware_platform_parameter_enabled
			ORDER BY project_id, aware_id, aware_parameter_key`

	aa := make([]awarePlatformParameterEnabled, 0)
	if err := db.Select(&aa, sql); err != nil {
		return make([]awarePlatformParameterEnabled, 0), err
	}
	return aa, nil
}

func ListAwarePlatformParameterConfig(db *sqlx.DB) ([]AwarePlatformParameterConfig, error) {
	ee, err := listAwarePlatformParameterEnabled(db)
	if err != nil {
		return make([]AwarePlatformParameterConfig, 0), err
	}
	// reorganize aware_parameter_key, timeseries_id into map for each instrument
	// Map of aware parameters to timeseries
	m1 := make(map[uuid.UUID]AwarePlatformParameterConfig)
	for _, e := range ee {
		if _, ok := m1[e.InstrumentID]; !ok {
			m1[e.InstrumentID] = AwarePlatformParameterConfig{
				ProjectID:       e.ProjectID,
				InstrumentID:    e.InstrumentID,
				AwareID:         e.AwareID,
				AwareParameters: make(map[string]*uuid.UUID),
			}
		}
		m1[e.InstrumentID].AwareParameters[e.AwareParameterKey] = e.TimeseriesID
	}

	cc := make([]AwarePlatformParameterConfig, 0)
	for k := range m1 {
		cc = append(cc, m1[k])
	}
	return cc, nil
}
