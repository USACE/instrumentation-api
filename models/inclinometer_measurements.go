package models

import (
	"encoding/json"
	"fmt"
	"time"

	ts "github.com/USACE/instrumentation-api/timeseries"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// InclinometerMeasurementCollectionCollection is a collection of inclinometer measurement collections
// i.e an array of structs, each containing inclinometer measurements not necessarily from the same time series
type InclinometerMeasurementCollectionCollection struct {
	Items []ts.InclinometerMeasurementCollection
}

// InclinometerTimeseriesIDs returns a slice of all timeseries IDs contained in the InclinometerMeasurementCollectionCollection
func (cc *InclinometerMeasurementCollectionCollection) InclinometerTimeseriesIDs() []uuid.UUID {

	dd := make([]uuid.UUID, 0)
	for _, item := range cc.Items {
		dd = append(dd, item.TimeseriesID)
	}
	return dd
}

// UnmarshalJSON implements UnmarshalJSON interface
func (cc *InclinometerMeasurementCollectionCollection) UnmarshalJSON(b []byte) error {
	switch JSONType(b) {
	case "ARRAY":
		if err := json.Unmarshal(b, &cc.Items); err != nil {
			return err
		}
	case "OBJECT":
		var mc ts.InclinometerMeasurementCollection
		if err := json.Unmarshal(b, &mc); err != nil {
			return err
		}
		cc.Items = []ts.InclinometerMeasurementCollection{mc}
	default:
		cc.Items = make([]ts.InclinometerMeasurementCollection, 0)
	}
	return nil
}

// ListInclinometersMeasurements returns a timeseries with slice of inclinometer measurements populated
func ListInclinometerMeasurements(db *sqlx.DB, timeseriesID *uuid.UUID, tw *ts.TimeWindow) (*ts.InclinometerMeasurementCollection, error) {

	mc := ts.InclinometerMeasurementCollection{TimeseriesID: *timeseriesID}
	// Get Timeseries Measurements
	if err := db.Select(
		&mc.Inclinometers,
		listInclinometerMeasurementsSQL()+" WHERE T.id = $1 AND M.time > $2 AND M.time < $3 ORDER BY M.time DESC",
		timeseriesID, tw.After, tw.Before,
	); err != nil {
		return nil, err
	}

	return &mc, nil
}

func ListInclinometerMeasurementValues(db *sqlx.DB, timeseriesID *uuid.UUID, time time.Time, inclinometerConstant float64) ([]*ts.InclinometerMeasurementValues, error) {
	constnat := fmt.Sprintf("%.0f", inclinometerConstant)
	v := []*ts.InclinometerMeasurementValues{}
	// Get Inclinometer Measurement values
	if err := db.Select(
		&v,
		inclinometerMeasurementsValuesSQL(constnat)+" WHERE timeseries_id = $1 AND time = $2 ORDER BY depth",
		timeseriesID, time,
	); err != nil {
		return nil, err
	}

	return v, nil
}

// DeleteInclinometerMeasurements deletes a inclinometer Measurement
func DeleteInclinometerMeasurements(db *sqlx.DB, id *uuid.UUID, time time.Time) error {
	if _, err := db.Exec("DELETE FROM inclinometer_measurement WHERE timeseries_id = $1 and time = $2", id, time); err != nil {
		return err
	}
	return nil
}

// CreateInclinometerMeasurements creates many inclinometer from an array of inclinometer
// If a inclinometer measurement already exists for a given timeseries_id and time, the values is updated
func CreateOrUpdateInclinometerMeasurements(db *sqlx.DB, im []ts.InclinometerMeasurementCollection, p *Profile, createDate time.Time) ([]ts.InclinometerMeasurementCollection, error) {

	txn, err := db.Begin()
	if err != nil {
		return nil, err
	}

	stmt, err := txn.Prepare(
		`INSERT INTO inclinometer_measurement (timeseries_id, time, values, creator, create_date) VALUES ($1, $2, $3, $4, $5)
		 ON CONFLICT ON CONSTRAINT inclinometer_unique_time DO UPDATE SET values = EXCLUDED.values; 
		`,
	)
	if err != nil {
		txn.Rollback()
		return nil, err
	}

	// Iterate All inclinometer Measurements
	for idx, _ := range im {
		for i, _ := range im[idx].Inclinometers {
			im[idx].Inclinometers[i].Creator = p.ID
			im[idx].Inclinometers[i].CreateDate = createDate
			if _, err := stmt.Exec(im[idx].TimeseriesID, im[idx].Inclinometers[i].Time, im[idx].Inclinometers[i].Values, p.ID, createDate); err != nil {
				txn.Rollback()
				return nil, err
			}
		}
	}
	if err := stmt.Close(); err != nil {
		txn.Rollback()
		return nil, err
	}
	if err := txn.Commit(); err != nil {
		txn.Rollback()
		return nil, err
	}

	return im, nil
}

func listInclinometerMeasurementsSQL() string {
	return `SELECT  M.timeseries_id,
			        M.time,
					M.creator,
					M.create_date
			FROM inclinometer_measurement M
			INNER JOIN timeseries T
    			    ON T.id = M.timeseries_id
	`
}

func inclinometerMeasurementsValuesSQL(inclinometerConstant string) string {
	return fmt.Sprintf(`select items.depth, 
					items.a0, 
					items.a180, 
					items.b0,
					items.b180,
					(items.a0 + items.a180) AS a_checksum,
					(items.a0 - items.a180)/2 AS a_comb,
					(items.a0 - items.a180) / 2 / %s * 24 AS a_increment,
					SUM((items.a0 - items.a180) / 2 / %s * 24) OVER (ORDER BY depth desc) AS a_cum_dev,
					(items.b0 + items.b180) AS b_checksum,
					(items.b0 - items.b180)/2 AS b_comb,
					(items.b0 - items.b180) / 2 / %s * 24 AS b_increment,
					SUM((items.b0 - items.b180) / 2 / %s * 24) OVER (ORDER BY depth desc) AS b_cum_dev
		from inclinometer_measurement, jsonb_to_recordset(inclinometer_measurement.values) as items(depth int, a0 real, a180 real, b0 real, b180 real)`, inclinometerConstant, inclinometerConstant, inclinometerConstant, inclinometerConstant)
}
