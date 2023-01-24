package models

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// Unit is a unit data structure
type Unit struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Abbreviation string    `json:"abbreviation"`
	UnitFamilyID uuid.UUID `json:"unit_family_id" db:"unit_family_id"`
	UnitFamily   string    `json:"unit_family" db:"unit_family"`
	MeasureID    uuid.UUID `json:"measure_id" db:"measure_id"`
	Measure      string    `json:"measure"`
}

// ListUnits returns a slice of units
func ListUnits(db *sqlx.DB) ([]Unit, error) {
	uu := make([]Unit, 0)
	if err := db.Select(
		&uu,
		`SELECT id, name, abbreviation, unit_family_id, unit_family, measure_id, measure
		 FROM v_unit
		 ORDER BY name`,
	); err != nil {
		return make([]Unit, 0), err
	}
	return uu, nil
}
