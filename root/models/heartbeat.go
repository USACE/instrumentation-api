package models

import (
	"time"

	"github.com/jmoiron/sqlx"
)

// Heartbeat is a timestamp
type Heartbeat struct {
	Time time.Time `json:"time"`
}

// DoHeartbeat does regular-interval tasks
func DoHeartbeat(db *sqlx.DB) (*Heartbeat, error) {
	var h Heartbeat
	err := db.Get(&h, "INSERT INTO heartbeat (time) VALUES ($1) RETURNING *", time.Now().In(time.UTC))
	if err != nil {
		return nil, err
	}
	return &h, nil
}

// GetLatestHeartbeat returns the most recent system heartbeat
func GetLatestHeartbeat(db *sqlx.DB) (*Heartbeat, error) {
	var h Heartbeat
	err := db.Get(&h, "SELECT MAX(time) AS time FROM heartbeat")
	if err != nil {
		return nil, err
	}
	return &h, nil
}
