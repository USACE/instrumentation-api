package models

import (
	"errors"
	"fmt"
	"log"
	"sort"
	"strconv"
	"time"

	"github.com/Knetic/govaluate"
	"github.com/USACE/instrumentation-api/timeseries"
)

//------------------------ calculate computed timeseries measurements ------------------------//

// Calculate performs calculations on a map of variables and values at given timestamps
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

//---------------------------------- create timestamp set ----------------------------------//

// createAllTimesSet creates a unique set of all times from an array of Timeseries
func createAllTimesSet(tss []Timeseries) []time.Time {
	// Get unique set of all measurement times of timeseries dependencies for non-regularized values
	tSet := make(map[time.Time]struct{})
	allTimes := make([]time.Time, 0)

	for _, ts := range tss {
		if ts.NextMeasurementLow != nil {
			tSet[ts.NextMeasurementLow.Time] = struct{}{}
		}
		if ts.NextMeasurementHigh != nil {
			tSet[ts.NextMeasurementHigh.Time] = struct{}{}
		}
		for _, m := range ts.Measurements {
			tSet[m.Time] = struct{}{}
		}
	}

	// Sort times from set
	for t := range tSet {
		allTimes = append(allTimes, t)
	}

	sort.Slice(allTimes, func(i, j int) bool { return allTimes[i].Before(allTimes[j]) })

	return allTimes
}

//---------------------------------- process timeseries ----------------------------------//

// 1. Need to identify when to use carry forward algorithm vs linear interpolation
// 2. Need to figure out resampling:
// 		- automatic resampling based on size / range queried?
//		- user provided?

// ProcessTimeseries performs calculations on a TimeseriesCollection
func ProcessComputedTimeseries(tss []Timeseries, tw *timeseries.TimeWindow, interp bool) ([]Timeseries, error) {

	tr := make([]Timeseries, 0)

	// a map of all available parameters for a given time slice
	variableMap := make(map[time.Time]map[string]interface{})
	ats := createAllTimesSet(tss)

	for _, ts := range tss {
		var tsAll Timeseries
		var err error

		// Stub out measurements for "missing" times, either use interpolated or last known value
		if interp {
			tsAll, err = ts.AggregateInterpolate(tw, ats)
			if err != nil {
				return make([]Timeseries, 0), err
			}
		} else {
			tsAll, err = ts.AggregateCarryForward(tw, ats)
			if err != nil {
				return make([]Timeseries, 0), err
			}
		}

		// Add All Measurements from Regularized Timeseries to Map
		for _, m := range tsAll.Measurements {
			if _, exists := variableMap[m.Time]; !exists {
				variableMap[m.Time] = make(map[string]interface{})
			}
			variableMap[m.Time][ts.Variable] = m.Value
		}

		// If not a computed timeseries do not add to returned timeseries
		if !ts.IsComputed {
			continue
		}

		// Calculations
		// It is known that all stored timeseries have been added to the Map and calculations
		// can now be run because calculated timeseries (identified by .IsComputed)
		// are returned from the database last in the query using ORDER BY is_computed
		err = ts.CalculateAggregate(variableMap)
		if err != nil {
			log.Printf("Error Computing Formula for Timeseries %s\n", ts.TimeseriesID)
			continue
		}

		// Sort by oldest to newest
		sort.Slice(ts.Measurements, func(i, j int) bool { return ts.Measurements[i].Time.Before(ts.Measurements[j].Time) })

		tr = append(tr, ts)
	}

	return tr, nil
}

//------------------------------ carry-forward algorithm ------------------------------//

// AggregateCarryForward creates an array of Measurments for a timeserires given an aggregate array of times
// This assumes that the provided aggregate set of times does not have any repeating times
// This algorithm will remember the last exisiting Measurement value in the Timeseries
func (ts Timeseries) AggregateCarryForward(w *timeseries.TimeWindow, allTimes []time.Time) (Timeseries, error) {
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

	wkIdx, lastIdx := 0, len(a)-1
	remember := a[0].Value

	for _, tm := range allTimes {
		// Time out of range, cannot compute
		if tm.Before(a[0].Time) || tm.After(w.Before) {
			continue
		}

		// Time allTimes buffer position caught up with working array index, add measurement and advance working index
		if wkIdx <= lastIdx && tm == a[wkIdx].Time {
			remember = a[wkIdx].Value
			wkIdx += 1
		}
		// allTimes buffer is behind the working array index, add measurement
		aggregateMeasurements = append(aggregateMeasurements, Measurement{tm, remember})
	}

	return Timeseries{
		TimeseriesInfo:      ts.TimeseriesInfo,
		Measurements:        aggregateMeasurements,
		NextMeasurementLow:  ts.NextMeasurementLow,
		NextMeasurementHigh: ts.NextMeasurementHigh,
		TimeWindow:          ts.TimeWindow,
	}, nil
}

