package model

import (
	"context"
	"encoding/json"
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
	// OverridePlotConfigSettings PlotConfigSettings      `json:"override_plot_config_settings" db:"override_plot_config_settings"`
}

type ReportConfigWithPlotConfigs struct {
	ReportConfig
	PlotConfigs []PlotConfig `json:"plot_configs"`
}

type ReportConfigJobMessage struct {
	ReportConfigID uuid.UUID `json:"report_config_id"`
}

func (rj ReportConfigJobMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(rj)
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

const listReportConfigPlotConfigs = `
	SELECT * FROM v_plot_configuration WHERE id = ANY(
		SELECT plot_config_id FROM report_config_plot_config WHERE report_config_id = $1
	)
`

func (q *Queries) ListReportConfigPlotConfigs(ctx context.Context, rcID uuid.UUID) ([]PlotConfig, error) {
	pcs := make([]PlotConfig, 0)
	err := q.db.SelectContext(ctx, &pcs, listReportConfigPlotConfigs, rcID)
	return pcs, err
}

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

const createReportDownloadJob = `
	INSERT INTO report_download_job (job_id) VALUES ($1)
`

const updateReportDownloadJob = `
	UPDATE report_download_job SET status_id=$2, update_date=$3 WHERE job_id=$1
`

// const getReportConfigChartData = `
// 	SELECT
// 		rcpc.report_config_id,
// 		pc.name AS plot_config_name,
// 		ii.name || ' - ' || ts.name AS series_name,
// 		uu.name AS y_axis_unit,
// 		COALESCE(xy.values, '[]') AS xy_values
// 	FROM report_config_plot_config rcpc
// 	INNER JOIN plot_configuration pc ON pc.id = rcpc.plot_config_id
// 	INNER JOIN plot_configuration_timeseries pcts ON pcts.plot_configuration_id = pc.id
// 	INNER JOIN timeseries ts ON ts.id = pcts.timeseries_id
// 	INNER JOIN unit uu ON uu.id = ts.unit_id
// 	LEFT JOIN LATERAL (
// 		SELECT json_agg(json_build_object(
// 			'x', mm.time,
// 			'y', mm.value
// 		))::text AS values
// 		FROM timeseries_measurement mm
// 		WHERE mm.timeseries_id = ts.id
// 		AND mm.time > $2
// 		AND mm.time < $3
// 	) xy ON true
// 	WHERE rcpc.id = $1
// `
//
// func (q *Queries) getReportConfigChartData(ctx context.Context, rcID uuid.UUID, tw TimeWindow) ([]ReportConfigChartData, error) {
// 	cd := make([]ReportConfigChartData, 0)
// 	err := q.db.SelectContext(ctx, &cd, listReportConfigPlotConfigs, rcID, tw.After, tw.Before)
// 	return cd, err
// }
