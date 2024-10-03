package service

import (
	"context"
	"errors"
	"strings"

	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
)

type ProfileService interface {
	GetProfileWithTokensForClaims(ctx context.Context, claims model.ProfileClaims) (model.Profile, error)
	GetProfileWithTokensForEDIPI(ctx context.Context, edipi int) (model.Profile, error)
	GetProfileForEmail(ctx context.Context, email string) (model.Profile, error)
	GetProfileWithTokensForUsername(ctx context.Context, username string) (model.Profile, error)
	GetProfileWithTokensForTokenID(ctx context.Context, tokenID string) (model.Profile, error)
	CreateProfile(ctx context.Context, n model.ProfileInfo) (model.Profile, error)
	CreateProfileToken(ctx context.Context, profileID uuid.UUID) (model.Token, error)
	GetTokenInfoByTokenID(ctx context.Context, tokenID string) (model.TokenInfo, error)
	UpdateProfileForClaims(ctx context.Context, p model.Profile, claims model.ProfileClaims) (model.Profile, error)
	DeleteToken(ctx context.Context, profileID uuid.UUID, tokenID string) error
}

type profileService struct {
	db *model.Database
	*model.Queries
}

func NewProfileService(db *model.Database, q *model.Queries) *profileService {
	return &profileService{db, q}
}

func (s profileService) GetProfileWithTokensForClaims(ctx context.Context, claims model.ProfileClaims) (model.Profile, error) {
	var p model.Profile
	var err error
	if claims.CacUID != nil {
		p, err = s.GetProfileWithTokensForEDIPI(ctx, *claims.CacUID)
	} else {
		p, err = s.GetProfileWithTokensForEmail(ctx, claims.Email)
	}
	if err != nil {
		return model.Profile{}, err
	}
	return p, nil
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

func (s profileService) GetProfileWithTokensForUsername(ctx context.Context, username string) (model.Profile, error) {
	p, err := s.GetProfileForUsername(ctx, username)
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

// UpdateProfileForClaims syncs a database profile to the provided token claims
// THe order of precence in which the function will attepmt to update profiles is edipi, email, username
func (s profileService) UpdateProfileForClaims(ctx context.Context, p model.Profile, claims model.ProfileClaims) (model.Profile, error) {
	var claimsMatchProfile bool = p.Username == claims.PreferredUsername &&
		strings.ToLower(p.Email) == strings.ToLower(claims.Email) &&
		p.DisplayName == claims.Name

	if claimsMatchProfile {
		return p, nil
	}

	if claims.CacUID != nil && !claimsMatchProfile {
		if err := s.UpdateProfileForEDIPI(ctx, *claims.CacUID, model.ProfileInfo{
			Username:    claims.PreferredUsername,
			DisplayName: claims.Name,
			Email:       claims.Email,
		}); err != nil {
			return p, err
		}
		p.Username = claims.PreferredUsername
		p.DisplayName = claims.Name
		p.Email = claims.Email

		return p, nil
	}

	if strings.ToLower(p.Email) == strings.ToLower(claims.Email) && !claimsMatchProfile {
		if err := s.UpdateProfileForEmail(ctx, claims.Email, model.ProfileInfo{
			Username:    claims.PreferredUsername,
			DisplayName: claims.Name,
		}); err != nil {
			return p, err
		}
		p.Username = claims.PreferredUsername
		p.DisplayName = claims.Name

		return p, nil
	}

	return p, errors.New("claims did not match profile and could not be updated")
}
