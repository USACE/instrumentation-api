package service

import (
	"context"

	"github.com/USACE/instrumentation-api/api/internal/model"
)

type alertConfigChangeService interface {
	CreateAlertConfigChange(ctx context.Context, ac model.AlertConfigChange) (model.AlertConfig, error)
	UpdateAlertConfigChange(ctx context.Context, ac model.AlertConfigChange) (model.AlertConfig, error)
}

func (s alertConfigService) CreateAlertConfigChange(ctx context.Context, ac model.AlertConfigChange) (model.AlertConfig, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return model.AlertConfig{}, err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)

	acID, err := qtx.CreateAlertConfig(ctx, ac.AlertConfig)
	if err != nil {
		return model.AlertConfig{}, err
	}
	ac.ID = acID

	if err := qtx.CreateAlertConfigChange(ctx, ac.ID, ac.Opts); err != nil {
		return model.AlertConfig{}, err
	}

	for _, aci := range ac.Timeseries {
		if err := qtx.AssignTimeseriesToAlertConfig(ctx, ac.ID, aci.TimeseriesID); err != nil {
			return model.AlertConfig{}, err
		}
	}

	if err := registerAndSubscribe(ctx, qtx, ac.ID, ac.AlertEmailSubscriptions); err != nil {
		return model.AlertConfig{}, err
	}

	acNew, err := qtx.GetOneAlertConfig(ctx, ac.ID)
	if err != nil {
		return model.AlertConfig{}, err
	}

	err = tx.Commit()

	return acNew, err
}

func (s alertConfigService) UpdateAlertConfigChange(ctx context.Context, ac model.AlertConfigChange) (model.AlertConfig, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return model.AlertConfig{}, err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)

	if err := qtx.UpdateAlertConfig(ctx, ac.AlertConfig); err != nil {
		return model.AlertConfig{}, err
	}

	if err := qtx.UpdateAlertConfigChange(ctx, ac.ID, ac.Opts); err != nil {
		return model.AlertConfig{}, err
	}

	if err := qtx.UnassignAllTimeseriesFromAlertConfig(ctx, ac.ID); err != nil {
		return model.AlertConfig{}, err
	}

	for _, aci := range ac.Timeseries {
		if err := qtx.AssignInstrumentToAlertConfig(ctx, ac.ID, aci.TimeseriesID); err != nil {
			return model.AlertConfig{}, err
		}
	}

	if err := qtx.UnsubscribeAllEmailsFromAlertConfig(ctx, ac.ID); err != nil {
		return model.AlertConfig{}, err
	}
	if err := registerAndSubscribe(ctx, qtx, ac.ID, ac.AlertEmailSubscriptions); err != nil {
		return model.AlertConfig{}, err
	}

	acNew, err := qtx.GetOneAlertConfig(ctx, ac.ID)
	if err != nil {
		return model.AlertConfig{}, err
	}

	err = tx.Commit()

	return acNew, err
}
