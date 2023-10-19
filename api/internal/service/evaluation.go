package service

import (
	"context"

	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
)

type EvaluationService interface {
	ListProjectEvaluations(ctx context.Context, projectID uuid.UUID) ([]model.Evaluation, error)
	ListProjectEvaluationsByAlertConfig(ctx context.Context, projectID, alertConfigID uuid.UUID) ([]model.Evaluation, error)
	ListInstrumentEvaluations(ctx context.Context, instrumentID uuid.UUID) ([]model.Evaluation, error)
	GetEvaluation(ctx context.Context, evaluationID uuid.UUID) (model.Evaluation, error)
	RecordEvaluationSubmittal(ctx context.Context, subID uuid.UUID) error
	CreateEvaluation(ctx context.Context, ev model.Evaluation) (model.Evaluation, error)
	UpdateEvaluation(ctx context.Context, evaluationID uuid.UUID, ev model.Evaluation) (model.Evaluation, error)
	DeleteEvaluation(ctx context.Context, evaluationID uuid.UUID) error
}

type evaluationService struct {
	db *model.Database
	*model.Queries
}

func NewEvaluationService(db *model.Database, q *model.Queries) *evaluationService {
	return &evaluationService{db, q}
}

func (s evaluationService) RecordEvaluationSubmittal(ctx context.Context, subID uuid.UUID) error {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)

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

func (s evaluationService) CreateEvaluation(ctx context.Context, ev model.Evaluation) (model.Evaluation, error) {
	var a model.Evaluation
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return a, err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)

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

func (s evaluationService) UpdateEvaluation(ctx context.Context, evaluationID uuid.UUID, ev model.Evaluation) (model.Evaluation, error) {
	var a model.Evaluation
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return a, err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)

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

func (s evaluationService) DeleteEvaluation(ctx context.Context, evaluationID uuid.UUID) error {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)

	if err := qtx.UnassignAllInstrumentsFromEvaluation(ctx, evaluationID); err != nil {
		return err
	}

	if err := qtx.DeleteEvaluation(ctx, evaluationID); err != nil {
		return err
	}

	return nil
}
