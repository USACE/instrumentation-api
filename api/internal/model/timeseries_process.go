package model

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/Knetic/govaluate"
	"github.com/USACE/instrumentation-api/api/internal/message"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx/types"
	"github.com/tidwall/btree"
)

type ProcessTimeseriesInfo struct {
	TimeseriesID uuid.UUID `json:"timeseries_id" db:"timeseries_id"`
	InstrumentID uuid.UUID `json:"instrument_id" db:"instrument_id"`
	Variable     string    `json:"variable" db:"variable"`
	IsComputed   bool      `json:"is_computed" db:"is_computed"`
	Formula      *string   `json:"formula" db:"formula"`
}

type DBProcessTimeseries struct {
	ProcessTimeseriesInfo
	Measurements        string  `json:"measurements" db:"measurements"`
	NextMeasurementLow  *string `json:"next_measurement_low" db:"next_measurement_low"`
	NextMeasurementHigh *string `json:"next_measurement_high" db:"next_measurement_high"`
}

type ProcessTimeseries struct {
	ProcessTimeseriesInfo
	Measurements        []ProcessMeasurement `json:"measurements" db:"measurements"`
	NextMeasurementLow  *ProcessMeasurement  `json:"next_measurement_low" db:"next_measurement_low"`
	NextMeasurementHigh *ProcessMeasurement  `json:"next_measurement_high" db:"next_measurement_high"`
	TimeWindow          TimeWindow           `json:"time_window"`
}

type ProcessMeasurementCollection struct {
	TimeseriesID uuid.UUID            `json:"timeseries_id" db:"timeseries_id"`
	Items        []ProcessMeasurement `json:"items"`
}

type ProcessTimeseriesResponseCollection []ProcessTimeseries

type ProcessMeasurement struct {
	Time  time.Time `json:"time"`
	Value float64   `json:"value"`
	Error string    `json:"error,omitempty"`
}

func (m ProcessMeasurement) Lean() map[time.Time]float64 {
	return map[time.Time]float64{m.Time: m.Value}
}

type ProcessInclinometerTimeseries struct {
	ProcessTimeseriesInfo
	Measurements        []ProcessInclinometerMeasurement `json:"measurements" db:"measurements"`
	NextMeasurementLow  *ProcessMeasurement              `json:"next_measurement_low" db:"next_measurement_low"`
	NextMeasurementHigh *ProcessMeasurement              `json:"next_measurement_high" db:"next_measurement_high"`
	TimeWindow          TimeWindow                       `json:"time_window"`
}

type ProcessInclinometerMeasurement struct {
	Time   time.Time      `json:"time"`
	Values types.JSONText `json:"values"`
}

type ProcessInclinometerTimeseriesResponseCollection []ProcessInclinometerTimeseries

func (m ProcessInclinometerMeasurement) InclinometerLean() map[time.Time]types.JSONText {
	return map[time.Time]types.JSONText{m.Time: m.Values}
}

// explorerResponseFactory returns the explorer-specific JSON response format
func explorerInclinometerResponseFactory(tt []ProcessInclinometerTimeseries) (map[uuid.UUID][]InclinometerMeasurementCollectionLean, error) {
	response := make(map[uuid.UUID][]InclinometerMeasurementCollectionLean)

	for _, t := range tt {
		if _, hasInstrument := response[t.InstrumentID]; !hasInstrument {
			response[t.InstrumentID] = make([]InclinometerMeasurementCollectionLean, 0)
		}
		mcl := InclinometerMeasurementCollectionLean{
			TimeseriesID: t.TimeseriesID,
			Items:        make([]InclinometerMeasurementLean, len(t.Measurements)),
		}
		for idx, m := range t.Measurements {
			mcl.Items[idx] = m.InclinometerLean()
		}
		response[t.InstrumentID] = append(response[t.InstrumentID], mcl)
	}

	return response, nil
}

