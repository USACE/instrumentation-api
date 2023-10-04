package service

import (
	"context"

	"github.com/USACE/instrumentation-api/api/internal/model"
)

type OpendcsService interface {
	ListOpendcsSites(ctx context.Context) ([]model.Site, error)
}

type opendcsService struct {
	db *model.Database
	*model.Queries
}

func NewOpendcsService(db *model.Database, q *model.Queries) *opendcsService {
	return &opendcsService{db, q}
}
