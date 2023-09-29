package store

import (
	"context"

	"github.com/USACE/instrumentation-api/api/internal/model"
)

type HeartbeatStore interface {
	DoHeartbeat(ctx context.Context) (model.Heartbeat, error)
	GetLatestHeartbeat(ctx context.Context) (model.Heartbeat, error)
	ListHeartbeats(ctx context.Context) ([]model.Heartbeat, error)
}

type heartbeatStore struct {
	db *model.Database
	*model.Queries
}

func NewHeartbeatStore(db *model.Database, q *model.Queries) *heartbeatStore {
	return &heartbeatStore{db, q}
}