// ProcessMeasurementFilter for conveniently passsing SQL query paramters to functions
type ProcessMeasurementFilter struct {
	TimeseriesID      *uuid.UUID  `db:"timeseries_id"`
	InstrumentID      *uuid.UUID  `db:"instrument_id"`
	InstrumentGroupID *uuid.UUID  `db:"instrument_group_id"`
	InstrumentIDs     []uuid.UUID `db:"instrument_ids"`
	After             time.Time   `db:"after"`
	Before            time.Time   `db:"before"`
}

// BTreeNode represents node for btree used for computing timeseries
type BTreeNode struct {
	Key   time.Time
	Value map[string]interface{}
}

func (mrc *ProcessTimeseriesResponseCollection) GroupByInstrument(threshold int) (map[uuid.UUID][]MeasurementCollectionLean, error) {
	if len(*mrc) == 0 {
		return make(map[uuid.UUID][]MeasurementCollectionLean), fmt.Errorf(message.NotFound)
	}

	tmp := make(map[uuid.UUID]map[uuid.UUID][]MeasurementLean)

	for _, t := range *mrc {
		if _, hasInstrument := tmp[t.InstrumentID]; !hasInstrument {
			tmp[t.InstrumentID] = make(map[uuid.UUID][]MeasurementLean, 0)
		}
		if _, hasTimeseries := tmp[t.InstrumentID][t.TimeseriesID]; !hasTimeseries {
			tmp[t.InstrumentID][t.TimeseriesID] = make([]MeasurementLean, 0)
		}
		for _, m := range t.Measurements {
			tmp[t.InstrumentID][t.TimeseriesID] = append(tmp[t.InstrumentID][t.TimeseriesID], MeasurementLean{m.Time: m.Value})
		}
	}

	res := make(map[uuid.UUID][]MeasurementCollectionLean)

	for instrumentID := range tmp {
		res[instrumentID] = make([]MeasurementCollectionLean, 0)

		for tsID := range tmp[instrumentID] {
			res[instrumentID] = append(res[instrumentID],
				MeasurementCollectionLean{
					TimeseriesID: tsID,
					Items:        LTTB(tmp[instrumentID][tsID], threshold),
				},
			)
		}
	}

	return res, nil
}

func (mrc *ProcessInclinometerTimeseriesResponseCollection) GroupByInstrument() (map[uuid.UUID][]InclinometerMeasurementCollectionLean, error) {
	if len(*mrc) == 0 {
		return make(map[uuid.UUID][]InclinometerMeasurementCollectionLean), fmt.Errorf(message.NotFound)
	}

	res := make(map[uuid.UUID][]InclinometerMeasurementCollectionLean)

	for _, t := range *mrc {
		if _, hasInstrument := res[t.InstrumentID]; !hasInstrument {
			res[t.InstrumentID] = make([]InclinometerMeasurementCollectionLean, 0)
		}
		mcl := InclinometerMeasurementCollectionLean{
			TimeseriesID: t.TimeseriesID,
			Items:        make([]InclinometerMeasurementLean, len(t.Measurements)),
		}
		for idx, m := range t.Measurements {
			mcl.Items[idx] = m.InclinometerLean()
		}
		res[t.InstrumentID] = append(res[t.InstrumentID], mcl)
	}
	return res, nil
}

func (mrc *ProcessTimeseriesResponseCollection) CollectSingleTimeseries(threshold int, tsID uuid.UUID) (MeasurementCollection, error) {
	if len(*mrc) == 0 {
		return MeasurementCollection{}, fmt.Errorf(message.NotFound)
	}

	for _, t := range *mrc {
		if t.TimeseriesID == tsID {
			mmts := make([]Measurement, len(t.Measurements))
			for i, m := range t.Measurements {
				mmts[i] = Measurement{
					TimeseriesID: t.TimeseriesID,
					Time:         m.Time,
					Value:        m.Value,
					Error:        m.Error,
				}
			}
			return MeasurementCollection{TimeseriesID: t.TimeseriesID, Items: LTTB(mmts, threshold)}, nil
		}
	}

	return MeasurementCollection{}, fmt.Errorf("requested timeseries does not match any in the result")
}

