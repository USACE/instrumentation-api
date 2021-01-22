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

func ListAwareParameters(db *sqlx.DB) ([]AwareParameter, error) {
	pp := make([]AwareParameter, 0)
	if err := db.Select(
		&pp, "SELECT id, key, parameter_id, unit_id FROM aware_parameter",
	); err != nil {
		return make([]AwareParameter, 0), err
	}
	return pp, nil
}

// func ListAwareEnabled(db *sqlx.DB) error {
// 	sql := `
// 	SELECT i.project_id          AS project_id,
// 	       i.id                  AS instrument_id,
// 	       e.aware_platform_id   AS platform_id,
// 	       e.key                 AS aware_parameter_key,
// 	       t.id                  AS timeseries_id
//     FROM aware_platform_parameter_enabled e
//     INNER JOIN aware_platform a ON a.platform_id = e.aware_platform_id
//     INNER JOIN instrument i ON i.id = a.instrument_id
//     INNER JOIN aware_parameter b ON b.id = e.aware_parameter_id
// 	LEFT JOIN timeseries t ON t.instrument_id=i.id, t.parameter_id=b.parameter_id, t.unit_id=b.parameter_id   `

// 	return nil
// }
