package models

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// Alert is an alert for an instrument
type Alert struct {
	ID           uuid.UUID `json:"id"`
	InstrumentID uuid.UUID `json:"instrument_id" db:"instrument_id"`
	Name         string    `json:"name"`
	Body         string    `json:"body"`
	Formula      string    `json:"formula"`
	Schedule     string    `json:"schedule"`
	AuditInfo
}

// AlertCollection holds one ore more alert items
type AlertCollection struct {
	Items []Alert `json:"items"`
}

// UnmarshalJSON implements the UnmarshalJSON Interface
func (c *AlertCollection) UnmarshalJSON(b []byte) error {

	switch JSONType(b) {
	case "ARRAY":
		if err := json.Unmarshal(b, &c.Items); err != nil {
			return err
		}
	case "OBJECT":
		var a Alert
		if err := json.Unmarshal(b, &a); err != nil {
			return err
		}
		c.Items = []Alert{a}
	default:
		c.Items = make([]Alert, 0)
	}
	return nil
}

// ListInstrumentAlerts lists all alerts for a single instrument
func ListInstrumentAlerts(db *sqlx.DB, instrumentID *uuid.UUID) ([]Alert, error) {
	var aa []Alert
	sql := `SELECT *
			FROM alert
			WHERE instrument_id = $1
	`
	err := db.Select(&aa, sql, instrumentID)
	if err != nil {
		return make([]Alert, 0), err
	}
	return aa, nil
}

// GetAlert gets a single alert
func GetAlert(db *sqlx.DB, alertID *uuid.UUID) (*Alert, error) {
	var a Alert
	sql := `SELECT *
			FROM alert
			WHERE id = $1
	`
	err := db.Get(&a, sql, alertID)
	if err != nil {
		return nil, err
	}
	return &a, nil
}

// CreateInstrumentAlerts creates one or more new alerts
func CreateInstrumentAlerts(db *sqlx.DB, action *Action, instrumentID *uuid.UUID, alerts []Alert) ([]Alert, error) {

	txn, err := db.Beginx()
	if err != nil {
		return make([]Alert, 0), err
	}

	// Instrument
	stmt1, err := txn.Preparex(
		`INSERT INTO alert
			(instrument_id, name, body, formula, schedule, creator, create_date, updater, update_date)
		VALUES
			 ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING *`,
	)
	if err != nil {
		return make([]Alert, 0), err
	}

	newAlerts := make([]Alert, len(alerts))
	for idx, a := range alerts {
		var aCreated Alert
		// Load Instrument
		if err := stmt1.Get(
			&aCreated,
			instrumentID, a.Name, a.Body, a.Formula, a.Schedule,
			action.Actor, action.Time, action.Actor, action.Time,
		); err != nil {
			return make([]Alert, 0), err
		}
		newAlerts[idx] = aCreated
	}
	if err := stmt1.Close(); err != nil {
		return make([]Alert, 0), err
	}
	if err := txn.Commit(); err != nil {
		return make([]Alert, 0), err
	}

	return newAlerts, nil
}

// UpdateInstrumentAlert updates an alert
func UpdateInstrumentAlert(db *sqlx.DB, action *Action, instrumentID *uuid.UUID, alertID *uuid.UUID, a *Alert) (*Alert, error) {

	var aUpdated Alert
	err := db.QueryRowx(
		`UPDATE alert SET name=$3, body=$4, formula=$5, schedule=$6, updater=$7, update_date=$8
		WHERE id=$1 AND instrument_id=$2
		RETURNING *`,
		alertID, instrumentID, a.Name, a.Body, a.Formula, a.Schedule, action.Actor, action.Time,
	).StructScan(&aUpdated)
	if err != nil {
		return nil, err
	}
	return &aUpdated, nil
}

// DeleteInstrumentAlert deletes an alert by ID
func DeleteInstrumentAlert(db *sqlx.DB, alertID *uuid.UUID, instrumentID *uuid.UUID) error {
	_, err := db.Exec(
		`DELETE FROM alert WHERE id = $1 AND instrument_id=$2`, alertID, instrumentID,
	)
	if err != nil {
		return err
	}
	return nil
}

// SubscribeAlertProfile

// UnsubscribeAlertProfile

// SubscribeAlertEmail

// UnsubscribeAlertEmail
