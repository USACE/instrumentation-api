package model

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

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

func (d *PlotConfigBullseyePlotDisplay) Scan(src interface{}) error {
	b, ok := src.(string)
	if !ok {
		return fmt.Errorf("type assertion failed")
	}
	return json.Unmarshal([]byte(b), d)
}

type PlotConfigMeasurementBullseyePlot struct {
	Time time.Time `json:"time" db:"time"`
	X    *float64  `json:"x" db:"x"`
	Y    *float64  `json:"y" db:"y"`
}

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

const listPlotConfigMeasurementsBullseyePlot = `
	SELECT
		t.time,
		locf(xm.value) AS x,
		locf(ym.value) AS y
	FROM plot_bullseye_config pc
	INNER JOIN timeseries_measurement t
	ON t.timeseries_id = pc.x_axis_timeseries_id
	OR t.timeseries_id = pc.y_axis_timeseries_id
	LEFT JOIN timeseries_measurement xm
	ON xm.timeseries_id = pc.x_axis_timeseries_id
	AND xm.time = t.time
	LEFT JOIN timeseries_measurement ym
	ON ym.timeseries_id = pc.y_axis_timeseries_id
	AND ym.time = t.time
	WHERE pc.plot_config_id = $1
	AND t.time > $2
	AND t.time < $3
	GROUP BY t.time
	ORDER BY t.time ASC
`

func (q *Queries) ListPlotConfigMeasurementsBullseyePlot(ctx context.Context, plotConfigID uuid.UUID, tw TimeWindow) ([]PlotConfigMeasurementBullseyePlot, error) {
	pcmm := make([]PlotConfigMeasurementBullseyePlot, 0)
	err := q.db.SelectContext(ctx, &pcmm, listPlotConfigMeasurementsBullseyePlot, plotConfigID, tw.After, tw.Before)
	return pcmm, err
}
