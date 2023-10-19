package model

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/util"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx/types"
)

// Inclinometer Measurement is a time and values associated with a timeseries
type InclinometerMeasurement struct {
	TimeseriesID uuid.UUID      `json:"-" db:"timeseries_id"`
	Time         time.Time      `json:"time"`
	Values       types.JSONText `json:"values"`
	Creator      uuid.UUID      `json:"creator"`
	CreateDate   time.Time      `json:"create_date" db:"create_date"`
}

// Values associated with a inclinometer measurement
type InclinometerMeasurementValues struct {
	Depth      int     `json:"depth" db:"depth"`
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

// InclinometerMeasurementLean is the minimalist representation of a timeseries measurement
// a key value pair where key is the timestamp, value is the measurement { <time.Time>: <types.JSONText> }
type InclinometerMeasurementLean map[time.Time]types.JSONText

// InclinometerMeasurementCollection is a collection of Inclinometer measurements
type InclinometerMeasurementCollection struct {
	TimeseriesID  uuid.UUID                 `json:"timeseries_id" db:"timeseries_id"`
	Inclinometers []InclinometerMeasurement `json:"inclinometers"`
}

// InclinometerMeasurementCollectionLean uses a minimalist representation of a Inclinometer timeseries measurement
type InclinometerMeasurementCollectionLean struct {
	TimeseriesID uuid.UUID                     `json:"timeseries_id" db:"timeseries_id"`
	Items        []InclinometerMeasurementLean `json:"items"`
}

// InclinometerMeasurementCollectionCollection is a collection of inclinometer measurement collections
// i.e an array of structs, each containing inclinometer measurements not necessarily from the same time series
type InclinometerMeasurementCollectionCollection struct {
	Items []InclinometerMeasurementCollection
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
	switch util.JSONType(b) {
	case "ARRAY":
		if err := json.Unmarshal(b, &cc.Items); err != nil {
			return err
		}
	case "OBJECT":
		var mc InclinometerMeasurementCollection
		if err := json.Unmarshal(b, &mc); err != nil {
			return err
		}
		cc.Items = []InclinometerMeasurementCollection{mc}
	default:
		cc.Items = make([]InclinometerMeasurementCollection, 0)
	}
	return nil
}

const listInclinometerMeasurements = `
	SELECT  M.timeseries_id,
		M.time,
		M.creator,
		M.create_date
	FROM inclinometer_measurement M
	INNER JOIN timeseries T
	ON T.id = M.timeseries_id
	WHERE T.id = $1 AND M.time > $2 AND M.time < $3 ORDER BY M.time DESC
`

// ListInclinometersMeasurements returns a timeseries with slice of inclinometer measurements populated
func (q *Queries) ListInclinometerMeasurements(ctx context.Context, timeseriesID uuid.UUID, tw TimeWindow) (*InclinometerMeasurementCollection, error) {
	mc := InclinometerMeasurementCollection{TimeseriesID: timeseriesID}
	if err := q.db.SelectContext(ctx, &mc.Inclinometers, listInclinometerMeasurements, timeseriesID, tw.Start, tw.End); err != nil {
		return nil, err
	}
	return &mc, nil
}

func listInclinometerMeasurementsValues(inclinometerConstant string) string {
	if inclinometerConstant == "0" {
		return `
			select items.depth, 
				items.a0, 
				items.a180, 
				items.b0,
				items.b180,
				(items.a0 + items.a180) AS a_checksum,
				(items.a0 - items.a180)/2 AS a_comb,
				0 AS a_increment,
				0 AS a_cum_dev,
				(items.b0 + items.b180) AS b_checksum,
				(items.b0 - items.b180)/2 AS b_comb,
				0 AS b_increment,
				0 AS b_cum_dev
			from inclinometer_measurement, jsonb_to_recordset(inclinometer_measurement.values) as items(depth int, a0 real, a180 real, b0 real, b180 real)
		`
	} else {
		return fmt.Sprintf(`
			select items.depth, 
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
			from inclinometer_measurement, jsonb_to_recordset(inclinometer_measurement.values) as items(depth int, a0 real, a180 real, b0 real, b180 real)
		`, inclinometerConstant, inclinometerConstant, inclinometerConstant, inclinometerConstant)
	}
}

func (q *Queries) ListInclinometerMeasurementValues(ctx context.Context, timeseriesID uuid.UUID, time time.Time, inclConstant float64) ([]*InclinometerMeasurementValues, error) {
	constant := fmt.Sprintf("%.0f", inclConstant)
	v := []*InclinometerMeasurementValues{}
	if err := q.db.SelectContext(ctx, &v, listInclinometerMeasurementsValues(constant)+" WHERE timeseries_id = $1 AND time = $2 ORDER BY depth", timeseriesID, time); err != nil {
		return nil, err
	}
	return v, nil
}

const deleteInclinometerMeasurement = `
	DELETE FROM inclinometer_measurement WHERE timeseries_id = $1 and time = $2
`

// DeleteInclinometerMeasurements deletes a inclinometer Measurement
func (q *Queries) DeleteInclinometerMeasurement(ctx context.Context, timeseriesID uuid.UUID, time time.Time) error {
	_, err := q.db.ExecContext(ctx, deleteInclinometerMeasurement, timeseriesID, time)
	return err
}

const createOrUpdateInclinometerMeasurement = `
	INSERT INTO inclinometer_measurement (timeseries_id, time, values, creator, create_date) VALUES ($1, $2, $3, $4, $5)
	ON CONFLICT ON CONSTRAINT inclinometer_unique_time DO UPDATE SET values = EXCLUDED.values; 
`

// CreateInclinometerMeasurements creates many inclinometer from an array of inclinometer
// If a inclinometer measurement already exists for a given timeseries_id and time, the values is updated
func (q *Queries) CreateOrUpdateInclinometerMeasurement(ctx context.Context, timeseriesID uuid.UUID, t time.Time, values types.JSONText, profileID uuid.UUID, createDate time.Time) error {
	_, err := q.db.ExecContext(ctx, createOrUpdateInclinometerMeasurement, timeseriesID, t, values, profileID, createDate)
	return err
}

const listInstrumentIDsFromTimeseriesID = `
	SELECT instrument_id FROM v_timeseries_stored WHERE id= $1
`

func (q *Queries) ListInstrumentIDsFromTimeseriesID(ctx context.Context, timeseriesID uuid.UUID) ([]uuid.UUID, error) {
	instrumentIDs := make([]uuid.UUID, 0)
	if err := q.db.SelectContext(ctx, &instrumentIDs, listInstrumentIDsFromTimeseriesID, timeseriesID); err != nil {
		return nil, err
	}
	return instrumentIDs, nil
}

const listParameterIDsFromParameterName = `
	SELECT id FROM parameter WHERE name = $1
`

func (q *Queries) ListParameterIDsFromParameterName(ctx context.Context, parameterName string) ([]uuid.UUID, error) {
	parameterIDs := make([]uuid.UUID, 0)
	if err := q.db.SelectContext(ctx, &parameterIDs, listParameterIDsFromParameterName, parameterName); err != nil {
		return nil, err
	}
	return parameterIDs, nil
}

const listUnitIDsFromUnitName = `
	SELECT id FROM unit WHERE name = $1
`

func (q *Queries) ListUnitIDsFromUnitName(ctx context.Context, unitName string) ([]uuid.UUID, error) {
	unitIDs := make([]uuid.UUID, 0)
	if err := q.db.SelectContext(ctx, &unitIDs, listUnitIDsFromUnitName, unitName); err != nil {
		return nil, err
	}
	return unitIDs, nil
}
