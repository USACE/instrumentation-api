package model

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type AlertConfig struct {
	ID                      uuid.UUID                         `json:"id" db:"id"`
	Name                    string                            `json:"name" db:"name"`
	Body                    string                            `json:"body" db:"body"`
	ProjectID               uuid.UUID                         `json:"project_id" db:"project_id"`
	ProjectName             string                            `json:"project_name" db:"project_name"`
	AlertTypeID             uuid.UUID                         `json:"alert_type_id" db:"alert_type_id"`
	AlertType               string                            `json:"alert_type" db:"alert_type"`
	StartDate               time.Time                         `json:"start_date" db:"start_date"`
	ScheduleInterval        string                            `json:"schedule_interval" db:"schedule_interval"`
	RemindInterval          string                            `json:"remind_interval" db:"remind_interval"`
	WarningInterval         string                            `json:"warning_interval" db:"warning_interval"`
	LastChecked             *time.Time                        `json:"last_checked" db:"last_checked"`
	LastReminded            *time.Time                        `json:"last_reminded" db:"last_reminded"`
	Instruments             AlertConfigInstrumentCollection   `json:"instruments" db:"instruments"`
	AlertEmailSubscriptions EmailAutocompleteResultCollection `json:"alert_email_subscriptions" db:"alert_email_subscriptions"`
	CreatorUsername         string                            `json:"creator_username" db:"creator_username"`
	UpdaterUsername         *string                           `json:"updater_username" db:"updater_username"`
	MuteConsecutiveAlerts   bool                              `json:"mute_consecutive_alerts" db:"mute_consecutive_alerts"`
	CreateNextSubmittalFrom *time.Time                        `json:"-" db:"-"`
	AuditInfo
}

type AlertConfigInstrument struct {
	InstrumentID   uuid.UUID `json:"instrument_id" db:"instrument_id"`
	InstrumentName string    `json:"instrument_name" db:"instrument_name"`
}

type AlertConfigInstrumentCollection []AlertConfigInstrument

func (a *AlertConfigInstrumentCollection) Scan(src interface{}) error {
	if err := json.Unmarshal([]byte(src.(string)), a); err != nil {
		return err
	}
	return nil
}

func (a *AlertConfig) GetToAddresses() []string {
	emails := make([]string, len(a.AlertEmailSubscriptions))
	for idx := range a.AlertEmailSubscriptions {
		emails[idx] = a.AlertEmailSubscriptions[idx].Email
	}
	return emails
}

// GetAllAlertConfigsForProject lists all alert configs for a single project
func (q *Queries) GetAllAlertConfigsForProject(ctx context.Context, projectID *uuid.UUID) ([]AlertConfig, error) {
	c := `
		SELECT *
		FROM v_alert_config
		WHERE project_id = $1
		ORDER BY name
	`
	aa := make([]AlertConfig, 0)
	if err := q.db.SelectContext(ctx, &aa, c, projectID); err != nil {
		return aa, err
	}
	return aa, nil
}

// GetAllAlertConfigsForProjectAndAlertType lists alert configs for a single project filetered by alert type
func (q *Queries) GetAllAlertConfigsForProjectAndAlertType(ctx context.Context, projectID, alertTypeID *uuid.UUID) ([]AlertConfig, error) {
	c := `
		SELECT *
		FROM v_alert_config
		WHERE project_id = $1
		AND alert_type_id = $2
		ORDER BY name
	`
	aa := make([]AlertConfig, 0)
	if err := q.db.SelectContext(ctx, &aa, c, projectID, alertTypeID); err != nil {
		return aa, err
	}
	return aa, nil
}

// GetAllAlertConfigsForInstrument lists all alerts for a single instrument
func (q *Queries) GetAllAlertConfigsForInstrument(ctx context.Context, instrumentID *uuid.UUID) ([]AlertConfig, error) {
	c := `
		SELECT *
		FROM v_alert_config
		WHERE id = ANY(
			SELECT alert_config_id
			FROM alert_config_instrument
			WHERE instrument_id = $1
		)
		ORDER BY name
	`
	aa := make([]AlertConfig, 0)
	if err := q.db.SelectContext(ctx, &aa, c, instrumentID); err != nil {
		return aa, err
	}
	return aa, nil
}

// GetOneAlertConfig gets a single alert
func (q *Queries) GetOneAlertConfig(ctx context.Context, alertConfigID *uuid.UUID) (*AlertConfig, error) {
	c := `
		SELECT * FROM v_alert_config WHERE id = $1
	`
	var a AlertConfig
	if err := q.db.GetContext(ctx, &a, c, alertConfigID); err != nil {
		return nil, err
	}
	return &a, nil
}

