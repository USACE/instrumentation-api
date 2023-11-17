package model

import (
	"context"
)

// Home is information for the homepage (landing page)
type Home struct {
	InstrumentCount     int `json:"instrument_count" db:"instrument_count"`
	InstrumetGroupCount int `json:"instrument_group_count" db:"instrument_group_count"`
	ProjectCount        int `json:"project_count" db:"project_count"`
	NewInstruments7D    int `json:"new_instruments_7d" db:"new_instruments_7d"`
	NewMeasurements2H   int `json:"new_measurements_2h" db:"new_measurements_2h"`
}

const getHome = `
	SELECT
		(SELECT COUNT(*) FROM instrument WHERE NOT deleted) AS instrument_count,
		(SELECT COUNT(*) FROM project WHERE NOT deleted) AS project_count,
		(SELECT COUNT(*) FROM instrument_group) AS instrument_group_count,
		(SELECT COUNT(*) FROM instrument WHERE NOT deleted AND create_date > NOW() - '7 days'::INTERVAL) AS new_instruments_7d,
		(SELECT COUNT(*) FROM timeseries_measurement WHERE time > NOW() - '2 hours'::INTERVAL) AS new_measurements_2h
`

// GetHome returns information for the homepage
func (q *Queries) GetHome(ctx context.Context) (Home, error) {
	var home Home
	err := q.db.GetContext(ctx, &home, getHome)
	return home, err
}
