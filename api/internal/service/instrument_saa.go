package service

import (
	"context"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
)

type SaaInstrumentService interface {
	GetAllSaaSegmentsForInstrument(ctx context.Context, instrumentID uuid.UUID) ([]model.SaaSegment, error)
	UpdateSaaSegment(ctx context.Context, seg model.SaaSegment) error
	UpdateSaaSegments(ctx context.Context, segs []model.SaaSegment) error
	GetSaaMeasurementsForInstrument(ctx context.Context, instrumentID uuid.UUID, tw model.TimeWindow) ([]model.SaaMeasurements, error)
}

type saaInstrumentService struct {
	db *model.Database
	*model.Queries
}

func NewSaaInstrumentService(db *model.Database, q *model.Queries) *saaInstrumentService {
	return &saaInstrumentService{db, q}
}

func (s saaInstrumentService) UpdateSaaSegments(ctx context.Context, segs []model.SaaSegment) error {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)

	for _, seg := range segs {
		if err := qtx.UpdateSaaSegment(ctx, seg); err != nil {
			return err
		}
		if seg.Length == nil {
			continue
		}
		if err := qtx.CreateTimeseriesMeasurement(ctx, seg.LengthTimeseriesID, time.Now(), *seg.Length); err != nil {
			return err
		}
	}
	return tx.Commit()
}
