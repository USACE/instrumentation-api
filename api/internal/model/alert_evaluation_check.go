package model

import (
	"context"
	"fmt"

	"github.com/USACE/instrumentation-api/api/internal/config"
	et "github.com/USACE/instrumentation-api/api/internal/email_template"
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
		return fmt.Errorf("must provide emailType")
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

func (q *Queries) GetAllIncompleteEvaluationSubmittals(ctx context.Context) ([]*EvaluationCheck, error) {
	ecs := make([]*EvaluationCheck, 0)
	if err := q.db.Select(&ecs, `
		SELECT * FROM v_alert_check_evaluation_submittal
		WHERE submittal_id = ANY(
			SELECT id FROM submittal
			WHERE completion_date IS NULL AND NOT marked_as_missing
		)
	`); err != nil {
		return ecs, err
	}
	return ecs, nil
}
