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
	WarnRateOfChange   *float64 `json:"warn_rate_of_change" db:"warn_rate_of_change"`
	AlertRateOfChange  float64  `json:"alert_rate_of_change" db:"alert_rate_of_change"`
	IgnoreRateOfChange *float64 `json:"ignore_rate_of_change" db:"ignore_rate_of_change"`
	LocfBackfill       *string  `json:"locf_backfill" db:"locf_backfill"`
	LocfBackfillMs     *int64   `json:"-" db:"locf_backfill_ms"`
}

func (o *AlertConfigChangeOpts) Scan(src interface{}) error {
	b, ok := src.(string)
	if !ok {
		return fmt.Errorf("type assertion failed")
	}
	return json.Unmarshal([]byte(b), o)
}

type TimeseriesMeasurementChange struct {
	BeforeMeasurement Measurement
	AfterMeasurement  Measurement
	Change            float64
}

const createAlertConfigChange = `
	INSERT INTO alert_config_change (alert_config_id, warn_rate_of_change, alert_rate_of_change, ignore_rate_of_change) VALUES ($1,$2,$3,$4)
`

func (q *Queries) CreateAlertConfigChange(ctx context.Context, alertConfigID uuid.UUID, opts AlertConfigChangeOpts) error {
	_, err := q.db.ExecContext(ctx, createAlertConfigChange, alertConfigID, opts.WarnRateOfChange, opts.AlertRateOfChange, opts.IgnoreRateOfChange)
	return err
}

const updateAlertConfigChange = `
	UPDATE alert_config_change SET
		warn_rate_of_change=$2,
		alert_rate_of_change=$3,
		ignore_rate_of_change=$4
	WHERE alert_config_id=$1
`

func (q *Queries) UpdateAlertConfigChange(ctx context.Context, alertConfigID uuid.UUID, opts AlertConfigChangeOpts) error {
	_, err := q.db.ExecContext(ctx, updateAlertConfigChange, alertConfigID, opts.WarnRateOfChange, opts.AlertRateOfChange, opts.IgnoreRateOfChange)
	return err
}
