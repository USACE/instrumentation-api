package model

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

type AlertConfigChange struct {
	AlertConfig
	Opts AlertConfigChangeOpts `json:"opts" db:"opts"`
}

type AlertConfigChangeOpts struct {
	RateOfChange float64 `json:"rate_of_change" db:"rate_of_change"`
}

func (o *AlertConfigChangeOpts) Scan(src interface{}) error {
	b, ok := src.(string)
	if !ok {
		return fmt.Errorf("type assertion failed")
	}
	return json.Unmarshal([]byte(b), o)
}

const createAlertConfigChange = `
	INSERT INTO alert_config_change (alert_config_id, rate_of_change) VALUES ($1,$2)
`

func (q *Queries) CreateAlertConfigChange(ctx context.Context, alertConfigID uuid.UUID, opts AlertConfigChangeOpts) error {
	_, err := q.db.ExecContext(ctx, createAlertConfigChange, alertConfigID, opts.RateOfChange)
	return err
}

const updateAlertConfigChange = `
	UPDATE alert_config_change SET
		rate_of_change=$2
	WHERE alert_config_id=$1
`

func (q *Queries) UpdateAlertConfigChange(ctx context.Context, alertConfigID uuid.UUID, opts AlertConfigChangeOpts) error {
	_, err := q.db.ExecContext(ctx, updateAlertConfigChange, alertConfigID, opts.RateOfChange)
	return err
}
