package models

import (
	_sql "database/sql"
	"encoding/json"
	"log"
	"time"

	"github.com/USACE/instrumentation-api/api/config"
	et "github.com/USACE/instrumentation-api/api/email_template"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type AlertCheck struct {
	AlertConfigID     uuid.UUID   `db:"alert_config_id"`
	ShouldWarn        bool        `db:"should_warn"`
	ShouldAlert       bool        `db:"should_alert"`
	ShouldRemind      bool        `db:"should_remind"`
	ExpectedSubmittal time.Time   `db:"expected_submittal"`
	AlertConfig       AlertConfig `db:"-"`
}

type EvaluationSubmittal struct {
	LastEvaluationTime *time.Time `db:"last_evaluation_time"`
	AlertCheck
}

type MeasurementSubmittal struct {
	AffectedInstruments MeasurementSubmittalInstrumentCollection `db:"affected_instruments"`
	AlertCheck
}

type MeasurementSubmittalInstrument struct {
	InstrumentName      string    `json:"instrument_name"`
	LastMeasurementTime time.Time `json:"last_measurement_time"`
}

type MeasurementSubmittalInstrumentCollection []MeasurementSubmittalInstrument

func (a *MeasurementSubmittalInstrumentCollection) Scan(src interface{}) error {
	if err := json.Unmarshal([]byte(src.(string)), a); err != nil {
		return err
	}
	return nil
}

type AlertChecker interface {
	GetAlertConfig() AlertConfig
	GetShouldWarn() bool
	GetShouldAlert() bool
	GetShouldRemind() bool
	DoEmail(string, *config.AlertCheckConfig, *config.SmtpConfig) error
}

func (ck AlertCheck) GetAlertConfig() AlertConfig {
	return ck.AlertConfig
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

func (es EvaluationSubmittal) DoEmail(emailType string, cfg *config.AlertCheckConfig, smtpCfg *config.SmtpConfig) error {
	preformatted := et.EmailContent{
		TextSubject: "-- DO NOT REPLY -- MIDAS " + emailType + ": {{.AlertConfig.ProjectName}} Evaluation Submittal \"{{.AlertConfig.Name}}\"",
		TextBody: "The following " + emailType + " has been triggered:\r\n\r\n" +
			"Project: {{.AlertConfig.ProjectName}}\r\n" +
			"Alert Type: Evaluation Submittal\r\n" +
			"Alert Name: \"{{.AlertConfig.Name}}\"\r\n" +
			"Description: \"{{.AlertConfig.Body}}\"\r\n" +
			"Expected Evaluation Submittal Time: {{.ExpectedSubmittal.Format \"Jan 02, 2006 15:04:05 UTC\" }}\r\n" +
			"{{if .LastEvaluationTime}}Last Evaluation Submittal Time: {{.LastEvaluationTime.Format \"Jan 02, 2006 15:04:05 UTC\" }}{{end}}\r\n",
	}
	templContent, err := et.CreateEmailTemplateContent(preformatted)
	if err != nil {
		return err
	}
	content, err := et.FormatAlertConfigTemplates(templContent, es)
	if err != nil {
		return err
	}
	content.To = es.AlertConfig.GetToAddresses()

	if err := et.ConstructAndSendEmail(content, cfg, smtpCfg); err != nil {
		return err
	}
	return nil
}

func (ms MeasurementSubmittal) DoEmail(emailType string, cfg *config.AlertCheckConfig, smtpCfg *config.SmtpConfig) error {
	preformatted := et.EmailContent{
		TextSubject: "-- DO NOT REPLY -- MIDAS " + emailType + ": {{.AlertConfig.ProjectName}} Timeseries Measurement Submittal \"{{.AlertConfig.Name}}\"",
		TextBody: "The following " + emailType + " has been triggered:\r\n\r\n" +
			"Project: {{.AlertConfig.ProjectName}}\r\n" +
			"Alert Type: Measurement Submittal\r\n" +
			"Alert Name: \"{{.AlertConfig.Name}}\"\r\n" +
			"Description: \"{{.AlertConfig.Body}}\"\r\n" +
			"Expected Measurement Submittal Time: {{.ExpectedSubmittal.Format \"Jan 02, 2006 15:04:05 UTC\" }}\r\n" +
			"Affected Instruments Last Measurement Time:\r\n" +
			"{{range .AffectedInstruments}}\t• {{.InstrumentName}}: {{.LastMeasurementTime.Format \"Jan 02, 2006 15:04:05 UTC\" }}\r\n{{end}}",
	}
	templContent, err := et.CreateEmailTemplateContent(preformatted)
	if err != nil {
		return err
	}
	content, err := et.FormatAlertConfigTemplates(templContent, ms)
	if err != nil {
		return err
	}
	content.To = ms.AlertConfig.GetToAddresses()

	if err := et.ConstructAndSendEmail(content, cfg, smtpCfg); err != nil {
		return err
	}
	return nil
}

func ListExpiredAlertConfigs(db *sqlx.DB) ([]AlertConfig, error) {
	aa := make([]AlertConfig, 0)

	sql := `
		UPDATE alert_config ac1
		SET last_checked = now()
		FROM  (
			SELECT *
			FROM v_alert_config a
			WHERE (
				COALESCE(
					last_checked,
					start_date
				) <= now() - schedule_interval::INTERVAL + warning_interval::INTERVAL
			)
			OR alert_status_id = '84a0f437-a20a-4ac2-8a5b-f8dc35e8489b'::UUID
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

func UpdateAlertConfigStatus(db *sqlx.DB, alertConfigs []AlertConfig) error {
	txn, err := db.Beginx()
	if err != nil {
		return err
	}
	defer txn.Rollback()

	stmt, err := txn.Preparex(`
		UPDATE alert_config SET
			last_reminded=$2,
			alert_status_id=$3
		WHERE id=$1
	`)
	if err != nil {
		return err
	}

	for _, ac := range alertConfigs {
		if _, err := stmt.Exec(ac.ID, ac.LastReminded, ac.AlertStatusID); err != nil {
			if err == _sql.ErrNoRows {
				log.Println(err.Error())
				return nil
			}
			return err
		}
	}
	if err := stmt.Close(); err != nil {
		return err
	}
	if err := txn.Commit(); err != nil {
		return err
	}

	return nil
}

func ListAlertCheckEvaluationSubmittals(db *sqlx.DB, alertConfigs []AlertConfig) ([]EvaluationSubmittal, error) {
	es := make([]EvaluationSubmittal, 0)
	if len(alertConfigs) == 0 {
		log.Println("no evaluation submittals to check")
		return es, nil
	}

	acIDs := make([]uuid.UUID, len(alertConfigs))
	for idx := range alertConfigs {
		acIDs[idx] = alertConfigs[idx].ID
	}

	query, args, err := sqlx.In(`SELECT * FROM v_alert_check_evaluation_submittal WHERE alert_config_id IN (?)`, acIDs)
	if err != nil {
		return es, err
	}
	query = db.Rebind(query)
	if err := db.Select(&es, query, args...); err != nil {
		if err == _sql.ErrNoRows {
			return es, nil
		}
		return es, err
	}

	acMap := make(map[uuid.UUID]AlertConfig)
	for _, a := range alertConfigs {
		acMap[a.ID] = a
	}
	for idx := range es {
		ac, ok := acMap[es[idx].AlertConfigID]
		if !ok {
			return es, err
		}
		es[idx].AlertConfig = ac
	}

	return es, nil
}

func ListAlertCheckMeasurementSubmittals(db *sqlx.DB, alertConfigs []AlertConfig) ([]MeasurementSubmittal, error) {
	ms := make([]MeasurementSubmittal, 0)
	if len(alertConfigs) == 0 {
		log.Println("no measurement submittals to check")
		return ms, nil
	}

	acIDs := make([]uuid.UUID, len(alertConfigs))
	for idx := range alertConfigs {
		acIDs[idx] = alertConfigs[idx].ID
	}

	query, args, err := sqlx.In(`SELECT * FROM v_alert_check_measurement_submittal WHERE alert_config_id IN (?)`, acIDs)
	if err != nil {
		return ms, err
	}

	query = db.Rebind(query)
	if err := db.Select(&ms, query, args...); err != nil {
		if err == _sql.ErrNoRows {
			return ms, nil
		}
		return ms, err
	}

	acMap := make(map[uuid.UUID]AlertConfig)
	for _, a := range alertConfigs {
		acMap[a.ID] = a
	}
	for idx := range ms {
		ac, ok := acMap[ms[idx].AlertConfigID]
		if !ok {
			return ms, err
		}
		ms[idx].AlertConfig = ac
	}

	return ms, nil
}
