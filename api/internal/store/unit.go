package store

import (
	"context"

	"github.com/USACE/instrumentation-api/api/internal/model"
)

type UnitStore interface {
	ListUnits(ctx context.Context) ([]model.Unit, error)
}

type unitStore struct {
	db *model.Database
	*model.Queries
}

func NewUnitStore(db *model.Database, q *model.Queries) *unitStore {
	return &unitStore{db, q}
}
