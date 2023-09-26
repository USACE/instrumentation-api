package store

import (
	"context"
	"log"

	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
)

type AlertStore interface {
	CreateAlerts(ctx context.Context, alertConfigIDs []uuid.UUID) error
	GetAllAlertsForProject(ctx context.Context, projectID uuid.UUID) ([]model.Alert, error)
	GetAllAlertsForInstrument(ctx context.Context, instrumentID uuid.UUID) ([]model.Alert, error)
	GetAllAlertsForProfile(ctx context.Context, profileID uuid.UUID) ([]model.Alert, error)
	GetOneAlertForProfile(ctx context.Context, profileID uuid.UUID, alertID uuid.UUID) (model.Alert, error)
	DoAlertRead(ctx context.Context, profileID uuid.UUID, alertID uuid.UUID) (model.Alert, error)
	DoAlertUnread(ctx context.Context, profileID uuid.UUID, alertID uuid.UUID) (model.Alert, error)
}

type alertStore struct {
	db *model.Database
	q  *model.Queries
}

func NewAlertStore(db *model.Database, q *model.Queries) *alertStore {
	return &alertStore{db, q}
}

// Create creates one or more new alerts
func (s alertStore) CreateAlerts(ctx context.Context, alertConfigIDs []uuid.UUID) error {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			log.Print("error rolling back changes")
		}
	}()

	qtx := s.q.WithTx(tx)
	for _, id := range alertConfigIDs {
		if err := qtx.CreateAlerts(ctx, id); err != nil {
			return err
		}
	}
	return tx.Commit()
}

// GetAllByProject lists all alerts for a given instrument ID
func (s alertStore) GetAllAlertsForProject(ctx context.Context, projectID uuid.UUID) ([]model.Alert, error) {
	return s.q.GetAllAlertsForProject(ctx, projectID)
}

// GetAllByInstrument lists all alerts for a given instrument ID
func (s alertStore) GetAllAlertsForInstrument(ctx context.Context, instrumentID uuid.UUID) ([]model.Alert, error) {
	return s.q.GetAllAlertsForInstrument(ctx, instrumentID)
}

// GetAllByProfile returns all alerts for which a profile is subscribed to the AlertConfig
func (s alertStore) GetAllAlertsForProfile(ctx context.Context, profileID uuid.UUID) ([]model.Alert, error) {
	return s.q.GetAllAlertsForProfile(ctx, profileID)
}

// GetOneByProfile returns a single alert for which a profile is subscribed
func (s alertStore) GetOneAlertForProfile(ctx context.Context, profileID, alertID uuid.UUID) (model.Alert, error) {
	return s.q.GetOneAlertForProfile(ctx, profileID, alertID)
}

// DoAlertRead marks an alert as read for a profile
func (s alertStore) DoAlertRead(ctx context.Context, profileID, alertID uuid.UUID) (model.Alert, error) {
	var a model.Alert
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return a, err
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			log.Print("error rolling back changes")
		}
	}()

	qtx := s.q.WithTx(tx)
	if err := qtx.DoAlertRead(ctx, profileID, alertID); err != nil {
		return a, err
	}
	b, err := qtx.GetOneAlertForProfile(ctx, profileID, alertID)
	if err != nil {
		return a, err
	}
	if err := tx.Commit(); err != nil {
		return a, err
	}

	return b, nil
}

// DoAlertUnread marks an alert as unread for a profile
func (s alertStore) DoAlertUnread(ctx context.Context, profileID, alertID uuid.UUID) (model.Alert, error) {
	var b model.Alert
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return b, err
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			log.Print("error rolling back changes")
		}
	}()

	qtx := s.q.WithTx(tx)
	if err := qtx.DoAlertUnread(ctx, profileID, alertID); err != nil {
		return b, err
	}
	b, err = qtx.GetOneAlertForProfile(ctx, profileID, alertID)
	if err != nil {
		return b, err
	}
	if err := tx.Commit(); err != nil {
		return b, err
	}

	return b, nil
}
