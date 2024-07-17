package model

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

const (
	ScatterLinePlotType = "scatter-line"
	ProfilePlotType     = "profile"
	ContourPlotType     = "contour"
	BullseyePlotType    = "bullseye"
)

type PlotConfig struct {
	ID            uuid.UUID               `json:"id"`
	Name          string                  `json:"name"`
	Slug          string                  `json:"slug"`
	ProjectID     uuid.UUID               `json:"project_id" db:"project_id"`
	ReportConfigs dbJSONSlice[IDSlugName] `json:"report_configs" db:"report_configs"`
	PlotType      string                  `json:"plot_type" db:"plot_type"`
	Display       Opts                    `json:"display" db:"display"`
	PlotConfigSettings
	AuditInfo
}

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

// DateRangeTimeWindow creates a TimeWindow from a date range string.
//
// Acceptable date range strings are "lifetime", "5 years", "1 year", or a fixed date in the
// format "YYYY-MM-DD YYYY-MM-DD" with after and before dates separated by a single whitespace.
func (pc *PlotConfig) DateRangeTimeWindow() (TimeWindow, error) {
	switch dr := strings.ToLower(pc.DateRange); dr {
	case "lifetime":
		return TimeWindow{After: time.Time{}, Before: time.Now()}, nil
	case "5 years":
		return TimeWindow{After: time.Now().AddDate(-5, 0, 0), Before: time.Now()}, nil
	case "1 year":
		return TimeWindow{After: time.Now().AddDate(-1, 0, 0), Before: time.Now()}, nil
	default:
		cdr := strings.Split(dr, " ")
		invalidDateErr := fmt.Errorf("invalid date range; custom date range must be in format \"YYYY-MM-DD YYYY-MM-DD\"")
		if len(cdr) != 2 {
			return TimeWindow{}, invalidDateErr
		}
		after, err := time.Parse("2006-01-02", cdr[0])
		if err != nil {
			return TimeWindow{}, invalidDateErr
		}
		before, err := time.Parse("2006-01-02", cdr[1])
		if err != nil {
			return TimeWindow{}, invalidDateErr
		}
		return TimeWindow{After: after, Before: before}, nil
	}
}

const listPlotConfigsSQL = `
	SELECT
		id,
		slug,
		name,
		project_id,
		report_configs,
		creator,
		create_date,
		updater,
		update_date,
		show_masked,
		show_nonvalidated,
		show_comments,
		auto_range,
		date_range,
		threshold,
		display
	FROM v_plot_configuration
`

// PlotConfig
const listPlotConfigs = listPlotConfigsSQL + `
	WHERE project_id = $1
`

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

func (q *Queries) GetPlotConfig(ctx context.Context, plotConfigID uuid.UUID) (PlotConfig, error) {
	var pc PlotConfig
	err := q.db.GetContext(ctx, &pc, getPlotConfig, plotConfigID)
	return pc, err
}

const createPlotConfig = `
	INSERT INTO plot_configuration (slug, name, project_id, creator, create_date, plot_type) VALUES (slugify($1, 'plot_configuration'), $1, $2, $3, $4, $5)
	RETURNING id
`

func (q *Queries) CreatePlotConfig(ctx context.Context, pc PlotConfig) (uuid.UUID, error) {
	var pcID uuid.UUID
	err := q.db.GetContext(ctx, &pcID, createPlotConfig, pc.Name, pc.ProjectID, pc.CreatorID, pc.CreateDate, pc.PlotType)
	return pcID, err
}

// PlotConfigSettings
const createPlotConfigSettings = `
	INSERT INTO plot_configuration_settings (id, show_masked, show_nonvalidated, show_comments, auto_range, date_range, threshold) 
	VALUES ($1, $2, $3, $4, $5, $6, $7)
`

func (q *Queries) CreatePlotConfigSettings(ctx context.Context, pcID uuid.UUID, pc PlotConfigSettings) error {
	_, err := q.db.ExecContext(ctx, createPlotConfigSettings, pcID, pc.ShowMasked, pc.ShowNonValidated, pc.ShowComments, pc.AutoRange, pc.DateRange, pc.Threshold)
	return err
}

const updatePlotConfig = `
	UPDATE plot_configuration SET name = $3, updater = $4, update_date = $5 WHERE project_id = $1 AND id = $2
`

func (q *Queries) UpdatePlotConfig(ctx context.Context, pc PlotConfig) error {
	_, err := q.db.ExecContext(ctx, updatePlotConfig, pc.ProjectID, pc.ID, pc.Name, pc.UpdaterID, pc.UpdateDate)
	return err
}

const deletePlotConfig = `
	DELETE from plot_configuration WHERE project_id = $1 AND id = $2
`

func (q *Queries) DeletePlotConfig(ctx context.Context, projectID, plotConfigID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deletePlotConfig, projectID, plotConfigID)
	return err
}

const deletePlotConfigSettings = `
	DELETE FROM plot_configuration_settings WHERE id = $1
`

func (q *Queries) DeletePlotConfigSettings(ctx context.Context, plotConfigID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deletePlotConfigSettings, plotConfigID)
	return err
}
