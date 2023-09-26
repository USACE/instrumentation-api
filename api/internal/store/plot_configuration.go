package store

import (
	"context"
	"log"

	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
)

type PlotConfigStore interface {
}

type plotConfigStore struct {
	db *model.Database
	q  *model.Queries
}

func NewPlotConfigStore(db *model.Database, q *model.Queries) *plotConfigStore {
	return &plotConfigStore{db, q}
}

// ListPlotConfigSlugs lists used instrument group slugs in the database
func (s plotConfigStore) ListPlotConfigSlugs(ctx context.Context) ([]string, error) {
	return s.q.ListPlotConfigSlugs(ctx)
}

// ListPlotConfigs returns a list of Plot groups
func (s plotConfigStore) ListPlotConfigs(ctx context.Context, projectID uuid.UUID) ([]model.PlotConfig, error) {
	return s.q.ListPlotConfigs(ctx, projectID)
}

// GetPlotConfig returns a single plot configuration
func (s plotConfigStore) GetPlotConfig(ctx context.Context, plotconfigID uuid.UUID) (model.PlotConfig, error) {
	return s.q.GetPlotConfig(ctx, plotconfigID)
}

// CreatePlotConfiguration add plot configuration for a project
func (s plotConfigStore) CreatePlotConfig(ctx context.Context, pc model.PlotConfig) (model.PlotConfig, error) {
	var a model.PlotConfig
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return a, err
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			log.Print(err.Error())
		}
	}()

	qtx := s.q.WithTx(tx)

	pcID, err := qtx.CreatePlotConfig(ctx, pc)
	if err != nil {
		return a, err
	}
	for _, tsid := range pc.TimeseriesID {
		if err := qtx.CreatePlotConfigTimeseries(ctx, pcID, tsid); err != nil {
			return a, err
		}
	}
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

const updatePlotConfig = `
	UPDATE plot_configuration SET name = $3, updater = $4, update_date = $5 WHERE project_id = $1 AND id = $2
`

func (s plotConfigStore) UpdatePlotConfig(ctx context.Context, pc model.PlotConfig) error {
	return s.q.UpdatePlotConfig(ctx, pc)
}

// UpdatePlotConfiguration update plot configuration for a project
func (s plotConfigStore) UpdatePlotConfiguration(ctx context.Context, pc model.PlotConfig) (model.PlotConfig, error) {
	var a model.PlotConfig
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return a, err
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			log.Print(err.Error())
		}
	}()

	qtx := s.q.WithTx(tx)

	if err != nil {
		return a, err
	}

	if err := qtx.UpdatePlotConfig(ctx, pc); err != nil {
		return a, err
	}

	if err := qtx.DeletePlotConfigTimeseries(ctx, pc.ID, pc.TimeseriesID); err != nil {
		return a, err
	}

	if err := qtx.DeletePlotConfigSettings(ctx, pc.ID); err != nil {
		return a, err
	}

	if err := qtx.CreatePlotConfigSettings(ctx, pc); err != nil {
		return a, err
	}

	for _, tsID := range pc.TimeseriesID {
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

// DeletePlotConfig delete plot configuration for a project
func (s plotConfigStore) DeletePlotConfig(ctx context.Context, projectID, plotConfigID uuid.UUID) error {
	return s.q.DeletePlotConfig(ctx, projectID, plotConfigID)
}
