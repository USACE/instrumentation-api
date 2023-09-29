package store

import (
	"context"

	"github.com/USACE/instrumentation-api/api/internal/model"
)

type OpendcsStore interface {
	ListOpendcsSites(ctx context.Context) ([]model.Site, error)
}

type opendcsStore struct {
	db *model.Database
	*model.Queries
}

func NewOpendcsStore(db *model.Database, q *model.Queries) *opendcsStore {
	return &opendcsStore{db, q}
}
