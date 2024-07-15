package model

import (
	"context"

	"github.com/google/uuid"
)

type PlotConfigBullseyePlot struct {
	PlotConfig
	Display PlotConfigBullseyePlotDisplay `json:"display" db:"display"`
}

type PlotConfigBullseyePlotDisplay struct {
	XAxisTimeseriesID uuid.UUID `json:"x_axis_timeseries_id" db:"x_axis_timeseries_id"`
	YAxisTimeseriesID uuid.UUID `json:"y_axis_timeseries_id" db:"y_axis_timeseries_id"`
}

func (d PlotConfigBullseyePlotDisplay) Display() {}

const createPlotBullseyeConfig = `
	INSERT INTO plot_bullseye_config (plot_config_id, x_axis_timeseries_id, y_axis_timeseries_id) VALUES ($1, $2, $3)
`

func (q *Queries) CreatePlotBullseyeConfig(ctx context.Context, plotConfigID uuid.UUID, cfg PlotConfigBullseyePlotDisplay) error {
	_, err := q.db.ExecContext(ctx, createPlotBullseyeConfig, plotConfigID, cfg.XAxisTimeseriesID, cfg.YAxisTimeseriesID)
	return err
}

const updatePlotBullseyeConfig = `
	UPDATE plot_bullseye_config SET x_axis_timeseries_id=$2, y_axis_timeseries_id=$3 WHERE plot_config_id=$1
`

func (q *Queries) UpdatePlotBullseyeConfig(ctx context.Context, plotConfigID uuid.UUID, cfg PlotConfigBullseyePlotDisplay) error {
	_, err := q.db.ExecContext(ctx, updatePlotBullseyeConfig, plotConfigID, cfg.XAxisTimeseriesID, cfg.YAxisTimeseriesID)
	return err
}

const deletePlotBullseyeConfig = `
	DELETE FROM plot_bullseye_config WHERE plog_config_id = $1
`

func (q *Queries) DeletePlotBullseyeConfig(ctx context.Context, plotConfigID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deletePlotBullseyeConfig, plotConfigID)
	return err
}
