package service

import (
	"context"

	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
)

type SaaInstrumentService interface {
	CreateSaaInstrument(ctx context.Context, si model.SaaInstrumentWithSegments) error
	CreateSaaSegments(ctx context.Context, segs []model.SaaSegment) error
	GetOneSaaInstrumentWithSegments(ctx context.Context, instrumentID uuid.UUID) (model.SaaInstrumentWithSegments, error)
	GetAllSaaInstrumentsWithSegmentsForProject(ctx context.Context, projectID uuid.UUID) ([]model.SaaInstrumentWithSegments, error)
	UpdateSaaInstrument(ctx context.Context, si model.SaaInstrumentWithSegments) error
	UpdateSaaInstrumentSegment(ctx context.Context, seg model.SaaSegment) error
}

type saaInstrumentService struct {
	db *model.Database
	*model.Queries
}

func NewSaaInstrumentService(db *model.Database, q *model.Queries) *saaInstrumentService {
	return &saaInstrumentService{db, q}
}

func (s saaInstrumentService) CreateSaaInstrument(ctx context.Context, si model.SaaInstrumentWithSegments) error {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)

	newInst, err := qtx.CreateInstrument(ctx, si.Instrument)
	si.ID = newInst.ID

	if err := qtx.CreateSaaInstrument(ctx, si.SaaInstrument); err != nil {
		return err
	}

	for i, seg := range si.Segments {
		seg.ID = i
		if err := qtx.CreateSaaSegment(ctx, seg); err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (s saaInstrumentService) CreateSaaSegments(ctx context.Context, segs []model.SaaSegment) error {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)

	for _, seg := range segs {
		if err := qtx.CreateSaaSegment(ctx, seg); err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (s saaInstrumentService) UpdateSaaInstrument(ctx context.Context, si model.SaaInstrumentWithSegments) error {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)

	if err := qtx.UpdateSaaInstrument(ctx, si.SaaInstrument); err != nil {
		return err
	}

	for _, seg := range si.Segments {
		if err := qtx.CreateSaaSegment(ctx, seg); err != nil {
			return err
		}
	}
	return tx.Commit()
}
