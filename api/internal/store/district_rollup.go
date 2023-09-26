package store

import (
	"context"

	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
)

type DistrictRollupStore interface {
}

type districtRollupStore struct {
	db *model.Database
	q  *model.Queries
}

func NewDistrictRollupStore(db *model.Database, q *model.Queries) *districtRollupStore {
	return &districtRollupStore{db, q}
}

// ListCollectionGroups lists all collection groups for a project
func (s districtRollupStore) ListEvaluationDistrictRollup(ctx context.Context, opID uuid.UUID, tw model.TimeWindow) ([]model.DistrictRollup, error) {
	return s.q.ListEvaluationDistrictRollup(ctx, opID, tw)
}

// ListCollectionGroups lists all collection groups for a project
func (s districtRollupStore) ListMeasurementDistrictRollup(ctx context.Context, opID uuid.UUID, tw model.TimeWindow) ([]model.DistrictRollup, error) {
	return s.q.ListEvaluationDistrictRollup(ctx, opID, tw)
}
