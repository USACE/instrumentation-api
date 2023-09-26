package store

import (
	"context"

	"github.com/USACE/instrumentation-api/api/internal/model"
)

type ProcessTimeseriesStore interface {
}

type processTimeseriesStore struct {
	db *model.Database
	q  *model.Queries
}

func NewProcessTimeseriesStore(db *model.Database, q *model.Queries) *processTimeseriesStore {
	return &processTimeseriesStore{db, q}
}

// SelectMeasurements returns measurements for the timeseries specified in the filter
func (s processTimeseriesStore) SelectMeasurements(ctx context.Context, f model.ProcessMeasurementFilter) (model.ProcessTimeseriesResponseCollection, error) {
	return s.q.SelectMeasurements(ctx, f)
}
