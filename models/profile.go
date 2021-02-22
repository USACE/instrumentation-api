package models

import (
	"time"

	"github.com/USACE/instrumentation-api/passwords"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// Profile is a user profile
type Profile struct {
	ID uuid.UUID `json:"id"`
	ProfileInfo
	Tokens []TokenInfoProfile `json:"tokens"`
}

// TokenInfoProfile is token information embedded in Profile
type TokenInfoProfile struct {
	TokenID string    `json:"token_id" db:"token_id"`
	Issued  time.Time `json:"issued"`
}

// ProfileInfo is information necessary to construct a profile
type ProfileInfo struct {
	EDIPI    int    `json:"-"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// TokenInfo represents the information held in the database about a token
type TokenInfo struct {
	ID        uuid.UUID `json:"-"`
	TokenID   string    `json:"token_id" db:"token_id"`
	ProfileID uuid.UUID `json:"profile_id" db:"profile_id"`
	Issued    time.Time `json:"issued"`
	Hash      string    `json:"-"`
}

// Token includes all TokenInfo and the actual token string generated for a user
// this is only returned the first time a token is generated
type Token struct {
	SecretToken string `json:"secret_token"`
	TokenInfo
}

// GetProfileFromEDIPI returns a profile given an edipi
func GetProfileFromEDIPI(db *sqlx.DB, e int) (*Profile, error) {
	// Would prefer to do this in one query using a join and postgres json/array aggregation
	// for now it's implemented with two queries
	var p Profile
	if err := db.Get(&p, "SELECT id, edipi, username, email FROM profile WHERE edipi=$1", e); err != nil {
		return nil, err
	}
	p.Tokens = make([]TokenInfoProfile, 0)
	if err := db.Select(&p.Tokens, "SELECT token_id, issued FROM profile_token WHERE profile_id=$1", p.ID); err != nil {
		return nil, err
	}
	return &p, nil
}

// GetProfileFromTokenID returns a profile given a token ID
func GetProfileFromTokenID(db *sqlx.DB, tokenID string) (*Profile, error) {
	var p Profile
	sql := `SELECT p.id, p.edipi, p.username, p.email
			FROM profile_token t
			LEFT JOIN profile p ON p.id = t.profile_id
			WHERE t.token_id=$1`
	if err := db.Get(&p, sql, tokenID); err != nil {
		return nil, err
	}
	return &p, nil
}

// CreateProfile creates a new profile
func CreateProfile(db *sqlx.DB, n *ProfileInfo) (*Profile, error) {
	sql := "INSERT INTO profile (edipi, username, email) VALUES ($1, $2, $3) RETURNING id, username, email"
	var p Profile
	if err := db.Get(&p, sql, n.EDIPI, n.Username, n.Email); err != nil {
		return nil, err
	}
	return &p, nil
}

// CreateProfileToken creates a secret token and stores the HASH (not the actual token)
// to the database. The return payload of this function is the first and last time you'll see
// the raw token unless the user writes it down or stores it somewhere safe.
func CreateProfileToken(db *sqlx.DB, profileID *uuid.UUID) (*Token, error) {
	secretToken := passwords.GenerateRandom(40)
	tokenID := passwords.GenerateRandom(40)

	hash, err := passwords.CreateHash(secretToken, passwords.DefaultParams)
	if err != nil {
		return nil, err
	}
	var t Token
	if err := db.Get(
		&t, "INSERT INTO profile_token (token_id, profile_id, hash) VALUES ($1,$2,$3) RETURNING *",
		tokenID, profileID, hash,
	); err != nil {
		return nil, err
	}
	t.SecretToken = secretToken
	return &t, nil
}

// GetTokenInfoByTokenID returns a single token by token id
func GetTokenInfoByTokenID(db *sqlx.DB, tokenID *string) (*TokenInfo, error) {
	var n TokenInfo
	if err := db.Get(
		&n, "SELECT * FROM profile_token WHERE token_id=$1 LIMIT 1", tokenID,
	); err != nil {
		return nil, err
	}
	return &n, nil
}

// DeleteToken deletes a token by token_id
func DeleteToken(db *sqlx.DB, profileID *uuid.UUID, tokenID *string) error {
	sql := "DELETE FROM profile_token WHERE profile_id=$1 AND token_id=$2"
	_, err := db.Exec(sql, profileID, tokenID)
	if err != nil {
		return err
	}
	return nil
}
