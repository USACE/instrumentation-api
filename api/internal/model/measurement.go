package model

import (
	"context"
	"encoding/json"
	"math"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/util"
	"github.com/google/uuid"
)

// TimeseriesMeasurementCollectionCollection is a collection of timeseries measurement collections
// i.e an array of structs, each containing timeseries measurements not necessarily from the same time series
type TimeseriesMeasurementCollectionCollection struct {
	Items []MeasurementCollection
}

// TimeseriesIDs returns a slice of all timeseries IDs contained in the MeasurementCollectionCollection
func (cc *TimeseriesMeasurementCollectionCollection) TimeseriesIDs() map[uuid.UUID]struct{} {
	dd := make(map[uuid.UUID]struct{})
	for _, item := range cc.Items {
		dd[item.TimeseriesID] = struct{}{}
	}
	return dd
}

// UnmarshalJSON implements UnmarshalJSON interface
func (cc *TimeseriesMeasurementCollectionCollection) UnmarshalJSON(b []byte) error {
	switch util.JSONType(b) {
	case "ARRAY":
		if err := json.Unmarshal(b, &cc.Items); err != nil {
			return err
		}
	case "OBJECT":
		var mc MeasurementCollection
		if err := json.Unmarshal(b, &mc); err != nil {
			return err
		}
		cc.Items = []MeasurementCollection{mc}
	default:
		cc.Items = make([]MeasurementCollection, 0)
	}
	return nil
}

// Measurement is a time and value associated with a timeseries
type Measurement struct {
	TimeseriesID uuid.UUID `json:"-" db:"timeseries_id"`
	Time         time.Time `json:"time"`
	Value        float64   `json:"value"`
	Error        string    `json:"error,omitempty"`
	TimeseriesNote
}

