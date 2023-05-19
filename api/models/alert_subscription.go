package models

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
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
func SubscribeProfileToAlerts(db *sqlx.DB, alertConfigID *uuid.UUID, profileID *uuid.UUID) (*AlertSubscription, error) {
	_, err := db.Exec(
		`INSERT INTO alert_profile_subscription (alert_config_id, profile_id) VALUES ($1, $2)
		 ON CONFLICT DO NOTHING`, alertConfigID, profileID,
	)
	if err != nil {
		return nil, err
	}
	return GetAlertSubscription(db, alertConfigID, profileID)
}

// UnsubscribeProfileToAlerts subscribes a profile to an instrument alert
func UnsubscribeProfileToAlerts(db *sqlx.DB, alertConfigID *uuid.UUID, profileID *uuid.UUID) error {
	if _, err := db.Exec(
		`DELETE FROM alert_profile_subscription WHERE alert_config_id = $1 AND profile_id = $2`, alertConfigID, profileID,
	); err != nil {
		return err
	}
	return nil
}

// GetAlertSubscription returns a AlertSubscription
func GetAlertSubscription(db *sqlx.DB, alertConfigID *uuid.UUID, profileID *uuid.UUID) (*AlertSubscription, error) {
	var pa AlertSubscription
	if err := db.Get(
		&pa, `SELECT * FROM alert_profile_subscription WHERE alert_config_id = $1 AND profile_id = $2`, alertConfigID, profileID,
	); err != nil {
		return nil, err
	}
	return &pa, nil
}

// GetAlertSubscriptionByID returns an alert subscription
func GetAlertSubscriptionByID(db *sqlx.DB, id *uuid.UUID) (*AlertSubscription, error) {
	var s AlertSubscription
	if err := db.Get(&s, `SELECT * FROM alert_profile_subscription WHERE id = $1`, id); err != nil {
		return nil, err
	}
	return &s, nil
}

// ListMyAlertSubscriptions returns all profile_alerts for a given profile ID
func ListMyAlertSubscriptions(db *sqlx.DB, profileID *uuid.UUID) ([]AlertSubscription, error) {
	ss := make([]AlertSubscription, 0)
	if err := db.Select(
		&ss, `SELECT * FROM alert_profile_subscription WHERE profile_id = $1`, profileID,
	); err != nil {
		return make([]AlertSubscription, 0), err
	}
	return ss, nil
}

// UpdateMyAlertSubscription updates properties on a AlertSubscription
func UpdateMyAlertSubscription(db *sqlx.DB, s *AlertSubscription) (*AlertSubscription, error) {

	_, err := db.Exec(
		"UPDATE alert_profile_subscription SET mute_ui=$1, mute_notify=$2 WHERE alert_config_id=$3 AND profile_id=$4",
		s.MuteUI, s.MuteNotify, s.AlertConfigID, s.ProfileID,
	)
	if err != nil {
		return nil, err
	}
	return GetAlertSubscription(db, &s.AlertConfigID, &s.ProfileID)
}

func SubscribeEmailsToAlertConfigTxn(txn *sqlx.Tx, alertConfigID *uuid.UUID, emails EmailAutocompleteResultCollection) error {
	registerStmt, err := txn.Preparex(`
		WITH e AS (
			INSERT INTO email (email) VALUES ($1)
			ON CONFLICT ON CONSTRAINT unique_email DO NOTHING
			RETURNING id
		)
		SELECT id FROM e
		UNION
		SELECT id from email WHERE email=$1
	`)
	if err != nil {
		return err
	}
	emailStmt, err := txn.Preparex(`
		INSERT INTO alert_email_subscription (alert_config_id, email_id) VALUES ($1,$2)
		ON CONFLICT ON CONSTRAINT email_unique_alert_config DO NOTHING
	`)
	if err != nil {
		return err
	}
	profileStmt, err := txn.Preparex(`
		INSERT INTO alert_profile_subscription (alert_config_id, profile_id) VALUES ($1,$2)
		ON CONFLICT ON CONSTRAINT profile_unique_alert_config DO NOTHING
	`)
	if err != nil {
		return err
	}

	// Register any emails that are not yet in system
	for idx, em := range emails {
		if em.UserType == "" {
			var newID uuid.UUID
			if err := registerStmt.Get(&newID, em.Email); err != nil {
				return err
			}
			emails[idx].ID = newID
			emails[idx].UserType = "email"
		}
	}
	// Subscribe emails
	for _, em := range emails {
		if em.UserType == "email" {
			if _, err := emailStmt.Exec(alertConfigID, em.ID); err != nil {
				return err
			}
		} else if em.UserType == "profile" {
			if _, err := profileStmt.Exec(alertConfigID, em.ID); err != nil {
				return err
			}
		} else {
			return fmt.Errorf("unable to unsubscribe email %s: user type %s does not exist, aborting transaction", em.Email, em.UserType)
		}
	}

	if err := registerStmt.Close(); err != nil {
		return err
	}
	if err := emailStmt.Close(); err != nil {
		return err
	}
	if err := profileStmt.Close(); err != nil {
		return err
	}
	return nil
}

