package store

import (
	"context"

	"github.com/USACE/instrumentation-api/api/internal/model"
)

type ProcessTimeseriesStore interface {
	SelectMeasurements(ctx context.Context, f model.ProcessMeasurementFilter) (model.ProcessTimeseriesResponseCollection, error)
}

type processTimeseriesStore struct {
	db *model.Database
	*model.Queries
}

func NewProcessTimeseriesStore(db *model.Database, q *model.Queries) *processTimeseriesStore {
	return &processTimeseriesStore{db, q}
}
