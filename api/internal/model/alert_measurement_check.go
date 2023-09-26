package model

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/USACE/instrumentation-api/api/internal/config"
	et "github.com/USACE/instrumentation-api/api/internal/email_template"
)

type AlertConfigMeasurementCheck struct {
	AlertConfig AlertConfig
	AlertChecks []*MeasurementCheck
}

type MeasurementCheck struct {
	AffectedTimeseries MeasurementSubmittalTimeseriesCollection `db:"affected_timeseries"`
	AlertCheck
}

type MeasurementSubmittalTimeseries struct {
	InstrumentName string `json:"instrument_name"`
	TimeseriesName string `json:"timeseries_name"`
	Status         string `json:"status"`
}

type MeasurementSubmittalTimeseriesCollection []MeasurementSubmittalTimeseries

func (a *MeasurementSubmittalTimeseriesCollection) Scan(src interface{}) error {
	if err := json.Unmarshal([]byte(src.(string)), a); err != nil {
		return err
	}
	return nil
}

func (a AlertConfigMeasurementCheck) GetAlertConfig() AlertConfig {
	return a.AlertConfig
}

func (a *AlertConfigMeasurementCheck) SetAlertConfig(ac AlertConfig) {
	a.AlertConfig = ac
}

func (a AlertConfigMeasurementCheck) GetChecks() []*MeasurementCheck {
	return a.AlertChecks
}

func (a *AlertConfigMeasurementCheck) SetChecks(mc []*MeasurementCheck) {
	a.AlertChecks = mc
}

func (ms AlertConfigMeasurementCheck) DoEmail(emailType string, cfg *config.AlertCheckConfig, smtpCfg *config.SmtpConfig) error {
	if emailType == "" {
		return fmt.Errorf("must provide emailType")
	}
	preformatted := et.EmailContent{
		TextSubject: "-- DO NOT REPLY -- MIDAS " + emailType + ": Timeseries Measurement Submittal",
		TextBody: "The following " + emailType + " has been triggered:\r\n\r\n" +
			"Project: {{.AlertConfig.ProjectName}}\r\n" +
			"Alert Type: Measurement Submittal\r\n" +
			"Alert Name: \"{{.AlertConfig.Name}}\"\r\n" +
			"Description: \"{{.AlertConfig.Body}}\"\r\n" +
			"Expected Measurement Submittals:\r\n" +
			"{{range .AlertChecks}}" +
			"\t• {{.Submittal.CreateDate.Format \"Jan 02 2006 15:04:05 UTC\"}} - {{.Submittal.DueDate.Format \"Jan 02 2006 15:04:05 UTC\"}}\r\n" +
			"{{range .AffectedTimeseries}}" +
			"\t\t• {{.InstrumentName}}: {{.TimeseriesName}} ({{.Status}})\r\n" +
			"{{end}}\r\n{{end}}",
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

const getAllIncompleteMeasurementSubmittals = `
	SELECT * FROM v_alert_check_measurement_submittal
	WHERE submittal_id = ANY(
		SELECT id FROM submittal
		WHERE completion_date IS NULL AND NOT marked_as_missing
	)
`

func (q *Queries) GetAllIncompleteMeasurementSubmittals(ctx context.Context) ([]*MeasurementCheck, error) {
	mcs := make([]*MeasurementCheck, 0)
	if err := q.db.SelectContext(ctx, &mcs, getAllIncompleteMeasurementSubmittals); err != nil {
		return nil, err
	}
	return mcs, nil
}
