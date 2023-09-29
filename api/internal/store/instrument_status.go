package store

import (
	"context"
	"log"

	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
)

type InstrumentStatusStore interface {
	ListInstrumentStatus(ctx context.Context, instrumentID uuid.UUID) ([]model.InstrumentStatus, error)
	GetInstrumentStatus(ctx context.Context, statusID uuid.UUID) (model.InstrumentStatus, error)
	CreateOrUpdateInstrumentStatus(ctx context.Context, instrumentID uuid.UUID, ss []model.InstrumentStatus) error
	DeleteInstrumentStatus(ctx context.Context, statusID uuid.UUID) error
}

type instrumentStatusStore struct {
	db *model.Database
	*model.Queries
}

func NewInstrumentStatusStore(db *model.Database, q *model.Queries) *instrumentStatusStore {
	return &instrumentStatusStore{db, q}
}

func (s instrumentStatusStore) CreateOrUpdateInstrumentStatus(ctx context.Context, instrumentID uuid.UUID, ss []model.InstrumentStatus) error {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			log.Print(err.Error())
		}
	}()

	qtx := s.WithTx(tx)

	for _, updateStatus := range ss {
		if err := qtx.CreateOrUpdateInstrumentStatus(ctx, instrumentID, updateStatus.StatusID, updateStatus.Time); err != nil {
			return err
		}
	}

	return tx.Commit()
}
