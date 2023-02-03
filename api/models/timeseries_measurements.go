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
func ListTimeseriesMeasurements(db *sqlx.DB, timeseriesID *uuid.UUID, tw *ts.TimeWindow) (*ts.MeasurementCollection, error) {

	mc := ts.MeasurementCollection{TimeseriesID: *timeseriesID}
	// Get Timeseries Measurements
	if err := db.Select(
		&mc.Items,
		listTimeseriesMeasurementsSQL()+" WHERE T.id = $1 AND M.time > $2 AND M.time < $3 ORDER BY M.time DESC",
		timeseriesID, tw.After, tw.Before,
	); err != nil {
		return nil, err
	}
	if mc.Items == nil {
		mc.Items = make([]ts.Measurement, 0)
	}

	return &mc, nil
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

	ms := []ts.Measurement{}

	if err := db.Select(
		&ms,
		constantMeasurementSQL()+" and P.name = $2",
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

func CreateOrUpdateTimeseriesMeasurementsTxn(txn *sqlx.Tx, mc []ts.MeasurementCollection) (*sqlx.Tx, error) {
	stmt_measurement, err := txn.Preparex(
		`INSERT INTO timeseries_measurement (timeseries_id, time, value) VALUES ($1, $2, $3)
		 ON CONFLICT ON CONSTRAINT timeseries_unique_time DO UPDATE SET value = EXCLUDED.value; 
		`,
	)
	if err != nil {
		return nil, err
	}
	stmt_notes, err := txn.Preparex(
		`INSERT INTO timeseries_notes (timeseries_id, time, masked, validated, annotation) VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT ON CONSTRAINT notes_unique_time DO UPDATE SET masked = EXCLUDED.masked, validated = EXCLUDED.validated, annotation = EXCLUDED.annotation;
		`,
	)
	if err != nil {
		return nil, err
	}

	// Iterate All Timeseries Measurements
	for _, c := range mc {
		for _, m := range c.Items {
			if _, err := stmt_measurement.Exec(c.TimeseriesID, m.Time, m.Value); err != nil {
				return nil, err
			}
			if _, err := stmt_notes.Exec(c.TimeseriesID, m.Time, m.Masked, m.Validated, m.Annotation); err != nil {
				return nil, err
			}
		}
	}
	if err := stmt_measurement.Close(); err != nil {
		return nil, err
	}
	if err := stmt_notes.Close(); err != nil {
		return nil, err
	}

	return txn, nil
}

// CreateOrUpdateTimeseriesMeasurements creates many timeseries from an array of timeseries
// If a timeseries measurement already exists for a given timeseries_id and time, the value is updated
func CreateOrUpdateTimeseriesMeasurements(db *sqlx.DB, mc []ts.MeasurementCollection) ([]ts.MeasurementCollection, error) {
	txn, err := db.Beginx()
	if err != nil {
		return nil, err
	}
	defer txn.Rollback()

	txn, err = CreateOrUpdateTimeseriesMeasurementsTxn(txn, mc)
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

	txn, err = CreateOrUpdateTimeseriesMeasurementsTxn(txn, mc)
	if err != nil {
		return nil, err
	}

	if err := txn.Commit(); err != nil {
		return nil, err
	}

	return mc, nil
}

// Stored timeseries
func listTimeseriesMeasurementsSQL() string {
	return `SELECT  M.timeseries_id,
			        M.time,
					M.value,
					COALESCE(N.masked, 'false') AS masked,
					COALESCE(N.validated, 'false') AS validated,
					COALESCE(N.annotation, '') AS annotation
			FROM timeseries_measurement M
			LEFT JOIN timeseries_notes N
					ON M.timeseries_id = N.timeseries_id
					AND M.time = N.time
			INNER JOIN timeseries T
    			    ON T.id = M.timeseries_id
	`
}

func constantMeasurementSQL() string {
	return `SELECT  M.timeseries_id,
				M.time,
				M.value
			FROM timeseries_measurement M
			INNER JOIN v_timeseries_stored T
				ON T.id = M.timeseries_id
			INNER JOIN parameter P
				ON P.id = T.parameter_id
			WHERE T.instrument_id in (
				SELECT  instrument_id
				FROM v_timeseries_stored T
				WHERE t.id= $1
			)
	`
}