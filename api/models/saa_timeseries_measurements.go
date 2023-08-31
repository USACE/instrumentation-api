package models

import (
	"time"

	ts "github.com/USACE/instrumentation-api/api/timeseries"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// TODO: Add create DB schema for SaaInstrument
// number segment may change, this property should not affect histroical readings
type SaaInstrument struct {
	NumSegments        int
	BottomElevation    float32
	SegmentLength      float32
	InitialMeasurement SaaMeasurement
}

type SaaMeasurement struct {
	TimeseriesID uuid.UUID  `json:"-" db:"timeseries_id"`
	Time         time.Time  `json:"time"`
	Value        []SaaValue `json:"value"`
}

// SAA needs to select depth and them populate corresponding X, Y, Temp
type SaaValue struct {
	ElevationChange float32 `json:"elevation_change" db:"elevation_change"`
	X               float32 `json:"x" db:"x"`
	Y               float32 `json:"y" db:"y"`
	Temperature     float32 `json:"temperature" db:"temperature"`
}

type SaaMeasurementCollection struct {
	TimeseriesID    uuid.UUID        `json:"timeseries_id" db:"timeseries_id"`
	SaaMeasurements []SaaMeasurement `json:"saa_measurements"`
}

func ListSaaMeasurements(db *sqlx.DB, timeseriesID *uuid.UUID, tw *ts.TimeWindow) (*SaaMeasurementCollection, error) {
	sql := `
		SELECT
			mmt.timeseries_id,
			mmt.time
		FROM saa_measurement mmt
		INNER JOIN timeseries ts
			ON ts.id = mmt.timeseries_id
		WHERE ts.id = $1
		AND mmt.time > $2
		AND mmt.time < $3
		ORDER BY mmt.time DESC
	`

	mc := SaaMeasurementCollection{TimeseriesID: *timeseriesID}
	// Get Timeseries Measurements
	if err := db.Select(&mc.SaaMeasurements, sql, timeseriesID, tw.Start, tw.End); err != nil {
		return nil, err
	}

	return &mc, nil
}

func ListSaaMeasurementValues(db *sqlx.DB, timeseriesID *uuid.UUID, time time.Time, saaConstant float64) ([]*SaaMeasurement, error) {
	sql := `SELECT * FROM WHERE timeseries_id = $1 AND time = $2 ORDER BY depth`

	mmt := make([]*SaaMeasurement, 0)
	if err := db.Select(&mmt, sql, timeseriesID, time); err != nil {
		return nil, err
	}

	return mmt, nil
}

func DeleteSaaMeasurements(db *sqlx.DB, id *uuid.UUID, time time.Time) error {
	if _, err := db.Exec("DELETE FROM saa_measurement WHERE timeseries_id = $1 and time = $2", id, time); err != nil {
		return err
	}
	return nil
}

func CreateOrUpdateSaaMeasurementsTxn(txn *sqlx.Tx, smc []SaaMeasurementCollection, doUpsert bool) error {
	doMmt := "DO NOTHING"

	if doUpsert {
		doMmt = "DO UPDATE SET values = EXCLUDED.values"
	}

	stmt, err := txn.Preparex(`
		INSERT INTO saa_measurement (timeseries_id, time, values) VALUES ($1, $2, $3)
		ON CONFLICT ON CONSTRAINT saa_unique_time ` + doMmt,
	)
	if err != nil {
		return err
	}

	for i := range smc {
		for j := range smc[i].SaaMeasurements {
			if _, err := stmt.Exec(smc[i].TimeseriesID, smc[i].SaaMeasurements[j].Time, smc[i].SaaMeasurements[j].Value); err != nil {
				return err
			}
		}
	}
	if err := stmt.Close(); err != nil {
		return err
	}

	return nil
}

func CreateSaaMeasurements(db *sqlx.DB, smc []SaaMeasurementCollection) error {
	txn, err := db.Beginx()
	if err != nil {
		return err
	}
	defer txn.Rollback()

	if err := CreateOrUpdateSaaMeasurementsTxn(txn, smc, false); err != nil {
		return err
	}

	if err := txn.Commit(); err != nil {
		return err
	}

	return nil
}

func CreateOrUpdateSaaMeasurements(db *sqlx.DB, smc []SaaMeasurementCollection) error {
	txn, err := db.Beginx()
	if err != nil {
		return err
	}
	defer txn.Rollback()

	if err := CreateOrUpdateSaaMeasurementsTxn(txn, smc, true); err != nil {
		return err
	}

	if err := txn.Commit(); err != nil {
		return err
	}

	return nil
}
