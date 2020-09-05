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

// CreateAlerts creates one or more new alerts
func CreateAlerts(db *sqlx.DB, action *Action, alerts []Alert) ([]Alert, error) {
	txn, err := db.Begin()
	if err != nil {
		return nil, err
	}

	// Instrument
	stmt1, err := txn.Prepare(
		`INSERT INTO alert
			(id, instrument_id, name, body, formula, schedule)
		VALUES
		 	($1, $2, $3, $4, $5, $6)`,
	)
	if err != nil {
		return nil, err
	}

	for _, a := range alerts {
		// Load Instrument
		if _, err := stmt1.Exec(
			a.ID, a.InstrumentID, a.Name, a.Body, a.Formula, a.Schedule,
		); err != nil {
			return nil, err
		}
	}
	if err := stmt1.Close(); err != nil {
		return nil, err
	}
	if err := txn.Commit(); err != nil {
		return nil, err
	}

	return nil, nil
}

// UpdateAlert

// SubscribeAlertProfile

// UnsubscribeAlertProfile

// SubscribeAlertEmail

// UnsubscribeAlertEmail
