package model

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type ReportConfig struct {
	ID              uuid.UUID                   `json:"id" db:"id"`
	Slug            string                      `json:"slug" db:"slug"`
	Name            string                      `json:"name" db:"name"`
	Description     string                      `json:"description" db:"description"`
	ProjectID       uuid.UUID                   `json:"project_id" db:"project_id"`
	ProjectName     string                      `json:"project_name" db:"project_name"`
	DistrictName    *string                     `json:"district_name" db:"district_name"`
	PlotConfigs     dbJSONSlice[IDSlugName]     `json:"plot_configs" db:"plot_configs"`
	GlobalOverrides ReportConfigGlobalOverrides `json:"global_overrides" db:"global_overrides"`
	AuditInfo
}

type ReportConfigGlobalOverrides struct {
	DateRange        TextOption   `json:"date_range" db:"date_range"`
	ShowMasked       ToggleOption `json:"show_masked" db:"show_masked"`
	ShowNonvalidated ToggleOption `json:"show_nonvalidated" db:"show_nonvalidated"`
}

type TextOption struct {
	Enabled bool   `json:"enabled" db:"enabled"`
	Value   string `json:"value" db:"value"`
}

type ToggleOption struct {
	Enabled bool `json:"enabled" db:"enabled"`
	Value   bool `json:"value" db:"value"`
}

type ReportDownloadJob struct {
	ID                 uuid.UUID  `json:"id" db:"id"`
	ReportConfigID     uuid.UUID  `json:"report_config_id" db:"report_config_id"`
	Creator            uuid.UUID  `json:"creator" db:"creator"`
	CreateDate         time.Time  `json:"create_date" db:"create_date"`
	Status             string     `json:"status" db:"status"`
	FileKey            *string    `json:"file_key" db:"file_key"`
	FileExpiry         *time.Time `json:"file_expiry" db:"file_expiry"`
	Progress           int        `json:"progress" db:"progress"`
	ProgressUpdateDate time.Time  `json:"progress_update_date" db:"progress_update_date"`
}

func (o *ReportConfigGlobalOverrides) Scan(src interface{}) error {
	b, ok := src.(string)
	if !ok {
		return fmt.Errorf("type assertion failed")
	}
	return json.Unmarshal([]byte(b), o)
}

type ReportConfigWithPlotConfigs struct {
	ReportConfig
	PlotConfigs []PlotConfigScatterLinePlot `json:"plot_configs"`
}

type ReportConfigJobMessage struct {
	ReportConfigID uuid.UUID `json:"report_config_id"`
	JobID          uuid.UUID `json:"job_id"`
	IsLandscape    bool      `json:"is_landscape"`
}

const createReportConfig = `
	INSERT INTO report_config (
		name, slug, project_id, creator, description, date_range, date_range_enabled,
		show_masked, show_masked_enabled, show_nonvalidated, show_nonvalidated_enabled
	)
	VALUES ($1, slugify($1, 'report_config'), $2, $3, $4, $5, $6, $7, $8, $9, $10)
	RETURNING id
`

func (q *Queries) CreateReportConfig(ctx context.Context, rc ReportConfig) (uuid.UUID, error) {
	var rcID uuid.UUID
	err := q.db.GetContext(
		ctx, &rcID, createReportConfig, rc.Name, rc.ProjectID, rc.CreatorID, rc.Description,
		rc.GlobalOverrides.DateRange.Value, rc.GlobalOverrides.DateRange.Enabled,
		rc.GlobalOverrides.ShowMasked.Value, rc.GlobalOverrides.ShowMasked.Enabled,
		rc.GlobalOverrides.ShowNonvalidated.Value, rc.GlobalOverrides.ShowNonvalidated.Enabled,
	)
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

const listReportConfigPlotConfigs = `
	SELECT * FROM v_plot_configuration WHERE id = ANY(
		SELECT plot_config_id FROM report_config_plot_config WHERE report_config_id = $1
	)
`

func (q *Queries) ListReportConfigPlotConfigs(ctx context.Context, rcID uuid.UUID) ([]PlotConfigScatterLinePlot, error) {
	pcs := make([]PlotConfigScatterLinePlot, 0)
	err := q.db.SelectContext(ctx, &pcs, listReportConfigPlotConfigs, rcID)
	return pcs, err
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
	UPDATE report_config SET name=$2,
	updater=$3, update_date=$4, description=$5, date_range=$6, date_range_enabled=$7, show_masked=$8,
	show_masked_enabled=$9, show_nonvalidated=$10, show_nonvalidated_enabled=$11 WHERE id=$1
`

func (q *Queries) UpdateReportConfig(ctx context.Context, rc ReportConfig) error {
	_, err := q.db.ExecContext(
		ctx, updateReportConfig, rc.ID, rc.Name, rc.UpdaterID, rc.UpdateDate, rc.Description,
		rc.GlobalOverrides.DateRange.Value, rc.GlobalOverrides.DateRange.Enabled,
		rc.GlobalOverrides.ShowMasked.Value, rc.GlobalOverrides.ShowMasked.Enabled,
		rc.GlobalOverrides.ShowNonvalidated.Value, rc.GlobalOverrides.ShowNonvalidated.Enabled,
	)
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
	_, err := q.db.ExecContext(ctx, unassignAllReportConfigPlotConfig, rcID)
	return err
}

const getReportDownloadJob = `
	SELECT * FROM report_download_job WHERE id=$1 AND creator=$2
`

func (q *Queries) GetReportDownloadJob(ctx context.Context, jobID, profileID uuid.UUID) (ReportDownloadJob, error) {
	var j ReportDownloadJob
	err := q.db.GetContext(ctx, &j, getReportDownloadJob, jobID, profileID)
	return j, err
}

const createReportDownloadJob = `
	INSERT INTO report_download_job (report_config_id, creator) VALUES ($1, $2) RETURNING *
`

func (q *Queries) CreateReportDownloadJob(ctx context.Context, rcID, profileID uuid.UUID) (ReportDownloadJob, error) {
	var jNew ReportDownloadJob
	err := q.db.GetContext(ctx, &jNew, createReportDownloadJob, rcID, profileID)
	return jNew, err
}

const updateReportDownloadJob = `
	UPDATE report_download_job SET status=$2, progress=$3, progress_update_date=$4, file_key=$5, file_expiry=$6 WHERE id=$1
`

func (q *Queries) UpdateReportDownloadJob(ctx context.Context, j ReportDownloadJob) error {
	_, err := q.db.ExecContext(
		ctx, updateReportDownloadJob,
		j.ID, j.Status, j.Progress, j.ProgressUpdateDate, j.FileKey, j.FileExpiry,
	)
	return err
}
