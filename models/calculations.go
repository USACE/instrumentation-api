package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"time"

	"github.com/Knetic/govaluate"
	"github.com/USACE/instrumentation-api/timeseries"
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
	Measurements        []Measurement 		  `json:"measurements" db:"measurements"`
	NextMeasurementLow  *Measurement  		  `json:"next_measurement_low" db:"next_measurement_low"`
	NextMeasurementHigh *Measurement  		  `json:"next_measurement_high" db:"next_measurement_high"`
	TimeWindow          timeseries.TimeWindow `json:"time_window"`
}

type MeasurementCollection struct {
	TimeseriesID uuid.UUID     `json:"timeseries_id" db:"timeseries_id"`
	Items        []Measurement `json:"items"`
}

type Measurement struct {
	Time  time.Time `json:"time"`
	Value float64   `json:"value"`
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
	FROM calculation
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

	rows, err := db.Queryx("SELECT slug from calculation")
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

	stmt := `
	INSERT INTO calculation (
		instrument_id,
		parameter_id,
		unit_id,
		slug,
		name,
		contents
	)
	VALUES
		($1, $2, $3, $4, $5, $6)
	RETURNING id
	`

	rows, err := db.Query(stmt, &formula.InstrumentID, &formula.ParameterID, &formula.UnitID, &formula.Slug, &formula.FormulaName, &formula.Formula)
	if err != nil {
		return err
	}
	if !rows.Next() {
		return errors.New("rows should be non-empty")
	}
	if err := rows.Scan(&formula.ID); err != nil {
		return err
	}
	if rows.Next() {
		return errors.New("rows should be exactly one")
	}
	return nil
}

