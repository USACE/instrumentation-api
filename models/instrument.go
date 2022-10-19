package models

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
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
	Formulas      []Formula        `json:"formulas"`
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

	switch JSONType(b) {
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

// ListInstrumentSlugs lists used instrument slugs in the database
func ListInstrumentSlugs(db *sqlx.DB) ([]string, error) {

	ss := make([]string, 0)
	if err := db.Select(&ss, "SELECT slug FROM instrument"); err != nil {
		return make([]string, 0), err
	}
	return ss, nil
}

// ListInstruments returns an array of instruments from the database
func ListInstruments(db *sqlx.DB) ([]Instrument, error) {

	rows, err := db.Queryx(listInstrumentsSQL + " WHERE NOT deleted")
	if err != nil {
		return make([]Instrument, 0), err
	}
	return InstrumentsFactory(rows)
}

// GetInstrument returns a single instrument
func GetInstrument(db *sqlx.DB, id *uuid.UUID) (*Instrument, error) {

	rows, err := db.Queryx(listInstrumentsSQL+" WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	ii, err := InstrumentsFactory(rows)
	if err != nil {
		return nil, err
	}

	return &ii[0], nil
}

// GetInstrumentCount returns the number of instruments in the database
func GetInstrumentCount(db *sqlx.DB) (int, error) {
	var count int
	if err := db.Get(&count, "SELECT COUNT(id) FROM instrument WHERE NOT deleted"); err != nil {
		return 0, err
	}
	return count, nil
}

// CreateInstruments creates many instruments from an array of instruments
func CreateInstruments(db *sqlx.DB, instruments []Instrument) ([]IDAndSlug, error) {

	txn, err := db.Beginx()
	if err != nil {
		return make([]IDAndSlug, 0), err
	}

	// Instrument
	stmt1, err := txn.Preparex(
		`INSERT INTO instrument
			(slug, name, type_id, geometry, station, station_offset, creator, create_date, project_id, formula, formula_name, nid_id, usgs_id)
		 VALUES
			 ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
		 RETURNING id, slug`,
	)
	if err != nil {
		return make([]IDAndSlug, 0), err
	}

	// Instrument Status
	stmt2, err := txn.Preparex(createInstrumentStatusSQL())
	if err != nil {
		return make([]IDAndSlug, 0), err
	}

	// AWARE Gage Info (if provided)
	stmt3, err := txn.Preparex(`INSERT INTO aware_platform (instrument_id, aware_id) VALUES ($1, $2)`)
	if err != nil {
		return make([]IDAndSlug, 0), err
	}

	ii := make([]IDAndSlug, len(instruments))
	for idx, i := range instruments {
		if err := stmt1.Get(
			&ii[idx],
			i.Slug, i.Name, i.TypeID, wkt.MarshalString(i.Geometry.Geometry()),
			i.Station, i.StationOffset, i.Creator, i.CreateDate, i.ProjectID, i.Formula, i.FormulaName, i.NIDID, i.USGSID,
		); err != nil {
			return make([]IDAndSlug, 0), err
		}
		if _, err := stmt2.Exec(ii[idx].ID, i.StatusID, i.StatusTime); err != nil {
			return make([]IDAndSlug, 0), err
		}
		// Store Aware ID if Provided
		if i.AwareID != nil {
			if _, err := stmt3.Exec(ii[idx].ID, i.AwareID); err != nil {
				return make([]IDAndSlug, 0), err
			}
		}
	}
	if err := stmt1.Close(); err != nil {
		return make([]IDAndSlug, 0), err
	}
	if err := stmt2.Close(); err != nil {
		return make([]IDAndSlug, 0), err
	}
	if err := stmt3.Close(); err != nil {
		return make([]IDAndSlug, 0), err
	}
	if err := txn.Commit(); err != nil {
		return make([]IDAndSlug, 0), err
	}
	return ii, nil
}

// ValidateCreateInstruments creates many instruments from an array of instruments
func ValidateCreateInstruments(db *sqlx.DB, instruments []Instrument) (CreateInstrumentsValidationResult, error) {

	validationResult := CreateInstrumentsValidationResult{Errors: make([]string, 0)}

	// Project IDs associated with instruments
	projectIDs := make([]uuid.UUID, 0)
	for idx := range instruments {
		projectIDs = append(projectIDs, *instruments[idx].ProjectID)
	}

	// Get Map of Taken Instrument Names by Project
	namesMap, err := projectInstrumentNamesMap(db, projectIDs)
	if err != nil {
		return validationResult, err
	}

	// Check that instrument names are unique name within project
	validationResult.IsValid = true // Start with assumption that POST is valid
	for _, n := range instruments {
		if namesMap[*n.ProjectID][strings.ToUpper(n.Name)] != true {
			continue
		}
		// Add message to Errors and make sure isValid is false
		validationResult.IsValid = false
		validationResult.Errors = append(
			validationResult.Errors,
			fmt.Sprintf("Instrument name '%s' is already taken. Instrument names must be unique within a project", n.Name),
		)
	}
	return validationResult, nil
}

// UpdateInstrument updates a single instrument
func UpdateInstrument(db *sqlx.DB, i *Instrument) (*Instrument, error) {

	txn, err := db.Begin()
	if err != nil {
		return nil, err
	}

	// Instrument
	stmt1, err := txn.Prepare(
		`UPDATE instrument
		 SET    name = $3,
			    type_id = $4,
			    geometry = ST_GeomFromWKB($5),
			    updater = $6,
				update_date = $7,
				project_id = $8,
				station = $9,
				station_offset = $10,
				formula = $11,
				formula_name = $12,
				nid_id = $13,
				usgs_id = $14
		 WHERE project_id = $1 AND id = $2
		 RETURNING id`,
	)
	// Update Instrument
	var updatedID uuid.UUID
	if err := stmt1.QueryRow(
		i.ProjectID, i.ID, i.Name, i.TypeID, wkb.Value(i.Geometry.Geometry()),
		i.Updater, i.UpdateDate, i.ProjectID, i.Station, i.StationOffset, i.Formula, i.FormulaName, i.NIDID, i.USGSID,
	).Scan(&updatedID); err != nil {
		return nil, err
	}
	if err := stmt1.Close(); err != nil {
		return nil, err
	}

	// Instrument Status
	stmt2, err := txn.Prepare(createInstrumentStatusSQL())
	if err != nil {
		return nil, err
	}
	if _, err := stmt2.Exec(i.ID, i.StatusID, i.StatusTime); err != nil {
		return nil, err
	}
	if err := stmt2.Close(); err != nil {
		return nil, err
	}

	if err := txn.Commit(); err != nil {
		return nil, err
	}

	// Get Updated Row
	return GetInstrument(db, &updatedID)
}

// UpdateInstrumentGeometry updates instrument geometry property
func UpdateInstrumentGeometry(db *sqlx.DB, projectID *uuid.UUID, instrumentID *uuid.UUID, geom *geojson.Geometry, p *Profile) (*Instrument, error) {
	var nID uuid.UUID
	if err := db.Get(
		&nID, `UPDATE instrument SET geometry=ST_GeomFromWKB($3), updater=$4, update_date=now()
			   WHERE project_id=$1 AND id=$2
			   RETURNING id`, projectID, instrumentID, wkb.Value(geom.Geometry()), p.ID,
	); err != nil {
		return nil, err
	}
	return GetInstrument(db, &nID)
}

// DeleteFlagInstrument changes delete flag to true
func DeleteFlagInstrument(db *sqlx.DB, projectID, instrumentID *uuid.UUID) error {

	if _, err := db.Exec(
		`UPDATE instrument SET deleted = true WHERE project_id = $1 AND id = $2`,
		projectID, instrumentID,
	); err != nil {
		return err
	}
	return nil
}

// InstrumentsFactory converts database rows to Instrument objects
func InstrumentsFactory(rows *sqlx.Rows) ([]Instrument, error) {
	defer rows.Close()
	ii := make([]Instrument, 0) // Instrument
	for rows.Next() {
		var i Instrument
		var p orb.Point
		err := rows.Scan(
			&i.ID, &i.Deleted, &i.StatusID, &i.Status, &i.StatusTime, &i.Slug, &i.Name, &i.TypeID, &i.Type, wkb.Scanner(&p), &i.Station, &i.StationOffset,
			&i.Creator, &i.CreateDate, &i.Updater, &i.UpdateDate, &i.ProjectID, pq.Array(&i.Constants), pq.Array(&i.Groups), pq.Array(&i.AlertConfigs),
			&i.Formula, &i.FormulaName, &i.NIDID, &i.USGSID,
		)
		if err != nil {
			return make([]Instrument, 0), err
		}
		// Set Geometry field
		i.Geometry = *geojson.NewGeometry(p)
		// Add
		ii = append(ii, i)
	}

	return ii, nil
}

// ListInstrumentsSQL is the base SQL to retrieve all instrumentsJSON
var listInstrumentsSQL = `SELECT id, deleted, status_id, status, status_time, slug,
	name, type_id, type, geometry, station, station_offset, creator, create_date,
	updater, update_date, project_id, constants, groups, alert_configs, formula, nid_id,
	usgs_id FROM v_instrument
	`
