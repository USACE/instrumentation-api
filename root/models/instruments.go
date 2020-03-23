package models

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// Instrument is an instrument data structure
type Instrument struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// InstrumentGroup holds information for entity instrument_group
type InstrumentGroup struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
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
	sql := "SELECT id, name, description FROM instrument_group WHERE id = ?"

	var result InstrumentGroup
	err := db.QueryRow(sql, 1).Scan(
		&result.ID, &result.Name, &result.Description,
	)
	if err != nil {
		log.Fatalf("Fail to query and scan row with ID %s", ID)
	}
	return result
}
