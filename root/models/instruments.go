package models

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// Instrument is an instrument data structure
type Instrument struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type InstrumentGroup struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// GetInstruments returns an array of instruments from the database
func GetInstruments(db *sql.DB) []Instrument {
	sql := "SELECT id, name FROM instrument"
	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()
	result := make([]Instrument, 0)
	for rows.Next() {
		n := Instrument{}
		err := rows.Scan(&n.ID, &n.Name)
		if err != nil {
			panic(err)
		}
		result = append(result, n)
	}
	return result
}

// GetInstrumentGroups returns a list of instrument groups
func GetInstrumentGroups(db *sql.DB) []InstrumentGroup {
	sql := "SELECT id, name FROM instrument_groups"
	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()
	result := make([]InstrumentGroup, 0)
	for rows.Next() {
		n := InstrumentGroup{}
		err := rows.Scan(&n.ID, &n.Name)
		if err != nil {
			panic(err)
		}
		result = append(result, n)
	}
	return result
}
