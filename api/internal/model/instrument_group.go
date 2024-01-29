package model

import (
	"context"
	"encoding/json"

	"github.com/USACE/instrumentation-api/api/internal/util"
	"github.com/google/uuid"
)

// InstrumentGroup holds information for entity instrument_group
type InstrumentGroup struct {
	ID              uuid.UUID  `json:"id"`
	Deleted         bool       `json:"-"`
	Slug            string     `json:"slug"`
	Name            string     `json:"name"`
	Description     string     `json:"description"`
	ProjectID       *uuid.UUID `json:"project_id" db:"project_id"`
	InstrumentCount int        `json:"instrument_count" db:"instrument_count"`
	TimeseriesCount int        `json:"timeseries_count" db:"timeseries_count"`
	AuditInfo
}

// InstrumentGroupCollection is a collection of Instrument items
type InstrumentGroupCollection struct {
	Items []InstrumentGroup
}

// Shorten returns an instrument collection with individual objects limited to ID and Struct fields
func (c InstrumentGroupCollection) Shorten() IDSlugCollection {
	ss := IDSlugCollection{Items: make([]IDSlug, 0)}
	for _, n := range c.Items {
		s := IDSlug{ID: n.ID, Slug: n.Slug}
		ss.Items = append(ss.Items, s)
	}
	return ss
}

// UnmarshalJSON implements UnmarshalJSON interface
// Allows unpacking object or array of objects into array of objects
func (c *InstrumentGroupCollection) UnmarshalJSON(b []byte) error {
	switch util.JSONType(b) {
	case "ARRAY":
		if err := json.Unmarshal(b, &c.Items); err != nil {
			return err
		}
	case "OBJECT":
		var g InstrumentGroup
		if err := json.Unmarshal(b, &g); err != nil {
			return err
		}
		c.Items = []InstrumentGroup{g}
	default:
		c.Items = make([]InstrumentGroup, 0)
	}
	return nil
}

const listInstrumentGroupsSQL = `
	SELECT
		id,
		slug,
		name,
		description,
		creator,
		create_date,
		updater,
		update_date,
		project_id,
		instrument_count,
		timeseries_count 
	FROM  v_instrument_group
`

const listInstrumentGroups = listInstrumentGroupsSQL + `
	WHERE NOT deleted
`

// ListInstrumentGroups returns a list of instrument groups
func (q *Queries) ListInstrumentGroups(ctx context.Context) ([]InstrumentGroup, error) {
	gg := make([]InstrumentGroup, 0)
	if err := q.db.SelectContext(ctx, &gg, listInstrumentGroups); err != nil {
		return make([]InstrumentGroup, 0), err
	}
	return gg, nil
}

const getInstrumentGroup = listInstrumentGroupsSQL + `
	WHERE id = $1
`

// GetInstrumentGroup returns a single instrument group
func (q *Queries) GetInstrumentGroup(ctx context.Context, instrumentGroupID uuid.UUID) (InstrumentGroup, error) {
	var g InstrumentGroup
	if err := q.db.GetContext(ctx, &g, getInstrumentGroup, instrumentGroupID); err != nil {
		return g, err
	}
	return g, nil
}

const createInstrumentGroup = `
	INSERT INTO instrument_group (slug, name, description, creator, create_date, project_id)
	VALUES (slugify($1, 'instrument_group'), $1, $2, $3, $4, $5)
	RETURNING id, slug, name, description, creator, create_date, updater, update_date, project_id
`

func (q *Queries) CreateInstrumentGroup(ctx context.Context, group InstrumentGroup) (InstrumentGroup, error) {
	var groupNew InstrumentGroup
	err := q.db.GetContext(
		ctx, &groupNew, createInstrumentGroup,
		group.Name, group.Description, group.CreatorID, group.CreateDate, group.ProjectID,
	)
	return groupNew, err
}

const updateInstrumentGroup = `
	UPDATE instrument_group SET
		name = $2,
		deleted = $3,
		description = $4,
		updater = $5,
		update_date = $6,
		project_id = $7
	 WHERE id = $1
	 RETURNING *
`

// UpdateInstrumentGroup updates an instrument group
func (q *Queries) UpdateInstrumentGroup(ctx context.Context, group InstrumentGroup) (InstrumentGroup, error) {
	var groupUpdated InstrumentGroup
	err := q.db.GetContext(
		ctx, &groupUpdated, updateInstrumentGroup,
		group.ID, group.Name, group.Deleted, group.Description, group.UpdaterID, group.UpdateDate, group.ProjectID,
	)
	return groupUpdated, err
}

const deleteFlagInstrumentGroup = `
	UPDATE instrument_group SET deleted = true WHERE id = $1
`

// DeleteFlagInstrumentGroup sets the deleted field to true
func (q *Queries) DeleteFlagInstrumentGroup(ctx context.Context, instrumentGroupID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteFlagInstrumentGroup, instrumentGroupID)
	return err
}

const listInstrumentGroupInstruments = `
	SELECT inst.*
	FROM   instrument_group_instruments igi
	INNER JOIN (` + listInstrumentsSQL + `) inst ON igi.instrument_id = inst.id
	WHERE igi.instrument_group_id = $1 and inst.deleted = false
`

// ListInstrumentGroupInstruments returns a list of instrument group instruments for a given instrument
func (q *Queries) ListInstrumentGroupInstruments(ctx context.Context, groupID uuid.UUID) ([]Instrument, error) {
	ii := make([]Instrument, 0)
	if err := q.db.SelectContext(ctx, &ii, listInstrumentGroupInstruments, groupID); err != nil {
		return nil, err
	}
	return ii, nil
}

const createInstrumentGroupInstruments = `
	INSERT INTO instrument_group_instruments (instrument_group_id, instrument_id) VALUES ($1, $2)
`

// CreateInstrumentGroupInstruments adds an instrument to an instrument group
func (q *Queries) CreateInstrumentGroupInstruments(ctx context.Context, instrumentGroupID uuid.UUID, instrumentID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, createInstrumentGroupInstruments, instrumentGroupID, instrumentID)
	return err
}

const deleteInstrumentGroupInstruments = `
	DELETE FROM instrument_group_instruments WHERE instrument_group_id = $1 and instrument_id = $2
`

// DeleteInstrumentGroupInstruments adds an instrument to an instrument group
func (q *Queries) DeleteInstrumentGroupInstruments(ctx context.Context, instrumentGroupID uuid.UUID, instrumentID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteInstrumentGroupInstruments, instrumentGroupID, instrumentID)
	return err
}
