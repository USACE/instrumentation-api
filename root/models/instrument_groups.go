package models

import (
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
