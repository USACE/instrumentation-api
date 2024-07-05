package model

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// CREATE VIEW v_plot_bullseye_measurement AS (
//   SELECT
//     pc.plot_config_id,
//     COALESCE(xm.time, ym.time) AS time,
//     locf(xm.value) OVER (ORDER BY xm.time) AS x,
//     locf(ym.value) OVER (ORDER BY ym.time) AS y
//   FROM plot_bullseye_config pc
//   LEFT JOIN timeseries_measurement xm ON xm.timeseries_id = pc.x_timeseries_id
//   LEFT JOIN timeseries_measurement ym ON ym.timeseries_id = pc.y_timeseries_id
// );
//
// CREATE VIEW v_plot_contour_measurement AS (
//   SELECT
//     pc.plot_config_id,
//     mm.time AS time,
//     ii.name AS instrument_name,
//     ts.name AS timeseries_name,
//     ST_X(ST_Centroid(ST_Transform(ii.geometry, 4326))) AS x,
//     ST_Y(ST_Centroid(ST_Transform(ii.geometry, 4326))) AS y,
//     mm.value AS z
//   FROM plot_contour_config pc
//   LEFT JOIN plot_contour_config_timeseries pcts ON pcts.plot_config_id = pc.plot_config_id
//   LEFT JOIN timeseries ts ON ts.id = pcts.timeseries_id
//   LEFT JOIN instrument ii ON ts.instrument_id = ii.id
//   LEFT JOIN measurement mm ON mm.timeseries_id = ts.id
//   GROUP BY pc.plot_config_id, mm.time
// );

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
