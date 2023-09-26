package store

import (
	"context"

	"github.com/USACE/instrumentation-api/api/internal/model"
)

type EmailAutocomplereStore interface {
	ListEmailAutocomplete(ctx context.Context, emailInput string, limit int) ([]model.EmailAutocompleteResult, error)
}

type emailAutocompleteStore struct {
	db *model.Database
	q  *model.Queries
}

func NewEmailAutoCompleteStore(db *model.Database, q *model.Queries) *emailAutocompleteStore {
	return &emailAutocompleteStore{db, q}
}

// ListEmailAutocomplete returns search results for email autocomplete
func (s emailAutocompleteStore) ListEmailAutocomplete(ctx context.Context, emailInput string, limit int) ([]model.EmailAutocompleteResult, error) {
	return s.q.ListEmailAutocomplete(ctx, emailInput, limit)
}
