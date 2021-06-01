package models

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// EmailAutocompleteResult stores search result in profiles and emails
type SearchResult struct {
	ID   uuid.UUID   `json:"id"`
	Type string      `json:"type"`
	Item interface{} `json:"item,omitempty"`
}

// ProjectSearch returns search result for projects
func ProjectSearch(db *sqlx.DB, str *string, limit *int) ([]SearchResult, error) {

	rows, err := db.Queryx(
		listProjectsSQL+` WHERE NOT deleted AND name ILIKE '%' || $1 || '%' LIMIT $2 ORDER BY name`, str, limit,
	)
	if err != nil {
		return make([]SearchResult, 0), err
	}
	projects, err := ProjectFactory(rows)
	if err != nil {
		return make([]SearchResult, 0), err
	}

	rr := make([]SearchResult, len(projects))
	for idx, p := range projects {
		rr[idx] = SearchResult{ID: p.ID, Type: "project", Item: p}
	}

	return rr, nil
}
