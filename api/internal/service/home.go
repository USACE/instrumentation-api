package service

import (
	"context"

	"github.com/USACE/instrumentation-api/api/internal/model"
)

type HomeService interface {
	GetHome(ctx context.Context) (model.Home, error)
}

type homeService struct {
	db *model.Database
	*model.Queries
}

func NewHomeService(db *model.Database, q *model.Queries) *homeService {
	return &homeService{db, q}
}
