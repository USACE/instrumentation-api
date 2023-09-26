package store

import (
	"context"
	"log"

	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
)

type TimeseriesStore interface {
}

type timeseriesStore struct {
	db *model.Database
	q  *model.Queries
}

func NewTimeseriesStore(db *model.Database, q *model.Queries) *timeseriesStore {
	return &timeseriesStore{db, q}
}

// ListTimeseries lists all timeseries
func (s timeseriesStore) ListTimeseries(ctx context.Context) ([]model.Timeseries, error) {
	return s.q.ListTimeseries(ctx)
}

// ValidateStoredTimeseries returns an error if the timeseries id does not exist or the timeseries is computed
func (s timeseriesStore) GetStoredTimeseriesExists(ctx context.Context, timeseriesID uuid.UUID) (bool, error) {
	return s.q.GetStoredTimeseriesExists(ctx, timeseriesID)
}

// ListTimeseriesSlugs lists used timeseries slugs in the database
func (s timeseriesStore) ListTimeseriesSlugs(ctx context.Context) ([]string, error) {
	return s.q.ListTimeseriesSlugs(ctx)
}

// GetTimeseriesProjectMap returns a map of { timeseries_id: project_id, }
func (s timeseriesStore) GetTimeseriesProjectMap(ctx context.Context, timeseriesIDs []uuid.UUID) (map[uuid.UUID]uuid.UUID, error) {
	return s.q.GetTimeseriesProjectMap(ctx, timeseriesIDs)
}

// ListTimeseriesSlugsForInstrument lists used timeseries slugs for a given instrument
func (s timeseriesStore) ListTimeseriesSlugsForInstrument(ctx context.Context, instrumentID uuid.UUID) ([]string, error) {
	return s.q.ListTimeseriesSlugsForInstrument(ctx, instrumentID)
}

// ListProjectTimeseries lists all timeseries for a given project
func (s timeseriesStore) ListProjectTimeseries(ctx context.Context, projectID uuid.UUID) ([]model.Timeseries, error) {
	return s.q.ListProjectTimeseries(ctx, projectID)
}

// ListInstrumentTimeseries returns an array of timeseries for an instrument
func (s timeseriesStore) ListInstrumentTimeseries(ctx context.Context, instrumentID uuid.UUID) ([]model.Timeseries, error) {
	return s.q.ListInstrumentTimeseries(ctx, instrumentID)
}

// ListInstrumentGroupTimeseries returns an array of timeseries for instruments that belong to an instrument_group
func (s timeseriesStore) ListInstrumentGroupTimeseries(ctx context.Context, instrumentGroupID uuid.UUID) ([]model.Timeseries, error) {
	return s.q.ListInstrumentGroupTimeseries(ctx, instrumentGroupID)
}

// GetTimeseries returns a single timeseries without measurements
func (s timeseriesStore) GetTimeseries(ctx context.Context, timeseriesID uuid.UUID) (model.Timeseries, error) {
	return s.q.GetTimeseries(ctx, timeseriesID)
}

// CreateTimeseries creates many timeseries from an array of timeseries
func (s timeseriesStore) CreateTimeseries(ctx context.Context, ts model.Timeseries) (model.Timeseries, error) {
	return s.q.CreateTimeseries(ctx, ts)
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

	qtx := s.q.WithTx(tx)

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

// UpdateTimeseries updates a timeseries
func (s timeseriesStore) UpdateTimeseries(ctx context.Context, ts model.Timeseries) (uuid.UUID, error) {
	return s.q.UpdateTimeseries(ctx, ts)
}

// DeleteTimeseries deletes a timeseries and cascade deletes all measurements
func (s timeseriesStore) DeleteTimeseries(ctx context.Context, timeseriesID uuid.UUID) error {
	return s.q.DeleteTimeseries(ctx, timeseriesID)
}
