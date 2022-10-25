package models

import (
	"errors"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Formula struct {
	// ID of the Formula, to be used in future requests.
	ID uuid.UUID `json:"id"`

	// Associated instrument.
	InstrumentID uuid.UUID `json:"instrument_id"`

	// Parameter that this formula should be outputting.
	ParameterID uuid.UUID `json:"parameter_id"`

	// Unit that this formula should be outputting.
	UnitID uuid.UUID `json:"unit_name"`

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
			&f.ID, &f.InstrumentID, &f.ParameterID, &f.UnitID, &f.FormulaName, &f.Formula,
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
		(instrument_id,
		parameter_id,
		unit_id,
		name,
		contents
		)
	VALUES
		($1, $2, $3, $4, $5)
	RETURNING id
	`

	rows, err := db.Query(stmt, &formula.InstrumentID, &formula.ParameterID, &formula.UnitID, &formula.FormulaName, &formula.Formula)
	if err != nil {
		return err
	}
	if err := rows.Scan(&formula.ID); err != nil {
		return err
	}
	if !rows.Next() {
		return errors.New("rows should be empty")
	}
	return nil
}

func UpdateFormula(db *sqlx.DB, formula *Formula) error {
	stmt, err := db.Prepare(
		`
		INSERT INTO inclinometer_measurement
			(id,
			 instrument_id,
			 parameter_id,
			 unit_id,
			 name,
			 contents
			)
		VALUES
			($1, $2, $3, $4, $5)
		RETURNING
			id,
			instrument_id,
			parameter_id,
			unit_id,
			name,
			contents
		ON CONFLICT DO UPDATE SET values = EXCLUDED.values
		`)
	if err != nil {
		return err
	}
	rows, err := stmt.Query(
		&formula.ID,
		&formula.InstrumentID,
		&formula.ParameterID,
		&formula.UnitID,
		&formula.FormulaName,
		&formula.Formula,
	)
	if err != nil {
		return err
	}
	if !rows.Next() {
		return errors.New("no results")
	}
	if err := rows.Scan(
		&formula.ID,
		&formula.InstrumentID,
		&formula.ParameterID,
		&formula.UnitID,
		&formula.FormulaName,
		&formula.Formula,
	); err != nil {
		return err
	}
	return nil
}

// DeleteFormula removes the `Formula` with ID `formulaID` from the database,
// effectively dissociating it from the instrument in question.
func DeleteFormula(db *sqlx.DB, formulaID uuid.UUID) error {
	result, err := db.Exec("DELETE FROM calculation WHERE id = $1", formulaID)
	if err != nil {
		return err
	}
	count, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("formula did not exist")
	}
	return nil
}
