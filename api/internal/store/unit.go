package store

import (
	"context"

	"github.com/USACE/instrumentation-api/api/internal/model"
)

type UnitStore interface {
}

type unitStore struct {
	db *model.Database
	q  *model.Queries
}

func NewUnitStore(db *model.Database, q *model.Queries) *unitStore {
	return &unitStore{db, q}
}

// ListUnits returns a slice of units
func (s unitStore) ListUnits(ctx context.Context) ([]model.Unit, error) {
	return s.q.ListUnits(ctx)
}
