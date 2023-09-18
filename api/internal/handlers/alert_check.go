package handlers

import (
	"errors"
	"log"

	"github.com/USACE/instrumentation-api/api/internal/config"
	"github.com/USACE/instrumentation-api/api/internal/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
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

	errs := make([]error, 0)

	if err := models.CheckMeasurements(db, subMap, acMap, cfg, smtpCfg); err != nil {
		errs = append(errs, err)
	}
	if err := models.CheckEvaluations(db, subMap, acMap, cfg, smtpCfg); err != nil {
		errs = append(errs, err)
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil
}
