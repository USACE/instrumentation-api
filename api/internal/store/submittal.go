package store

import (
	"context"

	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
)

type SubmittalStore interface {
	ListProjectSubmittals(ctx context.Context, projectID uuid.UUID, showMissing bool) ([]model.Submittal, error)
	ListInstrumentSubmittals(ctx context.Context, instrumentID uuid.UUID, showMissing bool) ([]model.Submittal, error)
	ListAlertConfigSubmittals(ctx context.Context, alertConfigID uuid.UUID, showMissing bool) ([]model.Submittal, error)
	ListUnverifiedMissingSubmittals(ctx context.Context) ([]model.Submittal, error)
	UpdateSubmittal(ctx context.Context, sub model.Submittal) error
	VerifyMissingSubmittal(ctx context.Context, submittalID uuid.UUID) error
	VerifyMissingAlertConfigSubmittals(ctx context.Context, alertConfigID uuid.UUID) error
}

type submittalStore struct {
	db *model.Database
	*model.Queries
}

func NewSubmittalStore(db *model.Database, q *model.Queries) *submittalStore {
	return &submittalStore{db, q}
}
