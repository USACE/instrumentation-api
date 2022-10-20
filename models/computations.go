package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/Knetic/govaluate"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/types"
)

// TimeWindow is a bounding box for time
type TimeWindow struct {
	After  time.Time `json:"after" query:"after"`
	Before time.Time `json:"before" query:"before"`
}

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
	Measurements        []Measurement `json:"measurements" db:"measurements"`
	NextMeasurementLow  *Measurement  `json:"next_measurement_low" db:"next_measurement_low"`
	NextMeasurementHigh *Measurement  `json:"next_measurement_high" db:"next_measurement_high"`
	TimeWindow          TimeWindow    `json:"time_window"`
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
	TimeWindow          TimeWindow                `json:"time_window"`
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

// RegularizeCarryForward converts potentially irregular timeseries measurements into a regular
// interval timeseries over the time window w with measurements spaced at interval d
// Missing values are filled-in using a carry forward algorithm (use previous known value in time for missing values)
func (ts Timeseries) RegularizeCarryForward(w TimeWindow, d time.Duration) (Timeseries, error) {

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

func (ts Timeseries) RegularizeInterpolate(w TimeWindow, d time.Duration) (Timeseries, error) {
	log.Println("Not Implemented")
	return Timeseries{}, errors.New("Method Not Implemented")
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
func ComputedTimeseries(db *sqlx.DB, instrumentIDs []uuid.UUID, tw *TimeWindow, interval *time.Duration) ([]Timeseries, error) {

	tt := make([]DBTimeseries, 0)
	sql := `
	-- Get Timeseries and Dependencies for Computations
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
	SELECT r.id                     AS timeseries_id,
		   ts.instrument_id         AS instrument_id,
		   i.slug || '.' || ts.slug AS variable,
		   false                    AS is_computed,
		   null                     AS formula,
		   COALESCE(m.measurements, '[]') AS measurements,
		   nl.measurement::text     AS next_measurement_low,
		   nh.measurement::text     AS next_measurement_high
	FROM required_timeseries r
	INNER JOIN timeseries ts ON ts.id = r.id
	INNER JOIN instrument i ON i.id = ts.instrument_id AND i.id IN (SELECT id FROM requested_instruments)
	LEFT JOIN measurements m ON m.timeseries_id = r.id
	LEFT JOIN next_low nl ON nl.timeseries_id = r.id
	LEFT JOIN next_high nh ON nh.timeseries_id = r.id
	UNION
	-- Computed Timeseries
	SELECT i.formula_id            AS timeseries_id,
		   i.id                    AS instrument_id,
		   i.slug || '.formula'    AS variable,
		   true                    AS is_computed,
		   i.formula               AS formula,
		   '[]'::text              AS measurements,
		   null                    AS next_measurement_low,
		   null                    AS next_measurement_high
	FROM instrument i
	WHERE i.formula IS NOT NULL AND i.id IN (SELECT id FROM requested_instruments)
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
	// only need to regularize those that will be used as computation dependencies
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

		// Computations
		// It is known that all stored timeseries have been added to the Map and computations
		// can now be run because alculated timeseries (identified by .IsComputed)
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
func ComputedInclinometerTimeseries(db *sqlx.DB, instrumentIDs []uuid.UUID, tw *TimeWindow, interval *time.Duration) ([]InclinometerTimeseries, error) {

	tt := make([]DBTimeseries, 0)
	sql := `
	-- Get Timeseries and Dependencies for Computations
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
	SELECT r.id                     AS timeseries_id,
		   ts.instrument_id         AS instrument_id,
		   i.slug || '.' || ts.slug AS variable,
		   false                    AS is_computed,
		   null                     AS formula,
		   COALESCE(m.measurements, '[]') AS measurements
	FROM required_timeseries r
	INNER JOIN timeseries ts ON ts.id = r.id
	INNER JOIN instrument i ON i.id = ts.instrument_id AND i.id IN (SELECT id FROM requested_instruments)
	LEFT JOIN measurements m ON m.timeseries_id = r.id
	UNION
	-- Computed Timeseries
	SELECT i.id                    AS timeseries_id,
		   i.instrument_id         AS instrument_id,
		   
		   -- TODO: make this component of the query a 'slug'-type.
		   i.name			       AS variable,
		   
		   true                    AS is_computed,
		   i.contents              AS formula,
		   '[]'::text              AS measurements,
		   null                    AS next_measurement_low,
		   null                    AS next_measurement_high
	FROM calculation i
	WHERE i.contents IS NOT NULL AND i.instrument_id IN (SELECT id FROM requested_instruments)
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

		for i, _ := range tt2[idx].Measurements {
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
