package model

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
)

// AlertSubscription is a profile subscription to an alert
type AlertSubscription struct {
	ID            uuid.UUID `json:"id"`
	AlertConfigID uuid.UUID `json:"alert_config_id" db:"alert_config_id"`
	ProfileID     uuid.UUID `json:"profile_id" db:"profile_id"`
	AlertSubscriptionSettings
}

// AlertSubscriptionSettings holds all settings for an AlertSubscription
type AlertSubscriptionSettings struct {
	MuteUI     bool `json:"mute_ui" db:"mute_ui"`
	MuteNotify bool `json:"mute_notify" db:"mute_notify"`
}

// AlertSubscriptionCollection is a collection of AlertSubscription items
type AlertSubscriptionCollection struct {
	Items []AlertSubscription `json:"items"`
}

// EmailAlert is an email subscription to an alert
type EmailAlert struct {
	ID            uuid.UUID `json:"id"`
	AlertConfigID uuid.UUID `json:"alert_config_id"`
	EmailID       uuid.UUID `json:"profile_id"`
	MuteNotify    bool      `json:"mute_notify" db:"mute_notify"`
}

type Email struct {
	ID    uuid.UUID `json:"id" db:"id"`
	Email string    `json:"email" db:"email"`
}

// UnmarshalJSON implements the UnmarshalJSON Interface for AlertSubscription
func (c *AlertSubscriptionCollection) UnmarshalJSON(b []byte) error {

	switch JSONType(b) {
	case "ARRAY":
		if err := json.Unmarshal(b, &c.Items); err != nil {
			return err
		}
	case "OBJECT":
		var a AlertSubscription
		if err := json.Unmarshal(b, &a); err != nil {
			return err
		}
		c.Items = []AlertSubscription{a}
	default:
		c.Items = make([]AlertSubscription, 0)
	}
	return nil
}

// SubscribeProfileToAlerts subscribes a profile to an instrument alert
func (q *Queries) SubscribeProfileToAlerts(ctx context.Context, alertConfigID, profileID *uuid.UUID) error {
	c := `
		INSERT INTO alert_profile_subscription (alert_config_id, profile_id)
		VALUES ($1, $2)
		ON CONFLICT DO NOTHING
	`
	if _, err := q.db.ExecContext(ctx, c, alertConfigID, profileID); err != nil {
		return err
	}
	return nil
}

// UnsubscribeProfileToAlerts subscribes a profile to an instrument alert
func (q *Queries) UnsubscribeProfileToAlerts(ctx context.Context, alertConfigID, profileID *uuid.UUID) error {
	c := `
		DELETE FROM alert_profile_subscription WHERE alert_config_id = $1 AND profile_id = $2
	`
	if _, err := q.db.ExecContext(ctx, c, alertConfigID, profileID); err != nil {
		return err
	}
	return nil
}

// GetAlertSubscription returns a AlertSubscription
func (q *Queries) GetAlertSubscription(ctx context.Context, alertConfigID, profileID *uuid.UUID) (*AlertSubscription, error) {
	c := `
		SELECT * FROM alert_profile_subscription WHERE alert_config_id = $1 AND profile_id = $2
	`
	var pa AlertSubscription
	if err := q.db.GetContext(ctx, &pa, c, alertConfigID, profileID); err != nil {
		return nil, err
	}
	return &pa, nil
}

// GetAlertSubscriptionByID returns an alert subscription
func (q *Queries) GetAlertSubscriptionByID(ctx context.Context, id *uuid.UUID) (*AlertSubscription, error) {
	c := `
		SELECT * FROM alert_profile_subscription WHERE id = $1
	`
	var s AlertSubscription
	if err := q.db.GetContext(ctx, &s, c, id); err != nil {
		return nil, err
	}
	return &s, nil
}

// ListMyAlertSubscriptions returns all profile_alerts for a given profile ID
func (q *Queries) ListMyAlertSubscriptions(ctx context.Context, profileID *uuid.UUID) ([]AlertSubscription, error) {
	c := `
		SELECT * FROM alert_profile_subscription WHERE profile_id = $1
	`
	ss := make([]AlertSubscription, 0)
	if err := q.db.SelectContext(ctx, &ss, c, profileID); err != nil {
		return make([]AlertSubscription, 0), err
	}
	return ss, nil
}

