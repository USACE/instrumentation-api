package model

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/util"
	"github.com/google/uuid"

	"github.com/paulmach/orb"
	"github.com/paulmach/orb/encoding/wkb"
	"github.com/paulmach/orb/geojson"
)

// Instrument is an instrument
type Instrument struct {
	ID            uuid.UUID          `json:"id"`
	AwareID       *uuid.UUID         `json:"aware_id,omitempty"`
	Groups        dbSlice[uuid.UUID] `json:"groups" db:"groups"`
	Constants     dbSlice[uuid.UUID] `json:"constants" db:"constants"`
	AlertConfigs  dbSlice[uuid.UUID] `json:"alert_configs" db:"alert_configs"`
	StatusID      uuid.UUID          `json:"status_id" db:"status_id"`
	Status        string             `json:"status"`
	StatusTime    time.Time          `json:"status_time" db:"status_time"`
	Deleted       bool               `json:"-"`
	Slug          string             `json:"slug"`
	Name          string             `json:"name"`
	TypeID        uuid.UUID          `json:"type_id" db:"type_id"`
	Type          string             `json:"type"`
	Geometry      Geometry           `json:"geometry,omitempty"`
	Station       *int               `json:"station"`
	StationOffset *int               `json:"offset" db:"station_offset"`
	ProjectID     *uuid.UUID         `json:"project_id" db:"project_id"`
	NIDID         *string            `json:"nid_id" db:"nid_id"`
	USGSID        *string            `json:"usgs_id" db:"usgs_id"`
	Opts          Opts               `json:"opts" db:"opts"`
	AuditInfo
}

// Optional instrument metadata based on type
// If there are no options defined for the instrument type, the object will be empty
type Opts map[string]interface{}

func (o *Opts) Scan(src interface{}) error {
	b, ok := src.(string)
	if !ok {
		return fmt.Errorf("type assertion failed")
	}
	return json.Unmarshal([]byte(b), o)
}

// CreateInstrumentsValidationResult holds results of checking InstrumentCollection POST
type CreateInstrumentsValidationResult struct {
	IsValid bool     `json:"is_valid"`
	Errors  []string `json:"errors"`
}

// InstrumentCollection is a collection of Instrument items
type InstrumentCollection struct {
	Items []Instrument
}

// Shorten returns an instrument collection with individual objects limited to ID and Struct fields
func (c InstrumentCollection) Shorten() IDAndSlugCollection {
	ss := IDAndSlugCollection{Items: make([]IDAndSlug, 0)}
	for _, n := range c.Items {
		s := IDAndSlug{ID: n.ID, Slug: n.Slug}

		ss.Items = append(ss.Items, s)
	}
	return ss
}

// UnmarshalJSON implements UnmarshalJSON interface
func (c *InstrumentCollection) UnmarshalJSON(b []byte) error {
	switch util.JSONType(b) {
	case "ARRAY":
		if err := json.Unmarshal(b, &c.Items); err != nil {
			return err
		}
	case "OBJECT":
		var n Instrument
		if err := json.Unmarshal(b, &n); err != nil {
			return err
		}
		c.Items = []Instrument{n}
	default:
		c.Items = make([]Instrument, 0)
	}
	return nil
}

type InstrumentCount struct {
	InstrumentCount int `json:"instrument_count"`
}

type Geometry geojson.Geometry

func (g Geometry) Value() (driver.Value, error) {
	og := geojson.Geometry(g)
	return wkb.Value(og.Geometry()), nil
}

func (g *Geometry) Scan(src interface{}) error {
	var p orb.Point
	if err := wkb.Scanner(&p).Scan(src); err != nil {
		return err
	}
	*g = Geometry(*geojson.NewGeometry(p))
	return nil
}

func (g Geometry) MarshalJSON() ([]byte, error) {
	gj := geojson.Geometry(g)
	return gj.MarshalJSON()
}

func (g *Geometry) UnmarshalJSON(data []byte) error {
	gj, err := geojson.UnmarshalGeometry(data)
	if err != nil {
		return err
	}
	if gj == nil {
		return fmt.Errorf("unable to unmarshal: geojson geometry is nil")
	}
	*g = Geometry(*gj)
	return nil
}

const listInstrumentsSQL = `
	SELECT 
		id,
		deleted,
		status_id,
		status,
		status_time,
		slug,
		name,
		type_id,
		type,
		geometry,
		station,
		station_offset,
		creator,
		create_date,
		updater,
		update_date,
		project_id,
		constants,
		groups,
		alert_configs,
		nid_id,
		usgs_id,
		opts
	FROM v_instrument
`

const listInstrumentSlugs = `
	SELECT slug FROM instrument
`

// ListInstrumentSlugs lists used instrument slugs in the database
func (q *Queries) ListInstrumentSlugs(ctx context.Context) ([]string, error) {
	ss := make([]string, 0)
	if err := q.db.SelectContext(ctx, &ss, listInstrumentSlugs); err != nil {
		return nil, err
	}
	return ss, nil
}

const listInstruments = listInstrumentsSQL + `
	WHERE NOT deleted
`

// ListInstruments returns an array of instruments from the database
func (q *Queries) ListInstruments(ctx context.Context) ([]Instrument, error) {
	ii := make([]Instrument, 0)
	if err := q.db.SelectContext(ctx, &ii, listInstruments); err != nil {
		return nil, err
	}
	return ii, nil
}