type TimeseriesMeasurement struct {
	TimeseriesID uuid.UUID `json:"timeseries_id" db:"timeseries_id"`
	Measurement
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

type MeasurementGetter interface {
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

const (
	createTimeseriesMeasurementSQL = `
		INSERT INTO timeseries_measurement (timeseries_id, time, value) VALUES ($1, $2, $3)
	`
	createTimeseriesNoteSQL = `
		INSERT INTO timeseries_notes (timeseries_id, time, masked, validated, annotation) VALUES ($1, $2, $3, $4, $5)
	`
)

const listTimeseriesMeasurements = `
	SELECT
		m.timeseries_id,
		m.time,
		m.value,
		n.masked,
		n.validated,
		n.annotation
	FROM timeseries_measurement m
	LEFT JOIN timeseries_notes n ON m.timeseries_id = n.timeseries_id AND m.time = n.time
	INNER JOIN timeseries t ON t.id = m.timeseries_id
	WHERE t.id = $1 AND m.time > $2 AND m.time < $3 ORDER BY m.time ASC
`

// ListTimeseriesMeasurements returns a stored timeseries with slice of timeseries measurements populated
func (q *Queries) ListTimeseriesMeasurements(ctx context.Context, timeseriesID uuid.UUID, tw TimeWindow, threshold int) (*MeasurementCollection, error) {
	items := make([]Measurement, 0)
	if err := q.db.SelectContext(ctx, &items, listTimeseriesMeasurements, timeseriesID, tw.After, tw.Before); err != nil {
		return nil, err
	}
	return &MeasurementCollection{TimeseriesID: timeseriesID, Items: LTTB(items, threshold)}, nil
}

const deleteTimeseriesMeasurements = `
	DELETE FROM timeseries_measurement WHERE timeseries_id = $1 and time = $2
`

// DeleteTimeserieMeasurements deletes a timeseries Measurement
func (q *Queries) DeleteTimeserieMeasurements(ctx context.Context, timeseriesID uuid.UUID, time time.Time) error {
	_, err := q.db.ExecContext(ctx, deleteTimeseriesMeasurements, timeseriesID, time)
	return err
}

const getTimeseriesConstantMeasurement = `
	SELECT
		M.timeseries_id,
		M.time,
		M.value
	FROM  timeseries_measurement M
	INNER JOIN v_timeseries_stored T ON T.id = M.timeseries_id
	INNER JOIN parameter P ON P.id = T.parameter_id
	WHERE T.instrument_id IN (
		SELECT  instrument_id
		FROM v_timeseries_stored T
		WHERE t.id= $1
	)
	AND P.name = $2
`

// GetTimeseriesConstantMeasurement returns a constant timeseries measurement for the same instrument by constant name
func (q *Queries) GetTimeseriesConstantMeasurement(ctx context.Context, timeseriesID uuid.UUID, constantName string) (Measurement, error) {
	var m Measurement
	ms := make([]Measurement, 0)
	if err := q.db.Select(&ms, getTimeseriesConstantMeasurement, timeseriesID, constantName); err != nil {
		return m, err
	}
	if len(ms) > 0 {
		m = ms[0]
	}
	return m, nil
}

const createTimeseriesMeasruement = createTimeseriesMeasurementSQL + `
	ON CONFLICT ON CONSTRAINT timeseries_unique_time DO NOTHING
`

func (q *Queries) CreateTimeseriesMeasurement(ctx context.Context, timeseriesID uuid.UUID, t time.Time, value float64) error {
	_, err := q.db.ExecContext(ctx, createTimeseriesMeasruement, timeseriesID, t, value)
	return err
}

const createOrUpdateTimeseriesMeasurement = createTimeseriesMeasurementSQL + `
	ON CONFLICT ON CONSTRAINT timeseries_unique_time DO UPDATE SET value = EXCLUDED.value
`

func (q *Queries) CreateOrUpdateTimeseriesMeasurement(ctx context.Context, timeseriesID uuid.UUID, t time.Time, value float64) error {
	_, err := q.db.ExecContext(ctx, createOrUpdateTimeseriesMeasurement, timeseriesID, t, value)
	return err
}

const createTimeseriesNote = createTimeseriesNoteSQL + `
	ON CONFLICT ON CONSTRAINT notes_unique_time DO NOTHING
`

func (q *Queries) CreateTimeseriesNote(ctx context.Context, timeseriesID uuid.UUID, t time.Time, n TimeseriesNote) error {
	_, err := q.db.ExecContext(ctx, createTimeseriesNote, timeseriesID, t, n.Masked, n.Validated, n.Annotation)
	return err
}

const createOrUpdateTimeseriesNote = createTimeseriesNoteSQL + `
	ON CONFLICT ON CONSTRAINT notes_unique_time DO UPDATE SET masked = EXCLUDED.masked, validated = EXCLUDED.validated, annotation = EXCLUDED.annotation
`

func (q *Queries) CreateOrUpdateTimeseriesNote(ctx context.Context, timeseriesID uuid.UUID, t time.Time, n TimeseriesNote) error {
	_, err := q.db.ExecContext(ctx, createOrUpdateTimeseriesNote, timeseriesID, t, n.Masked, n.Validated, n.Annotation)
	return err
}

const deleteTimeseriesMeasurement = `
	DELETE FROM timeseries_measurement WHERE timeseries_id = $1 AND time = $2
`

func (q *Queries) DeleteTimeseriesMeasurement(ctx context.Context, timeseriesID uuid.UUID, t time.Time) error {
	_, err := q.db.ExecContext(ctx, deleteTimeseriesMeasurementsRange, timeseriesID, t)
	return err
}

const deleteTimeseriesMeasurementsRange = `
	DELETE FROM timeseries_measurement WHERE timeseries_id = $1 AND time > $2 AND time < $3
`

func (q *Queries) DeleteTimeseriesMeasurementsByRange(ctx context.Context, timeseriesID uuid.UUID, start, end time.Time) error {
	_, err := q.db.ExecContext(ctx, deleteTimeseriesMeasurementsRange, timeseriesID, start, end)
	return err
}

const deleteTimeseriesNote = `
	DELETE FROM timeseries_notes WHERE timeseries_id = $1 AND time > $2 AND time < $3
`

func (q *Queries) DeleteTimeseriesNote(ctx context.Context, timeseriesID uuid.UUID, start, end time.Time) error {
	_, err := q.db.ExecContext(ctx, deleteTimeseriesNote, timeseriesID, start, end)
	return err
}

// A slightly modified LTTB (Largest-Triange-Three-Buckets) algorithm for downsampling timeseries measurements
// https://godoc.org/github.com/dgryski/go-lttb
func LTTB[T MeasurementGetter](data []T, threshold int) []T {
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
