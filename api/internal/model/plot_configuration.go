package model

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

// PlotConfigSettings describes options for displaying the plot consistently.
// Specifically, whether to ignore data entries in a timeseries that have been masked,
// or whether to display user comments.
type PlotConfigSettings struct {
	ShowMasked       bool   `json:"show_masked" db:"show_masked"`
	ShowNonValidated bool   `json:"show_nonvalidated" db:"show_nonvalidated"`
	ShowComments     bool   `json:"show_comments" db:"show_comments"`
	AutoRange        bool   `json:"auto_range" db:"auto_range"`
	DateRange        string `json:"date_range" db:"date_range"`
	Threshold        int    `json:"threshold" db:"threshold"`
}

// PlotConfig holds information for entity PlotConfig
type PlotConfig struct {
	ID            uuid.UUID          `json:"id"`
	Name          string             `json:"name"`
	Slug          string             `json:"slug"`
	ProjectID     uuid.UUID          `json:"project_id" db:"project_id"`
	TimeseriesIDs dbSlice[uuid.UUID] `json:"timeseries_id" db:"timeseries_id"`
	AuditInfo
	PlotConfigSettings
}

func (pc *PlotConfig) ValidateDateRange() error {
	dr := strings.ToLower(pc.DateRange)
	// Check for standard settings
	if dr == "lifetime" {
		return nil
	}
	if dr == "5 years" {
		return nil
	}
	if dr == "1 year" {
		return nil
	}
	cdr := strings.Split(dr, " - ")
	if len(cdr) == 2 {
		for _, v := range cdr {
			if _, err := time.Parse("01/02/2006", v); err != nil {
				return fmt.Errorf("custom date values must be in format \"MM/DD/YYYY - MM/DD/YYYY\"")
			}
		}
		return nil
	}
	return fmt.Errorf("invalid date range provided")
}

// listPlotConfigsSQL is the base SQL statement for above functions
const listPlotConfigsSQL = `
	SELECT
		id,
		slug,
		name,
		project_id,
		timeseries_id,
		creator,
		create_date,
		updater,
		update_date,
		show_masked,
		show_nonvalidated,
		show_comments,
		auto_range,
		date_range,
		threshold
	FROM v_plot_configuration
`

const listPlotConfigSlugs = `
	SELECT slug FROM plot_configuration
`

// ListPlotConfigSlugs lists used instrument group slugs in the database
func (q *Queries) ListPlotConfigSlugs(ctx context.Context) ([]string, error) {
	ss := make([]string, 0)
	if err := q.db.SelectContext(ctx, &ss, listPlotConfigSlugs); err != nil {
		return nil, err
	}
	return ss, nil
}

const listPlotConfigs = listPlotConfigsSQL + `
	WHERE project_id = $1
`

// ListPlotConfigs returns a list of Plot groups
func (q *Queries) ListPlotConfigs(ctx context.Context, projectID uuid.UUID) ([]PlotConfig, error) {
	ppc := make([]PlotConfig, 0)
	if err := q.db.SelectContext(ctx, &ppc, listPlotConfigs, projectID); err != nil {
		return make([]PlotConfig, 0), err
	}
	return ppc, nil
}

const getPlotConfig = listPlotConfigsSQL + `
	WHERE id = $1
`

// GetPlotConfig returns a single plot configuration
func (q *Queries) GetPlotConfig(ctx context.Context, plotconfigID uuid.UUID) (PlotConfig, error) {
	var pc PlotConfig
	err := q.db.GetContext(ctx, &pc, getPlotConfig, plotconfigID)
	return pc, err
}

const createPlotConfig = `
	INSERT INTO plot_configuration (slug, name, project_id, creator, create_date) VALUES ($1, $2, $3, $4, $5)
	RETURNING id
`

func (q *Queries) CreatePlotConfig(ctx context.Context, pc PlotConfig) (uuid.UUID, error) {
	var pcID uuid.UUID
	err := q.db.GetContext(ctx, &pcID, createPlotConfig, pc.Slug, pc.Name, pc.ProjectID, pc.Creator, pc.CreateDate)
	return pcID, err
}

const createPlotConfigTimeseries = `
	INSERT INTO plot_configuration_timeseries (plot_configuration_id, timeseries_id) VALUES ($1, $2)
	ON CONFLICT DO NOTHING
`

func (q *Queries) CreatePlotConfigTimeseries(ctx context.Context, pcID, timeseriesID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, createPlotConfigTimeseries, pcID, timeseriesID)
	return err
}

const createPlotConfigSettings = `
	INSERT INTO plot_configuration_settings (id, show_masked, show_nonvalidated, show_comments, auto_range, date_range, threshold) 
	VALUES ($1, $2, $3, $4, $5, $6, $7)
`

func (q *Queries) CreatePlotConfigSettings(ctx context.Context, pc PlotConfig) error {
	_, err := q.db.ExecContext(ctx, createPlotConfigSettings, pc.ID, pc.ShowMasked, pc.ShowNonValidated, pc.ShowComments, pc.AutoRange, pc.DateRange, pc.Threshold)
	return err
}

const updatePlotConfig = `
	UPDATE plot_configuration SET name = $3, updater = $4, update_date = $5 WHERE project_id = $1 AND id = $2
`

func (q *Queries) UpdatePlotConfig(ctx context.Context, pc PlotConfig) error {
	_, err := q.db.ExecContext(ctx, updatePlotConfig, pc.ProjectID, pc.ID, pc.Name, pc.Updater, pc.UpdateDate)
	return err
}

const deletePlotConfig = `
	DELETE from plot_configuration WHERE project_id = $1 AND id = $2
`

// DeletePlotConfig delete plot configuration for a project
func (q *Queries) DeletePlotConfig(ctx context.Context, projectID, plotConfigID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deletePlotConfig, projectID, plotConfigID)
	return err
}

const deletePlotConfigTimeseries = `
	DELETE FROM plot_configuration_timeseries WHERE plot_configuration_id = $1
`

func (q *Queries) DeletePlotConfigTimeseries(ctx context.Context, plotConfigID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deletePlotConfigTimeseries, plotConfigID)
	return err
}

const deletePlotConfigSettings = `
	DELETE FROM plot_configuration_settings WHERE id = $1
`

// DeletePlotConfiguration delete plot configuration for a project
func (q *Queries) DeletePlotConfigSettings(ctx context.Context, plotConfigID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deletePlotConfigSettings, plotConfigID)
	return err
}
