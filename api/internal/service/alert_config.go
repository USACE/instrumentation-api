package service

import (
	"context"

	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
)

type AlertConfigService interface {
	GetAllAlertConfigsForProject(ctx context.Context, projectID uuid.UUID) ([]model.AlertConfig, error)
	GetAllAlertConfigsForProjectAndAlertType(ctx context.Context, projectID, alertTypeID uuid.UUID) ([]model.AlertConfig, error)
	GetAllAlertConfigsForInstrument(ctx context.Context, instrumentID uuid.UUID) ([]model.AlertConfig, error)
	GetOneAlertConfig(ctx context.Context, alertConfigID uuid.UUID) (model.AlertConfig, error)
	CreateAlertConfig(ctx context.Context, ac model.AlertConfig) (model.AlertConfig, error)
	UpdateAlertConfig(ctx context.Context, alertConfigID uuid.UUID, ac model.AlertConfig) (model.AlertConfig, error)
	DeleteAlertConfig(ctx context.Context, alertConfigID uuid.UUID) error
}

type alertConfigService struct {
	db *model.Database
	*model.Queries
}

func NewAlertConfigService(db *model.Database, q *model.Queries) *alertConfigService {
	return &alertConfigService{db, q}
}

// CreateAlertConfig creates one new alert configuration
func (s alertConfigService) CreateAlertConfig(ctx context.Context, ac model.AlertConfig) (model.AlertConfig, error) {
	var a model.AlertConfig
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return a, err
	}
	defer model.TxDo(tx.Rollback)

	if ac.RemindInterval == "" {
		ac.RemindInterval = "PT0"
	}
	if ac.WarningInterval == "" {
		ac.WarningInterval = "PT0"
	}

	qtx := s.WithTx(tx)

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
func (s alertConfigService) UpdateAlertConfig(ctx context.Context, alertConfigID uuid.UUID, ac model.AlertConfig) (model.AlertConfig, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return model.AlertConfig{}, err
	}
	defer model.TxDo(tx.Rollback)

	if ac.RemindInterval == "" {
		ac.RemindInterval = "PT0"
	}
	if ac.WarningInterval == "" {
		ac.WarningInterval = "PT0"
	}

	qtx := s.WithTx(tx)

	if err := qtx.UpdateAlertConfig(ctx, ac); err != nil {
		return model.AlertConfig{}, err
	}

	if err := qtx.UnassignAllInstrumentsFromAlertConfig(ctx, alertConfigID); err != nil {
		return model.AlertConfig{}, err
	}

	for _, aci := range ac.Instruments {
		if err := qtx.AssignInstrumentToAlertConfig(ctx, alertConfigID, aci.InstrumentID); err != nil {
			return model.AlertConfig{}, err
		}
	}

	if err := qtx.UnsubscribeAllEmailsFromAlertConfig(ctx, alertConfigID); err != nil {
		return model.AlertConfig{}, err
	}
	if err := registerAndSubscribe(ctx, qtx, alertConfigID, ac.AlertEmailSubscriptions); err != nil {
		return model.AlertConfig{}, err
	}

	if err := qtx.UpdateFutureSubmittalForAlertConfig(ctx, alertConfigID); err != nil {
		return model.AlertConfig{}, err
	}

	acNew, err := qtx.GetOneAlertConfig(ctx, alertConfigID)
	if err != nil {
		return model.AlertConfig{}, err
	}

	if err := tx.Commit(); err != nil {
		return model.AlertConfig{}, err
	}

	return acNew, nil
}
