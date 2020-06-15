package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// TimeseriesMeasurement is a time series data structure
type TimeseriesMeasurement struct {
	ID           string    `json:"id,omitempty"`
	TimeseriesID uuid.UUID `json:"-" db:"timeseries_id"`
	Time         time.Time `json:"time"`
	Value        float32   `json:"value"`
}

// TimeseriesMeasurementCollection is a collection of timeseries measurements
type TimeseriesMeasurementCollection struct {
	TimeseriesID uuid.UUID               `json:"timeseries_id" db:"timeseries_id"`
	Items        []TimeseriesMeasurement `json:"items"`
}

// TimeseriesMeasurementCollectionCollection is a collection of timeseries measurement collections
// i.e an array of structs, each containing timeseries measurements not necessarily from the same time series
type TimeseriesMeasurementCollectionCollection struct {
	Items []TimeseriesMeasurementCollection
}

// UnmarshalJSON implements UnmarshalJSON interface
func (cc *TimeseriesMeasurementCollectionCollection) UnmarshalJSON(b []byte) error {
	switch JSONType(b) {
	case "ARRAY":
		if err := json.Unmarshal(b, &cc.Items); err != nil {
			return err
		}
	case "OBJECT":
		var mc TimeseriesMeasurementCollection
		if err := json.Unmarshal(b, &mc); err != nil {
			return err
		}
		cc.Items = []TimeseriesMeasurementCollection{mc}
	default:
		cc.Items = make([]TimeseriesMeasurementCollection, 0)
	}
	return nil
}

// TimeWindow is a bounding box for time
type TimeWindow struct {
	After  time.Time `json:"after"`
	Before time.Time `json:"before"`
}

// ListTimeseriesMeasurements returns a timeseries with slice of timeseries measurements populated
func ListTimeseriesMeasurements(db *sqlx.DB, timeseriesID *uuid.UUID, tw *TimeWindow) (*TimeseriesMeasurementCollection, error) {

	mc := TimeseriesMeasurementCollection{TimeseriesID: *timeseriesID}
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
func CreateOrUpdateTimeseriesMeasurements(db *sqlx.DB, mc []TimeseriesMeasurementCollection) error {

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
