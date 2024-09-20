package service

import (
	"context"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
)

type MeasurementService interface {
	ListTimeseriesMeasurements(ctx context.Context, timeseriesID uuid.UUID, tw model.TimeWindow, threshold int) (*model.MeasurementCollection, error)
	DeleteTimeserieMeasurements(ctx context.Context, timeseriesID uuid.UUID, t time.Time) error
	GetTimeseriesConstantMeasurement(ctx context.Context, timeseriesID uuid.UUID, constantName string) (model.Measurement, error)
	CreateTimeseriesMeasurement(ctx context.Context, timeseriesID uuid.UUID, t time.Time, value float64) error
	CreateOrUpdateTimeseriesMeasurement(ctx context.Context, timeseriesID uuid.UUID, t time.Time, value float64) error
	CreateTimeseriesNote(ctx context.Context, timeseriesID uuid.UUID, t time.Time, n model.TimeseriesNote) error
	CreateOrUpdateTimeseriesNote(ctx context.Context, timeseriesID uuid.UUID, t time.Time, n model.TimeseriesNote) error
	CreateTimeseriesMeasurements(ctx context.Context, mc []model.MeasurementCollection) ([]model.MeasurementCollection, error)
	CreateOrUpdateTimeseriesMeasurements(ctx context.Context, mc []model.MeasurementCollection) ([]model.MeasurementCollection, error)
	UpdateTimeseriesMeasurements(ctx context.Context, mc []model.MeasurementCollection, tw model.TimeWindow) ([]model.MeasurementCollection, error)
	DeleteTimeseriesMeasurementsByRange(ctx context.Context, timeseriesID uuid.UUID, start, end time.Time) error
	DeleteTimeseriesNote(ctx context.Context, timeseriesID uuid.UUID, start, end time.Time) error
}

type measurementService struct {
	db *model.Database
	*model.Queries
}

func NewMeasurementService(db *model.Database, q *model.Queries) *measurementService {
	return &measurementService{db, q}
}

type mmtCbk func(context.Context, uuid.UUID, time.Time, float64) error
type noteCbk func(context.Context, uuid.UUID, time.Time, model.TimeseriesNote) error

func createMeasurements(ctx context.Context, mc []model.MeasurementCollection, mmtFn mmtCbk, noteFn noteCbk) error {
	for _, c := range mc {
		for _, m := range c.Items {
			if err := mmtFn(ctx, c.TimeseriesID, m.Time, float64(m.Value)); err != nil {
				return err
			}
			if m.Masked != nil || m.Validated != nil || m.Annotation != nil {
				if err := noteFn(ctx, c.TimeseriesID, m.Time, m.TimeseriesNote); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// CreateTimeseriesMeasurements creates many timeseries from an array of timeseries
func (s measurementService) CreateTimeseriesMeasurements(ctx context.Context, mc []model.MeasurementCollection) ([]model.MeasurementCollection, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)

	if err := createMeasurements(ctx, mc, qtx.CreateTimeseriesMeasurement, qtx.CreateTimeseriesNote); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return mc, nil
}

// CreateOrUpdateTimeseriesMeasurements creates many timeseries from an array of timeseries
// If a timeseries measurement already exists for a given timeseries_id and time, the value is updated
func (s measurementService) CreateOrUpdateTimeseriesMeasurements(ctx context.Context, mc []model.MeasurementCollection) ([]model.MeasurementCollection, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)

	if err := createMeasurements(ctx, mc, qtx.CreateOrUpdateTimeseriesMeasurement, qtx.CreateOrUpdateTimeseriesNote); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return mc, nil
}

// UpdateTimeseriesMeasurements updates many timeseries measurements, "overwriting" time and values to match paylaod
func (s measurementService) UpdateTimeseriesMeasurements(ctx context.Context, mc []model.MeasurementCollection, tw model.TimeWindow) ([]model.MeasurementCollection, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)

	for _, c := range mc {
		if err := qtx.DeleteTimeseriesMeasurementsByRange(ctx, c.TimeseriesID, tw.After, tw.Before); err != nil {
			return nil, err
		}
		if err := qtx.DeleteTimeseriesNote(ctx, c.TimeseriesID, tw.After, tw.Before); err != nil {
			return nil, err
		}
	}

	if err := createMeasurements(ctx, mc, qtx.CreateTimeseriesMeasurement, qtx.CreateTimeseriesNote); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return mc, nil
}
