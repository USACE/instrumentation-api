package service

import (
	"context"

	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
)

type DistrictRollupService interface {
	ListEvaluationDistrictRollup(ctx context.Context, opID uuid.UUID, tw model.TimeWindow) ([]model.DistrictRollup, error)
	ListMeasurementDistrictRollup(ctx context.Context, opID uuid.UUID, tw model.TimeWindow) ([]model.DistrictRollup, error)
}

type districtRollupService struct {
	db *model.Database
	*model.Queries
}

func NewDistrictRollupService(db *model.Database, q *model.Queries) *districtRollupService {
	return &districtRollupService{db, q}
}
