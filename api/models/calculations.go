package models

import (
	"fmt"
	"reflect"
	"time"

	"github.com/USACE/instrumentation-api/api/timeseries"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/types"
)

type TimeseriesInfo struct {
	TimeseriesID uuid.UUID `json:"timeseries_id" db:"timeseries_id"`
	InstrumentID uuid.UUID `json:"instrument_id" db:"instrument_id"`
	Variable     string    `json:"variable" db:"variable"`
	IsComputed   bool      `json:"is_computed" db:"is_computed"`
	Formula      *string   `json:"formula" db:"formula"`
}

// Allows sending JSON Aggregated Data from the Database
// JSON Aggregated data not handled natively by sqlx
type DBTimeseries struct {
	TimeseriesInfo
	Measurements        string  `json:"measurements" db:"measurements"`
	NextMeasurementLow  *string `json:"next_measurement_low" db:"next_measurement_low"`
	NextMeasurementHigh *string `json:"next_measurement_high" db:"next_measurement_high"`
}

type Timeseries struct {
	TimeseriesInfo
	Measurements        []Measurement         `json:"measurements" db:"measurements"`
	NextMeasurementLow  *Measurement          `json:"next_measurement_low" db:"next_measurement_low"`
	NextMeasurementHigh *Measurement          `json:"next_measurement_high" db:"next_measurement_high"`
	TimeWindow          timeseries.TimeWindow `json:"time_window"`
}

type MeasurementCollection struct {
	TimeseriesID uuid.UUID     `json:"timeseries_id" db:"timeseries_id"`
	Items        []Measurement `json:"items"`
}

type Measurement struct {
	Time  time.Time `json:"time"`
	Value float64   `json:"value"`
	Error string    `json:"error,omitempty"`
}

type InclinometerTimeseries struct {
	TimeseriesInfo
	Measurements        []InclinometerMeasurement `json:"measurements" db:"measurements"`
	NextMeasurementLow  *Measurement              `json:"next_measurement_low" db:"next_measurement_low"`
	NextMeasurementHigh *Measurement              `json:"next_measurement_high" db:"next_measurement_high"`
	TimeWindow          timeseries.TimeWindow     `json:"time_window"`
}

type InclinometerMeasurement struct {
	Time   time.Time      `json:"time"`
	Values types.JSONText `json:"values"`
}

func (m Measurement) Lean() map[time.Time]float64 {
	return map[time.Time]float64{m.Time: m.Value}
}

func (m InclinometerMeasurement) InclinometerLean() map[time.Time]types.JSONText {
	return map[time.Time]types.JSONText{m.Time: m.Values}
}

type Calculation struct {
	// ID of the Formula, to be used in future requests.
	ID uuid.UUID `json:"id"`

	// Associated instrument.
	InstrumentID uuid.UUID `json:"instrument_id"`

	// Parameter that this formula should be outputting.
	ParameterID uuid.UUID `json:"parameter_id"`

	// Unit that this formula should be outputting.
	UnitID uuid.UUID `json:"unit_id"`

	Slug        string `json:"slug"`
	FormulaName string `json:"formula_name"`
	Formula     string `json:"formula"`
}

const listCalculationsSQL string = `
	SELECT
		id,
		instrument_id,
		parameter_id,
		unit_id,
		slug,
		name,
		COALESCE(contents, '') AS contents
	FROM v_timeseries_computed
`

// CalculationsFactory converts database rows to Calculation objects
func CalculationsFactory(rows *sqlx.Rows) ([]Calculation, error) {
	defer rows.Close()

	formulas := make([]Calculation, 0)
	for rows.Next() {
		var f Calculation
		err := rows.Scan(
			&f.ID, &f.InstrumentID, &f.ParameterID, &f.UnitID, &f.Slug, &f.FormulaName, &f.Formula,
		)
		if err != nil {
			return make([]Calculation, 0), err
		}
		formulas = append(formulas, f)
	}

	return formulas, nil
}

