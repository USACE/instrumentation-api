package service

import (
	"context"

	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
)

type ProfileService interface {
	GetProfileWithTokensFromEDIPI(ctx context.Context, edipi int) (model.Profile, error)
	GetProfileWithTokensFromTokenID(ctx context.Context, tokenID string) (model.Profile, error)
	CreateProfile(ctx context.Context, n model.ProfileInfo) (model.Profile, error)
	CreateProfileToken(ctx context.Context, profileID uuid.UUID) (model.Token, error)
	GetTokenInfoByTokenID(ctx context.Context, tokenID string) (model.TokenInfo, error)
	UpdateSubForEDIPI(ctx context.Context, sub uuid.UUID, edipi int) error
	DeleteToken(ctx context.Context, profileID uuid.UUID, tokenID string) error
}

type profileService struct {
	db *model.Database
	*model.Queries
}

func NewProfileService(db *model.Database, q *model.Queries) *profileService {
	return &profileService{db, q}
}

func (s profileService) GetProfileWithTokensFromEDIPI(ctx context.Context, edipi int) (model.Profile, error) {
	p, err := s.GetProfileFromEDIPI(ctx, edipi)
	if err != nil {
		return model.Profile{}, err
	}
	tokens, err := s.GetIssuedTokens(ctx, p.ID)
	if err != nil {
		return model.Profile{}, err
	}
	p.Tokens = tokens
	return p, nil
}

// GetProfileFromTokenID returns a profile given a token ID
func (s profileService) GetProfileWithTokensFromTokenID(ctx context.Context, tokenID string) (model.Profile, error) {
	p, err := s.GetProfileFromTokenID(ctx, tokenID)
	if err != nil {
		return model.Profile{}, err
	}
	tokens, err := s.GetIssuedTokens(ctx, p.ID)
	if err != nil {
		return model.Profile{}, err
	}
	p.Tokens = tokens
	return p, nil
}
