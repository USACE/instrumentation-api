package model

import (
	"context"
	"time"
)

// Heartbeat is a timestamp
type Heartbeat struct {
	Time time.Time `json:"time"`
}

const doHeartbeat = `
	INSERT INTO heartbeat (time) VALUES ($1) RETURNING *
`

// DoHeartbeat does regular-interval tasks
func (q *Queries) DoHeartbeat(ctx context.Context) (Heartbeat, error) {
	var h Heartbeat
	if err := q.db.GetContext(ctx, &h, doHeartbeat, time.Now().In(time.UTC)); err != nil {
		return h, err
	}
	return h, nil
}

const getLatestHeartbeat = `
	SELECT MAX(time) AS time FROM heartbeat
`

// GetLatestHeartbeat returns the most recent system heartbeat
func (q *Queries) GetLatestHeartbeat(ctx context.Context) (Heartbeat, error) {
	var h Heartbeat
	if err := q.db.GetContext(ctx, &h, getLatestHeartbeat); err != nil {
		return h, err
	}
	return h, nil
}

const listHeartbeats = `
	SELECT * FROM heartbeat
`

// ListHeartbeats returns all system heartbeats
func (q *Queries) ListHeartbeats(ctx context.Context) ([]Heartbeat, error) {
	var hh []Heartbeat
	if err := q.db.SelectContext(ctx, &hh, listHeartbeats); err != nil {
		return nil, err
	}
	return hh, nil
}
