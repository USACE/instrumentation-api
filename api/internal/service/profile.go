package service

import (
	"context"

	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
)

type ProfileService interface {
	GetProfileWithTokensForEDIPI(ctx context.Context, edipi int) (model.Profile, error)
	GetProfileWithTokensForEmail(ctx context.Context, email string) (model.Profile, error)
	GetProfileWithTokensForTokenID(ctx context.Context, tokenID string) (model.Profile, error)
	CreateProfile(ctx context.Context, n model.ProfileInfo) (model.Profile, error)
	CreateProfileToken(ctx context.Context, profileID uuid.UUID) (model.Token, error)
	GetTokenInfoByTokenID(ctx context.Context, tokenID string) (model.TokenInfo, error)
	UpdateProfileForEDIPI(ctx context.Context, username, email string, edipi int) error
	UpdateProfileForEmail(ctx context.Context, username, email string) error
	DeleteToken(ctx context.Context, profileID uuid.UUID, tokenID string) error
}

type profileService struct {
	db *model.Database
	*model.Queries
}

func NewProfileService(db *model.Database, q *model.Queries) *profileService {
	return &profileService{db, q}
}

func (s profileService) GetProfileWithTokensForEDIPI(ctx context.Context, edipi int) (model.Profile, error) {
	p, err := s.GetProfileForEDIPI(ctx, edipi)
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

func (s profileService) GetProfileWithTokensForEmail(ctx context.Context, email string) (model.Profile, error) {
	p, err := s.GetProfileForEmail(ctx, email)
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

// GetProfileForTokenID returns a profile given a token ID
func (s profileService) GetProfileWithTokensForTokenID(ctx context.Context, tokenID string) (model.Profile, error) {
	p, err := s.GetProfileForTokenID(ctx, tokenID)
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
