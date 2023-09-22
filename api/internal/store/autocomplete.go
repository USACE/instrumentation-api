package store

import (
	"context"

	"github.com/USACE/instrumentation-api/api/internal/model"
)

type EmailAutocomplereStore interface {
	ListEmailAutocomplete(ctx context.Context, emailInput *string, limit *int) ([]model.EmailAutocompleteResult, error)
}

type emailAutocompleteStore struct {
	db *model.Database
}

func NewEmailAutoCompleteStore(db *model.Database) *emailAutocompleteStore {
	return &emailAutocompleteStore{db}
}

// ListEmailAutocomplete returns search results for email autocomplete
func (s emailAutocompleteStore) ListEmailAutocomplete(ctx context.Context, emailInput *string, limit *int) ([]model.EmailAutocompleteResult, error) {
	q := model.NewQueries(s.db)

	rr, err := q.ListEmailAutocomplete(ctx, emailInput, limit)
	if err != nil {
		return rr, err
	}
	return rr, nil
}
