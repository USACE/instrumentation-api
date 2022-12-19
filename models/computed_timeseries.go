package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"sort"
	"strconv"
	"time"

	"github.com/Knetic/govaluate"
	"github.com/USACE/instrumentation-api/timeseries"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

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

// AggregateCarryForward creates an array of Measurments for a timeserires given an aggregate array of times
// This assumes that the provided aggregate set of times does not have any repeating times
// This algorithm will remember the last exisiting Measurement value in the Timeseries
func (ts Timeseries) AggregateCarryForward(w timeseries.TimeWindow, allTimes []time.Time) (Timeseries, error) {
	// Array to add additional "carry forward" measurements to
	aggregateMeasurements := make([]Measurement, 0)

	// Array of existing measurements
	a := make([]Measurement, 0)
	if ts.NextMeasurementLow != nil {
		a = append(a, *ts.NextMeasurementLow)
	}
	a = append(a, ts.Measurements...)
	if ts.NextMeasurementHigh != nil {
		a = append(a, *ts.NextMeasurementHigh)
	}

	if len(a) == 0 {
		return Timeseries{
			TimeseriesInfo: ts.TimeseriesInfo,
			Measurements:   make([]Measurement, 0),
			TimeWindow:     ts.TimeWindow,
		}, nil
	}

	tStart, tEnd, wkIdx, lastIdx := w.After, w.Before, 0, len(a)-1
	remember := a[0].Value

	for _, tm := range allTimes {
		// Time out of range, cannot compute
		if tm.Before(a[0].Time) || tm.Before(tStart) || tm.After(tEnd) || wkIdx == lastIdx {
			continue
		}

		// Time allTimes buffer position caught up with working array index, add measurement and advance working index
		if tm == a[wkIdx].Time {
			aggregateMeasurements = append(aggregateMeasurements, Measurement{tm, a[wkIdx].Value})
			remember = a[wkIdx].Value
			wkIdx += 1
			continue
		}
		// allTimes buffer is behind the working array index, add measurement
		aggregateMeasurements = append(aggregateMeasurements, Measurement{tm, remember})
	}

	return Timeseries{
		TimeseriesInfo:      ts.TimeseriesInfo,
		Measurements:        aggregateMeasurements,
		NextMeasurementLow:  ts.NextMeasurementLow,
		NextMeasurementHigh: ts.NextMeasurementHigh,
	}, nil
}

// Interpolate takes two arrays for the corresponding x and y of each point, returning the
// predicted value of y at the position of x using linear interpolation
func Interpolate(xs, ys []float64, x float64) (float64, error) {
	xsLen := len(xs)
	if len(ys) != xsLen {
		return 0, errors.New("xs and ys slices must be same length")
	}
	if xsLen < 2 {
		return 0, errors.New("xs length must be greater than 2")
	}
	if xs[0] > xs[1] {
		return 0, errors.New("xs array values must be increasing")
	}

	// y = y1 + ((x - x1) / (x2 - x1)) * (y2 - y1)
	return ys[0] + ((x-xs[0])/(xs[1]-xs[0]))*(ys[1]-ys[0]), nil
}

