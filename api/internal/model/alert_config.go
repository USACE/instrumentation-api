package model

import (
	"context"
	"fmt"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/config"
	et "github.com/USACE/instrumentation-api/api/internal/email"
	"github.com/google/uuid"
)

type AlertConfig struct {
	ID                      uuid.UUID                            `json:"id" db:"id"`
	Name                    string                               `json:"name" db:"name"`
	Body                    string                               `json:"body" db:"body"`
	ProjectID               uuid.UUID                            `json:"project_id" db:"project_id"`
	ProjectName             string                               `json:"project_name" db:"project_name"`
	AlertEmailSubscriptions dbJSONSlice[EmailAutocompleteResult] `json:"alert_email_subscriptions" db:"alert_email_subscriptions"`
	Instruments             dbJSONSlice[AlertConfigInstrument]   `json:"instruments" db:"instruments"`
	Timeseries              dbJSONSlice[AlertConfigTimeseries]   `json:"timeseries" db:"timeseries"`
	LastChecked             *time.Time                           `json:"last_checked" db:"last_checked"`
	AlertTypeID             uuid.UUID                            `json:"alert_type_id" db:"alert_type_id"`
	AlertType               string                               `json:"alert_type" db:"alert_type"`
	Opts                    Opts                                 `json:"opts" db:"opts"`
	Violations              []string                             `json:"-" db:"-"`
	AuditInfo
}

type AlertConfigInstrument struct {
	InstrumentID   uuid.UUID `json:"instrument_id" db:"instrument_id"`
	InstrumentName string    `json:"instrument_name" db:"instrument_name"`
}

type AlertConfigTimeseries struct {
	TimeseriesID   uuid.UUID `json:"timeseries_id" db:"timeseries_id"`
	TimeseriesName string    `json:"timeseries_name" db:"timeseries_name"`
}

type TimeseriesAlertConfig struct {
	TimeseriesID         uuid.UUID  `json:"timeseries_id" db:"timeseries_id"`
	LastMeasurementTime  *time.Time `json:"last_measurement_time" db:"last_measurement_time"`
	LastMeasurementValue *float64   `json:"last_measurement_value" db:"last_measurement_value"`
	AlertConfig
}

func (ac AlertConfig) DoEmail(emailType string, cfg config.EmailConfig) error {
	if emailType == "" {
		return fmt.Errorf("must provide emailType")
	}
	preformatted := et.EmailContent{
		TextSubject: "-- DO NOT REPLY -- MIDAS " + emailType + ": " + ac.AlertType,
		TextBody:    "The following " + emailType + " has been triggered:\r\n" + alertEmailTemplate,
	}
	templContent, err := et.CreateEmailTemplateContent(preformatted)
	if err != nil {
		return err
	}
	content, err := et.FormatAlertConfigTemplates(templContent, ac)
	if err != nil {
		return err
	}
	emails := make([]string, len(ac.AlertEmailSubscriptions))
	for idx := range ac.AlertEmailSubscriptions {
		emails[idx] = ac.AlertEmailSubscriptions[idx].Email
	}
	content.To = emails
	if err := et.ConstructAndSendEmail(content, cfg); err != nil {
		return err
	}
	return nil
}

const alertEmailTemplate = "Project: {{.ProjectName}}\r\n" +
	"Alert Type: {{.AlertType}}\r\n" +
	"Alert Name: \"{{.Name}}\"\r\n" +
	"Description: \"{{.Body}}\"\r\n" +
	"Voilations ({{len .Violations}} total):\r\n" +
	"{{range $i, $val := .Violations}}\r\n" +
	"{{if le $i 5}}\t• {{$val}}\r\n{{end}}" +
	"{{if eq $i 6}}\t  more...\r\n{{end}}{{end}}"

func (a *AlertConfig) GetToAddresses() []string {
	emails := make([]string, len(a.AlertEmailSubscriptions))
	for idx := range a.AlertEmailSubscriptions {
		emails[idx] = a.AlertEmailSubscriptions[idx].Email
	}
	return emails
}

