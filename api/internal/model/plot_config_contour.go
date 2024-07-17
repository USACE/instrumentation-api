package model

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type PlotConfigContourPlot struct {
	PlotConfig
	Display PlotConfigContourPlotDisplay `json:"display" db:"display"`
}

type PlotConfigContourPlotDisplay struct {
	TimeseiresIDs     dbSlice[uuid.UUID] `json:"timeseries_ids" db:"timeseries_ids"`
	Time              time.Time          `json:"time" db:"time"`
	LocfBackfill      string             `json:"locf_backfill" db:"locf_backfill"`
	GradientSmoothing bool               `json:"gradient_smoothing" db:"gradient_smoothing"`
	ContourSmoothing  bool               `json:"contour_smoothing" db:"contour_smoothing"`
	ShowLabels        bool               `json:"show_labels" db:"show_labels"`
}

func (d *PlotConfigContourPlotDisplay) Scan(src interface{}) error {
	b, ok := src.(string)
	if !ok {
		return fmt.Errorf("type assertion failed")
	}
	return json.Unmarshal([]byte(b), d)
}

type PlotConfigMeasurementContourPlot struct {
	X float64  `json:"x" db:"x"`
	Y float64  `json:"y" db:"y"`
	Z *float64 `json:"z" db:"z"`
}

type AggregatePlotConfigMeasurementsContourPlot struct {
	X []float64  `json:"x" db:"x"`
	Y []float64  `json:"y" db:"y"`
	Z []*float64 `json:"z" db:"z"`
}

const createPlotContourConfig = `
	INSERT INTO plot_contour_config (plot_config_id, "time", locf_backfill, gradient_smoothing, contour_smoothing, show_labels) 
	VALUES ($1, $2, $3, $4, $5, $6)
`

func (q *Queries) CreatePlotContourConfig(ctx context.Context, plotConfigID uuid.UUID, cfg PlotConfigContourPlotDisplay) error {
	_, err := q.db.ExecContext(ctx, createPlotContourConfig, plotConfigID, cfg.Time, cfg.LocfBackfill, cfg.GradientSmoothing, cfg.ContourSmoothing, cfg.ShowLabels)
	return err
}

const updatePlotContourConfig = `
	UPDATE plot_contour_config SET "time"=$2, locf_backfill=$3, gradient_smoothing=$4, contour_smoothing=$5, show_labels=$6 
	WHERE plot_config_id=$1
`

func (q *Queries) UpdatePlotContourConfig(ctx context.Context, plotConfigID uuid.UUID, cfg PlotConfigContourPlotDisplay) error {
	_, err := q.db.ExecContext(ctx, updatePlotContourConfig, plotConfigID, cfg.Time, cfg.LocfBackfill, cfg.GradientSmoothing, cfg.ContourSmoothing, cfg.ShowLabels)
	return err
}

const deletePlotContourConfig = `
	DELETE FROM plot_contour_config WHERE plog_config_id = $1
`

func (q *Queries) DeletePlotContourConfig(ctx context.Context, plotConfigID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deletePlotContourConfig, plotConfigID)
	return err
}

const createPlotContourConfigTimeseries = `
	INSERT INTO plot_contour_config_timeseries (plot_contour_config_id, timeseries_id) VALUES ($1, $2)
	ON CONFLICT (plot_contour_config_id, timeseries_id) DO NOTHING
`

func (q *Queries) CreatePlotContourConfigTimeseries(ctx context.Context, plotConfigID, timeseriesID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, createPlotContourConfigTimeseries, plotConfigID, timeseriesID)
	return err
}

const deleteAllPlotContourConfigTimeseries = `
	DELETE FROM plot_contour_config_timeseries WHERE plot_contour_config_id = $1
`

func (q *Queries) DeleteAllPlotContourConfigTimeseries(ctx context.Context, plotConfigID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteAllPlotContourConfigTimeseries, plotConfigID)
	return err
}

const listPlotContourConfigTimes = `
	SELECT DISTINCT mm.time
	FROM plot_contour_config_timeseries pcts
	INNER JOIN timeseries_measurement mm ON mm.timeseries_id = pcts.timeseries_id
	WHERE pcts.plot_contour_config_id = $1
	AND mm.time > $2
	AND mm.time < $3
`

func (q *Queries) ListPlotConfigTimesContourPlot(ctx context.Context, plotConfigID uuid.UUID, tw TimeWindow) ([]time.Time, error) {
	tt := make([]time.Time, 0)
	err := q.db.SelectContext(ctx, &tt, listPlotContourConfigTimes, plotConfigID, tw.After, tw.Before)
	return tt, err
}

const listPlotConfigMeasurementsContourPlot = `
	SELECT
		oi.x,
		oi.y,
		locf(mm.value) AS z
	FROM plot_contour_config pc
	LEFT JOIN plot_contour_config_timeseries pcts ON pcts.plot_contour_config_id = pc.plot_config_id
	LEFT JOIN timeseries_measurement mm ON mm.timeseries_id = pcts.timeseries_id
	INNER JOIN timeseries ts ON ts.id = pcts.timeseries_id
	INNER JOIN (
		SELECT
			ii.id,
			ST_X(ST_Centroid(ST_Transform(ii.geometry, 4326))) AS x,
			ST_Y(ST_Centroid(ST_Transform(ii.geometry, 4326))) AS y
		FROM instrument ii
	) oi ON oi.id = ts.instrument_id
	WHERE plot_config_id = $1
	AND mm.time = $2
	GROUP BY pc.plot_config_id, pcts.timeseries_id, oi.x, oi.y
`

func (q *Queries) ListPlotConfigMeasurementsContourPlot(ctx context.Context, plotConfigID uuid.UUID, t time.Time) ([]PlotConfigMeasurementContourPlot, error) {
	var pcmm []PlotConfigMeasurementContourPlot
	err := q.db.SelectContext(ctx, pcmm, listPlotConfigMeasurementsContourPlot, plotConfigID, t)
	return pcmm, err
}
