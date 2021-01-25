package models

import (
	"fmt"

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
// { projectID: <uuid4>, instrument_id: <uuid4>, aware_platform_id: <uuid4>, aware_parameters: { <string>: <uuid4> } }
// aware_parameters is a map of <aware_parameter_key> : <timeseries_id>
type AwarePlatformParameterConfig struct {
	ProjectID       uuid.UUID             `json:"project_id" db:"project_id"`
	InstrumentID    uuid.UUID             `json:"instrument_id" db:"instrument_id"`
	AwarePlatformID uuid.UUID             `json:"aware_platform_id" db:"aware_platform_id"`
	AwareParameters map[string]*uuid.UUID `json:"aware_parameters"`
}

type awarePlatformParameterEnabled struct {
	ProjectID         uuid.UUID  `json:"project_id" db:"project_id"`
	InstrumentID      uuid.UUID  `json:"instrument_id" db:"instrument_id"`
	AwarePlatformID   uuid.UUID  `json:"aware_platform_id" db:"aware_platform_id"`
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

	sql := `SELECT i.project_id          AS project_id,
	               i.id                  AS instrument_id,
				   e.aware_platform_id   AS aware_platform_id,
				   b.key                 AS aware_parameter_key,
				   t.id                  AS timeseries_id
            FROM aware_platform_parameter_enabled e
            INNER JOIN aware_platform a ON a.id = e.aware_platform_id
            INNER JOIN instrument i ON i.id = a.instrument_id
            INNER JOIN aware_parameter b ON b.id = e.aware_parameter_id
			LEFT JOIN timeseries t ON t.instrument_id=i.id AND t.parameter_id=b.parameter_id AND t.unit_id=b.unit_id
			ORDER BY e.aware_platform_id, b.key`
	aa := make([]awarePlatformParameterEnabled, 0)
	if err := db.Select(&aa, sql); err != nil {
		return make([]awarePlatformParameterEnabled, 0), err
	}
	return aa, nil
}

func ListAwarePlatformParameterConfig(db *sqlx.DB) ([]AwarePlatformParameterConfig, error) {
	ee, err := listAwarePlatformParameterEnabled(db)
	fmt.Println(len(ee))
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
				AwarePlatformID: e.AwarePlatformID,
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