const getAllAlertConfigsForProject = `
	SELECT *
	FROM v_alert_config
	WHERE project_id = $1
	ORDER BY name
`

// GetAllAlertConfigsForProject lists all alert configs for a single project
func (q *Queries) GetAllAlertConfigsForProject(ctx context.Context, projectID uuid.UUID) ([]AlertConfig, error) {
	aa := make([]AlertConfig, 0)
	err := q.db.SelectContext(ctx, &aa, getAllAlertConfigsForProject, projectID)
	return aa, err
}

const getAllAlertConfigsForProjectAndAlertType = `
	SELECT *
	FROM v_alert_config
	WHERE project_id = $1
	AND alert_type_id = $2
	ORDER BY name
`

// GetAllAlertConfigsForProjectAndAlertType lists alert configs for a single project filetered by alert type
func (q *Queries) GetAllAlertConfigsForProjectAndAlertType(ctx context.Context, projectID, alertTypeID uuid.UUID) ([]AlertConfig, error) {
	aa := make([]AlertConfig, 0)
	err := q.db.SelectContext(ctx, &aa, getAllAlertConfigsForProjectAndAlertType, projectID, alertTypeID)
	return aa, err
}

const getAllAlertConfigsForInstrument = `
	SELECT ac.*
	FROM v_alert_config ac
	INNER JOIN alert_config_instrument aci ON aci.alert_config_id = ac.id
	WHERE instrument_id = $1
	ORDER BY name
`

// GetAllAlertConfigsForInstrument lists all alerts for a single instrument
func (q *Queries) GetAllAlertConfigsForInstrument(ctx context.Context, instrumentID uuid.UUID) ([]AlertConfig, error) {
	aa := make([]AlertConfig, 0)
	err := q.db.SelectContext(ctx, &aa, getAllAlertConfigsForInstrument, instrumentID)
	return aa, err
}

const listAlertConfigsForTimeseries = `
	WITH input AS (SELECT * FROM json_to_recordset(?) AS r(timeseries_id uuid, time timestamptz))
	SELECT
		acts.timeseries_id,
		mm.time AS last_measurement_time,
		mm.value AS last_measurement_value,
		ac.*
	FROM v_alert_config ac
	INNER JOIN alert_config_timeseries acts ON acts.alert_config_id = ac.id
	INNER JOIN input ON true
	LEFT JOIN LATERAL (
		SELECT imm.time, imm.value
		FROM timeseries_measurement imm
		WHERE imm.timeseries_id = acts.timeseries_id
		AND imm.time < input.time
		ORDER BY imm.time DESC
		LIMIT 1
	) mm ON true
	WHERE acts.timeseries_id = input.timeseries_id
	AND ac.alert_type_id IN (?)
`

func (q *Queries) GetTimeseriesAlertConfigsForTimeseriesAndAlertTypes(ctx context.Context, rr string, alertTypeIDs []uuid.UUID) ([]TimeseriesAlertConfig, error) {
	query, args, err := sqlIn(listAlertConfigsForTimeseries, rr, alertTypeIDs)
	if err != nil {
		return nil, err
	}
	query = q.db.Rebind(query)
	acc := make([]TimeseriesAlertConfig, 0)
	err = q.db.SelectContext(ctx, &acc, query, args...)
	return acc, err
}

const getOneAlertConfig = `
	SELECT * FROM v_alert_config WHERE id = $1
`

// GetOneAlertConfig gets a single alert
func (q *Queries) GetOneAlertConfig(ctx context.Context, alertConfigID uuid.UUID) (AlertConfig, error) {
	var a AlertConfig
	err := q.db.GetContext(ctx, &a, getOneAlertConfig, alertConfigID)
	return a, err
}

const createAlertConfig = `
	INSERT INTO alert_config (
		project_id,
		name,
		body,
		alert_type_id,
		creator,
		create_date
	) VALUES ($1,$2,$3,$4,$5,$6)
	RETURNING id
`