// GetInstrumentCalculations returns all formulas associated to a given instrument ID.
func GetInstrumentCalculations(db *sqlx.DB, instrument *Instrument) ([]Calculation, error) {

	rows, err := db.Queryx(listCalculationsSQL+" WHERE instrument_id = $1", instrument.ID)
	if err != nil {
		return nil, err
	}
	ff, err := CalculationsFactory(rows)
	if err != nil {
		return nil, err
	}

	return ff, nil
}

func ListCalculationSlugs(db *sqlx.DB) ([]string, error) {

	rows, err := db.Queryx("SELECT slug from v_timeseries_computed")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	slugs := make([]string, 0)
	for rows.Next() {
		var slug string
		err := rows.Scan(
			&slug,
		)
		if err != nil {
			return make([]string, 0), err
		}
		slugs = append(slugs, slug)
	}

	return slugs, nil
}

// CreateCalculation accepts a single Calculation instance and attempts to create it in
// the database, returning an error if anything goes wrong.
//
// Generating a UUID for the Calculation is not required. In the case that a Calculation
// is passed to this function **without** a set UUID field (i.e., `nil`), this function
// will set the UUID field to the one given to it by the database if the operation
// completes successfully. In the event that the function returns an error, the UUID
// field will remain unchanged.
func CreateCalculation(db *sqlx.DB, formula *Calculation) error {
	if reflect.ValueOf(formula.ParameterID).IsZero() {
		formula.ParameterID = uuid.Must(uuid.Parse("2b7f96e1-820f-4f61-ba8f-861640af6232"))
	}
	if reflect.ValueOf(formula.UnitID).IsZero() {
		formula.UnitID = uuid.Must(uuid.Parse("4a999277-4cf5-4282-93ce-23b33c65e2c8"))
	}
	txn, err := db.Beginx()
	if err != nil {
		return err
	}
	defer txn.Rollback()

	stmt1, err := txn.Preparex(`
		INSERT INTO timeseries (
			instrument_id,
			parameter_id,
			unit_id,
			slug,
			name
		)
		VALUES
			($1, $2, $3, $4, $5)
		RETURNING id
	`)
	if err != nil {
		return err
	}

	rows, err := stmt1.Queryx(&formula.InstrumentID, &formula.ParameterID, &formula.UnitID, &formula.Slug, &formula.FormulaName)
	if err != nil {
		return err
	}
	if !rows.Next() {
		return fmt.Errorf("rows should be non-empty")
	}
	if err := rows.Scan(&formula.ID); err != nil {
		return err
	}
	if rows.Next() {
		return fmt.Errorf("rows should be exactly one")
	}

	stmt2, err := txn.Preparex(`
		INSERT INTO calculation (timeseries_id, contents)
		VALUES ($1, $2)
		RETURNING timeseries_id
	`)
	if err != nil {
		return err
	}

	rows2, err := stmt2.Queryx(&formula.ID, &formula.Formula)
	if err != nil {
		return err
	}
	if !rows2.Next() {
		return fmt.Errorf("rows should be non-empty")
	}
	if err := rows2.Scan(&formula.ID); err != nil {
		return err
	}
	if rows2.Next() {
		return fmt.Errorf("rows should be exactly one")
	}

	if err := txn.Commit(); err != nil {
		return err
	}

	return nil
}

