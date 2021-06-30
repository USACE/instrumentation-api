package timeseries

import (
	"time"

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
	//ACumD      float64 `json:"aCumD" db:"a_cum_d"`
	BChecksum  float32 `json:"bChecksum" db:"b_checksum"`
	BComb      float32 `json:"bComb" db:"b_comb"`
	BIncrement float32 `json:"bIncrement" db:"b_increment"`
	BCumDev    float32 `json:"bCumDev" db:"b_cum_dev"`
	//BCumD      float64 `json:"bCumD" db:"b_cum_d"`
}

// InclinometerMeasurementCollection is a collection of Inclinometer measurements
type InclinometerMeasurementCollection struct {
	TimeseriesID  uuid.UUID                 `json:"timeseries_id" db:"timeseries_id"`
	Inclinometers []InclinometerMeasurement `json:"inclinometers"`
}
