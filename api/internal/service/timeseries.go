package service

import (
	"context"
	"errors"

	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
)

type TimeseriesService interface {
	GetStoredTimeseriesExists(ctx context.Context, timeseriesID uuid.UUID) (bool, error)
	AssertTimeseriesLinkedToProject(ctx context.Context, projectID uuid.UUID, dd map[uuid.UUID]struct{}) error
	ListProjectTimeseries(ctx context.Context, projectID uuid.UUID) ([]model.Timeseries, error)
	ListInstrumentTimeseries(ctx context.Context, instrumentID uuid.UUID) ([]model.Timeseries, error)
	ListInstrumentGroupTimeseries(ctx context.Context, instrumentGroupID uuid.UUID) ([]model.Timeseries, error)
	ListPlotConfigTimeseries(ctx context.Context, plotConfigID uuid.UUID) ([]model.Timeseries, error)
	GetTimeseries(ctx context.Context, timeseriesID uuid.UUID) (model.Timeseries, error)
	CreateTimeseries(ctx context.Context, ts model.Timeseries) (model.Timeseries, error)
	CreateTimeseriesBatch(ctx context.Context, tt []model.Timeseries) ([]model.Timeseries, error)
	UpdateTimeseries(ctx context.Context, ts model.Timeseries) (uuid.UUID, error)
	DeleteTimeseries(ctx context.Context, timeseriesID uuid.UUID) error
}

type timeseriesService struct {
	db *model.Database
	*model.Queries
}

func NewTimeseriesService(db *model.Database, q *model.Queries) *timeseriesService {
	return &timeseriesService{db, q}
}

func (s timeseriesService) CreateTimeseriesBatch(ctx context.Context, tt []model.Timeseries) ([]model.Timeseries, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)

	uu := make([]model.Timeseries, len(tt))
	for idx, ts := range tt {
		ts.Type = model.StandardTimeseriesType
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

func (s timeseriesService) AssertTimeseriesLinkedToProject(ctx context.Context, projectID uuid.UUID, dd map[uuid.UUID]struct{}) error {
	ddc := make(map[uuid.UUID]struct{}, len(dd))
	dds := make([]uuid.UUID, len(dd))
	idx := 0
	for k := range ddc {
		ddc[k] = struct{}{}
		dds[idx] = k
		idx++
	}

	q := s.db.Queries()

	m, err := q.GetTimeseriesProjectMap(ctx, dds)
	if err != nil {
		return err
	}
	for tID := range ddc {
		ppID, ok := m[tID]
		if ok && ppID == projectID {
			delete(ddc, tID)
		}
	}
	if len(ddc) != 0 {
		return errors.New("instruments for all timeseries must be linked to project")
	}
	return nil
}