// AggregateInterpolate creates an array of Measurments for a timeseries given an aggregate array of times.
// This assumes that the provided aggregate set of times does not have any repeating times. This algorithm
// will predict Measurement values given an x postion to predict and xy values of the neighboring points
func (ts Timeseries) AggregateInterpolate(w timeseries.TimeWindow, allTimes []time.Time) (Timeseries, error) {

	interpolated := make([]Measurement, 0)

	// Array of existing measurements
	a := make([]Measurement, 0)
	if ts.NextMeasurementLow != nil {
		a = append(a, *ts.NextMeasurementLow)
	}
	a = append(a, ts.Measurements...)
	if ts.NextMeasurementHigh != nil {
		a = append(a, *ts.NextMeasurementHigh)
	}

	if len(a) == 0 {
		return Timeseries{
			TimeseriesInfo: ts.TimeseriesInfo,
			Measurements:   make([]Measurement, 0),
			TimeWindow:     ts.TimeWindow,
		}, nil
	}

	sort.Slice(a, func(i, j int) bool { return a[i].Time.Before(a[j].Time) })

	tStart, tEnd, wkIdx, lastIdx := w.After, w.Before, 0, len(a)-1

	for _, tm := range allTimes {
		// Time out of range, cannot compute
		if tm.Before(a[0].Time) || tm.Before(tStart) || tm.After(tEnd) || wkIdx > lastIdx {
			continue
		}

		// Time allTimes buffer caught up with working array index, add measurement and advance working index
		if tm == a[wkIdx].Time {
			interpolated = append(interpolated, Measurement{tm, a[wkIdx].Value})
			wkIdx += 1
			continue
		}

		// At this point, the current index i should be at least i > 0 and at most i < len(a)-1
		// Fill in interpolated values
		prevX := float64(a[wkIdx-1].Time.Unix())
		nextX := float64(a[wkIdx].Time.Unix())

		prevY := a[wkIdx-1].Value
		nextY := a[wkIdx].Value

		currentX := float64(tm.Unix())

		// allTimes buffer is behind the working array index, add interpolated measurement
		currentY, err := Interpolate([]float64{prevX, nextX}, []float64{prevY, nextY}, currentX)
		if err != nil {
			log.Println(err)
			continue
		}

		interpolated = append(interpolated, Measurement{tm, currentY})
	}

	return Timeseries{
		TimeseriesInfo:      ts.TimeseriesInfo,
		Measurements:        interpolated,
		NextMeasurementLow:  ts.NextMeasurementLow,
		NextMeasurementHigh: ts.NextMeasurementHigh,
	}, nil
}

func (ts Timeseries) RegularizeInterpolate(w timeseries.TimeWindow, d time.Duration) (Timeseries, error) {
	log.Println("Not Implemented")
	return Timeseries{}, errors.New("method not implemented")
}

func (ts *Timeseries) Calculate(variableMap map[time.Time]map[string]interface{}, interval *time.Duration) error {
	expression, err := govaluate.NewEvaluableExpression(*ts.Formula)
	if err != nil {
		return err
	}
	t, end := ts.TimeWindow.After, ts.TimeWindow.Before
	for !t.After(end) {
		if params, exists := variableMap[t]; exists {
			valStr, err := expression.Evaluate(params)
			if err != nil {
				t = t.Add(*interval)
				continue
			}
			val64, err := strconv.ParseFloat(fmt.Sprint(valStr), 64)
			if err != nil {
				t = t.Add(*interval)
				continue
			}
			ts.Measurements = append(ts.Measurements, Measurement{Time: t, Value: val64})
		}
		t = t.Add(*interval)
	}
	return nil
}

// CalculateAggregate computes aggregate, possibly irregular intervals of all timeseires
// The provided variableMap should include multiple variables for each key (time) provided
func (ts *Timeseries) CalculateAggregate(variableMap map[time.Time]map[string]interface{}) error {
	expression, err := govaluate.NewEvaluableExpression(*ts.Formula)
	if err != nil {
		return err
	}

	for k, v := range variableMap {
		valStr, err := expression.Evaluate(v)
		if err != nil {
			continue
		}

		val64, err := strconv.ParseFloat(fmt.Sprint(valStr), 64)
		if err != nil {
			continue
		}
		ts.Measurements = append(ts.Measurements, Measurement{Time: k, Value: val64})
	}

	return nil
}

