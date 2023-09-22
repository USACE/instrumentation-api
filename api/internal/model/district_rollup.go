package model

import (
	"time"

	"github.com/USACE/instrumentation-api/api/internal/timeseries"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
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

// ListCollectionGroups lists all collection groups for a project
func ListEvaluationDistrictRollup(db *sqlx.DB, opID uuid.UUID, tw timeseries.TimeWindow) ([]DistrictRollup, error) {
	dr := make([]DistrictRollup, 0)
	if err := db.Select(&dr, `
		SELECT * FROM v_district_rollup
		WHERE alert_type_id = 'da6ee89e-58cc-4d85-8384-43c3c33a68bd'::UUID
		AND project_id = $1
		AND the_month >= DATE_TRUNC('month', $2::TIMESTAMPTZ)
		AND the_month <= DATE_TRUNC('month', $3::TIMESTAMPTZ)
	`, opID, tw.Start, tw.End); err != nil {
		return dr, err
	}
	return dr, nil
}

// ListCollectionGroups lists all collection groups for a project
func ListMeasurementDistrictRollup(db *sqlx.DB, opID uuid.UUID, tw timeseries.TimeWindow) ([]DistrictRollup, error) {
	dr := make([]DistrictRollup, 0)
	if err := db.Select(&dr, `
		SELECT * FROM v_district_rollup
		WHERE alert_type_id = '97e7a25c-d5c7-4ded-b272-1bb6e5914fe3'::UUID
		AND project_id = $1
		AND the_month >= DATE_TRUNC('month', $2::TIMESTAMPTZ)
		AND the_month <= DATE_TRUNC('month', $3::TIMESTAMPTZ)
	`, opID, tw.Start, tw.End); err != nil {
		return dr, err
	}
	return dr, nil
}
