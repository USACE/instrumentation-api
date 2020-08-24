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
