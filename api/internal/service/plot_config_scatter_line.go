package service

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
)

type plotConfigScatterLinePlotService interface {
	CreatePlotConfigScatterLinePlot(ctx context.Context, pc model.PlotConfigScatterLinePlot) (model.PlotConfig, error)
	UpdatePlotConfigScatterLinePlot(ctx context.Context, pc model.PlotConfigScatterLinePlot) (model.PlotConfig, error)
}

func (s plotConfigService) CreatePlotConfigScatterLinePlot(ctx context.Context, pc model.PlotConfigScatterLinePlot) (model.PlotConfig, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return model.PlotConfig{}, err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)

	pc.PlotType = model.ScatterLinePlotType
	pcID, err := qtx.CreatePlotConfig(ctx, pc.PlotConfig)
	if err != nil {
		return model.PlotConfig{}, err
	}

	if err := qtx.CreatePlotConfigSettings(ctx, pcID, pc.PlotConfigSettings); err != nil {
		return model.PlotConfig{}, err
	}

	if err := validateCreateTraces(ctx, qtx, pcID, pc.Display.Traces); err != nil {
		return model.PlotConfig{}, err
	}

	if err := qtx.CreatePlotConfigScatterLineLayout(ctx, pcID, pc.Display.Layout); err != nil {
		return model.PlotConfig{}, err
	}

	if err := validateCreateCustomShapes(ctx, qtx, pcID, pc.Display.Layout.CustomShapes); err != nil {
		return model.PlotConfig{}, err
	}
	pcNew, err := qtx.GetPlotConfig(ctx, pcID)
	if err != nil {
		return model.PlotConfig{}, err
	}

	err = tx.Commit()

	return pcNew, err
}

func (s plotConfigService) UpdatePlotConfigScatterLinePlot(ctx context.Context, pc model.PlotConfigScatterLinePlot) (model.PlotConfig, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return model.PlotConfig{}, err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)

	if err := qtx.UpdatePlotConfig(ctx, pc.PlotConfig); err != nil {
		return model.PlotConfig{}, err
	}

	if err := qtx.DeletePlotConfigSettings(ctx, pc.ID); err != nil {
		log.Printf("fails on delete %s", pc.ID)
		return model.PlotConfig{}, err
	}

	if err := qtx.DeleteAllPlotConfigTimeseriesTraces(ctx, pc.ID); err != nil {
		return model.PlotConfig{}, err
	}

	if err := qtx.DeleteAllPlotConfigCustomShapes(ctx, pc.ID); err != nil {
		return model.PlotConfig{}, err
	}

	if err := qtx.CreatePlotConfigSettings(ctx, pc.ID, pc.PlotConfigSettings); err != nil {
		log.Printf("fails on create %s, %+v", pc.ID, pc.PlotConfigSettings)
		return model.PlotConfig{}, err
	}

	if err := validateCreateTraces(ctx, qtx, pc.ID, pc.Display.Traces); err != nil {
		return model.PlotConfig{}, err
	}

	if err := qtx.UpdatePlotConfigScatterLineLayout(ctx, pc.ID, pc.Display.Layout); err != nil {
		return model.PlotConfig{}, err
	}

	if err := validateCreateCustomShapes(ctx, qtx, pc.ID, pc.Display.Layout.CustomShapes); err != nil {
		return model.PlotConfig{}, err
	}

	pcNew, err := qtx.GetPlotConfig(ctx, pc.ID)
	if err != nil {
		return model.PlotConfig{}, err
	}

	err = tx.Commit()

	return pcNew, err
}

func validateCreateTraces(ctx context.Context, q *model.Queries, pcID uuid.UUID, trs []model.PlotConfigScatterLineTimeseriesTrace) error {
	for _, tr := range trs {
		tr.PlotConfigurationID = pcID

		if err := validateColor(tr.Color); err != nil {
			return err
		}
		if tr.LineStyle == "" {
			tr.LineStyle = "solid"
		}
		if tr.YAxis == "" {
			tr.YAxis = "y1"
		}

		if err := q.CreatePlotConfigTimeseriesTrace(ctx, tr); err != nil {
			return err
		}
	}
	return nil
}

func validateCreateCustomShapes(ctx context.Context, q *model.Queries, pcID uuid.UUID, css []model.PlotConfigScatterLineCustomShape) error {
	for _, cs := range css {
		cs.PlotConfigurationID = pcID

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
	for _, r := range parts[1] {
		if !(r >= '0' && r <= '9' || r >= 'a' && r <= 'f' || r >= 'A' && r <= 'F') {
			return invalidHexErr
		}
	}
	return nil
}
