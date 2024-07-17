package service

import (
	"context"

	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
)

type plotConfigBullseyePlotService interface {
	CreatePlotConfigBullseyePlot(ctx context.Context, pc model.PlotConfigBullseyePlot) (model.PlotConfig, error)
	UpdatePlotConfigBullseyePlot(ctx context.Context, pc model.PlotConfigBullseyePlot) (model.PlotConfig, error)
	ListPlotConfigMeasurementsBullseyePlot(ctx context.Context, plotConfigID uuid.UUID, tw model.TimeWindow) ([]model.PlotConfigMeasurementBullseyePlot, error)
}

func (s plotConfigService) CreatePlotConfigBullseyePlot(ctx context.Context, pc model.PlotConfigBullseyePlot) (model.PlotConfig, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return model.PlotConfig{}, err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)

	pc.PlotType = model.BullseyePlotType
	pcID, err := qtx.CreatePlotConfig(ctx, pc.PlotConfig)
	if err != nil {
		return model.PlotConfig{}, err
	}

	if err := qtx.CreatePlotConfigSettings(ctx, pcID, pc.PlotConfigSettings); err != nil {
		return model.PlotConfig{}, err
	}

	if err := qtx.CreatePlotBullseyeConfig(ctx, pcID, pc.Display); err != nil {
		return model.PlotConfig{}, err
	}

	pcNew, err := qtx.GetPlotConfig(ctx, pcID)
	if err != nil {
		return model.PlotConfig{}, err
	}

	return pcNew, nil
}

func (s plotConfigService) UpdatePlotConfigBullseyePlot(ctx context.Context, pc model.PlotConfigBullseyePlot) (model.PlotConfig, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return model.PlotConfig{}, err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)

	if err := qtx.UpdatePlotConfig(ctx, pc.PlotConfig); err != nil {
		return model.PlotConfig{}, err
	}

	if err := qtx.UpdatePlotBullseyeConfig(ctx, pc.ID, pc.Display); err != nil {
		return model.PlotConfig{}, err
	}

	if err := qtx.DeletePlotConfigSettings(ctx, pc.ID); err != nil {
		return model.PlotConfig{}, err
	}

	if err := qtx.CreatePlotConfigSettings(ctx, pc.ID, pc.PlotConfigSettings); err != nil {
		return model.PlotConfig{}, err
	}

	pcNew, err := qtx.GetPlotConfig(ctx, pc.ID)
	if err != nil {
		return model.PlotConfig{}, err
	}

	return pcNew, nil
}
