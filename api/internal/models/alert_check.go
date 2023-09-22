package models

import (
	_sql "database/sql"
	"errors"
	"sync"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/config"
	"github.com/USACE/instrumentation-api/api/internal/utils"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
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

func UpdateAlertConfigChecks[T AlertChecker, PT AlertConfigChecker[T]](txn *sqlx.Tx, accs []PT) error {
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

	return nil
}

// there should always be at least one "missing" submittal within an alert config. Submittals are created:
//  1. when an alert config is created (first submittal)
//  2. when a submittal is completed (next submittal created)
//  3. when a submittals due date has passed if it is not completed
//
// for evaluations, the next is submittal created manually when the evaluation is made
// for measurements, the next submittal is created the first time this function runs after the due date
//
// No "Yellow" Status Submittals should be passed to this function as it implies the submittal has been completed
//
// TODO: smtp.SendMail esablishes a new connection for each batch of emails sent. I would be better to aggregate
// the contents of each email, then create a connection pool to reuse and send all emails at once, with any errors wrapped and returned
func HandleChecks[T AlertChecker, PT AlertConfigChecker[T]](txn *sqlx.Tx, accs []PT, cfg *config.AlertCheckConfig, smtpCfg *config.SmtpConfig) error {
	defer utils.Timer()()

	mu := &sync.Mutex{}
	aaccs := make([]PT, len(accs))
	errs := make([]error, 0)
	t := time.Now()

	wg := sync.WaitGroup{}
	for i, p := range accs {
		wg.Add(1)
		go func(idx int, acc PT) {
			defer wg.Done()

			ac := acc.GetAlertConfig()
			checks := acc.GetChecks()

			// If ANY "missing" submittals are within an alert config, aggregate missing submittals and send an alert
			acAlert := false
			sendAlertEmail := false
			// If ANY missing submittals previously existed within an alert config, send them in a "reminder" instead of an alert
			acReminder := false
			sendReminderEmail := false
			// If a reminder exists when at least one submittal "shouldAlert", the alert should be aggregated into the next reminder
			// instead of sending a new reminder email. If NO alerts exist for an alert config, the reminder can be reset to NULL.
			// Reminders should be set when the first alert for an alert config is triggered, or at each reminder interval
			resetReminders := len(checks) != 0

			for j, c := range checks {
				shouldWarn := c.GetShouldWarn()
				shouldAlert := c.GetShouldAlert()
				shouldRemind := c.GetShouldRemind()
				sub := c.GetSubmittal()

				// if no submittal alerts or warnings are found, no emails should be sent
				if !shouldAlert && !shouldWarn {
					// if submittal status was previously red, update status to yellow and
					// completion_date to current timestamp
					if sub.SubmittalStatusID == RedSubmittalStatusID {
						sub.SubmittalStatusID = YellowSubmittalStatusID
						sub.CompletionDate = &t
						ac.CreateNextSubmittalFrom = &t
					} else

					// if submittal status is green and the current time is not before the submittal due date,
					// complete the submittal at that due date and prepare the next submittal interval
					if sub.SubmittalStatusID == GreenSubmittalStatusID && !t.Before(sub.DueDate) {
						sub.CompletionDate = &sub.DueDate
						ac.CreateNextSubmittalFrom = &sub.DueDate
					}
				} else

				// if any submittal warning is triggered, immediately send a
				// warning email, since submittal due dates are unique within alert configs
				if shouldWarn && !sub.WarningSent {
					if !ac.MuteConsecutiveAlerts || ac.LastReminded == nil {
						mu.Lock()
						if err := acc.DoEmail(warning, cfg, smtpCfg); err != nil {
							errs = append(errs, err)
						}
						mu.Unlock()
					}
					sub.SubmittalStatusID = GreenSubmittalStatusID
					sub.WarningSent = true
				} else

				// if any submittal alert is triggered after a warning has been sent within an
				// alert config, aggregate missing submittals and send their contents in an alert email
				if shouldAlert {
					if sub.SubmittalStatusID != RedSubmittalStatusID {
						sub.SubmittalStatusID = RedSubmittalStatusID
						acAlert = true
						ac.CreateNextSubmittalFrom = &sub.DueDate
					}
					resetReminders = false
				}

				// if any reminder is triggered, aggregate missing
				// submittals and send their contents in an email
				if shouldRemind {
					acReminder = true
				}

				c.SetSubmittal(sub)
				checks[j] = c
			}

			// if there are no alerts, there should also be no reminders sent. "last_reminded" is used to determine
			// if an alert has already been sent for an alert config, and send a reminder if so
			if resetReminders {
				ac.LastReminded = nil
			}

			// if there are any reminders within an alert config, they will override the alerts if MuteConsecutiveAlerts is true
			if acAlert && ((!acReminder && ac.LastReminded == nil) || !ac.MuteConsecutiveAlerts) {
				ac.LastReminded = &t
				sendAlertEmail = true
			}
			if acReminder && ac.LastReminded != nil {
				ac.LastReminded = &t
				sendReminderEmail = true
			}

			acc.SetAlertConfig(ac)
			acc.SetChecks(checks)

			if sendAlertEmail {
				mu.Lock()
				if err := acc.DoEmail(alert, cfg, smtpCfg); err != nil {
					errs = append(errs, err)
				}
				mu.Unlock()
			}
			if sendReminderEmail {
				mu.Lock()
				if err := acc.DoEmail(reminder, cfg, smtpCfg); err != nil {
					errs = append(errs, err)
				}
				mu.Unlock()
			}

			aaccs[idx] = acc
		}(i, p)
	}
	wg.Wait()

	if err := UpdateAlertConfigChecks[T, PT](txn, aaccs); err != nil {
		errs = append(errs, err)
		return errors.Join(errs...)
	}
	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil
}
