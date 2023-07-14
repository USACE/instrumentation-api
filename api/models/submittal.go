package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Submittal struct {
	ID                  uuid.UUID  `json:"id" db:"id"`
	AlertConfigID       uuid.UUID  `json:"alert_config_id" db:"alert_config_id"`
	AlertConfigName     string     `json:"alert_config_name" db:"alert_config_name"`
	AlertTypeID         uuid.UUID  `json:"alert_type_id" db:"alert_type_id"`
	AlertTypeName       string     `json:"alert_type_name" db:"alert_type_name"`
	ProjectID           uuid.UUID  `json:"project_id" db:"project_id"`
	SubmittalStatusID   uuid.UUID  `json:"submittal_status_id" db:"submittal_status_id"`
	SubmittalStatusName string     `json:"submittal_status_name" db:"submittal_status_name"`
	CompletionDate      *time.Time `json:"completion_date,omitempty" db:"completion_date"`
	CreateDate          time.Time  `json:"create_date" db:"create_date"`
	DueDate             time.Time  `json:"due_date" db:"due_date"`
	MarkedAsMissing     bool       `json:"marked_as_missing" db:"marked_as_missing"`
	WarningSent         bool       `json:"warning_sent" db:"warning_sent"`
}

func ListProjectSubmittals(db *sqlx.DB, projectID *uuid.UUID) ([]Submittal, error) {
	aa := make([]Submittal, 0)
	sql := `
		SELECT *
		FROM v_submittal
		WHERE project_id = $1
		ORDER BY due_date DESC, alert_type ASC
	`
	if err := db.Select(&aa, sql, projectID); err != nil {
		return aa, err
	}

	return aa, nil
}

func ListInstrumentSubmittals(db *sqlx.DB, instrumentID *uuid.UUID) ([]Submittal, error) {
	aa := make([]Submittal, 0)
	sql := `
		SELECT *
		FROM v_submittal
		WHERE id = ANY(
			SELECT sub.id
			FROM submittal sub
			INNER JOIN alert_config_instrument aci ON aci.alert_config_id = sub.alert_config_id
			WHERE aci.instrument_id = $1
		)
		ORDER BY due_date DESC
	`
	if err := db.Select(&aa, sql, instrumentID); err != nil {
		return aa, err
	}

	return aa, nil
}

func ListAlertConfigSubmittals(db *sqlx.DB, alertConfigID *uuid.UUID) ([]Submittal, error) {
	aa := make([]Submittal, 0)
	sql := `
		SELECT *
		FROM v_submittal
		WHERE alert_config_id = $1
		ORDER BY due_date DESC
	`
	if err := db.Select(&aa, sql, alertConfigID); err != nil {
		return aa, err
	}

	return aa, nil
}

func ListUnverifiedMissingSubmittals(db *sqlx.DB) ([]Submittal, error) {
	aa := make([]Submittal, 0)
	sql := `SELECT * FROM v_submittal WHERE completion_date IS NULL AND NOT marked_as_missing`

	if err := db.Select(&aa, sql); err != nil {
		return aa, err
	}

	return aa, nil
}

func UpdateSubmittal(db *sqlx.DB, sub Submittal) error {
	sql := `
		UPDATE submittal SET
			submittal_status_id = $2,
			completion_date = $3,
			warning_sent = $4
		WHERE id = $1
	`
	if _, err := db.Exec(sql, sub.ID, sub.SubmittalStatusID, sub.CompletionDate, sub.WarningSent); err != nil {
		return err
	}

	return nil
}

func MarkMissingSubmittal(db *sqlx.DB, submittalID *uuid.UUID) error {
	sql := `
		UPDATE submittal SET
			-- red submittal status
			submittal_status_id = '84a0f437-a20a-4ac2-8a5b-f8dc35e8489b'::UUID,
			marked_as_missing = true
		WHERE id = $1
	`
	if _, err := db.Exec(sql, submittalID); err != nil {
		return err
	}

	return nil
}
