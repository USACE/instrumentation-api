package store

import (
	"context"
	"log"

	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
)

type InstrumentConstantStore interface {
	ListInstrumentConstants(ctx context.Context, instrumentID uuid.UUID) ([]model.Timeseries, error)
	CreateInstrumentConstant(ctx context.Context, instrumentID, timeseriesID uuid.UUID) error
	CreateInstrumentConstants(ctx context.Context, tt []model.Timeseries) ([]model.Timeseries, error)
	DeleteInstrumentConstant(ctx context.Context, instrumentID, timeseriesID uuid.UUID) error
}

type instrumentConstantStore struct {
	db *model.Database
	*model.Queries
}

func NewInstrumentConstantStore(db *model.Database, q *model.Queries) *instrumentConstantStore {
	return &instrumentConstantStore{db, q}
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

	qtx := s.WithTx(tx)

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

	qtx := s.WithTx(tx)

	if err := qtx.DeleteInstrumentConstant(ctx, instrumentID, timeseriesID); err != nil {
		return err
	}

	if err := qtx.DeleteTimeseries(ctx, timeseriesID); err != nil {
		return err
	}

	return tx.Commit()
}
