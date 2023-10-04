package model

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/message"
	"github.com/USACE/instrumentation-api/api/internal/util"
	"github.com/google/uuid"
	"github.com/lib/pq"

	"github.com/paulmach/orb"
	"github.com/paulmach/orb/encoding/wkb"
	"github.com/paulmach/orb/encoding/wkt"
	"github.com/paulmach/orb/geojson"
)

// Instrument is an instrument
type Instrument struct {
	ID            uuid.UUID        `json:"id"`
	AwareID       *uuid.UUID       `json:"aware_id,omitempty"`
	Groups        []uuid.UUID      `json:"groups"`
	Constants     []uuid.UUID      `json:"constants"`
	AlertConfigs  []uuid.UUID      `json:"alert_configs"`
	StatusID      uuid.UUID        `json:"status_id" db:"status_id"`
	Status        string           `json:"status"`
	StatusTime    time.Time        `json:"status_time" db:"status_time"`
	Deleted       bool             `json:"-"`
	Slug          string           `json:"slug"`
	Name          string           `json:"name"`
	TypeID        uuid.UUID        `json:"type_id" db:"type_id"`
	Type          string           `json:"type"`
	Geometry      geojson.Geometry `json:"geometry,omitempty"`
	Station       *int             `json:"station"`
	StationOffset *int             `json:"offset" db:"station_offset"`
	ProjectID     *uuid.UUID       `json:"project_id" db:"project_id"`
	NIDID         *string          `json:"nid_id" db:"nid_id"`
	USGSID        *string          `json:"usgs_id" db:"usgs_id"`
	AuditInfo
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

// instrumentFactory converts database rows to Instrument objects
func instrumentFactory(rows DBRows) ([]Instrument, error) {
	defer rows.Close()
	ii := make([]Instrument, 0)
	for rows.Next() {
		var i Instrument
		var p orb.Point
		err := rows.Scan(
			&i.ID, &i.Deleted, &i.StatusID, &i.Status, &i.StatusTime, &i.Slug, &i.Name, &i.TypeID, &i.Type, wkb.Scanner(&p), &i.Station, &i.StationOffset,
			&i.Creator, &i.CreateDate, &i.Updater, &i.UpdateDate, &i.ProjectID, pq.Array(&i.Constants), pq.Array(&i.Groups), pq.Array(&i.AlertConfigs),
			&i.NIDID, &i.USGSID,
		)
		if err != nil {
			return nil, err
		}
		i.Geometry = *geojson.NewGeometry(p)
		ii = append(ii, i)
	}
	return ii, nil
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
		usgs_id
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
	rows, err := q.db.QueryxContext(ctx, listInstruments)
	if err != nil {
		return nil, err
	}
	return instrumentFactory(rows)
}

const getInstrument = listInstrumentsSQL + `
	WHERE id = $1
`

// GetInstrument returns a single instrument
func (q *Queries) GetInstrument(ctx context.Context, instrumentID uuid.UUID) (Instrument, error) {
	e := Instrument{}
	rows, err := q.db.QueryxContext(ctx, getInstrument, instrumentID)
	if err != nil {
		return e, err
	}
	ii, err := instrumentFactory(rows)
	if err != nil {
		return e, err
	}
	if len(ii) == 0 {
		return e, fmt.Errorf(message.NotFound)
	}
	return ii[0], nil
}

const getInstrumentCount = `
	SELECT COUNT(id) FROM instrument WHERE NOT deleted
`

// GetInstrumentCount returns the number of instruments in the database
func (q *Queries) GetInstrumentCount(ctx context.Context) (int, error) {
	var count int
	if err := q.db.GetContext(ctx, &count, getInstrumentCount); err != nil {
		return 0, err
	}
	return count, nil
}

const createInstrument = `
	INSERT INTO instrument (slug, name, type_id, geometry, station, station_offset, creator, create_date, project_id, nid_id, usgs_id)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	RETURNING id, slug
`

func (q *Queries) CreateInstrument(ctx context.Context, i Instrument) (IDAndSlug, error) {
	var aa IDAndSlug
	if err := q.db.GetContext(
		ctx, &aa, createInstrument,
		i.Slug, i.Name, i.TypeID, wkt.MarshalString(i.Geometry.Geometry()),
		i.Station, i.StationOffset, i.Creator, i.CreateDate, i.ProjectID, i.NIDID, i.USGSID,
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
	UPDATE instrument
	SET
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
		i.ProjectID, i.ID, i.Name, i.TypeID, wkb.Value(i.Geometry.Geometry()),
		i.Updater, i.UpdateDate, i.ProjectID, i.Station, i.StationOffset, i.NIDID, i.USGSID,
	)
	return err
}

const updateInstrumentGeometry = `
	UPDATE instrument SET geometry=ST_GeomFromWKB($3), updater= $4, update_date=now()
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
