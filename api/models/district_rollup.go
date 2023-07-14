package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type DistrictRollup struct {
	AlertConfigID           uuid.UUID `json:"alert_config_id" db:"alert_config_id"`
	ProjectName             string    `json:"project_name" db:"project_name"`
	AlertTypeID             uuid.UUID `json:"-" db:"alert_type_id"`
	Month                   time.Time `json:"month" db:"the_month"`
	ExpectedTotalSubmittals int       `json:"expected_total_submittals" db:"expected_total_submittals"`
	ActualTotalSubmittals   int       `json:"actual_total_submittals" db:"actual_total_submittals"`
	RedSubmittals           int       `json:"red_submittals" db:"red_submittals"`
	YellowSubmittals        int       `json:"yellow_submittals" db:"yellow_submittals"`
	GreenSubmittals         int       `json:"green_submittals" db:"green_submittals"`
}

// ListCollectionGroups lists all collection groups for a project
func ListEvaluationDistrictRollup(db *sqlx.DB) ([]DistrictRollup, error) {
	cc := make([]DistrictRollup, 0)
	if err := db.Select(&cc, `SELECT * FROM v_district_rollup WHERE alert_type_id = 'da6ee89e-58cc-4d85-8384-43c3c33a68bd'::UUID`); err != nil {
		return make([]DistrictRollup, 0), err
	}
	return cc, nil
}

// ListCollectionGroups lists all collection groups for a project
func ListMeasurementDistrictRollup(db *sqlx.DB) ([]DistrictRollup, error) {
	cc := make([]DistrictRollup, 0)
	if err := db.Select(&cc, `SELECT * FROM v_district_rollup WHERE alert_type_id = '97e7a25c-d5c7-4ded-b272-1bb6e5914fe3'::UUID`); err != nil {
		return make([]DistrictRollup, 0), err
	}
	return cc, nil
}
