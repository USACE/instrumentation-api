package model

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type SaaInstrument struct {
	Instrument
	NumSegments     int
	BottomElevation float32
	InitialTime     *time.Time
}

type SaaInstrumentWithSegments struct {
	SaaInstrument
	Segments []SaaSegment
}

type SaaSegment struct {
	ID               int
	InstrumentID     uuid.UUID
	Length           float32
	XTimeseriesID    uuid.UUID
	YTimeseriesID    uuid.UUID
	ZTimeseriesID    uuid.UUID
	TempTimeseriesID uuid.UUID
}

type SaaInstrumentMeasurements struct {
	InstrumentID uuid.UUID
	Time         time.Time
	Measurements dbJSONSlice[SAASegmentMeasurement]
}

type SAASegmentMeasurement struct {
	SegmentID uuid.UUID
	X         *float64
	Y         *float64
	Z         *float64
	Temp      *float64
}

// TODO: when creating new timeseries, any depth based instruments should not be available for assignment

const createSaaInstrument = `
	INSERT INTO saa_instrument (instrument_id, num_segments, bottom_elevation, initial_time)
	VALUES ($1, $2, $3, $4)
`

func (q *Queries) CreateSaaInstrument(ctx context.Context, si SaaInstrument) error {
	_, err := q.db.ExecContext(ctx, createSaaInstrument, si.ID, si.NumSegments, si.BottomElevation, si.InitialTime)
	return err
}

const createSaaSegment = `
	INSERT INTO saa_segment (
		instrument_id,
		length,
		x_timeseries_id,
		y_timeseries_id,
		z_timeseries_id,
		temp_timeseries_id
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
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

const getOneSaaInstrumentWithSegments = `
	SELECT * FROM v_saa_instrument WHERE id = $1
`

func (q *Queries) GetOneSaaInstrumentWithSegments(ctx context.Context, instrumentID uuid.UUID) (SaaInstrumentWithSegments, error) {
	var si SaaInstrumentWithSegments
	err := q.db.GetContext(ctx, &si, createSaaSegment, instrumentID)
	return si, err
}

const getAllSaaInstrumentsWithSegmentsForProject = `
	SELECT * FROM v_saa_instrument WHERE project_id = $1
`

func (q *Queries) GetAllSaaInstrumentsWithSegmentsForProject(ctx context.Context, projectID uuid.UUID) ([]SaaInstrumentWithSegments, error) {
	ssi := make([]SaaInstrumentWithSegments, 0)
	err := q.db.SelectContext(ctx, &ssi, createSaaSegment, projectID)
	return ssi, err
}

const updateSaaInstrument = `
	UPDATE saa_instrument SET
		num_segments = $2,
		bottom_elevation = $3,
		initial_time = $4
	WHERE instrument_id = $1
`

func (q *Queries) UpdateSaaInstrument(ctx context.Context, si SaaInstrument) error {
	_, err := q.db.ExecContext(ctx, createSaaInstrument, si.ID, si.NumSegments, si.BottomElevation, si.InitialTime)
	return err
}

const updateSaaInstrumentSegment = `
	UPDATE saa_segment SET
		length = $3,
		x_timeseries_id = $4,
		y_timeseries_id = $5,
		z_timeseries_id = $6,
		temp_timeseries_id = $7
	WHERE id = $1 AND instrument_id = $2
`

func (q *Queries) UpdateSaaInstrumentSegment(ctx context.Context, seg SaaSegment) error {
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
