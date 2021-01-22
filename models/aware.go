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
