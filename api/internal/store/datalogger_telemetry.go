package store

import (
	"context"
	"log"

	"github.com/USACE/instrumentation-api/api/internal/model"
)

type DataloggerTelemetryStore interface {
	GetDataloggerByModelSN(ctx context.Context, modelName, sn string) (model.Datalogger, error)
	GetDataloggerHashByModelSN(ctx context.Context, modelName, sn string) (string, error)
	UpdateDataloggerPreview(ctx context.Context, dlp model.DataloggerPreview) error
	UpdateDataloggerError(ctx context.Context, e *model.DataloggerError) error
}

type dataloggerTelemetryStore struct {
	db *model.Database
	*model.Queries
}

func NewDataloggerTelemetryStore(db *model.Database, q *model.Queries) *dataloggerTelemetryStore {
	return &dataloggerTelemetryStore{db, q}
}

func (s dataloggerTelemetryStore) UpdateDataloggerError(ctx context.Context, e *model.DataloggerError) error {
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
