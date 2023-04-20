package timeseries

import (
	"math"
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
	IsComputed     bool          `json:"is_computed" db:"is_computed"`
}

// Measurement is a time and value associated with a timeseries
type Measurement struct {
	TimeseriesID uuid.UUID `json:"-" db:"timeseries_id"`
	Time         time.Time `json:"time"`
	Value        float64   `json:"value"`
	Error        string    `json:"error,omitempty"`
	TimeseriesNote
}

type TimeseriesNote struct {
	Masked     *bool   `json:"masked,omitempty"`
	Validated  *bool   `json:"validated,omitempty"`
	Annotation *string `json:"annotation,omitempty"`
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

type MeasurementLike interface {
	getTime() time.Time
	getValue() float64
}

func (m Measurement) getTime() time.Time {
	return m.Time
}

func (m Measurement) getValue() float64 {
	return m.Value
}

// Should only ever be one
func (ml MeasurementLean) getTime() time.Time {
	var t time.Time
	for k := range ml {
		t = k
	}
	return t
}

// Should only ever be one
func (ml MeasurementLean) getValue() float64 {
	var m float64
	for _, v := range ml {
		m = v
	}
	return m
}

// A slightly modified LTTB (Largest-Triange-Three-Buckets) algorithm for downsampling timeseries measurements while keeping
// https://godoc.org/github.com/dgryski/go-lttb
func LTTB[T MeasurementLike](data []T, threshold int) []T {
	if threshold == 0 || threshold >= len(data) {
		return data // Nothing to do
	}

	if threshold < 3 {
		threshold = 3
	}

	sampled := make([]T, 0, threshold)

	// Bucket size. Leave room for start and end data points
	every := float64(len(data)-2) / float64(threshold-2)

	sampled = append(sampled, data[0]) // Always add the first point

	bucketStart := 1
	bucketCenter := int(math.Floor(every)) + 1

	var a int

	for i := 0; i < threshold-2; i++ {

		bucketEnd := int(math.Floor(float64(i+2)*every)) + 1

		// Calculate point average for next bucket (containing c)
		avgRangeStart := bucketCenter
		avgRangeEnd := bucketEnd

		if avgRangeEnd >= len(data) {
			avgRangeEnd = len(data)
		}

		avgRangeLength := float64(avgRangeEnd - avgRangeStart)

		var avgX, avgY float64
		for ; avgRangeStart < avgRangeEnd; avgRangeStart++ {
			avgX += time.Duration(data[avgRangeStart].getTime().Unix()).Seconds()
			avgY += data[avgRangeStart].getValue()
		}
		avgX /= avgRangeLength
		avgY /= avgRangeLength

		// Get the range for this bucket
		rangeOffs := bucketStart
		rangeTo := bucketCenter

		// Point a
		pointAX := time.Duration(data[a].getTime().UnixNano()).Seconds()
		pointAY := data[a].getValue()

		maxArea := float64(-1.0)

		var nextA int
		for ; rangeOffs < rangeTo; rangeOffs++ {
			// Calculate triangle area over three buckets
			area := (pointAX-avgX)*(data[rangeOffs].getValue()-pointAY) - (pointAX-time.Duration(data[rangeOffs].getTime().Unix()).Seconds())*(avgY-pointAY)
			// We only care about the relative area here.
			// Calling math.Abs() is slower than squaring
			area *= area
			if area > maxArea {
				maxArea = area
				nextA = rangeOffs // Next a is this b
			}
		}

		sampled = append(sampled, data[nextA]) // Pick this point from the bucket
		a = nextA                              // This a is the next a (chosen b)

		bucketStart = bucketCenter
		bucketCenter = bucketEnd
	}

	sampled = append(sampled, data[len(data)-1]) // Always add last

	return sampled
}
