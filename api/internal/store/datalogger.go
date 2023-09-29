package store

import (
	"context"
	"log"

	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
)

type DataloggerStore interface {
	GetDataloggerModelName(ctx context.Context, modelID uuid.UUID) (string, error)
	ListProjectDataloggers(ctx context.Context, projectID uuid.UUID) ([]model.Datalogger, error)
	ListAllDataloggers(ctx context.Context) ([]model.Datalogger, error)
	GetDataloggerIsActive(ctx context.Context, modelName, sn string) (bool, error)
	VerifyDataloggerExists(ctx context.Context, dlID uuid.UUID) error
	CreateDatalogger(ctx context.Context, n model.Datalogger) (model.DataloggerWithKey, error)
	CycleDataloggerKey(ctx context.Context, u model.Datalogger) (model.DataloggerWithKey, error)
	GetOneDatalogger(ctx context.Context, dataloggerID uuid.UUID) (model.Datalogger, error)
	UpdateDatalogger(ctx context.Context, u model.Datalogger) (model.Datalogger, error)
	DeleteDatalogger(ctx context.Context, d model.Datalogger) error
	GetDataloggerPreview(ctx context.Context, dlID uuid.UUID) (model.DataloggerPreview, error)
	CreateUniqueSlugDatalogger(ctx context.Context, dataloggerName string) (string, error)
}

type dataloggerStore struct {
	db *model.Database
	*model.Queries
}

func NewDataloggerStore(db *model.Database, q *model.Queries) *dataloggerStore {
	return &dataloggerStore{db, q}
}

func (s dataloggerStore) CreateDatalogger(ctx context.Context, n model.Datalogger) (model.DataloggerWithKey, error) {
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

	qtx := s.WithTx(tx)

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

func (s dataloggerStore) CycleDataloggerKey(ctx context.Context, u model.Datalogger) (model.DataloggerWithKey, error) {
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

	qtx := s.WithTx(tx)

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

func (s dataloggerStore) UpdateDatalogger(ctx context.Context, u model.Datalogger) (model.Datalogger, error) {
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

	qtx := s.WithTx(tx)

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
