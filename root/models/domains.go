package models

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Domain struct {
	ID    string `json:"id"`
	Group string `json:"group"`
	Value string `json:"value"`
}

// GetDomains returns a UNION of all domain tables in the database
func GetDomains(db *sql.DB) []Domain {
	sql := `SELECT id, 
	               'instrument_type' AS group, 
	               name              AS value 
            FROM   instrument_type 
            UNION 
            SELECT id, 
            	   'parameter' AS group, 
            	   name        AS value 
            FROM   parameter 
            UNION 
            SELECT id, 
            	   'unit' AS group, 
            	   name   AS value 
            FROM   unit 
	`
	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()
	result := make([]Domain, 0)
	for rows.Next() {
		d := Domain{}
		err := rows.Scan(&d.ID, &d.Group, &d.Value)
		if err != nil {
			log.Panic(err)
		}
		result = append(result, d)
	}
	return result
}
