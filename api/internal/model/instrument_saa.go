package model

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type SaaOpts struct {
	InstrumentID    uuid.UUID  `json:"-" db:"instrument_id"`
	NumSegments     int        `json:"num_segments" db:"num_segments"`
	BottomElevation float32    `json:"bottom_elevation" db:"bottom_elevation"`
	InitialTime     *time.Time `json:"initial_time" db:"initial_time"`
}

type SaaSegment struct {
	ID               int        `json:"id" db:"id"`
	InstrumentID     uuid.UUID  `json:"instrument_id" db:"instrument_id"`
	Length           *float32   `json:"length" db:"length"`
	XTimeseriesID    *uuid.UUID `json:"x_timeseries_id" db:"x_timeseries_id"`
	YTimeseriesID    *uuid.UUID `json:"y_timeseries_id" db:"y_timeseries_id"`
	ZTimeseriesID    *uuid.UUID `json:"z_timeseries_id" db:"z_timeseries_id"`
	TempTimeseriesID *uuid.UUID `json:"temp_timeseries_id" db:"temp_timeseries_id"`
}

type SaaMeasurements struct {
	InstrumentID uuid.UUID                          `json:"-" db:"instrument_id"`
	Time         time.Time                          `json:"time" db:"time"`
	Measurements dbJSONSlice[SaaSegmentMeasurement] `json:"measurements" db:"measurements"`
}

type SaaSegmentMeasurement struct {
	SegmentID     int      `json:"segment_id" db:"segment_id"`
	X             *float64 `json:"x" db:"x"`
	Y             *float64 `json:"y" db:"y"`
	Z             *float64 `json:"z" db:"z"`
	Temp          *float64 `json:"temp" db:"temp"`
	XIncrement    *float64 `json:"x_increment" db:"x_increment"`
	YIncrement    *float64 `json:"y_increment" db:"y_increment"`
	ZIncrement    *float64 `json:"z_increment" db:"z_increment"`
	TempIncrement *float64 `json:"temp_increment" db:"temp_increment"`
	XCumDev       *float64 `json:"x_cum_dev" db:"x_cum_dev"`
	YCumDev       *float64 `json:"y_cum_dev" db:"y_cum_dev"`
	ZCumDev       *float64 `json:"z_cum_dev" db:"z_cum_dev"`
	TempCumDev    *float64 `json:"temp_cum_dev" db:"temp_cum_dev"`
}

// TODO: when creating new timeseries, any depth based instruments should not be available for assignment

const createSaaOpts = `
	INSERT INTO saa_opts (instrument_id, num_segments, bottom_elevation, initial_time)
	VALUES ($1, $2, $3, $4)
`

func (q *Queries) CreateSaaOpts(ctx context.Context, instrumentID uuid.UUID, si SaaOpts) error {
	_, err := q.db.ExecContext(ctx, createSaaOpts, instrumentID, si.NumSegments, si.BottomElevation, si.InitialTime)
	return err
}

const updateSaaOpts = `
	UPDATE saa_opts SET
		bottom_elevation = $2,
		initial_time = $3
	WHERE instrument_id = $1
`

func (q *Queries) UpdateSaaOpts(ctx context.Context, instrumentID uuid.UUID, si SaaOpts) error {
	_, err := q.db.ExecContext(ctx, updateSaaOpts, si.InstrumentID, si.BottomElevation, si.InitialTime)
	return err
}

const getAllSaaSegmentsForInstrument = `
	SELECT * FROM v_saa_segment WHERE instrument_id = $1
`

func (q *Queries) GetAllSaaSegmentsForInstrument(ctx context.Context, instrumentID uuid.UUID) ([]SaaSegment, error) {
	ssi := make([]SaaSegment, 0)
	err := q.db.SelectContext(ctx, &ssi, getAllSaaSegmentsForInstrument, instrumentID)
	return ssi, err
}

const createSaaSegment = `
	INSERT INTO saa_segment (
		instrument_id,
		length,
		x_timeseries_id,
		y_timeseries_id,
		z_timeseries_id,
		temp_timeseries_id
	) VALUES ($1, $2, $3, $4, $5, $6)
`

func (q *Queries) CreateSaaSegment(ctx context.Context, seg SaaSegment) error {
	_, err := q.db.ExecContext(ctx, createSaaSegment,
		seg.InstrumentID,
		seg.Length,
		seg.XTimeseriesID,
		seg.YTimeseriesID,
		seg.ZTimeseriesID,
		seg.TempTimeseriesID,
	)
	return err
}

const updateSaaSegment = `
	UPDATE saa_segment SET
		length = $3,
		x_timeseries_id = $4,
		y_timeseries_id = $5,
		z_timeseries_id = $6,
		temp_timeseries_id = $7
	WHERE id = $1 AND instrument_id = $2
`

func (q *Queries) UpdateSaaSegment(ctx context.Context, seg SaaSegment) error {
	_, err := q.db.ExecContext(ctx, updateSaaSegment,
		seg.ID,
		seg.InstrumentID,
		seg.Length,
		seg.XTimeseriesID,
		seg.YTimeseriesID,
		seg.ZTimeseriesID,
		seg.TempTimeseriesID,
	)
	return err
}

const getSaaMeasurementsForInstrument = `
	SELECT instrument_id, time, measurements
	FROM v_saa_measurement
	WHERE instrument_id = $1 AND time >= $2 AND time <= $3
	OR time IN (SELECT initial_time FROM saa_opts WHERE instrument_id = $1)
`

func (q *Queries) GetSaaMeasurementsForInstrument(ctx context.Context, instrumentID uuid.UUID, tw TimeWindow) ([]SaaMeasurements, error) {
	mm := make([]SaaMeasurements, 0)
	err := q.db.SelectContext(ctx, &mm, getSaaMeasurementsForInstrument, instrumentID, tw.Start, tw.End)
	return mm, err
}
