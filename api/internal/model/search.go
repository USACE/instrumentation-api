package model

import (
	"github.com/google/uuid"
)

// EmailAutocompleteResult stores search result in profiles and emails
type SearchResult struct {
	ID   uuid.UUID   `json:"id"`
	Type string      `json:"type"`
	Item interface{} `json:"item,omitempty"`
}
