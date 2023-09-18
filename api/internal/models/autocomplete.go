package models

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
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
func ListEmailAutocomplete(db *sqlx.DB, str *string, limit *int) ([]EmailAutocompleteResult, error) {

	rr := make([]EmailAutocompleteResult, 0)
	sql := `SELECT id, user_type, username, email
			FROM v_email_autocomplete
			WHERE username_email ILIKE '%'||$1||'%'
			LIMIT $2
	`
	if err := db.Select(&rr, sql, str, limit); err != nil {
		return make([]EmailAutocompleteResult, 0), err
	}
	return rr, nil
}