// UpdateMyAlertSubscription updates properties on a AlertSubscription
func (q *Queries) UpdateMyAlertSubscription(ctx context.Context, s *AlertSubscription) error {
	c := `
		UPDATE alert_profile_subscription SET mute_ui=$1, mute_notify=$2 WHERE alert_config_id=$3 AND profile_id=$4
	`
	_, err := q.db.ExecContext(ctx, c, s.MuteUI, s.MuteNotify, s.AlertConfigID, s.ProfileID)
	if err != nil {
		return err
	}
	return nil
}

func (q *Queries) RegisterEmail(ctx context.Context, emailAddress string) (*uuid.UUID, error) {
	c := `
		WITH e AS (
			INSERT INTO email (email) VALUES ($1)
			ON CONFLICT ON CONSTRAINT unique_email DO NOTHING
			RETURNING id
		)
		SELECT id FROM e
		UNION
		SELECT id from email WHERE email = $1
	`
	var newID uuid.UUID
	if err := q.db.GetContext(ctx, &newID, c, emailAddress); err != nil {
		return nil, err
	}
	return &newID, nil
}

func (q *Queries) UnregisterEmail(ctx context.Context, emailID *uuid.UUID) error {
	c := `
		DELETE FROM email WHERE id = $1
	`
	if _, err := q.db.ExecContext(ctx, c, emailID); err != nil {
		return err
	}
	return nil
}

func (q *Queries) SubscribeEmailToAlertConfig(ctx context.Context, alertConfigID, emailID *uuid.UUID) error {
	c := `
		INSERT INTO alert_email_subscription (alert_config_id, email_id) VALUES ($1,$2)
		ON CONFLICT ON CONSTRAINT email_unique_alert_config DO NOTHING
	`
	if _, err := q.db.ExecContext(ctx, c, alertConfigID, emailID); err != nil {
		return err
	}
	return nil
}

func (q *Queries) SubscribeProfileToAlertConfig(ctx context.Context, alertConfigID, emailID *uuid.UUID) error {
	c := `
		INSERT INTO alert_profile_subscription (alert_config_id, profile_id) VALUES ($1,$2)
		ON CONFLICT ON CONSTRAINT profile_unique_alert_config DO NOTHING
	`
	if _, err := q.db.ExecContext(ctx, c, alertConfigID, emailID); err != nil {
		return err
	}
	return nil
}

func (q *Queries) UnsubscribeEmailFromAlertConfig(ctx context.Context, alertConfigID, emailID *uuid.UUID) error {
	c := `
		DELETE FROM alert_email_subscription WHERE alert_config_id = $1 AND email_id = $2
	`
	if _, err := q.db.ExecContext(ctx, c, alertConfigID, emailID); err != nil {
		return err
	}
	return nil
}

func (q *Queries) UnsubscribeProfileFromAlertConfig(ctx context.Context, alertConfigID, emailID *uuid.UUID) error {
	c := `
		DELETE FROM alert_profile_subscription WHERE alert_config_id = $1 AND profile_id = $2
	`
	if _, err := q.db.ExecContext(ctx, c, alertConfigID, emailID); err != nil {
		return err
	}
	return nil
}

func (q *Queries) UnsubscribeAllEmailsFromAlertConfig(ctx context.Context, alertConfigID *uuid.UUID) error {
	c := `
		DELETE FROM alert_email_subscription WHERE alert_config_id = $1
	`
	if _, err := q.db.ExecContext(ctx, c, alertConfigID); err != nil {
		return err
	}
	return nil
}

func (q *Queries) UnsubscribeAllProfilesFromAlertConfig(ctx context.Context, alertConfigID *uuid.UUID) error {
	c := `
		DELETE FROM alert_profile_subscription WHERE alert_config_id = $1
	`
	if _, err := q.db.ExecContext(ctx, c, alertConfigID); err != nil {
		return err
	}
	return nil
}
