package service

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/model"
)

type alertConfigSchedulerService interface {
	CreateAlertConfigScheduler(ctx context.Context, ac model.AlertConfigScheduler) (model.AlertConfig, error)
	UpdateAlertConfigScheduler(ctx context.Context, ac model.AlertConfigScheduler) (model.AlertConfig, error)
}

func (s alertConfigService) CreateAlertConfigScheduler(ctx context.Context, ac model.AlertConfigScheduler) (model.AlertConfig, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return model.AlertConfig{}, err
	}
	defer model.TxDo(tx.Rollback)

	if ac.Opts.RemindInterval == "" {
		ac.Opts.RemindInterval = "PT0"
	}
	if ac.Opts.WarningInterval == "" {
		ac.Opts.WarningInterval = "PT0"
	}

	qtx := s.WithTx(tx)

	acID, err := qtx.CreateAlertConfig(ctx, ac.AlertConfig)
	if err != nil {
		return model.AlertConfig{}, err
	}
	ac.ID = acID

	if err := qtx.CreateAlertConfigScheduler(ctx, ac.ID, ac.Opts); err != nil {
		return model.AlertConfig{}, err
	}

	for _, aci := range ac.Instruments {
		if err := qtx.AssignInstrumentToAlertConfig(ctx, ac.ID, aci.InstrumentID); err != nil {
			return model.AlertConfig{}, err
		}
	}

	if err := registerAndSubscribe(ctx, qtx, ac.ID, ac.AlertEmailSubscriptions); err != nil {
		return model.AlertConfig{}, err
	}

	if err := qtx.CreateNextSubmittalFromNewAlertConfigDate(ctx, ac.ID); err != nil {
		return model.AlertConfig{}, err
	}

	acNew, err := qtx.GetOneAlertConfig(ctx, ac.ID)
	if err != nil {
		return model.AlertConfig{}, err
	}

	err = tx.Commit()

	return acNew, err
}

func (s alertConfigService) UpdateAlertConfigScheduler(ctx context.Context, ac model.AlertConfigScheduler) (model.AlertConfig, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return model.AlertConfig{}, err
	}
	defer model.TxDo(tx.Rollback)

	if ac.Opts.RemindInterval == "" {
		ac.Opts.RemindInterval = "PT0"
	}
	if ac.Opts.WarningInterval == "" {
		ac.Opts.WarningInterval = "PT0"
	}

	qtx := s.WithTx(tx)

	if err := qtx.UpdateAlertConfig(ctx, ac.AlertConfig); err != nil {
		return model.AlertConfig{}, err
	}

	if err := qtx.UpdateAlertConfigScheduler(ctx, ac.ID, ac.Opts); err != nil {
		return model.AlertConfig{}, err
	}

	if err := qtx.UnassignAllInstrumentsFromAlertConfig(ctx, ac.ID); err != nil {
		return model.AlertConfig{}, err
	}

	for _, aci := range ac.Instruments {
		if err := qtx.AssignInstrumentToAlertConfig(ctx, ac.ID, aci.InstrumentID); err != nil {
			return model.AlertConfig{}, err
		}
	}

	if err := qtx.UnsubscribeAllEmailsFromAlertConfig(ctx, ac.ID); err != nil {
		return model.AlertConfig{}, err
	}
	if err := registerAndSubscribe(ctx, qtx, ac.ID, ac.AlertEmailSubscriptions); err != nil {
		return model.AlertConfig{}, err
	}

	if err := qtx.UpdateFutureSubmittalForAlertConfig(ctx, ac.ID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			t := time.Now()
			if err := qtx.CreateNextSubmittalFromExistingAlertConfigDate(ctx, ac.ID, &t); err != nil {
				return model.AlertConfig{}, err
			}
		} else {
			return model.AlertConfig{}, err
		}
	}

	acNew, err := qtx.GetOneAlertConfig(ctx, ac.ID)
	if err != nil {
		return model.AlertConfig{}, err
	}

	err = tx.Commit()

	return acNew, err
}
