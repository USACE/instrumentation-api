package service

import (
	"context"
	"net/http"

	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
)

type TimeseriesCwmsService interface {
	ListTimeseriesCwms(ctx context.Context, instrumentID uuid.UUID) ([]model.TimeseriesCwms, error)
	CreateTimeseriesCwmsBatch(ctx context.Context, instrumentID uuid.UUID, tcc []model.TimeseriesCwms) ([]model.TimeseriesCwms, error)
	UpdateTimeseriesCwms(ctx context.Context, tsCwms model.TimeseriesCwms) error
}

type timeseriesCwmsService struct {
	cwmsClient  *http.Client
	cwmsDataUrl string
	db          *model.Database
	*model.Queries
}

func NewTimeseriesCwmsService(cwmsClient *http.Client, cwmsDataUrl string, db *model.Database, q *model.Queries) *timeseriesCwmsService {
	return &timeseriesCwmsService{cwmsClient, cwmsDataUrl, db, q}
}

func (s timeseriesCwmsService) CreateTimeseriesCwmsBatch(ctx context.Context, instrumentID uuid.UUID, tcc []model.TimeseriesCwms) ([]model.TimeseriesCwms, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return tcc, err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)

	for idx := range tcc {
		tcc[idx].Type = model.CwmsTimeseriesType
		tcc[idx].InstrumentID = instrumentID
		tsNew, err := qtx.CreateTimeseries(ctx, tcc[idx].Timeseries)
		if err != nil {
			return tcc, err
		}
		tcc[idx].Timeseries = tsNew
		if err := qtx.CreateTimeseriesCwms(ctx, tcc[idx]); err != nil {
			return tcc, err
		}
	}

	if err := tx.Commit(); err != nil {
		return tcc, err
	}

	return tcc, nil
}

func (s timeseriesCwmsService) UpdateTimeseriesCwms(ctx context.Context, tsCwms model.TimeseriesCwms) error {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)

	if _, err := qtx.UpdateTimeseries(ctx, tsCwms.Timeseries); err != nil {
		return err
	}

	if err := qtx.UpdateTimeseriesCwms(ctx, tsCwms); err != nil {
		return err
	}

	return tx.Commit()
}
