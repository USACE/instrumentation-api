package model

import (
	"time"

	"github.com/google/uuid"
)

// AuditInfo holds common information about object creator and updater
type AuditInfo struct {
	Creator    uuid.UUID  `json:"creator"`
	CreateDate time.Time  `json:"create_date" db:"create_date"`
	Updater    *uuid.UUID `json:"updater"`
	UpdateDate *time.Time `json:"update_date" db:"update_date"`
}

// Action captures an API action, including user, user's roles,
// type (GET, POST, etc.), and Time
type Action struct {
	Actor int       `json:"actor"`
	Type  string    `json:"action"`
	Time  time.Time `json:"action_time"`
}
