package service

import (
	"context"

	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
	"github.com/paulmach/orb/geojson"
)

type InstrumentService interface {
	ListInstrumentSlugs(ctx context.Context) ([]string, error)
	ListInstruments(ctx context.Context) ([]model.Instrument, error)
	GetInstrument(ctx context.Context, instrumentID uuid.UUID) (model.Instrument, error)
	GetInstrumentCount(ctx context.Context) (int, error)
	CreateInstrument(ctx context.Context, i model.Instrument) (model.IDAndSlug, error)
	CreateInstruments(ctx context.Context, instruments []model.Instrument) ([]model.IDAndSlug, error)
	ValidateCreateInstruments(ctx context.Context, instruments []model.Instrument) (model.CreateInstrumentsValidationResult, error)
	UpdateInstrument(ctx context.Context, i model.Instrument) (model.Instrument, error)
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

func (s instrumentService) CreateInstruments(ctx context.Context, instruments []model.Instrument) ([]model.IDAndSlug, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)

	ii := make([]model.IDAndSlug, len(instruments))
	for idx, i := range instruments {
		newInstrument, err := qtx.CreateInstrument(ctx, i)
		if err != nil {
			return nil, err
		}
		ii[idx] = newInstrument

		if err := qtx.CreateOrUpdateInstrumentStatus(ctx, newInstrument.ID, i.StatusID, i.StatusTime); err != nil {
			return nil, err
		}

		if i.AwareID != nil {
			if err := qtx.CreateAwarePlatform(ctx, newInstrument.ID, *i.AwareID); err != nil {
				return nil, err
			}
		}
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return ii, nil
}

// UpdateInstrument updates a single instrument
func (s instrumentService) UpdateInstrument(ctx context.Context, i model.Instrument) (model.Instrument, error) {
	e := model.Instrument{}
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return e, err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)

	if err := qtx.UpdateInstrument(ctx, i); err != nil {
		return e, err
	}
	if err := qtx.CreateOrUpdateInstrumentStatus(ctx, i.ID, i.StatusID, i.StatusTime); err != nil {
		return e, err
	}

	aa, err := qtx.GetInstrument(ctx, i.ID)
	if err != nil {
		return e, err
	}

	if err := tx.Commit(); err != nil {
		return e, err
	}

	return aa, nil
}

func (s instrumentService) UpdateInstrumentGeometry(ctx context.Context, projectID, instrumentID uuid.UUID, geom geojson.Geometry, p model.Profile) (model.Instrument, error) {
	e := model.Instrument{}
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return e, err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)

	if err := qtx.UpdateInstrumentGeometry(ctx, projectID, instrumentID, geom, p); err != nil {
		return e, err
	}

	aa, err := qtx.GetInstrument(ctx, instrumentID)
	if err != nil {
		return e, err
	}

	if err := tx.Commit(); err != nil {
		return e, err
	}

	return aa, nil
}
