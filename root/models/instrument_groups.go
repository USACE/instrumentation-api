package models

import (
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"github.com/lib/pq"

	"github.com/paulmach/orb"
	"github.com/paulmach/orb/encoding/wkb"
	"github.com/paulmach/orb/geojson"
)

// InstrumentGroup holds information for entity instrument_group
type InstrumentGroup struct {
	ID          uuid.UUID  `json:"id"`
	Deleted     bool       `json:"-"`
	Slug        string     `json:"slug"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Creator     int        `json:"creator"`
	CreateDate  time.Time  `json:"create_date" db:"create_date"`
	Updater     int        `json:"updater"`
	UpdateDate  time.Time  `json:"update_date" db:"update_date"`
	ProjectID   *uuid.UUID `json:"project_id" db:"project_id"`
}

// ListInstrumentGroupSlugs lists used instrument group slugs in the database
func ListInstrumentGroupSlugs(db *sqlx.DB) ([]string, error) {

	ss := make([]string, 0)
	if err := db.Select(&ss, "SELECT slug FROM instrument_group"); err != nil {
		return make([]string, 0), err
	}
	return ss, nil
}

// ListInstrumentGroups returns a list of instrument groups
func ListInstrumentGroups(db *sqlx.DB) ([]InstrumentGroup, error) {

	gg := make([]InstrumentGroup, 0)
	if err := db.Select(
		&gg,
		`SELECT id, slug, name, description, creator, create_date, updater, update_date, project_id
		 FROM   instrument_group
		 WHERE NOT deleted
		`,
	); err != nil {
		return make([]InstrumentGroup, 0), err
	}

	return gg, nil
}

// GetInstrumentGroup returns a single instrument group
func GetInstrumentGroup(db *sqlx.DB, ID uuid.UUID) (*InstrumentGroup, error) {
	sql := "SELECT id, slug, name, description, creator, create_date, updater, update_date, project_id FROM instrument_group WHERE id = $1"

	var g InstrumentGroup
	if err := db.QueryRowx(sql, ID).StructScan(&g); err != nil {
		return nil, err
	}

	return &g, nil
}

// CreateInstrumentGroupBulk creates many instruments from an array of instruments
func CreateInstrumentGroupBulk(db *sqlx.DB, groups []InstrumentGroup) error {

	txn, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := txn.Prepare(pq.CopyIn(
		"instrument_group",
		"id", "slug", "name", "description", "creator", "create_date", "updater", "update_date", "project_id",
	))

	if err != nil {
		log.Fatal(err)
	}

	for _, g := range groups {

		_, err := stmt.Exec(
			g.ID, g.Slug, g.Name, g.Description, g.Creator, g.CreateDate, g.Updater, g.UpdateDate, g.ProjectID,
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

// DeleteFlagInstrumentGroup sets the deleted field to true
func DeleteFlagInstrumentGroup(db *sqlx.DB, ID uuid.UUID) error {
	if _, err := db.Exec(`UPDATE instrument_group SET deleted = true WHERE id = $1`, ID); err != nil {
		return err
	}
	return nil
}

// ListInstrumentGroupInstruments returns a list of instrument group instruments for a given instrument
func ListInstrumentGroupInstruments(db *sqlx.DB, ID uuid.UUID) []Instrument {

	sql := `SELECT A.instrument_id,
	               instrument.active,
	               instrument.slug,
				   instrument.NAME,
				   instrument.INSTRUMENT_TYPE_ID,
	        	   instrument_type.NAME              AS instrument_type,
	               instrument.height,
				   ST_AsBinary(instrument.geometry) AS geometry,
				   instrument.creator,
				   instrument.create_date,
				   instrument.updater,
				   instrument.update_date,
				   instrument.project_id
            FROM   instrument_group_instruments A
	               INNER JOIN instrument instrument
	               		   ON instrument.id = A.instrument_id
	               INNER JOIN instrument_type
	               		   ON instrument_type.id = instrument.instrument_type_id
			WHERE  instrument_group_id = $1
			`

	rows, err := db.Query(sql, ID)

	if err != nil {
		panic(err)
	}

	defer rows.Close()
	result := make([]Instrument, 0)
	for rows.Next() {
		var p orb.Point
		var n Instrument
		err := rows.Scan(
			&n.ID, &n.Active, &n.Slug, &n.Name, &n.TypeID, &n.Type, &n.Height, wkb.Scanner(&p),
			&n.Creator, &n.CreateDate, &n.Updater, &n.UpdateDate, &n.ProjectID,
		)
		n.Geometry = *geojson.NewGeometry(p)
		if err != nil {
			panic(err)
		}

		result = append(result, n)
	}
	return result
}

// CreateInstrumentGroupInstruments adds an instrument to an instrument group
func CreateInstrumentGroupInstruments(db *sqlx.DB, instrumentGroupID uuid.UUID, instrumentID uuid.UUID) error {

	if _, err := db.Exec(
		`INSERT INTO instrument_group_instruments (instrument_group_id, instrument_id) VALUES ($1, $2)`,
		instrumentGroupID,
		instrumentID,
	); err != nil {
		return err
	}

	return nil
}

// DeleteInstrumentGroupInstruments adds an instrument to an instrument group
func DeleteInstrumentGroupInstruments(db *sqlx.DB, instrumentGroupID uuid.UUID, instrumentID uuid.UUID) error {

	if _, err := db.Exec(
		`DELETE FROM instrument_group_instruments WHERE instrument_group_id = $1 and instrument_id = $2`,
		instrumentGroupID, instrumentID,
	); err != nil {
		return err
	}

	return nil
}
