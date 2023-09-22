package store

import (
	"context"
	"log"

	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
)

type AlertConfigStore interface {
	GetAllAlertConfigsForProject(ctx context.Context, projectID *uuid.UUID) ([]model.AlertConfig, error)
	GetAllAlertConfigsForProjectAndAlertType(ctx context.Context, projectID, alertTypeID *uuid.UUID) ([]model.AlertConfig, error)
	GetAllAlertConfigsForInstrument(ctx context.Context, instrumentID *uuid.UUID) ([]model.AlertConfig, error)
	GetOneAlertConfig(ctx context.Context, alertConfigID *uuid.UUID) (*model.AlertConfig, error)
}

type alertConfigStore struct {
	db *model.Database
}

func NewAlertConfigStore(db *model.Database) *alertConfigStore {
	return &alertConfigStore{db}
}

// GetAllAlertConfigsForProject lists all alert configs for a single project
func (s alertConfigStore) GetAllAlertConfigsForProject(ctx context.Context, projectID *uuid.UUID) ([]model.AlertConfig, error) {
	q := model.NewQueries(s.db)
	aa, err := q.GetAllAlertConfigsForProject(ctx, projectID)
	if err != nil {
		return aa, err
	}
	return aa, nil
}

// GetAllAlertConfigsForProjectAndAlertType lists alert configs for a single project filetered by alert type
func (s alertConfigStore) GetAllAlertConfigsForProjectAndAlertType(ctx context.Context, projectID, alertTypeID *uuid.UUID) ([]model.AlertConfig, error) {
	q := model.NewQueries(s.db)
	aa, err := q.GetAllAlertConfigsForProjectAndAlertType(ctx, projectID, alertTypeID)
	if err != nil {
		return aa, err
	}
	return aa, nil
}

// ListInstrumentAlertConfigs lists all alerts for a single instrument
func (s alertConfigStore) GetAllAlertConfigsForInstrument(ctx context.Context, instrumentID *uuid.UUID) ([]model.AlertConfig, error) {
	q := model.NewQueries(s.db)
	aa, err := q.GetAllAlertConfigsForInstrument(ctx, instrumentID)
	if err != nil {
		return aa, err
	}
	return aa, nil
}

// GetOneAlertConfig gets a single alert config
func (s alertConfigStore) GetOneAlertConfig(ctx context.Context, alertConfigID *uuid.UUID) (*model.AlertConfig, error) {
	q := model.NewQueries(s.db)
	aa, err := q.GetOneAlertConfig(ctx, alertConfigID)
	if err != nil {
		return aa, err
	}
	return aa, nil
}

// CreateAlertConfig creates one new alert configuration
func (s alertConfigStore) CreateAlertConfig(ctx context.Context, ac *model.AlertConfig) (*model.AlertConfig, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, err
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

	q := model.NewQueries(s.db).WithTx(tx)

	acID, err := q.CreateAlertConfig(ctx, ac)
	if err != nil {
		return nil, err
	}

	for _, aci := range ac.Instruments {
		if err := q.AssignInstrumentToAlertConfig(ctx, acID, &aci.InstrumentID); err != nil {
			return nil, err
		}
	}

	if err := registerAndSubscribe(ctx, q, acID, ac.AlertEmailSubscriptions); err != nil {
		return nil, err
	}

	if err := q.CreateNextSubmittalFromExistingAlertConfigDate(ctx, acID); err != nil {
		return nil, err
	}

	acNew, err := q.GetOneAlertConfig(ctx, acID)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return acNew, nil
}

// UpdateAlertConfig updates an alert config
func (s alertConfigStore) UpdateAlertConfig(ctx context.Context, alertConfigID *uuid.UUID, ac *model.AlertConfig) (*model.AlertConfig, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, err
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

	q := model.NewQueries(s.db).WithTx(tx)

	if err := q.UpdateAlertConfig(ctx, ac); err != nil {
		return nil, err
	}

	if err := q.UnassignAllInstrumentsFromAlertConfig(ctx, &ac.ID); err != nil {
		return nil, err
	}

	for _, aci := range ac.Instruments {
		if err := q.AssignInstrumentToAlertConfig(ctx, &ac.ID, &aci.InstrumentID); err != nil {
			return nil, err
		}
	}

	if err := q.UnsubscribeAllEmailsFromAlertConfig(ctx, alertConfigID); err != nil {
		return nil, err
	}
	if err := registerAndSubscribe(ctx, q, alertConfigID, ac.AlertEmailSubscriptions); err != nil {
		return nil, err
	}

	if err := q.UpdateFutureSubmittalForAlertConfig(ctx, &ac.ID); err != nil {
		return nil, err
	}

	acNew, err := q.GetOneAlertConfig(ctx, &ac.ID)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return acNew, nil
}

// DeleteAlertConfig deletes an alert by ID
func (s alertConfigStore) DeleteAlertConfig(ctx context.Context, alertConfigID *uuid.UUID) error {
	q := model.NewQueries(s.db)

	if err := q.DeleteAlertConfig(ctx, alertConfigID); err != nil {
		return err
	}

	return nil
}
