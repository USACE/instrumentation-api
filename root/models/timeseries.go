package models

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Instrument is an instrument data structure
type Timeseries struct {
	ID       	string      `json:"id"`
	Name     	string      `json:"name"`
	Instrument  string		`json:"instrument`
	Parameter 	string		`json:"parameter`
	Unit		string		`json:"unit`
}

// GetInstruments returns an array of instruments from the database
func GetTimeseries(db *sqlx.DB) []Timeseries {
	sql := `SELECT  timeseries.id, 
				    timeseries.NAME,
				    instrument.Name as instrument,
				    parameter.Name as parameter,
				    unit.Name as unit
			FROM    timeseries
				    INNER JOIN instrument
							  	ON instrument.id = timeseries.instrument_id
				    INNER JOIN parameter
								ON parameter.id = timeseries.parameter_id
					INNER JOIN unit
								ON unit.id = timeseries.unit_id
					
			`
	
	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()
	result := make([]Timeseries, 0)
	for rows.Next() {
		n := Timeseries{}
		err := rows.Scan(&n.ID, &n.Name, &n.Instrument, &n.Parameter, &n.Unit)
		if err != nil {
			panic(err)
		}

		result = append(result, n)
	}
	return result
}

