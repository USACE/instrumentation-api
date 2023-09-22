package store

import (
	"context"
	"log"

	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
)

type AlertStore interface {
	CreateAlerts(ctx context.Context, alertConfigIDs []uuid.UUID) error
	GetAllAlertsForProject(ctx context.Context, projectID *uuid.UUID) ([]model.Alert, error)
	GetAllAlertsForInstrument(ctx context.Context, instrumentID *uuid.UUID) ([]model.Alert, error)
	GetAllAlertsForProfile(ctx context.Context, profileID *uuid.UUID) ([]model.Alert, error)
	GetOneAlertForProfile(ctx context.Context, profileID *uuid.UUID, alertID *uuid.UUID) (*model.Alert, error)
	DoAlertRead(ctx context.Context, profileID *uuid.UUID, alertID *uuid.UUID) (*model.Alert, error)
	DoAlertUnread(ctx context.Context, profileID *uuid.UUID, alertID *uuid.UUID) (*model.Alert, error)
}

type alertStore struct {
	db *model.Database
}

func NewAlertStore(db *model.Database) *alertStore {
	return &alertStore{db}
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

	q := model.NewQueries(s.db).WithTx(tx)
	for _, id := range alertConfigIDs {
		if err := q.CreateAlerts(ctx, id); err != nil {
			return err
		}
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

// GetAllByProject lists all alerts for a given instrument ID
func (s alertStore) GetAllAlertsForProject(ctx context.Context, projectID *uuid.UUID) ([]model.Alert, error) {
	q := model.NewQueries(s.db)
	aa, err := q.GetAllAlertsForProject(ctx, projectID)
	if err != nil {
		return aa, err
	}
	return aa, nil
}

// GetAllByInstrument lists all alerts for a given instrument ID
func (s alertStore) GetAllAlertsForInstrument(ctx context.Context, instrumentID *uuid.UUID) ([]model.Alert, error) {
	q := model.NewQueries(s.db)
	aa, err := q.GetAllAlertsForInstrument(ctx, instrumentID)
	if err != nil {
		return aa, err
	}
	return aa, nil
}

// GetAllByProfile returns all alerts for which a profile is subscribed to the AlertConfig
func (s alertStore) GetAllAlertsForProfile(ctx context.Context, profileID *uuid.UUID) ([]model.Alert, error) {
	q := model.NewQueries(s.db)
	aa, err := q.GetAllAlertsForProfile(ctx, profileID)
	if err != nil {
		return aa, err
	}
	return aa, nil
}

// GetOneByProfile returns a single alert for which a profile is subscribed
func (s alertStore) GetOneAlertForProfile(ctx context.Context, profileID *uuid.UUID, alertID *uuid.UUID) (*model.Alert, error) {
	q := model.NewQueries(s.db)
	aa, err := q.GetOneAlertForProfile(ctx, profileID, alertID)
	if err != nil {
		return aa, err
	}
	return aa, nil
}

// DoAlertRead marks an alert as read for a profile
func (s alertStore) DoAlertRead(ctx context.Context, profileID *uuid.UUID, alertID *uuid.UUID) (*model.Alert, error) {
	q := model.NewQueries(s.db)
	aa, err := q.DoAlertRead(ctx, profileID, alertID)
	if err != nil {
		return aa, err
	}
	return aa, nil
}

// DoAlertUnread marks an alert as unread for a profile
func (s alertStore) DoAlertUnread(ctx context.Context, profileID *uuid.UUID, alertID *uuid.UUID) (*model.Alert, error) {
	q := model.NewQueries(s.db)
	aa, err := q.DoAlertUnread(ctx, profileID, alertID)
	if err != nil {
		return aa, err
	}
	return aa, nil
}
