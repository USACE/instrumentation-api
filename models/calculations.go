package models

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Formula struct {
	InstrumentID uuid.UUID
	ParameterID  uuid.UUID
	Unit         uuid.UUID `json:"unit_name"`
	FormulaName  string    `json:"formula_name"`
	Formula      string    `json:"formula"`
}

const listFormulasSQL string = `
	SELECT *
	FROM calculation
`

// FormulasFactory converts database rows to Formula objects
func FormulasFactory(rows *sqlx.Rows) ([]Formula, error) {
	defer rows.Close()

	formulas := make([]Formula, 0)
	for rows.Next() {
		var f Formula
		err := rows.Scan(
			&f.InstrumentID, &f.ParameterID, &f.Unit, &f.FormulaName, &f.Formula,
		)
		if err != nil {
			return make([]Formula, 0), err
		}
		formulas = append(formulas, f)
	}

	return formulas, nil
}

// GetFormulas returns all formulas associated to a given instrument ID.
func GetFormulas(db *sqlx.DB, instrument *Instrument) ([]Formula, error) {

	rows, err := db.Queryx(listFormulasSQL+" WHERE instrument_id = $1", instrument.ID)
	if err != nil {
		return nil, err
	}
	ff, err := FormulasFactory(rows)
	if err != nil {
		return nil, err
	}

	return ff, nil
}
