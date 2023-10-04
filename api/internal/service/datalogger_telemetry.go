package service

import (
	"context"

	"github.com/USACE/instrumentation-api/api/internal/model"
)

type DataloggerTelemetryService interface {
	GetDataloggerByModelSN(ctx context.Context, modelName, sn string) (model.Datalogger, error)
	GetDataloggerHashByModelSN(ctx context.Context, modelName, sn string) (string, error)
	UpdateDataloggerPreview(ctx context.Context, dlp model.DataloggerPreview) error
	UpdateDataloggerError(ctx context.Context, e *model.DataloggerError) error
}

type dataloggerTelemetryService struct {
	db *model.Database
	*model.Queries
}

func NewDataloggerTelemetryService(db *model.Database, q *model.Queries) *dataloggerTelemetryService {
	return &dataloggerTelemetryService{db, q}
}

func (s dataloggerTelemetryService) UpdateDataloggerError(ctx context.Context, e *model.DataloggerError) error {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)

	if err := qtx.DeleteDataloggerError(ctx, e.DataloggerID); err != nil {
		return err
	}

	for _, m := range e.Errors {
		if err := qtx.CreateDataloggerError(ctx, e.DataloggerID, m); err != nil {
			return err
		}
	}

	return tx.Commit()
}
