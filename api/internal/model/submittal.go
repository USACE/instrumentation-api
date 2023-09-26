package model

import (
	"context"
	"time"

	"github.com/google/uuid"
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
	CompletionDate      *time.Time `json:"completion_date" db:"completion_date"`
	CreateDate          time.Time  `json:"create_date" db:"create_date"`
	DueDate             time.Time  `json:"due_date" db:"due_date"`
	MarkedAsMissing     bool       `json:"marked_as_missing" db:"marked_as_missing"`
	WarningSent         bool       `json:"warning_sent" db:"warning_sent"`
}

const missingFilter = `
	AND completion_date IS NULL AND NOT marked_as_missing
`

func (q *Queries) ListProjectSubmittals(ctx context.Context, projectID uuid.UUID, showMissing bool) ([]Submittal, error) {
	var filter string
	if showMissing {
		filter = missingFilter
	}
	listProjectSubmittals := `
		SELECT *
		FROM v_submittal
		WHERE project_id = $1
		` + filter + `
		ORDER BY due_date DESC, alert_type_name ASC
	`

	aa := make([]Submittal, 0)
	if err := q.db.SelectContext(ctx, &aa, listProjectSubmittals, projectID); err != nil {
		return aa, err
	}
	return aa, nil
}

func (q *Queries) ListInstrumentSubmittals(ctx context.Context, instrumentID uuid.UUID, showMissing bool) ([]Submittal, error) {
	var filter string
	if showMissing {
		filter = missingFilter
	}
	listInstrumentSubmittals := `
		SELECT *
		FROM v_submittal
		WHERE id = ANY(
			SELECT sub.id
			FROM submittal sub
			INNER JOIN alert_config_instrument aci ON aci.alert_config_id = sub.alert_config_id
			WHERE aci.instrument_id = $1
		)
		` + filter + `
		ORDER BY due_date DESC
	`
	aa := make([]Submittal, 0)
	if err := q.db.SelectContext(ctx, &aa, listInstrumentSubmittals, instrumentID); err != nil {
		return aa, err
	}
	return aa, nil
}

func (q *Queries) ListAlertConfigSubmittals(ctx context.Context, alertConfigID uuid.UUID, showMissing bool) ([]Submittal, error) {
	var filter string
	if showMissing {
		filter = missingFilter
	}
	listAlertConfigSubmittals := `
		SELECT *
		FROM v_submittal
		WHERE alert_config_id = $1
		` + filter + `
		ORDER BY due_date DESC
	`
	aa := make([]Submittal, 0)
	if err := q.db.SelectContext(ctx, &aa, listAlertConfigSubmittals, alertConfigID); err != nil {
		return aa, err
	}
	return aa, nil
}

const listUnverifiedMissingSubmittals = `
	SELECT *
	FROM v_submittal
	WHERE completion_date IS NULL
	AND NOT marked_as_missing
	ORDER BY due_date DESC
`

func (q *Queries) ListUnverifiedMissingSubmittals(ctx context.Context) ([]Submittal, error) {
	aa := make([]Submittal, 0)
	if err := q.db.SelectContext(ctx, &aa, listUnverifiedMissingSubmittals); err != nil {
		return nil, err
	}
	return aa, nil
}

const updateSubmittal = `
	UPDATE submittal SET
		submittal_status_id = $2,
		completion_date = $3,
		warning_sent = $4
	WHERE id = $1
`

func (q *Queries) UpdateSubmittal(ctx context.Context, sub Submittal) error {
	_, err := q.db.ExecContext(ctx, updateSubmittal, sub.ID, sub.SubmittalStatusID, sub.CompletionDate, sub.WarningSent)
	return err
}

const verifyMissingSubmittal = `
	UPDATE submittal SET
		-- red submittal status
		submittal_status_id = '84a0f437-a20a-4ac2-8a5b-f8dc35e8489b'::UUID,
		marked_as_missing = true
	WHERE id = $1
	AND completion_date IS NULL
	AND NOW() > due_date
`

func (q *Queries) VerifyMissingSubmittal(ctx context.Context, submittalID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, verifyMissingSubmittal, submittalID)
	return err
}

const verifyMissingAlertConfigSubmittals = `
	UPDATE submittal SET
		submittal_status_id = '84a0f437-a20a-4ac2-8a5b-f8dc35e8489b'::UUID,
		marked_as_missing = true
	WHERE alert_config_id = $1
	AND completion_date IS NULL
	AND NOW() > due_date
`

func (q *Queries) VerifyMissingAlertConfigSubmittals(ctx context.Context, alertConfigID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, verifyMissingAlertConfigSubmittals, alertConfigID)
	return err
}
