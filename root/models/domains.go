package models

import (
	"github.com/jmoiron/sqlx"
	// pq library
	_ "github.com/lib/pq"
)

// Domain is a struct for returning all database domain values
type Domain struct {
	ID          string  `json:"id"`
	Group       string  `json:"group"`
	Value       string  `json:"value"`
	Description *string `json:"description"`
}

// GetDomains returns a UNION of all domain tables in the database
func GetDomains(db *sqlx.DB) ([]Domain, error) {

	dd := make([]Domain, 0)

	sql := `SELECT id, 
	               'instrument_type' AS group, 
				   name              AS value,
				   null              AS description
            FROM   instrument_type 
            UNION 
            SELECT id, 
            	   'parameter' AS group, 
				   name        AS value,
				   null        AS description

            FROM   parameter 
            UNION 
            SELECT id, 
            	   'unit' AS group, 
				   name   AS value,
				   null   AS description
			FROM   unit
			UNION
			SELECT id,
				   'status' AS group,
				   name                AS value,
				   description
			FROM   status
			UNION
			SELECT id,
				   'zreference_datum' AS group,
				   name               AS value,
				   null               AS description
			FROM zreference_datum
	`
	if err := db.Select(&dd, sql); err != nil {
		return make([]Domain, 0), err
	}
	return dd, nil
}
