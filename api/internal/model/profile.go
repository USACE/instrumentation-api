package model

import (
	"context"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/password"
	"github.com/google/uuid"
)

// Profile is a user profile
type Profile struct {
	ID uuid.UUID `json:"id"`
	ProfileInfo
	Tokens  []TokenInfoProfile `json:"tokens"`
	IsAdmin bool               `json:"is_admin" db:"is_admin"`
	Roles   dbSlice[string]    `json:"roles" db:"roles"`
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

const getProfileForEDIPI = `
	SELECT * FROM v_profile WHERE edipi = $1
`

func (q *Queries) GetProfileForEDIPI(ctx context.Context, edipi int) (Profile, error) {
	var p Profile
	err := q.db.GetContext(ctx, &p, getProfileForEDIPI, edipi)
	return p, err
}

const getProfileForEmail = `
	SELECT * FROM v_profile WHERE email = $1
`

func (q *Queries) GetProfileForEmail(ctx context.Context, email string) (Profile, error) {
	var p Profile
	err := q.db.GetContext(ctx, &p, getProfileForEDIPI, email)
	return p, err
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
	err := q.db.GetContext(ctx, getProfileFromTokenID, tokenID)
	return p, err
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
	secretToken := password.GenerateRandom(40)
	tokenID := password.GenerateRandom(40)
	hash, err := password.CreateHash(secretToken, password.DefaultParams)
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

const updateProfileForEDIPI = `UPDATE profile SET username=$1, email=$2 WHERE edipi=$3`

func (q *Queries) UpdateProfileForEDIPI(ctx context.Context, username, email string, edipi int) error {
	_, err := q.db.ExecContext(ctx, updateProfileForEDIPI, username, email, edipi)
	return err
}

const updateProfileForEmail = `UPDATE profile SET username=$1 WHERE email=$2`

func (q *Queries) UpdateProfileForEmail(ctx context.Context, username, email string) error {
	_, err := q.db.ExecContext(ctx, updateProfileForEmail, username, email)
	return err
}

const deleteToken = `
	DELETE FROM profile_token WHERE profile_id=$1 AND token_id=$2
`

// DeleteToken deletes a token by token_id
func (q *Queries) DeleteToken(ctx context.Context, profileID uuid.UUID, tokenID string) error {
	_, err := q.db.ExecContext(ctx, deleteToken, profileID, tokenID)
	return err
}