func UpdateCalculation(db *sqlx.DB, formula *Calculation) error {
	txn, err := db.Beginx()
	if err != nil {
		return err
	}
	defer txn.Rollback()

	var defaults Calculation
	row := txn.QueryRowx(listCalculationsSQL+` WHERE id = $1`, &formula.ID)
	if err := row.Scan(
		&defaults.ID,
		&defaults.InstrumentID,
		&defaults.ParameterID,
		&defaults.UnitID,
		&defaults.Slug,
		&defaults.FormulaName,
		&defaults.Formula,
	); err != nil {
		return err
	}

	// TODO: there is a better way of doing this using Golang's
	// [reflect](https://pkg.go.dev/reflect) package (other than
	// the way we are currently doing it).
	if reflect.ValueOf(formula.InstrumentID).IsZero() {
		formula.InstrumentID = defaults.InstrumentID
	}
	if reflect.ValueOf(formula.ParameterID).IsZero() {
		formula.ParameterID = defaults.ParameterID
	}
	if reflect.ValueOf(formula.UnitID).IsZero() {
		formula.UnitID = defaults.UnitID
	}
	if reflect.ValueOf(formula.Slug).IsZero() {
		formula.Slug = defaults.Slug
	}
	if reflect.ValueOf(formula.FormulaName).IsZero() {
		formula.FormulaName = defaults.FormulaName
	}
	if reflect.ValueOf(formula.Formula).IsZero() {
		formula.Formula = defaults.Formula
	}

	stmt, err := txn.Preparex(`
		INSERT INTO timeseries
			(
			 id,
			 instrument_id,
			 parameter_id,
			 unit_id,
			 slug,
			 name
			)
		VALUES
			($1, $2, $3, $4, $5, $6)
		ON CONFLICT (id) DO UPDATE SET
			instrument_id = COALESCE(EXCLUDED.instrument_id, $7),
			parameter_id = COALESCE(EXCLUDED.parameter_id, $8),
			unit_id = COALESCE(EXCLUDED.unit_id, $9),
			slug = COALESCE(EXCLUDED.slug, $10),
			name = COALESCE(EXCLUDED.name, $11)
		RETURNING
			id,
			instrument_id,
			parameter_id,
			unit_id,
			slug,
			name
	`)
	if err != nil {
		return err
	}
	rows, err := stmt.Queryx(
		&formula.ID,
		&formula.InstrumentID,
		&formula.ParameterID,
		&formula.UnitID,
		&formula.Slug,
		&formula.FormulaName,
		&defaults.InstrumentID,
		&defaults.ParameterID,
		&defaults.UnitID,
		&defaults.Slug,
		&defaults.FormulaName,
	)
	if err != nil {
		return err
	}
	if !rows.Next() {
		return fmt.Errorf("no results")
	}
	if err := rows.Scan(
		&formula.ID,
		&formula.InstrumentID,
		&formula.ParameterID,
		&formula.UnitID,
		&formula.Slug,
		&formula.FormulaName,
	); err != nil {
		return err
	}
	if err := rows.Close(); err != nil {
		return err
	}

	stmt2, err := txn.Preparex(`
		INSERT INTO calculation (timeseries_id, contents) VALUES ($1, $2)
		ON CONFLICT (timeseries_id) DO UPDATE SET
			contents = COALESCE(EXCLUDED.contents, $3)
		RETURNING contents
	`)
	if err != nil {
		return err
	}
	rows2, err := stmt2.Queryx(
		&formula.ID,
		&formula.Formula,
		&defaults.Formula,
	)
	if err != nil {
		return err
	}
	if !rows2.Next() {
		return fmt.Errorf("no results")
	}
	if err := rows2.Scan(
		&formula.Formula,
	); err != nil {
		return err
	}
	if err := rows2.Close(); err != nil {
		return err
	}

	if err := txn.Commit(); err != nil {
		return err
	}

	return nil
}

// DeleteCalculation removes the `Calculation` with ID `formulaID` from the database,
// effectively dissociating it from the instrument in question.
func DeleteCalculation(db *sqlx.DB, formulaID uuid.UUID) error {
	result, err := db.Exec("DELETE FROM timeseries WHERE id = $1 AND id IN (SELECT timeseries_id FROM calculation)", formulaID)
	if err != nil {
		return err
	}
	count, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("formula did not exist")
	}
	return nil
}
