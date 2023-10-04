package service

import (
	"context"

	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
)

type AlertService interface {
	CreateAlerts(ctx context.Context, alertConfigIDs []uuid.UUID) error
	GetAllAlertsForProject(ctx context.Context, projectID uuid.UUID) ([]model.Alert, error)
	GetAllAlertsForInstrument(ctx context.Context, instrumentID uuid.UUID) ([]model.Alert, error)
	GetAllAlertsForProfile(ctx context.Context, profileID uuid.UUID) ([]model.Alert, error)
	GetOneAlertForProfile(ctx context.Context, profileID uuid.UUID, alertID uuid.UUID) (model.Alert, error)
	DoAlertRead(ctx context.Context, profileID uuid.UUID, alertID uuid.UUID) (model.Alert, error)
	DoAlertUnread(ctx context.Context, profileID uuid.UUID, alertID uuid.UUID) (model.Alert, error)
}

type alertService struct {
	db *model.Database
	*model.Queries
}

func NewAlertService(db *model.Database, q *model.Queries) *alertService {
	return &alertService{db, q}
}

// Create creates one or more new alerts
func (s alertService) CreateAlerts(ctx context.Context, alertConfigIDs []uuid.UUID) error {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)
	for _, id := range alertConfigIDs {
		if err := qtx.CreateAlerts(ctx, id); err != nil {
			return err
		}
	}
	return tx.Commit()
}

// DoAlertRead marks an alert as read for a profile
func (s alertService) DoAlertRead(ctx context.Context, profileID, alertID uuid.UUID) (model.Alert, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return model.Alert{}, err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)
	if err := qtx.DoAlertRead(ctx, profileID, alertID); err != nil {
		return model.Alert{}, err
	}
	b, err := qtx.GetOneAlertForProfile(ctx, profileID, alertID)
	if err != nil {
		return model.Alert{}, err
	}
	if err := tx.Commit(); err != nil {
		return model.Alert{}, err
	}

	return b, nil
}

// DoAlertUnread marks an alert as unread for a profile
func (s alertService) DoAlertUnread(ctx context.Context, profileID, alertID uuid.UUID) (model.Alert, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return model.Alert{}, err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)
	if err := qtx.DoAlertUnread(ctx, profileID, alertID); err != nil {
		return model.Alert{}, err
	}
	a, err := qtx.GetOneAlertForProfile(ctx, profileID, alertID)
	if err != nil {
		return model.Alert{}, err
	}
	if err := tx.Commit(); err != nil {
		return model.Alert{}, err
	}

	return a, nil
}
