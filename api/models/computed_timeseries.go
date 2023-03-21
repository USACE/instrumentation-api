package models

import (
	"encoding/json"
	"log"
	"time"

	"github.com/USACE/instrumentation-api/api/timeseries"
	"github.com/google/uuid"
	"github.com/jackc/pgtype"
	"github.com/jmoiron/sqlx"
)

type MeasurementsFromStream struct {
	Time             time.Time    `db:"time"`
	TimeseriesID     uuid.UUID    `db:"timeseries_id"`
	IsComputed       bool         `db:"is_computed"`
	Formula          string       `db:"formula"`
	MeasurementsJSON pgtype.JSONB `db:"measurements_json"`
}

type MeasurementsFilter struct {
	TimeseriesID      uuid.UUID        `db:"timeseries_id"`
	InstrumentID      uuid.UUID        `db:"instrument_id"`
	InstrumentGroupID uuid.UUID        `db:"instrument_group_id"`
	InstrumentIDs     pgtype.UUIDArray `db:"instrument_ids"`
	After             time.Time        `db:"after"`
	Before            time.Time        `db:"before"`
}

func QueryTimeseriesMeasurementsRows(db *sqlx.DB, f *MeasurementsFilter) (*sqlx.Rows, error) {
	sql := `
	WITH required_timeseries AS (
		SELECT
			dependency_timeseries_id AS id,
			timeseries_id AS computed_timeseries_id,
			instrument_id,
			parsed_variable AS variable,
			true AS is_computed
		FROM v_timeseries_dependency
		WHERE
			timeseries_id = $1
			OR instrument_id = $2
			OR instrument_id IN (
				SELECT instrument_id
				FROM instrument_group_instruments
				WHERE instrument_group_id = $3
			)
			-- OR instrument_id = ANY(:instrument_ids)
		UNION
		SELECT
			id,
			NULL AS computed_timeseries_id,
			instrument_id,
			NULL AS variable,
			false AS is_computed
		FROM v_timeseries_stored
		WHERE
			id = $1
			OR instrument_id = $2
			OR instrument_id IN (
				SELECT instrument_id
				FROM instrument_group_instruments
				WHERE instrument_group_id = $3
			)
			-- OR instrument_id = ANY(:instrument_ids)
	),
	next_low AS (
		SELECT
			timeseries_id,
			MAX(time) AS time
		FROM timeseries_measurement
		WHERE
			timeseries_id IN (SELECT id FROM required_timeseries)
			AND time < $4
		GROUP BY timeseries_id
	),
	next_high AS (
		SELECT
			timeseries_id,
			MIN(time) AS time
		FROM timeseries_measurement
		WHERE
			timeseries_id IN (SELECT id FROM required_timeseries)
			AND time > $5
		GROUP BY timeseries_id
	)
	
	SELECT
		tm.time AS time,
		COALESCE(rt.computed_timeseries_id, rt.id) AS timeseries_id,
		rt.is_computed AS is_computed,
		cc.contents AS formula,
		jsonb_object_agg(
			COALESCE(rt.variable, 'value'), tm.value
		) || (
			CASE rt.is_computed
			  WHEN NOT true THEN
				jsonb_build_object(
					'masked', COALESCE(tn.masked, false),
					'validated', COALESCE(tn.validated, false),
					'annotation', COALESCE(tn.annotation, '')
				)
			  ELSE '{}'::jsonb
		END) AS measurements_json
	FROM required_timeseries rt
	
	INNER JOIN timeseries_measurement tm
		ON rt.id = tm.timeseries_id
	INNER JOIN instrument i
		ON i.id = rt.instrument_id
	LEFT JOIN calculation cc
		ON rt.computed_timeseries_id = cc.timeseries_id
	LEFT JOIN next_low nl
		ON nl.timeseries_id = rt.id
	LEFT JOIN next_high nh
		ON nh.timeseries_id = rt.id
	LEFT JOIN timeseries_notes tn
		ON tm.timeseries_id = tn.timeseries_id
		AND tm.time = tn.time
	WHERE
		(nl.time IS NULL OR tm.time >= nl.time)
		AND (nh.time IS NULL OR tm.time <= nh.time)
	GROUP BY
		rt.computed_timeseries_id,
		rt.id,
		rt.is_computed,
		tm.time,
		tn.masked,
		tn.validated,
		tn.annotation,
		cc.contents
	ORDER BY tm.time ASC, rt.is_computed DESC
	`

	rows, err := db.Queryx(sql, &f.TimeseriesID, &f.InstrumentID, &f.InstrumentGroupID, &f.After, &f.Before)
	if err != nil {
		return nil, err
	}

	return rows, nil
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
