package models

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	// pq library
	_ "github.com/lib/pq"
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/encoding/wkb"
	"github.com/paulmach/orb/encoding/wkt"
	"github.com/paulmach/orb/geojson"
)

// Instrument is an instrument
type Instrument struct {
	ID                uuid.UUID        `json:"id"`
	Groups            []uuid.UUID      `json:"groups"`
	StatusID          uuid.UUID        `json:"status_id" db:"status_id"`
	Status            string           `json:"status"`
	StatusTime        time.Time        `json:"status_time" db:"status_time"`
	Deleted           bool             `json:"-"`
	Slug              string           `json:"slug"`
	Name              string           `json:"name"`
	TypeID            uuid.UUID        `json:"type_id" db:"type_id"`
	Type              string           `json:"type"`
	Geometry          geojson.Geometry `json:"geometry,omitempty"`
	Station           *int             `json:"station"`
	StationOffset     *int             `json:"offset" db:"station_offset"`
	ProjectID         *uuid.UUID       `json:"project_id" db:"project_id"`
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

	rows, err := db.Queryx(listInstrumentsSQL() + " WHERE NOT I.deleted")
	if err != nil {
		return make([]Instrument, 0), err
	}
	return InstrumentsFactory(db, rows)
}

// GetInstrument returns a single instrument
func GetInstrument(db *sqlx.DB, id *uuid.UUID) (*Instrument, error) {

	rows, err := db.Queryx(listInstrumentsSQL()+" WHERE I.id = $1", id)
	if err != nil {
		return nil, err
	}
	ii, err := InstrumentsFactory(db, rows)
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
func CreateInstruments(db *sqlx.DB, a *Action, instruments []Instrument) error {

	txn, err := db.Begin()
	if err != nil {
		return err
	}

	// Instrument
	stmt1, err := txn.Prepare(
		`INSERT INTO instrument
			(id, slug, name, type_id, geometry, station, station_offset, creator, create_date, updater, update_date, project_id)
		 VALUES
		 	($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`,
	)
	if err != nil {
		return err
	}

	// Instrument Status
	stmt2, err := txn.Prepare(createInstrumentStatusSQL())
	if err != nil {
		return err
	}

	for _, i := range instruments {
		// Load Instrument
		if _, err := stmt1.Exec(
			i.ID, i.Slug, i.Name, i.TypeID, wkt.MarshalString(i.Geometry.Geometry()),
			i.Station, i.StationOffset, a.Actor, a.Time, a.Actor, a.Time, i.ProjectID,
		); err != nil {
			return err
		}
		if _, err := stmt2.Exec(i.ID, i.StatusID, i.StatusTime); err != nil {
			return err
		}
	}
	if err := stmt1.Close(); err != nil {
		return err
	}
	if err := stmt2.Close(); err != nil {
		return err
	}
	if err := txn.Commit(); err != nil {
		return err
	}

	return nil
}

// ValidateCreateInstruments creates many instruments from an array of instruments
func ValidateCreateInstruments(db *sqlx.DB, a *Action, instruments []Instrument) (CreateInstrumentsValidationResult, error) {

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
func UpdateInstrument(db *sqlx.DB, a *Action, i *Instrument) (*Instrument, error) {

	txn, err := db.Begin()
	if err != nil {
		return nil, err
	}

	// Instrument
	stmt1, err := txn.Prepare(
		`UPDATE instrument
		 SET    name = $2,
			    type_id = $3,
			    geometry = ST_GeomFromWKB($4),
			    updater = $5,
				update_date = $6,
				project_id = $7,
				station = $8,
				station_offset = $9
		 WHERE id = $1
		 RETURNING id`,
	)
	// Update Instrument
	var updatedID uuid.UUID
	if err := stmt1.QueryRow(
		i.ID, i.Name, i.TypeID, wkb.Value(i.Geometry.Geometry()),
		a.Actor, a.Time, i.ProjectID, i.Station, i.StationOffset,
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

// DeleteFlagInstrument changes delete flag to true
func DeleteFlagInstrument(db *sqlx.DB, id *uuid.UUID) error {

	if _, err := db.Exec(`UPDATE instrument SET deleted = true WHERE id = $1`, id); err != nil {
		return err
	}

	return nil
}

// InstrumentsFactory converts database rows to Instrument objects
func InstrumentsFactory(db *sqlx.DB, rows *sqlx.Rows) ([]Instrument, error) {
	defer rows.Close()
	_IDs := make([]uuid.UUID, 0)           // Instrument IDs (used to get groups)
	ii := make([]Instrument, 0) // Instrument
	for rows.Next() {
		var _ID uuid.UUID // InstrumentID
		var i Instrument
		var p orb.Point
		err := rows.Scan(
			&_ID, &i.Deleted, &i.StatusID, &i.Status, &i.StatusTime, &i.Slug, &i.Name, &i.TypeID, &i.Type, wkb.Scanner(&p), &i.Station, &i.StationOffset,
			&i.Creator, &i.CreateDate, &i.Updater, &i.UpdateDate, &i.ProjectID,
		)
		if err != nil {
			return make([]Instrument, 0), err
		}
		// Add _ID to list of IDs; Used to fetch instrument_group_instruments
		_IDs = append(_IDs, _ID)
		// Set InstrumentInfo ID Field
		i.ID = _ID
		// Set Geometry field
		i.Geometry = *geojson.NewGeometry(p)
		// Add
		ii = append(ii, i)
	}

	// Add groups
	groupMap, err := instrumentGroupMap(db, _IDs)
	if err != nil {
		return make([]Instrument, 0), err
	}

	// Assign Array of Group IDs to Each Instrument
	for idx := range ii {
		groups := groupMap[ii[idx].ID]
		if groups == nil {
			// Instrument is not member of any groups
			ii[idx].Groups = make([]uuid.UUID, 0)
			continue
		}
		ii[idx].Groups = groups
	}
	return ii, nil
}

// instrumentGroupMap takes a list of instrument IDs and returns a map of
// <instrument_id>: [<group_id>, <group_id>, <group_id>] for all instrumentIDs passed in
// this allows nesting instrument groups inside the instrument struct using only 2 database hits.
// (avoids nested SQL queries or numerous database queries.
func instrumentGroupMap(db *sqlx.DB, instrumentIDs []uuid.UUID) (map[uuid.UUID][]uuid.UUID, error) {
	sql := func(inClause string) string {
		return fmt.Sprintf(
			`SELECT a.instrument_id,
					a.instrument_group_id
			FROM instrument_group_instruments a
			INNER JOIN instrument i ON i.id = a.instrument_id
			INNER JOIN instrument_group g ON g.id = a.instrument_group_id
			WHERE NOT (i.deleted OR g.deleted)
			%s
			ORDER BY a.instrument_id
			`,
			inClause,
		)
	}

	var gg []struct {
		InstrumentID uuid.UUID `db:"instrument_id"`
		GroupID      uuid.UUID `db:"instrument_group_id"`
	}
	switch {
	// Fetch a subset of the instrument_group_instruments table
	case len(instrumentIDs) > 0:
		s := sql("AND a.instrument_id IN (?)")
		query, args, err := sqlx.In(s, instrumentIDs)
		if err != nil {
			return nil, err
		}
		err = db.Select(&gg, db.Rebind(query), args...)
		if err != nil {
			return nil, err
		}
	// Fetch entire instrument_group_instruments table (Listing all instruments)
	default:
		if err := db.Select(&gg, sql("")); err != nil {
			return nil, err
		}
	}

	// Sort Instrument Group IDs into Arrays by InstrumentID
	var _ID uuid.UUID                    // Working Instrument ID
	m := make(map[uuid.UUID][]uuid.UUID) // Instrument: GroupIDs Map
	for idx := range gg {
		if gg[idx].InstrumentID != _ID {
			// Started on a new instrument; Create empty array in map
			m[gg[idx].InstrumentID] = make([]uuid.UUID, 0)
		}
		// Add instrument_group_id to map
		m[gg[idx].InstrumentID] = append(m[gg[idx].InstrumentID], gg[idx].GroupID)
	}

	return m, nil
}

// ListInstrumentsSQL is the base SQL to retrieve all instrumentsJSON
func listInstrumentsSQL() string {
	return `SELECT I.id,
			       I.deleted,
				   S.status_id,
				   S.status,
				   S.status_time,
	               I.slug,
	               I.name,
	               I.type_id,
	               T.name                  AS type, 
				   ST_AsBinary(I.geometry) AS geometry,
				   I.station,
				   I.station_offset,
	               I.creator,
	               I.create_date,
	               I.updater,
	               I.update_date,
				   I.project_id
			FROM   instrument I
			INNER JOIN instrument_type T
			   ON T.id = I.type_id
			INNER JOIN (
				SELECT
                	DISTINCT ON (instrument_id) instrument_id, 
					a.time                 AS status_time,
					a.status_id            AS status_id,
					d.name                 AS status
				FROM instrument_status a
				INNER JOIN status d ON d.id = a.status_id
				WHERE a.time <= now()
				ORDER BY instrument_id, a.time DESC
			) S ON S.instrument_id = I.id
			`
}
