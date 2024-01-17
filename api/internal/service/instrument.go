package service

import (
	"context"
	"fmt"

	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
	"github.com/paulmach/orb/geojson"
)

type InstrumentService interface {
	ListInstruments(ctx context.Context) ([]model.Instrument, error)
	GetInstrument(ctx context.Context, instrumentID uuid.UUID) (model.Instrument, error)
	GetInstrumentCount(ctx context.Context) (model.InstrumentCount, error)
	CreateInstrument(ctx context.Context, i model.Instrument) (model.IDSlugName, error)
	CreateInstruments(ctx context.Context, instruments []model.Instrument) ([]model.IDSlugName, error)
	AssignerIsAuthorized(ctx context.Context, profileID, instrumentID uuid.UUID) (bool, error)
	AssignInstrumentToProject(ctx context.Context, projectID, instrumentID uuid.UUID) error
	UnassignInstrumentFromProject(ctx context.Context, projectID, instrumentID uuid.UUID) error
	ValidateCreateInstruments(ctx context.Context, instruments []model.Instrument) (model.CreateInstrumentsValidationResult, error)
	UpdateInstrument(ctx context.Context, projectID uuid.UUID, i model.Instrument) (model.Instrument, error)
	UpdateInstrumentGeometry(ctx context.Context, projectID, instrumentID uuid.UUID, geom geojson.Geometry, p model.Profile) (model.Instrument, error)
	DeleteFlagInstrument(ctx context.Context, projectID, instrumentID uuid.UUID) error
}

type instrumentService struct {
	db *model.Database
	*model.Queries
}

func NewInstrumentService(db *model.Database, q *model.Queries) *instrumentService {
	return &instrumentService{db, q}
}

var (
	saaTypeID = uuid.MustParse("07b91c5c-c1c5-428d-8bb9-e4c93ab2b9b9")
	ipiTypeID = uuid.MustParse("c81f3a5d-fc5f-47fd-b545-401fe6ee63bb")
)

type requestType int

const (
	create requestType = iota
	update
)

func createInstrument(ctx context.Context, q *model.Queries, instrument model.Instrument) (model.IDSlugName, error) {
	newInstrument, err := q.CreateInstrument(ctx, instrument)
	if err != nil {
		return model.IDSlugName{}, err
	}
	if err := q.CreateOrUpdateInstrumentStatus(ctx, newInstrument.ID, instrument.StatusID, instrument.StatusTime); err != nil {
		return model.IDSlugName{}, err
	}
	if instrument.AwareID != nil {
		if err := q.CreateAwarePlatform(ctx, newInstrument.ID, *instrument.AwareID); err != nil {
			return model.IDSlugName{}, err
		}
	}
	instrument.ID = newInstrument.ID
	if err := handleOpts(ctx, q, instrument, create); err != nil {
		return model.IDSlugName{}, err
	}
	return newInstrument, nil
}

func (s instrumentService) CreateInstrument(ctx context.Context, instrument model.Instrument) (model.IDSlugName, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return model.IDSlugName{}, err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)

	newInstrument, err := createInstrument(ctx, qtx, instrument)
	if err != nil {
		return model.IDSlugName{}, err
	}

	if err := tx.Commit(); err != nil {
		return model.IDSlugName{}, err
	}
	return newInstrument, nil
}

func (s instrumentService) CreateInstruments(ctx context.Context, instruments []model.Instrument) ([]model.IDSlugName, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)

	ii := make([]model.IDSlugName, len(instruments))
	for idx, i := range instruments {
		newInstrument, err := createInstrument(ctx, qtx, i)
		if err != nil {
			return nil, err
		}
		ii[idx] = newInstrument
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return ii, nil
}

func (s instrumentService) AssignerIsAuthorized(ctx context.Context, profileID, instrumentID uuid.UUID) (bool, error) {
	q := s.db.Queries()

	instrumentProjectIDs, err := q.ListInstrumentProjects(ctx, instrumentID)
	if err != nil {
		return false, err
	}
	if len(instrumentProjectIDs) == 0 {
		return false, fmt.Errorf("invalid instrument %s has no projects assigned (instrument may be deleted or not exist)", instrumentID)
	}

	adminProjectIDs, err := q.ListAdminProjects(ctx, profileID)
	if err != nil {
		return false, err
	}

	pIDSet := make(map[uuid.UUID]struct{})
	for _, pID := range adminProjectIDs {
		pIDSet[pID] = struct{}{}
	}

	for _, ipID := range instrumentProjectIDs {
		if _, exists := pIDSet[ipID]; !exists {
			return false, nil
		}
	}

	return true, nil
}

// UpdateInstrument updates a single instrument
func (s instrumentService) UpdateInstrument(ctx context.Context, projectID uuid.UUID, i model.Instrument) (model.Instrument, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return model.Instrument{}, err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)

	if err := qtx.UpdateInstrument(ctx, projectID, i); err != nil {
		return model.Instrument{}, err
	}
	if err := qtx.CreateOrUpdateInstrumentStatus(ctx, i.ID, i.StatusID, i.StatusTime); err != nil {
		return model.Instrument{}, err
	}

	if err := handleOpts(ctx, qtx, i, update); err != nil {
		return model.Instrument{}, err
	}

	aa, err := qtx.GetInstrument(ctx, i.ID)
	if err != nil {
		return model.Instrument{}, err
	}

	if err := tx.Commit(); err != nil {
		return model.Instrument{}, err
	}

	return aa, nil
}

func (s instrumentService) UpdateInstrumentGeometry(ctx context.Context, projectID, instrumentID uuid.UUID, geom geojson.Geometry, p model.Profile) (model.Instrument, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return model.Instrument{}, err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)

	if err := qtx.UpdateInstrumentGeometry(ctx, projectID, instrumentID, geom, p); err != nil {
		return model.Instrument{}, err
	}

	aa, err := qtx.GetInstrument(ctx, instrumentID)
	if err != nil {
		return model.Instrument{}, err
	}

	if err := tx.Commit(); err != nil {
		return model.Instrument{}, err
	}

	return aa, nil
}
