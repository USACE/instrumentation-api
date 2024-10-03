package model

import "github.com/google/uuid"

type Job struct {
	ID      uuid.UUID
	Process string
	Payload JobPayload
}

type JobPayload ReportJobPayload

type ReportJobPayload struct {
	ReportConfigID uuid.UUID
}
