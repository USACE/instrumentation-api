package store

import (
	"context"

	"github.com/USACE/instrumentation-api/api/internal/model"
)

type DomainStore interface {
	GetDomains(ctx context.Context) ([]model.Domain, error)
}

type domainStore struct {
	db *model.Database
	*model.Queries
}

func NewDomainStore(db *model.Database, q *model.Queries) *domainStore {
	return &domainStore{db, q}
}
