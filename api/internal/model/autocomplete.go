package model

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
)

// EmailAutocompleteResult stores search result in profiles and emails
type EmailAutocompleteResult struct {
	ID       uuid.UUID `json:"id"`
	UserType string    `json:"user_type" db:"user_type"`
	Username *string   `json:"username"`
	Email    string    `json:"email"`
}

type EmailAutocompleteResultCollection []EmailAutocompleteResult

func (a *EmailAutocompleteResultCollection) Scan(src interface{}) error {
	if err := json.Unmarshal([]byte(src.(string)), a); err != nil {
		return err
	}
	return nil
}

// ListEmailAutocomplete returns search results for email autocomplete
func (q *Queries) ListEmailAutocomplete(ctx context.Context, emailInput *string, limit *int) ([]EmailAutocompleteResult, error) {
	c := `
		SELECT id, user_type, username, email
		FROM v_email_autocomplete
		WHERE username_email ILIKE '%'||$1||'%'
		LIMIT $2
	`
	aa := make([]EmailAutocompleteResult, 0)
	if err := q.db.SelectContext(ctx, &aa, c, emailInput, limit); err != nil {
		return make([]EmailAutocompleteResult, 0), err
	}
	return aa, nil
}