func (q *Queries) CreateAlertConfig(ctx context.Context, ac *AlertConfig) (*uuid.UUID, error) {
	c := `
		INSERT INTO alert_config (
			project_id,
			name,
			body,
			alert_type_id,
			start_date,
			schedule_interval,
			mute_consecutive_alerts,
			remind_interval,
			warning_interval,
			creator,
			create_date
		) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)
		RETURNING id
	`
	var alertConfigID uuid.UUID
	if err := q.db.GetContext(ctx, &alertConfigID, c,
		ac.ProjectID,
		ac.Name,
		ac.Body,
		ac.AlertTypeID,
		ac.StartDate,
		ac.ScheduleInterval,
		ac.MuteConsecutiveAlerts,
		ac.RemindInterval,
		ac.WarningInterval,
		ac.Creator,
		ac.CreateDate,
	); err != nil {
		return nil, err
	}
	return &alertConfigID, nil
}

func (q *Queries) AssignInstrumentToAlertConfig(ctx context.Context, alertConfigID, instrumentID *uuid.UUID) error {
	c := `
		INSERT INTO alert_config_instrument (alert_config_id, instrument_id) VALUES ($1, $2)
	`
	if _, err := q.db.ExecContext(ctx, c, alertConfigID, instrumentID); err != nil {
		return err
	}
	return nil
}

func (q *Queries) UnassignAllInstrumentsFromAlertConfig(ctx context.Context, alertConfigID *uuid.UUID) error {
	c := `
		DELETE FROM alert_config_instrument WHERE alert_config_id = $1
	`
	if _, err := q.db.ExecContext(ctx, c, alertConfigID); err != nil {
		return err
	}
	return nil
}

func (q *Queries) CreateNextSubmittalFromExistingAlertConfigDate(ctx context.Context, alertConfigID *uuid.UUID) error {
	c := `
		INSERT INTO submittal (alert_config_id, due_date)
		SELECT id, create_date + schedule_interval
		FROM alert_config
		WHERE id = $1
	`
	if _, err := q.db.ExecContext(ctx, c, alertConfigID); err != nil {
		return err
	}
	return nil
}

func (q *Queries) UpdateAlertConfig(ctx context.Context, ac *AlertConfig) error {
	c := `
		UPDATE alert_config SET
			name = $3,
			body = $4,
			start_date = $5,
			schedule_interval = $6,
			mute_consecutive_alerts = $7,
			remind_interval = $8,
			warning_interval = $9,
			updater = $10,
			update_date = $11
		WHERE id = $1 AND project_id = $2
	`
	if _, err := q.db.ExecContext(ctx, c,
		ac.ID,
		ac.ProjectID,
		ac.Name,
		ac.Body,
		ac.StartDate,
		ac.ScheduleInterval,
		ac.MuteConsecutiveAlerts,
		ac.RemindInterval,
		ac.WarningInterval,
		ac.Updater,
		ac.UpdateDate,
	); err != nil {
		return err
	}
	return nil

}

func (q *Queries) UpdateFutureSubmittalForAlertConfig(ctx context.Context, alertConfigID *uuid.UUID) error {
	c := `
		UPDATE submittal
		SET due_date = sq.new_due_date
		FROM (
			SELECT
				sub.id AS submittal_id,
				sub.create_date + ac.schedule_interval AS new_due_date
			FROM submittal sub
			INNER JOIN alert_config ac ON sub.alert_config_id = ac.id
			WHERE sub.alert_config_id = $1
			AND sub.due_date > NOW()
			AND sub.completion_date IS NULL
			AND NOT sub.marked_as_missing
		) sq
		WHERE id = sq.submittal_id
		AND sq.new_due_date > NOW()
		RETURNING id
	`
	var updatedSubID uuid.UUID
	if err := q.db.GetContext(ctx, &updatedSubID, c, alertConfigID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("updated alert config new due date must be in the futute! complete the current submittal before updating")
		}
		return err
	}
	return nil
}

// DeleteAlertConfig deletes an alert by ID
func (q *Queries) DeleteAlertConfig(ctx context.Context, alertConfigID *uuid.UUID) error {
	c := `
		UPDATE alert_config SET deleted=true WHERE id = $1
	`
	if _, err := q.db.ExecContext(ctx, c, alertConfigID); err != nil {
		return err
	}
	return nil
}