func UpdateCalculation(db *sqlx.DB, formula *Calculation) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	var defaults Calculation
	row := tx.QueryRow("SELECT instrument_id, parameter_id, unit_id, slug, name, contents FROM calculation WHERE id = $1", &formula.ID)
	if err := row.Scan(
		&defaults.InstrumentID,
		&defaults.ParameterID,
		&defaults.UnitID,
		&defaults.Slug,
		&defaults.FormulaName,
		&defaults.Formula,
	); err != nil {
		tx.Rollback()
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

	stmt, err := tx.Prepare(
		`
		INSERT INTO calculation
			(
			 id,
			 instrument_id,
			 parameter_id,
			 unit_id,
			 slug,
			 name,
			 contents
			)
		VALUES
			($1, $2, $3, $4, $5, $6, $7)
		ON CONFLICT (id) DO UPDATE SET
			instrument_id = COALESCE(EXCLUDED.instrument_id, $8),
			parameter_id = COALESCE(EXCLUDED.parameter_id, $9),
			unit_id = COALESCE(EXCLUDED.unit_id, $10),
			slug = COALESCE(EXCLUDED.slug, $11),
			name = COALESCE(EXCLUDED.name, $12),
			contents = COALESCE(EXCLUDED.contents, $13)
		RETURNING
			id,
			instrument_id,
			parameter_id,
			unit_id,
			slug,
			name,
			contents
		`)
	if err != nil {
		tx.Rollback()
		return err
	}
	rows, err := stmt.Query(
		&formula.ID,
		&formula.InstrumentID,
		&formula.ParameterID,
		&formula.UnitID,
		&formula.Slug,
		&formula.FormulaName,
		&formula.Formula,
		&defaults.InstrumentID,
		&defaults.ParameterID,
		&defaults.UnitID,
		&defaults.Slug,
		&defaults.FormulaName,
		&defaults.Formula,
	)
	if err != nil {
		tx.Rollback()
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
		&formula.Slug,
		&formula.FormulaName,
		&formula.Formula,
	); err != nil {
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

// DeleteCalculation removes the `Calculation` with ID `formulaID` from the database,
// effectively dissociating it from the instrument in question.
func DeleteCalculation(db *sqlx.DB, formulaID uuid.UUID) error {
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

// RegularizeCarryForward converts potentially irregular timeseries measurements into a regular
// interval timeseries over the time window w with measurements spaced at interval d
// Missing values are filled-in using a carry forward algorithm (use previous known value in time for missing values)
func (ts Timeseries) RegularizeCarryForward(w timeseries.TimeWindow, d time.Duration) (Timeseries, error) {

	regularized := make([]Measurement, 0)

	a := make([]Measurement, 0)
	// Array of Measurements to Work Against
	if ts.NextMeasurementLow != nil {
		a = append(a, *ts.NextMeasurementLow)
	}
	a = append(a, ts.Measurements...)
	if ts.NextMeasurementHigh != nil {
		a = append(a, *ts.NextMeasurementHigh)
	}

	t, tEnd, wkIdx := w.After, w.Before, 0

	if len(a) == 0 {
		return Timeseries{
			TimeseriesInfo: ts.TimeseriesInfo,
			Measurements:   make([]Measurement, 0),
			TimeWindow:     ts.TimeWindow,
		}, nil
	}

	for !t.After(tEnd) {
		// log.Printf("Working Index: %d; Working on time: %s; Comparing to: Time: %s ; Value: %f\n", wkIdx, t.Format(time.RFC3339), a[wkIdx].Time, a[wkIdx].Value)
		if t.Before(a[0].Time) {
			// log.Printf("time %s is below the minimum comparison time: %s; TS Value Literally Not Computable\n", t.Format(time.RFC3339), a[wkIdx].Time)
			t = t.Add(d)
			continue
		}
		if !t.Before(a[wkIdx].Time) {
			// log.Printf("time %s is equal to or greater than comparison time: %s\n", t.Format(time.RFC3339), a[wkIdx].Time)
			// If Already at the end of the interpable array
			if wkIdx == len(a)-1 {
				// log.Println("Already on the last index")
				// log.Printf("Set Value for time %s; Time: %s; Value: %f\n", t.Format(time.RFC3339), a[wkIdx].Time, a[wkIdx].Value)
				regularized = append(regularized, Measurement{t, a[wkIdx].Value})
				t = t.Add(d)
				continue
			}
			// log.Printf("ts: %s; Array Length: %d; Bump Working Index From %d --> %d\n", ts.TimeseriesID, len(a), wkIdx, wkIdx+1)
			wkIdx += 1
			continue
		}
		regularized = append(regularized, Measurement{t, a[wkIdx-1].Value})
		t = t.Add(d)
	}
	return Timeseries{
		TimeseriesInfo:      ts.TimeseriesInfo,
		Measurements:        regularized,
		NextMeasurementLow:  ts.NextMeasurementLow,
		NextMeasurementHigh: ts.NextMeasurementHigh,
	}, nil
}

func (ts Timeseries) RegularizeInterpolate(w timeseries.TimeWindow, d time.Duration) (Timeseries, error) {
	log.Println("Not Implemented")
	return Timeseries{}, errors.New("method not implemented")
}

func (ts *Timeseries) Calculate(variableMap map[time.Time]map[string]interface{}) error {
	expression, err := govaluate.NewEvaluableExpression(*ts.Formula)
	if err != nil {
		return err
	}
	t, end, interval := ts.TimeWindow.After, ts.TimeWindow.Before, time.Hour
	for !t.After(end) {
		if params, exists := variableMap[t]; exists {
			valStr, err := expression.Evaluate(params)
			if err != nil {
				t = t.Add(interval)
				continue
			}
			val64, err := strconv.ParseFloat(fmt.Sprint(valStr), 64)
			if err != nil {
				t = t.Add(interval)
				continue
			}
			ts.Measurements = append(ts.Measurements, Measurement{Time: t, Value: val64})
		}
		t = t.Add(interval)
	}
	return nil
}

// ComputedTimeseries returns computed and stored timeseries for a specified array of instrument IDs
func ComputedTimeseries(db *sqlx.DB, instrumentIDs []uuid.UUID, tw *timeseries.TimeWindow, interval *time.Duration) ([]Timeseries, error) {

	tt := make([]DBTimeseries, 0)
	sql := `
	-- Get Timeseries and Dependencies for Calculations
	-- timeseries required based on requested instrument
	WITH requested_instruments AS (
		SELECT id
		FROM instrument
		WHERE id IN (?)
	), required_timeseries AS (
	-- 	Timeseries for Instrument
		SELECT id FROM timeseries WHERE instrument_id IN (SELECT id FROM requested_instruments)
		UNION
	-- Dependencies for Instrument Timeseries
		SELECT dependency_timeseries_id AS id
		FROM v_timeseries_dependency
		WHERE instrument_id IN (SELECT id from requested_instruments)
	),
	-- Next Timeseries Measurement Outside Time Window (Earlier); Needed for Calculation Interpolation
	next_low AS (
		SELECT nlm.timeseries_id AS timeseries_id, json_build_object('time', nlm.time, 'value', m1.value) AS measurement
		FROM (
			SELECT timeseries_id, MAX(time) AS time
			FROM timeseries_measurement
			WHERE timeseries_id IN (SELECT id FROM required_timeseries) AND time < ?
			GROUP BY timeseries_id
		) nlm
		INNER JOIN timeseries_measurement m1 ON m1.time = nlm.time AND m1.timeseries_id = nlm.timeseries_id
	),
	-- Next Timeseries Measurement Outside Time Window (Later); Needed For Calculation Interpolation
	next_high AS (
		SELECT nhm.timeseries_id AS timeseries_id, json_build_object('time', nhm.time, 'value', m2.value) AS measurement
		FROM (
			SELECT timeseries_id, MIN(time) AS time
			FROM timeseries_measurement
			WHERE timeseries_id IN (SELECT id FROM required_timeseries) AND time > ?
			GROUP BY timeseries_id
		) nhm
		INNER JOIN timeseries_measurement m2 ON m2.time = nhm.time AND m2.timeseries_id = nhm.timeseries_id
	),
	-- Measurements Within Time Window by timeseries_id;
	measurements AS (
		SELECT timeseries_id,
			   json_agg(json_build_object('time', time, 'value', value) ORDER BY time ASC)::text AS measurements
		FROM timeseries_measurement
		WHERE timeseries_id IN (SELECT id FROM required_timeseries) AND time >= ? AND time <= ?
		GROUP BY timeseries_id
	)
	-- Stored Timeseries
	SELECT rt.id                          AS timeseries_id,
		   ts.instrument_id               AS instrument_id,
		   i.slug || '.' || ts.slug       AS variable,
		   false                          AS is_computed,
		   null                           AS formula,
		   COALESCE(m.measurements, '[]') AS measurements,
		   nl.measurement::text           AS next_measurement_low,
		   nh.measurement::text           AS next_measurement_high
	FROM required_timeseries rt
	INNER JOIN timeseries ts ON ts.id = rt.id
	INNER JOIN instrument i ON i.id = ts.instrument_id AND i.id IN (SELECT id FROM requested_instruments)
	LEFT JOIN measurements m ON m.timeseries_id = rt.id
	LEFT JOIN next_low nl ON nl.timeseries_id = rt.id
	LEFT JOIN next_high nh ON nh.timeseries_id = rt.id
	UNION
	-- Computed Timeseries
	SELECT cc.id                AS timeseries_id,
		cc.instrument_id        AS instrument_id,
		cc.slug			        AS variable,
		true                    AS is_computed,
		cc.contents             AS formula,
		'[]'::text              AS measurements,
		null                    AS next_measurement_low,
		null                    AS next_measurement_high
	FROM calculation cc
	WHERE cc.contents IS NOT NULL AND cc.instrument_id IN (SELECT id FROM requested_instruments)
	ORDER BY is_computed
	`

	query, args, err := sqlx.In(sql, instrumentIDs, tw.After, tw.Before, tw.After, tw.Before)
	if err != nil {
		return make([]Timeseries, 0), err
	}
	query = db.Rebind(query)
	if err := db.Select(&tt, query, args...); err != nil {
		return make([]Timeseries, 0), err
	}

	// Unmarshal JSON Strings
	tt2 := make([]Timeseries, len(tt))
	for idx, t := range tt {
		tt2[idx] = Timeseries{
			TimeseriesInfo: t.TimeseriesInfo,
			Measurements:   make([]Measurement, 0),
			TimeWindow:     *tw,
		}
		if err := json.Unmarshal([]byte(t.Measurements), &tt2[idx].Measurements); err != nil {
			log.Println(err)
		}
		if t.NextMeasurementHigh != nil {
			if err := json.Unmarshal([]byte(*t.NextMeasurementHigh), &tt2[idx].NextMeasurementHigh); err != nil {
				log.Println(err)
			}
		}
		if t.NextMeasurementLow != nil {
			if err := json.Unmarshal([]byte(*t.NextMeasurementLow), &tt2[idx].NextMeasurementLow); err != nil {
				log.Println(err)
			}
		}
	}

	// Final Timeseries to be returned
	tt3 := make([]Timeseries, 0)

	// a map of all available parameters for a given time slice
	variableMap := make(map[time.Time]map[string]interface{})

	// todo: Optimization - do not need to regularize all timeseries
	// only need to regularize those that will be used as calculation dependencies
	for _, ts := range tt2 {
		tsReg, err := ts.RegularizeCarryForward(*tw, *interval)
		if err != nil {
			return make([]Timeseries, 0), err
		}
		// Add All Measurements from Regularized Timeseries to Map
		for _, m := range tsReg.Measurements {
			if _, exists := variableMap[m.Time]; !exists {
				variableMap[m.Time] = make(map[string]interface{})
			}
			variableMap[m.Time][ts.Variable] = m.Value
		}
		// If not a computed timeseries, add raw version of timeseries to response
		if !ts.IsComputed {
			tt3 = append(tt3, ts)
			continue
		}

		// Calculations
		// It is known that all stored timeseries have been added to the Map and calculations
		// can now be run because calculated timeseries (identified by .IsComputed)
		// are returned from the database last in the query using ORDER BY is_computed
		err = ts.Calculate(variableMap)
		if err != nil {
			log.Printf("Error Computing Formula for Timeseries %s\n", ts.TimeseriesID)
			continue
		}
		tt3 = append(tt3, ts)
	}

	return tt3, nil
}

// ComputedInclinometerTimeseries returns computed and stored inclinometer timeseries for a specified array of instrument IDs
func ComputedInclinometerTimeseries(db *sqlx.DB, instrumentIDs []uuid.UUID, tw *timeseries.TimeWindow, interval *time.Duration) ([]InclinometerTimeseries, error) {

	tt := make([]DBTimeseries, 0)
	sql := `
	-- Get Timeseries and Dependencies for Calculations
	-- timeseries required based on requested instrument
	WITH requested_instruments AS (
		SELECT id
		FROM instrument
		WHERE id IN (?)
	), required_timeseries AS (
	-- 	Timeseries for Instrument
		SELECT id FROM timeseries WHERE instrument_id IN (SELECT id FROM requested_instruments)
		UNION
	-- Dependencies for Instrument Timeseries
		SELECT dependency_timeseries_id AS id
		FROM v_timeseries_dependency
		WHERE instrument_id IN (SELECT id from requested_instruments)
	),
	-- Measurements Within Time Window by timeseries_id;
	measurements AS (
		SELECT timeseries_id,
			   json_agg(json_build_object('time', time, 'values', values) ORDER BY time ASC)::text AS measurements
		FROM inclinometer_measurement
		WHERE timeseries_id IN (SELECT id FROM required_timeseries) AND time >= ? AND time <= ?
		GROUP BY timeseries_id
	)
	-- Stored Timeseries
	SELECT rt.id                          AS timeseries_id,
		   ts.instrument_id               AS instrument_id,
		   i.slug || '.' || ts.slug       AS variable,
		   false                          AS is_computed,
		   null                           AS formula,
		   COALESCE(m.measurements, '[]') AS measurements
	FROM required_timeseries rt
	INNER JOIN timeseries ts ON ts.id = rt.id
	INNER JOIN instrument i ON i.id = ts.instrument_id AND i.id IN (SELECT id FROM requested_instruments)
	LEFT JOIN measurements m ON m.timeseries_id = rt.id
	UNION
	-- Computed Timeseries
	SELECT cc.id                   AS timeseries_id,
		   cc.instrument_id        AS instrument_id,
		   
		   -- TODO: make this component of the query a 'slug'-type.
		   cc.name			       AS variable,
		   
		   true                    AS is_computed,
		   cc.contents             AS formula,
		   '[]'::text              AS measurements,
		   null                    AS next_measurement_low,
		   null                    AS next_measurement_high
	FROM calculation cc
	WHERE cc.contents IS NOT NULL AND cc.instrument_id IN (SELECT id FROM requested_instruments)
	ORDER BY is_computed
	`

	query, args, err := sqlx.In(sql, instrumentIDs, tw.After, tw.Before)
	if err != nil {
		return make([]InclinometerTimeseries, 0), err
	}
	query = db.Rebind(query)
	if err := db.Select(&tt, query, args...); err != nil {
		return make([]InclinometerTimeseries, 0), err
	}

	// Unmarshal JSON Strings
	tt2 := make([]InclinometerTimeseries, len(tt))
	for idx, t := range tt {
		tt2[idx] = InclinometerTimeseries{
			TimeseriesInfo: t.TimeseriesInfo,
			Measurements:   make([]InclinometerMeasurement, 0),
			TimeWindow:     *tw,
		}

		cm, err := ConstantMeasurement(db, &t.TimeseriesID, "inclinometer-constant")
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal([]byte(t.Measurements), &tt2[idx].Measurements); err != nil {
			log.Println(err)
		}

		for i := range tt2[idx].Measurements {
			values, err := ListInclinometerMeasurementValues(db, &t.TimeseriesID, tt2[idx].Measurements[i].Time, cm.Value)
			if err != nil {
				return nil, err
			}

			jsonValues, err := json.Marshal(values)
			if err != nil {
				return nil, err
			}
			tt2[idx].Measurements[i].Values = jsonValues
		}
	}

	return tt2, nil
}
