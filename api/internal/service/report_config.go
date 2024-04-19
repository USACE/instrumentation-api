package service

import (
	"context"

	"github.com/USACE/instrumentation-api/api/internal/cloud"
	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
)

type ReportConfigService interface {
	ListProjectReportConfigs(ctx context.Context, projectID uuid.UUID) ([]model.ReportConfig, error)
	CreateReportConfig(ctx context.Context, rc model.ReportConfig) (model.ReportConfig, error)
	UpdateReportConfig(ctx context.Context, rc model.ReportConfig) error
	DeleteReportConfig(ctx context.Context, reportConfigID uuid.UUID) error
	GetReportConfigWithPlotConfigs(ctx context.Context, rcID uuid.UUID) (model.ReportConfigWithPlotConfigs, error)
	CreateReportDownloadJob(ctx context.Context, reportConfigID uuid.UUID) (string, error)
}

type reportConfigService struct {
	db *model.Database
	*model.Queries
	pubsub cloud.Pubsub
}

func NewReportConfigService(db *model.Database, q *model.Queries, ps cloud.Pubsub) *reportConfigService {
	return &reportConfigService{db, q, ps}
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

func (s reportConfigService) GetReportConfigWithPlotConfigs(ctx context.Context, rcID uuid.UUID) (model.ReportConfigWithPlotConfigs, error) {
	q := s.db.Queries()

	rc, err := q.GetReportConfigByID(ctx, rcID)
	if err != nil {
		return model.ReportConfigWithPlotConfigs{}, err
	}
	pcs, err := q.ListReportConfigPlotConfigs(ctx, rcID)
	if err != nil {
		return model.ReportConfigWithPlotConfigs{}, err
	}
	return model.ReportConfigWithPlotConfigs{
		ReportConfig: rc,
		PlotConfigs:  pcs,
	}, nil
}

func (s reportConfigService) CreateReportDownloadJob(ctx context.Context, rcID uuid.UUID) (string, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return "", err
	}
	defer model.TxDo(tx.Rollback)

	// qtx := s.WithTx(tx)

	msgID, err := s.pubsub.PublishMessage(ctx, model.ReportConfigJobMessage{ReportConfigID: rcID})
	if err != nil {
		return "", err
	}

	return msgID, nil
}

// func (s reportConfigService) GetReportDownloadJob(ctx context.Context, jobID string) error {
//
// }
