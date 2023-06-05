package handlers

import (
	"log"
	"time"

	"github.com/USACE/instrumentation-api/api/models"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
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

func DoAlertChecks(db *sqlx.DB, svc *ses.SES, sender string) echo.HandlerFunc {
	return func(c echo.Context) error {
		txn, err := db.Beginx()
		if err != nil {
			return err
		}
		defer txn.Rollback()

		aa, err := models.ListAndRenewExpiredAlertConfigsTxn(txn)
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

		measurementChecks, err := models.ListAlertCheckMeasurementSubmittalsTxn(txn, aa)
		if err != nil {
			return err
		}
		evaluationChecks, err := models.ListAlertCheckEvaluationSubmittalsTxn(txn, aa)
		if err != nil {
			return err
		}
		if err := handleChecks(txn, svc, measurementChecks, aa, sender); err != nil {
			return err
		}
		if err := handleChecks(txn, svc, evaluationChecks, aa, sender); err != nil {
			return err
		}
		if err := txn.Commit(); err != nil {
			return err
		}

		return nil
	}
}

func handleChecks[T models.AlertChecker](txn *sqlx.Tx, svc *ses.SES, checks []T, alertConfigs []models.AlertConfig, sender string) error {
	check := func(err error) {
		if err != nil {
			log.Println(err.Error())
		}
	}
	acIDs := make([]uuid.UUID, 0)
	aa := make([]models.AlertConfig, len(checks))

	for idx, c := range checks {
		ac := c.GetAlertConfig()
		shouldWarn := c.GetShouldWarn()
		shouldAlert := c.GetShouldAlert()
		shouldRemind := c.GetShouldRemind()

		switch ac.AlertStatusID {
		case GreenAlertStatusID:
			if shouldWarn && !shouldAlert {
				ac.AlertStatusID = YellowAlertStatusID
				check(c.DoEmail(svc, warning, sender))
				_ = append(acIDs, ac.ID)
			} else if shouldAlert {
				ac.AlertStatusID = RedAlertStatusID
				check(c.DoEmail(svc, alert, sender))
				_ = append(acIDs, ac.ID)
			}
		case YellowAlertStatusID:
			if shouldAlert {
				ac.AlertStatusID = RedAlertStatusID
				check(c.DoEmail(svc, alert, sender))
				_ = append(acIDs, ac.ID)
			} else if !shouldWarn {
				ac.AlertStatusID = GreenAlertStatusID
			}
		case RedAlertStatusID:
			if shouldRemind {
				t := time.Now()
				ac.LastReminded = &t
				check(c.DoEmail(svc, reminder, sender))
				_ = append(acIDs, ac.ID)
			} else if !shouldAlert && shouldWarn {
				// edge case may happen where if an submittal is very late, the next
				// scheduled submittal may go directly into warning or alert status
				ac.AlertStatusID = YellowAlertStatusID
				check(c.DoEmail(svc, warning, sender))
				_ = append(acIDs, ac.ID)
			} else if !shouldAlert && !shouldWarn {
				ac.AlertStatusID = GreenAlertStatusID
			}
		default:
			log.Printf("invalid alert_status_id: %+v", ac.AlertStatusID)
		}
		aa[idx] = ac
	}

	if err := models.UpdateAlertConfigStatusAndLastRemindedTxn(txn, aa); err != nil {
		return err
	}
	if err := models.CreateAlertsTxn(txn, acIDs); err != nil {
		return err
	}

	return nil
}
