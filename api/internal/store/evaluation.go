package store

import (
	"context"
	"log"

	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
)

type EvaluationStore interface {
}

type evaluationStore struct {
	db *model.Database
	q  *model.Queries
}

func NewEvaluationStore(db *model.Database, q *model.Queries) *evaluationStore {
	return &evaluationStore{db, q}
}

func (s evaluationStore) ListProjectEvaluations(ctx context.Context, projectID uuid.UUID) ([]model.Evaluation, error) {
	return s.q.ListProjectEvaluations(ctx, projectID)
}

func (s evaluationStore) ListProjectEvaluationsByAlertConfig(ctx context.Context, projectID, alertConfigID uuid.UUID) ([]model.Evaluation, error) {
	return s.q.ListProjectEvaluationsByAlertConfig(ctx, projectID, alertConfigID)
}

func (s evaluationStore) ListInstrumentEvaluations(ctx context.Context, instrumentID uuid.UUID) ([]model.Evaluation, error) {
	return s.q.ListInstrumentEvaluations(ctx, instrumentID)
}

func (s evaluationStore) GetEvaluation(ctx context.Context, evaluationID uuid.UUID) (model.Evaluation, error) {
	return s.q.GetEvaluation(ctx, evaluationID)
}

func (s evaluationStore) RecordEvaluationSubmittal(ctx context.Context, subID uuid.UUID) error {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			log.Print(err.Error())
		}
	}()

	qtx := s.q.WithTx(tx)

	sub, err := qtx.CompleteEvaluationSubmittal(ctx, subID)
	if err != nil {
		return err
	}

	// Create next submittal if submitted on-time
	// late submittals will have already generated next submittal
	if sub.SubmittalStatusID == GreenSubmittalStatusID {
		if err := qtx.CreateNextEvaluationSubmittal(ctx, subID); err != nil {
			return err
		}
	}
	return tx.Commit()
}

func (s evaluationStore) CreateEvaluation(ctx context.Context, ev model.Evaluation) (model.Evaluation, error) {
	var a model.Evaluation
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return a, err
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			log.Print(err.Error())
		}
	}()

	qtx := s.q.WithTx(tx)

	if ev.SubmittalID != nil {
		sub, err := qtx.CompleteEvaluationSubmittal(ctx, *ev.SubmittalID)
		if err != nil {
			return a, err
		}
		// Create next submittal if submitted on-time
		// late submittals will have already generated next submittal
		if sub.SubmittalStatusID == GreenSubmittalStatusID {
			qtx.CreateNextEvaluationSubmittal(ctx, *ev.SubmittalID)
		}
	}

	evID, err := qtx.CreateEvaluation(ctx, ev)
	if err != nil {
		return a, err
	}

	for _, aci := range ev.Instruments {
		if err := qtx.CreateEvaluationInstrument(ctx, evID, aci.InstrumentID); err != nil {
			return a, err
		}
	}

	evNew, err := qtx.GetEvaluation(ctx, evID)
	if err != nil {
		return a, err
	}

	if err := tx.Commit(); err != nil {
		return a, err
	}

	return evNew, nil
}

func (s evaluationStore) UpdateEvaluation(ctx context.Context, evaluationID *uuid.UUID, ev model.Evaluation) (model.Evaluation, error) {
	var a model.Evaluation
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return a, err
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			log.Print(err.Error())
		}
	}()

	qtx := s.q.WithTx(tx)

	if err := qtx.UpdateEvaluation(ctx, ev); err != nil {
		return a, err
	}

	if err := qtx.UnassignAllInstrumentsFromEvaluation(ctx, ev.ID); err != nil {
		return a, err
	}

	for _, aci := range ev.Instruments {
		if err := qtx.CreateEvaluationInstrument(ctx, ev.ID, aci.InstrumentID); err != nil {
			return a, err
		}
	}

	evUpdated, err := qtx.GetEvaluation(ctx, ev.ID)
	if err != nil {
		return a, err
	}

	if err := tx.Commit(); err != nil {
		return a, err
	}
	return evUpdated, nil
}

func (s evaluationStore) DeleteEvaluation(ctx context.Context, evaluationID uuid.UUID) error {
	return s.q.DeleteEvaluation(ctx, evaluationID)
}
