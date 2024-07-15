package model

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type PlotBullseyeMeasurements struct {
	PlotConfigID uuid.UUID `json:"plot_config_id" db:"-"`
	Measurements []PlotBullseyeMeasurement
}

type PlotBullseyeMeasurement struct {
	Time time.Time `json:"time" db:"time"`
	X    *float64  `json:"x" db:"x"`
	Y    *float64  `json:"y" db:"y"`
}

type PlotContourMeasurements struct {
	PlotConfigID   uuid.UUID `json:"plot_config_id" db:"plot_config_id"`
	Time           time.Time `json:"time" db:"time"`
	InstrumentName string    `json:"instrument_name" db:"instrument_name"`
	TimeseriesName string    `json:"timeseries_name" db:"timeseries_name"`
	X              float64   `json:"x" db:"x"`
	Y              float64   `json:"y" db:"y"`
	Z              *float64  `json:"z" db:"z"`
}

const listPlotBullseyeMeasurements = `
	SELECT time, x, y FROM v_plot_bullseye_measurement WHERE plot_config_id = $1 AND time > $2 AND time < $3
`

func (q *Queries) GetPlotBullseyeMeasurementsForPlotConfig(ctx context.Context, plotConfigID uuid.UUID, tw TimeWindow) (PlotBullseyeMeasurements, error) {
	var pcmm []PlotBullseyeMeasurement
	err := q.db.SelectContext(ctx, pcmm, listPlotBullseyeMeasurements, plotConfigID, tw.After, tw.Before)
	return PlotBullseyeMeasurements{PlotConfigID: plotConfigID, Measurements: pcmm}, err
}

const getPlotContourMeasurements = `
	SELECT
		plot_config_id,
		timeseries_id,
		name,
		long,
		lat,
		locf(time) FILTER (WHERE time >= ($3::timestamptz - $4::interval)) OVER (ORDER BY time ASC) AS time,
		locf(value) FILTER (WHERE time >= ($3::timestamptz - $4::interval)) OVER (ORDER BY time ASC) AS value
	FROM v_plot_contour_measurement WHERE plot_config_id = $1 AND timeseries_id = $2 AND time = $3
`
