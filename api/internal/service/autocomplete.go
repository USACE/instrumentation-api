package service

import (
	"context"

	"github.com/USACE/instrumentation-api/api/internal/model"
)

type EmailAutocompleteService interface {
	ListEmailAutocomplete(ctx context.Context, emailInput string, limit int) ([]model.EmailAutocompleteResult, error)
}

type emailAutocompleteService struct {
	db *model.Database
	*model.Queries
}

func NewEmailAutocompleteService(db *model.Database, q *model.Queries) *emailAutocompleteService {
	return &emailAutocompleteService{db, q}
}
