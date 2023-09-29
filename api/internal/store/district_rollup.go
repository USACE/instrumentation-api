package store

import (
	"context"

	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
)

type DistrictRollupStore interface {
	ListEvaluationDistrictRollup(ctx context.Context, opID uuid.UUID, tw model.TimeWindow) ([]model.DistrictRollup, error)
	ListMeasurementDistrictRollup(ctx context.Context, opID uuid.UUID, tw model.TimeWindow) ([]model.DistrictRollup, error)
}

type districtRollupStore struct {
	db *model.Database
	*model.Queries
}

func NewDistrictRollupStore(db *model.Database, q *model.Queries) *districtRollupStore {
	return &districtRollupStore{db, q}
}
