package store

import (
	"context"

	"github.com/USACE/instrumentation-api/api/internal/model"
)

type OpendcsStore interface {
}

type opendcsStore struct {
	db *model.Database
	q  *model.Queries
}

func NewOpendcsStore(db *model.Database, q *model.Queries) *opendcsStore {
	return &opendcsStore{db, q}
}

// ListOpendcsSites returns an array of instruments from the database
// And formats them as OpenDCS Sites
func (s opendcsStore) ListOpendcsSites(ctx context.Context) ([]model.Site, error) {
	return s.ListOpendcsSites(ctx)
}