// SelectMeasurements returns measurements for the timeseries specified in the filter
func (q *Queries) SelectMeasurements(ctx context.Context, f ProcessMeasurementFilter) (ProcessTimeseriesResponseCollection, error) {
	tss, err := queryTimeseriesMeasurements(ctx, q, f)
	if err != nil {
		return tss, err
	}
	tss, err = processLOCF(tss)
	if err != nil {
		return tss, err
	}
	return tss, nil
}

// SelectInclinometerMeasurements returns inclinometer measurements for the instruments specified in the filter
func (q *Queries) SelectInclinometerMeasurements(ctx context.Context, f ProcessMeasurementFilter) (ProcessInclinometerTimeseriesResponseCollection, error) {
	tss, err := queryInclinometerTimeseriesMeasurements(ctx, q, f)
	if err != nil {
		return tss, err
	}
	return tss, nil
}

// collectAggregate creates a btree of all sorted times (key) and measurements (value; as variable map) from an array of Timeseries
func collectAggregate(tss *ProcessTimeseriesResponseCollection) *btree.BTreeG[BTreeNode] {
	// Get unique set of all measurement times of timeseries dependencies for non-regularized values
	btm := btree.NewBTreeG(func(a, b BTreeNode) bool { return a.Key.Before(b.Key) })
	for _, ts := range *tss {
		if ts.NextMeasurementLow != nil {
			if item, exists := btm.Get(BTreeNode{Key: ts.NextMeasurementLow.Time}); !exists {
				btm.Set(BTreeNode{Key: ts.NextMeasurementLow.Time, Value: map[string]interface{}{ts.Variable: ts.NextMeasurementLow.Value}})
			} else {
				item.Value[ts.Variable] = ts.NextMeasurementLow.Value
				btm.Set(item)
			}
		}
		for _, m := range ts.Measurements {
			if item, exists := btm.Get(BTreeNode{Key: m.Time}); !exists {
				btm.Set(BTreeNode{Key: m.Time, Value: map[string]interface{}{ts.Variable: m.Value}})
			} else {
				item.Value[ts.Variable] = m.Value
				btm.Set(item)
			}
		}
		if ts.NextMeasurementHigh != nil {
			if item, exists := btm.Get(BTreeNode{Key: ts.NextMeasurementHigh.Time}); !exists {
				btm.Set(BTreeNode{Key: ts.NextMeasurementHigh.Time, Value: map[string]interface{}{ts.Variable: ts.NextMeasurementHigh.Value}})
			} else {
				item.Value[ts.Variable] = ts.NextMeasurementHigh.Value
				btm.Set(item)
			}
		}
	}
	return btm
}

