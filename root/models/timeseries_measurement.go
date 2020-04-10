package models

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// TimeseriesMeasurement is an time series data structure
type TimeseriesMeasurement struct {
	ID       	string      `json:"id"`
	Time     	string      `json:"time"`
	Value  		string		`json:"value"`
	Timeseries 	string		`json:"timeseries"`
}

// GetTimeseriesMeasurements returns all time series measurements from the database
func GetTimeseriesMeasurements(db *sqlx.DB) []TimeseriesMeasurement {
	sql := `SELECT  timeseries_measurement.id, 
				    timeseries_measurement.time,
				    timeseries_measurement.value,
				    timeseries.id as timeseries
			FROM    timeseries_measurement
				    INNER JOIN timeseries
							  	ON timeseries.id = timeseries_measurement.timeseries_id
					
			`
	
	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()
	result := make([]TimeseriesMeasurement, 0)
	for rows.Next() {
		n := TimeseriesMeasurement{}
		err := rows.Scan(&n.ID, &n.Time, &n.Value, &n.Timeseries)
		if err != nil {
			panic(err)
		}

		result = append(result, n)
	}
	return result
}
