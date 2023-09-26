package store

import (
	"context"

	"github.com/USACE/instrumentation-api/api/internal/model"
)

type HeartbeatStore interface {
}

type heartbeatStore struct {
	db *model.Database
	q  *model.Queries
}

func NewHeartbeatStore(db *model.Database, q *model.Queries) *heartbeatStore {
	return &heartbeatStore{db, q}
}

// DoHeartbeat does regular-interval tasks
func (s heartbeatStore) DoHeartbeat(ctx context.Context) (model.Heartbeat, error) {
	return s.q.DoHeartbeat(ctx)
}

// GetLatestHeartbeat returns the most recent system heartbeat
func (s heartbeatStore) GetLatestHeartbeat(ctx context.Context) (model.Heartbeat, error) {
	return s.q.GetLatestHeartbeat(ctx)
}

// ListHeartbeats returns all system heartbeats
func (s heartbeatStore) ListHeartbeats(ctx context.Context) ([]model.Heartbeat, error) {
	return s.q.ListHeartbeats(ctx)
}
