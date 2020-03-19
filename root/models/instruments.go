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

// GetInstruments returns an array of instruments from the database
func GetInstruments(db *sql.DB) []Instrument {
	sql := "SELECT id, name FROM instrument"
	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()
	var result []Instrument
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
