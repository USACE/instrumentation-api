package service

import (
	"context"

	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
)

type Survey123Service interface {
	CreateSurvey123Preview(ctx context.Context, survey123EquivalencyTableID uuid.UUID, previewRawJson []byte) error
	CreateOrUpdateSurvey123Measurements(ctx context.Context, sp model.Survey123Payload) error
}

type survey123Service struct {
	db *model.Database
	*model.Queries
}

func NewSurvey123Service(db *model.Database, q *model.Queries) *survey123Service {
	return &survey123Service{db, q}
}

func (s survey123Service) CreateSurvey123Preview(ctx context.Context, survey123EquivalencyTableID uuid.UUID, previewRawJson []byte) error {

	return nil
}

func (s survey123Service) CreateOrUpdateSurvey123Measurements(ctx context.Context, sp model.Survey123Payload) error {
	return nil
}
