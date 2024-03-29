package model

import (
	"time"

	"github.com/google/uuid"
)

// AuditInfo holds common information about object creator and updater
type AuditInfo struct {
	CreatorID       uuid.UUID  `json:"creator_id" db:"creator"`
	CreatorUsername *string    `json:"creator_username,omitempty" db:"creator_username"`
	CreateDate      time.Time  `json:"create_date" db:"create_date"`
	UpdaterID       *uuid.UUID `json:"updater_id" db:"updater"`
	UpdaterUsername *string    `json:"updater_username,omitempty" db:"updater_username"`
	UpdateDate      *time.Time `json:"update_date" db:"update_date"`
}

type IDSlug struct {
	ID   uuid.UUID `json:"id"`
	Slug string    `json:"slug"`
}

type IDName struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type IDSlugName struct {
	IDSlug
	Name string `json:"name,omitempty"`
}

type IDSlugCollection struct {
	Items []IDSlug `json:"items"`
}

// Shortener allows a shorter representation of an object. Typically, ID and Slug fields
type Shortener[T any] interface {
	Shorten() T
}
