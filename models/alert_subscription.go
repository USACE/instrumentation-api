package models

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// AlertSubscription is a profile subscription to an alert
type AlertSubscription struct {
	ID        uuid.UUID `json:"id"`
	AlertID   uuid.UUID `json:"alert_id" db:"alert_id"`
	ProfileID uuid.UUID `json:"profile_id" db:"profile_id"`
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
	ID         uuid.UUID `json:"id"`
	AlertID    uuid.UUID `json:"alert_id"`
	EmailID    uuid.UUID `json:"profile_id"`
	MuteNotify bool      `json:"mute_notify" db:"mute_notify"`
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
func SubscribeProfileToAlerts(db *sqlx.DB, alertID *uuid.UUID, profileID *uuid.UUID) (*AlertSubscription, error) {
	var pa AlertSubscription
	err := db.QueryRowx(
		`INSERT INTO profile_alerts (alert_id, profile_id) VALUES ($1, $2) RETURNING *`, alertID, profileID,
	).StructScan(&pa)
	if err != nil {
		return nil, err
	}
	return &pa, nil
}

// UnsubscribeProfileToAlerts subscribes a profile to an instrument alert
func UnsubscribeProfileToAlerts(db *sqlx.DB, alertID *uuid.UUID, profileID *uuid.UUID) error {
	if _, err := db.Exec(
		`DELETE FROM profile_alerts WHERE alert_id = $1 AND profile_id = $2`, alertID, profileID,
	); err != nil {
		return err
	}
	return nil
}

// GetAlertSubscription returns a AlertSubscription
func GetAlertSubscription(db *sqlx.DB, alertID *uuid.UUID, profileID *uuid.UUID) (*AlertSubscription, error) {
	var pa AlertSubscription
	if err := db.Get(
		&pa, `SELECT * FROM profile_alerts WHERE alert_id = $1 AND profile_id = $2`, alertID, profileID,
	); err != nil {
		return nil, err
	}
	return &pa, nil
}

// GetAlertSubscriptionByID returns an alert subscription
func GetAlertSubscriptionByID(db *sqlx.DB, id *uuid.UUID) (*AlertSubscription, error) {
	var s AlertSubscription
	if err := db.Get(&s, `SELECT * FROM profile_alerts WHERE id = $1`, id); err != nil {
		return nil, err
	}
	return &s, nil
}

// ListMyAlertSubscriptions returns all profile_alerts for a given profile ID
func ListMyAlertSubscriptions(db *sqlx.DB, profileID *uuid.UUID) ([]AlertSubscription, error) {
	ss := make([]AlertSubscription, 0)
	if err := db.Select(
		&ss, `SELECT * FROM profile_alerts WHERE profile_id = $1`, profileID,
	); err != nil {
		return make([]AlertSubscription, 0), err
	}
	return ss, nil
}

// UpdateMyAlertSubscription updates properties on a AlertSubscription
func UpdateMyAlertSubscription(db *sqlx.DB, s *AlertSubscription) (*AlertSubscription, error) {

	_, err := db.Exec(
		"UPDATE profile_alerts SET mute_ui=$1, mute_notify=$2 WHERE alert_id=$3 AND profile_id=$4",
		s.MuteUI, s.MuteNotify, s.AlertID, s.ProfileID,
	)
	if err != nil {
		return nil, err
	}
	return GetAlertSubscription(db, &s.AlertID, &s.ProfileID)
}
