package handlers

import (
	"errors"
	"log"
	"sync"
	"time"

	"github.com/USACE/instrumentation-api/api/config"
	"github.com/USACE/instrumentation-api/api/dbutils"
	"github.com/USACE/instrumentation-api/api/models"
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

func DoAlertChecks(db *sqlx.DB, cfg *config.AlertCheckConfig, smtpCfg *config.SmtpConfig) error {
	subs, err := models.ListUnverifiedMissingSubmittals(db)
	if err != nil {
		return err
	}
	acs, err := models.ListAndCheckAlertConfigs(db)
	if err != nil {
		return err
	}
	if len(acs) == 0 {
		log.Println("no alert configs to check")
		return nil
	}

	subMap := make(map[uuid.UUID]models.Submittal)
	for _, s := range subs {
		subMap[s.ID] = s
	}
	acMap := make(map[uuid.UUID]models.AlertConfig)
	for _, a := range acs {
		acMap[a.ID] = a
	}

	measurementChecks, err := models.ListMeasurementChecks(db, acMap, subMap)
	if err != nil {
		return err
	}
	evaluationChecks, err := models.ListEvaluationChecks(db, acMap, subMap)
	if err != nil {
		return err
	}

	errs := make([]error, 0)
	if err := handleChecks[*models.MeasurementCheck, *models.AlertConfigMeasurementCheck](db, measurementChecks, cfg, smtpCfg); err != nil {
		errs = append(errs, err)
	}
	if err := handleChecks[*models.EvaluationCheck, *models.AlertConfigEvaluationCheck](db, evaluationChecks, cfg, smtpCfg); err != nil {
		errs = append(errs, err)
	}
	if len(errs) > 0 {
		return errors.Join(errs...)
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
func handleChecks[T models.AlertChecker, PT models.AlertConfigChecker[T]](db *sqlx.DB, accs []PT, cfg *config.AlertCheckConfig, smtpCfg *config.SmtpConfig) error {
	defer dbutils.Timer()()

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
			resetReminders := true

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
					sub.SubmittalStatusID = GreenSubmittalStatusID
					mu.Lock()
					if err := acc.DoEmail(warning, cfg, smtpCfg); err != nil {
						errs = append(errs, err)
					}
					mu.Unlock()
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

			// if there are any reminders within an alert config, they will override the alerts
			if acAlert && ((!acReminder && ac.LastReminded == nil) || !ac.MuteConsecutiveAlerts) {
				ac.LastReminded = &t
				sendAlertEmail = true
			}
			if acReminder {
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

	if err := models.UpdateAlertConfigChecks[T, PT](db, aaccs); err != nil {
		errs = append(errs, err)
		return errors.Join(errs...)
	}
	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil
}
