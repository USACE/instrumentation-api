package store

import (
	"context"
	"log"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
)

type MeasurementStore interface {
	ListTimeseriesMeasurements(ctx context.Context, timeseriesID uuid.UUID, tw model.TimeWindow, threshold int) (*model.MeasurementCollection, error)
	DeleteTimeserieMeasurements(ctx context.Context, timeseriesID uuid.UUID, t time.Time) error
	GetTimeseriesConstantMeasurement(ctx context.Context, timeseriesID uuid.UUID, constantName string) (model.Measurement, error)
	CreateTimeseriesMeasurement(ctx context.Context, timeseriesID uuid.UUID, t time.Time, value float64) error
	CreateOrUpdateTimeseriesMeasurement(ctx context.Context, timeseriesID uuid.UUID, t time.Time, value float64) error
	CreateTimeseriesNote(ctx context.Context, timeseriesID uuid.UUID, t time.Time, n model.TimeseriesNote) error
	CreateOrUpdateTimeseriesNote(ctx context.Context, timeseriesID uuid.UUID, t time.Time, n model.TimeseriesNote) error
	CreateTimeseriesMeasurements(ctx context.Context, mc []model.MeasurementCollection) ([]model.MeasurementCollection, error)
	CreateOrUpdateTimeseriesMeasurements(ctx context.Context, mc []model.MeasurementCollection) ([]model.MeasurementCollection, error)
	DeleteTimeseriesMeasurementsByRange(ctx context.Context, timeseriesID uuid.UUID, start, end time.Time) error
	DeleteTimeseriesNote(ctx context.Context, timeseriesID uuid.UUID, start, end time.Time) error
}

type measurementStore struct {
	db *model.Database
	*model.Queries
}

func NewMeasurementStore(db *model.Database, q *model.Queries) *measurementStore {
	return &measurementStore{db, q}
}

type mmtCbk func(context.Context, uuid.UUID, time.Time, float64) error
type noteCbk func(context.Context, uuid.UUID, time.Time, model.TimeseriesNote) error

func createMeasurements(ctx context.Context, mc []model.MeasurementCollection, mmtFn mmtCbk, noteFn noteCbk) error {
	for _, c := range mc {
		for _, m := range c.Items {
			if err := mmtFn(ctx, c.TimeseriesID, m.Time, m.Value); err != nil {
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
func (s measurementStore) CreateTimeseriesMeasurements(ctx context.Context, mc []model.MeasurementCollection) ([]model.MeasurementCollection, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			log.Print(err.Error())
		}
	}()

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
func (s measurementStore) CreateOrUpdateTimeseriesMeasurements(ctx context.Context, mc []model.MeasurementCollection) ([]model.MeasurementCollection, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			log.Print(err.Error())
		}
	}()

	qtx := s.WithTx(tx)

	if err := createMeasurements(ctx, mc, qtx.CreateOrUpdateTimeseriesMeasurement, qtx.CreateOrUpdateTimeseriesNote); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return mc, nil
}
