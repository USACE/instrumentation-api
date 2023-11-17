package model

import (
	"context"
	"encoding/json"

	"github.com/USACE/instrumentation-api/api/internal/util"
	"github.com/google/uuid"
)

// Timeseries is a timeseries data structure
type Timeseries struct {
	ID             uuid.UUID     `json:"id"`
	Slug           string        `json:"slug"`
	Name           string        `json:"name"`
	Variable       string        `json:"variable"`
	ProjectID      uuid.UUID     `json:"project_id" db:"project_id"`
	ProjectSlug    string        `json:"project_slug" db:"project_slug"`
	Project        string        `json:"project,omitempty" db:"project"`
	InstrumentID   uuid.UUID     `json:"instrument_id" db:"instrument_id"`
	InstrumentSlug string        `json:"instrument_slug" db:"instrument_slug"`
	Instrument     string        `json:"instrument,omitempty"`
	ParameterID    uuid.UUID     `json:"parameter_id" db:"parameter_id"`
	Parameter      string        `json:"parameter,omitempty"`
	UnitID         uuid.UUID     `json:"unit_id" db:"unit_id"`
	Unit           string        `json:"unit,omitempty"`
	Values         []Measurement `json:"values,omitempty"`
	IsComputed     bool          `json:"is_computed" db:"is_computed"`
}

type TimeseriesNote struct {
	Masked     *bool   `json:"masked,omitempty"`
	Validated  *bool   `json:"validated,omitempty"`
	Annotation *string `json:"annotation,omitempty"`
}

type TimeseriesCollectionItems struct {
	Items []Timeseries
}

// UnmarshalJSON implements UnmarshalJSON interface
func (c *TimeseriesCollectionItems) UnmarshalJSON(b []byte) error {
	switch util.JSONType(b) {
	case "ARRAY":
		if err := json.Unmarshal(b, &c.Items); err != nil {
			return err
		}
	case "OBJECT":
		var t Timeseries
		if err := json.Unmarshal(b, &t); err != nil {
			return err
		}
		c.Items = []Timeseries{t}
	default:
		c.Items = make([]Timeseries, 0)
	}
	return nil
}

var (
	unknownParameterID = uuid.MustParse("2b7f96e1-820f-4f61-ba8f-861640af6232")
	unknownUnitID      = uuid.MustParse("4a999277-4cf5-4282-93ce-23b33c65e2c8")
)

const listTimeseries = `
	SELECT
		id, slug, name, variable, project_id, project_slug, project, instrument_id,
		instrument_slug, instrument, parameter_id, parameter, unit_id, unit, is_computed
	FROM v_timeseries
`

// ListTimeseries lists all timeseries
func (q *Queries) ListTimeseries(ctx context.Context) ([]Timeseries, error) {
	tt := make([]Timeseries, 0)
	if err := q.db.SelectContext(ctx, &tt, listTimeseries); err != nil {
		return make([]Timeseries, 0), err
	}
	return tt, nil
}

const getStoredTimeseriesExists = `
	SELECT EXISTS (SELECT id FROM v_timeseries_stored WHERE id = $1)
`

// ValidateStoredTimeseries returns an error if the timeseries id does not exist or the timeseries is computed
func (q *Queries) GetStoredTimeseriesExists(ctx context.Context, timeseriesID uuid.UUID) (bool, error) {
	var isStored bool
	if err := q.db.GetContext(ctx, &isStored, getStoredTimeseriesExists, &timeseriesID); err != nil {
		return false, err
	}
	return isStored, nil
}

const listTimeseriesSlugs = `
	SELECT slug FROM v_timeseries
`

// ListTimeseriesSlugs lists used timeseries slugs in the database
func (q *Queries) ListTimeseriesSlugs(ctx context.Context) ([]string, error) {
	ss := make([]string, 0)
	if err := q.db.SelectContext(ctx, &ss, listTimeseriesSlugs); err != nil {
		return make([]string, 0), err
	}
	return ss, nil
}

const getTimeseriesProjectMap = `
	SELECT timeseries_id, project_id
	FROM v_timeseries_project_map
	WHERE timeseries_id IN (?)
`

// GetTimeseriesProjectMap returns a map of { timeseries_id: project_id, }
func (q *Queries) GetTimeseriesProjectMap(ctx context.Context, timeseriesIDs []uuid.UUID) (map[uuid.UUID]uuid.UUID, error) {
	query, args, err := sqlIn(getTimeseriesProjectMap, timeseriesIDs)
	if err != nil {
		return nil, err
	}
	query = q.db.Rebind(query)
	var result []struct {
		TimeseriesID uuid.UUID `db:"timeseries_id"`
		ProjectID    uuid.UUID `db:"project_id"`
	}
	if err = q.db.SelectContext(ctx, &result, query, args...); err != nil {
		return nil, err
	}
	m := make(map[uuid.UUID]uuid.UUID)
	for _, r := range result {
		m[r.TimeseriesID] = r.ProjectID
	}
	return m, nil
}

