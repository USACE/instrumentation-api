package models

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgtype"
	"github.com/jmoiron/sqlx"
)

type AlertConfig struct {
	ID                      uuid.UUID                              `json:"id" db:"id"`
	Name                    string                                 `json:"name" db:"name"`
	Body                    string                                 `json:"body" db:"body"`
	ProjectID               uuid.UUID                              `json:"project_id" db:"project_id"`
	ProjectName             string                                 `json:"project_name" db:"project_name"`
	AlertTypeID             uuid.UUID                              `json:"alert_type_id" db:"alert_type_id"`
	AlertType               string                                 `json:"alert_type" db:"alert_type"`
	StartDate               time.Time                              `json:"start_date" db:"start_date"`
	ScheduleInterval        Duration                               `json:"schedule_interval" db:"schedule_interval"`
	NMissedBeforeAlert      int                                    `json:"n_missed_before_alert" db:"n_missed_before_alert"`
	RemindInterval          Duration                               `json:"remind_interval" db:"remind_interval"`
	WarningInterval         *Duration                              `json:"warning_interval" db:"warning_interval"`
	LastChecked             *time.Time                             `json:"last_checked" db:"last_checked"`
	AlertStatus             string                                 `json:"alert_status" db:"alert_status"`
	LastReminded            *time.Time                             `json:"last_reminded" db:"last_reminded"`
	Instruments             AlertConfigInstrumentCollection        `json:"instruments" db:"instruments"`
	AlertEmailSubscriptions AlertConfigEmailSubscriptionCollection `json:"alert_email_subscriptions" db:"alert_email_subscriptions"`
	CreatorUsername         string                                 `json:"creator_username" db:"creator_username"`
	UpdaterUsername         *string                                `json:"updater_username" db:"updater_username"`
	AuditInfo
}

type AlertConfigInstrument struct {
	InstrumentID   uuid.UUID `json:"instrument_id" db:"instrument_id"`
	InstrumentName string    `json:"instrument_name" db:"instrument_name"`
}

type AlertConfigEmailSubscription struct {
	Username *string `json:"username" db:"username"`
	Email    string  `json:"email" db:"email"`
}

type Duration string

func (a *Duration) Scan(src interface{}) error {
	var iv pgtype.Interval
	var buf []byte

	if err := iv.Scan(src); err != nil {
		return err
	}
	buf, err := iv.EncodeText(nil, buf)
	if err != nil {
		return err
	}
	*a = Duration(buf)
	return nil
}
func (a Duration) Value() (driver.Value, error) {
	var iv pgtype.Interval
	if err := iv.DecodeText(nil, []byte(a)); err != nil {
		return nil, err
	}
	return iv.Value()
}
func (a Duration) ToNsDuration() (time.Duration, error) {
	var iv pgtype.Interval
	var d time.Duration

	if err := iv.DecodeText(nil, []byte(a)); err != nil {
		return 0, err
	}
	if err := iv.AssignTo(&d); err != nil {
		return 0, err
	}
	return d, nil
}

type AlertConfigInstrumentCollection []AlertConfigInstrument

func (a *AlertConfigInstrumentCollection) Scan(src interface{}) error {
	if err := json.Unmarshal([]byte(src.(string)), a); err != nil {
		return err
	}
	return nil
}

type AlertConfigEmailSubscriptionCollection []AlertConfigEmailSubscription

func (a *AlertConfigEmailSubscriptionCollection) Scan(src interface{}) error {
	if err := json.Unmarshal([]byte(src.(string)), a); err != nil {
		return err
	}
	return nil
}

// ListProjectAlertConfigs lists all alerts for a single project
func ListProjectAlertConfigs(db *sqlx.DB, projectID *uuid.UUID) ([]AlertConfig, error) {
	var aa []AlertConfig
	sql := `SELECT *
			FROM v_alert_config
			WHERE project_id = $1
	`
	err := db.Select(&aa, sql, projectID)
	if err != nil {
		return make([]AlertConfig, 0), err
	}

	return aa, nil
}

// ListInstrumentAlertConfigs lists all alerts for a single instrument
func ListInstrumentAlertConfigs(db *sqlx.DB, instrumentID *uuid.UUID) ([]AlertConfig, error) {
	var aa []AlertConfig
	sql := `
		SELECT *
		FROM v_alert_config
		WHERE id = ANY(
			SELECT alert_config_id
			FROM alert_config_instrument
			WHERE instrument_id = $1
		)
	`
	err := db.Select(&aa, sql, instrumentID)
	if err != nil {
		return make([]AlertConfig, 0), err
	}

	return aa, nil
}

