package models

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	geojson "github.com/paulmach/go.geojson"
)

// InstrumentGroup holds information for entity instrument_group
type InstrumentGroup struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// GetInstrumentGroups returns a list of instrument groups
func GetInstrumentGroups(db *sql.DB) []InstrumentGroup {
	sql := "SELECT id, name, description FROM instrument_group"
	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()
	result := make([]InstrumentGroup, 0)
	for rows.Next() {
		n := InstrumentGroup{}
		err := rows.Scan(&n.ID, &n.Name, &n.Description)
		if err != nil {
			panic(err)
		}
		result = append(result, n)
	}
	return result
}

// GetInstrumentGroup returns a single instrument group
func GetInstrumentGroup(db *sql.DB, ID string) InstrumentGroup {
	sql := "SELECT id, name, description FROM instrument_group WHERE id = $1"

	var result InstrumentGroup
	err := db.QueryRow(sql, ID).Scan(
		&result.ID, &result.Name, &result.Description,
	)
	if err != nil {
		log.Printf("Fail to query and scan row with ID %s; %s", ID, err)
	}
	return result
}

// GetInstrumentGroupInstruments returns a list of instrument group instruments for a given instrument
func GetInstrumentGroupInstruments(db *sql.DB, ID string) []Instrument {

	sql := `SELECT A.instrument_id, 
	        	   instrument.NAME,
	        	   instrument_type.NAME              AS instrument_type, 
	               instrument.height, 
	               ST_AsGeoJSON(instrument.geometry)::json AS geometry 
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
