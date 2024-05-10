package service

import (
	"context"
	"fmt"

	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
)

type InstrumentAssignService interface {
	AssignProjectsToInstrument(ctx context.Context, profileID, instrumentID uuid.UUID, projectIDs []uuid.UUID, dryRun bool) (model.InstrumentsValidation, error)
	UnassignProjectsFromInstrument(ctx context.Context, profileID, instrumentID uuid.UUID, projectIDs []uuid.UUID, dryRun bool) (model.InstrumentsValidation, error)
	AssignInstrumentsToProject(ctx context.Context, profileID, projectID uuid.UUID, instrumentIDs []uuid.UUID, dryRun bool) (model.InstrumentsValidation, error)
	UnassignInstrumentsFromProject(ctx context.Context, profileID, projectID uuid.UUID, instrumentIDs []uuid.UUID, dryRun bool) (model.InstrumentsValidation, error)
	ValidateInstrumentNamesProjectUnique(ctx context.Context, projectID uuid.UUID, instrumentNames []string) (model.InstrumentsValidation, error)
	ValidateProjectsInstrumentNameUnique(ctx context.Context, instrumentName string, projectIDs []uuid.UUID) (model.InstrumentsValidation, error)
}

type instrumentAssignService struct {
	db *model.Database
	*model.Queries
}

func NewInstrumentAssignService(db *model.Database, q *model.Queries) *instrumentAssignService {
	return &instrumentAssignService{db, q}
}

func validateAssignInstrumentsToProject(ctx context.Context, q *model.Queries, profileID, projectID uuid.UUID, instruments []model.Instrument) (model.InstrumentsValidation, error) {
	iIDs := make([]uuid.UUID, len(instruments))
	iNames := make([]string, len(instruments))
	for idx := range instruments {
		iIDs[idx] = instruments[idx].ID
		iNames[idx] = instruments[idx].Name
	}
	v, err := q.ValidateInstrumentsAssignerAuthorized(ctx, profileID, iIDs)
	if err != nil || !v.IsValid {
		return v, err
	}
	return q.ValidateInstrumentNamesProjectUnique(ctx, projectID, iNames)
}

func validateAssignProjectsToInstrument(ctx context.Context, q *model.Queries, profileID uuid.UUID, instrument model.Instrument, projectIDs []uuid.UUID) (model.InstrumentsValidation, error) {
	v, err := q.ValidateProjectsAssignerAuthorized(ctx, profileID, instrument.ID, projectIDs)
	if err != nil || !v.IsValid {
		return v, err
	}
	return q.ValidateProjectsInstrumentNameUnique(ctx, instrument.Name, projectIDs)
}

func assignProjectsToInstrument(ctx context.Context, q *model.Queries, profileID, instrumentID uuid.UUID, projectIDs []uuid.UUID) (model.InstrumentsValidation, error) {
	instrument, err := q.GetInstrument(ctx, instrumentID)
	if err != nil {
		return model.InstrumentsValidation{}, err
	}
	v, err := validateAssignProjectsToInstrument(ctx, q, profileID, instrument, projectIDs)
	if err != nil || !v.IsValid {
		return v, err
	}
	for _, pID := range projectIDs {
		if err := q.AssignInstrumentToProject(ctx, pID, instrumentID); err != nil {
			return model.InstrumentsValidation{}, err
		}
	}
	return v, nil
}

func unassignProjectsFromInstrument(ctx context.Context, q *model.Queries, profileID, instrumentID uuid.UUID, projectIDs []uuid.UUID) (model.InstrumentsValidation, error) {
	v, err := q.ValidateProjectsAssignerAuthorized(ctx, profileID, instrumentID, projectIDs)
	if err != nil || !v.IsValid {
		return v, err
	}
	for _, pID := range projectIDs {
		if err := q.UnassignInstrumentFromProject(ctx, pID, instrumentID); err != nil {
			return v, err
		}
	}
	return v, nil
}

