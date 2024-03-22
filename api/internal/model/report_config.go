package model

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type ReportConfig struct {
	ID          uuid.UUID               `json:"id" db:"id"`
	Slug        string                  `json:"slug" db:"slug"`
	Name        string                  `json:"name" db:"name"`
	Description string                  `json:"description" db:"description"`
	ProjectID   uuid.UUID               `json:"project_id" db:"project_id"`
	ProjectName string                  `json:"project_name" db:"project_name"`
	PlotConfigs dbJSONSlice[IDSlugName] `json:"plot_configs" db:"plot_configs"`
	After       *time.Time              `json:"after" db:"after"`
	Before      *time.Time              `json:"before" db:"before"`
	AuditInfo
}

const createReportConfig = `
	INSERT INTO report_config (name, slug, project_id, after, before, creator, description)
	VALUES ($1, slugify($1, 'report_config'), $2, $3, $4, $5, $6)
	RETURNING id
`

func (q *Queries) CreateReportConfig(ctx context.Context, rc ReportConfig) (uuid.UUID, error) {
	var rcID uuid.UUID
	err := q.db.GetContext(ctx, &rcID, createReportConfig, rc.Name, rc.ProjectID, rc.After, rc.Before, rc.CreatorID, rc.Description)
	return rcID, err
}

const listProjectReportConfigs = `
	SELECT * FROM v_report_config WHERE project_id = $1
`

func (q *Queries) ListProjectReportConfigs(ctx context.Context, projectID uuid.UUID) ([]ReportConfig, error) {
	rcs := make([]ReportConfig, 0)
	err := q.db.SelectContext(ctx, &rcs, listProjectReportConfigs, projectID)
	return rcs, err
}

const getReportConfigByID = `
	SELECT * FROM v_report_config WHERE id = $1
`

func (q *Queries) GetReportConfigByID(ctx context.Context, rcID uuid.UUID) (ReportConfig, error) {
	var rc ReportConfig
	err := q.db.GetContext(ctx, &rc, getReportConfigByID, rcID)
	return rc, err
}

const updateReportConfig = `
	UPDATE report_config SET name=$2, after=$3, before=$4, updater=$5, update_date=$6, description=$7 WHERE id=$1
`

func (q *Queries) UpdateReportConfig(ctx context.Context, rc ReportConfig) error {
	_, err := q.db.ExecContext(ctx, updateReportConfig, rc.ID, rc.Name, rc.After, rc.Before, rc.UpdaterID, rc.UpdateDate, rc.Description)
	return err
}

const deleteReportConfig = `
	DELETE FROM report_config WHERE id=$1
`

func (q *Queries) DeleteReportConfig(ctx context.Context, rcID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteReportConfig, rcID)
	return err
}

const assignReportConfigPlotConfig = `
	INSERT INTO report_config_plot_config (report_config_id, plot_config_id) VALUES ($1, $2)
`

func (q *Queries) AssignReportConfigPlotConfig(ctx context.Context, rcID, pcID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, assignReportConfigPlotConfig, rcID, pcID)
	return err
}

const unassignReportConfigPlotConfig = `
	DELETE FROM report_config_plot_config WHERE report_config_id=$1 AND plot_config_id=$2
`

func (q *Queries) UnassignReportConfigPlotConfig(ctx context.Context, rcID, pcID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, assignReportConfigPlotConfig, rcID, pcID)
	return err
}

const unassignAllReportConfigPlotConfig = `
	DELETE FROM report_config_plot_config WHERE report_config_id=$1
`

func (q *Queries) UnassignAllReportConfigPlotConfig(ctx context.Context, rcID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, unassignAllInstrumentsFromAlertConfig, rcID)
	return err
}
