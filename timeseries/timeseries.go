package timeseries

import (
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx/types"
)

// Timeseries is a timeseries data structure
type Timeseries struct {
	ID             uuid.UUID     `json:"id"`
	Slug           string        `json:"slug"`
	Name           string        `json:"name"`
	Variable       string        `json:"variable"`
	ProjectID      uuid.UUID     `json:"project_id" db:"project_id"`
	ProjectSlug    string        `json:"project_slug" db:"project_slug"`
	Project        string        `json:"project,omitempty" db:"project"`
	InstrumentID   uuid.UUID     `json:"instrument_id" db:"instrument_id"`
	InstrumentSlug string        `json:"instrument_slug" db:"instrument_slug"`
	Instrument     string        `json:"instrument,omitempty"`
	ParameterID    uuid.UUID     `json:"parameter_id" db:"parameter_id"`
	Parameter      string        `json:"parameter,omitempty"`
	UnitID         uuid.UUID     `json:"unit_id" db:"unit_id"`
	Unit           string        `json:"unit,omitempty"`
	Values         []Measurement `json:"values,omitempty"`
	IsComputed     bool          `json:"is_computed" db:"is_computed"`
}

// Measurement is a time and value associated with a timeseries
type Measurement struct {
	TimeseriesID uuid.UUID `json:"-" db:"timeseries_id"`
	Time         time.Time `json:"time"`
	Value        float64   `json:"value"`
}

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
	ACumD      float32 `json:"aCumD" db:"a_cum_d"`
	BChecksum  float32 `json:"bChecksum" db:"b_checksum"`
	BComb      float32 `json:"bComb" db:"b_comb"`
	BIncrement float32 `json:"bIncrement" db:"b_increment"`
	BCumDev    float32 `json:"bCumDev" db:"b_cum_dev"`
	BCumD      float32 `json:"bCumD" db:"b_cum_d"`
}

// MeasurementLean is the minimalist representation of a timeseries measurement
// a key value pair where key is the timestamp, value is the measurement { <time.Time>: <float32> }
type MeasurementLean map[time.Time]float64

// MeasurementCollection is a collection of timeseries measurements
type MeasurementCollection struct {
	TimeseriesID uuid.UUID     `json:"timeseries_id" db:"timeseries_id"`
	Items        []Measurement `json:"items"`
}

// MeasurementCollectionLean uses a minimalist representation of a timeseries measurement
type MeasurementCollectionLean struct {
	TimeseriesID uuid.UUID         `json:"timeseries_id" db:"timeseries_id"`
	Items        []MeasurementLean `json:"items"`
}

// MeasurementCollection is a collection of timeseries measurements
type InclinometerMeasurementCollection struct {
	TimeseriesID  uuid.UUID                 `json:"timeseries_id" db:"timeseries_id"`
	Inclinometers []InclinometerMeasurement `json:"inclinometers"`
}
