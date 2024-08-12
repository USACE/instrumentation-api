package service

import (
	"context"

	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
)

type InstrumentConstantService interface {
	ListInstrumentConstants(ctx context.Context, instrumentID uuid.UUID) ([]model.Timeseries, error)
	CreateInstrumentConstant(ctx context.Context, instrumentID, timeseriesID uuid.UUID) error
	CreateInstrumentConstants(ctx context.Context, tt []model.Timeseries) ([]model.Timeseries, error)
	DeleteInstrumentConstant(ctx context.Context, instrumentID, timeseriesID uuid.UUID) error
}

type instrumentConstantService struct {
	db *model.Database
	*model.Queries
}

func NewInstrumentConstantService(db *model.Database, q *model.Queries) *instrumentConstantService {
	return &instrumentConstantService{db, q}
}

// CreateInstrumentConstants creates many instrument constants from an array of instrument constants
// An InstrumentConstant is structurally the same as a timeseries and saved in the same tables
func (s instrumentConstantService) CreateInstrumentConstants(ctx context.Context, tt []model.Timeseries) ([]model.Timeseries, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)

	uu := make([]model.Timeseries, len(tt))
	for idx, t := range tt {
		t.Type = model.ConstantTimeseriesType
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
func (s instrumentConstantService) DeleteInstrumentConstant(ctx context.Context, instrumentID, timeseriesID uuid.UUID) error {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)

	if err := qtx.DeleteInstrumentConstant(ctx, instrumentID, timeseriesID); err != nil {
		return err
	}

	if err := qtx.DeleteTimeseries(ctx, timeseriesID); err != nil {
		return err
	}

	return tx.Commit()
}
