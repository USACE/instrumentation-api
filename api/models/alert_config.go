package models

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// AlertConfig is an alert configuration for an instrument
type AlertConfig struct {
	ID           uuid.UUID `json:"id"`
	InstrumentID uuid.UUID `json:"instrument_id" db:"instrument_id"`
	Name         string    `json:"name"`
	Body         string    `json:"body"`
	Formula      string    `json:"formula"`
	Schedule     string    `json:"schedule"`
	AuditInfo
}

// AlertConfigCollection holds one ore more alert items
type AlertConfigCollection struct {
	Items []AlertConfig `json:"items"`
}

// UnmarshalJSON implements the UnmarshalJSON Interface
func (c *AlertConfigCollection) UnmarshalJSON(b []byte) error {

	switch JSONType(b) {
	case "ARRAY":
		if err := json.Unmarshal(b, &c.Items); err != nil {
			return err
		}
	case "OBJECT":
		var a AlertConfig
		if err := json.Unmarshal(b, &a); err != nil {
			return err
		}
		c.Items = []AlertConfig{a}
	default:
		c.Items = make([]AlertConfig, 0)
	}
	return nil
}

// ListInstrumentAlertConfigs lists all alerts for a single instrument
func ListInstrumentAlertConfigs(db *sqlx.DB, instrumentID *uuid.UUID) ([]AlertConfig, error) {
	var aa []AlertConfig
	sql := `SELECT *
			FROM alert_config
			WHERE instrument_id = $1
	`
	err := db.Select(&aa, sql, instrumentID)
	if err != nil {
		return make([]AlertConfig, 0), err
	}
	return aa, nil
}

// GetAlertConfig gets a single alert
func GetAlertConfig(db *sqlx.DB, alertID *uuid.UUID) (*AlertConfig, error) {
	var a AlertConfig
	sql := `SELECT *
			FROM alert_config
			WHERE id = $1
	`
	err := db.Get(&a, sql, alertID)
	if err != nil {
		return nil, err
	}
	return &a, nil
}

// CreateInstrumentAlertConfigs creates one or more new alert configurations
func CreateInstrumentAlertConfigs(db *sqlx.DB, instrumentID *uuid.UUID, alertConfigs []AlertConfig) ([]AlertConfig, error) {

	txn, err := db.Beginx()
	if err != nil {
		return make([]AlertConfig, 0), err
	}
	defer txn.Rollback()

	// Instrument
	stmt1, err := txn.Preparex(
		`INSERT INTO alert_config
			(instrument_id, name, body, formula, schedule, creator, create_date)
		VALUES
			 ($1, $2, $3, $4, $5, $6, $7)
		RETURNING *`,
	)
	if err != nil {
		return make([]AlertConfig, 0), err
	}

	newAlertConfigs := make([]AlertConfig, len(alertConfigs))
	for idx, c := range alertConfigs {
		var aCreated AlertConfig
		// Load Instrument
		if err := stmt1.Get(&aCreated, instrumentID, c.Name, c.Body, c.Formula, c.Schedule, c.Creator, c.CreateDate); err != nil {
			return make([]AlertConfig, 0), err
		}
		newAlertConfigs[idx] = aCreated
	}
	if err := stmt1.Close(); err != nil {
		return make([]AlertConfig, 0), err
	}
	if err := txn.Commit(); err != nil {
		return make([]AlertConfig, 0), err
	}
	return newAlertConfigs, nil
}

// UpdateInstrumentAlertConfig updates an alert
func UpdateInstrumentAlertConfig(db *sqlx.DB, instrumentID *uuid.UUID, alertConfigID *uuid.UUID, ac *AlertConfig) (*AlertConfig, error) {

	var cUpdated AlertConfig
	err := db.QueryRowx(
		`UPDATE alert_config SET name=$3, body=$4, formula=$5, schedule=$6, updater=$7, update_date=$8
		WHERE id=$1 AND instrument_id=$2
		RETURNING *`,
		alertConfigID, instrumentID, ac.Name, ac.Body, ac.Formula, ac.Schedule, ac.Updater, ac.UpdateDate,
	).StructScan(&cUpdated)
	if err != nil {
		return nil, err
	}
	return &cUpdated, nil
}

// DeleteInstrumentAlertConfig deletes an alert by ID
func DeleteInstrumentAlertConfig(db *sqlx.DB, alertConfigID *uuid.UUID, instrumentID *uuid.UUID) error {
	_, err := db.Exec(
		`DELETE FROM alert_config WHERE id = $1 AND instrument_id=$2`, alertConfigID, instrumentID,
	)
	if err != nil {
		return err
	}
	return nil
}
