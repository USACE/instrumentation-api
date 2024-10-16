package model

import (
	"context"
	"encoding/json"

	"github.com/USACE/instrumentation-api/api/internal/util"
	"github.com/google/uuid"
)

const (
	StandardTimeseriesType = "standard"
	ConstantTimeseriesType = "constant"
	ComputedTimeseriesType = "computed"
	CwmsTimeseriesType     = "cwms"
)

type Timeseries struct {
	ID             uuid.UUID     `json:"id"`
	Slug           string        `json:"slug"`
	Name           string        `json:"name"`
	Variable       string        `json:"variable"`
	InstrumentID   uuid.UUID     `json:"instrument_id" db:"instrument_id"`
	InstrumentSlug string        `json:"instrument_slug" db:"instrument_slug"`
	Instrument     string        `json:"instrument,omitempty"`
	ParameterID    uuid.UUID     `json:"parameter_id" db:"parameter_id"`
	Parameter      string        `json:"parameter,omitempty"`
	UnitID         uuid.UUID     `json:"unit_id" db:"unit_id"`
	Unit           string        `json:"unit,omitempty"`
	Values         []Measurement `json:"values,omitempty"`
	Type           string        `json:"type" db:"type"`
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

const getTimeseriesProjectMap = `
	SELECT timeseries_id, project_id
	FROM v_timeseries_project_map
	WHERE timeseries_id IN (?)
`

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

const listProjectTimeseries = `
	SELECT t.* FROM v_timeseries t
	INNER JOIN project_instrument p ON p.instrument_id = t.instrument_id
	WHERE p.project_id = $1
`

// ListProjectTimeseries lists all timeseries for a given project
func (q *Queries) ListProjectTimeseries(ctx context.Context, projectID uuid.UUID) ([]Timeseries, error) {
	tt := make([]Timeseries, 0)
	if err := q.db.SelectContext(ctx, &tt, listProjectTimeseries, projectID); err != nil {
		return make([]Timeseries, 0), err
	}

	return tt, nil
}

const listInstrumentTimeseries = `
	SELECT * FROM v_timeseries
	WHERE instrument_id = $1
`

func (q *Queries) ListInstrumentTimeseries(ctx context.Context, instrumentID uuid.UUID) ([]Timeseries, error) {
	tt := make([]Timeseries, 0)
	if err := q.db.Select(&tt, listInstrumentTimeseries, instrumentID); err != nil {
		return nil, err
	}
	return tt, nil
}

const listPlotConfigTimeseries = `
	SELECT t.* FROM v_timeseries t
	INNER JOIN plot_configuration_timeseries pct ON pct.timeseries_id = t.id
	WHERE pct.plot_configuration_id = $1
`

func (q *Queries) ListPlotConfigTimeseries(ctx context.Context, plotConfigID uuid.UUID) ([]Timeseries, error) {
	tt := make([]Timeseries, 0)
	if err := q.db.Select(&tt, listPlotConfigTimeseries, plotConfigID); err != nil {
		return nil, err
	}
	return tt, nil
}

const listInstrumentGroupTimeseries = `
	SELECT t.* FROM v_timeseries t
	INNER JOIN instrument_group_instruments gi ON gi.instrument_id = t.instrument_id
	WHERE gi.instrument_group_id = $1
`

func (q *Queries) ListInstrumentGroupTimeseries(ctx context.Context, instrumentGroupID uuid.UUID) ([]Timeseries, error) {
	tt := make([]Timeseries, 0)
	if err := q.db.SelectContext(ctx, &tt, listInstrumentGroupTimeseries, instrumentGroupID); err != nil {
		return nil, err
	}
	return tt, nil
}

const getTimeseries = `
	SELECT * FROM v_timeseries WHERE id = $1
`

func (q *Queries) GetTimeseries(ctx context.Context, timeseriesID uuid.UUID) (Timeseries, error) {
	var t Timeseries
	err := q.db.GetContext(ctx, &t, getTimeseries, timeseriesID)
	return t, err
}

const createTimeseries = `
	INSERT INTO timeseries (instrument_id, slug, name, parameter_id, unit_id, type)
	VALUES ($1, slugify($2, 'timeseries'), $2, $3, $4, $5)
	RETURNING id, instrument_id, slug, name, parameter_id, unit_id, type
`

func (q *Queries) CreateTimeseries(ctx context.Context, ts Timeseries) (Timeseries, error) {
	if ts.ParameterID == uuid.Nil {
		ts.ParameterID = unknownParameterID
	}
	if ts.UnitID == uuid.Nil {
		ts.UnitID = unknownUnitID
	}
	if ts.Type == "" {
		ts.Type = StandardTimeseriesType
	}
	var tsNew Timeseries
	err := q.db.GetContext(ctx, &tsNew, createTimeseries, ts.InstrumentID, ts.Name, ts.ParameterID, ts.UnitID, ts.Type)
	return tsNew, err
}

const updateTimeseries = `
	UPDATE timeseries SET name = $2, instrument_id = $3, parameter_id = $4, unit_id = $5
	WHERE id = $1
	RETURNING id
`

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

func (q *Queries) DeleteTimeseries(ctx context.Context, timeseriesID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteTimeseries, timeseriesID)
	return err
}
