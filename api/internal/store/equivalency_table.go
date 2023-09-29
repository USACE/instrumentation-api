package store

import (
	"context"
	"log"

	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
)

type EquivalencyTableStore interface {
	GetEquivalencyTable(ctx context.Context, dlID uuid.UUID) (model.EquivalencyTable, error)
	CreateEquivalencyTable(ctx context.Context, t model.EquivalencyTable) error
	UpdateEquivalencyTable(ctx context.Context, t *model.EquivalencyTable) error
	DeleteEquivalencyTable(ctx context.Context, dataloggerID uuid.UUID) error
	DeleteEquivalencyTableRow(ctx context.Context, dataloggerID, rowID uuid.UUID) error
}

type equivalencyTableStore struct {
	db *model.Database
	*model.Queries
}

func NewEquivalencyTableStore(db *model.Database, q *model.Queries) *equivalencyTableStore {
	return &equivalencyTableStore{db, q}
}

// CreateEquivalencyTable creates EquivalencyTable rows
// If a row with the given datalogger id or field name already exists the row will be ignored
func (s equivalencyTableStore) CreateEquivalencyTable(ctx context.Context, t model.EquivalencyTable) error {
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

	for _, r := range t.Rows {
		if r.TimeseriesID == nil {
			continue
		}
		if err = qtx.GetIsValidEquivalencyTableTimeseries(ctx, *r.TimeseriesID); err != nil {
			return err
		}
		if err := qtx.CreateEquivalencyTableRow(ctx, r); err != nil {
			return err
		}
	}
	return tx.Commit()
}

// UpdateEquivalencyTable updates rows of an EquivalencyTable
func (s equivalencyTableStore) UpdateEquivalencyTable(ctx context.Context, t *model.EquivalencyTable) error {
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

	for _, r := range t.Rows {
		if r.TimeseriesID == nil {
			continue
		}
		if err = qtx.GetIsValidEquivalencyTableTimeseries(ctx, *r.TimeseriesID); err != nil {
			return err
		}
		if err := qtx.UpdateEquivalencyTableRow(ctx, r); err != nil {
			return err
		}
	}
	return tx.Commit()
}
