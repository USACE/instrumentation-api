package models

import (
	ts "github.com/USACE/instrumentation-api/timeseries"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	// pq database driver
	_ "github.com/lib/pq"
)

// TimeseriesMeasurementCollectionCollection is a collection of timeseries measurement collections
// i.e an array of structs, each containing timeseries measurements not necessarily from the same time series
type TimeseriesMeasurementCollectionCollection struct {
	Items []ts.MeasurementCollection
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

// ListTimeseriesMeasurements returns a timeseries with slice of timeseries measurements populated
func ListTimeseriesMeasurements(db *sqlx.DB, timeseriesID *uuid.UUID, tw *ts.TimeWindow) (*ts.MeasurementCollection, error) {

	mc := ts.MeasurementCollection{TimeseriesID: *timeseriesID}
	// Get Timeseries Measurements
	if err := db.Select(
		&mc.Items,
		listTimeseriesMeasurementsSQL()+" WHERE T.id = $1 AND M.time > $2 AND M.time < $3",
		timeseriesID, tw.After, tw.Before,
	); err != nil {
		return nil, err
	}

	return &mc, nil
}

// CreateOrUpdateTimeseriesMeasurements creates many timeseries from an array of timeseries
// If a timeseries measurement already exists for a given timeseries_id and time, the value is updated
func CreateOrUpdateTimeseriesMeasurements(db *sqlx.DB, mc []ts.MeasurementCollection) error {

	txn, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := txn.Prepare(
		`INSERT INTO timeseries_measurement (timeseries_id, time, value) VALUES ($1, $2, $3)
		 ON CONFLICT ON CONSTRAINT timeseries_unique_time DO UPDATE SET value = EXCLUDED.value; 
		`,
	)
	if err != nil {
		return err
	}

	// Iterate All Timeseries Measurements
	for _, c := range mc {
		for _, m := range c.Items {
			if _, err := stmt.Exec(c.TimeseriesID, m.Time, m.Value); err != nil {
				return err
			}
		}
	}
	if err := stmt.Close(); err != nil {
		return err
	}
	if err := txn.Commit(); err != nil {
		return err
	}

	return nil
}

func listTimeseriesMeasurementsSQL() string {
	return `SELECT  M.id,
	                M.timeseries_id,
			        M.time,
					M.value
			FROM timeseries_measurement M
			INNER JOIN timeseries T
    			    ON T.id = M.timeseries_id
	`
}
