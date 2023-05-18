package models

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// Domain is a struct for returning all database domain values
type Domain struct {
	ID          uuid.UUID `json:"id"`
	Group       string    `json:"group"`
	Value       string    `json:"value"`
	Description *string   `json:"description"`
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
			       'role' AS group,
				   name   AS value,
				   null   AS description
			FROM   role
			UNION
			SELECT id,
				   'datalogger_model' AS group,
				   model			  AS value,
				   null 			  AS description
			FROM   datalogger_model
			UNION
			SELECT id,
				   'alert_status'	AS group,
				   name				AS value,
				   null				AS description
			FROM   alert_status
			order by "group", value
			UNION
			SELECT id,
				   'alert_type'		AS group,
				   name				AS value,
				   null				AS description
			FROM   alert_type
			order by "group", value
	`
	if err := db.Select(&dd, sql); err != nil {
		return make([]Domain, 0), err
	}
	return dd, nil
}
