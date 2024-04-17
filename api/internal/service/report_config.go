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
	GetReportConfigWithPlotConfigs(ctx context.Context, rcID uuid.UUID) (model.ReportConfigWithPlotConfigs, error)
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

// func (s reportConfigService) DownloadReport(ctx context.Context, rcID uuid.UUID) ([]byte, error) {
// 	q := s.db.Queries()
//
// 	rc, err := q.GetReportConfigByID(ctx, rcID)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	pcs, err := q.ListReportConfigPlotConfigs(ctx, rc.ID)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	// TODO: impose a limit on the number of plots allowed in a report?
// 	// Otherwise this could go off the rails
// 	for _, pc := range pcs {
// 		var tw model.TimeWindow
// 		pctw, err := pc.DateRangeTimeWindow()
// 		if err != nil {
// 			continue
// 		}
// 		if rc.After != nil {
// 			tw.After = *rc.After
// 		} else {
// 			tw.After = pctw.After
// 		}
// 		if rc.Before != nil {
// 			tw.Before = *rc.Before
// 		} else {
// 			tw.Before = pctw.Before
// 		}
//
// 		// TODO: is this the same API that queries masked/validated/etc?
// 		// Stored timeseries should adhere to these filters
// 		// It is not applicable to computed timeseries
// 		// pgx has automatic prepared statement caching, this is ok
// 		mm, err := q.SelectMeasurements(ctx, model.ProcessMeasurementFilter{
// 			TimeseriesIDs: pc.TimeseriesIDs,
// 			After:         tw.After,
// 			Before:        tw.Before,
// 		})
// 		if err != nil {
// 			return nil, err
// 		}
//
// 		log.Print(mm)
// 		// TODO: render plot to pdf to allow measurements to be gc'd
// 	}
//
// 	// TODO: return pdf
//
// 	return nil, nil
// }

func (s reportConfigService) DownloadReport(ctx context.Context, rcID uuid.UUID) ([]byte, error) {
	return nil, nil
}
