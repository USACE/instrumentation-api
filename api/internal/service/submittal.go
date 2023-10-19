package service

import (
	"context"

	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
)

type SubmittalService interface {
	ListProjectSubmittals(ctx context.Context, projectID uuid.UUID, showMissing bool) ([]model.Submittal, error)
	ListInstrumentSubmittals(ctx context.Context, instrumentID uuid.UUID, showMissing bool) ([]model.Submittal, error)
	ListAlertConfigSubmittals(ctx context.Context, alertConfigID uuid.UUID, showMissing bool) ([]model.Submittal, error)
	ListUnverifiedMissingSubmittals(ctx context.Context) ([]model.Submittal, error)
	UpdateSubmittal(ctx context.Context, sub model.Submittal) error
	VerifyMissingSubmittal(ctx context.Context, submittalID uuid.UUID) error
	VerifyMissingAlertConfigSubmittals(ctx context.Context, alertConfigID uuid.UUID) error
}

type submittalService struct {
	db *model.Database
	*model.Queries
}

func NewSubmittalService(db *model.Database, q *model.Queries) *submittalService {
	return &submittalService{db, q}
}
