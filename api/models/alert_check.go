package models

import (
	_sql "database/sql"

	"github.com/USACE/instrumentation-api/api/config"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type AlertCheck struct {
	AlertConfigID uuid.UUID `db:"alert_config_id"`
	SubmittalID   uuid.UUID `db:"submittal_id"`
	ShouldWarn    bool      `db:"should_warn"`
	ShouldAlert   bool      `db:"should_alert"`
	ShouldRemind  bool      `db:"should_remind"`
	Submittal     Submittal `db:"-"`
}

type AlertConfigChecker[T AlertChecker] interface {
	GetAlertConfig() AlertConfig
	SetAlertConfig(AlertConfig)
	GetChecks() []T
	SetChecks([]T)
	DoEmail(string, *config.AlertCheckConfig, *config.SmtpConfig) error
}

type AlertChecker interface {
	GetShouldWarn() bool
	GetShouldAlert() bool
	GetShouldRemind() bool
	GetSubmittal() Submittal
	SetSubmittal(Submittal)
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

func ListAndCheckAlertConfigs(db *sqlx.DB) ([]AlertConfig, error) {
	aa := make([]AlertConfig, 0)

	sql := `
		UPDATE alert_config ac1
		SET last_checked = now()
		FROM (
			SELECT *
			FROM v_alert_config
		) ac2
		WHERE  ac1.id = ac2.id
		RETURNING ac2.*
	`

	if err := db.Select(&aa, sql); err != nil {
		if err == _sql.ErrNoRows {
			return aa, nil
		}
		return aa, err
	}

	return aa, nil
}

func UpdateAlertConfigChecks[T AlertChecker, PT AlertConfigChecker[T]](db *sqlx.DB, accs []PT) error {
	txn, err := db.Beginx()
	if err != nil {
		return err
	}
	defer txn.Rollback()

	stmt1, err := txn.Preparex(`
		UPDATE alert_config SET
			last_reminded = $2
		WHERE id = $1
	`)
	if err != nil {
		return err
	}

	stmt2, err := txn.Preparex(`
		UPDATE submittal SET
			submittal_status_id = $2,
			completion_date = $3,
			warning_sent = $4
		WHERE id = $1
	`)
	if err != nil {
		return err
	}

	stmt3, err := txn.Preparex(`
		INSERT INTO submittal (alert_config_id, create_date, due_date)
		SELECT
			ac.id,
			$2::TIMESTAMPTZ,
			$2::TIMESTAMPTZ + ac.schedule_interval
		FROM alert_config ac
		WHERE ac.id = $1
	`)
	if err != nil {
		return err
	}

	for _, acc := range accs {
		ac := acc.GetAlertConfig()
		if _, err := stmt1.Exec(ac.ID, ac.LastReminded); err != nil {
			return err
		}
		checks := acc.GetChecks()
		for _, c := range checks {
			sub := c.GetSubmittal()
			if _, err := stmt2.Exec(sub.ID, sub.SubmittalStatusID, sub.CompletionDate, sub.WarningSent); err != nil {
				return err
			}
		}
		if ac.CreateNextSubmittalFrom != nil {
			if _, err := stmt3.Exec(ac.ID, ac.CreateNextSubmittalFrom); err != nil {
				return err
			}
		}
	}

	if err := stmt1.Close(); err != nil {
		return err
	}
	if err := stmt2.Close(); err != nil {
		return err
	}
	if err := stmt3.Close(); err != nil {
		return err
	}
	if err := txn.Commit(); err != nil {
		return err
	}

	return nil
}
