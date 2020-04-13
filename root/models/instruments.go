package models

import (
	"api/root/dbutils"
	"log"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/paulmach/orb"
	"github.com/paulmach/orb/encoding/wkb"
	"github.com/paulmach/orb/geojson"
)

// Instrument is an instrument data structure
type Instrument struct {
	ID       uuid.UUID        `json:"id"`
	Slug     string           `json:"slug"`
	Name     string           `json:"name"`
	Type     string           `json:"type"`
	Height   string           `json:"height"`
	Geometry geojson.Geometry `json:"geometry"`
}

// ListInstruments returns an array of instruments from the database
func ListInstruments(db *sqlx.DB) []Instrument {
	sql := `SELECT instrument.id,
	               instrument.slug,
	        	   instrument.NAME,
	        	   instrument_type.NAME              AS instrument_type, 
	               instrument.height, 
	               ST_AsBinary(instrument.geometry) AS geometry 
            FROM   instrument
	               INNER JOIN instrument_type
	               		   ON instrument_type.id = instrument.instrument_type_id
			`

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()
	result := make([]Instrument, 0)
	for rows.Next() {
		var p orb.Point
		var i Instrument
		err := rows.Scan(&i.ID, &i.Slug, &i.Name, &i.Type, &i.Height, wkb.Scanner(&p))
		i.Geometry = *geojson.NewGeometry(p)

		if err != nil {
			panic(err)
		}

		result = append(result, i)
	}
	return result
}

// GetInstrument returns a single instrument
func GetInstrument(db *sqlx.DB, id uuid.UUID) Instrument {
	sql := `SELECT instrument.id,
	               instrument.slug,
	        	   instrument.NAME,
	        	   instrument_type.NAME              AS instrument_type,
	               instrument.height,
	               ST_AsBinary(instrument.geometry) AS geometry
            FROM   instrument
	               INNER JOIN instrument_type
							  ON instrument_type.id = instrument.instrument_type_id
			WHERE instrument.id = $1
			`

	var i Instrument
	var p orb.Point
	err := db.QueryRow(sql, id).Scan(&i.ID, &i.Slug, &i.Name, &i.Type, &i.Height, wkb.Scanner(&p))
	if err != nil {
		log.Printf("Fail to query and scan row with ID %s; %s", id, err)
	}
	i.Geometry = *geojson.NewGeometry(p)

	return i
}

// CreateInstrument creates a single instrument
func CreateInstrument(db *sqlx.DB, i *Instrument) error {

	// unique slug
	slug, err := dbutils.NextUniqueSlug(db, i.Name, "instrument", "slug")
	if err != nil {
		return err
	}
	if _, err := db.Exec(
		`INSERT INTO instrument (id, slug, name, height, instrument_type_id, geometry) VALUES ($1, $2, $3, $4, $5, $6)`,
		i.ID, slug, i.Name, i.Height, i.Type, wkb.Value(i.Geometry.Geometry()),
	); err != nil {
		return err
	}

	return nil
}

// UpdateInstrument updates a single instrument
func UpdateInstrument(db *sqlx.DB, i *Instrument) error {

	_, err := db.Exec(
		`UPDATE instrument SET name = $1, height = $2, instrument_type_id = $3, geometry = ST_GeomFromWKB($4) WHERE id = $5`,
		i.Name, i.Height, i.Type, wkb.Value(i.Geometry.Geometry()), i.ID,
	)

	if err != nil {
		return err
	}

	return nil

}

// DeleteInstrument deletes a single instrument
func DeleteInstrument(db *sqlx.DB, id uuid.UUID) error {
	_, err := db.Exec(`DELETE FROM instrument WHERE id = $1`, id)

	if err != nil {
		return err
	}

	return nil
}
