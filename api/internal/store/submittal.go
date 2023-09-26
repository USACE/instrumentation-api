package store

import (
	"context"

	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
)

type SubmittalStore interface {
}

type submittalStore struct {
	db *model.Database
	q  *model.Queries
}

func NewSubmittalStore(db *model.Database, q *model.Queries) *submittalStore {
	return &submittalStore{db, q}
}

func (s submittalStore) ListProjectSubmittals(ctx context.Context, projectID uuid.UUID, showMissing bool) ([]model.Submittal, error) {
	return s.q.ListProjectSubmittals(ctx, projectID, showMissing)
}

func (s submittalStore) ListInstrumentSubmittals(ctx context.Context, instrumentID uuid.UUID, showMissing bool) ([]model.Submittal, error) {
	return s.q.ListInstrumentSubmittals(ctx, instrumentID, showMissing)
}

func (s submittalStore) ListAlertConfigSubmittals(ctx context.Context, alertConfigID uuid.UUID, showMissing bool) ([]model.Submittal, error) {
	return s.q.ListAlertConfigSubmittals(ctx, alertConfigID, showMissing)
}

func (s submittalStore) ListUnverifiedMissingSubmittals(ctx context.Context) ([]model.Submittal, error) {
	return s.q.ListUnverifiedMissingSubmittals(ctx)
}

func (s submittalStore) UpdateSubmittal(ctx context.Context, sub model.Submittal) error {
	return s.q.UpdateSubmittal(ctx, sub)
}

func (s submittalStore) VerifyMissingSubmittal(ctx context.Context, submittalID uuid.UUID) error {
	return s.q.VerifyMissingSubmittal(ctx, submittalID)
}

func (s submittalStore) VerifyMissingAlertConfigSubmittals(ctx context.Context, alertConfigID uuid.UUID) error {
	return s.q.VerifyMissingAlertConfigSubmittals(ctx, alertConfigID)
}
