package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
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

func ListProjectEvaluations(db *sqlx.DB, projectID *uuid.UUID) ([]Evaluation, error) {
	var aa []Evaluation
	sql := `
		SELECT *
		FROM v_evaluation
		WHERE project_id = $1
	`
	err := db.Select(&aa, sql, projectID)
	if err != nil {
		return make([]Evaluation, 0), err
	}

	return aa, nil
}

func ListProjectEvaluationsByAlertConfig(db *sqlx.DB, projectID, alertConfigID *uuid.UUID) ([]Evaluation, error) {
	var aa []Evaluation
	sql := `
		SELECT *
		FROM v_evaluation
		WHERE project_id = $1
		AND alert_config_id IS NOT NULL
		AND alert_config_id = $2
	`
	err := db.Select(&aa, sql, projectID, alertConfigID)
	if err != nil {
		return make([]Evaluation, 0), err
	}

	return aa, nil
}

func ListInstrumentEvaluations(db *sqlx.DB, instrumentID *uuid.UUID) ([]Evaluation, error) {
	aa := make([]Evaluation, 0)
	sql := `
		SELECT *
		FROM v_evaluation
		WHERE id = ANY(
			SELECT evaluation_id
			FROM evaluation_instrument
			WHERE instrument_id = $1
		)
	`
	err := db.Select(&aa, sql, instrumentID)
	if err != nil {
		return make([]Evaluation, 0), err
	}

	return aa, nil
}

func GetEvaluation(db *sqlx.DB, evaluationID *uuid.UUID) (*Evaluation, error) {
	var a Evaluation
	sql := `SELECT * FROM v_evaluation WHERE id = $1`
	err := db.Get(&a, sql, evaluationID)
	if err != nil {
		return nil, err
	}

	return &a, nil
}

func RecordEvaluationSubmittalTxn(txn *sqlx.Tx, ev *Evaluation) error {
	if ev.AlertConfigID != nil {
		stmt, err := txn.Preparex(`
			UPDATE submittal SET
				submittal_status_id = sq.submittal_status_id,
				completion_date = NOW()
			FROM (
				SELECT
					CASE
						-- if completed before due date, mark submittal as green id
						WHEN NOW() <= due_date THEN '0c0d6487-3f71-4121-8575-19514c7b9f03'::UUID
						-- if completed after due date, mark as yellow
						ELSE 'ef9a3235-f6e2-4e6c-92f6-760684308f7f'::UUID
					END
				FROM submittal
				WHERE id = $1
			) sq
			WHERE id = $1
		`)
		if err != nil {
			return err
		}
		if _, err := stmt.Exec(ev.AlertConfigID); err != nil {
			return err
		}
		if err := stmt.Close(); err != nil {
			return err
		}
	}
	return nil
}

func CreateNextSubmittalTxn(txn *sqlx.Tx, ev *Evaluation) error {
	stmt, err := txn.Preparex(`
		INSERT INTO submittal (alert_config_id, due_date)
		SELECT
			ac.id,
			NOW() + ac.schedule_interval
		FROM alert_config ac
		WHERE ac.id = $1
	`)
	if err != nil {
		return err
	}
	if _, err := stmt.Exec(ev.AlertConfigID); err != nil {
		return err
	}
	if err := stmt.Close(); err != nil {
		return err
	}
	return nil
}

func CreateEvaluation(db *sqlx.DB, ev *Evaluation) (*Evaluation, error) {
	txn, err := db.Beginx()
	if err != nil {
		return nil, err
	}
	defer txn.Rollback()

	if err := RecordEvaluationSubmittalTxn(txn, ev); err != nil {
		return nil, err
	}
	if err := CreateNextSubmittalTxn(txn, ev); err != nil {
		return nil, err
	}

	stmt1, err := txn.Preparex(`
		INSERT INTO evaluation
			(
				project_id,
				alert_config_id,
				name,
				body,
				start_date,
				end_date,
				creator,
				create_date
			) VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
			RETURNING id
	`)
	if err != nil {
		return nil, err
	}
	stmt2, err := txn.Preparex(`
		INSERT INTO evaluation_instrument (evaluation_id, instrument_id) VALUES ($1, $2)
	`)
	if err != nil {
		return nil, err
	}

	var evaluationID uuid.UUID
	if err := stmt1.Get(
		&evaluationID,
		ev.ProjectID,
		ev.AlertConfigID,
		ev.Name,
		ev.Body,
		ev.StartDate,
		ev.EndDate,
		ev.Creator,
		ev.CreateDate,
	); err != nil {
		return nil, err
	}

	for _, aci := range ev.Instruments {
		if _, err := stmt2.Exec(&evaluationID, aci.InstrumentID); err != nil {
			return nil, err
		}
	}

	if err := stmt1.Close(); err != nil {
		return nil, err
	}
	if err := stmt2.Close(); err != nil {
		return nil, err
	}
	if err := txn.Commit(); err != nil {
		return nil, err
	}
	return GetEvaluation(db, &evaluationID)
}

func UpdateEvaluation(db *sqlx.DB, evaluationID *uuid.UUID, ev *Evaluation) (*Evaluation, error) {
	txn, err := db.Beginx()
	if err != nil {
		return nil, err
	}
	defer txn.Rollback()

	stmt1, err := txn.Preparex(`
		UPDATE evaluation SET
			name=$3,
			body=$4,
			start_date=$5,
			end_date=$6,
			updater=$7,
			update_date=$8
		WHERE id=$1 AND project_id=$2
	`)
	if err != nil {
		return nil, err
	}

	stmt2, err := txn.Preparex(`
		DELETE FROM evaluation_instrument WHERE evaluation_id=$1
	`)
	if err != nil {
		return nil, err
	}
	stmt3, err := txn.Preparex(`
		INSERT INTO evaluation_instrument (evaluation_id, instrument_id) VALUES ($1, $2)
	`)
	if err != nil {
		return nil, err
	}

	if _, err := stmt1.Exec(
		evaluationID,
		ev.ProjectID,
		ev.Name,
		ev.Body,
		ev.StartDate,
		ev.EndDate,
		ev.Updater,
		ev.UpdateDate,
	); err != nil {
		return nil, err
	}
	if _, err := stmt2.Exec(evaluationID); err != nil {
		return nil, err
	}
	for _, aci := range ev.Instruments {
		if _, err := stmt3.Exec(evaluationID, aci.InstrumentID); err != nil {
			return nil, err
		}
	}

	if err := stmt1.Close(); err != nil {
		return nil, err
	}
	if err := stmt2.Close(); err != nil {
		return nil, err
	}
	if err := stmt3.Close(); err != nil {
		return nil, err
	}
	if err := txn.Commit(); err != nil {
		return nil, err
	}
	return GetEvaluation(db, evaluationID)
}

func DeleteEvaluation(db *sqlx.DB, evaluationID *uuid.UUID) error {
	_, err := db.Exec(
		`DELETE FROM evaluation WHERE id = $1`, evaluationID,
	)
	if err != nil {
		return err
	}
	return nil
}
