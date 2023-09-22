package model

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

var (
	GreenSubmittalStatusID  uuid.UUID = uuid.MustParse("0c0d6487-3f71-4121-8575-19514c7b9f03")
	YellowSubmittalStatusID uuid.UUID = uuid.MustParse("ef9a3235-f6e2-4e6c-92f6-760684308f7f")
	RedSubmittalStatusID    uuid.UUID = uuid.MustParse("84a0f437-a20a-4ac2-8a5b-f8dc35e8489b")

	MeasurementSubmittalAlertTypeID uuid.UUID = uuid.MustParse("97e7a25c-d5c7-4ded-b272-1bb6e5914fe3")
	EvaluationSubmittalAlertTypeID  uuid.UUID = uuid.MustParse("da6ee89e-58cc-4d85-8384-43c3c33a68bd")
)

const (
	warning  = "Warning"
	alert    = "Alert"
	reminder = "Reminder"
)

type AlertCheck struct {
	AlertConfigID uuid.UUID `db:"alert_config_id"`
	SubmittalID   uuid.UUID `db:"submittal_id"`
	ShouldWarn    bool      `db:"should_warn"`
	ShouldAlert   bool      `db:"should_alert"`
	ShouldRemind  bool      `db:"should_remind"`
	Submittal     Submittal `db:"-"`
}

func (ck AlertCheck) GetShouldWarn() bool {
	return ck.ShouldWarn
}

func (ck AlertCheck) GetShouldAlert() bool {
	return ck.ShouldAlert
}

func (ck AlertCheck) GetShouldRemind() bool {
	return ck.ShouldRemind
}

func (ck AlertCheck) GetSubmittal() Submittal {
	return ck.Submittal
}

func (ck *AlertCheck) SetSubmittal(sub Submittal) {
	ck.Submittal = sub
}

type AlertConfigMap map[uuid.UUID]AlertConfig

type SubmittalMap map[uuid.UUID]Submittal

func (q *Queries) ListAndCheckAlertConfigs(ctx context.Context) ([]AlertConfig, error) {
	aa := make([]AlertConfig, 0)
	if err := q.db.SelectContext(ctx, &aa, `
		UPDATE alert_config ac1
		SET last_checked = now()
		FROM (
			SELECT *
			FROM v_alert_config
		) ac2
		WHERE  ac1.id = ac2.id
		RETURNING ac2.*
	`); err != nil {
		if err == sql.ErrNoRows {
			return aa, nil
		}
		return aa, err
	}
	return aa, nil
}

func (q *Queries) UpdateAlertConfigLastReminded(ctx context.Context, ac *AlertConfig) error {
	if _, err := q.db.ExecContext(ctx, `
		UPDATE alert_config SET
			last_reminded = $2
		WHERE id = $1
	`, ac.ID, ac.LastReminded); err != nil {
		return err
	}
	return nil
}

func (q *Queries) UpdateSubmittalCompletionDateOrWarningSent(ctx context.Context, sub *Submittal) error {
	if _, err := q.db.ExecContext(ctx, `
		UPDATE submittal SET
			submittal_status_id = $2,
			completion_date = $3,
			warning_sent = $4
		WHERE id = $1
	`, sub.ID, sub.SubmittalStatusID, sub.CompletionDate, sub.WarningSent); err != nil {
		return err
	}
	return nil
}

func (q *Queries) CreateNextSubmittalFromNewAlertConfigDate(ctx context.Context, ac *AlertConfig) error {
	if _, err := q.db.ExecContext(ctx, `
		INSERT INTO submittal (alert_config_id, create_date, due_date)
		SELECT
			ac.id,
			$2::TIMESTAMPTZ,
			$2::TIMESTAMPTZ + ac.schedule_interval
		FROM alert_config ac
		WHERE ac.id = $1
	`, ac.ID, ac.CreateNextSubmittalFrom); err != nil {
		return err
	}
	return nil
}