// processLOCF calculates computed timeseries using "Last-Observation-Carried-Forward" algorithm
func processLOCF(tss ProcessTimeseriesResponseCollection) (ProcessTimeseriesResponseCollection, error) {
	tssFinal := make(ProcessTimeseriesResponseCollection, 0)
	var variableMap *btree.BTreeG[BTreeNode]
	// Check if any computed timeseries present, collect aggregates used for calculations if so
	for _, ts := range tss {
		if ts.IsComputed {
			variableMap = collectAggregate(&tss)
			break
		}
	}
	// Add any stored timeseries to the result
	// Do calculations for computed timeseries and add to result
	for _, ts := range tss {
		// Array of existing measurements
		a1 := make([]ProcessMeasurement, 0)
		if ts.NextMeasurementLow != nil {
			a1 = append(a1, *ts.NextMeasurementLow)
		}
		a1 = append(a1, ts.Measurements...)
		if ts.NextMeasurementHigh != nil {
			a1 = append(a1, *ts.NextMeasurementHigh)
		}

		// Could do some additional checks before adding, like if the
		// timeseries was actual requested or if it was just in the result as a
		// dependency of the computed timeseries, just returning them all for now
		if !ts.IsComputed {
			tssFinal = append(tssFinal, ProcessTimeseries{
				ProcessTimeseriesInfo: ts.ProcessTimeseriesInfo,
				Measurements:          a1,
				TimeWindow:            ts.TimeWindow,
			})
			continue
		}

		// By now, all of the stored timeseries have been processed;
		// the query is ordered in a way that priortizes stored timeseries
		expr, err := govaluate.NewEvaluableExpression(*ts.Formula)
		if err != nil {
			continue
		}

		// Do calculations
		remember := make(map[string]interface{})
		a2 := make([]ProcessMeasurement, 0)

		it := variableMap.Iter()
		for it.Next() {
			item := it.Item()

			// fill in any missing gaps of data
			for k, v := range remember {
				if _, exists := item.Value[k]; !exists {
					item.Value[k] = v
				}
			}
			// Add/Update the most recent values
			for k, v := range item.Value {
				remember[k] = v
			}

			val, err := expr.Evaluate(item.Value)
			if err != nil {
				continue
			}
			val64, err := strconv.ParseFloat(fmt.Sprint(val), 64)
			if err != nil {
				continue
			}

			a2 = append(a2, ProcessMeasurement{Time: item.Key, Value: val64})
		}
		it.Release()

		tssFinal = append(tssFinal, ProcessTimeseries{
			ProcessTimeseriesInfo: ts.ProcessTimeseriesInfo,
			Measurements:          a2,
			TimeWindow:            ts.TimeWindow,
		})
	}

	return tssFinal, nil
}

