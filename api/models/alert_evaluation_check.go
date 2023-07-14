package models

import (
	"github.com/USACE/instrumentation-api/api/config"
	et "github.com/USACE/instrumentation-api/api/email_template"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type AlertConfigEvaluationCheck struct {
	AlertConfig AlertConfig
	AlertChecks []*EvaluationCheck
}

type EvaluationCheck struct {
	AlertCheck
}

func (a AlertConfigEvaluationCheck) GetAlertConfig() AlertConfig {
	return a.AlertConfig
}

func (a *AlertConfigEvaluationCheck) SetAlertConfig(ac AlertConfig) {
	a.AlertConfig = ac
}

func (a AlertConfigEvaluationCheck) GetChecks() []*EvaluationCheck {
	return a.AlertChecks
}

func (a *AlertConfigEvaluationCheck) SetChecks(ec []*EvaluationCheck) {
	a.AlertChecks = ec
}

func (acc AlertConfigEvaluationCheck) DoEmail(emailType string, cfg *config.AlertCheckConfig, smtpCfg *config.SmtpConfig) error {
	if emailType == "" {
		return nil
	}

	preformatted := et.EmailContent{
		TextSubject: "-- DO NOT REPLY -- MIDAS " + emailType + ": Evaluation Submittal",
		TextBody: "The following " + emailType + " has been triggered:\r\n\r\n" +
			"Project: {{.AlertConfig.ProjectName}}\r\n" +
			"Alert Type: Evaluation Submittal\r\n" +
			"Alert Name: \"{{.AlertConfig.Name}}\"\r\n" +
			"Description: \"{{.AlertConfig.Body}}\"\r\n" +
			"Expected Evaluation Submittals:\r\n" +
			"{{range .AlertChecks}}{{if or .ShouldAlert .ShouldWarn}}" +
			"\t• {{.Submittal.CreateDate.Format \"Jan 02 2006 15:04:05 UTC\"}} - {{.Submittal.DueDate.Format \"Jan 02 2006 15:04:05 UTC\"}}" +
			"{{if .ShouldAlert}} (missing) {{else if .ShouldWarn}} (warning) {{end}}\r\n{{end}}{{end}}",
	}
	templContent, err := et.CreateEmailTemplateContent(preformatted)
	if err != nil {
		return err
	}
	content, err := et.FormatAlertConfigTemplates(templContent, acc)
	if err != nil {
		return err
	}
	content.To = acc.AlertConfig.GetToAddresses()
	if err := et.ConstructAndSendEmail(content, cfg, smtpCfg); err != nil {
		return err
	}
	return nil
}

func ListEvaluationChecks(db *sqlx.DB, acMap map[uuid.UUID]AlertConfig, subMap map[uuid.UUID]Submittal) ([]*AlertConfigEvaluationCheck, error) {
	ecs := make([]*EvaluationCheck, 0)
	accs := make([]*AlertConfigEvaluationCheck, 0)

	if err := db.Select(&ecs, `
		SELECT * FROM v_alert_check_evaluation_submittal
		WHERE submittal_id = ANY(
			SELECT id FROM submittal
			WHERE completion_date IS NULL AND NOT marked_as_missing
		)
	`); err != nil {
		return accs, err
	}

	ecMap := make(map[uuid.UUID][]*EvaluationCheck)
	for k := range acMap {
		ecMap[k] = make([]*EvaluationCheck, 0)
	}

	for idx := range ecs {
		if sub, ok := subMap[ecs[idx].SubmittalID]; ok {
			ecs[idx].Submittal = sub
			ecMap[ecs[idx].AlertConfigID] = append(ecMap[ecs[idx].AlertConfigID], ecs[idx])
		}
	}

	for k, v := range acMap {
		acc := AlertConfigEvaluationCheck{
			AlertConfig: v,
			AlertChecks: ecMap[k],
		}
		accs = append(accs, &acc)
	}

	return accs, nil
}
