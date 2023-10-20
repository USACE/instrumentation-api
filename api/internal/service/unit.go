package service

import (
	"context"

	"github.com/USACE/instrumentation-api/api/internal/model"
)

type UnitService interface {
	ListUnits(ctx context.Context) ([]model.Unit, error)
}

type unitService struct {
	db *model.Database
	*model.Queries
}

func NewUnitService(db *model.Database, q *model.Queries) *unitService {
	return &unitService{db, q}
}
