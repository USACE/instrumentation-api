package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
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
	NMissedBeforeAlert      int                               `json:"n_missed_before_alert" db:"n_missed_before_alert"`
	RemindInterval          string                            `json:"remind_interval" db:"remind_interval"`
	WarningInterval         string                            `json:"warning_interval" db:"warning_interval"`
	LastChecked             *time.Time                        `json:"last_checked" db:"last_checked"`
	LastReminded            *time.Time                        `json:"last_reminded" db:"last_reminded"`
	Instruments             AlertConfigInstrumentCollection   `json:"instruments" db:"instruments"`
	AlertEmailSubscriptions EmailAutocompleteResultCollection `json:"alert_email_subscriptions" db:"alert_email_subscriptions"`
	CreatorUsername         string                            `json:"creator_username" db:"creator_username"`
	UpdaterUsername         *string                           `json:"updater_username" db:"updater_username"`
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

// ListProjectAlertConfigs lists all alert configs for a single project
func ListProjectAlertConfigs(db *sqlx.DB, projectID *uuid.UUID) ([]AlertConfig, error) {
	aa := make([]AlertConfig, 0)
	sql := `
		SELECT *
		FROM v_alert_config
		WHERE project_id = $1
		ORDER BY name
	`
	if err := db.Select(&aa, sql, projectID); err != nil {
		return aa, err
	}

	return aa, nil
}

// ListProjectAlertConfigsByAlertType lists alert configs for a single project filetered by alert type
func ListProjectAlertConfigsByAlertType(db *sqlx.DB, projectID, alertTypeID *uuid.UUID) ([]AlertConfig, error) {
	aa := make([]AlertConfig, 0)
	sql := `
		SELECT *
		FROM v_alert_config
		WHERE project_id = $1
		AND alert_type_id = $2
		ORDER BY name
	`
	if err := db.Select(&aa, sql, projectID, alertTypeID); err != nil {
		return aa, err
	}

	return aa, nil
}

// ListInstrumentAlertConfigs lists all alerts for a single instrument
func ListInstrumentAlertConfigs(db *sqlx.DB, instrumentID *uuid.UUID) ([]AlertConfig, error) {
	aa := make([]AlertConfig, 0)
	sql := `
		SELECT *
		FROM v_alert_config
		WHERE id = ANY(
			SELECT alert_config_id
			FROM alert_config_instrument
			WHERE instrument_id = $1
		)
		ORDER BY name
	`
	if err := db.Select(&aa, sql, instrumentID); err != nil {
		return aa, err
	}

	return aa, nil
}

// GetAlertConfig gets a single alert
func GetAlertConfig(db *sqlx.DB, alertConfigID *uuid.UUID) (*AlertConfig, error) {
	var a AlertConfig
	sql := `SELECT * FROM v_alert_config WHERE id = $1`
	if err := db.Get(&a, sql, alertConfigID); err != nil {
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

	if ac.RemindInterval == "" {
		ac.RemindInterval = "PT0"
	}
	if ac.WarningInterval == "" {
		ac.WarningInterval = "PT0"
	}

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
	stmt3, err := txn.Preparex(`
		INSERT INTO submittal (alert_config_id, due_date)
		SELECT id, create_date + (schedule_interval * n_missed_before_alert)
		FROM alert_config
		WHERE alert_config_id=$1
		LIMIT 1
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
	if err := SubscribeEmailsToAlertConfigTxn(txn, &alertConfigID, ac.AlertEmailSubscriptions); err != nil {
		return nil, err
	}
	if _, err := stmt3.Exec(&alertConfigID); err != nil {
		return nil, err
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
	return GetAlertConfig(db, &alertConfigID)
}

// UpdateAlertConfig updates an alert
func UpdateAlertConfig(db *sqlx.DB, alertConfigID *uuid.UUID, ac *AlertConfig) (*AlertConfig, error) {
	txn, err := db.Beginx()
	if err != nil {
		return nil, err
	}
	defer txn.Rollback()

	if ac.RemindInterval == "" {
		ac.RemindInterval = "PT0"
	}
	if ac.WarningInterval == "" {
		ac.WarningInterval = "PT0"
	}

	stmt1, err := txn.Preparex(`
		UPDATE alert_config SET
			name=$3,
			body=$4,
			start_date=$5,
			schedule_interval=$6,
			n_missed_before_alert=$7,
			remind_interval=$8,
			warning_interval=$9,
			updater=$10,
			update_date=$11
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
	// delete future alert check
	stmt4, err := txn.Preparex(`
		DELETE FROM submittal
		WHERE alert_config_id=$1
		AND due_date > NOW()
	`)
	if err != nil {
		return nil, err
	}
	// update alert check with new interval
	stmt5, err := txn.Preparex(`
		INSERT INTO submittal (alert_config_id, due_date)
		SELECT ac.id, COALESCE(MAX(acs.create_date), ac.create_date) + (ac.schedule_interval * ac.n_missed_before_alert)
		FROM alert_config ac
		INNER JOIN submittal acs ON ac.id = acs.alert_config_id
		WHERE ac.id=$1
		LIMIT 1
	`)
	if err != nil {
		return nil, err
	}

	if _, err := stmt1.Exec(
		alertConfigID,
		ac.ProjectID,
		ac.Name,
		ac.Body,
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

	if err := UnsubscribeAllFromAlertConfigTxn(txn, alertConfigID); err != nil {
		return nil, err
	}
	if err := SubscribeEmailsToAlertConfigTxn(txn, alertConfigID, ac.AlertEmailSubscriptions); err != nil {
		return nil, err
	}

	if _, err := stmt4.Exec(alertConfigID); err != nil {
		return nil, err
	}
	if _, err := stmt5.Exec(alertConfigID); err != nil {
		return nil, err
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
	if err := stmt4.Close(); err != nil {
		return nil, err
	}
	if err := stmt5.Close(); err != nil {
		return nil, err
	}
	if err := txn.Commit(); err != nil {
		return nil, err
	}
	return GetAlertConfig(db, alertConfigID)
}

// DeleteAlertConfig deletes an alert by ID
func DeleteAlertConfig(db *sqlx.DB, alertConfigID *uuid.UUID) error {
	if _, err := db.Exec(`UPDATE alert_config SET deleted=true WHERE id=$1`, alertConfigID); err != nil {
		return err
	}
	return nil
}
