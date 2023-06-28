package models

import (
	"encoding/json"
	"time"

	ts "github.com/USACE/instrumentation-api/api/timeseries"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// TimeseriesMeasurementCollectionCollection is a collection of timeseries measurement collections
// i.e an array of structs, each containing timeseries measurements not necessarily from the same time series
type TimeseriesMeasurementCollectionCollection struct {
	Items []ts.MeasurementCollection
}

// TimeseriesIDs returns a slice of all timeseries IDs contained in the MeasurementCollectionCollection
func (cc *TimeseriesMeasurementCollectionCollection) TimeseriesIDs() []uuid.UUID {

	dd := make([]uuid.UUID, 0)
	for _, item := range cc.Items {
		dd = append(dd, item.TimeseriesID)
	}
	return dd
}

// UnmarshalJSON implements UnmarshalJSON interface
func (cc *TimeseriesMeasurementCollectionCollection) UnmarshalJSON(b []byte) error {
	switch JSONType(b) {
	case "ARRAY":
		if err := json.Unmarshal(b, &cc.Items); err != nil {
			return err
		}
	case "OBJECT":
		var mc ts.MeasurementCollection
		if err := json.Unmarshal(b, &mc); err != nil {
			return err
		}
		cc.Items = []ts.MeasurementCollection{mc}
	default:
		cc.Items = make([]ts.MeasurementCollection, 0)
	}
	return nil
}

// ListTimeseriesMeasurements returns a stored timeseries with slice of timeseries measurements populated
func ListTimeseriesMeasurements(db *sqlx.DB, tsID *uuid.UUID, tw *ts.TimeWindow, threshold int) (*ts.MeasurementCollection, error) {
	sql := `
	SELECT M.timeseries_id,
		   M.time,
		   M.value,
		   COALESCE(N.masked, 'false') AS masked,
		   COALESCE(N.validated, 'false') AS validated,
		   COALESCE(N.annotation, '') AS annotation
	FROM timeseries_measurement M
	LEFT JOIN timeseries_notes N ON M.timeseries_id = N.timeseries_id AND M.time = N.time
	INNER JOIN timeseries T ON T.id = M.timeseries_id
	WHERE T.id = $1 AND M.time > $2 AND M.time < $3 ORDER BY M.time ASC
	`

	// Get Timeseries Measurements
	items := make([]ts.Measurement, 0)
	if err := db.Select(
		&items,
		sql,
		tsID, tw.After, tw.Before,
	); err != nil {
		return nil, err
	}

	return &ts.MeasurementCollection{TimeseriesID: *tsID, Items: ts.LTTB(items, threshold)}, nil
}

// DeleteTimeserieMeasurements deletes a timeseries Measurement
func DeleteTimeserieMeasurements(db *sqlx.DB, id *uuid.UUID, time time.Time) error {
	if _, err := db.Exec("DELETE FROM timeseries_measurement WHERE timeseries_id = $1 and time = $2", id, time); err != nil {
		return err
	}
	return nil
}

// ConstantMeasurement returns a constant timeseries measurement for the same instrument by constant name
func ConstantMeasurement(db *sqlx.DB, tsID *uuid.UUID, constantName string) (*ts.Measurement, error) {
	sql := `
	SELECT M.timeseries_id,
		   M.time,
		   M.value
	FROM  timeseries_measurement M
	INNER JOIN v_timeseries_stored T ON T.id = M.timeseries_id
	INNER JOIN parameter P ON P.id = T.parameter_id
	WHERE T.instrument_id IN (
		SELECT  instrument_id
		FROM v_timeseries_stored T
		WHERE t.id= $1
	)
	AND P.name = $2
	`

	ms := make([]ts.Measurement, 0)
	if err := db.Select(
		&ms,
		sql,
		tsID, constantName,
	); err != nil {
		return nil, err
	}

	m := ts.Measurement{}
	if len(ms) > 0 {
		m = ms[0]
	}

	return &m, nil
}

func CreateOrUpdateTimeseriesMeasurementsTxn(txn *sqlx.Tx, mc []ts.MeasurementCollection, doUpsert bool) (*sqlx.Tx, error) {
	doMmt := "DO NOTHING"
	doNotes := "DO NOTHING"
	if doUpsert {
		doMmt = `DO UPDATE SET value = EXCLUDED.value`
		doNotes = `DO UPDATE SET masked = EXCLUDED.masked, validated = EXCLUDED.validated, annotation = EXCLUDED.annotation`
	}

	stmt_measurement, err := txn.Preparex(
		`INSERT INTO timeseries_measurement (timeseries_id, time, value) VALUES ($1, $2, $3)
		 ON CONFLICT ON CONSTRAINT timeseries_unique_time ` + doMmt,
	)
	if err != nil {
		return nil, err
	}
	stmt_notes, err := txn.Preparex(
		`INSERT INTO timeseries_notes (timeseries_id, time, masked, validated, annotation) VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT ON CONSTRAINT notes_unique_time ` + doNotes,
	)
	if err != nil {
		return nil, err
	}

	tss := make([]uuid.UUID, len(mc))

	// Iterate All Timeseries Measurements
	for idx, c := range mc {
		tss[idx] = c.TimeseriesID
		for _, m := range c.Items {
			if _, err := stmt_measurement.Exec(c.TimeseriesID, m.Time, m.Value); err != nil {
				return nil, err
			}

			if m.Masked != nil || m.Validated != nil || m.Annotation != nil {
				if _, err := stmt_notes.Exec(c.TimeseriesID, m.Time, m.Masked, m.Validated, m.Annotation); err != nil {
					return nil, err
				}
			}
		}
	}

	query, args, err := sqlx.In(`
		SELECT alert_config_id, alert_status_id, NOW() AS submitted_at FROM alert_config_instrument
		WHERE instrument_id = ANY(SELECT intrument_id FROM timeseries WHERE id IN (?))
	`, tss)
	if err != nil {
		return nil, err
	}

	stmt_submittal1, err := txn.Preparex(txn.Rebind(query))
	if err != nil {
		return nil, err
	}
	statuses := make([]SubmittalStatusHistory, 0)
	stmt_submittal1.Select(&statuses, args...)

	if len(statuses) > 0 {
		stmt_submittal2, err := txn.Preparex(`
			INSERT INTO submittal_status_history (alert_config_id, alert_status_id, submitted_at) VALUES ($1, $2, $3)
		`)
		if err != nil {
			return nil, err
		}
		for _, s := range statuses {
			if _, err := stmt_submittal2.Exec(s.AlertConfigID, s.AlertStatusID, s.SubmittedAt); err != nil {
				return nil, err
			}
		}
		if err := stmt_submittal2.Close(); err != nil {
			return nil, err
		}
	}

	if err := stmt_measurement.Close(); err != nil {
		return nil, err
	}
	if err := stmt_notes.Close(); err != nil {
		return nil, err
	}
	if err := stmt_submittal1.Close(); err != nil {
		return nil, err
	}

	return txn, nil
}

// CreateTimeseriesMeasurements creates many timeseries from an array of timeseries
func CreateTimeseriesMeasurements(db *sqlx.DB, mc []ts.MeasurementCollection) ([]ts.MeasurementCollection, error) {
	txn, err := db.Beginx()
	if err != nil {
		return nil, err
	}
	defer txn.Rollback()

	txn, err = CreateOrUpdateTimeseriesMeasurementsTxn(txn, mc, false)
	if err != nil {
		return nil, err
	}

	if err := txn.Commit(); err != nil {
		return nil, err
	}

	return mc, nil
}

// CreateOrUpdateTimeseriesMeasurements creates many timeseries from an array of timeseries
// If a timeseries measurement already exists for a given timeseries_id and time, the value is updated
func CreateOrUpdateTimeseriesMeasurements(db *sqlx.DB, mc []ts.MeasurementCollection) ([]ts.MeasurementCollection, error) {
	txn, err := db.Beginx()
	if err != nil {
		return nil, err
	}
	defer txn.Rollback()

	txn, err = CreateOrUpdateTimeseriesMeasurementsTxn(txn, mc, true)
	if err != nil {
		return nil, err
	}

	if err := txn.Commit(); err != nil {
		return nil, err
	}

	return mc, nil
}

// UpdateTimeseriesMeasurements updates many timeseries measurements, "overwriting" time and values to match paylaod
func UpdateTimeseriesMeasurements(db *sqlx.DB, mc []ts.MeasurementCollection, tw *ts.TimeWindow) ([]ts.MeasurementCollection, error) {

	txn, err := db.Beginx()
	if err != nil {
		return nil, err
	}
	defer txn.Rollback()

	stmt_measurement, err := txn.Preparex(
		`DELETE FROM timeseries_measurement WHERE timeseries_id = $1 AND time > $2 AND time < $3; 
		`,
	)
	if err != nil {
		return nil, err
	}
	stmt_notes, err := txn.Preparex(
		`DELETE FROM timeseries_notes WHERE timeseries_id = $1 AND time > $2 AND time < $3; 
		`,
	)
	if err != nil {
		return nil, err
	}

	for _, c := range mc {
		if _, err := stmt_measurement.Exec(c.TimeseriesID, tw.After, tw.Before); err != nil {
			return nil, err
		}
		if _, err := stmt_notes.Exec(c.TimeseriesID, tw.After, tw.Before); err != nil {
			return nil, err
		}
	}

	if err := stmt_measurement.Close(); err != nil {
		return nil, err
	}
	if err := stmt_notes.Close(); err != nil {
		return nil, err
	}

	txn, err = CreateOrUpdateTimeseriesMeasurementsTxn(txn, mc, true)
	if err != nil {
		return nil, err
	}

	if err := txn.Commit(); err != nil {
		return nil, err
	}

	return mc, nil
}
