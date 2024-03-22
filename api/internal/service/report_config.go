package service

import (
	"context"

	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
)

type ReportConfigService interface {
	ListProjectReportConfigs(ctx context.Context, projectID uuid.UUID) ([]model.ReportConfig, error)
	CreateReportConfig(ctx context.Context, rc model.ReportConfig) (model.ReportConfig, error)
	UpdateReportConfig(ctx context.Context, rc model.ReportConfig) error
	DeleteReportConfig(ctx context.Context, reportConfigID uuid.UUID) error
	DownloadReport(ctx context.Context, reportConfigID uuid.UUID) ([]byte, error)
}

type reportConfigService struct {
	db *model.Database
	*model.Queries
}

func NewReportConfigService(db *model.Database, q *model.Queries) *reportConfigService {
	return &reportConfigService{db, q}
}

func (s reportConfigService) CreateReportConfig(ctx context.Context, rc model.ReportConfig) (model.ReportConfig, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return model.ReportConfig{}, err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)

	rcID, err := qtx.CreateReportConfig(ctx, rc)
	if err != nil {
		return model.ReportConfig{}, err
	}

	for _, pc := range rc.PlotConfigs {
		if err := qtx.AssignReportConfigPlotConfig(ctx, rcID, pc.ID); err != nil {
			return model.ReportConfig{}, err
		}
	}

	rcNew, err := qtx.GetReportConfigByID(ctx, rcID)
	if err != nil {
		return model.ReportConfig{}, err
	}

	if err := tx.Commit(); err != nil {
		return model.ReportConfig{}, err
	}
	return rcNew, nil
}

func (s reportConfigService) UpdateReportConfig(ctx context.Context, rc model.ReportConfig) error {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)

	if err := qtx.UpdateReportConfig(ctx, rc); err != nil {
		return err
	}

	if err := qtx.UnassignAllReportConfigPlotConfig(ctx, rc.ID); err != nil {
		return err
	}

	for _, pc := range rc.PlotConfigs {
		if err := qtx.AssignReportConfigPlotConfig(ctx, rc.ID, pc.ID); err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (s reportConfigService) DownloadReport(ctx context.Context, projectID uuid.UUID) ([]byte, error) {
	q := s.db.Queries()
	_, err := q.ListPlotConfigs(ctx, projectID)
	if err != nil {
		return nil, err
	}

	// TODO render pdf plots based on plot configs and time window

	return nil, nil
}
