package models

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/USACE/instrumentation-api/api/dbutils"
	"github.com/USACE/instrumentation-api/api/messages"
	"github.com/USACE/instrumentation-api/api/timeseries"
	"github.com/google/uuid"
	"github.com/jackc/pgtype"
	"github.com/jmoiron/sqlx"
)

// MeasurementsFromRow binds to each row returned from QueryTimeseriesMeasurementsRows
type MeasurementsFromRow struct {
	Time             time.Time    `db:"time"`
	TimeseriesID     uuid.UUID    `db:"timeseries_id"`
	InstrumentID     uuid.UUID    `db:"instrument_id"`
	IsComputed       bool         `db:"is_computed"`
	Formula          string       `db:"formula"`
	MeasurementsJSON pgtype.JSONB `db:"measurements_json"`
}

type MeasurementsResponseCollection []MeasurementsResponse

// MeasurementsResponse basic type for responses, can be transformed with methods based on MeasurementsFilter
type MeasurementsResponse struct {
	TimeseriesID uuid.UUID `json:"timeseries_id"`
	InstrumentID uuid.UUID `json:"instrument_id"`
	Measurement
	TimeseriesNote
}

type TimeseriesNote struct {
	Masked     bool   `json:"masked,omitempty"`
	Validated  bool   `json:"validated,omitempty"`
	Annotation string `json:"annotation,omitempty"`
}

// MeasurementsFilter for conveniently passsing SQL query paramters to functions
type MeasurementsFilter struct {
	TimeseriesID       *uuid.UUID  `db:"timeseries_id"`
	InstrumentID       *uuid.UUID  `db:"instrument_id"`
	InstrumentGroupID  *uuid.UUID  `db:"instrument_group_id"`
	InstrumentIDs      []uuid.UUID `db:"instrument_ids"`
	After              time.Time   `db:"after"`
	Before             time.Time   `db:"before"`
	TemporalResolution int         `db:"temporal_resolution"`
}

func (mrc *MeasurementsResponseCollection) GroupByInstrument() (map[uuid.UUID][]timeseries.MeasurementCollectionLean, error) {
	defer dbutils.Timer()()
	if len(*mrc) == 0 {
		return make(map[uuid.UUID][]timeseries.MeasurementCollectionLean), fmt.Errorf(messages.NotFound)
	}

	tmp := make(map[uuid.UUID]map[uuid.UUID][]timeseries.MeasurementLean)

	for _, t := range *mrc {
		if _, hasInstrument := tmp[t.InstrumentID]; !hasInstrument {
			tmp[t.InstrumentID] = make(map[uuid.UUID][]timeseries.MeasurementLean, 0)
		}
		if _, hasTimeseries := tmp[t.InstrumentID][t.TimeseriesID]; !hasTimeseries {
			tmp[t.InstrumentID][t.TimeseriesID] = make([]timeseries.MeasurementLean, 0)
		}
		if len(t.Error) != 0 {
			continue
		}
		tmp[t.InstrumentID][t.TimeseriesID] = append(tmp[t.InstrumentID][t.TimeseriesID], timeseries.MeasurementLean{t.Time: t.Value})
	}

	res := make(map[uuid.UUID][]timeseries.MeasurementCollectionLean)

	for instrumentID := range tmp {
		res[instrumentID] = make([]timeseries.MeasurementCollectionLean, 0)

		for tsID := range tmp[instrumentID] {
			res[instrumentID] = append(res[instrumentID],
				timeseries.MeasurementCollectionLean{
					TimeseriesID: tsID,
					Items:        tmp[instrumentID][tsID],
				},
			)
		}
	}

	return res, nil
}

func (mrc *MeasurementsResponseCollection) CollectSingleTimeseries() (timeseries.MeasurementCollection, error) {
	if len(*mrc) == 0 {
		return timeseries.MeasurementCollection{}, fmt.Errorf(messages.NotFound)
	}

	mmts := make([]timeseries.Measurement, 0)

	for _, t := range *mrc {
		mmts = append(mmts, timeseries.Measurement{
			TimeseriesID: t.TimeseriesID,
			Time:         t.Time,
			Value:        t.Value,
			Masked:       t.Masked,
			Validated:    t.Validated,
			Annotation:   t.Annotation,
			Error:        t.Error,
		})
	}

	return timeseries.MeasurementCollection{TimeseriesID: mmts[0].TimeseriesID, Items: mmts}, nil
}

