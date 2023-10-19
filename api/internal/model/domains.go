package model

import (
	"context"

	"github.com/google/uuid"
)

// Domain is a struct for returning all database domain values
type Domain struct {
	ID          uuid.UUID `json:"id"`
	Group       string    `json:"group"`
	Value       string    `json:"value"`
	Description *string   `json:"description"`
}

const getDomains = `
	SELECT
		id, 
		'instrument_type'		AS group, 
		name				AS value,
		null				AS description
	FROM instrument_type 
	UNION 
	SELECT
		id, 
		'parameter'			AS group, 
		name       			AS value,
		null       			AS description
	FROM parameter 
	UNION 
	SELECT
		id, 
		'unit'				AS group, 
		name  				AS value,
		null  				AS description
	FROM unit
	UNION
	SELECT
		id,
		'status'			AS group,
		name				AS value,
		description			AS description
	FROM status
	UNION
	SELECT
		id,
		'role' 				AS group,
		name   				AS value,
		null   				AS description
	FROM role
	UNION
	SELECT
		id,
		'datalogger_model'		AS group,
		model			  	AS value,
		null 			  	AS description
	FROM datalogger_model
	UNION
	SELECT
		id,
		'submittal_status'		AS group,
		name				AS value,
		null				AS description
	FROM submittal_status
	UNION
	SELECT
		id,
		'alert_type'			AS group,
		name				AS value,
		null				AS description
	FROM alert_type
	ORDER BY "group", value
`

// GetDomains returns a UNION of all domain tables in the database
func (q *Queries) GetDomains(ctx context.Context) ([]Domain, error) {
	dd := make([]Domain, 0)
	if err := q.db.SelectContext(ctx, &dd, getDomains); err != nil {
		return nil, err
	}
	return dd, nil
}
