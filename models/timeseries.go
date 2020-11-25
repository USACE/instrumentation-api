package models

import (
	"encoding/json"

	ts "github.com/USACE/instrumentation-api/timeseries"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	// pq database driver
	_ "github.com/lib/pq"
)

// TimeseriesCollection is a collection of Timeseries items
type TimeseriesCollection struct {
	Items []ts.Timeseries
}

// UnmarshalJSON implements UnmarshalJSON interface
func (c *TimeseriesCollection) UnmarshalJSON(b []byte) error {
	switch JSONType(b) {
	case "ARRAY":
		if err := json.Unmarshal(b, &c.Items); err != nil {
			return err
		}
	case "OBJECT":
		var t ts.Timeseries
		if err := json.Unmarshal(b, &t); err != nil {
			return err
		}
		c.Items = []ts.Timeseries{t}
	default:
		c.Items = make([]ts.Timeseries, 0)
	}
	return nil
}

// ListTimeseriesSlugs lists used timeseries slugs in the database
func ListTimeseriesSlugs(db *sqlx.DB) ([]string, error) {

	ss := make([]string, 0)
	if err := db.Select(&ss, "SELECT slug FROM v_timeseries"); err != nil {
		return make([]string, 0), err
	}
	return ss, nil
}

// GetTimeseriesProjectMap returns a map of { timeseries_id: project_id, }
func GetTimeseriesProjectMap(db *sqlx.DB, timeseriesIDS []uuid.UUID) (map[uuid.UUID]uuid.UUID, error) {
	query, args, err := sqlx.In(
		`SELECT timeseries_id, project_id
		 FROM v_timeseries_project_map
		 WHERE timeseries_id IN (?);`,
		timeseriesIDS,
	)
	if err != nil {
		return nil, err
	}
	query = db.Rebind(query)

	// struct to temporarily hold SQL Query Result
	var result []struct {
		TimeseriesID uuid.UUID `db:"timeseries_id"`
		ProjectID    uuid.UUID `db:"project_id"`
	}
	if err = db.Select(&result, query, args...); err != nil {
		return nil, err
	}
	m := make(map[uuid.UUID]uuid.UUID)
	for _, r := range result {
		m[r.TimeseriesID] = r.ProjectID
	}
	return m, nil
}

// ListTimeseriesSlugsForInstrument lists used timeseries slugs for a given instrument
func ListTimeseriesSlugsForInstrument(db *sqlx.DB, id *uuid.UUID) ([]string, error) {

	ss := make([]string, 0)
	if err := db.Select(&ss, "SELECT slug FROM v_timeseries WHERE instrument_id = $1", id); err != nil {
		return make([]string, 0), err
	}
	return ss, nil
}

// ListTimeseries lists all timeseries
func ListTimeseries(db *sqlx.DB) ([]ts.Timeseries, error) {

	tt := make([]ts.Timeseries, 0)
	if err := db.Select(&tt, "SELECT * FROM v_timeseries"); err != nil {
		return make([]ts.Timeseries, 0), err
	}
	return tt, nil
}

// ListInstrumentTimeseries returns an array of timeseries for an instrument
func ListInstrumentTimeseries(db *sqlx.DB, projectID *uuid.UUID, instrumentID *uuid.UUID) ([]ts.Timeseries, error) {
	tt := make([]ts.Timeseries, 0)
	if err := db.Select(&tt, "SELECT * FROM v_timeseries WHERE project_id = $1 AND instrument_id = $2", projectID, instrumentID); err != nil {
		return make([]ts.Timeseries, 0), err
	}
	return tt, nil
}

// ListInstrumentGroupTimeseries returns an array of timeseries for instruments that belong to an instrument_group
func ListInstrumentGroupTimeseries(db *sqlx.DB, instrumentGroupID *uuid.UUID) ([]ts.Timeseries, error) {

	var tt []ts.Timeseries
	if err := db.Select(
		&tt,
		`SELECT *
		 FROM   v_timeseries
		 WHERE  instrument_id IN (
			SELECT instrument_id
			FROM   instrument_group_instruments
			WHERE  instrument_group_id = $1
		)`, instrumentGroupID,
	); err != nil {
		return make([]ts.Timeseries, 0), err
	}
	return tt, nil
}

// GetTimeseries returns a single timeseries without measurements
func GetTimeseries(db *sqlx.DB, id *uuid.UUID) (*ts.Timeseries, error) {

	var t ts.Timeseries
	if err := db.Get(&t, "SELECT * FROM v_timeseries WHERE id = $1", id); err != nil {
		return nil, err
	}
	return &t, nil
}

// CreateTimeseries creates many timeseries from an array of timeseries
func CreateTimeseries(db *sqlx.DB, tt []ts.Timeseries) ([]ts.Timeseries, error) {

	txn, err := db.Beginx()
	if err != nil {
		return nil, err
	}

	// Insert Timeseries
	stmt, err := txn.Preparex(
		`INSERT INTO timeseries (instrument_id, slug, name, parameter_id, unit_id)
		 VALUES ($1, $2, $3, $4, $5)
		 RETURNING id, instrument_id, slug, name, parameter_id, unit_id`,
	)
	if err != nil {
		return make([]ts.Timeseries, 0), err
	}

	// Insert
	uu := make([]ts.Timeseries, len(tt))
	for idx, t := range tt {
		if err := stmt.Get(&uu[idx], t.InstrumentID, t.Slug, t.Name, t.ParameterID, t.UnitID); err != nil {
			return make([]ts.Timeseries, 0), err
		}
	}

	if err := stmt.Close(); err != nil {
		return make([]ts.Timeseries, 0), err
	}

	if err := txn.Commit(); err != nil {
		return make([]ts.Timeseries, 0), err
	}

	return uu, nil
}

// UpdateTimeseries updates a timeseries
func UpdateTimeseries(db *sqlx.DB, t *ts.Timeseries) (*ts.Timeseries, error) {

	var tUpdated ts.Timeseries
	if err := db.QueryRowx(
		`UPDATE timeseries
		 SET    name = $2,
			    instrument_id = $3,
			    parameter_id = $4,
			    unit_id = $5
		 WHERE id = $1
		 RETURNING *
		`, t.ID, t.Name, t.InstrumentID, t.ParameterID, t.UnitID,
	).StructScan(&tUpdated); err != nil {
		return nil, err
	}

	return &tUpdated, nil
}

// DeleteTimeseries deletes a timeseries and cascade deletes all measurements
func DeleteTimeseries(db *sqlx.DB, id *uuid.UUID) error {
	if _, err := db.Exec("DELETE FROM timeseries WHERE id = $1", id); err != nil {
		return err
	}
	return nil
}
