package timeseries

import (
	"time"

	"github.com/google/uuid"
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
}

// Measurement is a time and value associated with a timeseries
type Measurement struct {
	TimeseriesID uuid.UUID `json:"-" db:"timeseries_id"`
	Time         time.Time `json:"time"`
	Value        float32   `json:"value"`
}

// MeasurementLean is the minimalist representation of a timeseries measurement
// a key value pair where key is the timestamp, value is the measurement { <time.Time>: <float32> }
type MeasurementLean map[time.Time]float32

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
