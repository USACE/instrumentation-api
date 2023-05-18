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
	sql := `SELECT *
			FROM v_evaluation
			WHERE project_id = $1
	`
	err := db.Select(&aa, sql, projectID)
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

func CreateEvaluation(db *sqlx.DB, ev *Evaluation) (*Evaluation, error) {
	txn, err := db.Beginx()
	if err != nil {
		return nil, err
	}
	defer txn.Rollback()

	stmt1, err := txn.Preparex(`
		INSERT INTO evaluation
			(
				project_id,
				name,
				body,
				start_date,
				end_date,
				creator,
				create_date
			) VALUES ($1,$2,$3,$4,$5,$6,$7)
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
