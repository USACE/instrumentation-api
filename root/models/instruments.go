package models

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	geojson "github.com/paulmach/go.geojson"
)

// Instrument is an instrument data structure
type Instrument struct {
	ID       string           `json:"id"`
	Name     string           `json:"name"`
	Type     string           `json:"type"`
	Height   string           `json:"height"`
	Geometry geojson.Geometry `json:"geometry"`
}

// GetInstruments returns an array of instruments from the database
func GetInstruments(db *sql.DB) []Instrument {
	sql := `SELECT instrument.id, 
	        	   instrument.NAME,
	        	   instrument_type.NAME              AS instrument_type, 
	               instrument.height, 
	               ST_AsGeoJSON(instrument.geometry) AS geometry 
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
		n := Instrument{}
		var geometry *geojson.Geometry
		err := rows.Scan(&n.ID, &n.Name, &n.Type, &n.Height, &geometry)
		if err != nil {
			panic(err)
		}
		n.Geometry = *geometry

		result = append(result, n)
	}
	return result
}

// GetInstrument returns a single instrument
func GetInstrument(db *sql.DB, ID string) Instrument {
	sql := `SELECT instrument.id, 
	        	   instrument.NAME,
	        	   instrument_type.NAME              AS instrument_type, 
	               instrument.height, 
	               ST_AsGeoJSON(instrument.geometry) AS geometry 
            FROM   instrument
	               INNER JOIN instrument_type
							  ON instrument_type.id = instrument.instrument_type_id
			WHERE instrument.id = $1
			`
	var result Instrument
	var geom *geojson.Geometry
	err := db.QueryRow(sql, ID).Scan(
		&result.ID, &result.Name, &result.Type, &result.Height, &geom,
	)
	result.Geometry = *geom
	if err != nil {
		log.Printf("Fail to query and scan row with ID %s; %s", ID, err)
	}
	return result
}
