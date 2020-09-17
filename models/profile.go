package models

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// Profile is a user profile
type Profile struct {
	ID    uuid.UUID `json:"id"`
	EDIPI int       `json:"edipi"`
	ProfileData
}

// ProfileData holds the incoming data for a create profile request
type ProfileData struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

// CreateProfileRequest holds the data required to create a profile
type CreateProfileRequest struct {
	Action
	ProfileData
}

// Email is like a profile, but for a user that will never log-in.
// Useful for subscribing any email to alerts
type Email struct {
	ID    uuid.UUID `json:"id"`
	Email string    `json:"email"`
}

// EmailAutocompleteResult stores search result in profiles and emails
type EmailAutocompleteResult struct {
	ID       uuid.UUID `json:"id"`
	UserType string    `json:"user_type" db:"user_type"`
	Username *string   `json:"username"`
	Email    string    `json:"email"`
}

// GetMyProfile returns profile
func GetMyProfile(db *sqlx.DB, action *Action) (*Profile, error) {
	var p Profile
	if err := db.Get(&p, "SELECT * FROM profile WHERE edipi=$1", action.Actor); err != nil {
		return nil, err
	}
	return &p, nil
}

// CreateProfile creates a new profile
func CreateProfile(db *sqlx.DB, r *CreateProfileRequest) (*Profile, error) {
	sql := "INSERT INTO profile (edipi, username, email) VALUES ($1, $2, $3) RETURNING *"
	var p Profile
	if err := db.Get(&p, sql, r.Actor, r.Username, r.Email); err != nil {
		return nil, err
	}
	return &p, nil
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
