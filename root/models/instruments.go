package models

import (
	"encoding/json"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"

	"github.com/paulmach/orb"
	"github.com/paulmach/orb/encoding/wkb"
	"github.com/paulmach/orb/encoding/wkt"
	"github.com/paulmach/orb/geojson"
)

// Instrument is an instrument data structure
type Instrument struct {
	ID            uuid.UUID        `json:"id"`
	Active        bool             `json:"active"`
	Deleted       bool             `json:"-"`
	Slug          string           `json:"slug"`
	Name          string           `json:"name"`
	TypeID        string           `json:"type_id"`
	Type          string           `json:"type"`
	Height        float32          `json:"height"`
	Geometry      geojson.Geometry `json:"geometry,omitempty"`
	Station       *int             `json:"station"`
	StationOffset *int             `json:"station_offset" db:"station_offset"`
	Creator       int              `json:"creator"`
	CreateDate    time.Time        `json:"create_date" db:"create_date"`
	Updater       int              `json:"updater"`
	UpdateDate    time.Time        `json:"update_date" db:"update_date"`
	ProjectID     *string          `json:"project_id" db:"project_id"`
}

// InstrumentCollection is a collection of Instrument items
type InstrumentCollection struct {
	Items []Instrument
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

	rows, err := db.Queryx(listInstrumentsSQL() + " WHERE NOT instrument.deleted")
	if err != nil {
		return make([]Instrument, 0), err
	}
	return InstrumentsFactory(rows)
}

// GetInstrument returns a single instrument
func GetInstrument(db *sqlx.DB, id uuid.UUID) (*Instrument, error) {

	rows, err := db.Queryx(listInstrumentsSQL()+" WHERE instrument.id = $1", id)
	if err != nil {
		return nil, err
	}
	ii, err := InstrumentsFactory(rows)
	if err != nil {
		return nil, err
	}
	return &ii[0], nil
}

// CreateInstrumentBulk creates many instruments from an array of instruments
func CreateInstrumentBulk(db *sqlx.DB, instruments []Instrument) error {

	txn, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := txn.Prepare(pq.CopyIn(
		"instrument",
		"id", "active", "slug", "name", "height", "instrument_type_id",
		"geometry", "station", "station_offset", "creator", "create_date", "updater", "update_date", "project_id",
	))

	if err != nil {
		log.Fatal(err)
	}

	for _, i := range instruments {

		_, err = stmt.Exec(
			i.ID, i.Active, i.Slug, i.Name, i.Height, i.TypeID, wkt.MarshalString(i.Geometry.Geometry()),
			i.Station, i.StationOffset, i.Creator, i.CreateDate, i.Updater, i.UpdateDate, i.ProjectID,
		)

		if err != nil {
			return err
		}
	}

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	err = stmt.Close()
	if err != nil {
		return err
	}

	err = txn.Commit()
	if err != nil {
		return err
	}

	return nil
}

// UpdateInstrument updates a single instrument
func UpdateInstrument(db *sqlx.DB, i *Instrument) (*Instrument, error) {

	var updatedID struct{ ID uuid.UUID }
	if err := db.QueryRowx(
		`UPDATE instrument
		 SET    name = $2,
			    active = $3,
			    height = $4,
			    instrument_type_id = $5,
			    geometry = ST_GeomFromWKB($6),
			    updater = $7,
				update_date = $8,
				project_id = $9,
				station = $10,
				station_offset = $11
		 WHERE id = $1
		 RETURNING id
		`, i.ID, i.Name, i.Active, i.Height, i.TypeID, wkb.Value(i.Geometry.Geometry()), i.Updater, i.UpdateDate, i.ProjectID, i.Station, i.StationOffset,
	).StructScan(&updatedID); err != nil {
		return nil, err
	}
	// Get Updated Row
	return GetInstrument(db, updatedID.ID)
}

// DeleteFlagInstrument changes delete flag to true
func DeleteFlagInstrument(db *sqlx.DB, id uuid.UUID) error {

	if _, err := db.Exec(`UPDATE instrument SET deleted = true WHERE id = $1`, id); err != nil {
		return err
	}

	return nil
}

// InstrumentsFactory returns a slice of instruments from a string of SQL
func InstrumentsFactory(rows *sqlx.Rows) ([]Instrument, error) {

	defer rows.Close()
	ii := make([]Instrument, 0)
	for rows.Next() {
		var i Instrument
		var p orb.Point
		err := rows.Scan(
			&i.ID, &i.Active, &i.Slug, &i.Name, &i.TypeID, &i.Type, &i.Height, wkb.Scanner(&p), &i.Station, &i.StationOffset,
			&i.Creator, &i.CreateDate, &i.Updater, &i.UpdateDate, &i.ProjectID,
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

// ListInstrumentsSQL is the base SQL to retrieve all instruments
func listInstrumentsSQL() string {
	return `SELECT instrument.id,
	               instrument.active,
	               instrument.slug,
	               instrument.NAME,
	               instrument.INSTRUMENT_TYPE_ID,
	               instrument_type.NAME              AS instrument_type, 
	               instrument.height, 
				   ST_AsBinary(instrument.geometry) AS geometry,
				   instrument.station,
				   instrument.station_offset,
	               instrument.creator,
	               instrument.create_date,
	               instrument.updater,
	               instrument.update_date,
	               instrument.project_id
			FROM   instrument
			INNER JOIN instrument_type
			   ON instrument_type.id = instrument.instrument_type_id
			`
}
