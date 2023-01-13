package models

import (
	"encoding/json"
	"log"
	"time"

	"github.com/USACE/instrumentation-api/timeseries"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// ListInstrumentsMeasurements returns all stored and computed timeseries for a specified array of instrument IDs
func ListInstrumentsMeasurements(db *sqlx.DB, instrumentIDs []uuid.UUID, tw *timeseries.TimeWindow, interval time.Duration) ([]Timeseries, error) {

	tr := make([]Timeseries, 0)

	for _, iID := range instrumentIDs {
		its, err := ListInstrumentTimeseries(db, &iID)
		if err != nil {
			return nil, err
		}

		for _, ts := range its {
			tss, err := queryTimeseriesMeasurements(db, &ts.ID, &iID, tw, ts.IsComputed)
			if err != nil {
				return tss, err
			}

			if ts.IsComputed {
				tss, err = ProcessComputedTimeseries(tss, tw, true)
				if err != nil {
					return tss, err
				}
			}

			if interval != 0 {
				resampled := make([]Timeseries, len(tss))

				for i, t := range tss {
					t, err = t.ResampleTimeseriesMeasurements(tw, interval)
					if err != nil {
						return tss, err
					}

					resampled[i] = t
				}
				tss = resampled
			}
			tr = append(tr, tss...)
		}
	}

	return tr, nil
}

// ComputedTimeseriesWithMeasurements returns computed for a specified instrument ID
func ComputedTimeseriesWithMeasurements(db *sqlx.DB, timeseriesID *uuid.UUID, instrumentID *uuid.UUID, tw *timeseries.TimeWindow, interval time.Duration) ([]Timeseries, error) {

	tr := make([]Timeseries, 0)

	// Query and unmarshal JSON
	tss, err := queryTimeseriesMeasurements(db, timeseriesID, instrumentID, tw, true)
	if err != nil {
		return tss, err
	}

	tss, err = ProcessComputedTimeseries(tss, tw, true)
	if err != nil {
		return tss, err
	}

	if interval != 0 {
		resampled := make([]Timeseries, len(tss))

		for i, t := range tss {
			t, err = t.ResampleTimeseriesMeasurements(tw, interval)
			if err != nil {
				return tss, err
			}

			resampled[i] = t
		}
		tss = resampled
	}

	tr = append(tr, tss...)

	return tr, nil
}

// Helper function for getting timeseries by instruments
func queryTimeseriesMeasurements(db *sqlx.DB, timeseriesID *uuid.UUID, instrumentID *uuid.UUID, tw *timeseries.TimeWindow, computed bool) ([]Timeseries, error) {
	reqTimseriesSql := `
	-- Regular stored timeseries
	SELECT id
	FROM v_timeseries_stored
	WHERE id = $1
	`

	if computed {
		reqTimseriesSql = `
		-- Dependencies for computed timeseries
		SELECT dependency_timeseries_id AS id
		FROM v_timeseries_dependency
		WHERE instrument_id = $2
		`
	}

	sql := `
	-- Get Timeseries and Dependencies for Calculations
	WITH required_timeseries AS (` + reqTimseriesSql + `),
	-- Next Timeseries Measurement Outside Time Window (Earlier); Needed for Calculation Interpolation
	next_low AS (
		SELECT nlm.timeseries_id AS timeseries_id, json_build_object('time', nlm.time, 'value', m1.value) AS measurement
		FROM (
			SELECT timeseries_id, MAX(time) AS time
			FROM timeseries_measurement
			WHERE timeseries_id IN (SELECT id FROM required_timeseries) AND time < $3
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
			WHERE timeseries_id IN (SELECT id FROM required_timeseries) AND time > $4
			GROUP BY timeseries_id
		) nhm
		INNER JOIN timeseries_measurement m2 ON m2.time = nhm.time AND m2.timeseries_id = nhm.timeseries_id
	),
	-- Measurements Within Time Window by timeseries_id;
	measurements AS (
		SELECT timeseries_id,
			   json_agg(json_build_object('time', time, 'value', value) ORDER BY time ASC)::text AS measurements
		FROM timeseries_measurement
		WHERE timeseries_id IN (SELECT id FROM required_timeseries) AND time >= $3 AND time <= $4
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
	INNER JOIN instrument i ON i.id = ts.instrument_id AND i.id = $2
	LEFT JOIN measurements m ON m.timeseries_id = rt.id
	LEFT JOIN next_low nl ON nl.timeseries_id = rt.id
	LEFT JOIN next_high nh ON nh.timeseries_id = rt.id
	`
	if computed {
		sql += `
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
		AND cc.instrument_id = $2
		AND cc.id = $1
		ORDER BY is_computed
		`
	}

	tt := make([]DBTimeseries, 0)
	if err := db.Select(&tt, sql, timeseriesID, instrumentID, tw.After, tw.Before); err != nil {
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