//------------------------------ interpolation algorithm ------------------------------//

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
func (ts Timeseries) AggregateInterpolate(w *timeseries.TimeWindow, allTimes []time.Time) (Timeseries, error) {

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

	wkIdx, lastIdx := 0, len(a)-1

	for _, tm := range allTimes {
		// Time out of range, cannot compute
		if tm.Before(a[0].Time) || tm.After(a[lastIdx].Time) || wkIdx > lastIdx {
			continue
		}

		// Time allTimes buffer caught up with working array index, add measurement and advance working index
		if tm == a[wkIdx].Time {
			interpolated = append(interpolated, Measurement{tm, a[wkIdx].Value})
			wkIdx += 1
			continue
		}

		if wkIdx-1 < 0 {
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
			return Timeseries{
				TimeseriesInfo: ts.TimeseriesInfo,
				Measurements:   make([]Measurement, 0),
				TimeWindow:     ts.TimeWindow,
			}, err
		}

		interpolated = append(interpolated, Measurement{tm, currentY})
	}

	return Timeseries{
		TimeseriesInfo:      ts.TimeseriesInfo,
		Measurements:        interpolated,
		NextMeasurementLow:  ts.NextMeasurementLow,
		NextMeasurementHigh: ts.NextMeasurementHigh,
		TimeWindow:          ts.TimeWindow,
	}, nil
}

// ------------------------------ resample ----------------------------------//

// ResampleTimeseriesMeasurements provides values at a fixed, regularized interval based a provided duration
// The method of resampling (interpolated or carried forward) is specified via the `interp` function parameter
func (ts *Timeseries) ResampleTimeseriesMeasurements(w *timeseries.TimeWindow, d time.Duration, interp bool) (Timeseries, error) {
	if interp {
		return ts.ResampleInterpolate(w, d)
	} else {
		return ts.ResampleCarryForward(w, d)
	}
}

//------------------------------ interpolated resampling ------------------------------//

// ResampleInterpolate resamples using interpolation
func (ts *Timeseries) ResampleInterpolate(w *timeseries.TimeWindow, d time.Duration) (Timeseries, error) {

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
			TimeseriesInfo:      ts.TimeseriesInfo,
			Measurements:        make([]Measurement, 0),
			TimeWindow:          ts.TimeWindow,
			NextMeasurementLow:  ts.NextMeasurementLow,
			NextMeasurementHigh: ts.NextMeasurementHigh,
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
				t = t.Add(d)
				continue
			}
			wkIdx += 1
		}

		if wkIdx-1 < 0 {
			continue
		}

		// Resample using linear interpolation
		prevX := float64(a[wkIdx-1].Time.Unix())
		nextX := float64(a[wkIdx].Time.Unix())
		prevY := a[wkIdx-1].Value
		nextY := a[wkIdx].Value
		currentX := float64(t.Unix())

		currentY, err := Interpolate([]float64{prevX, nextX}, []float64{prevY, nextY}, currentX)
		if err != nil {
			return Timeseries{
				TimeseriesInfo:      ts.TimeseriesInfo,
				Measurements:        make([]Measurement, 0),
				TimeWindow:          ts.TimeWindow,
				NextMeasurementLow:  ts.NextMeasurementLow,
				NextMeasurementHigh: ts.NextMeasurementHigh,
			}, err
		}

		resampled = append(resampled, Measurement{t, currentY})
		t = t.Add(d)
	}

	return Timeseries{
		TimeseriesInfo:      ts.TimeseriesInfo,
		Measurements:        resampled,
		TimeWindow:          ts.TimeWindow,
		NextMeasurementLow:  ts.NextMeasurementLow,
		NextMeasurementHigh: ts.NextMeasurementHigh,
	}, nil
}

//------------------------------ carry-forward resampling ------------------------------//

// ResampleCarryForward resamples using the last known value
func (ts *Timeseries) ResampleCarryForward(w *timeseries.TimeWindow, d time.Duration) (Timeseries, error) {

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
			TimeseriesInfo:      ts.TimeseriesInfo,
			Measurements:        make([]Measurement, 0),
			TimeWindow:          ts.TimeWindow,
			NextMeasurementLow:  ts.NextMeasurementLow,
			NextMeasurementHigh: ts.NextMeasurementHigh,
		}, nil
	}

	sort.Slice(a, func(i, j int) bool { return a[i].Time.Before(a[j].Time) })

	wkIdx, lastIdx := 0, len(a)-1
	remember := a[0].Value

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
				remember = a[wkIdx].Value
				resampled = append(resampled, Measurement{t, remember})
				t = t.Add(d)
				continue
			}
			wkIdx += 1
		}

		if wkIdx-1 < 0 {
			continue
		}

		resampled = append(resampled, Measurement{t, remember})
		t = t.Add(d)
	}

	return Timeseries{
		TimeseriesInfo:      ts.TimeseriesInfo,
		Measurements:        resampled,
		TimeWindow:          ts.TimeWindow,
		NextMeasurementLow:  ts.NextMeasurementLow,
		NextMeasurementHigh: ts.NextMeasurementHigh,
	}, nil
}
