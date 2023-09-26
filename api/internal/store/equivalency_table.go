package store

import (
	"context"
	"log"

	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
)

type EquivalencyTableStore interface {
}

type equivalencyTableStore struct {
	db *model.Database
	q  *model.Queries
}

func NewEquivalencyTableStore(db *model.Database, q *model.Queries) *equivalencyTableStore {
	return &equivalencyTableStore{db, q}
}

// GetEquivalencyTable returns a single DataLogger EquivalencyTable
func (s equivalencyTableStore) GetEquivalencyTable(ctx context.Context, dlID uuid.UUID) (model.EquivalencyTable, error) {
	return s.q.GetEquivalencyTable(ctx, dlID)
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

	qtx := s.q.WithTx(tx)

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

	qtx := s.q.WithTx(tx)

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

// DeleteEquivalencyTable clears all rows of the EquivalencyTable for a Datalogger
func (s equivalencyTableStore) DeleteEquivalencyTable(ctx context.Context, dataloggerID uuid.UUID) error {
	return s.q.DeleteEquivalencyTable(ctx, dataloggerID)
}

// DeleteEquivalencyTableRow deletes a single EquivalencyTable row by row id
func (s equivalencyTableStore) DeleteEquivalencyTableRow(ctx context.Context, dataloggerID, rowID uuid.UUID) error {
	return s.q.DeleteEquivalencyTableRow(ctx, dataloggerID, rowID)
}
