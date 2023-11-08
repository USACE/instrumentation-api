package service

import (
	"context"

	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
)

type DataloggerTelemetryService interface {
	GetDataloggerByModelSN(ctx context.Context, modelName, sn string) (model.Datalogger, error)
	GetDataloggerHashByModelSN(ctx context.Context, modelName, sn string) (string, error)
	UpdateDataloggerTablePreview(ctx context.Context, dataloggerID uuid.UUID, tableName string, dlp model.DataloggerPreview) error
	UpdateDataloggerTableError(ctx context.Context, dataloggerID uuid.UUID, tableName *string, e *model.DataloggerError) error
	GetOrCreateDataloggerTable(ctx context.Context, dataloggerID uuid.UUID, tableName string) (uuid.UUID, error)
}

type dataloggerTelemetryService struct {
	db *model.Database
	*model.Queries
}

func NewDataloggerTelemetryService(db *model.Database, q *model.Queries) *dataloggerTelemetryService {
	return &dataloggerTelemetryService{db, q}
}

func (s dataloggerTelemetryService) UpdateDataloggerTableError(ctx context.Context, dataloggerID uuid.UUID, tableName *string, e *model.DataloggerError) error {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)

	if err := qtx.DeleteDataloggerTableError(ctx, e.DataloggerTableID, tableName); err != nil {
		return err
	}

	for _, m := range e.Errors {
		if err := qtx.CreateDataloggerTableError(ctx, e.DataloggerTableID, tableName, m); err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (s dataloggerTelemetryService) GetOrCreateDataloggerTable(ctx context.Context, dataloggerID uuid.UUID, tableName string) (uuid.UUID, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return uuid.Nil, err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)

	if err := qtx.RenameEmptyDataloggerTableName(ctx, dataloggerID, tableName); err != nil {
		return uuid.Nil, err
	}

	dataloggerTableID, err := qtx.GetOrCreateDataloggerTable(ctx, dataloggerID, tableName)
	if err != nil {
		return uuid.Nil, err
	}

	if err := tx.Commit(); err != nil {
		return uuid.Nil, err
	}

	return dataloggerTableID, nil
}
