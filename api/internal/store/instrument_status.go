package store

import (
	"context"
	"log"

	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
)

type InstrumentStatusStore interface {
}

type instrumentStatusStore struct {
	db *model.Database
	q  *model.Queries
}

func NewInstrumentStatusStore(db *model.Database, q *model.Queries) *instrumentStatusStore {
	return &instrumentStatusStore{db, q}
}

// ListInstrumentStatus returns all status values for an instrument
func (s instrumentStatusStore) ListInstrumentStatus(ctx context.Context, instrumentID uuid.UUID) ([]model.InstrumentStatus, error) {
	return s.q.ListInstrumentStatus(ctx, instrumentID)
}

// GetInstrumentStatus gets a single status
func (s instrumentStatusStore) GetInstrumentStatus(ctx context.Context, statusID uuid.UUID) (model.InstrumentStatus, error) {
	return s.q.GetInstrumentStatus(ctx, statusID)
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

	qtx := s.q.WithTx(tx)

	for _, updateStatus := range ss {
		if err := qtx.CreateOrUpdateInstrumentStatus(ctx, instrumentID, updateStatus.StatusID, updateStatus.Time); err != nil {
			return err
		}
	}

	return tx.Commit()
}

// DeleteInstrumentStatus deletes a status for an instrument
func (s instrumentStatusStore) DeleteInstrumentStatus(ctx context.Context, statusID uuid.UUID) error {
	return s.q.DeleteInstrumentStatus(ctx, statusID)
}
