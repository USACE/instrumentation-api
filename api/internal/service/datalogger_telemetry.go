package service

import (
	"context"
	"database/sql"
	"errors"

	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
)

type DataloggerTelemetryService interface {
	GetDataloggerByModelSN(ctx context.Context, modelName, sn string) (model.Datalogger, error)
	GetDataloggerHashByModelSN(ctx context.Context, modelName, sn string) (string, error)
	CreateDataloggerTablePreview(ctx context.Context, prv model.DataloggerTablePreview) error
	UpdateDataloggerTablePreview(ctx context.Context, dataloggerID uuid.UUID, tableName string, prv model.DataloggerTablePreview) (uuid.UUID, error)
	UpdateDataloggerTableError(ctx context.Context, dataloggerID uuid.UUID, tableName *string, e *model.DataloggerError) error
}

type dataloggerTelemetryService struct {
	db *model.Database
	*model.Queries
}

func NewDataloggerTelemetryService(db *model.Database, q *model.Queries) *dataloggerTelemetryService {
	return &dataloggerTelemetryService{db, q}
}

// UpdateDataloggerTablePreview attempts to update a table preview by datalogger_id and table_name, creates the
// datalogger table and corresponding preview if it doesn't exist
func (s dataloggerTelemetryService) UpdateDataloggerTablePreview(ctx context.Context, dataloggerID uuid.UUID, tableName string, prv model.DataloggerTablePreview) (uuid.UUID, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return uuid.Nil, err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)

	tableID, err := qtx.GetOrCreateDataloggerTable(ctx, dataloggerID, tableName)
	if err != nil {
		return uuid.Nil, err
	}
	if err := qtx.UpdateDataloggerTablePreview(ctx, dataloggerID, tableName, prv); err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return uuid.Nil, err
		}
		prv.DataloggerTableID = tableID
		if err := qtx.CreateDataloggerTablePreview(ctx, prv); err != nil {
		}
	}

	return tableID, tx.Commit()
}

func (s dataloggerTelemetryService) UpdateDataloggerTableError(ctx context.Context, dataloggerID uuid.UUID, tableName *string, e *model.DataloggerError) error {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)

	if err := qtx.DeleteDataloggerTableError(ctx, dataloggerID, tableName); err != nil {
		return err
	}

	for _, m := range e.Errors {
		if err := qtx.CreateDataloggerTableError(ctx, dataloggerID, tableName, m); err != nil {
			return err
		}
	}

	return tx.Commit()
}
