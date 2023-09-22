package model

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// Alert is an alert, triggered by an AlertConfig evaluating to true
type Alert struct {
	Read          *bool                           `json:"read,omitempty"`
	ID            uuid.UUID                       `json:"id"`
	AlertConfigID uuid.UUID                       `json:"alert_config_id" db:"alert_config_id"`
	ProjectID     uuid.UUID                       `json:"project_id" db:"project_id"`
	ProjectName   string                          `json:"project_name" db:"project_name"`
	Name          string                          `json:"name"`
	Body          string                          `json:"body"`
	CreateDate    time.Time                       `json:"create_date" db:"create_date"`
	Instruments   AlertConfigInstrumentCollection `json:"instruments" db:"instruments"`
}

// ListMyAlertsSQL returns all alerts for a profile's alert_profile_subscriptions
var listMyAlertsSQL = `
	SELECT a.*,
		CASE WHEN r.alert_id IS NOT NULL THEN true ELSE false
		END AS read
	FROM v_alert a
	LEFT JOIN alert_read r ON r.alert_id = a.id
	WHERE a.alert_config_id IN (
		SELECT alert_config_id
		FROM alert_profile_subscription
		WHERE profile_id = $1
	)
`

// CreateAlerts creates one or more new alerts
func (q *Queries) CreateAlerts(ctx context.Context, id uuid.UUID) error {
	if _, err := q.db.ExecContext(ctx, `INSERT INTO alert (alert_config_id) VALUES ($1)`, id); err != nil {
		return err
	}
	return nil
}

// GetAllAlertsForProject lists all alerts for a given instrument ID
func (q *Queries) GetAllAlertsForProject(ctx context.Context, projectID *uuid.UUID) ([]Alert, error) {
	aa := make([]Alert, 0)
	if err := q.db.SelectContext(ctx, &aa, `SELECT * FROM v_alert WHERE project_id = $1`, projectID); err != nil {
		return aa, err
	}
	return aa, nil
}

// GetAllAlertsForInstrument lists all alerts for a given instrument ID
func (q *Queries) GetAllAlertsForInstrument(ctx context.Context, instrumentID *uuid.UUID) ([]Alert, error) {
	aa := make([]Alert, 0)
	err := q.db.SelectContext(ctx, &aa, `
		SELECT * FROM v_alert
		WHERE alert_config_id = ANY(
			SELECT id FROM alert_config_instrument
			WHERE instrument_id = $1
		)
	`, instrumentID)
	if err != nil {
		return make([]Alert, 0), err
	}
	return aa, nil
}

// GetAllAlertsForProfile returns all alerts for which a profile is subscribed to the AlertConfig
func (q *Queries) GetAllAlertsForProfile(ctx context.Context, profileID *uuid.UUID) ([]Alert, error) {
	aa := make([]Alert, 0)
	if err := q.db.SelectContext(ctx, &aa, listMyAlertsSQL, profileID); err != nil {
		return aa, err
	}
	return aa, nil
}

// GetOneAlertForProfile returns a single alert for which a profile is subscribed
func (q *Queries) GetOneAlertForProfile(ctx context.Context, profileID *uuid.UUID, alertID *uuid.UUID) (*Alert, error) {
	var a Alert
	if err := q.db.GetContext(ctx, &a, listMyAlertsSQL+`AND a.id = $2`, profileID, alertID); err != nil {
		return nil, err
	}
	return &a, nil
}

// DoAlertRead marks an alert as read for a profile
func (q *Queries) DoAlertRead(ctx context.Context, profileID *uuid.UUID, alertID *uuid.UUID) (*Alert, error) {
	if _, err := q.db.ExecContext(ctx, `
		INSERT INTO alert_read (profile_id, alert_id) VALUES ($1, $2)
		ON CONFLICT DO NOTHING
	`, profileID, alertID); err != nil {
		return nil, err
	}
	return q.GetOneAlertForProfile(ctx, profileID, alertID)
}

// DoAlertUnread marks an alert as unread for a profile
func (q *Queries) DoAlertUnread(ctx context.Context, profileID *uuid.UUID, alertID *uuid.UUID) (*Alert, error) {
	if _, err := q.db.ExecContext(ctx, `
		DELETE FROM alert_read WHERE profile_id = $1 AND alert_id = $2
	`, profileID, alertID); err != nil {
		return nil, err
	}
	return q.GetOneAlertForProfile(ctx, profileID, alertID)
}