func (q *Queries) CreateAlertConfig(ctx context.Context, ac AlertConfig) (uuid.UUID, error) {
	var alertConfigID uuid.UUID
	err := q.db.GetContext(ctx, &alertConfigID, createAlertConfig,
		ac.ProjectID,
		ac.Name,
		ac.Body,
		ac.AlertTypeID,
		ac.CreatorID,
		ac.CreateDate,
	)
	return alertConfigID, err
}

const assignInstrumentToAlertConfig = `
	INSERT INTO alert_config_instrument (alert_config_id, instrument_id) VALUES ($1, $2)
`

func (q *Queries) AssignInstrumentToAlertConfig(ctx context.Context, alertConfigID, instrumentID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, assignInstrumentToAlertConfig, alertConfigID, instrumentID)
	return err
}

const unassignAllInstrumentsFromAlertConfig = `
	DELETE FROM alert_config_instrument WHERE alert_config_id = $1
`

func (q *Queries) UnassignAllInstrumentsFromAlertConfig(ctx context.Context, alertConfigID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, unassignAllInstrumentsFromAlertConfig, alertConfigID)
	return err
}

const assignTimeseriesToAlertConfig = `
	INSERT INTO alert_config_timeseries (alert_config_id, timeseries_id) VALUES ($1, $2)
`

func (q *Queries) AssignTimeseriesToAlertConfig(ctx context.Context, alertConfigID, timeseriesID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, assignTimeseriesToAlertConfig, alertConfigID, timeseriesID)
	return err
}

const unassignAllTimeseriesFromAlertConfig = `
	DELETE FROM alert_config_timeseries WHERE alert_config_id = $1
`

func (q *Queries) UnassignAllTimeseriesFromAlertConfig(ctx context.Context, alertConfigID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, unassignAllTimeseriesFromAlertConfig, alertConfigID)
	return err
}

const updateAlertConfig = `
	UPDATE alert_config SET
		name = $3,
		body = $4,
		updater = $5,
		update_date = $6
	WHERE id = $1 AND project_id = $2
`

func (q *Queries) UpdateAlertConfig(ctx context.Context, ac AlertConfig) error {
	_, err := q.db.ExecContext(ctx, updateAlertConfig,
		ac.ID,
		ac.ProjectID,
		ac.Name,
		ac.Body,
		ac.UpdaterID,
		ac.UpdateDate,
	)
	return err
}

const updateFutureSubmittalForAlertConfig = `
	UPDATE submittal
	SET due_date = sq.new_due_date
	FROM (
		SELECT
			sub.id AS submittal_id,
			sub.create_date + acs.schedule_interval AS new_due_date
		FROM submittal sub
		INNER JOIN alert_config_scheduler acs ON sub.alert_config_id = acs.alert_config_id
		WHERE sub.alert_config_id = $1
		AND sub.due_date > now()
		AND sub.completion_date IS NULL
		AND NOT sub.marked_as_missing
	) sq
	WHERE id = sq.submittal_id
	AND sq.new_due_date > now()
	RETURNING id
`

func (q *Queries) UpdateFutureSubmittalForAlertConfig(ctx context.Context, alertConfigID uuid.UUID) error {
	var updatedSubID uuid.UUID
	return q.db.GetContext(ctx, &updatedSubID, updateFutureSubmittalForAlertConfig, alertConfigID)
}

const deleteAlertConfig = `
	UPDATE alert_config SET deleted=true WHERE id = $1
`

// DeleteAlertConfig deletes an alert by ID
func (q *Queries) DeleteAlertConfig(ctx context.Context, alertConfigID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteAlertConfig, alertConfigID)
	return err
}

const updateAlertConfigStatus = `
	UPDATE alert_config SET status=$2, last_checked=$3 WHERE id=$1
`

func (q *Queries) UpdateAlertConfigStatus(ctx context.Context, alertConfigID uuid.UUID, status *string) error {
	_, err := q.db.ExecContext(ctx, updateAlertConfigStatus, alertConfigID, status, time.Now())
	return err
}
