package service

import (
	"context"

	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
)

type PlotConfigService interface {
	ListPlotConfigs(ctx context.Context, projectID uuid.UUID) ([]model.PlotConfig, error)
	GetPlotConfig(ctx context.Context, plotconfigID uuid.UUID) (model.PlotConfig, error)
	CreatePlotConfig(ctx context.Context, pc model.PlotConfig) (model.PlotConfig, error)
	UpdatePlotConfig(ctx context.Context, pc model.PlotConfig) (model.PlotConfig, error)
	DeletePlotConfig(ctx context.Context, projectID, plotConfigID uuid.UUID) error
}

type plotConfigService struct {
	db *model.Database
	*model.Queries
}

func NewPlotConfigService(db *model.Database, q *model.Queries) *plotConfigService {
	return &plotConfigService{db, q}
}

// CreatePlotConfiguration add plot configuration for a project
func (s plotConfigService) CreatePlotConfig(ctx context.Context, pc model.PlotConfig) (model.PlotConfig, error) {
	var a model.PlotConfig
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return a, err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)

	pcID, err := qtx.CreatePlotConfig(ctx, pc)
	if err != nil {
		return a, err
	}

	for _, tsid := range pc.TimeseriesIDs {
		if err := qtx.CreatePlotConfigTimeseries(ctx, pcID, tsid); err != nil {
			return a, err
		}
	}

	pc.ID = pcID
	if err := qtx.CreatePlotConfigSettings(ctx, pc); err != nil {
		return a, err
	}

	pcNew, err := qtx.GetPlotConfig(ctx, pcID)
	if err != nil {
		return a, err
	}

	if err := tx.Commit(); err != nil {
		return a, err
	}

	return pcNew, nil
}

// UpdatePlotConfiguration update plot configuration for a project
func (s plotConfigService) UpdatePlotConfig(ctx context.Context, pc model.PlotConfig) (model.PlotConfig, error) {
	var a model.PlotConfig
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return a, err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)

	if err := qtx.UpdatePlotConfig(ctx, pc); err != nil {
		return a, err
	}

	if err := qtx.DeletePlotConfigTimeseries(ctx, pc.ID); err != nil {
		return a, err
	}

	if err := qtx.DeletePlotConfigSettings(ctx, pc.ID); err != nil {
		return a, err
	}

	if err := qtx.CreatePlotConfigSettings(ctx, pc); err != nil {
		return a, err
	}

	for _, tsID := range pc.TimeseriesIDs {
		if err := qtx.CreatePlotConfigTimeseries(ctx, pc.ID, tsID); err != nil {
			return a, err
		}
	}

	pcNew, err := qtx.GetPlotConfig(ctx, pc.ID)
	if err != nil {
		return a, err
	}

	if err := tx.Commit(); err != nil {
		return a, err
	}

	return pcNew, nil
}
