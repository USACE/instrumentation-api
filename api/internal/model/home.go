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
		(SELECT count(id) FROM instrument WHERE NOT deleted)        AS instrument_count,
		(SELECT count(id) FROM project WHERE NOT deleted)           AS project_count,
		(SELECT count(id) FROM instrument_group)                    AS instrument_group_count,
		(SELECT count(id) FROM instrument WHERE NOT deleted and (now() - create_date) < '7 Days') AS new_instruments_7d,
		(SELECT count(timeseries_id) FROM timeseries_measurement WHERE (now() - timeseries_measurement.time) < '2 Hours' ) AS new_measurements_2h
`

// GetHome returns information for the homepage
func (q *Queries) GetHome(ctx context.Context) (Home, error) {
	var home Home
	err := q.db.GetContext(ctx, &home, getHome)
	return home, err
}
