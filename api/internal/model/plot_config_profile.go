package model

import (
	"context"

	"github.com/google/uuid"
)

type PlotConfigProfilePlot struct {
	PlotConfig
	Display PlotConfigProfilePlotDisplay `json:"display" db:"display"`
}

type PlotConfigProfilePlotDisplay struct {
	InstrumentID uuid.UUID `json:"instrument_id" db:"instrument_id"`
}

const createPlotProfileConfig = `
	INSERT INTO plot_profile_config (plot_config_id, instrument_id) VALUES ($1, $2)
`

func (q *Queries) CreatePlotProfileConfig(ctx context.Context, plotConfigID uuid.UUID, d PlotConfigProfilePlotDisplay) error {
	_, err := q.db.ExecContext(ctx, createPlotProfileConfig, plotConfigID, d.InstrumentID)
	return err
}

const updatePlotProfileConfig = `
	UPDATE plot_profile_config SET instrument_id=$2 WHERE plot_config_id=$1
`

func (q *Queries) UpdatePlotProfileConfig(ctx context.Context, plotConfigID uuid.UUID, d PlotConfigProfilePlotDisplay) error {
	_, err := q.db.ExecContext(ctx, updatePlotProfileConfig, plotConfigID, d.InstrumentID)
	return err
}
