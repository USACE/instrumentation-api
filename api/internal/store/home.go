package store

import (
	"context"

	"github.com/USACE/instrumentation-api/api/internal/model"
)

type HomeStore interface {
	GetHome(ctx context.Context) (model.Home, error)
}

type homeStore struct {
	db *model.Database
	*model.Queries
}

func NewHomeStore(db *model.Database, q *model.Queries) *homeStore {
	return &homeStore{db, q}
}
