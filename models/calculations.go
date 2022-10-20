package models

import (
	"errors"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Formula struct {
	// ID of the Formula, to be used in future requests.
	ID uuid.UUID

	// Associated instrument.
	InstrumentID uuid.UUID

	// Parameter that this formula should be outputting.
	ParameterID uuid.UUID

	// Unit that this formula should be outputting.
	Unit uuid.UUID `json:"unit_name"`

	FormulaName string `json:"formula_name"`
	Formula     string `json:"formula"`
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

// GetFormulasAsTimeseries returns all Formulas associated to a given Instrument
// in their computed timeseries form. If a Timeseries is not yet computed,
// this function has the side-effect of computing it.
func GetFormulasAsTimeseries(db *sqlx.DB, instrument *Instrument) ([]Timeseries, error) {
	return nil, nil
}

// CreateFormula accepts a single Formula instance and attempts to create it in
// the database, returning an error if anything goes wrong.
//
// Generating a UUID for the Formula is not required. In the case that a Formula
// is passed to this function **without** a set UUID field (i.e., `nil`), this function
// will set the UUID field to the one given to it by the database if the operation
// completes successfully. In the event that the function returns an error, the UUID
// field will remain unchanged.
func CreateFormula(db *sqlx.DB, formula *Formula) error {
	stmt := `
	INSERT INTO calculation
		(id,
		instrument_id,
		parameter_id,
		unit_id,
		name,
		contents
		)
	VALUES
		($1, $2, $3, $4, $5, $6)
	RETURNING id
	`

	rows, err := db.Query(stmt, formula.ID, formula.InstrumentID, formula.ParameterID, formula.Unit.ID, formula.FormulaName, formula.Formula)
	if err != nil {
		return err
	}
	if err := rows.Scan(formula.ID); err != nil {
		return err
	}
	if !rows.Next() {
		return errors.New("rows should be empty")
	}
	return nil
}
