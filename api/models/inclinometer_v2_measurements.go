package models

import (
	"encoding/json"
	"log"
	"time"

	ts "github.com/USACE/instrumentation-api/api/timeseries"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type InclinometerV2MeasurementCollections struct {
	Items []InclinometerV2MeasurementCollection
}

type InclinometerV2MeasurementCollection struct {
	TimeseriesID               uuid.UUID                   `json:"timeseries_id" db:"timeseries_id"`
	InclinometerV2Measurements []InclinometerV2Measurement `json:"inclinometer_measurements"`
}

type DBInclinometerV2Measurement struct {
	TimeseriesID uuid.UUID `json:"-" db:"timeseries_id"`
	Time         time.Time `json:"time" db:"time"`
	Value        string    `json:"value" db:"value"`
}

type InclinometerV2Measurement struct {
	TimeseriesID uuid.UUID             `json:"-" db:"timeseries_id"`
	Time         time.Time             `json:"time"`
	Value        []InclinometerV2Value `json:"value,omitempty"`
}

type InclinometerV2Value struct {
	Depth      float32 `json:"depth" db:"depth"`
	A0         float32 `json:"a0" db:"a0"`
	A180       float32 `json:"a180" db:"a180"`
	B0         float32 `json:"b0" db:"b0"`
	B180       float32 `json:"b180" db:"b180"`
	AChecksum  float32 `json:"aChecksum" db:"a_checksum"`
	AComb      float32 `json:"aComb" db:"a_comb"`
	AIncrement float32 `json:"aIncrement" db:"a_increment"`
	ACumDev    float32 `json:"aCumDev" db:"a_cum_dev"`
	BChecksum  float32 `json:"bChecksum" db:"b_checksum"`
	BComb      float32 `json:"bComb" db:"b_comb"`
	BIncrement float32 `json:"bIncrement" db:"b_increment"`
	BCumDev    float32 `json:"bCumDev" db:"b_cum_dev"`
}

func ListInclinometerV2Measurements(db *sqlx.DB, tsID *uuid.UUID, tw *ts.TimeWindow) (*InclinometerV2MeasurementCollection, error) {
	sql := `
		SELECT * FROM 
	`
	mc := InclinometerV2MeasurementCollection{TimeseriesID: *tsID}
	dbmmts := make([]DBInclinometerV2Measurement, 0)
	if err := db.Select(&dbmmts, sql, tsID, tw.Start, tw.End); err != nil {
		return nil, err
	}

	mmts := make([]InclinometerV2Measurement, len(dbmmts))
	for idx, m := range dbmmts {
		mmts[idx] = InclinometerV2Measurement{
			TimeseriesID: m.TimeseriesID,
			Time:         m.Time,
			Value:        make([]InclinometerV2Value, 0),
		}
		if err := json.Unmarshal([]byte(m.Value), &mmts[idx].Value); err != nil {
			log.Println(err)
		}
	}

	return &mc, nil
}

func ListInclinometerV2MeasurementValue(db *sqlx.DB, timeseriesID *uuid.UUID, time time.Time, inclinometerConstant float64) ([]*InclinometerV2Measurement, error) {
	sql := `SELECT * FROM WHERE timeseries_id = $1 AND time = $2 ORDER BY depth`

	mmt := make([]*InclinometerV2Measurement, 0)
	if err := db.Select(&mmt, sql, timeseriesID, time); err != nil {
		return nil, err
	}

	return mmt, nil
}

func DeleteInclinometerV2Measurements(db *sqlx.DB, id *uuid.UUID, time time.Time) error {
	if _, err := db.Exec("DELETE FROM inclinometer_measurement WHERE timeseries_id = $1 AND time = $2", id, time); err != nil {
		return err
	}
	return nil
}

func CreateOrUpdateInclinometerV2MeasurementsTxn(txn *sqlx.Tx, smc InclinometerV2MeasurementCollections, doUpsert bool) error {
	doMmt := "DO NOTHING"

	if doUpsert {
		doMmt = "DO UPDATE SET values = EXCLUDED.values"
	}

	stmt, err := txn.Preparex(`
		INSERT INTO inclinometer_measurement (timeseries_id, time, depth, a0, a80, b0, b180) VALUES ($1,$2,$3,$4,$5,$6,$7)
		ON CONFLICT ON CONSTRAINT inclinometer_measurement_time_depth_key ` + doMmt,
	)
	if err != nil {
		return err
	}

	for i := range smc.Items {
		for j := range smc.Items[i].InclinometerV2Measurements {
			for k := range smc.Items[i].InclinometerV2Measurements[j].Value {
				if _, err := stmt.Exec(
					smc.Items[i].TimeseriesID,
					smc.Items[i].InclinometerV2Measurements[j].Time,
					smc.Items[i].InclinometerV2Measurements[j].Value[k].Depth,
					smc.Items[i].InclinometerV2Measurements[j].Value[k].A0,
					smc.Items[i].InclinometerV2Measurements[j].Value[k].A180,
					smc.Items[i].InclinometerV2Measurements[j].Value[k].B0,
					smc.Items[i].InclinometerV2Measurements[j].Value[k].B180,
				); err != nil {
					return err
				}
			}
		}
	}
	if err := stmt.Close(); err != nil {
		return err
	}

	return nil
}

func CreateInclinometerV2Measurements(db *sqlx.DB, mcs InclinometerV2MeasurementCollections) error {
	txn, err := db.Beginx()
	if err != nil {
		return err
	}
	defer txn.Rollback()

	if err := CreateOrUpdateInclinometerV2MeasurementsTxn(txn, mcs, false); err != nil {
		return err
	}

	if err := txn.Commit(); err != nil {
		return err
	}

	return nil
}

// TODO: make this sensible
func CreateOrUpdateInclinometerV2Measurements(db *sqlx.DB, mcs InclinometerV2MeasurementCollections) error {
	txn, err := db.Beginx()
	if err != nil {
		return err
	}
	defer txn.Rollback()

	if err := CreateOrUpdateInclinometerV2MeasurementsTxn(txn, mcs, true); err != nil {
		return err
	}

	if len(mcs.Items) > 0 {
		mmt, err := GetTimeseriesConstantTxn(txn, &mcs.Items[0].TimeseriesID, "inclinometer-constant")
		if err != nil {
			return err
		}

		if mmt.TimeseriesID == uuid.Nil {
			err := CreateTimeseriesConstantTxn(txn, &mcs.Items[0].TimeseriesID, "inclinometer-constant", "Meters", 20000)
			if err != nil {
				return err
			}
		}
	}

	if err := txn.Commit(); err != nil {
		return err
	}

	return nil
}
