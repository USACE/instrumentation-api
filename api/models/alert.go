package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// Alert is an alert, triggered by an AlertConfig evaluating to true
type Alert struct {
	Read          *bool     `json:"read,omitempty"`
	ID            uuid.UUID `json:"id"`
	AlertConfigID uuid.UUID `json:"alert_config_id" db:"alert_config_id"`
	ProjectID     uuid.UUID `json:"project_id" db:"project_id"`
	ProjectName   string    `json:"project_name" db:"project_name"`
	Name          string    `json:"name"`
	Body          string    `json:"body"`
	CreateDate    time.Time `json:"create_date" db:"create_date"`
}

// CreateAlerts creates one or more new alerts
func CreateAlerts(db *sqlx.DB, alertConfigIDS []uuid.UUID) error {
	txn, err := db.Beginx()
	if err != nil {
		return err
	}
	defer txn.Rollback()

	// Create Alert (CreateDate is a default now() in the database)
	stmt1, err := txn.Preparex(`INSERT INTO alert (alert_config_id) VALUES ($1)`)
	if err != nil {
		return err
	}
	for _, id := range alertConfigIDS {
		// Load Alert
		if _, err := stmt1.Exec(id); err != nil {
			return err
		}
	}
	if err := stmt1.Close(); err != nil {
		return err
	}
	if err := txn.Commit(); err != nil {
		return err
	}
	return nil
}

// ListProjectAlerts lists all alerts for a given instrument ID
func ListProjectAlerts(db *sqlx.DB, projectID *uuid.UUID) ([]Alert, error) {
	var aa []Alert
	err := db.Select(&aa, `
		SELECT * FROM v_alerts WHERE project_id = $1
	`, projectID)
	if err != nil {
		return make([]Alert, 0), err
	}
	return aa, nil
}

// ListAlertsForInstrument lists all alerts for a given instrument ID
func ListAlertsForInstrument(db *sqlx.DB, instrumentID *uuid.UUID) ([]Alert, error) {
	var aa []Alert
	err := db.Select(&aa, `
		SELECT * FROM v_alerts
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

// ListMyAlerts returns all alerts for which a profile is subscribed to the AlertConfig
func ListMyAlerts(db *sqlx.DB, profileID *uuid.UUID) ([]Alert, error) {
	aa := make([]Alert, 0)
	if err := db.Select(&aa, listMyAlertsSQL, profileID); err != nil {
		return make([]Alert, 0), err
	}
	return aa, nil
}

// GetMyAlert returns a single alert for which a profile is subscribed
func GetMyAlert(db *sqlx.DB, profileID *uuid.UUID, alertID *uuid.UUID) (*Alert, error) {
	var a Alert
	if err := db.Get(&a, getMyAlertSQL, profileID, alertID); err != nil {
		return nil, err
	}
	return &a, nil
}

// DoAlertRead marks an alert as read for a profile
func DoAlertRead(db *sqlx.DB, profileID *uuid.UUID, alertID *uuid.UUID) (*Alert, error) {
	if _, err := db.Exec(
		`INSERT INTO alert_read (profile_id, alert_id) VALUES ($1, $2)
		 ON CONFLICT DO NOTHING`, profileID, alertID,
	); err != nil {
		return nil, err
	}
	return GetMyAlert(db, profileID, alertID)
}

// DoAlertUnread marks an alert as unread for a profile
func DoAlertUnread(db *sqlx.DB, profileID *uuid.UUID, alertID *uuid.UUID) (*Alert, error) {
	if _, err := db.Exec(
		`DELETE FROM alert_read WHERE profile_id = $1 AND alert_id = $2`, profileID, alertID,
	); err != nil {
		return nil, err
	}
	return GetMyAlert(db, profileID, alertID)
}

// ListMyAlertsSQL returns all alerts for a profile's alert_profile_subscriptions
var listMyAlertsSQL = `SELECT a.*,
                              CASE WHEN r.alert_id IS NOT NULL THEN true
	                               ELSE false
                              END AS read
					   FROM v_alert a
					       LEFT JOIN alert_read r ON r.alert_id = a.id
                       WHERE a.alert_config_id IN (
                           SELECT alert_config_id
                           FROM alert_profile_subscription
                           WHERE profile_id = $1
					   )`

// GetMyAlertSQL returns a single alert
var getMyAlertSQL = listMyAlertsSQL + " AND a.id = $2"