// ResampleTimeseriesMeasurements provides values at a fixed, regularized interval based a provided duration
// These resampled values are interpolated from the nearest points in the aggregate calculation's curve
func (ts *Timeseries) ResampleTimeseriesMeasurements(w *timeseries.TimeWindow, d *time.Duration) (Timeseries, error) {
	resampled := make([]Measurement, 0)

	// Computed timeseries working array
	a := make([]Measurement, 0)
	if ts.NextMeasurementLow != nil {
		a = append(a, *ts.NextMeasurementLow)
	}
	a = append(a, ts.Measurements...)
	if ts.NextMeasurementHigh != nil {
		a = append(a, *ts.NextMeasurementHigh)
	}

	if len(a) == 0 {
		return Timeseries{
			TimeseriesInfo: ts.TimeseriesInfo,
			Measurements:   make([]Measurement, 0),
			TimeWindow:     ts.TimeWindow,
		}, nil
	}

	sort.Slice(a, func(i, j int) bool { return a[i].Time.Before(a[j].Time) })

	wkIdx, lastIdx := 0, len(a)-1

	// Max time between time window and measured time
	t := func() time.Time {
		if a[0].Time.After(w.After) {
			return a[0].Time
		}
		return w.After
	}()
	// Min time between time window and measured time
	tEnd := func() time.Time {
		if a[lastIdx].Time.Before(w.Before) {
			return a[lastIdx].Time
		}
		return w.Before
	}()

	for !t.After(tEnd) {
		if !t.Before(a[wkIdx].Time) {
			if wkIdx == lastIdx {
				resampled = append(resampled, Measurement{t, a[wkIdx].Value})
				t = t.Add(*d)
				continue
			}
			wkIdx += 1
		}

		// Resample using linear interpolation
		prevX := float64(a[wkIdx-1].Time.Unix())
		nextX := float64(a[wkIdx].Time.Unix())
		prevY := a[wkIdx-1].Value
		nextY := a[wkIdx].Value
		currentX := float64(t.Unix())

		currentY, err := Interpolate([]float64{prevX, nextX}, []float64{prevY, nextY}, currentX)
		if err != nil {
			log.Println(err)
			continue
		}

		resampled = append(resampled, Measurement{t, currentY})
		t = t.Add(*d)
	}

	return Timeseries{
		TimeseriesInfo:      ts.TimeseriesInfo,
		Measurements:        resampled,
		NextMeasurementLow:  ts.NextMeasurementLow,
		NextMeasurementHigh: ts.NextMeasurementHigh,
	}, nil
}

