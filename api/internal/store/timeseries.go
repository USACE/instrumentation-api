package store

import (
	"context"
	"log"

	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
)

type TimeseriesStore interface {
	ListTimeseries(ctx context.Context) ([]model.Timeseries, error)
	GetStoredTimeseriesExists(ctx context.Context, timeseriesID uuid.UUID) (bool, error)
	ListTimeseriesSlugs(ctx context.Context) ([]string, error)
	GetTimeseriesProjectMap(ctx context.Context, timeseriesIDs []uuid.UUID) (map[uuid.UUID]uuid.UUID, error)
	ListTimeseriesSlugsForInstrument(ctx context.Context, instrumentID uuid.UUID) ([]string, error)
	ListProjectTimeseries(ctx context.Context, projectID uuid.UUID) ([]model.Timeseries, error)
	ListInstrumentTimeseries(ctx context.Context, instrumentID uuid.UUID) ([]model.Timeseries, error)
	ListInstrumentGroupTimeseries(ctx context.Context, instrumentGroupID uuid.UUID) ([]model.Timeseries, error)
	GetTimeseries(ctx context.Context, timeseriesID uuid.UUID) (model.Timeseries, error)
	CreateTimeseries(ctx context.Context, ts model.Timeseries) (model.Timeseries, error)
	CreateTimeseriesBatch(ctx context.Context, tt []model.Timeseries) ([]model.Timeseries, error)
	UpdateTimeseries(ctx context.Context, ts model.Timeseries) (uuid.UUID, error)
	DeleteTimeseries(ctx context.Context, timeseriesID uuid.UUID) error
}

type timeseriesStore struct {
	db *model.Database
	*model.Queries
}

func NewTimeseriesStore(db *model.Database, q *model.Queries) *timeseriesStore {
	return &timeseriesStore{db, q}
}

func (s timeseriesStore) CreateTimeseriesBatch(ctx context.Context, tt []model.Timeseries) ([]model.Timeseries, error) {
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

	uu := make([]model.Timeseries, len(tt))
	for idx, ts := range tt {
		tsNew, err := qtx.CreateTimeseries(ctx, ts)
		if err != nil {
			return nil, err
		}
		uu[idx] = tsNew
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return uu, nil
}
