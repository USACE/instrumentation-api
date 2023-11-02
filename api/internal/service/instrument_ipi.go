package service

import (
	"context"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
)

type IpiInstrumentService interface {
	GetAllIpiSegmentsForInstrument(ctx context.Context, instrumentID uuid.UUID) ([]model.IpiSegment, error)
	UpdateIpiSegment(ctx context.Context, seg model.IpiSegment) error
	UpdateIpiSegments(ctx context.Context, segs []model.IpiSegment) error
	GetIpiMeasurementsForInstrument(ctx context.Context, instrumentID uuid.UUID, tw model.TimeWindow) ([]model.IpiMeasurements, error)
}

type ipiInstrumentService struct {
	db *model.Database
	*model.Queries
}

func NewIpiInstrumentService(db *model.Database, q *model.Queries) *ipiInstrumentService {
	return &ipiInstrumentService{db, q}
}

func (s ipiInstrumentService) UpdateIpiSegments(ctx context.Context, segs []model.IpiSegment) error {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)

	for _, seg := range segs {
		if err := qtx.UpdateIpiSegment(ctx, seg); err != nil {
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
