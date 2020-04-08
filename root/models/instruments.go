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

// Instrument is an instrument data structure
type Instrument struct {
	ID       uuid.UUID        `json:"id"`
	Name     string           `json:"name"`
	Type     string           `json:"type"`
	Height   string           `json:"height"`
	Geometry geojson.Geometry `json:"geometry"`
}

// ListInstruments returns an array of instruments from the database
func ListInstruments(db *sqlx.DB) []Instrument {
	sql := `SELECT instrument.id, 
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
		var n Instrument
		err := rows.Scan(&n.ID, &n.Name, &n.Type, &n.Height, wkb.Scanner(&p))
		n.Geometry = *geojson.NewGeometry(p)

		if err != nil {
			panic(err)
		}

		result = append(result, n)
	}
	return result
}

// GetInstrument returns a single instrument
func GetInstrument(db *sqlx.DB, id uuid.UUID) Instrument {
	sql := `SELECT instrument.id,
	        	   instrument.NAME,
	        	   instrument_type.NAME              AS instrument_type,
	               instrument.height,
	               ST_AsBinary(instrument.geometry) AS geometry
            FROM   instrument
	               INNER JOIN instrument_type
							  ON instrument_type.id = instrument.instrument_type_id
			WHERE instrument.id = $1
			`
	var result Instrument
	var p orb.Point
	err := db.QueryRow(sql, id).Scan(&result.ID, &result.Name, &result.Type, &result.Height, wkb.Scanner(&p))
	result.Geometry = *geojson.NewGeometry(p)
	if err != nil {
		log.Printf("Fail to query and scan row with ID %s; %s", id, err)
	}
	return result
}

// CreateInstrument creates a single instrument
func CreateInstrument(db *sqlx.DB, i *Instrument) error {

	if _, err := db.Exec(
		`INSERT INTO instrument (id, name, height, instrument_type_id, geometry) VALUES ($1, $2, $3, $4, $5)`,
		i.ID, i.Name, i.Height, i.Type, wkb.Value(i.Geometry.Geometry()),
	); err != nil {
		return err
	}

	return nil
}

// UpdateInstrument updates a single instrument
func UpdateInstrument(db *sqlx.DB, i *Instrument) error {

	_, err := db.Exec(
		`UPDATE instrument SET name = $1, height = $2, instrument_type_id = $3, geometry = $4 WHERE id = $5`,
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
