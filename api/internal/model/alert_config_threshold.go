package model

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

type AlertConfigThreshold struct {
	AlertConfig
	Opts AlertConfigThresholdOpts `json:"opts" db:"opts"`
}

type AlertConfigThresholdOpts struct {
	AlertLowValue   *float64 `json:"alert_low_value" db:"alert_low_value"`
	AlertHighValue  *float64 `json:"alert_high_value" db:"alert_high_value"`
	WarnLowValue    *float64 `json:"warn_low_value" db:"warn_low_value"`
	WarnHighValue   *float64 `json:"warn_high_value" db:"warn_high_value"`
	IgnoreLowValue  *float64 `json:"ignore_low_value" db:"ignore_low_value"`
	IgnoreHighValue *float64 `json:"ignore_high_value" db:"ignore_high_value"`
	Variance        float64  `json:"variance" db:"variance"`
}

func (o *AlertConfigThresholdOpts) Scan(src interface{}) error {
	b, ok := src.(string)
	if !ok {
		return fmt.Errorf("type assertion failed")
	}
	return json.Unmarshal([]byte(b), o)
}

const createAlertConfigThreshold = `
	INSERT INTO alert_config_threshold
	(alert_config_id, alert_low_value, alert_high_value, warn_low_value, warn_high_value, ignore_low_value, ignore_high_value, variance) VALUES
	($1,$2,$3,$4,$5,$6,$7,$8)
`

func (q *Queries) CreateAlertConfigThreshold(ctx context.Context, alertConfigID uuid.UUID, opts AlertConfigThresholdOpts) error {
	_, err := q.db.ExecContext(ctx, createAlertConfigThreshold, alertConfigID,
		opts.AlertLowValue, opts.AlertHighValue,
		opts.WarnLowValue, opts.WarnHighValue,
		opts.IgnoreLowValue, opts.IgnoreHighValue, opts.Variance,
	)
	return err
}

const updateAlertConfigThreshold = `
	UPDATE alert_config_threshold SET
		alert_low_value=$2,
		alert_high_value=$3,
		warn_low_value=$4,
		warn_high_value=$5,
		ignore_low_value=$6,
		ignore_high_value=$7,
		variance=$8
	WHERE alert_config_id=$1
`

func (q *Queries) UpdateAlertConfigThreshold(ctx context.Context, alertConfigID uuid.UUID, opts AlertConfigThresholdOpts) error {
	_, err := q.db.ExecContext(ctx, updateAlertConfigThreshold, alertConfigID,
		opts.AlertLowValue, opts.AlertHighValue,
		opts.WarnLowValue, opts.WarnHighValue,
		opts.IgnoreLowValue, opts.IgnoreHighValue, opts.Variance,
	)
	return err
}
