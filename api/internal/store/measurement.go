package store

import (
	"context"
	"log"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
)

type MeasurementStore interface {
}

type measurementStore struct {
	db *model.Database
	q  *model.Queries
}

func NewMeasurementStore(db *model.Database, q *model.Queries) *measurementStore {
	return &measurementStore{db, q}
}

// ListTimeseriesMeasurements returns a stored timeseries with slice of timeseries measurements populated
func (s measurementStore) ListTimeseriesMeasurements(ctx context.Context, timeseriesID uuid.UUID, tw model.TimeWindow, threshold int) (*model.MeasurementCollection, error) {
	return s.q.ListTimeseriesMeasurements(ctx, timeseriesID, tw, threshold)
}

// DeleteTimeserieMeasurements deletes a timeseries Measurement
func (s measurementStore) DeleteTimeserieMeasurements(ctx context.Context, timeseriesID uuid.UUID, t time.Time) error {
	return s.q.DeleteTimeseriesMeasurement(ctx, timeseriesID, t)
}

// GetTimeseriesConstantMeasurement returns a constant timeseries measurement for the same instrument by constant name
func (s measurementStore) GetTimeseriesConstantMeasurement(ctx context.Context, timeseriesID uuid.UUID, constantName string) (model.Measurement, error) {
	return s.q.GetTimeseriesConstantMeasurement(ctx, timeseriesID, constantName)
}

func (s measurementStore) CreateTimeseriesMeasurement(ctx context.Context, timeseriesID uuid.UUID, t time.Time, value float64) error {
	return s.q.CreateTimeseriesMeasurement(ctx, timeseriesID, t, value)
}

func (s measurementStore) CreateOrUpdateTimeseriesMeasurement(ctx context.Context, timeseriesID uuid.UUID, t time.Time, value float64) error {
	return s.q.CreateOrUpdateTimeseriesMeasurement(ctx, timeseriesID, t, value)
}

func (s measurementStore) CreateTimeseriesNote(ctx context.Context, timeseriesID uuid.UUID, t time.Time, n model.TimeseriesNote) error {
	return s.q.CreateTimeseriesNote(ctx, timeseriesID, t, n)
}

func (s measurementStore) CreateOrUpdateTimeseriesNote(ctx context.Context, timeseriesID uuid.UUID, t time.Time, n model.TimeseriesNote) error {
	return s.q.CreateOrUpdateTimeseriesNote(ctx, timeseriesID, t, n)
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

	qtx := s.q.WithTx(tx)

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

	qtx := s.q.WithTx(tx)

	if err := createMeasurements(ctx, mc, qtx.CreateOrUpdateTimeseriesMeasurement, qtx.CreateOrUpdateTimeseriesNote); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return mc, nil
}

func (s measurementStore) DeleteTimeseriesMeasurementRange(ctx context.Context, timeseriesID uuid.UUID, start, end time.Time) error {
	return s.q.DeleteTimeseriesMeasurementsRange(ctx, timeseriesID, start, end)
}

func (s measurementStore) DeleteTimeseriesNote(ctx context.Context, timeseriesID uuid.UUID, start, end time.Time) error {
	return s.q.DeleteTimeseriesNote(ctx, timeseriesID, start, end)
}