// SelectTimeseriesMeasurements selects stored measurements and dependencies for computed measurements
func queryTimeseriesMeasurements(ctx context.Context, q *Queries, f ProcessMeasurementFilter) (ProcessTimeseriesResponseCollection, error) {
	var filterSQL string
	var filterArg interface{}
	// short circuiting before executing SQL query greatly improves query perfomance,
	// rather than adding all parameters to the query with logical OR
	if f.TimeseriesID != nil {
		filterSQL = `id = ?`
		filterArg = f.TimeseriesID
	} else if f.InstrumentID != nil {
		filterSQL = `instrument_id = ?`
		filterArg = f.InstrumentID
	} else if f.InstrumentGroupID != nil {
		filterSQL = `
		instrument_id IN (
			SELECT instrument_id
			FROM instrument_group_instruments
			WHERE instrument_group_id = ?
		)`
		filterArg = f.InstrumentGroupID
	} else if len(f.InstrumentIDs) > 0 {
		filterSQL = `instrument_id IN (?)`
		filterArg = f.InstrumentIDs
	} else {
		return nil, fmt.Errorf("must supply valid filter for timeseries_measurement query")
	}
	listTimeseriesMeasurments := `
	WITH required_timeseries AS (
		(
			SELECT id
			FROM v_timeseries_stored
			WHERE ` + filterSQL + `
		)
		UNION ALL
		(
			SELECT dependency_timeseries_id AS id
			FROM v_timeseries_dependency
			WHERE ` + filterSQL + `
		)
	),
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
	next_high AS (
		SELECT nhm.timeseries_id AS timeseries_id, json_build_object('time', nhm.time, 'value', m2.value) AS measurement
		FROM (
			SELECT timeseries_id, MIN(time) AS time
			FROM timeseries_measurement
			WHERE timeseries_id IN (SELECT id FROM required_timeseries) AND time > ?
			GROUP BY timeseries_id
		) nhm
		INNER JOIN timeseries_measurement m2 ON m2.time = nhm.time AND m2.timeseries_id = nhm.timeseries_id
	)
	(
		SELECT rt.id                          AS timeseries_id,
			   ts.instrument_id               AS instrument_id,
			   i.slug || '.' || ts.slug       AS variable,
			   false                          AS is_computed,
			   null                           AS formula,
			   COALESCE((
					SELECT json_agg(json_build_object('time', time, 'value', value) ORDER BY time ASC)::text
					FROM timeseries_measurement
					WHERE timeseries_id = rt.id AND time >= ? AND time <= ?
			   ), '[]')			  AS measurements,
			   nl.measurement::text           AS next_measurement_low,
			   nh.measurement::text           AS next_measurement_high
		FROM required_timeseries rt
		INNER JOIN timeseries ts ON ts.id = rt.id
		INNER JOIN instrument i ON i.id = ts.instrument_id
		LEFT JOIN next_low nl ON nl.timeseries_id = rt.id
		LEFT JOIN next_high nh ON nh.timeseries_id = rt.id
	)
	UNION ALL
	(
		SELECT id                		  	  AS timeseries_id,
			   instrument_id        		  AS instrument_id,
			   slug			        	  AS variable,
			   true                    		  AS is_computed,
			   contents             		  AS formula,
			   '[]'::text              		  AS measurements,
			   null                    		  AS next_measurement_low,
			   null                    		  AS next_measurement_high
		FROM v_timeseries_computed
		WHERE ` + filterSQL + ` AND contents IS NOT NULL
	)
	ORDER BY is_computed
	`
	query, args, err := sqlIn(listTimeseriesMeasurments, filterArg, filterArg, f.After, f.Before, f.After, f.Before, filterArg)
	if err != nil {
		return nil, err
	}
	query = q.db.Rebind(query)
	tt := make([]DBProcessTimeseries, 0)
	if err := q.db.SelectContext(ctx, &tt, query, args...); err != nil {
		return make(ProcessTimeseriesResponseCollection, 0), err
	}
	tt2 := make(ProcessTimeseriesResponseCollection, len(tt))
	for idx, t := range tt {
		tt2[idx] = ProcessTimeseries{
			ProcessTimeseriesInfo: t.ProcessTimeseriesInfo,
			Measurements:          make([]ProcessMeasurement, 0),
			TimeWindow:            TimeWindow{Start: f.After, End: f.Before},
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
func queryInclinometerTimeseriesMeasurements(ctx context.Context, q *Queries, f ProcessMeasurementFilter) ([]ProcessInclinometerTimeseries, error) {
	tt := make([]DBProcessTimeseries, 0)
	listInclinometerTimeseriesMeasurements := `
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

	query, args, err := sqlIn(listInclinometerTimeseriesMeasurements, f.InstrumentIDs, f.After, f.Before)
	if err != nil {
		return make([]ProcessInclinometerTimeseries, 0), err
	}
	query = q.db.Rebind(query)
	if err := q.db.Select(&tt, query, args...); err != nil {
		return make([]ProcessInclinometerTimeseries, 0), err
	}

	// Unmarshal JSON Strings
	tt2 := make([]ProcessInclinometerTimeseries, len(tt))
	for idx, t := range tt {
		tt2[idx] = ProcessInclinometerTimeseries{
			ProcessTimeseriesInfo: t.ProcessTimeseriesInfo,
			Measurements:          make([]ProcessInclinometerMeasurement, 0),
			TimeWindow:            TimeWindow{Start: f.After, End: f.Before},
		}
		cm, err := q.GetTimeseriesConstantMeasurement(ctx, t.TimeseriesID, "inclinometer-constant")
		if err != nil {
			return nil, err
		}
		if err := json.Unmarshal([]byte(t.Measurements), &tt2[idx].Measurements); err != nil {
			log.Println(err)
		}
		for i := range tt2[idx].Measurements {
			values, err := q.ListInclinometerMeasurementValues(ctx, t.TimeseriesID, tt2[idx].Measurements[i].Time, cm.Value)
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
