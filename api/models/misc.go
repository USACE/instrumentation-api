package models

import (
	"github.com/google/uuid"
)

// IDAndSlug is a UUID4 and Slug representation of something
type IDAndSlug struct {
	ID   uuid.UUID `json:"id"`
	Slug string    `json:"slug"`
}

// IDAndSlugCollection is a collection of objects with ID and Slug properties
type IDAndSlugCollection struct {
	Items []IDAndSlug `json:"items"`
}

// Shortener allows a shorter representation of an object. Typically, ID and Slug fields only
type Shortener interface {
	shorten()
}
