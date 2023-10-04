package service

import (
	"context"

	"github.com/USACE/instrumentation-api/api/internal/model"
)

type DomainService interface {
	GetDomains(ctx context.Context) ([]model.Domain, error)
}

type domainService struct {
	db *model.Database
	*model.Queries
}

func NewDomainService(db *model.Database, q *model.Queries) *domainService {
	return &domainService{db, q}
}
