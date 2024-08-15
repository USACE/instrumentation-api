package model

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type DistrictRollup struct {
	AlertTypeID             uuid.UUID  `json:"alert_type_id" db:"alert_type_id"`
	OfficeID                *uuid.UUID `json:"office_id" db:"office_id"`
	DistrictInitials        *string    `json:"district_initials" db:"district_initials"`
	ProjectName             string     `json:"project_name" db:"project_name"`
	ProjectID               uuid.UUID  `json:"project_id" db:"project_id"`
	Month                   time.Time  `json:"month" db:"the_month"`
	ExpectedTotalSubmittals int        `json:"expected_total_submittals" db:"expected_total_submittals"`
	ActualTotalSubmittals   int        `json:"actual_total_submittals" db:"actual_total_submittals"`
	RedSubmittals           int        `json:"red_submittals" db:"red_submittals"`
	YellowSubmittals        int        `json:"yellow_submittals" db:"yellow_submittals"`
	GreenSubmittals         int        `json:"green_submittals" db:"green_submittals"`
}

const listEvaluationDistrictRollup = `
	SELECT * FROM v_district_rollup
	WHERE alert_type_id = 'da6ee89e-58cc-4d85-8384-43c3c33a68bd'::UUID
	AND project_id = $1
	AND the_month >= DATE_TRUNC('month', $2::TIMESTAMPTZ)
	AND the_month <= DATE_TRUNC('month', $3::TIMESTAMPTZ)
`

// ListCollectionGroups lists all collection groups for a project
func (q *Queries) ListEvaluationDistrictRollup(ctx context.Context, opID uuid.UUID, tw TimeWindow) ([]DistrictRollup, error) {
	dr := make([]DistrictRollup, 0)
	if err := q.db.SelectContext(ctx, &dr, listEvaluationDistrictRollup, opID, tw.After, tw.Before); err != nil {
		return nil, err
	}
	return dr, nil
}

const listMeasurementDistrictRollup = `
	SELECT * FROM v_district_rollup
	WHERE alert_type_id = '97e7a25c-d5c7-4ded-b272-1bb6e5914fe3'::UUID
	AND project_id = $1
	AND the_month >= DATE_TRUNC('month', $2::TIMESTAMPTZ)
	AND the_month <= DATE_TRUNC('month', $3::TIMESTAMPTZ)
`

// ListCollectionGroups lists all collection groups for a project
func (q *Queries) ListMeasurementDistrictRollup(ctx context.Context, opID uuid.UUID, tw TimeWindow) ([]DistrictRollup, error) {
	dr := make([]DistrictRollup, 0)
	if err := q.db.SelectContext(ctx, &dr, listMeasurementDistrictRollup, opID, tw.After, tw.Before); err != nil {
		return nil, err
	}
	return dr, nil
}
