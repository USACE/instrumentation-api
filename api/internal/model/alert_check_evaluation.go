package model

import (
	"context"
	"fmt"

	"github.com/USACE/instrumentation-api/api/internal/config"
	"github.com/USACE/instrumentation-api/api/internal/email"
)

type AlertConfigSchedulerEvaluationCheck struct {
	AlertConfig AlertConfigScheduler
	AlertChecks []*EvaluationCheck
}

type EvaluationCheck struct {
	AlertCheck
}

func (a AlertConfigSchedulerEvaluationCheck) GetAlertConfigScheduler() AlertConfigScheduler {
	return a.AlertConfig
}

func (a *AlertConfigSchedulerEvaluationCheck) SetAlertConfigScheduler(ac AlertConfigScheduler) {
	a.AlertConfig = ac
}

func (a AlertConfigSchedulerEvaluationCheck) GetChecks() []*EvaluationCheck {
	return a.AlertChecks
}

func (a *AlertConfigSchedulerEvaluationCheck) SetChecks(ec []*EvaluationCheck) {
	a.AlertChecks = ec
}

func (acc AlertConfigSchedulerEvaluationCheck) DoEmail(emailType string, cfg config.AlertCheckSchedulerConfig) error {
	if emailType == "" {
		return fmt.Errorf("must provide emailType")
	}
	preformatted := email.EmailContent{
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
	templContent, err := email.CreateEmailTemplateContent(preformatted)
	if err != nil {
		return err
	}
	content, err := email.FormatAlertConfigTemplates(templContent, acc)
	if err != nil {
		return err
	}
	content.To = acc.AlertConfig.GetToAddresses()
	if err := email.ConstructAndSendEmail(content, cfg.EmailConfig); err != nil {
		return err
	}
	return nil
}

const getAllIncompleteEvaluationSubmittals = `
	SELECT * FROM v_alert_check_evaluation_submittal
	WHERE submittal_id = ANY(
		SELECT id FROM submittal
		WHERE completion_date IS NULL AND NOT marked_as_missing
	)
`

func (q *Queries) GetAllIncompleteEvaluationSubmittals(ctx context.Context) ([]*EvaluationCheck, error) {
	ecs := make([]*EvaluationCheck, 0)
	if err := q.db.SelectContext(ctx, &ecs, getAllIncompleteEvaluationSubmittals); err != nil {
		return nil, err
	}
	return ecs, nil
}
