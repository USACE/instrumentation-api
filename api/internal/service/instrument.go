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

var (
	saaTypeID = uuid.MustParse("07b91c5c-c1c5-428d-8bb9-e4c93ab2b9b9")
)

type requestType int

const (
	create requestType = iota
	update
)

func handleOpts(ctx context.Context, q *model.Queries, inst model.Instrument, rt requestType) error {
	switch inst.TypeID {
	case saaTypeID:
		opts, err := model.MapToStruct[model.SaaOpts](inst.Opts)
		if err != nil {
			return err
		}
		if rt == create {
			if err := q.CreateSaaOpts(ctx, inst.ID, opts); err != nil {
				return err
			}
			for i := 1; i <= opts.NumSegments; i++ {
				if err := q.CreateSaaSegment(ctx, model.SaaSegment{ID: i, InstrumentID: inst.ID}); err != nil {
					return err
				}
			}
		}
		if rt == update {
			if err := q.UpdateSaaOpts(ctx, inst.ID, opts); err != nil {
				return err
			}
		}
	default:
	}
	return nil
}

func createInstrument(ctx context.Context, q *model.Queries, instrument model.Instrument) (model.IDAndSlug, error) {
	newInstrument, err := q.CreateInstrument(ctx, instrument)
	if err != nil {
		return model.IDAndSlug{}, err
	}
	if err := q.CreateOrUpdateInstrumentStatus(ctx, newInstrument.ID, instrument.StatusID, instrument.StatusTime); err != nil {
		return model.IDAndSlug{}, err
	}
	if instrument.AwareID != nil {
		if err := q.CreateAwarePlatform(ctx, newInstrument.ID, *instrument.AwareID); err != nil {
			return model.IDAndSlug{}, err
		}
	}
	if err := handleOpts(ctx, q, instrument, create); err != nil {
		return model.IDAndSlug{}, err
	}
	return newInstrument, nil
}

func (s instrumentService) CreateInstrument(ctx context.Context, instrument model.Instrument) (model.IDAndSlug, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return model.IDAndSlug{}, err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)

	newInstrument, err := createInstrument(ctx, qtx, instrument)
	if err != nil {
		return model.IDAndSlug{}, err
	}

	if err := tx.Commit(); err != nil {
		return model.IDAndSlug{}, err
	}
	return newInstrument, nil
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

// UpdateInstrument updates a single instrument
func (s instrumentService) UpdateInstrument(ctx context.Context, i model.Instrument) (model.Instrument, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return model.Instrument{}, err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)

	if err := qtx.UpdateInstrument(ctx, i); err != nil {
		return model.Instrument{}, err
	}
	if err := qtx.CreateOrUpdateInstrumentStatus(ctx, i.ID, i.StatusID, i.StatusTime); err != nil {
		return model.Instrument{}, err
	}

	aa, err := qtx.GetInstrument(ctx, i.ID)
	if err != nil {
		return model.Instrument{}, err
	}

	if err := handleOpts(ctx, qtx, i, update); err != nil {
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
