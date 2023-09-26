package store

import (
	"context"
	"log"

	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
)

type InstrumentConstantStore interface {
}

type instrumentConstantStore struct {
	db *model.Database
	q  *model.Queries
}

func NewInstrumentConstantStore(db *model.Database, q *model.Queries) *instrumentConstantStore {
	return &instrumentConstantStore{db, q}
}

// ListInstrumentConstants lists constants for a given instrument id
func (s instrumentConstantStore) ListInstrumentConstants(ctx context.Context, instrumentID uuid.UUID) ([]model.Timeseries, error) {
	return s.q.ListInstrumentConstants(ctx, instrumentID)
}

func (s instrumentConstantStore) CreateInstrumentConstant(ctx context.Context, instrumentID, timeseriesID uuid.UUID) error {
	return s.q.CreateInstrumentConstant(ctx, instrumentID, timeseriesID)
}

// CreateInstrumentConstants creates many instrument constants from an array of instrument constants
// An InstrumentConstant is structurally the same as a timeseries and saved in the same tables
func (s instrumentConstantStore) CreateInstrumentConstants(ctx context.Context, tt []model.Timeseries) ([]model.Timeseries, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			log.Print(err.Error())
		}
	}()

	qtx := s.q.WithTx(tx)

	uu := make([]model.Timeseries, len(tt))
	for idx, t := range tt {
		tsNew, err := qtx.CreateTimeseries(ctx, t)
		if err != nil {
			return nil, err
		}
		if err := qtx.CreateInstrumentConstant(ctx, tsNew.InstrumentID, tsNew.ID); err != nil {
			return nil, err
		}
		uu[idx] = tsNew
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return uu, nil
}

// DeleteInstrumentConstant removes a timeseries as an Instrument Constant; Does not delete underlying timeseries
func (s instrumentConstantStore) DeleteInstrumentConstant(ctx context.Context, instrumentID, timeseriesID uuid.UUID) error {
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

	if err := qtx.DeleteInstrumentConstant(ctx, instrumentID, timeseriesID); err != nil {
		return err
	}

	if err := qtx.DeleteTimeseries(ctx, timeseriesID); err != nil {
		return err
	}

	return tx.Commit()
}