// AllTimeseriesWithMeasurements returns all stored and computed timeseries for a specified array of instrument IDs
func AllTimeseriesWithMeasurements(db *sqlx.DB, instrumentIDs []uuid.UUID, tw *timeseries.TimeWindow, interval *time.Duration) ([]Timeseries, error) {
	tt2, err := queryAllTimeseriesForInstruments(db, instrumentIDs, tw)
	if err != nil {
		return tt2, err
	}

	// Final Timeseries to be returned
	tt3 := make([]Timeseries, 0)

	// a map of all available parameters for a given time slice
	variableMap := make(map[time.Time]map[string]interface{})

	// todo: Optimization - do not need to regularize all timeseries
	// only need to regularize those that will be used as calculation dependencies
	for _, ts := range tt2 {
		tsAll, err := ts.RegularizeCarryForward(*tw, *interval)
		if err != nil {
			return make([]Timeseries, 0), err
		}

		// Add All Measurements from Regularized Timeseries to Map
		for _, m := range tsAll.Measurements {
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
		err = ts.Calculate(variableMap, interval)
		if err != nil {
			log.Printf("Error Computing Formula for Timeseries %s\n", ts.TimeseriesID)
			continue
		}
		tt3 = append(tt3, ts)
	}

	return tt3, nil
}

// ComputedTimeseriesWithMeasurements returns computed for a specified instrument ID
func ComputedTimeseriesWithMeasurements(db *sqlx.DB, timeseriesID *uuid.UUID, instrumentID *uuid.UUID, tw *timeseries.TimeWindow, interval *time.Duration) ([]Timeseries, error) {
	// Query and unmarshal JSON
	tt2, err := queryComputedTimeseries(db, timeseriesID, instrumentID, tw)
	if err != nil {
		return tt2, err
	}

	// Final Timeseries to be returned
	tt3 := make([]Timeseries, 0)

	// a map of all available parameters for a given time slice
	variableMap := make(map[time.Time]map[string]interface{})

	// Get unique set of all measurement times of timeseries dependencies for non-regularized values
	tSet := make(map[time.Time]struct{})
	allTimes := make([]time.Time, 0)

	for _, ts := range tt2 {
		for _, m := range ts.Measurements {
			tSet[m.Time] = struct{}{}
		}
	}

	// Sort times from set
	for t := range tSet {
		allTimes = append(allTimes, t)
	}

	sort.Slice(allTimes, func(i, j int) bool { return allTimes[i].Before(allTimes[j]) })

	for _, ts := range tt2 {

		// Aggregate of measurements added to timeseries
		// ts, err := ts.AggregateCarryForward(*tw, allTimes)
		// if err != nil {
		// 	return make([]Timeseries, 0), err
		// }

		// Aggregate of measurements added to timeseries
		ts, err := ts.AggregateInterpolate(*tw, allTimes)
		if err != nil {
			return make([]Timeseries, 0), err
		}

		// Add All Measurements from Timeseries to Map
		for _, m := range ts.Measurements {
			if _, exists := variableMap[m.Time]; !exists {
				variableMap[m.Time] = make(map[string]interface{})
			}
			variableMap[m.Time][ts.Variable] = m.Value
		}

		// If not a computed timeseries, do not calculate yet
		// Timeseries dependencies should be processed first
		// due to ORDER BY is_computed in SQL query
		if !ts.IsComputed {
			continue
		}

		err = ts.CalculateAggregate(variableMap)
		if err != nil {
			log.Printf("error computing formula for timeseries %s\n", ts.TimeseriesID)
			continue
		}

		if *interval != 0 {
			ts, err = ts.ResampleTimeseriesMeasurements(tw, interval)

			if err != nil {
				log.Printf("error resampling computed timeseries %s\n", ts.TimeseriesID)
				continue
			}
		}
		tt3 = append(tt3, ts)
	}

	return tt3, nil
}

// Helper function for getting timeseries by instruments
func queryAllTimeseriesForInstruments(db *sqlx.DB, instrumentIDs []uuid.UUID, tw *timeseries.TimeWindow) ([]Timeseries, error) {
	sql := `
	-- Get Timeseries and Dependencies for Calculations
	-- timeseries required based on requested instrument
	WITH requested_instruments AS (
		SELECT id
		FROM instrument
		WHERE id IN (?)
	), required_timeseries AS (
	-- 	Timeseries for Instrument
		SELECT id FROM v_timeseries_stored WHERE instrument_id IN (SELECT id FROM requested_instruments)
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
	FROM v_timeseries_computed cc
	WHERE cc.contents IS NOT NULL AND cc.instrument_id IN (SELECT id FROM requested_instruments)
	ORDER BY is_computed
	`

	query, args, err := sqlx.In(sql, instrumentIDs, tw.After, tw.Before, tw.After, tw.Before)
	if err != nil {
		return make([]Timeseries, 0), err
	}

	tt := make([]DBTimeseries, 0)
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
	return tt2, nil
}

// Helper function for getting timeseries by instruments
func queryComputedTimeseries(db *sqlx.DB, timeseriesID *uuid.UUID, instrumentID *uuid.UUID, tw *timeseries.TimeWindow) ([]Timeseries, error) {
	sql := `
	-- Get Timeseries and Dependencies for Calculations
	-- timeseries required based on requested instrument
	WITH requested_instruments AS (
		SELECT id
		FROM instrument
		WHERE id = ?
	), required_timeseries AS (
	-- Dependencies for computed timeseries
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
	-- Timeseries Dependencies
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
	FROM v_timeseries_computed cc
	WHERE cc.contents IS NOT NULL
	AND cc.instrument_id IN (SELECT id FROM requested_instruments)
	AND cc.id = ?
	ORDER BY is_computed
	`

	query, args, err := sqlx.In(sql, instrumentID, tw.After, tw.Before, tw.After, tw.Before, timeseriesID)
	if err != nil {
		return make([]Timeseries, 0), err
	}
	query = db.Rebind(query)

	tt := make([]DBTimeseries, 0)
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
	return tt2, nil
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
		SELECT id FROM v_timeseries_stored WHERE instrument_id IN (SELECT id FROM requested_instruments)
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
		   '[]'::text              AS measurements
	FROM v_timeseries_computed cc
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