func assignInstrumentsToProject(ctx context.Context, q *model.Queries, profileID, projectID uuid.UUID, instrumentIDs []uuid.UUID) (model.InstrumentsValidation, error) {
	v, err := q.ValidateInstrumentsAssignerAuthorized(ctx, profileID, instrumentIDs)
	if err != nil || !v.IsValid {
		return v, err
	}
	for _, iID := range instrumentIDs {
		if err := q.AssignInstrumentToProject(ctx, projectID, iID); err != nil {
			return v, err
		}
	}
	return v, nil
}

func unassignInstrumentsFromProject(ctx context.Context, q *model.Queries, profileID, projectID uuid.UUID, instrumentIDs []uuid.UUID) (model.InstrumentsValidation, error) {
	v, err := q.ValidateInstrumentsAssignerAuthorized(ctx, profileID, instrumentIDs)
	if err != nil || !v.IsValid {
		return v, err
	}
	cc, err := q.GetProjectCountForInstruments(ctx, instrumentIDs)
	if err != nil {
		return model.InstrumentsValidation{}, err
	}

	for _, count := range cc {
		if count.ProjectCount < 1 {
			// invalid instrument, skipping
			continue
		}
		if count.ProjectCount == 1 {
			v.IsValid = false
			v.ReasonCode = model.InvalidUnassign
			v.Errors = append(v.Errors, fmt.Sprintf("cannot unassign instruments from project, all instruments must have at least one project assinment (%s is only assign to this project)", count.InstrumentName))
		}
		if err := q.UnassignInstrumentFromProject(ctx, projectID, count.InstrumentID); err != nil {
			return v, err
		}
	}
	return v, nil
}

func (s instrumentAssignService) AssignProjectsToInstrument(ctx context.Context, profileID, instrumentID uuid.UUID, projectIDs []uuid.UUID, dryRun bool) (model.InstrumentsValidation, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return model.InstrumentsValidation{}, err
	}
	defer model.TxDo(tx.Rollback)
	qtx := s.WithTx(tx)

	v, err := assignProjectsToInstrument(ctx, qtx, profileID, instrumentID, projectIDs)
	if err != nil || !v.IsValid || dryRun {
		return v, err
	}
	return v, tx.Commit()
}

func (s instrumentAssignService) UnassignProjectsFromInstrument(ctx context.Context, profileID, instrumentID uuid.UUID, projectIDs []uuid.UUID, dryRun bool) (model.InstrumentsValidation, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return model.InstrumentsValidation{}, err
	}
	defer model.TxDo(tx.Rollback)
	qtx := s.WithTx(tx)

	v, err := unassignProjectsFromInstrument(ctx, qtx, profileID, instrumentID, projectIDs)
	if err != nil || !v.IsValid || dryRun {
		return v, err
	}
	return v, tx.Commit()
}

func (s instrumentAssignService) AssignInstrumentsToProject(ctx context.Context, profileID, projectID uuid.UUID, instrumentIDs []uuid.UUID, dryRun bool) (model.InstrumentsValidation, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return model.InstrumentsValidation{}, err
	}
	defer model.TxDo(tx.Rollback)
	qtx := s.WithTx(tx)

	v, err := assignInstrumentsToProject(ctx, qtx, profileID, projectID, instrumentIDs)
	if err != nil || !v.IsValid || dryRun {
		return v, err
	}
	return v, tx.Commit()
}

func (s instrumentAssignService) UnassignInstrumentsFromProject(ctx context.Context, profileID, projectID uuid.UUID, instrumentIDs []uuid.UUID, dryRun bool) (model.InstrumentsValidation, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return model.InstrumentsValidation{}, err
	}
	defer model.TxDo(tx.Rollback)
	qtx := s.WithTx(tx)

	v, err := unassignInstrumentsFromProject(ctx, qtx, profileID, projectID, instrumentIDs)
	if err != nil || !v.IsValid || dryRun {
		return v, err
	}
	return v, tx.Commit()
}