const getInstrument = listInstrumentsSQL + `
	WHERE id = $1
`

// GetInstrument returns a single instrument
func (q *Queries) GetInstrument(ctx context.Context, instrumentID uuid.UUID) (Instrument, error) {
	var i Instrument
	err := q.db.GetContext(ctx, &i, getInstrument, instrumentID)
	return i, err
}

const getInstrumentCount = `
	SELECT COUNT(id) FROM instrument WHERE NOT deleted
`

// GetInstrumentCount returns the number of instruments in the database
func (q *Queries) GetInstrumentCount(ctx context.Context) (InstrumentCount, error) {
	var ic InstrumentCount
	if err := q.db.GetContext(ctx, &ic.InstrumentCount, getInstrumentCount); err != nil {
		return ic, err
	}
	return ic, nil
}

const createInstrument = `
	INSERT INTO instrument (slug, name, type_id, geometry, station, station_offset, creator, create_date, project_id, nid_id, usgs_id)
	VALUES ($1, $2, $3, ST_GeomFromWKB($4), $5, $6, $7, $8, $9, $10, $11)
	RETURNING id, slug
`

func (q *Queries) CreateInstrument(ctx context.Context, i Instrument) (IDAndSlug, error) {
	var aa IDAndSlug
	if err := q.db.GetContext(
		ctx, &aa, createInstrument,
		i.Slug, i.Name, i.TypeID, i.Geometry, i.Station, i.StationOffset, i.Creator, i.CreateDate, i.ProjectID, i.NIDID, i.USGSID,
	); err != nil {
		return aa, err
	}
	return aa, nil
}

const validateCreateInstruments = ` 
	SELECT project_id, name
	FROM instrument
	WHERE project_id IN (?)
	AND NOT deleted
	ORDER BY project_id
`

// ValidateCreateInstruments creates many instruments from an array of instruments
func (q *Queries) ValidateCreateInstruments(ctx context.Context, instruments []Instrument) (CreateInstrumentsValidationResult, error) {
	validationResult := CreateInstrumentsValidationResult{Errors: make([]string, 0)}
	projectIDs := make([]uuid.UUID, 0)
	for idx := range instruments {
		projectIDs = append(projectIDs, *instruments[idx].ProjectID)
	}
	query, args, err := sqlIn(validateCreateInstruments, projectIDs)
	if err != nil {
		return validationResult, err
	}
	var nn []struct {
		ProjectID      uuid.UUID `db:"project_id"`
		InstrumentName string    `db:"name"`
	}
	if err := q.db.SelectContext(ctx, &nn, q.db.Rebind(query), args...); err != nil {
		return validationResult, err
	}
	m := make(map[uuid.UUID]map[string]bool)
	var _pID uuid.UUID
	for _, n := range nn {
		if n.ProjectID != _pID {
			m[n.ProjectID] = make(map[string]bool)
			_pID = n.ProjectID
		}
		m[n.ProjectID][strings.ToUpper(n.InstrumentName)] = true
	}
	validationResult.IsValid = true
	for _, n := range instruments {
		if !m[*n.ProjectID][strings.ToUpper(n.Name)] {
			continue
		}
		validationResult.IsValid = false
		validationResult.Errors = append(
			validationResult.Errors,
			fmt.Sprintf("Instrument name '%s' is already taken. Instrument names must be unique within a project", n.Name),
		)
	}
	return validationResult, nil
}

const updateInstrument = `
	UPDATE instrument SET
		name = $3,
		type_id = $4,
		geometry = ST_GeomFromWKB($5),
		updater = $6,
		update_date = $7,
		project_id = $8,
		station = $9,
		station_offset = $10,
		nid_id = $11,
		usgs_id = $12
	WHERE project_id = $1 AND id = $2
`

func (q *Queries) UpdateInstrument(ctx context.Context, i Instrument) error {
	_, err := q.db.ExecContext(
		ctx, updateInstrument,
		i.ProjectID, i.ID, i.Name, i.TypeID, i.Geometry,
		i.Updater, i.UpdateDate, i.ProjectID, i.Station, i.StationOffset, i.NIDID, i.USGSID,
	)
	return err
}

const updateInstrumentGeometry = `
	UPDATE instrument SET
		geometry = ST_GeomFromWKB($3),
		updater = $4,
		update_date = NOW()
	WHERE project_id = $1 AND id = $2
	RETURNING id
`

// UpdateInstrumentGeometry updates instrument geometry property
func (q *Queries) UpdateInstrumentGeometry(ctx context.Context, projectID, instrumentID uuid.UUID, geom geojson.Geometry, p Profile) error {
	_, err := q.db.ExecContext(ctx, updateInstrumentGeometry, projectID, instrumentID, wkb.Value(geom.Geometry()), p.ID)
	return err
}

const deleteFlagInstrument = `
	UPDATE instrument SET deleted = true WHERE project_id = $1 AND id = $2
`

// DeleteFlagInstrument changes delete flag to true
func (q *Queries) DeleteFlagInstrument(ctx context.Context, projectID, instrumentID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteFlagInstrument, projectID, instrumentID)
	return err
}
