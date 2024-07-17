package service

import (
	"context"

	"github.com/USACE/instrumentation-api/api/internal/model"
)

type plotConfigProfilePlotService interface {
	CreatePlotConfigProfilePlot(ctx context.Context, pc model.PlotConfigProfilePlot) (model.PlotConfig, error)
	UpdatePlotConfigProfilePlot(ctx context.Context, pc model.PlotConfigProfilePlot) (model.PlotConfig, error)
}

func (s plotConfigService) CreatePlotConfigProfilePlot(ctx context.Context, pc model.PlotConfigProfilePlot) (model.PlotConfig, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return model.PlotConfig{}, err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)

	pc.PlotType = model.ProfilePlotType
	pcID, err := qtx.CreatePlotConfig(ctx, pc.PlotConfig)
	if err != nil {
		return model.PlotConfig{}, err
	}

	if err := qtx.CreatePlotConfigSettings(ctx, pcID, pc.PlotConfigSettings); err != nil {
		return model.PlotConfig{}, err
	}

	if err := qtx.CreatePlotProfileConfig(ctx, pcID, pc.Display); err != nil {
		return model.PlotConfig{}, err
	}

	pcNew, err := qtx.GetPlotConfig(ctx, pcID)
	if err != nil {
		return model.PlotConfig{}, err
	}

	return pcNew, nil
}

func (s plotConfigService) UpdatePlotConfigProfilePlot(ctx context.Context, pc model.PlotConfigProfilePlot) (model.PlotConfig, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return model.PlotConfig{}, err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)

	if err := qtx.UpdatePlotConfig(ctx, pc.PlotConfig); err != nil {
		return model.PlotConfig{}, err
	}

	if err := qtx.UpdatePlotProfileConfig(ctx, pc.ID, pc.Display); err != nil {
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
