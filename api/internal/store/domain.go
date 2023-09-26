package store

import (
	"context"

	"github.com/USACE/instrumentation-api/api/internal/model"
)

type DomainStore interface {
}

type domainStore struct {
	db *model.Database
	q  *model.Queries
}

func NewDomainStore(db *model.Database, q *model.Queries) *districtRollupStore {
	return &districtRollupStore{db, q}
}

// ListCollectionGroups lists all collection groups for a project
func (s domainStore) GetDomains(ctx context.Context) ([]model.Domain, error) {
	return s.q.GetDomains(ctx)
}
