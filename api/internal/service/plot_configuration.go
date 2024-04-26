package service

import (
	"context"
	"fmt"
	"strconv"
	"strings"

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

	pc.ID = pcID
	if err := qtx.CreatePlotConfigSettings(ctx, pc); err != nil {
		return a, err
	}

	if err := validateCreateTraces(ctx, qtx, pc.ID, pc.Display.Traces); err != nil {
		return pc, err
	}

	if err := validateCreateCustomShapes(ctx, qtx, pc.ID, pc.Display.Layout.CustomShapes); err != nil {
		return pc, err
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
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return pc, err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)

	if err := qtx.UpdatePlotConfig(ctx, pc); err != nil {
		return pc, err
	}

	if err := qtx.DeletePlotConfigSettings(ctx, pc.ID); err != nil {
		return pc, err
	}

	if err := qtx.DeleteAllPlotConfigTimeseriesTraces(ctx, pc.ID); err != nil {
		return pc, err
	}

	if err := qtx.DeleteAllPlotConfigCustomShapes(ctx, pc.ID); err != nil {
		return pc, err
	}

	if err := qtx.CreatePlotConfigSettings(ctx, pc); err != nil {
		return pc, err
	}

	if err := validateCreateTraces(ctx, qtx, pc.ID, pc.Display.Traces); err != nil {
		return pc, err
	}

	if err := validateCreateCustomShapes(ctx, qtx, pc.ID, pc.Display.Layout.CustomShapes); err != nil {
		return pc, err
	}

	pcNew, err := qtx.GetPlotConfig(ctx, pc.ID)
	if err != nil {
		return pc, err
	}

	if err := tx.Commit(); err != nil {
		return pc, err
	}

	return pcNew, nil
}

func validateCreateTraces(ctx context.Context, q *model.Queries, pcID uuid.UUID, trs []model.PlotConfigTimeseriesTrace) error {
	for _, tr := range trs {
		tr.PlotConfigurationID = pcID

		// TODO validate input
		if err := validateColor(tr.Color); err != nil {
			return err
		}

		if err := q.CreatePlotConfigTimeseriesTrace(ctx, tr); err != nil {
			return err
		}
	}
	return nil
}

func validateCreateCustomShapes(ctx context.Context, q *model.Queries, pcID uuid.UUID, css []model.PlotConfigCustomShape) error {
	for _, cs := range css {
		cs.PlotConfigurationID = pcID

		// TODO validate input
		if err := validateColor(cs.Color); err != nil {
			return err
		}

		if err := q.CreatePlotConfigCustomShape(ctx, cs); err != nil {
			return err
		}
	}
	return nil
}

func validateColor(colorHex string) error {
	parts := strings.SplitAfter(colorHex, "#")
	invalidHexErr := fmt.Errorf("invalid hex code format: %s; format must be '#000000'", colorHex)
	if len(parts) != 2 {
		return invalidHexErr
	}
	if len(parts[0]) != 1 && len(parts[1]) != 6 {
		return invalidHexErr
	}
	if _, err := strconv.Atoi(parts[1]); err != nil {
		return invalidHexErr
	}
	return nil
}
