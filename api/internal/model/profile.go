package model

import (
	"context"
	"fmt"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/passwords"
	"github.com/google/uuid"
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

func profileFactory(rows DBRows) ([]Profile, error) {
	defer rows.Close()
	pp := make([]Profile, 0)
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
		pp = append(pp, p)
	}
	return pp, nil
}

const getProfileFromEDIPI = `
	SELECT id, edipi, username, email, is_admin, roles FROM v_profile WHERE edipi = $1
`

// GetProfileFromEDIPI returns a profile given an edipi
func (q *Queries) GetProfileFromEDIPI(ctx context.Context, edipi int) (Profile, error) {
	var p Profile
	rows, err := q.db.Queryx(getProfileFromEDIPI, edipi)
	if err != nil {
		return p, err
	}
	pp, err := profileFactory(rows)
	if err != nil {
		return p, err
	}
	if len(pp) == 0 {
		return p, fmt.Errorf("Profile Does Not Exist for User")
	}
	return pp[0], nil
}

const getIssuedTokens = `
	SELECT token_id, issued FROM profile_token WHERE profile_id = $1
`

func (q *Queries) GetIssuedTokens(ctx context.Context, profileID uuid.UUID) ([]TokenInfoProfile, error) {
	tokens := make([]TokenInfoProfile, 0)
	err := q.db.SelectContext(ctx, &tokens, getIssuedTokens, profileID)
	return tokens, err
}

const getProfileFromTokenID = `
	SELECT p.id, p.edipi, p.username, p.email, p.is_admin
	FROM profile_token t
	LEFT JOIN v_profile p ON p.id = t.profile_id
	WHERE t.token_id = $1
`

func (q *Queries) GetProfileFromTokenID(ctx context.Context, tokenID string) (Profile, error) {
	var p Profile
	rows, err := q.db.QueryxContext(ctx, getProfileFromTokenID, tokenID)
	if err != nil {
		return p, err
	}
	pp, err := profileFactory(rows)
	if err != nil {
		return p, err
	}
	if len(pp) == 0 {
		return p, fmt.Errorf("Profile Does Not Exist for User")
	}
	return pp[0], nil
}

const createProfile = `
	INSERT INTO profile (edipi, username, email) VALUES ($1, $2, $3) RETURNING id, username, email
`

// CreateProfile creates a new profile
func (q *Queries) CreateProfile(ctx context.Context, n ProfileInfo) (Profile, error) {
	p := Profile{
		Tokens: make([]TokenInfoProfile, 0),
		Roles:  make([]string, 0),
	}
	err := q.db.GetContext(ctx, &p, createProfile, n.EDIPI, n.Username, n.Email)
	return p, err
}

const createProfileToken = `
	INSERT INTO profile_token (token_id, profile_id, hash) VALUES ($1,$2,$3) RETURNING *
`

// CreateProfileToken creates a secret token and stores the HASH (not the actual token)
// to the database. The return payload of this function is the first and last time you'll see
// the raw token unless the user writes it down or stores it somewhere safe.
func (q *Queries) CreateProfileToken(ctx context.Context, profileID uuid.UUID) (Token, error) {
	var t Token
	secretToken := passwords.GenerateRandom(40)
	tokenID := passwords.GenerateRandom(40)
	hash, err := passwords.CreateHash(secretToken, passwords.DefaultParams)
	if err != nil {
		return t, err
	}
	if err := q.db.GetContext(ctx, &t, createProfileToken, tokenID, profileID, hash); err != nil {
		return t, err
	}
	t.SecretToken = secretToken
	return t, nil
}

const getTokenInfoByTokenID = `
	SELECT id, token_id, profile_id, issued, hash FROM profile_token WHERE token_id=$1 LIMIT 1
`

// GetTokenInfoByTokenID returns a single token by token id
func (q *Queries) GetTokenInfoByTokenID(ctx context.Context, tokenID string) (TokenInfo, error) {
	var n TokenInfo
	err := q.db.GetContext(ctx, &n, getTokenInfoByTokenID, tokenID)
	return n, err
}

const deleteToken = `
	DELETE FROM profile_token WHERE profile_id=$1 AND token_id=$2
`

// DeleteToken deletes a token by token_id
func (q *Queries) DeleteToken(ctx context.Context, profileID uuid.UUID, tokenID string) error {
	_, err := q.db.ExecContext(ctx, deleteToken, profileID, tokenID)
	return err
}