// QueryTimeseriesMeasurementsRows returns an aggregate of stored and computed timeseries rows to be iterated over
// instead of binding to an object like with Get or Select.
//
// This allows a buffer to only load on row into memory at a time, thus reducing the large memory spikes that come with
// calling large queries. The schema of the JSONB ('measurements_json' column) of the row returned vary depending on
// whether the is_computed column is true or false.
//
// Stored Timeseries measurements will have the following structure:
//
// {"value": float, "validated": bool, "masked": bool, "annotation": string}
//
// Computed Timeseries will look different, because they need to be fed into the arbitrary expression evaluation library.
//
// {"variable-1": float, "variable-2": float}
//
// Once the Computed Timeseries are processed however, they are bound
// to the same struct as the stored measurements: MeasurementsStreamResponse
func QueryTimeseriesMeasurementsRows(db *sqlx.DB, f *MeasurementsFilter) (*sqlx.Rows, error) {
	defer dbutils.Timer()()
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

	sql := `
	WITH required_timeseries AS (
		SELECT
			dependency_timeseries_id AS stored_timeseries_id,
		  	id AS computed_timeseries_id,
		  	instrument_id,
		  	parsed_variable AS variable,
		  	true AS is_computed
		FROM v_timeseries_dependency
		WHERE ` + filterSQL + `
		UNION
		SELECT
		  	id AS stored_timeseries_id,
		  	NULL AS computed_timeseries_id,
		  	instrument_id,
		  	NULL AS variable,
		  	false AS is_computed
		FROM v_timeseries_stored
		WHERE ` + filterSQL + `
	),
	next_low AS (
	SELECT
		timeseries_id,
		MAX(time) AS time
	FROM timeseries_measurement
	WHERE
		timeseries_id IN (SELECT stored_timeseries_id FROM required_timeseries)
		AND time < ?
	GROUP BY timeseries_id
	),
	next_high AS (
		SELECT
		  	timeseries_id,
		  	MIN(time) AS time
		FROM timeseries_measurement
		WHERE
		  	timeseries_id IN (SELECT stored_timeseries_id FROM required_timeseries)
		  	AND time > ?
		GROUP BY timeseries_id
	),
	proc AS (
		SELECT
			tab.timeseries_id,
			array_agg(row(tab.timeseries_id, tab.time, tab.value)::ts_measurement) AS measurements
		FROM timeseries_measurement tab (time,value)
		LEFT JOIN next_low nl
			ON nl.timeseries_id = tab.timeseries_id
		LEFT JOIN next_high nh
			ON nh.timeseries_id = tab.timeseries_id
		WHERE
			tab.timeseries_id IN (SELECT stored_timeseries_id FROM required_timeseries)
			AND (nl.time IS NULL OR tab.time >= nl.time)
			AND (nh.time IS NULL OR tab.time <= nh.time)
		GROUP BY tab.timeseries_id
	)
	SELECT
		ds.t AS time,
		xx.timeseries_id AS timeseries_id,
		rt.instrument_id AS instrument_id,
		rt.is_computed AS is_computed,
		COALESCE(cc.contents, '') AS formula,
		jsonb_object_agg(
		  	COALESCE(rt.variable, 'value'), ds.v
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
	CROSS JOIN LATERAL
		(VALUES (COALESCE(rt.computed_timeseries_id, rt.stored_timeseries_id))) AS xx(timeseries_id)
	INNER JOIN proc p
		ON p.timeseries_id = rt.stored_timeseries_id
	INNER JOIN lttb(?, p.measurements) ds
		ON rt.stored_timeseries_id = ds.id
	INNER JOIN instrument i
		ON i.id = rt.instrument_id
	LEFT JOIN calculation cc
		ON rt.computed_timeseries_id = cc.timeseries_id
	LEFT JOIN timeseries_notes tn
		ON rt.stored_timeseries_id = tn.timeseries_id
		AND ds.t = tn.time
	GROUP BY
		ds.t,
		xx.timeseries_id,
		rt.instrument_id,
		rt.is_computed,
		cc.contents,
		tn.masked,
		tn.validated,
		tn.annotation
	ORDER BY ds.t ASC, rt.is_computed DESC
	`

	query, args, err := sqlx.In(sql, filterArg, filterArg, f.After, f.Before, f.TemporalResolution)
	if err != nil {
		return nil, err
	}
	query = db.Rebind(query)
	rows, err := db.Queryx(query, args...)
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
