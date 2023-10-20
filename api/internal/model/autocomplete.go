package model

import (
	"context"

	"github.com/google/uuid"
)

// EmailAutocompleteResult stores search result in profiles and emails
type EmailAutocompleteResult struct {
	ID       uuid.UUID `json:"id"`
	UserType string    `json:"user_type" db:"user_type"`
	Username *string   `json:"username"`
	Email    string    `json:"email"`
}

const listEmailAutocomplete = `
	SELECT id, user_type, username, email
	FROM v_email_autocomplete
	WHERE username_email ILIKE '%'||$1||'%'
	LIMIT $2
`

// ListEmailAutocomplete returns search results for email autocomplete
func (q *Queries) ListEmailAutocomplete(ctx context.Context, emailInput string, limit int) ([]EmailAutocompleteResult, error) {
	aa := make([]EmailAutocompleteResult, 0)
	if err := q.db.SelectContext(ctx, &aa, listEmailAutocomplete, emailInput, limit); err != nil {
		return nil, err
	}
	return aa, nil
}
