package model

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
)

type AlertConfigScheduler struct {
	AlertConfig
	Opts AlertConfigSchedulerOpts `json:"opts" db:"opts"`
}

type AlertConfigSchedulerOpts struct {
	StartDate               time.Time  `json:"start_date" db:"start_date"`
	ScheduleInterval        string     `json:"schedule_interval" db:"schedule_interval"`
	RemindInterval          string     `json:"remind_interval" db:"remind_interval"`
	WarningInterval         string     `json:"warning_interval" db:"warning_interval"`
	LastReminded            *time.Time `json:"last_reminded" db:"last_reminded"`
	MuteConsecutiveAlerts   bool       `json:"mute_consecutive_alerts" db:"mute_consecutive_alerts"`
	CreateNextSubmittalFrom *time.Time `json:"-" db:"-"`
}

func (o *AlertConfigSchedulerOpts) Scan(src interface{}) error {
	b, ok := src.(string)
	if !ok {
		return fmt.Errorf("type assertion failed")
	}
	return json.Unmarshal([]byte(b), o)
}

const createAlertConfigScheduler = `
	INSERT INTO alert_config_scheduler (
		alert_config_id,
		start_date,
		schedule_interval,
		mute_consecutive_alerts,
		remind_interval,
		warning_interval
	) VALUES ($1,$2,$3,$4,$5,$6)
`

func (q *Queries) CreateAlertConfigScheduler(ctx context.Context, alertConfigID uuid.UUID, ac AlertConfigSchedulerOpts) error {
	_, err := q.db.ExecContext(ctx, createAlertConfigScheduler,
		alertConfigID,
		ac.StartDate,
		ac.ScheduleInterval,
		ac.MuteConsecutiveAlerts,
		ac.RemindInterval,
		ac.WarningInterval,
	)
	return err
}

const updateAlertConfigScheduler = `
	UPDATE alert_config_scheduler SET
		start_date = $2,
		schedule_interval = $3,
		mute_consecutive_alerts = $4,
		remind_interval = $5,
		warning_interval = $6
	WHERE alert_config_id = $1
`

func (q *Queries) UpdateAlertConfigScheduler(ctx context.Context, alertConfigID uuid.UUID, ac AlertConfigSchedulerOpts) error {
	_, err := q.db.ExecContext(ctx, updateAlertConfigScheduler,
		alertConfigID,
		ac.StartDate,
		ac.ScheduleInterval,
		ac.MuteConsecutiveAlerts,
		ac.RemindInterval,
		ac.WarningInterval,
	)
	return err
}

const listAndCheckAlertConfigSchedulers = `
	UPDATE alert_config ac1
	SET last_checked = now()
	FROM (
		SELECT *
		FROM v_alert_config
		WHERE alert_type_id = '97e7a25c-d5c7-4ded-b272-1bb6e5914fe3'::uuid
		OR alert_type_id = 'da6ee89e-58cc-4d85-8384-43c3c33a68bd'::uuid
	) ac2
	WHERE  ac1.id = ac2.id
	RETURNING ac2.*
`

func (q *Queries) ListAndCheckAlertConfigSchedulers(ctx context.Context) ([]AlertConfigScheduler, error) {
	aa := make([]AlertConfigScheduler, 0)
	if err := q.db.SelectContext(ctx, &aa, listAndCheckAlertConfigSchedulers); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return make([]AlertConfigScheduler, 0), nil
		}
		log.Print("this is the error")
		return nil, err
	}
	return aa, nil
}

const updateAlertConfigLastReminded = `
	UPDATE alert_config_scheduler SET
		last_reminded = $2
	WHERE alert_config_id = $1
`

func (q *Queries) UpdateAlertConfigLastReminded(ctx context.Context, alertConfigID uuid.UUID, lastReminded *time.Time) error {
	_, err := q.db.ExecContext(ctx, updateAlertConfigLastReminded, alertConfigID, lastReminded)
	return err
}
