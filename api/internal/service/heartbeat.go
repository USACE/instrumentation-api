package service

import (
	"context"

	"github.com/USACE/instrumentation-api/api/internal/model"
)

type HeartbeatService interface {
	DoHeartbeat(ctx context.Context) (model.Heartbeat, error)
	GetLatestHeartbeat(ctx context.Context) (model.Heartbeat, error)
	ListHeartbeats(ctx context.Context) ([]model.Heartbeat, error)
}

type heartbeatService struct {
	db *model.Database
	*model.Queries
}

func NewHeartbeatService(db *model.Database, q *model.Queries) *heartbeatService {
	return &heartbeatService{db, q}
}
