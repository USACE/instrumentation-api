package models

import (
	"api/root/dbutils"
	"database/sql"
	"log"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

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
}

// ListInstrumentGroups returns a list of instrument groups
func ListInstrumentGroups(db *sqlx.DB) []InstrumentGroup {
	sql := "SELECT id, slug, name, description FROM instrument_group"
	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()
	result := make([]InstrumentGroup, 0)
	for rows.Next() {
		g := InstrumentGroup{}
		err := rows.Scan(&g.ID, &g.Slug, &g.Name, &g.Description)
		if err != nil {
			panic(err)
		}
		result = append(result, g)
	}
	return result
}

// GetInstrumentGroup returns a single instrument group
func GetInstrumentGroup(db *sqlx.DB, ID uuid.UUID) InstrumentGroup {
	sql := "SELECT id, slug, name, description FROM instrument_group WHERE id = $1"

	var g InstrumentGroup
	err := db.QueryRow(sql, ID).Scan(
		&g.ID, &g.Slug, &g.Name, &g.Description,
	)
	if err != nil {
		log.Printf("Fail to query and scan row with ID %s; %s", ID, err)
	}
	return g
}

// CreateInstrumentGroup creates a single instrument group
func CreateInstrumentGroup(db *sqlx.DB, g *InstrumentGroup) (uuid.UUID, error) {

	// UUID
	id := uuid.Must(uuid.NewRandom())
	// unique slug
	slug, err := dbutils.NextUniqueSlug(db, g.Name, "instrument_group", "slug")
	if err != nil {
		return uuid.UUID{}, err
	}
	_, err = db.Exec(
		`INSERT INTO instrument_group (id, slug, name, description) VALUES ($1, $2, $3, $4)`,
		id, slug, g.Name, g.Description,
	)

	return id, err
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
	        	   instrument_type.NAME              AS instrument_type,
	               instrument.height,
	               ST_AsBinary(instrument.geometry) AS geometry
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
		err := rows.Scan(&n.ID, &n.Slug, &n.Name, &n.Type, &n.Height, wkb.Scanner(&p))
		n.Geometry = *geojson.NewGeometry(p)
		if err != nil {
			panic(err)
		}

		result = append(result, n)
	}
	return result
}

// CreateInstrumentGroupInstruments adds an instrument to an instrument group
func CreateInstrumentGroupInstruments(db *sqlx.DB, instrumentGroupID uuid.UUID, instrumentID uuid.UUID) (sql.Result, error) {
	result, err := db.Exec(
		`INSERT INTO instrument_group_instruments (instrument_group_id, instrument_id) VALUES ($1, $2)`,
		instrumentGroupID,
		instrumentID,
	)
	return result, err
}

// DeleteInstrumentGroupInstruments adds an instrument to an instrument group
func DeleteInstrumentGroupInstruments(db *sqlx.DB, instrumentGroupID uuid.UUID, instrumentID uuid.UUID) (sql.Result, error) {
	result, err := db.Exec(
		`DELETE FROM instrument_group_instruments WHERE instrument_group_id = $1 and instrument_id = $2`,
		instrumentGroupID,
		instrumentID,
	)
	return result, err
}