const listTimeseriesSlugsForInstrument = `
	SELECT slug FROM v_timeseries WHERE instrument_id = $1
`

// ListTimeseriesSlugsForInstrument lists used timeseries slugs for a given instrument
func (q *Queries) ListTimeseriesSlugsForInstrument(ctx context.Context, instrumentID uuid.UUID) ([]string, error) {
	ss := make([]string, 0)
	if err := q.db.SelectContext(ctx, &ss, listTimeseriesSlugsForInstrument, instrumentID); err != nil {
		return nil, err
	}
	return ss, nil
}

const listProjectTimeseries = listTimeseries + `
	WHERE project_id = $1
`

// ListProjectTimeseries lists all timeseries for a given project
func (q *Queries) ListProjectTimeseries(ctx context.Context, projectID uuid.UUID) ([]Timeseries, error) {
	tt := make([]Timeseries, 0)
	if err := q.db.SelectContext(ctx, &tt, listProjectTimeseries, projectID); err != nil {
		return make([]Timeseries, 0), err
	}
	return tt, nil
}

const listInstrumentTimeseries = listTimeseries + `
	WHERE instrument_id = $1
`

// ListInstrumentTimeseries returns an array of timeseries for an instrument
func (q *Queries) ListInstrumentTimeseries(ctx context.Context, instrumentID uuid.UUID) ([]Timeseries, error) {
	tt := make([]Timeseries, 0)
	if err := q.db.Select(&tt, listInstrumentTimeseries, instrumentID); err != nil {
		return nil, err
	}
	return tt, nil
}

const listInstrumentGroupTimeseries = listTimeseries + `
	WHERE  instrument_id IN (
		SELECT instrument_id
		FROM   instrument_group_instruments
		WHERE  instrument_group_id = $1
	)
`

// ListInstrumentGroupTimeseries returns an array of timeseries for instruments that belong to an instrument_group
func (q *Queries) ListInstrumentGroupTimeseries(ctx context.Context, instrumentGroupID uuid.UUID) ([]Timeseries, error) {
	tt := make([]Timeseries, 0)
	if err := q.db.SelectContext(ctx, &tt, listInstrumentGroupTimeseries, instrumentGroupID); err != nil {
		return nil, err
	}
	return tt, nil
}

const getTimeseries = listTimeseries + `
	WHERE id = $1
`

// GetTimeseries returns a single timeseries without measurements
func (q *Queries) GetTimeseries(ctx context.Context, timeseriesID uuid.UUID) (Timeseries, error) {
	var t Timeseries
	err := q.db.GetContext(ctx, &t, getTimeseries, timeseriesID)
	return t, err
}

const createTimeseries = `
	INSERT INTO timeseries (instrument_id, slug, name, parameter_id, unit_id)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id, instrument_id, slug, name, parameter_id, unit_id
`

// CreateTimeseries creates many timeseries from an array of timeseries
func (q *Queries) CreateTimeseries(ctx context.Context, ts Timeseries) (Timeseries, error) {
	if ts.ParameterID == uuid.Nil {
		ts.ParameterID = unknownParameterID
	}
	if ts.UnitID == uuid.Nil {
		ts.UnitID = unknownUnitID
	}
	var tsNew Timeseries
	err := q.db.GetContext(ctx, &tsNew, createTimeseries, ts.InstrumentID, ts.Slug, ts.Name, ts.ParameterID, ts.UnitID)
	return tsNew, err
}

const updateTimeseries = `
	UPDATE timeseries SET name = $2, instrument_id = $3, parameter_id = $4, unit_id = $5
	WHERE id = $1
	RETURNING id
`

// UpdateTimeseries updates a timeseries
func (q *Queries) UpdateTimeseries(ctx context.Context, ts Timeseries) (uuid.UUID, error) {
	if ts.ParameterID == uuid.Nil {
		ts.ParameterID = unknownParameterID
	}
	if ts.UnitID == uuid.Nil {
		ts.UnitID = unknownUnitID
	}
	var tID uuid.UUID
	err := q.db.GetContext(ctx, &tID, updateTimeseries, ts.ID, ts.Name, ts.InstrumentID, ts.ParameterID, ts.UnitID)
	return tID, err
}

const deleteTimeseries = `
	DELETE FROM timeseries WHERE id = $1
`

// DeleteTimeseries deletes a timeseries and cascade deletes all measurements
func (q *Queries) DeleteTimeseries(ctx context.Context, timeseriesID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteTimeseries, timeseriesID)
	return err
}