// GetAlertConfig gets a single alert
func GetAlertConfig(db *sqlx.DB, alertConfigID *uuid.UUID) (*AlertConfig, error) {
	var a AlertConfig
	sql := `SELECT * FROM v_alert_config WHERE id = $1`
	err := db.Get(&a, sql, alertConfigID)
	if err != nil {
		return nil, err
	}

	return &a, nil
}

// CreateAlertConfig creates one new alert configuration
func CreateAlertConfig(db *sqlx.DB, ac *AlertConfig) (*AlertConfig, error) {
	txn, err := db.Beginx()
	if err != nil {
		return nil, err
	}
	defer txn.Rollback()

	stmt1, err := txn.Preparex(`
		INSERT INTO alert_config
			(
				project_id,
				name,
				body,
				alert_type_id,
				start_date,
				schedule_interval,
				n_missed_before_alert,
				remind_interval,
				warning_interval,
				creator,
				create_date
			) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)
			RETURNING id
	`)
	if err != nil {
		return nil, err
	}
	stmt2, err := txn.Preparex(`
		INSERT INTO alert_config_instrument (alert_config_id, instrument_id) VALUES ($1, $2)
	`)
	if err != nil {
		return nil, err
	}

	var alertConfigID uuid.UUID
	if err := stmt1.Get(
		&alertConfigID,
		ac.ProjectID,
		ac.Name,
		ac.Body,
		ac.AlertTypeID,
		ac.StartDate,
		ac.ScheduleInterval,
		ac.NMissedBeforeAlert,
		ac.RemindInterval,
		ac.WarningInterval,
		ac.Creator,
		ac.CreateDate,
	); err != nil {
		return nil, err
	}

	for _, aci := range ac.Instruments {
		if _, err := stmt2.Exec(&alertConfigID, aci.InstrumentID); err != nil {
			return nil, err
		}
	}
	if err := stmt1.Close(); err != nil {
		return nil, err
	}
	if err := stmt2.Close(); err != nil {
		return nil, err
	}
	if err := txn.Commit(); err != nil {
		return nil, err
	}
	return GetAlertConfig(db, &alertConfigID)
}

// UpdateAlertConfig updates an alert
func UpdateAlertConfig(db *sqlx.DB, alertConfigID *uuid.UUID, ac *AlertConfig) (*AlertConfig, error) {
	txn, err := db.Beginx()
	if err != nil {
		return nil, err
	}
	defer txn.Rollback()

	stmt1, err := txn.Preparex(`
		UPDATE alert_config SET
			name=$3,
			body=$4,
			alert_type_id=$5,
			start_date=$6,
			schedule_interval=$7,
			n_missed_before_alert=$8,
			remind_interval=$9,
			warning_interval=$10,
			updater=$11,
			update_date=$12
		WHERE id=$1 AND project_id=$2
	`)
	if err != nil {
		return nil, err
	}

	stmt2, err := txn.Preparex(`
		DELETE FROM alert_config_instrument WHERE alert_config_id=$1
	`)
	if err != nil {
		return nil, err
	}
	stmt3, err := txn.Preparex(`
		INSERT INTO alert_config_instrument (alert_config_id, instrument_id) VALUES ($1, $2)
	`)
	if err != nil {
		return nil, err
	}

	if _, err := stmt1.Exec(
		alertConfigID,
		ac.ProjectID,
		ac.Name,
		ac.Body,
		ac.AlertTypeID,
		ac.StartDate,
		ac.ScheduleInterval,
		ac.NMissedBeforeAlert,
		ac.RemindInterval,
		ac.WarningInterval,
		ac.Updater,
		ac.UpdateDate,
	); err != nil {
		return nil, err
	}
	if _, err := stmt2.Exec(alertConfigID); err != nil {
		return nil, err
	}
	for _, aci := range ac.Instruments {
		if _, err := stmt3.Exec(alertConfigID, aci.InstrumentID); err != nil {
			return nil, err
		}
	}

	if err := stmt1.Close(); err != nil {
		return nil, err
	}
	if err := stmt2.Close(); err != nil {
		return nil, err
	}
	if err := stmt3.Close(); err != nil {
		return nil, err
	}
	if err := txn.Commit(); err != nil {
		return nil, err
	}
	return GetAlertConfig(db, alertConfigID)
}

// DeleteAlertConfig deletes an alert by ID
func DeleteAlertConfig(db *sqlx.DB, alertConfigID *uuid.UUID) error {
	_, err := db.Exec(
		`DELETE FROM alert_config WHERE id = $1`, alertConfigID,
	)
	if err != nil {
		return err
	}
	return nil
}
