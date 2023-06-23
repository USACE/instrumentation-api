package handlers

import (
	"errors"
	"fmt"
	"time"

	"github.com/USACE/instrumentation-api/api/config"
	"github.com/USACE/instrumentation-api/api/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

var (
	GreenAlertStatusID  uuid.UUID = uuid.MustParse("0c0d6487-3f71-4121-8575-19514c7b9f03")
	YellowAlertStatusID uuid.UUID = uuid.MustParse("ef9a3235-f6e2-4e6c-92f6-760684308f7f")
	RedAlertStatusID    uuid.UUID = uuid.MustParse("84a0f437-a20a-4ac2-8a5b-f8dc35e8489b")

	MeasurementSubmittalAlertTypeID uuid.UUID = uuid.MustParse("97e7a25c-d5c7-4ded-b272-1bb6e5914fe3")
	EvaluationSubmittalAlertTypeID  uuid.UUID = uuid.MustParse("da6ee89e-58cc-4d85-8384-43c3c33a68bd")
)

const (
	warning  = "Warning"
	alert    = "Alert"
	reminder = "Reminder"
)

func DoAlertChecks(db *sqlx.DB, cfg *config.AlertCheckConfig, smtpCfg *config.SmtpConfig) error {
	aa, err := models.ListExpiredAlertConfigs(db)
	if err != nil {
		return err
	}

	acMap := make(map[uuid.UUID][]models.AlertConfig)
	for _, a := range aa {
		if _, exists := acMap[a.AlertTypeID]; !exists {
			acMap[a.AlertTypeID] = make([]models.AlertConfig, 0)
		}
		acMap[a.AlertTypeID] = append(acMap[a.AlertTypeID], a)
	}

	measurementChecks, err := models.ListAlertCheckMeasurementSubmittals(db, aa)
	if err != nil {
		return err
	}
	evaluationChecks, err := models.ListAlertCheckEvaluationSubmittals(db, aa)
	if err != nil {
		return err
	}

	errs := make([]error, 0)
	if err := handleChecks(db, measurementChecks, aa, cfg, smtpCfg); err != nil {
		errs = append(errs, err)
	}
	if err := handleChecks(db, evaluationChecks, aa, cfg, smtpCfg); err != nil {
		errs = append(errs, err)
	}
	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil
}

// TODO: smtp.SendMail esablishes a new connection for each batch of emails sent. I would be better to aggregate
// the contents of each email, then create a connection pool to reuse and send all emails at once, with any errors wrapped and returned
func handleChecks[T models.AlertChecker](db *sqlx.DB, checks []T, alertConfigs []models.AlertConfig, cfg *config.AlertCheckConfig, smtpCfg *config.SmtpConfig) error {
	acIDs := make([]uuid.UUID, 0)
	aa := make([]models.AlertConfig, len(checks))
	errs := make([]error, 0)

	for idx, c := range checks {
		ac := c.GetAlertConfig()
		shouldWarn := c.GetShouldWarn()
		shouldAlert := c.GetShouldAlert()
		shouldRemind := c.GetShouldRemind()

		switch ac.AlertStatusID {
		case GreenAlertStatusID:
			if shouldWarn && !shouldAlert {
				if err := c.DoEmail(warning, cfg, smtpCfg); err != nil {
					errs = append(errs, err) // aggregate errors
				}
				ac.AlertStatusID = YellowAlertStatusID // update alert config status
				acIDs = append(acIDs, ac.ID)           // add for in-app notification
			} else if shouldAlert {
				if err := c.DoEmail(alert, cfg, smtpCfg); err != nil {
					errs = append(errs, err)
				}
				ac.AlertStatusID = RedAlertStatusID
				t := time.Now()
				ac.LastReminded = &t
				acIDs = append(acIDs, ac.ID)
			}
		case YellowAlertStatusID:
			if shouldAlert {
				if err := c.DoEmail(alert, cfg, smtpCfg); err != nil {
					errs = append(errs, err)
				}
				ac.AlertStatusID = RedAlertStatusID
				t := time.Now()
				ac.LastReminded = &t
				acIDs = append(acIDs, ac.ID)
			} else if !shouldWarn {
				ac.AlertStatusID = GreenAlertStatusID
			}
		case RedAlertStatusID:
			if shouldRemind {
				if err := c.DoEmail(reminder, cfg, smtpCfg); err != nil {
					errs = append(errs, err)
				}
				t := time.Now()
				ac.LastReminded = &t
				acIDs = append(acIDs, ac.ID)
			} else if !shouldAlert && shouldWarn {
				// edge case may happen where if an submittal is very late, the next
				// scheduled submittal may go directly into warning or alert status
				if err := c.DoEmail(warning, cfg, smtpCfg); err != nil {
					errs = append(errs, err)
				}
				ac.AlertStatusID = YellowAlertStatusID
				acIDs = append(acIDs, ac.ID)
			} else if !shouldAlert && !shouldWarn {
				ac.AlertStatusID = GreenAlertStatusID
			}
		default:
			errs = append(errs, fmt.Errorf("invalid alert_status_id: %+v", ac.AlertStatusID))
		}
		aa[idx] = ac
	}

	if err := models.UpdateAlertConfigStatus(db, aa); err != nil {
		errs = append(errs, err)
		return errors.Join(errs...)
	}
	if len(acIDs) > 0 {
		if err := models.CreateAlerts(db, acIDs); err != nil {
			errs = append(errs, err)
			return errors.Join(errs...)
		}
	}
	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil
}
