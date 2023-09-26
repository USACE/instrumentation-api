package store

import (
	"context"
	"log"

	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
)

type CalculatedTimeseriesStore interface {
	GetAllCalculationsForInstrument(ctx context.Context, instrumentID *uuid.UUID) ([]model.CalculatedTimeseries, error)
}

type calculatedTimeseriesStore struct {
	db *model.Database
	q  *model.Queries
}

func NewCalculatedTimeseriesStore(db *model.Database, q *model.Queries) *calculatedTimeseriesStore {
	return &calculatedTimeseriesStore{db, q}
}

// GetInstrumentCalculations returns all formulas associated to a given instrument ID.
func (s calculatedTimeseriesStore) GetAllCalculatedTimeseriesForInstrument(ctx context.Context, instrumentID uuid.UUID) ([]model.CalculatedTimeseries, error) {
	return s.q.GetAllCalculatedTimeseriesForInstrument(ctx, instrumentID)
}

func (s calculatedTimeseriesStore) ListCalculatedTimeseriesSlugs(ctx context.Context) ([]string, error) {
	return s.q.ListCalculatedTimeseriesSlugs(ctx)
}

func (s calculatedTimeseriesStore) CreateCalculatedTimeseries(ctx context.Context, cc model.CalculatedTimeseries) error {
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
	defer func() {
		if err := tx.Rollback(); err != nil {
			log.Print("error rolling back changes")
		}
	}()

	qtx := s.q.WithTx(tx)

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

func (s calculatedTimeseriesStore) UpdateCalculatedTimeseries(ctx context.Context, cts model.CalculatedTimeseries) error {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			log.Print("error rolling back changes")
		}
	}()

	qtx := s.q.WithTx(tx)

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

	if err := qtx.CreateOrUpdateCalculatedTimeseries(ctx, cts, defaultCts); err != nil {
		return err
	}

	if err := qtx.CreateOrUpdateCalculation(ctx, cts.ID, cts.Formula, defaultCts.Formula); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

// DeleteCalculation removes the `Calculation` with ID `formulaID` from the database,
// effectively dissociating it from the instrument in question.
func (s calculatedTimeseriesStore) DeleteCalculatedTimeseries(ctx context.Context, ctsID uuid.UUID) error {
	return s.q.DeleteCalculatedTimeseries(ctx, ctsID)
}
