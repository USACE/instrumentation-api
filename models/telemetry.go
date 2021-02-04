package models

import (
	"github.com/google/uuid"
)

type Telemetry struct {
	ID       uuid.UUID
	TypeID   string
	TypeSlug string
	TypeName string
}
