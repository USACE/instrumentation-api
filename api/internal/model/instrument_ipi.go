package model

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type IpiOpts struct {
	InstrumentID                uuid.UUID  `json:"-" db:"instrument_id"`
	NumSegments                 int        `json:"num_segments" db:"num_segments"`
	BottomElevationTimeseriesID uuid.UUID  `json:"bottom_elevation_timeseries_id" db:"bottom_elevation_timeseries_id"`
	BottomElevation             float64    `json:"bottom_elevation" db:"bottom_elevation"`
	InitialTime                 *time.Time `json:"initial_time" db:"initial_time"`
}

type IpiSegment struct {
	ID                 int        `json:"id" db:"id"`
	InstrumentID       uuid.UUID  `json:"instrument_id" db:"instrument_id"`
	Length             *float64   `json:"length" db:"length"`
	LengthTimeseriesID uuid.UUID  `json:"length_timeseries_id" db:"length_timeseries_id"`
	TiltTimeseriesID   *uuid.UUID `json:"tilt_timeseries_id" db:"tilt_timeseries_id"`
	IncDevTimeseriesID *uuid.UUID `json:"inc_dev_timeseries_id" db:"inc_dev_timeseries_id"`
	TempTimeseriesID   *uuid.UUID `json:"temp_timeseries_id" db:"temp_timeseries_id"`
}

type IpiMeasurements struct {
	InstrumentID uuid.UUID                          `json:"-" db:"instrument_id"`
	Time         time.Time                          `json:"time" db:"time"`
	Measurements dbJSONSlice[IpiSegmentMeasurement] `json:"measurements" db:"measurements"`
}

type IpiSegmentMeasurement struct {
	SegmentID  int      `json:"segment_id" db:"segment_id"`
	Tilt       *float64 `json:"tilt" db:"tilt"`
	IncDev     *float64 `json:"inc_dev" db:"inc_dev"`
	CumDev     *float64 `json:"cum_dev" db:"cum_dev"`
	Temp       *float64 `json:"temp" db:"temp"`
	Elelvation *float64 `json:"elevation" db:"elevation"`
}

var (
	IpiParameterID = uuid.MustParse("a9a5ad45-b2e5-4744-816e-d3184f2c08bd")
)

// TODO: when creating new timeseries, any depth based instruments should not be available for assignment

const createIpiOpts = `
	INSERT INTO ipi_opts (instrument_id, num_segments, bottom_elevation_timeseries_id, initial_time)
	VALUES ($1, $2, $3, $4)
`

func (q *Queries) CreateIpiOpts(ctx context.Context, instrumentID uuid.UUID, si IpiOpts) error {
	_, err := q.db.ExecContext(ctx, createIpiOpts, instrumentID, si.NumSegments, si.BottomElevationTimeseriesID, si.InitialTime)
	return err
}

const updateIpiOpts = `
	UPDATE ipi_opts SET
		bottom_elevation_timeseries_id = $2,
		initial_time = $3
	WHERE instrument_id = $1
`

func (q *Queries) UpdateIpiOpts(ctx context.Context, instrumentID uuid.UUID, si IpiOpts) error {
	_, err := q.db.ExecContext(ctx, updateIpiOpts, instrumentID, si.BottomElevationTimeseriesID, si.InitialTime)
	return err
}

const getAllIpiSegmentsForInstrument = `
	SELECT * FROM v_ipi_segment WHERE instrument_id = $1
`

func (q *Queries) GetAllIpiSegmentsForInstrument(ctx context.Context, instrumentID uuid.UUID) ([]IpiSegment, error) {
	ssi := make([]IpiSegment, 0)
	err := q.db.SelectContext(ctx, &ssi, getAllIpiSegmentsForInstrument, instrumentID)
	return ssi, err
}

const createIpiSegment = `
	INSERT INTO ipi_segment (
		id,
		instrument_id,
		length_timeseries_id,
		tilt_timeseries_id,
		inc_dev_timeseries_id,
		temp_timeseries_id
	) VALUES ($1, $2, $3, $4, $5, $6)
`

func (q *Queries) CreateIpiSegment(ctx context.Context, seg IpiSegment) error {
	_, err := q.db.ExecContext(ctx, createIpiSegment,
		seg.ID,
		seg.InstrumentID,
		seg.LengthTimeseriesID,
		seg.TiltTimeseriesID,
		seg.IncDevTimeseriesID,
		seg.TempTimeseriesID,
	)
	return err
}

const updateIpiSegment = `
	UPDATE ipi_segment SET
		length_timeseries_id = $3,
		tilt_timeseries_id = $4,
		inc_dev_timeseries_id = $5,
		temp_timeseries_id = $6
	WHERE id = $1 AND instrument_id = $2
`

func (q *Queries) UpdateIpiSegment(ctx context.Context, seg IpiSegment) error {
	_, err := q.db.ExecContext(ctx, updateIpiSegment,
		seg.ID,
		seg.InstrumentID,
		seg.LengthTimeseriesID,
		seg.TiltTimeseriesID,
		seg.IncDevTimeseriesID,
		seg.TempTimeseriesID,
	)
	return err
}

const getIpiMeasurementsForInstrument = `
	SELECT instrument_id, time, measurements
	FROM v_ipi_measurement
	WHERE instrument_id = $1 AND time >= $2 AND time <= $3
	UNION
	SELECT instrument_id, time, measurements
	FROM v_ipi_measurement
	WHERE time IN (SELECT initial_time FROM ipi_opts WHERE instrument_id = $1)
	AND instrument_id = $1
	ORDER BY time ASC
`

func (q *Queries) GetIpiMeasurementsForInstrument(ctx context.Context, instrumentID uuid.UUID, tw TimeWindow) ([]IpiMeasurements, error) {
	mm := make([]IpiMeasurements, 0)
	err := q.db.SelectContext(ctx, &mm, getIpiMeasurementsForInstrument, instrumentID, tw.After, tw.Before)
	return mm, err
}
