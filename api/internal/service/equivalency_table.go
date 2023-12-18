package service

import (
	"context"

	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
)

type EquivalencyTableService interface {
	GetEquivalencyTable(ctx context.Context, dataloggerTableID uuid.UUID) (model.EquivalencyTable, error)
	CreateEquivalencyTable(ctx context.Context, t model.EquivalencyTable) (model.EquivalencyTable, error)
	UpdateEquivalencyTable(ctx context.Context, t model.EquivalencyTable) (model.EquivalencyTable, error)
	DeleteEquivalencyTable(ctx context.Context, dataloggerTableID uuid.UUID) error
	DeleteEquivalencyTableRow(ctx context.Context, rowID uuid.UUID) error
	GetIsValidDataloggerTable(ctx context.Context, dataloggerTableID uuid.UUID) error
}

type equivalencyTableService struct {
	db *model.Database
	*model.Queries
}

func NewEquivalencyTableService(db *model.Database, q *model.Queries) *equivalencyTableService {
	return &equivalencyTableService{db, q}
}

// CreateEquivalencyTable creates EquivalencyTable rows
// If a row with the given datalogger id or field name already exists the row will be ignored
func (s equivalencyTableService) CreateEquivalencyTable(ctx context.Context, t model.EquivalencyTable) (model.EquivalencyTable, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return model.EquivalencyTable{}, err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)

	for _, r := range t.Rows {
		if r.TimeseriesID != nil {
			if err = qtx.GetIsValidEquivalencyTableTimeseries(ctx, *r.TimeseriesID); err != nil {
				return model.EquivalencyTable{}, err
			}
		}
		if err := qtx.CreateEquivalencyTableRow(ctx, t.DataloggerID, t.DataloggerTableID, r); err != nil {
			return model.EquivalencyTable{}, err
		}
	}

	eqt, err := qtx.GetEquivalencyTable(ctx, t.DataloggerTableID)
	if err != nil {
		return model.EquivalencyTable{}, err
	}

	if err := tx.Commit(); err != nil {
		return model.EquivalencyTable{}, err
	}

	return eqt, nil
}

// UpdateEquivalencyTable updates rows of an EquivalencyTable
func (s equivalencyTableService) UpdateEquivalencyTable(ctx context.Context, t model.EquivalencyTable) (model.EquivalencyTable, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return model.EquivalencyTable{}, err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)

	for _, r := range t.Rows {
		if r.TimeseriesID != nil {
			if err = qtx.GetIsValidEquivalencyTableTimeseries(ctx, *r.TimeseriesID); err != nil {
				return model.EquivalencyTable{}, err
			}
		}
		if err := qtx.UpdateEquivalencyTableRow(ctx, r); err != nil {
			return model.EquivalencyTable{}, err
		}
	}

	eqt, err := qtx.GetEquivalencyTable(ctx, t.DataloggerTableID)

	if err := tx.Commit(); err != nil {
		return model.EquivalencyTable{}, err
	}

	return eqt, nil
}
