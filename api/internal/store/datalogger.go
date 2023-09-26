package store

import (
	"context"
	"log"

	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
)

type DataloggerStore interface {
}

type dataloggerStore struct {
	db *model.Database
	q  *model.Queries
}

func NewDataloggerStore(db *model.Database, q *model.Queries) *dataloggerStore {
	return &dataloggerStore{db, q}
}

func (s dataloggerStore) GetDataLoggerModel(ctx context.Context, modelID uuid.UUID) (string, error) {
	return s.q.GetDataloggerModelName(ctx, modelID)
}

func (s dataloggerStore) ListProjectDataLoggers(ctx context.Context, projectID uuid.UUID) ([]model.Datalogger, error) {
	return s.q.ListProjectDataloggers(ctx, projectID)
}

func (s dataloggerStore) ListAllDataLoggers(ctx context.Context) ([]model.Datalogger, error) {
	return s.q.ListAllDataLoggers(ctx)
}

func (s dataloggerStore) DataLoggerActive(ctx context.Context, modelName, sn string) (bool, error) {
	return s.q.GetDataloggerIsActive(ctx, modelName, sn)
}

func (s dataloggerStore) VerifyDataLoggerExists(ctx context.Context, dlID uuid.UUID) error {
	return s.q.VerifyDataloggerExists(ctx, dlID)
}

func (s dataloggerStore) CreateDataLogger(ctx context.Context, n model.Datalogger) (model.DataloggerWithKey, error) {
	var a model.DataloggerWithKey
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return a, err
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			log.Print(err.Error())
		}
	}()

	qtx := s.q.WithTx(tx)

	dataloggerID, err := qtx.CreateDatalogger(ctx, n)
	if err != nil {
		return a, err
	}

	key, err := qtx.CreateDataloggerHash(ctx, dataloggerID)
	if err != nil {
		return a, err
	}

	if err := qtx.CreateDataloggerPreview(ctx, dataloggerID); err != nil {
		return a, err
	}

	dl, err := qtx.GetOneDatalogger(ctx, dataloggerID)
	if err != nil {
		return a, err
	}

	if err := tx.Commit(); err != nil {
		return a, err
	}

	dk := model.DataloggerWithKey{
		Datalogger: &dl,
		Key:        key,
	}

	return dk, nil
}

func (s dataloggerStore) CycleDataLoggerKey(ctx context.Context, u model.Datalogger) (model.DataloggerWithKey, error) {
	var a model.DataloggerWithKey
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return a, err
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			log.Print(err.Error())
		}
	}()

	qtx := s.q.WithTx(tx)

	key, err := qtx.UpdateDataloggerHash(ctx, u.ID)
	if err != nil {
		return a, err
	}

	if err := qtx.UpdateDataloggerUpdater(ctx, u); err != nil {
		return a, err
	}

	dl, err := qtx.GetOneDatalogger(ctx, u.ID)
	if err != nil {
		return a, err
	}

	if err := tx.Commit(); err != nil {
		return a, err
	}

	dk := model.DataloggerWithKey{
		Datalogger: &dl,
		Key:        key,
	}

	return dk, nil
}

func (s dataloggerStore) GetOneDataLogger(ctx context.Context, dataloggerID uuid.UUID) (model.Datalogger, error) {
	return s.q.GetOneDatalogger(ctx, dataloggerID)
}

func (s dataloggerStore) UpdateDataLogger(ctx context.Context, u model.Datalogger) (model.Datalogger, error) {
	var a model.Datalogger
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return a, err
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			log.Print(err.Error())
		}
	}()

	qtx := s.q.WithTx(tx)

	if err := qtx.UpdateDatalogger(ctx, u); err != nil {
		return a, err
	}

	dlUpdated, err := qtx.GetOneDatalogger(ctx, u.ID)
	if err != nil {
		return a, err
	}

	if err := tx.Commit(); err != nil {
		return a, err
	}

	return dlUpdated, nil
}

func (s dataloggerStore) DeleteDataLogger(ctx context.Context, d model.Datalogger) error {
	return s.q.DeleteDatalogger(ctx, d)
}

func (s dataloggerStore) GetDataLoggerPreview(ctx context.Context, dlID uuid.UUID) (model.DataloggerPreview, error) {
	return s.q.GetDataloggerPreview(ctx, dlID)
}
