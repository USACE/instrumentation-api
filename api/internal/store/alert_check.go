package store

import (
	"context"
	"errors"
	"log"
	"sync"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/config"
	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/USACE/instrumentation-api/api/internal/utils"
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

type AlertCheckStore interface {
	CheckEvaluations(ctx context.Context, subMap model.SubmittalMap, acMap model.AlertConfigMap) error
	CheckMeasurements(ctx context.Context, subMap model.SubmittalMap, acMap model.AlertConfigMap) error
}

type alertConfigChecker[T alertChecker] interface {
	GetAlertConfig() model.AlertConfig
	SetAlertConfig(model.AlertConfig)
	GetChecks() []T
	SetChecks([]T)
	DoEmail(string, *config.AlertCheckConfig, *config.SmtpConfig) error
}

type alertChecker interface {
	GetShouldWarn() bool
	GetShouldAlert() bool
	GetShouldRemind() bool
	GetSubmittal() model.Submittal
	SetSubmittal(model.Submittal)
}

type alertCheckStore struct {
	db      *model.Database
	cfg     *config.AlertCheckConfig
	smtpCfg *config.SmtpConfig
}

func NewAlertCheckStore(db *model.Database, cfg *config.AlertCheckConfig, smtpCfg *config.SmtpConfig) *alertCheckStore {
	return &alertCheckStore{db, cfg, smtpCfg}
}

func (s alertCheckStore) CheckEvaluations(ctx context.Context, subMap model.SubmittalMap, acMap model.AlertConfigMap) error {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			log.Print(err.Error())
		}
	}()

	q := model.NewQueries(s.db).WithTx(tx)

	accs := make([]*model.AlertConfigEvaluationCheck, 0)
	ecs, err := q.GetAllIncompleteEvaluationSubmittals(ctx)
	if err != nil {
		return err
	}

	ecMap := make(map[uuid.UUID][]*model.EvaluationCheck)
	for k := range acMap {
		ecMap[k] = make([]*model.EvaluationCheck, 0)
	}
	for idx := range ecs {
		if sub, ok := subMap[ecs[idx].SubmittalID]; ok {
			ecs[idx].Submittal = sub
			ecMap[ecs[idx].AlertConfigID] = append(ecMap[ecs[idx].AlertConfigID], ecs[idx])
		}
	}
	for k, v := range acMap {
		if v.AlertTypeID != EvaluationSubmittalAlertTypeID {
			continue
		}
		acc := model.AlertConfigEvaluationCheck{
			AlertConfig: v,
			AlertChecks: ecMap[k],
		}
		accs = append(accs, &acc)
	}

	// handleChecks should not rollback txn but should bubble up errors after txn committed
	alertCheckErr := handleChecks[*model.EvaluationCheck, *model.AlertConfigEvaluationCheck](ctx, q, accs, s.cfg, s.smtpCfg)

	if err := tx.Commit(); err != nil {
		return err
	}
	if alertCheckErr != nil {
		return alertCheckErr
	}

	return nil
}

func (s alertCheckStore) CheckMeasurements(ctx context.Context, subMap model.SubmittalMap, acMap model.AlertConfigMap) error {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			log.Print(err.Error())
		}
	}()

	q := model.NewQueries(s.db).WithTx(tx)

	accs := make([]*model.AlertConfigMeasurementCheck, 0)

	mcs, err := q.GetAllIncompleteMeasurementSubmittals(ctx)
	if err != nil {
		return err
	}

	mcMap := make(map[uuid.UUID][]*model.MeasurementCheck)
	for k := range acMap {
		mcMap[k] = make([]*model.MeasurementCheck, 0)
	}

	for idx := range mcs {
		if sub, ok := subMap[mcs[idx].SubmittalID]; ok {
			mcs[idx].Submittal = sub
			mcMap[mcs[idx].AlertConfigID] = append(mcMap[mcs[idx].AlertConfigID], mcs[idx])
		}
	}

	for k, v := range acMap {
		if v.AlertTypeID != MeasurementSubmittalAlertTypeID {
			continue
		}
		acc := model.AlertConfigMeasurementCheck{
			AlertConfig: v,
			AlertChecks: mcMap[k],
		}

		accs = append(accs, &acc)
	}

	// HandleChecks should not rollback txn but should bubble up errors after txn committed
	alertCheckErr := handleChecks[*model.MeasurementCheck, *model.AlertConfigMeasurementCheck](ctx, q, accs, s.cfg, s.smtpCfg)

	if err := tx.Commit(); err != nil {
		return err
	}
	if alertCheckErr != nil {
		return alertCheckErr
	}

	return nil
}

func updateAlertConfigChecks[T alertChecker, PT alertConfigChecker[T]](ctx context.Context, q *model.Queries, accs []PT) error {
	for _, acc := range accs {
		ac := acc.GetAlertConfig()
		if err := q.UpdateAlertConfigLastReminded(ctx, &ac); err != nil {
			return err
		}
		checks := acc.GetChecks()
		for _, c := range checks {
			sub := c.GetSubmittal()
			if err := q.UpdateSubmittalCompletionDateOrWarningSent(ctx, &sub); err != nil {
				return err
			}
		}
		if ac.CreateNextSubmittalFrom != nil {
			if err := q.CreateNextSubmittalFromNewAlertConfigDate(ctx, &ac); err != nil {
				return err
			}
		}
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
// p.s. Dear future me/someone else: I'm sorry
func handleChecks[T alertChecker, PT alertConfigChecker[T]](ctx context.Context, q *model.Queries, accs []PT, cfg *config.AlertCheckConfig, smtpCfg *config.SmtpConfig) error {
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

	if err := updateAlertConfigChecks[T, PT](ctx, q, aaccs); err != nil {
		errs = append(errs, err)
		return errors.Join(errs...)
	}
	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil
}
