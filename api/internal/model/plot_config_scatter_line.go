package model

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

// PlotConfigScatterLinePlot holds information for entity PlotConfigScatterLinePlot
type PlotConfigScatterLinePlot struct {
	PlotConfig
	Display PlotConfigScatterLineDisplay `json:"display" db:"display"`
	// TODO AlertConfigIDs []string
}

type PlotConfigScatterLineDisplay struct {
	Traces []PlotConfigScatterLineTimeseriesTrace `json:"traces"`
	Layout PlotConfigScatterLineLayout            `json:"layout"`
}

func (d *PlotConfigScatterLineDisplay) Scan(src interface{}) error {
	b, ok := src.(string)
	if !ok {
		return fmt.Errorf("type assertion failed")
	}
	return json.Unmarshal([]byte(b), d)
}

type PlotConfigScatterLineTimeseriesTrace struct {
	PlotConfigurationID uuid.UUID `json:"plot_configuration_id"`
	TimeseriesID        uuid.UUID `json:"timeseries_id"`
	Name                string    `json:"name"`      // read-only
	Parameter           string    `json:"parameter"` // read-only
	TraceOrder          int       `json:"trace_order"`
	TraceType           string    `json:"trace_type"`
	Color               string    `json:"color"`
	LineStyle           string    `json:"line_style"`
	Width               float32   `json:"width"`
	ShowMarkers         bool      `json:"show_markers"`
	YAxis               string    `json:"y_axis"` // y1 or y2, default y1
}

type PlotConfigScatterLineLayout struct {
	CustomShapes       []PlotConfigScatterLineCustomShape `json:"custom_shapes"`
	YAxisTitle         *string                            `json:"yaxis_title"`
	SecondaryAxisTitle *string                            `json:"secondary_axis_title"`
}

type PlotConfigScatterLineCustomShape struct {
	PlotConfigurationID uuid.UUID `json:"plot_configuration_id"`
	Enabled             bool      `json:"enabled"`
	Name                string    `json:"name"`
	DataPoint           float32   `json:"data_point"`
	Color               string    `json:"color"`
}

const createPlotConfigScatterLineLayout = `INSERT INTO plot_scatter_line_config (plot_config_id, y_axis_title, y2_axis_title) VALUES ($1, $2, $3)`

func (q *Queries) CreatePlotConfigScatterLineLayout(ctx context.Context, pcID uuid.UUID, layout PlotConfigScatterLineLayout) error {
	_, err := q.db.ExecContext(ctx, createPlotConfigScatterLineLayout, pcID, layout.YAxisTitle, layout.SecondaryAxisTitle)
	return err
}

const updatePlotConfigScatterLineLayout = `UPDATE plot_scatter_line_config SET y_axis_title=$2, y2_axis_title=$3 WHERE plot_config_id=$1`

func (q *Queries) UpdatePlotConfigScatterLineLayout(ctx context.Context, pcID uuid.UUID, layout PlotConfigScatterLineLayout) error {
	_, err := q.db.ExecContext(ctx, updatePlotConfigScatterLineLayout, pcID, layout.YAxisTitle, layout.SecondaryAxisTitle)
	return err
}

// PlotConfigTimeseriesTrace
const createPlotConfigTimeseriesTrace = `
	INSERT INTO plot_configuration_timeseries_trace
	(plot_configuration_id, timeseries_id, trace_order, color, line_style, width, show_markers, y_axis) VALUES
	($1, $2, $3, $4, $5, $6, $7, $8)
`

func (q *Queries) CreatePlotConfigTimeseriesTrace(ctx context.Context, tr PlotConfigScatterLineTimeseriesTrace) error {
	_, err := q.db.ExecContext(
		ctx, createPlotConfigTimeseriesTrace,
		tr.PlotConfigurationID, tr.TimeseriesID, tr.TraceOrder, tr.Color, tr.LineStyle, tr.Width, tr.ShowMarkers, tr.YAxis,
	)
	return err
}

const updatePlotConfigTimeseriesTrace = `
	UPDATE plot_configuration_timeseries_trace
	SET trace_order=$3, color=$4, line_style=$5, width=$6, show_markers=$7, y_axis=$8
	WHERE plot_configuration_id=$1 AND timeseries_id=$2
`

func (q *Queries) UpdatePlotConfigTimeseriesTrace(ctx context.Context, tr PlotConfigScatterLineTimeseriesTrace) error {
	_, err := q.db.ExecContext(
		ctx, createPlotConfigTimeseriesTrace,
		tr.PlotConfigurationID, tr.TimeseriesID, tr.TraceOrder, tr.Color, tr.LineStyle, tr.Width, tr.ShowMarkers, tr.YAxis,
	)
	return err
}

const deleteAllPlotConfigTimeseriesTraces = `
	DELETE FROM plot_configuration_timeseries_trace WHERE plot_configuration_id=$1
`

func (q *Queries) DeleteAllPlotConfigTimeseriesTraces(ctx context.Context, pcID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteAllPlotConfigTimeseriesTraces, pcID)
	return err
}

// PlotConfigCustomShape
const createPlotConfigCustomShape = `
	INSERT INTO plot_configuration_custom_shape
	(plot_configuration_id, enabled, name, data_point, color) VALUES ($1, $2, $3, $4, $5)
`

func (q *Queries) CreatePlotConfigCustomShape(ctx context.Context, cs PlotConfigScatterLineCustomShape) error {
	_, err := q.db.ExecContext(
		ctx, createPlotConfigCustomShape,
		cs.PlotConfigurationID, cs.Enabled, cs.Name, cs.DataPoint, cs.Color,
	)
	return err
}

const updatePlotConfigCustomShape = `
	UPDATE plot_configuration_custom_shape
	SET enabled=$2, name=$3, data_point=$4, color=$5 WHERE plot_configuration_id=$1
`

func (q *Queries) UpdatePlotConfigCustomShape(ctx context.Context, cs PlotConfigScatterLineCustomShape) error {
	_, err := q.db.ExecContext(
		ctx, updatePlotConfigCustomShape,
		cs.PlotConfigurationID, cs.Enabled, cs.Name, cs.DataPoint, cs.Color,
	)
	return err
}

const deleteAllPlotConfigCustomShapes = `
	DELETE FROM plot_configuration_custom_shape WHERE plot_configuration_id=$1
`

func (q *Queries) DeleteAllPlotConfigCustomShapes(ctx context.Context, pcID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteAllPlotConfigCustomShapes, pcID)
	return err
}
