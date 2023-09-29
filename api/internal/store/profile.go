package store

import (
	"context"
	"log"

	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
)

type ProfileStore interface {
	GetProfileWithTokensFromEDIPI(ctx context.Context, edipi int) (model.Profile, error)
	GetProfileFromTokenID(ctx context.Context, tokenID string) (model.Profile, error)
	CreateProfile(ctx context.Context, n model.ProfileInfo) (model.Profile, error)
	CreateProfileToken(ctx context.Context, profileID uuid.UUID) (model.Token, error)
	GetTokenInfoByTokenID(ctx context.Context, tokenID string) (model.TokenInfo, error)
	DeleteToken(ctx context.Context, profileID uuid.UUID, tokenID string) error
}

type profileStore struct {
	db *model.Database
	*model.Queries
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

	qtx := s.WithTx(tx)

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

	qtx := s.WithTx(tx)

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