func UnsubscribeEmailsToAlertConfigTxn(txn *sqlx.Tx, alertConfigID *uuid.UUID, emails EmailAutocompleteResultCollection) error {
	emailStmt, err := txn.Preparex(`
		DELETE FROM alert_email_subscription WHERE alert_config_id=$1 AND email_id=$2
	`)
	if err != nil {
		return err
	}
	profileStmt, err := txn.Preparex(`
		DELETE FROM alert_profile_subscription WHERE alert_config_id=$1 AND profile_id=$2
	`)
	if err != nil {
		return err
	}

	for _, em := range emails {
		if em.UserType == "" {
			return fmt.Errorf("required field user_type is null, aborting transaction")
		} else if em.UserType == "email" {
			if _, err := emailStmt.Exec(alertConfigID, em.ID); err != nil {
				return err
			}
		} else if em.UserType == "profile" {
			if _, err := profileStmt.Exec(alertConfigID, em.ID); err != nil {
				return err
			}
		} else {
			return fmt.Errorf("unable to unsubscribe email %s: user type %s does not exist, aborting transaction", em.Email, em.UserType)
		}
	}
	if err := emailStmt.Close(); err != nil {
		return err
	}
	if err := profileStmt.Close(); err != nil {
		return err
	}
	return nil
}

func UnsubscribeAllEmailsToAlertConfigTxn(txn *sqlx.Tx, alertConfigID *uuid.UUID, emails EmailAutocompleteResultCollection) error {
	emailStmt, err := txn.Preparex(`
		DELETE FROM alert_email_subscription WHERE alert_config_id=$1
	`)
	if err != nil {
		return err
	}
	profileStmt, err := txn.Preparex(`
		DELETE FROM alert_profile_subscription WHERE alert_config_id=$1
	`)
	if err != nil {
		return err
	}

	for _, em := range emails {
		if em.UserType == "" {
			return fmt.Errorf("required field user_type is null, aborting transaction")
		} else if em.UserType == "email" {
			if _, err := emailStmt.Exec(alertConfigID); err != nil {
				return err
			}
		} else if em.UserType == "profile" {
			if _, err := profileStmt.Exec(alertConfigID); err != nil {
				return err
			}
		} else {
			return fmt.Errorf("unable to unsubscribe email %s: user type %s does not exist, aborting transaction", em.Email, em.UserType)
		}
	}
	if err := emailStmt.Close(); err != nil {
		return err
	}
	if err := profileStmt.Close(); err != nil {
		return err
	}
	return nil
}

func SubscribeEmailsToAlertConfig(db *sqlx.DB, alertConfigID *uuid.UUID, emails *EmailAutocompleteResultCollection) (*AlertConfig, error) {
	txn, err := db.Beginx()
	if err != nil {
		return nil, err
	}
	defer txn.Rollback()

	if err := SubscribeEmailsToAlertConfigTxn(txn, alertConfigID, *emails); err != nil {
		return nil, err
	}
	if err := txn.Commit(); err != nil {
		return nil, err
	}
	return GetAlertConfig(db, alertConfigID)
}

func UnsubscribeEmailsToAlertConfig(db *sqlx.DB, alertConfigID *uuid.UUID, emails *EmailAutocompleteResultCollection) (*AlertConfig, error) {
	txn, err := db.Beginx()
	if err != nil {
		return nil, err
	}
	defer txn.Rollback()

	if err := UnsubscribeEmailsToAlertConfigTxn(txn, alertConfigID, *emails); err != nil {
		return nil, err
	}
	if err := txn.Commit(); err != nil {
		return nil, err
	}
	return GetAlertConfig(db, alertConfigID)
}

func DeleteEmail(db *sqlx.DB, emailID *uuid.UUID) error {
	if _, err := db.Exec(`DELETE FROM email WHERE id = $1`, emailID); err != nil {
		return err
	}
	return nil
}
