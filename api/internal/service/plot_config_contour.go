package service

import (
	"context"

	"github.com/USACE/instrumentation-api/api/internal/model"
)

type plotConfigContourPlotService interface {
	CreatePlotConfigContourPlot(ctx context.Context, pc model.PlotConfigContourPlot) (model.PlotConfig, error)
	UpdatePlotConfigContourPlot(ctx context.Context, pc model.PlotConfigContourPlot) (model.PlotConfig, error)
}

func (s plotConfigService) CreatePlotConfigContourPlot(ctx context.Context, pc model.PlotConfigContourPlot) (model.PlotConfig, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return model.PlotConfig{}, err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)

	pcID, err := qtx.CreatePlotConfig(ctx, pc.PlotConfig)
	if err != nil {
		return model.PlotConfig{}, err
	}

	if err := qtx.CreatePlotContourConfig(ctx, pcID, pc.Display); err != nil {
		return model.PlotConfig{}, err
	}

	for _, tsID := range pc.Display.TimeseiresIDs {
		if err := qtx.CreatePlotContourConfigTimeseries(ctx, pc.ID, tsID); err != nil {
			return model.PlotConfig{}, err
		}
	}

	pcNew, err := qtx.GetPlotConfig(ctx, pc.ID)
	if err != nil {
		return model.PlotConfig{}, err
	}

	return pcNew, nil
}

func (s plotConfigService) UpdatePlotConfigContourPlot(ctx context.Context, pc model.PlotConfigContourPlot) (model.PlotConfig, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return model.PlotConfig{}, err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)

	if err := qtx.UpdatePlotConfig(ctx, pc.PlotConfig); err != nil {
		return model.PlotConfig{}, err
	}

	if err := qtx.UpdatePlotContourConfig(ctx, pc.ID, pc.Display); err != nil {
		return model.PlotConfig{}, nil
	}

	if err := qtx.DeletePlotConfigSettings(ctx, pc.ID); err != nil {
		return model.PlotConfig{}, err
	}

	if err := qtx.DeleteAllPlotContourConfigTimeseries(ctx, pc.ID); err != nil {
		return model.PlotConfig{}, nil
	}

	if err := qtx.CreatePlotConfigSettings(ctx, pc.ID, pc.PlotConfigSettings); err != nil {
		return model.PlotConfig{}, err
	}

	for _, tsID := range pc.Display.TimeseiresIDs {
		if err := qtx.CreatePlotContourConfigTimeseries(ctx, pc.ID, tsID); err != nil {
			return model.PlotConfig{}, err
		}
	}

	pcNew, err := qtx.GetPlotConfig(ctx, pc.ID)
	if err != nil {
		return model.PlotConfig{}, err
	}

	return pcNew, nil
}
