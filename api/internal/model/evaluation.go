package model

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Evaluation struct {
	ID              uuid.UUID                      `json:"id" db:"id"`
	ProjectID       uuid.UUID                      `json:"project_id" db:"project_id"`
	ProjectName     string                         `json:"project_name" db:"project_name"`
	AlertConfigID   *uuid.UUID                     `json:"alert_config_id" db:"alert_config_id"`
	AlertConfigName *string                        `json:"alert_config_name" db:"alert_config_name"`
	SubmittalID     *uuid.UUID                     `json:"submittal_id" db:"submittal_id"`
	Name            string                         `json:"name" db:"name"`
	Body            string                         `json:"body" db:"body"`
	StartDate       time.Time                      `json:"start_date" db:"start_date"`
	EndDate         time.Time                      `json:"end_date" db:"end_date"`
	Instruments     EvaluationInstrumentCollection `json:"instruments" db:"instruments"`
	CreatorUsername string                         `json:"creator_username" db:"creator_username"`
	UpdaterUsername *string                        `json:"updater_username" db:"updater_username"`
	AuditInfo
}

type EvaluationInstrument struct {
	InstrumentID   uuid.UUID `json:"instrument_id" db:"instrument_id"`
	InstrumentName string    `json:"instrument_name" db:"instrument_name"`
}

type EvaluationInstrumentCollection []EvaluationInstrument

func (a *EvaluationInstrumentCollection) Scan(src interface{}) error {
	if err := json.Unmarshal([]byte(src.(string)), a); err != nil {
		return err
	}
	return nil
}

const listProjectEvaluations = `
	SELECT *
	FROM v_evaluation
	WHERE project_id = $1
`

func (q *Queries) ListProjectEvaluations(ctx context.Context, projectID uuid.UUID) ([]Evaluation, error) {
	var aa []Evaluation
	if err := q.db.SelectContext(ctx, &aa, listProjectEvaluations, projectID); err != nil {
		return nil, err
	}
	return aa, nil
}

const listProjectEvaluationsByAlertConfig = `
	SELECT * FROM v_evaluation
	WHERE project_id = $1
	AND alert_config_id IS NOT NULL
	AND alert_config_id = $2
`

func (q *Queries) ListProjectEvaluationsByAlertConfig(ctx context.Context, projectID, alertConfigID uuid.UUID) ([]Evaluation, error) {
	var aa []Evaluation
	err := q.db.SelectContext(ctx, &aa, listProjectEvaluationsByAlertConfig, projectID, alertConfigID)
	if err != nil {
		return make([]Evaluation, 0), err
	}
	return aa, nil
}

const listInstrumentEvaluations = `
	SELECT * FROM v_evaluation
	WHERE id = ANY(
		SELECT evaluation_id
		FROM evaluation_instrument
		WHERE instrument_id = $1
	)
`

func (q *Queries) ListInstrumentEvaluations(ctx context.Context, instrumentID uuid.UUID) ([]Evaluation, error) {
	aa := make([]Evaluation, 0)
	if err := q.db.SelectContext(ctx, &aa, listInstrumentEvaluations, instrumentID); err != nil {
		return nil, err
	}
	return aa, nil
}

const getEvaluation = `
	SELECT * FROM v_evaluation WHERE id = $1
`

func (q *Queries) GetEvaluation(ctx context.Context, evaluationID uuid.UUID) (Evaluation, error) {
	var a Evaluation
	if err := q.db.GetContext(ctx, &a, getEvaluation, evaluationID); err != nil {
		return a, err
	}
	return a, nil
}

const completeEvaluationSubmittal = `
	UPDATE submittal sub1 SET
		submittal_status_id = sq.submittal_status_id,
		completion_date = NOW()
	FROM (
		SELECT
			sub2.id AS submittal_id,
			CASE
				-- if completed before due date, mark submittal as green id
				WHEN NOW() <= sub2.due_date THEN '0c0d6487-3f71-4121-8575-19514c7b9f03'::UUID
				-- if completed after due date, mark as yellow
				ELSE 'ef9a3235-f6e2-4e6c-92f6-760684308f7f'::UUID
			END AS submittal_status_id
		FROM submittal sub2
		INNER JOIN alert_config ac ON sub2.alert_config_id = ac.id
		WHERE sub2.id = $1
		AND sub2.completion_date IS NULL
		AND NOT sub2.marked_as_missing
		AND ac.alert_type_id = 'da6ee89e-58cc-4d85-8384-43c3c33a68bd'::UUID
	) sq
	WHERE sub1.id = sq.submittal_id
	RETURNING sub1.*
`

func (q *Queries) CompleteEvaluationSubmittal(ctx context.Context, submittalID uuid.UUID) (Submittal, error) {
	var sub Submittal
	if err := q.db.GetContext(ctx, &sub, completeEvaluationSubmittal, submittalID); err != nil {
		if err == sql.ErrNoRows {
			return sub, fmt.Errorf("submittal must exist, be of evaluation type, and before due date or unvalidated missing")
		}
		return sub, err
	}
	return sub, nil
}

const createNextEvaluationSubmittal = `
	INSERT INTO submittal (alert_config_id, due_date)
	SELECT
		ac.id,
		NOW() + ac.schedule_interval
	FROM alert_config ac
	WHERE ac.id IN (SELECT alert_config_id FROM submittal WHERE id = $1)
`

func (q *Queries) CreateNextEvaluationSubmittal(ctx context.Context, submittalID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, createNextEvaluationSubmittal, submittalID)
	return err
}

const createEvaluation = `
	INSERT INTO evaluation (
		project_id,
		submittal_id,
		name,
		body,
		start_date,
		end_date,
		creator,
		create_date
	) VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
	RETURNING id
`

func (q *Queries) CreateEvaluation(ctx context.Context, ev Evaluation) (uuid.UUID, error) {
	var evaluationID uuid.UUID
	err := q.db.GetContext(
		ctx,
		&evaluationID,
		createEvaluation,
		ev.ProjectID,
		ev.SubmittalID,
		ev.Name,
		ev.Body,
		ev.StartDate,
		ev.EndDate,
		ev.Creator,
		ev.CreateDate,
	)
	return evaluationID, err
}

const createEvalationInstrument = `
	INSERT INTO evaluation_instrument (evaluation_id, instrument_id) VALUES ($1,$2)
`

func (q *Queries) CreateEvaluationInstrument(ctx context.Context, evaluationID, instrumentID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, createEvalationInstrument, evaluationID, instrumentID)
	return err
}

const updateEvaluation = `
	UPDATE evaluation SET
		name=$3,
		body=$4,
		start_date=$5,
		end_date=$6,
		updater=$7,
		update_date=$8
	WHERE id=$1 AND project_id=$2
`

func (q *Queries) UpdateEvaluation(ctx context.Context, ev Evaluation) error {
	_, err := q.db.ExecContext(
		ctx,
		updateEvaluation,
		ev.ID,
		ev.ProjectID,
		ev.Name,
		ev.Body,
		ev.StartDate,
		ev.EndDate,
		ev.Updater,
		ev.UpdateDate,
	)
	return err
}

const unassignAllInstrumentsFromEvaluation = `
	DELETE FROM evaluation_instrument WHERE evaluation_id = $1
`

func (q *Queries) UnassignAllInstrumentsFromEvaluation(ctx context.Context, evaluationID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, unassignAllInstrumentsFromEvaluation, evaluationID)
	return err
}

const deleteEvaluation = `
	DELETE FROM evaluation WHERE id = $1
`

func (q *Queries) DeleteEvaluation(ctx context.Context, evaluationID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteEvaluation, evaluationID)
	return err
}
