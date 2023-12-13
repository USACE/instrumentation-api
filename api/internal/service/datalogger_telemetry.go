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
