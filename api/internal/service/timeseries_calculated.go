package service

import (
	"context"
	"database/sql"
	"errors"

	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
)

type CalculatedTimeseriesService interface {
	GetAllCalculatedTimeseriesForInstrument(ctx context.Context, instrumentID uuid.UUID) ([]model.CalculatedTimeseries, error)
	ListCalculatedTimeseriesSlugs(ctx context.Context) ([]string, error)
	CreateCalculatedTimeseries(ctx context.Context, cc model.CalculatedTimeseries) error
	UpdateCalculatedTimeseries(ctx context.Context, cts model.CalculatedTimeseries) error
	DeleteCalculatedTimeseries(ctx context.Context, ctsID uuid.UUID) error
}

type calculatedTimeseriesService struct {
	db *model.Database
	*model.Queries
}

func NewCalculatedTimeseriesService(db *model.Database, q *model.Queries) *calculatedTimeseriesService {
	return &calculatedTimeseriesService{db, q}
}

func (s calculatedTimeseriesService) CreateCalculatedTimeseries(ctx context.Context, cc model.CalculatedTimeseries) error {
	if cc.ParameterID == uuid.Nil {
		cc.ParameterID = uuid.Must(uuid.Parse("2b7f96e1-820f-4f61-ba8f-861640af6232")) // unknown parameter
	}
	if cc.UnitID == uuid.Nil {
		cc.UnitID = uuid.Must(uuid.Parse("4a999277-4cf5-4282-93ce-23b33c65e2c8")) // unknown unit
	}

	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)

	tsID, err := qtx.CreateCalculatedTimeseries(ctx, cc)
	if err != nil {
		return err
	}

	if err := qtx.CreateCalculation(ctx, tsID, cc.Formula); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (s calculatedTimeseriesService) UpdateCalculatedTimeseries(ctx context.Context, cts model.CalculatedTimeseries) error {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)

	defaultCts, err := qtx.GetOneCalculation(ctx, &cts.ID)
	if err != nil {
		return err
	}

	if cts.InstrumentID == uuid.Nil {
		cts.InstrumentID = defaultCts.InstrumentID
	}
	if cts.ParameterID == uuid.Nil {
		cts.ParameterID = defaultCts.ParameterID
	}
	if cts.UnitID == uuid.Nil {
		cts.UnitID = defaultCts.UnitID
	}
	if cts.Slug == "" {
		cts.Slug = defaultCts.Slug
	}
	if cts.FormulaName == "" {
		cts.FormulaName = defaultCts.FormulaName
	}
	if cts.Formula == "" {
		cts.Formula = defaultCts.Formula
	}

	if err := qtx.CreateOrUpdateCalculatedTimeseries(ctx, cts, defaultCts); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	}

	if err := qtx.CreateOrUpdateCalculation(ctx, cts.ID, cts.Formula, defaultCts.Formula); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
