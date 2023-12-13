package service

import (
	"context"

	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
)

type DataloggerService interface {
	GetDataloggerModelName(ctx context.Context, modelID uuid.UUID) (string, error)
	ListProjectDataloggers(ctx context.Context, projectID uuid.UUID) ([]model.Datalogger, error)
	ListAllDataloggers(ctx context.Context) ([]model.Datalogger, error)
	ListDataloggerSlugs(ctx context.Context) ([]string, error)
	GetDataloggerIsActive(ctx context.Context, modelName, sn string) (bool, error)
	VerifyDataloggerExists(ctx context.Context, dlID uuid.UUID) error
	CreateDatalogger(ctx context.Context, n model.Datalogger) (model.DataloggerWithKey, error)
	CycleDataloggerKey(ctx context.Context, u model.Datalogger) (model.DataloggerWithKey, error)
	GetOneDatalogger(ctx context.Context, dataloggerID uuid.UUID) (model.Datalogger, error)
	UpdateDatalogger(ctx context.Context, u model.Datalogger) (model.Datalogger, error)
	DeleteDatalogger(ctx context.Context, d model.Datalogger) error
	GetDataloggerTablePreview(ctx context.Context, dataloggerTableID uuid.UUID) (model.DataloggerPreview, error)
	ResetDataloggerTableName(ctx context.Context, dataloggerTableID uuid.UUID) error
	GetOrCreateDataloggerTable(ctx context.Context, dataloggerID uuid.UUID, tableName string) (uuid.UUID, error)
	DeleteDataloggerTable(ctx context.Context, dataloggerTableID uuid.UUID) error
}

type dataloggerService struct {
	db *model.Database
	*model.Queries
}

func NewDataloggerService(db *model.Database, q *model.Queries) *dataloggerService {
	return &dataloggerService{db, q}
}

func (s dataloggerService) CreateDatalogger(ctx context.Context, n model.Datalogger) (model.DataloggerWithKey, error) {
	var a model.DataloggerWithKey
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return a, err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)

	dataloggerID, err := qtx.CreateDatalogger(ctx, n)
	if err != nil {
		return a, err
	}

	key, err := qtx.CreateDataloggerHash(ctx, dataloggerID)
	if err != nil {
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
		Datalogger: dl,
		Key:        key,
	}

	return dk, nil
}

func (s dataloggerService) CycleDataloggerKey(ctx context.Context, u model.Datalogger) (model.DataloggerWithKey, error) {
	var a model.DataloggerWithKey
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return a, err
	}
	defer model.TxDo(tx.Rollback)

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
		Datalogger: dl,
		Key:        key,
	}

	return dk, nil
}

func (s dataloggerService) UpdateDatalogger(ctx context.Context, u model.Datalogger) (model.Datalogger, error) {
	var a model.Datalogger
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return a, err
	}
	defer model.TxDo(tx.Rollback)

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
