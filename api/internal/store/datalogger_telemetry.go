package store

import (
	"context"
	"log"

	"github.com/USACE/instrumentation-api/api/internal/model"
)

type DataloggerTelemetryStore interface {
}

type dataloggerTelemetryStore struct {
	db *model.Database
	q  *model.Queries
}

func NewDataloggerTelemetryStore(db *model.Database, q *model.Queries) *dataloggerTelemetryStore {
	return &dataloggerTelemetryStore{db, q}
}

func (s dataloggerTelemetryStore) GetDataLoggerByModelSN(ctx context.Context, modelName, sn string) (model.Datalogger, error) {
	return s.q.GetDataloggerByModelSN(ctx, modelName, sn)
}

func (s dataloggerTelemetryStore) GetDataLoggerHashByModelSN(ctx context.Context, modelName, sn string) (string, error) {
	return s.q.GetDataloggerHashByModelSN(ctx, modelName, sn)
}

func (s dataloggerTelemetryStore) UpdateDataLoggerPreview(ctx context.Context, dlp model.DataloggerPreview) error {
	return s.q.UpdateDataloggerPreview(ctx, dlp)
}

func (s dataloggerTelemetryStore) UpdateDataLoggerError(ctx context.Context, e *model.DataloggerError) error {
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
