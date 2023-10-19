package model

import (
	"context"

	"github.com/google/uuid"
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

const listUnits = `
	SELECT id, name, abbreviation, unit_family_id, unit_family, measure_id, measure
	FROM v_unit
	ORDER BY name
`

// ListUnits returns a slice of units
func (q *Queries) ListUnits(ctx context.Context) ([]Unit, error) {
	uu := make([]Unit, 0)
	if err := q.db.SelectContext(ctx, &uu, listUnits); err != nil {
		return nil, err
	}
	return uu, nil
}
