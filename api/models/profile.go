package models

import (
	"errors"
	"time"

	"github.com/USACE/instrumentation-api/api/passwords"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

// Profile is a user profile
type Profile struct {
	ID uuid.UUID `json:"id"`
	ProfileInfo
	Tokens  []TokenInfoProfile `json:"tokens"`
	IsAdmin bool               `json:"is_admin" db:"is_admin"`
	Roles   []string           `json:"roles"`
}

func (p *Profile) attachTokens(db *sqlx.DB) error {

	if err := db.Select(&p.Tokens, "SELECT token_id, issued FROM profile_token WHERE profile_id=$1", p.ID); err != nil {
		return err
	}
	return nil
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

// ProfilesFactory converts database rows to Profile objects
func ProfilesFactory(rows *sqlx.Rows) ([]Profile, error) {
	defer rows.Close()
	pp := make([]Profile, 0) // Profiles
	for rows.Next() {
		var p Profile
		p.Tokens = make([]TokenInfoProfile, 0)
		p.Roles = make([]string, 0)
		err := rows.Scan(
			&p.ID, &p.EDIPI, &p.Username, &p.Email, &p.IsAdmin, pq.Array(&p.Roles),
		)
		if err != nil {
			return make([]Profile, 0), err
		}
		// Add
		pp = append(pp, p)
	}

	return pp, nil
}

// GetProfileFromEDIPI returns a profile given an edipi
func GetProfileFromEDIPI(db *sqlx.DB, e int) (*Profile, error) {
	// Would prefer to do this in one query using a join and postgres json/array aggregation
	// for now it's implemented with two queries
	rows, err := db.Queryx("SELECT id, edipi, username, email, is_admin, roles FROM v_profile WHERE edipi=$1", e)
	if err != nil {
		return nil, err
	}
	pp, err := ProfilesFactory(rows)
	if err != nil {
		return nil, err
	}
	if len(pp) == 0 {
		return nil, errors.New("Profile Does Not Exist for User")
	}
	if err := pp[0].attachTokens(db); err != nil {
		return nil, err
	}
	return &pp[0], nil
}

// GetProfileFromTokenID returns a profile given a token ID
func GetProfileFromTokenID(db *sqlx.DB, tokenID string) (*Profile, error) {
	rows, err := db.Queryx(
		`SELECT p.id, p.edipi, p.username, p.email, p.is_admin
		 FROM profile_token t
		 LEFT JOIN v_profile p ON p.id = t.profile_id
		 WHERE t.token_id=$1`, tokenID,
	)
	if err != nil {
		return nil, err
	}
	pp, err := ProfilesFactory(rows)
	if err != nil {
		return nil, err
	}
	if err := pp[0].attachTokens(db); err != nil {
		return nil, err
	}
	return &pp[0], nil
}

// CreateProfile creates a new profile
func CreateProfile(db *sqlx.DB, n *ProfileInfo) (*Profile, error) {
	sql := "INSERT INTO profile (edipi, username, email) VALUES ($1, $2, $3) RETURNING id, username, email"
	p := Profile{
		Tokens: make([]TokenInfoProfile, 0),
		Roles:  make([]string, 0),
	}
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
		&n, "SELECT id, token_id, profile_id, issued, hash FROM profile_token WHERE token_id=$1 LIMIT 1", tokenID,
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
