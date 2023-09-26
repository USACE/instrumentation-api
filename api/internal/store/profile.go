package store

import (
	"context"
	"log"

	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
)

type ProfileStore interface {
}

type profileStore struct {
	db *model.Database
	q  *model.Queries
}

func NewProfileStore(db *model.Database, q *model.Queries) *profileStore {
	return &profileStore{db, q}
}

func (s profileStore) GetProfileWithTokensFromEDIPI(ctx context.Context, edipi int) (model.Profile, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return model.Profile{}, err
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			log.Print(err.Error())
		}
	}()

	qtx := s.q.WithTx(tx)

	p, err := qtx.GetProfileFromEDIPI(ctx, edipi)
	if err != nil {
		return model.Profile{}, err
	}
	tokens, err := qtx.GetIssuedTokens(ctx, p.ID)
	if err != nil {
		return model.Profile{}, err
	}
	if err := tx.Commit(); err != nil {
		return model.Profile{}, err
	}

	p.Tokens = tokens
	return p, nil
}

const getProfileFromTokenID = `
	SELECT p.id, p.edipi, p.username, p.email, p.is_admin
	FROM profile_token t
	LEFT JOIN v_profile p ON p.id = t.profile_id
	WHERE t.token_id=$1
`

// GetProfileFromTokenID returns a profile given a token ID
func (s profileStore) GetProfileFromTokenID(ctx context.Context, tokenID string) (model.Profile, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return model.Profile{}, err
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			log.Print(err.Error())
		}
	}()

	qtx := s.q.WithTx(tx)

	p, err := qtx.GetProfileFromTokenID(ctx, tokenID)
	if err != nil {
		return model.Profile{}, err
	}
	tokens, err := qtx.GetIssuedTokens(ctx, p.ID)
	if err != nil {
		return model.Profile{}, err
	}
	if err := tx.Commit(); err != nil {
		return model.Profile{}, err
	}

	p.Tokens = tokens
	return p, nil
}

// CreateProfile creates a new profile
func (s profileStore) CreateProfile(ctx context.Context, n model.ProfileInfo) (model.Profile, error) {
	return s.q.CreateProfile(ctx, n)
}

// CreateProfileToken creates a secret token and stores the HASH (not the actual token)
// to the database. The return payload of this function is the first and last time you'll see
// the raw token unless the user writes it down or stores it somewhere safe.
func (s profileStore) CreateProfileToken(ctx context.Context, profileID uuid.UUID) (model.Token, error) {
	return s.q.CreateProfileToken(ctx, profileID)
}

// GetTokenInfoByTokenID returns a single token by token id
func (s profileStore) GetTokenInfoByTokenID(ctx context.Context, tokenID string) (model.TokenInfo, error) {
	return s.q.GetTokenInfoByTokenID(ctx, tokenID)
}

// DeleteToken deletes a token by token_id
func (s profileStore) DeleteToken(ctx context.Context, profileID uuid.UUID, tokenID string) error {
	return s.q.DeleteToken(ctx, profileID, tokenID)
}
