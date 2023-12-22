package model

import (
	"context"

	"github.com/google/uuid"
)

// Domain is a struct for returning all database domain values
type Domain struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Group       string    `json:"group" db:"group"`
	Value       string    `json:"value" db:"value"`
	Description *string   `json:"description" db:"description"`
}

type DomainGroup struct {
	Group string                         `json:"group" db:"group"`
	Opts  dbJSONSlice[DomainGroupOption] `json:"opts" db:"opts"`
}

type DomainGroupOption struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Value       string    `json:"value" db:"value"`
	Description *string   `json:"description" db:"description"`
}

type DomainGroupCollection []DomainGroup

type DomainMap map[string][]DomainGroupOption

const getDomains = `
	SELECT * FROM v_domain
`

// GetDomains returns a UNION of all domain tables in the database
func (q *Queries) GetDomains(ctx context.Context) ([]Domain, error) {
	dd := make([]Domain, 0)
	if err := q.db.SelectContext(ctx, &dd, getDomains); err != nil {
		return nil, err
	}
	return dd, nil
}

const getDomainMap = `
	SELECT * FROM v_domain_group
`

// GetDomainsV2 returns all domains grouped by table
func (q *Queries) GetDomainMap(ctx context.Context) (DomainMap, error) {
	dd := make([]DomainGroup, 0)
	if err := q.db.SelectContext(ctx, &dd, getDomainMap); err != nil {
		return nil, err
	}
	m := make(DomainMap)
	for i := range dd {
		m[dd[i].Group] = dd[i].Opts
	}
	return m, nil
}
