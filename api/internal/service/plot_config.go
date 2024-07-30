package service

import (
	"context"

	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
)

type PlotConfigService interface {
	ListPlotConfigs(ctx context.Context, projectID uuid.UUID) ([]model.PlotConfig, error)
	GetPlotConfig(ctx context.Context, plotconfigID uuid.UUID) (model.PlotConfig, error)
	DeletePlotConfig(ctx context.Context, projectID, plotConfigID uuid.UUID) error
	plotConfigBullseyePlotService
	plotConfigContourPlotService
	plotConfigProfilePlotService
	plotConfigScatterLinePlotService
}

type plotConfigService struct {
	db *model.Database
	*model.Queries
}

func NewPlotConfigService(db *model.Database, q *model.Queries) *plotConfigService {
	return &plotConfigService{db, q}
}
