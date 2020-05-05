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
	ID          uuid.UUID `json:"id"`
	Slug        string    `json:"slug"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Creator     int       `json:"creator"`
	CreateDate  time.Time `json:"create_date" db:"create_date"`
	Updater     int       `json:"updater"`
	UpdateDate  time.Time `json:"update_date" db:"update_date"`
}

// ListInstrumentGroupSlugs lists used instrument group slugs in the database
func ListInstrumentGroupSlugs(db *sqlx.DB) []string {

	rows, err := db.Query(`SELECT slug from instrument_group`)

	if err != nil {
		log.Panic(err)
	}

	defer rows.Close()
	result := make([]string, 0)
	for rows.Next() {
		var slug string
		err := rows.Scan(&slug)
		if err != nil {
			log.Panic(err)
		}
		result = append(result, slug)
	}
	return result
}

// ListInstrumentGroups returns a list of instrument groups
func ListInstrumentGroups(db *sqlx.DB) []InstrumentGroup {

	sql := "SELECT id, slug, name, description, creator, create_date, updater, update_date FROM instrument_group"
	rows, err := db.Queryx(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()
	result := make([]InstrumentGroup, 0)
	for rows.Next() {
		g := InstrumentGroup{}
		err = rows.StructScan(&g)
		if err != nil {
			panic(err)
		}
		result = append(result, g)
	}

	return result
}

// GetInstrumentGroup returns a single instrument group
func GetInstrumentGroup(db *sqlx.DB, ID uuid.UUID) InstrumentGroup {
	sql := "SELECT id, slug, name, description, creator, create_date, updater, update_date FROM instrument_group WHERE id = $1"

	var g InstrumentGroup
	err := db.QueryRowx(sql, ID).StructScan(&g)
	if err != nil {
		log.Printf("Fail to query and scan row with ID %s; %s", ID, err)
	}
	return g
}

// CreateInstrumentGroup creates a single instrument group
func CreateInstrumentGroup(db *sqlx.DB, g *InstrumentGroup) error {

	_, err := db.NamedExec(
		`INSERT INTO instrument_group (id, slug, name, description, creator, create_date, updater, update_date)
		VALUES (:id, :slug, :name, :description, :creator, :create_date, :updater, :update_date)`,
		g,
	)

	return err
}

// CreateInstrumentGroupBulk creates many instruments from an array of instruments
func CreateInstrumentGroupBulk(db *sqlx.DB, groups []InstrumentGroup) error {

	txn, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := txn.Prepare(pq.CopyIn(
		"instrument_group",
		"id", "slug", "name", "description", "creator", "create_date", "updater", "update_date",
	))

	if err != nil {
		log.Fatal(err)
	}

	for _, g := range groups {

		_, err := stmt.Exec(
			g.ID, g.Slug, g.Name, g.Description, g.Creator, g.CreateDate, g.Updater, g.UpdateDate,
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

// DeleteInstrumentGroup deletes an instrument group and any associations in instrument_group_instruments
func DeleteInstrumentGroup(db *sqlx.DB, id uuid.UUID) error {

	tx, err := db.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}
	// delete instrument_group_instruments first to avoid foreign key constraint
	if _, err = tx.Exec(
		`DELETE FROM instrument_group_instruments WHERE instrument_group_id = $1`,
		id,
	); err != nil {
		tx.Rollback()
		return err
	}
	if _, err = tx.Exec(`DELETE FROM instrument_group WHERE id = $1`, id); err != nil {
		tx.Rollback()
		return err
	}
	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}

// ListInstrumentGroupInstruments returns a list of instrument group instruments for a given instrument
func ListInstrumentGroupInstruments(db *sqlx.DB, ID uuid.UUID) []Instrument {

	sql := `SELECT A.instrument_id,
	               instrument.slug,
				   instrument.NAME,
				   instrument.INSTRUMENT_TYPE_ID,
	        	   instrument_type.NAME              AS instrument_type,
	               instrument.height,
				   ST_AsBinary(instrument.geometry) AS geometry,
				   instrument.creator,
				   instrument.create_date,
				   instrument.updater,
				   instrument.update_date
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
			&n.ID, &n.Slug, &n.Name, &n.TypeID, &n.Type, &n.Height, wkb.Scanner(&p),
			&n.Creator, &n.CreateDate, &n.Updater, &n.UpdateDate,
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
