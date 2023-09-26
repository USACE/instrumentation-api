package store

import (
	"context"
	"log"

	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
)

type AlertConfigStore interface {
	GetAllAlertConfigsForProject(ctx context.Context, projectID *uuid.UUID) ([]model.AlertConfig, error)
	GetAllAlertConfigsForProjectAndAlertType(ctx context.Context, projectID, alertTypeID uuid.UUID) ([]model.AlertConfig, error)
	GetAllAlertConfigsForInstrument(ctx context.Context, instrumentID uuid.UUID) ([]model.AlertConfig, error)
	GetOneAlertConfig(ctx context.Context, alertConfigID uuid.UUID) (model.AlertConfig, error)
}

type alertConfigStore struct {
	db *model.Database
	q  *model.Queries
}

func NewAlertConfigStore(db *model.Database, q *model.Queries) *alertConfigStore {
	return &alertConfigStore{db, q}
}

// GetAllAlertConfigsForProject lists all alert configs for a single project
func (s alertConfigStore) GetAllAlertConfigsForProject(ctx context.Context, projectID uuid.UUID) ([]model.AlertConfig, error) {
	return s.q.GetAllAlertConfigsForProject(ctx, projectID)
}

// GetAllAlertConfigsForProjectAndAlertType lists alert configs for a single project filetered by alert type
func (s alertConfigStore) GetAllAlertConfigsForProjectAndAlertType(ctx context.Context, projectID, alertTypeID uuid.UUID) ([]model.AlertConfig, error) {
	return s.q.GetAllAlertConfigsForProjectAndAlertType(ctx, projectID, alertTypeID)
}

// ListInstrumentAlertConfigs lists all alerts for a single instrument
func (s alertConfigStore) GetAllAlertConfigsForInstrument(ctx context.Context, instrumentID uuid.UUID) ([]model.AlertConfig, error) {
	return s.q.GetAllAlertConfigsForInstrument(ctx, instrumentID)
}

// GetOneAlertConfig gets a single alert config
func (s alertConfigStore) GetOneAlertConfig(ctx context.Context, alertConfigID uuid.UUID) (model.AlertConfig, error) {
	return s.q.GetOneAlertConfig(ctx, alertConfigID)
}

// CreateAlertConfig creates one new alert configuration
func (s alertConfigStore) CreateAlertConfig(ctx context.Context, ac model.AlertConfig) (model.AlertConfig, error) {
	var a model.AlertConfig
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return a, err
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			log.Print(err.Error())
		}
	}()

	if ac.RemindInterval == "" {
		ac.RemindInterval = "PT0"
	}
	if ac.WarningInterval == "" {
		ac.WarningInterval = "PT0"
	}

	qtx := s.q.WithTx(tx)

	acID, err := qtx.CreateAlertConfig(ctx, ac)
	if err != nil {
		return a, err
	}

	for _, aci := range ac.Instruments {
		if err := qtx.AssignInstrumentToAlertConfig(ctx, acID, aci.InstrumentID); err != nil {
			return a, err
		}
	}

	if err := registerAndSubscribe(ctx, qtx, acID, ac.AlertEmailSubscriptions); err != nil {
		return a, err
	}

	if err := qtx.CreateNextSubmittalFromExistingAlertConfigDate(ctx, acID); err != nil {
		return a, err
	}

	acNew, err := qtx.GetOneAlertConfig(ctx, acID)
	if err != nil {
		return a, err
	}

	if err := tx.Commit(); err != nil {
		return a, err
	}

	return acNew, nil
}

// UpdateAlertConfig updates an alert config
func (s alertConfigStore) UpdateAlertConfig(ctx context.Context, alertConfigID uuid.UUID, ac model.AlertConfig) (model.AlertConfig, error) {
	var a model.AlertConfig
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return a, err
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			log.Print(err.Error())
		}
	}()

	if ac.RemindInterval == "" {
		ac.RemindInterval = "PT0"
	}
	if ac.WarningInterval == "" {
		ac.WarningInterval = "PT0"
	}

	qtx := s.q.WithTx(tx)

	if err := qtx.UpdateAlertConfig(ctx, ac); err != nil {
		return a, err
	}

	if err := qtx.UnassignAllInstrumentsFromAlertConfig(ctx, ac.ID); err != nil {
		return a, err
	}

	for _, aci := range ac.Instruments {
		if err := qtx.AssignInstrumentToAlertConfig(ctx, ac.ID, aci.InstrumentID); err != nil {
			return a, err
		}
	}

	if err := qtx.UnsubscribeAllEmailsFromAlertConfig(ctx, alertConfigID); err != nil {
		return a, err
	}
	if err := registerAndSubscribe(ctx, qtx, alertConfigID, ac.AlertEmailSubscriptions); err != nil {
		return a, err
	}

	if err := qtx.UpdateFutureSubmittalForAlertConfig(ctx, ac.ID); err != nil {
		return a, err
	}

	acNew, err := qtx.GetOneAlertConfig(ctx, ac.ID)
	if err != nil {
		return a, err
	}

	if err := tx.Commit(); err != nil {
		return a, err
	}

	return acNew, nil
}

// DeleteAlertConfig deletes an alert by ID
func (s alertConfigStore) DeleteAlertConfig(ctx context.Context, alertConfigID uuid.UUID) error {
	err := s.q.DeleteAlertConfig(ctx, alertConfigID)
	return err
}
