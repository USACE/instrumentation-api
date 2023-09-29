package store

import (
	"context"

	"github.com/USACE/instrumentation-api/api/internal/model"
)

type EmailAutocompleteStore interface {
	ListEmailAutocomplete(ctx context.Context, emailInput string, limit int) ([]model.EmailAutocompleteResult, error)
}

type emailAutocompleteStore struct {
	db *model.Database
	*model.Queries
}

func NewEmailAutocompleteStore(db *model.Database, q *model.Queries) *emailAutocompleteStore {
	return &emailAutocompleteStore{db, q}
}
