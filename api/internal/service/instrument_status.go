package service

import (
	"context"

	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
)

type InstrumentStatusService interface {
	ListInstrumentStatus(ctx context.Context, instrumentID uuid.UUID) ([]model.InstrumentStatus, error)
	GetInstrumentStatus(ctx context.Context, statusID uuid.UUID) (model.InstrumentStatus, error)
	CreateOrUpdateInstrumentStatus(ctx context.Context, instrumentID uuid.UUID, ss []model.InstrumentStatus) error
	DeleteInstrumentStatus(ctx context.Context, statusID uuid.UUID) error
}

type instrumentStatusService struct {
	db *model.Database
	*model.Queries
}

func NewInstrumentStatusService(db *model.Database, q *model.Queries) *instrumentStatusService {
	return &instrumentStatusService{db, q}
}

func (s instrumentStatusService) CreateOrUpdateInstrumentStatus(ctx context.Context, instrumentID uuid.UUID, ss []model.InstrumentStatus) error {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)

	for _, updateStatus := range ss {
		if err := qtx.CreateOrUpdateInstrumentStatus(ctx, instrumentID, updateStatus.StatusID, updateStatus.Time); err != nil {
			return err
		}
	}

	return tx.Commit()
}
